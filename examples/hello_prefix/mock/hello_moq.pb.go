// Code generated by protoc-gen-moq. DO NOT EDIT.
// (See github.com/matryer/moq for original template)

package hello

import (
	context "context"
	hello "github.com/alexandrevilain/protoc-gen-moq/examples/hello"
	grpc "google.golang.org/grpc"
	sync "sync"
)

// Ensure, that HelloServiceClientMock does implement hello.HelloServiceClient.
// If this is not the case, regenerate this file with moq.
var _ hello.HelloServiceClient = &HelloServiceClientMock{}

// HelloServiceClientMock is a mock implementation of hello.HelloServiceClient.
//
// 	func TestSomethingThatUseshello.HelloServiceClient(t *testing.T) {
//
// 		// make and configure a mocked hello.HelloServiceClient
// 		mockedhello.HelloServiceClient := &HelloServiceClientMock{
// 			SayHelloFunc: func(ctx context.Context, in *hello.SayHelloRequest, opts ...grpc.CallOption) (*hello.SayHelloResponse, error) {
// 				panic("mock out the SayHello method")
// 			},
// 		}
//
// 		// use mockedhello.HelloServiceClient in code that requires hello.HelloServiceClient
// 		// and then make assertions.
//
// 	}
type HelloServiceClientMock struct {
	// SayHelloFunc mocks the SayHello method.
	SayHelloFunc func(ctx context.Context, in *hello.SayHelloRequest, opts ...grpc.CallOption) (*hello.SayHelloResponse, error)

	// calls tracks calls to the methods.
	calls struct {
		// SayHello holds details about calls to the SayHello method.
		SayHello []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *hello.SayHelloRequest
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
	}
	lockSayHello sync.RWMutex
}

// SayHello calls SayHelloFunc.
func (mock *HelloServiceClientMock) SayHello(ctx context.Context, in *hello.SayHelloRequest, opts ...grpc.CallOption) (*hello.SayHelloResponse, error) {
	if mock.SayHelloFunc == nil {
		panic("HelloServiceClientMock.SayHelloFunc: method is nil but hello.HelloServiceClient.SayHello was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *hello.SayHelloRequest
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockSayHello.Lock()
	mock.calls.SayHello = append(mock.calls.SayHello, callInfo)
	mock.lockSayHello.Unlock()
	return mock.SayHelloFunc(ctx, in, opts...)
}

// SayHelloCalls gets all the calls that were made to SayHello.
// Check the length with:
//     len(mockedhello.HelloServiceClient.SayHelloCalls())
func (mock *HelloServiceClientMock) SayHelloCalls() []struct {
	Ctx  context.Context
	In   *hello.SayHelloRequest
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *hello.SayHelloRequest
		Opts []grpc.CallOption
	}
	mock.lockSayHello.RLock()
	calls = mock.calls.SayHello
	mock.lockSayHello.RUnlock()
	return calls
}
