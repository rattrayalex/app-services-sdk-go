// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package registrymgmtclient

import (
	_context "context"
	_nethttp "net/http"
	"sync"
)

// Ensure, that ErrorsApiMock does implement ErrorsApi.
// If this is not the case, regenerate this file with moq.
var _ ErrorsApi = &ErrorsApiMock{}

// ErrorsApiMock is a mock implementation of ErrorsApi.
//
// 	func TestSomethingThatUsesErrorsApi(t *testing.T) {
//
// 		// make and configure a mocked ErrorsApi
// 		mockedErrorsApi := &ErrorsApiMock{
// 			GetErrorFunc: func(ctx _context.Context, id int32) ApiGetErrorRequest {
// 				panic("mock out the GetError method")
// 			},
// 			GetErrorExecuteFunc: func(r ApiGetErrorRequest) (Error, *_nethttp.Response, error) {
// 				panic("mock out the GetErrorExecute method")
// 			},
// 			GetErrorsFunc: func(ctx _context.Context) ApiGetErrorsRequest {
// 				panic("mock out the GetErrors method")
// 			},
// 			GetErrorsExecuteFunc: func(r ApiGetErrorsRequest) (ErrorList, *_nethttp.Response, error) {
// 				panic("mock out the GetErrorsExecute method")
// 			},
// 		}
//
// 		// use mockedErrorsApi in code that requires ErrorsApi
// 		// and then make assertions.
//
// 	}
type ErrorsApiMock struct {
	// GetErrorFunc mocks the GetError method.
	GetErrorFunc func(ctx _context.Context, id int32) ApiGetErrorRequest

	// GetErrorExecuteFunc mocks the GetErrorExecute method.
	GetErrorExecuteFunc func(r ApiGetErrorRequest) (Error, *_nethttp.Response, error)

	// GetErrorsFunc mocks the GetErrors method.
	GetErrorsFunc func(ctx _context.Context) ApiGetErrorsRequest

	// GetErrorsExecuteFunc mocks the GetErrorsExecute method.
	GetErrorsExecuteFunc func(r ApiGetErrorsRequest) (ErrorList, *_nethttp.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetError holds details about calls to the GetError method.
		GetError []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// ID is the id argument value.
			ID int32
		}
		// GetErrorExecute holds details about calls to the GetErrorExecute method.
		GetErrorExecute []struct {
			// R is the r argument value.
			R ApiGetErrorRequest
		}
		// GetErrors holds details about calls to the GetErrors method.
		GetErrors []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// GetErrorsExecute holds details about calls to the GetErrorsExecute method.
		GetErrorsExecute []struct {
			// R is the r argument value.
			R ApiGetErrorsRequest
		}
	}
	lockGetError         sync.RWMutex
	lockGetErrorExecute  sync.RWMutex
	lockGetErrors        sync.RWMutex
	lockGetErrorsExecute sync.RWMutex
}

// GetError calls GetErrorFunc.
func (mock *ErrorsApiMock) GetError(ctx _context.Context, id int32) ApiGetErrorRequest {
	if mock.GetErrorFunc == nil {
		panic("ErrorsApiMock.GetErrorFunc: method is nil but ErrorsApi.GetError was just called")
	}
	callInfo := struct {
		Ctx _context.Context
		ID  int32
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetError.Lock()
	mock.calls.GetError = append(mock.calls.GetError, callInfo)
	mock.lockGetError.Unlock()
	return mock.GetErrorFunc(ctx, id)
}

// GetErrorCalls gets all the calls that were made to GetError.
// Check the length with:
//     len(mockedErrorsApi.GetErrorCalls())
func (mock *ErrorsApiMock) GetErrorCalls() []struct {
	Ctx _context.Context
	ID  int32
} {
	var calls []struct {
		Ctx _context.Context
		ID  int32
	}
	mock.lockGetError.RLock()
	calls = mock.calls.GetError
	mock.lockGetError.RUnlock()
	return calls
}

// GetErrorExecute calls GetErrorExecuteFunc.
func (mock *ErrorsApiMock) GetErrorExecute(r ApiGetErrorRequest) (Error, *_nethttp.Response, error) {
	if mock.GetErrorExecuteFunc == nil {
		panic("ErrorsApiMock.GetErrorExecuteFunc: method is nil but ErrorsApi.GetErrorExecute was just called")
	}
	callInfo := struct {
		R ApiGetErrorRequest
	}{
		R: r,
	}
	mock.lockGetErrorExecute.Lock()
	mock.calls.GetErrorExecute = append(mock.calls.GetErrorExecute, callInfo)
	mock.lockGetErrorExecute.Unlock()
	return mock.GetErrorExecuteFunc(r)
}

// GetErrorExecuteCalls gets all the calls that were made to GetErrorExecute.
// Check the length with:
//     len(mockedErrorsApi.GetErrorExecuteCalls())
func (mock *ErrorsApiMock) GetErrorExecuteCalls() []struct {
	R ApiGetErrorRequest
} {
	var calls []struct {
		R ApiGetErrorRequest
	}
	mock.lockGetErrorExecute.RLock()
	calls = mock.calls.GetErrorExecute
	mock.lockGetErrorExecute.RUnlock()
	return calls
}

// GetErrors calls GetErrorsFunc.
func (mock *ErrorsApiMock) GetErrors(ctx _context.Context) ApiGetErrorsRequest {
	if mock.GetErrorsFunc == nil {
		panic("ErrorsApiMock.GetErrorsFunc: method is nil but ErrorsApi.GetErrors was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetErrors.Lock()
	mock.calls.GetErrors = append(mock.calls.GetErrors, callInfo)
	mock.lockGetErrors.Unlock()
	return mock.GetErrorsFunc(ctx)
}

// GetErrorsCalls gets all the calls that were made to GetErrors.
// Check the length with:
//     len(mockedErrorsApi.GetErrorsCalls())
func (mock *ErrorsApiMock) GetErrorsCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockGetErrors.RLock()
	calls = mock.calls.GetErrors
	mock.lockGetErrors.RUnlock()
	return calls
}

// GetErrorsExecute calls GetErrorsExecuteFunc.
func (mock *ErrorsApiMock) GetErrorsExecute(r ApiGetErrorsRequest) (ErrorList, *_nethttp.Response, error) {
	if mock.GetErrorsExecuteFunc == nil {
		panic("ErrorsApiMock.GetErrorsExecuteFunc: method is nil but ErrorsApi.GetErrorsExecute was just called")
	}
	callInfo := struct {
		R ApiGetErrorsRequest
	}{
		R: r,
	}
	mock.lockGetErrorsExecute.Lock()
	mock.calls.GetErrorsExecute = append(mock.calls.GetErrorsExecute, callInfo)
	mock.lockGetErrorsExecute.Unlock()
	return mock.GetErrorsExecuteFunc(r)
}

// GetErrorsExecuteCalls gets all the calls that were made to GetErrorsExecute.
// Check the length with:
//     len(mockedErrorsApi.GetErrorsExecuteCalls())
func (mock *ErrorsApiMock) GetErrorsExecuteCalls() []struct {
	R ApiGetErrorsRequest
} {
	var calls []struct {
		R ApiGetErrorsRequest
	}
	mock.lockGetErrorsExecute.RLock()
	calls = mock.calls.GetErrorsExecute
	mock.lockGetErrorsExecute.RUnlock()
	return calls
}