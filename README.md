# CBAPIClient

CBAPIClient is an opinionated **Go** HTTP client that provides built-in support for the following common InVision design patterns:

* Circuit Breakers
* New Relic Transactions in Context
* Logrus logger in Context

## Usage

Go straight to the [examples](#examples)

### Setting Defaults

It is recommended that you initialize `cbapiclient` with package-level defaults when your application bootstraps.  
To do this, use the `Defaults` struct in the `SetDefaults` function.  The functions assigned in `Defaults` match 
the function signatures provided by [jelly](https://github.com/InVisionApp/jelly), but any function that returns 
the requested item with a boolean will work.  In our examples, we use the 
[jelly](https://github.com/InVisionApp/jelly) functions.

```go
type Defaults struct {
	// ServiceName is the name of the calling service
	ServiceName string

	// ContextLoggerProviderFunc is a function that provides
	// a logger from the current context.  If this function
	// is not set, the client will create a new logger for
	// the Request.
	ContextLoggerProviderFunc func(ctx context.Context) (*logrus.Entry, bool)

	// RequestIDProviderFunc is a function that provides the
	// parent Request id used in tracing the caller's Request.
	// If this function is not set, the client will generate
	// a new UUID for the Request id.
	RequestIDProviderFunc func(ctx context.Context) (string, bool)
}
```

It is strongly recommended that you provide these defaults.  However, if they are not provided, `cbapiclient` 
will provide its own defaults:

* `ServiceName` - if the service name is not provided, `cbapiclient` will look for an environment variable 
called `SERVICE_NAME`.  If that variable doesnt exist, `cbapiclient` will fall back to `HOSTNAME`
* `ContextLoggerProviderFunc` - if this function is not provided, `cbapiclient` will created a new `logrus.Entry` 
to use for the request
* `RequestIDProviderFunc` - if this function is not provided, `cbapiclient` will generate a new UUID to use for 
the request

Typical usage, with [jelly](https://github.com/InVisionApp/jelly) functions:

```go
package main

import (
	"fmt"
	
	cbapi "github.com/InVisionApp/cbapiclient"
	"github.com/InVisionApp/jelly"
)

const serviceName = "my-service"

func main() {
	cbapi.SetDefaults(&cbapi.Defaults{
		ServiceName: serviceName,
		NewRelicTransactionProviderFunc: jelly.GetNewRelicTransaction,
		ContextLoggerProviderFunc: jelly.GetContextLogger,
		RequestIDProviderFunc: jelly.GetRequestID,
	})
}
```

### Making Requests

#### Using Circuit Breakers

`cbapiclient` does not force you to use a circuit breaker, or, if you do, which library to use.  `cbapiclient` 
is built to implicitly support the [Sony GoBreaker](https://github.com/sony/gobreaker) library.  The circuit 
breaker implementation must conform to the following interface:

```go
type CircuitBreakerPrototype interface {
	Execute(func() (interface{}, error)) (interface{}, error)
}
```

To set the circuit breaker, use the `SetCircuitBreaker` function.

#### Content Type

By default, `cbapiclient` sets the `Content-Type` header to `application/json`.  You may override this header if 
you are sending a different content type by using the `SetContentType` function.

#### Response Payloads

There are two ways to access the response payload from `cbapiclient`.  If you want to access the raw bytes 
returned in the response, or if you are not expecting a JSON response payload, you can use the `RawResponse`
function to access the response bytes.

The more common way to access the payload is to use one of the `Saturate` functions.  You provide `cbapiclient` 
with an empty struct pointer, and it will be saturated when the response returns.  Each works slightly 
differently depending on what you're trying to do:

* `WillSaturate`
    * Use this function to saturate the struct you expect when the request succeeds.  **A request is considered 
    successful if the response code is between 200 and 299, inclusive.**
* `WillSaturateOnError`
	* Use this function to saturate the struct you expect when the request fails.  **A request is considered a 
	failure if the response code is below 200 or above 299**.  Also note that this struct will only be saturated 
	if the actual response is an error.  **`cbapiclient` _will not saturate any response if the error originated 
	at the caller_**.
* `WillSaturateWithStatusCode`
	* Use this function to saturate the struct you expect when a specific status code is returned.  
	Structs provided here take precedence over `WillSaturate` and `WillSaturateOnError`.  For example, if a 
	specific response is expected for a `401` error, and a struct is provided for `401`, then that struct will 
	be saturated when the response is `401`, regardless what was provided in `WillSaturateOnError`	 
	
#### Convenience Functions	

`cbapiclient` provides the following functions for convenience:

* `Get` - perform an HTTP GET request with no outgoing payload
* `Post` - perform an HTTP POST request with an outgoing payload
* `Put` - perform an HTTP PUT request with an outgoing payload
* `Patch` - perform an HTTP PATCH request with an outgoing payload
* `Delete` - perform an HTTP DELETE request with no outgoing payload

### Request/Response Customization

#### Headers

The following headers are set for every request:

* `Request-ID`
* `Content-Type`
* `User-Agent` - the user agent is in the form `{{service name}}-{{namespace}}-{{tenancy}}`
* `Calling-Service`

Additionally, the `Content-Length` header is set for all requests with an outgoing payload (`POST`, `PUT`, `PATCH`)

You can set additional headers by using the `SetHeader` function:

```go
c.SetHeader("X-Forwarded-For", "127.0.0.1")
```

### <a name="examples">Examples</a>

#### Get with known payload
```go
package main

import (
	"context"
	"log"

	cbapi "github.com/InVisionApp/cbapiclient"
)

type ResponsePayload struct {
	Foo  string `json:"foo"`
	Fizz string `json:"fizz"`
}

func main() {
	ctx := context.Background()

	// make you a client
	c, err := cbapi.NewClient("http://localhost:8080/foo/bar")
	if err != nil {
		log.Fatalln(err.Error())
	}

	// make empty struct pointer
	payload := &ResponsePayload{}
	
	// we will saturate the response with this GET
	c.WillSaturate(payload)
	statusCode, err := c.Get(ctx)

	log.Println(statusCode)
	log.Printf("%+v", payload)
}

```

#### Get with raw bytes
```go
package main

import (
	"context"
	"log"

	cbapi "github.com/InVisionApp/cbapiclient"
)

func main() {
	ctx := context.Background()

	// make you a client
	c, err := cbapi.NewClient("http://localhost:8080/foo/bar/xml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	
	// xml? it could happen ;)
	c.SetContentType("application/xml")
	c.KeepRawResponse()

	// run the request 
	statusCode, err := c.Get(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	
	// get the raw byte slice 
	data := c.RawResponse()

	log.Println(statusCode)
	log.Printf("%+v", data)
}

```

#### Post with known payload
```go
package main

import (
	"context"
	"log"

	cbapi "github.com/InVisionApp/cbapiclient"
)

type Payload struct {
	Foo  string `json:"foo"`
	Fizz string `json:"fizz"`
}

type ResponsePayload struct {
	Bar string `json:"bar"`
	Baz string `json:"baz"`
}

func main() {
	payload := &Payload{
		Foo: "this is foo",
		Fizz: "this is fizz",
	}
	ctx := context.Background()

	// make you a client
	c, err := cbapi.NewClient("http://localhost:8080/foo/bar")
	if err != nil {
		log.Fatalln(err.Error())
	}

	// make the empty struct pointer
	response := &ResponsePayload{}
	
	// we will saturate the response with this post
	c.WillSaturate(response)
	statusCode, err := c.Post(ctx, payload)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(statusCode)
	log.Printf("%+v", response)
}

```

#### Get with differing responses
```go
package main

import (
	"context"
	"log"

	cbapi "github.com/InVisionApp/cbapiclient"
)

type SuccessPayload struct {
	Foo  string `json:"foo"`
	Fizz string `json:"fizz"`
}

type ErrorPayload struct {
	Reason  string `json:"reason"`
	Code int `json:"code"`
}

type ForbiddenPayload struct {
	Domain  string `json:"domain"`
	UserID string `json:"userID"`
}

func main() {
	ctx := context.Background()

	// make you a client
	c, err := cbapi.NewClient("http://localhost:8080/foo/bar")
	if err != nil {
		log.Fatalln(err.Error())
	}

	// make empty struct pointers
	payload := &SuccessPayload{}
	errorPayload := &ErrorPayload{}
	forbiddenPayload := &ForbiddenPayload{}
	
	// this gets saturated on error
	c.WillSaturateOnError(errorPayload)
	
	// this gets saturated on 403
	c.WillSaturateWithStatusCode(403, forbiddenPayload)
	
	// we will saturate the response with this GET
	c.WillSaturate(payload)
	statusCode, err := c.Get(ctx)

	log.Println(statusCode)
	
	if statusCode == 403 {
		log.Printf("%+v", forbiddenPayload)
	} else if c.StatusCodeIsError() {
		log.Printf("%+v", errorPayload)
	} else {
		log.Printf("%+v", payload)
	}
}

```

## Future enhancements

* Add a `statsd` provider using the same convention established in `Defaults`
