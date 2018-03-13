// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"
	"time"

	"github.com/InVisionApp/cbapiclient"
)

type FakeIClient struct {
	DeleteStub        func(ctx context.Context) (int, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		ctx context.Context
	}
	deleteReturns struct {
		result1 int
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	DurationStub        func() time.Duration
	durationMutex       sync.RWMutex
	durationArgsForCall []struct{}
	durationReturns     struct {
		result1 time.Duration
	}
	durationReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	DoStub        func(ctx context.Context, method string, payload interface{}) (int, error)
	doMutex       sync.RWMutex
	doArgsForCall []struct {
		ctx     context.Context
		method  string
		payload interface{}
	}
	doReturns struct {
		result1 int
		result2 error
	}
	doReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	GetStub        func(ctx context.Context) (int, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		ctx context.Context
	}
	getReturns struct {
		result1 int
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	KeepRawResponseStub        func()
	keepRawResponseMutex       sync.RWMutex
	keepRawResponseArgsForCall []struct{}
	PostStub                   func(ctx context.Context, payload interface{}) (int, error)
	postMutex                  sync.RWMutex
	postArgsForCall            []struct {
		ctx     context.Context
		payload interface{}
	}
	postReturns struct {
		result1 int
		result2 error
	}
	postReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	PutStub        func(ctx context.Context, payload interface{}) (int, error)
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		ctx     context.Context
		payload interface{}
	}
	putReturns struct {
		result1 int
		result2 error
	}
	putReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	PatchStub        func(ctx context.Context, payload interface{}) (int, error)
	patchMutex       sync.RWMutex
	patchArgsForCall []struct {
		ctx     context.Context
		payload interface{}
	}
	patchReturns struct {
		result1 int
		result2 error
	}
	patchReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	RawResponseStub        func() []byte
	rawResponseMutex       sync.RWMutex
	rawResponseArgsForCall []struct{}
	rawResponseReturns     struct {
		result1 []byte
	}
	rawResponseReturnsOnCall map[int]struct {
		result1 []byte
	}
	SetCircuitBreakerStub        func(cb cbapiclient.CircuitBreakerPrototype)
	setCircuitBreakerMutex       sync.RWMutex
	setCircuitBreakerArgsForCall []struct {
		cb cbapiclient.CircuitBreakerPrototype
	}
	SetStatsdClientWithTagsStub        func(sd cbapiclient.StatsdClientPrototype, tags []string)
	setStatsdClientWithTagsMutex       sync.RWMutex
	setStatsdClientWithTagsArgsForCall []struct {
		sd   cbapiclient.StatsdClientPrototype
		tags []string
	}
	SetContentTypeStub        func(ct string)
	setContentTypeMutex       sync.RWMutex
	setContentTypeArgsForCall []struct {
		ct string
	}
	SetHeaderStub        func(key string, value string)
	setHeaderMutex       sync.RWMutex
	setHeaderArgsForCall []struct {
		key   string
		value string
	}
	SetNRTxnNameStub        func(name string)
	setNRTxnNameMutex       sync.RWMutex
	setNRTxnNameArgsForCall []struct {
		name string
	}
	SetTimeoutMSStub        func(timeout time.Duration)
	setTimeoutMSMutex       sync.RWMutex
	setTimeoutMSArgsForCall []struct {
		timeout time.Duration
	}
	StatusCodeIsErrorStub        func() bool
	statusCodeIsErrorMutex       sync.RWMutex
	statusCodeIsErrorArgsForCall []struct{}
	statusCodeIsErrorReturns     struct {
		result1 bool
	}
	statusCodeIsErrorReturnsOnCall map[int]struct {
		result1 bool
	}
	WillSaturateStub        func(proto interface{})
	willSaturateMutex       sync.RWMutex
	willSaturateArgsForCall []struct {
		proto interface{}
	}
	WillSaturateOnErrorStub        func(proto interface{})
	willSaturateOnErrorMutex       sync.RWMutex
	willSaturateOnErrorArgsForCall []struct {
		proto interface{}
	}
	WillSaturateWithStatusCodeStub        func(statusCode int, proto interface{})
	willSaturateWithStatusCodeMutex       sync.RWMutex
	willSaturateWithStatusCodeArgsForCall []struct {
		statusCode int
		proto      interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIClient) Delete(ctx context.Context) (int, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		ctx context.Context
	}{ctx})
	fake.recordInvocation("Delete", []interface{}{ctx})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(ctx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteReturns.result1, fake.deleteReturns.result2
}

func (fake *FakeIClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeIClient) DeleteArgsForCall(i int) context.Context {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].ctx
}

func (fake *FakeIClient) DeleteReturns(result1 int, result2 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) DeleteReturnsOnCall(i int, result1 int, result2 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) Duration() time.Duration {
	fake.durationMutex.Lock()
	ret, specificReturn := fake.durationReturnsOnCall[len(fake.durationArgsForCall)]
	fake.durationArgsForCall = append(fake.durationArgsForCall, struct{}{})
	fake.recordInvocation("Duration", []interface{}{})
	fake.durationMutex.Unlock()
	if fake.DurationStub != nil {
		return fake.DurationStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.durationReturns.result1
}

func (fake *FakeIClient) DurationCallCount() int {
	fake.durationMutex.RLock()
	defer fake.durationMutex.RUnlock()
	return len(fake.durationArgsForCall)
}

func (fake *FakeIClient) DurationReturns(result1 time.Duration) {
	fake.DurationStub = nil
	fake.durationReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeIClient) DurationReturnsOnCall(i int, result1 time.Duration) {
	fake.DurationStub = nil
	if fake.durationReturnsOnCall == nil {
		fake.durationReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.durationReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeIClient) Do(ctx context.Context, method string, payload interface{}) (int, error) {
	fake.doMutex.Lock()
	ret, specificReturn := fake.doReturnsOnCall[len(fake.doArgsForCall)]
	fake.doArgsForCall = append(fake.doArgsForCall, struct {
		ctx     context.Context
		method  string
		payload interface{}
	}{ctx, method, payload})
	fake.recordInvocation("Do", []interface{}{ctx, method, payload})
	fake.doMutex.Unlock()
	if fake.DoStub != nil {
		return fake.DoStub(ctx, method, payload)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.doReturns.result1, fake.doReturns.result2
}

func (fake *FakeIClient) DoCallCount() int {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return len(fake.doArgsForCall)
}

func (fake *FakeIClient) DoArgsForCall(i int) (context.Context, string, interface{}) {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return fake.doArgsForCall[i].ctx, fake.doArgsForCall[i].method, fake.doArgsForCall[i].payload
}

func (fake *FakeIClient) DoReturns(result1 int, result2 error) {
	fake.DoStub = nil
	fake.doReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) DoReturnsOnCall(i int, result1 int, result2 error) {
	fake.DoStub = nil
	if fake.doReturnsOnCall == nil {
		fake.doReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.doReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) Get(ctx context.Context) (int, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		ctx context.Context
	}{ctx})
	fake.recordInvocation("Get", []interface{}{ctx})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(ctx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getReturns.result1, fake.getReturns.result2
}

func (fake *FakeIClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeIClient) GetArgsForCall(i int) context.Context {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].ctx
}

func (fake *FakeIClient) GetReturns(result1 int, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) GetReturnsOnCall(i int, result1 int, result2 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) KeepRawResponse() {
	fake.keepRawResponseMutex.Lock()
	fake.keepRawResponseArgsForCall = append(fake.keepRawResponseArgsForCall, struct{}{})
	fake.recordInvocation("KeepRawResponse", []interface{}{})
	fake.keepRawResponseMutex.Unlock()
	if fake.KeepRawResponseStub != nil {
		fake.KeepRawResponseStub()
	}
}

func (fake *FakeIClient) KeepRawResponseCallCount() int {
	fake.keepRawResponseMutex.RLock()
	defer fake.keepRawResponseMutex.RUnlock()
	return len(fake.keepRawResponseArgsForCall)
}

func (fake *FakeIClient) Post(ctx context.Context, payload interface{}) (int, error) {
	fake.postMutex.Lock()
	ret, specificReturn := fake.postReturnsOnCall[len(fake.postArgsForCall)]
	fake.postArgsForCall = append(fake.postArgsForCall, struct {
		ctx     context.Context
		payload interface{}
	}{ctx, payload})
	fake.recordInvocation("Post", []interface{}{ctx, payload})
	fake.postMutex.Unlock()
	if fake.PostStub != nil {
		return fake.PostStub(ctx, payload)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.postReturns.result1, fake.postReturns.result2
}

func (fake *FakeIClient) PostCallCount() int {
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	return len(fake.postArgsForCall)
}

func (fake *FakeIClient) PostArgsForCall(i int) (context.Context, interface{}) {
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	return fake.postArgsForCall[i].ctx, fake.postArgsForCall[i].payload
}

func (fake *FakeIClient) PostReturns(result1 int, result2 error) {
	fake.PostStub = nil
	fake.postReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) PostReturnsOnCall(i int, result1 int, result2 error) {
	fake.PostStub = nil
	if fake.postReturnsOnCall == nil {
		fake.postReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.postReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) Put(ctx context.Context, payload interface{}) (int, error) {
	fake.putMutex.Lock()
	ret, specificReturn := fake.putReturnsOnCall[len(fake.putArgsForCall)]
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		ctx     context.Context
		payload interface{}
	}{ctx, payload})
	fake.recordInvocation("Put", []interface{}{ctx, payload})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(ctx, payload)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.putReturns.result1, fake.putReturns.result2
}

func (fake *FakeIClient) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeIClient) PutArgsForCall(i int) (context.Context, interface{}) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].ctx, fake.putArgsForCall[i].payload
}

func (fake *FakeIClient) PutReturns(result1 int, result2 error) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) PutReturnsOnCall(i int, result1 int, result2 error) {
	fake.PutStub = nil
	if fake.putReturnsOnCall == nil {
		fake.putReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.putReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) Patch(ctx context.Context, payload interface{}) (int, error) {
	fake.patchMutex.Lock()
	ret, specificReturn := fake.patchReturnsOnCall[len(fake.patchArgsForCall)]
	fake.patchArgsForCall = append(fake.patchArgsForCall, struct {
		ctx     context.Context
		payload interface{}
	}{ctx, payload})
	fake.recordInvocation("Patch", []interface{}{ctx, payload})
	fake.patchMutex.Unlock()
	if fake.PatchStub != nil {
		return fake.PatchStub(ctx, payload)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.patchReturns.result1, fake.patchReturns.result2
}

func (fake *FakeIClient) PatchCallCount() int {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	return len(fake.patchArgsForCall)
}

func (fake *FakeIClient) PatchArgsForCall(i int) (context.Context, interface{}) {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	return fake.patchArgsForCall[i].ctx, fake.patchArgsForCall[i].payload
}

func (fake *FakeIClient) PatchReturns(result1 int, result2 error) {
	fake.PatchStub = nil
	fake.patchReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) PatchReturnsOnCall(i int, result1 int, result2 error) {
	fake.PatchStub = nil
	if fake.patchReturnsOnCall == nil {
		fake.patchReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.patchReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) RawResponse() []byte {
	fake.rawResponseMutex.Lock()
	ret, specificReturn := fake.rawResponseReturnsOnCall[len(fake.rawResponseArgsForCall)]
	fake.rawResponseArgsForCall = append(fake.rawResponseArgsForCall, struct{}{})
	fake.recordInvocation("RawResponse", []interface{}{})
	fake.rawResponseMutex.Unlock()
	if fake.RawResponseStub != nil {
		return fake.RawResponseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.rawResponseReturns.result1
}

func (fake *FakeIClient) RawResponseCallCount() int {
	fake.rawResponseMutex.RLock()
	defer fake.rawResponseMutex.RUnlock()
	return len(fake.rawResponseArgsForCall)
}

func (fake *FakeIClient) RawResponseReturns(result1 []byte) {
	fake.RawResponseStub = nil
	fake.rawResponseReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeIClient) RawResponseReturnsOnCall(i int, result1 []byte) {
	fake.RawResponseStub = nil
	if fake.rawResponseReturnsOnCall == nil {
		fake.rawResponseReturnsOnCall = make(map[int]struct {
			result1 []byte
		})
	}
	fake.rawResponseReturnsOnCall[i] = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeIClient) SetCircuitBreaker(cb cbapiclient.CircuitBreakerPrototype) {
	fake.setCircuitBreakerMutex.Lock()
	fake.setCircuitBreakerArgsForCall = append(fake.setCircuitBreakerArgsForCall, struct {
		cb cbapiclient.CircuitBreakerPrototype
	}{cb})
	fake.recordInvocation("SetCircuitBreaker", []interface{}{cb})
	fake.setCircuitBreakerMutex.Unlock()
	if fake.SetCircuitBreakerStub != nil {
		fake.SetCircuitBreakerStub(cb)
	}
}

func (fake *FakeIClient) SetCircuitBreakerCallCount() int {
	fake.setCircuitBreakerMutex.RLock()
	defer fake.setCircuitBreakerMutex.RUnlock()
	return len(fake.setCircuitBreakerArgsForCall)
}

func (fake *FakeIClient) SetCircuitBreakerArgsForCall(i int) cbapiclient.CircuitBreakerPrototype {
	fake.setCircuitBreakerMutex.RLock()
	defer fake.setCircuitBreakerMutex.RUnlock()
	return fake.setCircuitBreakerArgsForCall[i].cb
}

func (fake *FakeIClient) SetStatsdClientWithTags(sd cbapiclient.StatsdClientPrototype, tags []string) {
	var tagsCopy []string
	if tags != nil {
		tagsCopy = make([]string, len(tags))
		copy(tagsCopy, tags)
	}
	fake.setStatsdClientWithTagsMutex.Lock()
	fake.setStatsdClientWithTagsArgsForCall = append(fake.setStatsdClientWithTagsArgsForCall, struct {
		sd   cbapiclient.StatsdClientPrototype
		tags []string
	}{sd, tagsCopy})
	fake.recordInvocation("SetStatsdClientWithTags", []interface{}{sd, tagsCopy})
	fake.setStatsdClientWithTagsMutex.Unlock()
	if fake.SetStatsdClientWithTagsStub != nil {
		fake.SetStatsdClientWithTagsStub(sd, tags)
	}
}

func (fake *FakeIClient) SetStatsdClientWithTagsCallCount() int {
	fake.setStatsdClientWithTagsMutex.RLock()
	defer fake.setStatsdClientWithTagsMutex.RUnlock()
	return len(fake.setStatsdClientWithTagsArgsForCall)
}

func (fake *FakeIClient) SetStatsdClientWithTagsArgsForCall(i int) (cbapiclient.StatsdClientPrototype, []string) {
	fake.setStatsdClientWithTagsMutex.RLock()
	defer fake.setStatsdClientWithTagsMutex.RUnlock()
	return fake.setStatsdClientWithTagsArgsForCall[i].sd, fake.setStatsdClientWithTagsArgsForCall[i].tags
}

func (fake *FakeIClient) SetContentType(ct string) {
	fake.setContentTypeMutex.Lock()
	fake.setContentTypeArgsForCall = append(fake.setContentTypeArgsForCall, struct {
		ct string
	}{ct})
	fake.recordInvocation("SetContentType", []interface{}{ct})
	fake.setContentTypeMutex.Unlock()
	if fake.SetContentTypeStub != nil {
		fake.SetContentTypeStub(ct)
	}
}

func (fake *FakeIClient) SetContentTypeCallCount() int {
	fake.setContentTypeMutex.RLock()
	defer fake.setContentTypeMutex.RUnlock()
	return len(fake.setContentTypeArgsForCall)
}

func (fake *FakeIClient) SetContentTypeArgsForCall(i int) string {
	fake.setContentTypeMutex.RLock()
	defer fake.setContentTypeMutex.RUnlock()
	return fake.setContentTypeArgsForCall[i].ct
}

func (fake *FakeIClient) SetHeader(key string, value string) {
	fake.setHeaderMutex.Lock()
	fake.setHeaderArgsForCall = append(fake.setHeaderArgsForCall, struct {
		key   string
		value string
	}{key, value})
	fake.recordInvocation("SetHeader", []interface{}{key, value})
	fake.setHeaderMutex.Unlock()
	if fake.SetHeaderStub != nil {
		fake.SetHeaderStub(key, value)
	}
}

func (fake *FakeIClient) SetHeaderCallCount() int {
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	return len(fake.setHeaderArgsForCall)
}

func (fake *FakeIClient) SetHeaderArgsForCall(i int) (string, string) {
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	return fake.setHeaderArgsForCall[i].key, fake.setHeaderArgsForCall[i].value
}

func (fake *FakeIClient) SetNRTxnName(name string) {
	fake.setNRTxnNameMutex.Lock()
	fake.setNRTxnNameArgsForCall = append(fake.setNRTxnNameArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("SetNRTxnName", []interface{}{name})
	fake.setNRTxnNameMutex.Unlock()
	if fake.SetNRTxnNameStub != nil {
		fake.SetNRTxnNameStub(name)
	}
}

func (fake *FakeIClient) SetNRTxnNameCallCount() int {
	fake.setNRTxnNameMutex.RLock()
	defer fake.setNRTxnNameMutex.RUnlock()
	return len(fake.setNRTxnNameArgsForCall)
}

func (fake *FakeIClient) SetNRTxnNameArgsForCall(i int) string {
	fake.setNRTxnNameMutex.RLock()
	defer fake.setNRTxnNameMutex.RUnlock()
	return fake.setNRTxnNameArgsForCall[i].name
}

func (fake *FakeIClient) SetTimeoutMS(timeout time.Duration) {
	fake.setTimeoutMSMutex.Lock()
	fake.setTimeoutMSArgsForCall = append(fake.setTimeoutMSArgsForCall, struct {
		timeout time.Duration
	}{timeout})
	fake.recordInvocation("SetTimeoutMS", []interface{}{timeout})
	fake.setTimeoutMSMutex.Unlock()
	if fake.SetTimeoutMSStub != nil {
		fake.SetTimeoutMSStub(timeout)
	}
}

func (fake *FakeIClient) SetTimeoutMSCallCount() int {
	fake.setTimeoutMSMutex.RLock()
	defer fake.setTimeoutMSMutex.RUnlock()
	return len(fake.setTimeoutMSArgsForCall)
}

func (fake *FakeIClient) SetTimeoutMSArgsForCall(i int) time.Duration {
	fake.setTimeoutMSMutex.RLock()
	defer fake.setTimeoutMSMutex.RUnlock()
	return fake.setTimeoutMSArgsForCall[i].timeout
}

func (fake *FakeIClient) StatusCodeIsError() bool {
	fake.statusCodeIsErrorMutex.Lock()
	ret, specificReturn := fake.statusCodeIsErrorReturnsOnCall[len(fake.statusCodeIsErrorArgsForCall)]
	fake.statusCodeIsErrorArgsForCall = append(fake.statusCodeIsErrorArgsForCall, struct{}{})
	fake.recordInvocation("StatusCodeIsError", []interface{}{})
	fake.statusCodeIsErrorMutex.Unlock()
	if fake.StatusCodeIsErrorStub != nil {
		return fake.StatusCodeIsErrorStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.statusCodeIsErrorReturns.result1
}

func (fake *FakeIClient) StatusCodeIsErrorCallCount() int {
	fake.statusCodeIsErrorMutex.RLock()
	defer fake.statusCodeIsErrorMutex.RUnlock()
	return len(fake.statusCodeIsErrorArgsForCall)
}

func (fake *FakeIClient) StatusCodeIsErrorReturns(result1 bool) {
	fake.StatusCodeIsErrorStub = nil
	fake.statusCodeIsErrorReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeIClient) StatusCodeIsErrorReturnsOnCall(i int, result1 bool) {
	fake.StatusCodeIsErrorStub = nil
	if fake.statusCodeIsErrorReturnsOnCall == nil {
		fake.statusCodeIsErrorReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.statusCodeIsErrorReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeIClient) WillSaturate(proto interface{}) {
	fake.willSaturateMutex.Lock()
	fake.willSaturateArgsForCall = append(fake.willSaturateArgsForCall, struct {
		proto interface{}
	}{proto})
	fake.recordInvocation("WillSaturate", []interface{}{proto})
	fake.willSaturateMutex.Unlock()
	if fake.WillSaturateStub != nil {
		fake.WillSaturateStub(proto)
	}
}

func (fake *FakeIClient) WillSaturateCallCount() int {
	fake.willSaturateMutex.RLock()
	defer fake.willSaturateMutex.RUnlock()
	return len(fake.willSaturateArgsForCall)
}

func (fake *FakeIClient) WillSaturateArgsForCall(i int) interface{} {
	fake.willSaturateMutex.RLock()
	defer fake.willSaturateMutex.RUnlock()
	return fake.willSaturateArgsForCall[i].proto
}

func (fake *FakeIClient) WillSaturateOnError(proto interface{}) {
	fake.willSaturateOnErrorMutex.Lock()
	fake.willSaturateOnErrorArgsForCall = append(fake.willSaturateOnErrorArgsForCall, struct {
		proto interface{}
	}{proto})
	fake.recordInvocation("WillSaturateOnError", []interface{}{proto})
	fake.willSaturateOnErrorMutex.Unlock()
	if fake.WillSaturateOnErrorStub != nil {
		fake.WillSaturateOnErrorStub(proto)
	}
}

func (fake *FakeIClient) WillSaturateOnErrorCallCount() int {
	fake.willSaturateOnErrorMutex.RLock()
	defer fake.willSaturateOnErrorMutex.RUnlock()
	return len(fake.willSaturateOnErrorArgsForCall)
}

func (fake *FakeIClient) WillSaturateOnErrorArgsForCall(i int) interface{} {
	fake.willSaturateOnErrorMutex.RLock()
	defer fake.willSaturateOnErrorMutex.RUnlock()
	return fake.willSaturateOnErrorArgsForCall[i].proto
}

func (fake *FakeIClient) WillSaturateWithStatusCode(statusCode int, proto interface{}) {
	fake.willSaturateWithStatusCodeMutex.Lock()
	fake.willSaturateWithStatusCodeArgsForCall = append(fake.willSaturateWithStatusCodeArgsForCall, struct {
		statusCode int
		proto      interface{}
	}{statusCode, proto})
	fake.recordInvocation("WillSaturateWithStatusCode", []interface{}{statusCode, proto})
	fake.willSaturateWithStatusCodeMutex.Unlock()
	if fake.WillSaturateWithStatusCodeStub != nil {
		fake.WillSaturateWithStatusCodeStub(statusCode, proto)
	}
}

func (fake *FakeIClient) WillSaturateWithStatusCodeCallCount() int {
	fake.willSaturateWithStatusCodeMutex.RLock()
	defer fake.willSaturateWithStatusCodeMutex.RUnlock()
	return len(fake.willSaturateWithStatusCodeArgsForCall)
}

func (fake *FakeIClient) WillSaturateWithStatusCodeArgsForCall(i int) (int, interface{}) {
	fake.willSaturateWithStatusCodeMutex.RLock()
	defer fake.willSaturateWithStatusCodeMutex.RUnlock()
	return fake.willSaturateWithStatusCodeArgsForCall[i].statusCode, fake.willSaturateWithStatusCodeArgsForCall[i].proto
}

func (fake *FakeIClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.durationMutex.RLock()
	defer fake.durationMutex.RUnlock()
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.keepRawResponseMutex.RLock()
	defer fake.keepRawResponseMutex.RUnlock()
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	fake.rawResponseMutex.RLock()
	defer fake.rawResponseMutex.RUnlock()
	fake.setCircuitBreakerMutex.RLock()
	defer fake.setCircuitBreakerMutex.RUnlock()
	fake.setStatsdClientWithTagsMutex.RLock()
	defer fake.setStatsdClientWithTagsMutex.RUnlock()
	fake.setContentTypeMutex.RLock()
	defer fake.setContentTypeMutex.RUnlock()
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	fake.setNRTxnNameMutex.RLock()
	defer fake.setNRTxnNameMutex.RUnlock()
	fake.setTimeoutMSMutex.RLock()
	defer fake.setTimeoutMSMutex.RUnlock()
	fake.statusCodeIsErrorMutex.RLock()
	defer fake.statusCodeIsErrorMutex.RUnlock()
	fake.willSaturateMutex.RLock()
	defer fake.willSaturateMutex.RUnlock()
	fake.willSaturateOnErrorMutex.RLock()
	defer fake.willSaturateOnErrorMutex.RUnlock()
	fake.willSaturateWithStatusCodeMutex.RLock()
	defer fake.willSaturateWithStatusCodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ cbapiclient.IClient = new(FakeIClient)
