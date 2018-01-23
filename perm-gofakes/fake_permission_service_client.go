// Code generated by counterfeiter. DO NOT EDIT.
package permgofakes

import (
	"sync"

	perm_go "code.cloudfoundry.org/perm-go"
	_ "github.com/gogo/protobuf/gogoproto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type FakePermissionServiceClient struct {
	HasPermissionStub        func(ctx context.Context, in *perm_go.HasPermissionRequest, opts ...grpc.CallOption) (*perm_go.HasPermissionResponse, error)
	hasPermissionMutex       sync.RWMutex
	hasPermissionArgsForCall []struct {
		ctx  context.Context
		in   *perm_go.HasPermissionRequest
		opts []grpc.CallOption
	}
	hasPermissionReturns struct {
		result1 *perm_go.HasPermissionResponse
		result2 error
	}
	hasPermissionReturnsOnCall map[int]struct {
		result1 *perm_go.HasPermissionResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePermissionServiceClient) HasPermission(ctx context.Context, in *perm_go.HasPermissionRequest, opts ...grpc.CallOption) (*perm_go.HasPermissionResponse, error) {
	fake.hasPermissionMutex.Lock()
	ret, specificReturn := fake.hasPermissionReturnsOnCall[len(fake.hasPermissionArgsForCall)]
	fake.hasPermissionArgsForCall = append(fake.hasPermissionArgsForCall, struct {
		ctx  context.Context
		in   *perm_go.HasPermissionRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("HasPermission", []interface{}{ctx, in, opts})
	fake.hasPermissionMutex.Unlock()
	if fake.HasPermissionStub != nil {
		return fake.HasPermissionStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.hasPermissionReturns.result1, fake.hasPermissionReturns.result2
}

func (fake *FakePermissionServiceClient) HasPermissionCallCount() int {
	fake.hasPermissionMutex.RLock()
	defer fake.hasPermissionMutex.RUnlock()
	return len(fake.hasPermissionArgsForCall)
}

func (fake *FakePermissionServiceClient) HasPermissionArgsForCall(i int) (context.Context, *perm_go.HasPermissionRequest, []grpc.CallOption) {
	fake.hasPermissionMutex.RLock()
	defer fake.hasPermissionMutex.RUnlock()
	return fake.hasPermissionArgsForCall[i].ctx, fake.hasPermissionArgsForCall[i].in, fake.hasPermissionArgsForCall[i].opts
}

func (fake *FakePermissionServiceClient) HasPermissionReturns(result1 *perm_go.HasPermissionResponse, result2 error) {
	fake.HasPermissionStub = nil
	fake.hasPermissionReturns = struct {
		result1 *perm_go.HasPermissionResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePermissionServiceClient) HasPermissionReturnsOnCall(i int, result1 *perm_go.HasPermissionResponse, result2 error) {
	fake.HasPermissionStub = nil
	if fake.hasPermissionReturnsOnCall == nil {
		fake.hasPermissionReturnsOnCall = make(map[int]struct {
			result1 *perm_go.HasPermissionResponse
			result2 error
		})
	}
	fake.hasPermissionReturnsOnCall[i] = struct {
		result1 *perm_go.HasPermissionResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePermissionServiceClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.hasPermissionMutex.RLock()
	defer fake.hasPermissionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePermissionServiceClient) recordInvocation(key string, args []interface{}) {
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

var _ perm_go.PermissionServiceClient = new(FakePermissionServiceClient)
