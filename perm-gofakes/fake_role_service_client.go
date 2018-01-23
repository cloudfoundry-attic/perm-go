// Code generated by counterfeiter. DO NOT EDIT.
package permgofakes

import (
	"sync"

	perm_go "code.cloudfoundry.org/perm-go"
	_ "github.com/gogo/protobuf/gogoproto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type FakeRoleServiceClient struct {
	CreateRoleStub        func(ctx context.Context, in *perm_go.CreateRoleRequest, opts ...grpc.CallOption) (*perm_go.CreateRoleResponse, error)
	createRoleMutex       sync.RWMutex
	createRoleArgsForCall []struct {
		ctx  context.Context
		in   *perm_go.CreateRoleRequest
		opts []grpc.CallOption
	}
	createRoleReturns struct {
		result1 *perm_go.CreateRoleResponse
		result2 error
	}
	createRoleReturnsOnCall map[int]struct {
		result1 *perm_go.CreateRoleResponse
		result2 error
	}
	GetRoleStub        func(ctx context.Context, in *perm_go.GetRoleRequest, opts ...grpc.CallOption) (*perm_go.GetRoleResponse, error)
	getRoleMutex       sync.RWMutex
	getRoleArgsForCall []struct {
		ctx  context.Context
		in   *perm_go.GetRoleRequest
		opts []grpc.CallOption
	}
	getRoleReturns struct {
		result1 *perm_go.GetRoleResponse
		result2 error
	}
	getRoleReturnsOnCall map[int]struct {
		result1 *perm_go.GetRoleResponse
		result2 error
	}
	DeleteRoleStub        func(ctx context.Context, in *perm_go.DeleteRoleRequest, opts ...grpc.CallOption) (*perm_go.DeleteRoleResponse, error)
	deleteRoleMutex       sync.RWMutex
	deleteRoleArgsForCall []struct {
		ctx  context.Context
		in   *perm_go.DeleteRoleRequest
		opts []grpc.CallOption
	}
	deleteRoleReturns struct {
		result1 *perm_go.DeleteRoleResponse
		result2 error
	}
	deleteRoleReturnsOnCall map[int]struct {
		result1 *perm_go.DeleteRoleResponse
		result2 error
	}
	AssignRoleStub        func(ctx context.Context, in *perm_go.AssignRoleRequest, opts ...grpc.CallOption) (*perm_go.AssignRoleResponse, error)
	assignRoleMutex       sync.RWMutex
	assignRoleArgsForCall []struct {
		ctx  context.Context
		in   *perm_go.AssignRoleRequest
		opts []grpc.CallOption
	}
	assignRoleReturns struct {
		result1 *perm_go.AssignRoleResponse
		result2 error
	}
	assignRoleReturnsOnCall map[int]struct {
		result1 *perm_go.AssignRoleResponse
		result2 error
	}
	UnassignRoleStub        func(ctx context.Context, in *perm_go.UnassignRoleRequest, opts ...grpc.CallOption) (*perm_go.UnassignRoleResponse, error)
	unassignRoleMutex       sync.RWMutex
	unassignRoleArgsForCall []struct {
		ctx  context.Context
		in   *perm_go.UnassignRoleRequest
		opts []grpc.CallOption
	}
	unassignRoleReturns struct {
		result1 *perm_go.UnassignRoleResponse
		result2 error
	}
	unassignRoleReturnsOnCall map[int]struct {
		result1 *perm_go.UnassignRoleResponse
		result2 error
	}
	HasRoleStub        func(ctx context.Context, in *perm_go.HasRoleRequest, opts ...grpc.CallOption) (*perm_go.HasRoleResponse, error)
	hasRoleMutex       sync.RWMutex
	hasRoleArgsForCall []struct {
		ctx  context.Context
		in   *perm_go.HasRoleRequest
		opts []grpc.CallOption
	}
	hasRoleReturns struct {
		result1 *perm_go.HasRoleResponse
		result2 error
	}
	hasRoleReturnsOnCall map[int]struct {
		result1 *perm_go.HasRoleResponse
		result2 error
	}
	ListActorRolesStub        func(ctx context.Context, in *perm_go.ListActorRolesRequest, opts ...grpc.CallOption) (*perm_go.ListActorRolesResponse, error)
	listActorRolesMutex       sync.RWMutex
	listActorRolesArgsForCall []struct {
		ctx  context.Context
		in   *perm_go.ListActorRolesRequest
		opts []grpc.CallOption
	}
	listActorRolesReturns struct {
		result1 *perm_go.ListActorRolesResponse
		result2 error
	}
	listActorRolesReturnsOnCall map[int]struct {
		result1 *perm_go.ListActorRolesResponse
		result2 error
	}
	ListRolePermissionsStub        func(ctx context.Context, in *perm_go.ListRolePermissionsRequest, opts ...grpc.CallOption) (*perm_go.ListRolePermissionsResponse, error)
	listRolePermissionsMutex       sync.RWMutex
	listRolePermissionsArgsForCall []struct {
		ctx  context.Context
		in   *perm_go.ListRolePermissionsRequest
		opts []grpc.CallOption
	}
	listRolePermissionsReturns struct {
		result1 *perm_go.ListRolePermissionsResponse
		result2 error
	}
	listRolePermissionsReturnsOnCall map[int]struct {
		result1 *perm_go.ListRolePermissionsResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRoleServiceClient) CreateRole(ctx context.Context, in *perm_go.CreateRoleRequest, opts ...grpc.CallOption) (*perm_go.CreateRoleResponse, error) {
	fake.createRoleMutex.Lock()
	ret, specificReturn := fake.createRoleReturnsOnCall[len(fake.createRoleArgsForCall)]
	fake.createRoleArgsForCall = append(fake.createRoleArgsForCall, struct {
		ctx  context.Context
		in   *perm_go.CreateRoleRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("CreateRole", []interface{}{ctx, in, opts})
	fake.createRoleMutex.Unlock()
	if fake.CreateRoleStub != nil {
		return fake.CreateRoleStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createRoleReturns.result1, fake.createRoleReturns.result2
}

func (fake *FakeRoleServiceClient) CreateRoleCallCount() int {
	fake.createRoleMutex.RLock()
	defer fake.createRoleMutex.RUnlock()
	return len(fake.createRoleArgsForCall)
}

func (fake *FakeRoleServiceClient) CreateRoleArgsForCall(i int) (context.Context, *perm_go.CreateRoleRequest, []grpc.CallOption) {
	fake.createRoleMutex.RLock()
	defer fake.createRoleMutex.RUnlock()
	return fake.createRoleArgsForCall[i].ctx, fake.createRoleArgsForCall[i].in, fake.createRoleArgsForCall[i].opts
}

func (fake *FakeRoleServiceClient) CreateRoleReturns(result1 *perm_go.CreateRoleResponse, result2 error) {
	fake.CreateRoleStub = nil
	fake.createRoleReturns = struct {
		result1 *perm_go.CreateRoleResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) CreateRoleReturnsOnCall(i int, result1 *perm_go.CreateRoleResponse, result2 error) {
	fake.CreateRoleStub = nil
	if fake.createRoleReturnsOnCall == nil {
		fake.createRoleReturnsOnCall = make(map[int]struct {
			result1 *perm_go.CreateRoleResponse
			result2 error
		})
	}
	fake.createRoleReturnsOnCall[i] = struct {
		result1 *perm_go.CreateRoleResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) GetRole(ctx context.Context, in *perm_go.GetRoleRequest, opts ...grpc.CallOption) (*perm_go.GetRoleResponse, error) {
	fake.getRoleMutex.Lock()
	ret, specificReturn := fake.getRoleReturnsOnCall[len(fake.getRoleArgsForCall)]
	fake.getRoleArgsForCall = append(fake.getRoleArgsForCall, struct {
		ctx  context.Context
		in   *perm_go.GetRoleRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("GetRole", []interface{}{ctx, in, opts})
	fake.getRoleMutex.Unlock()
	if fake.GetRoleStub != nil {
		return fake.GetRoleStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getRoleReturns.result1, fake.getRoleReturns.result2
}

func (fake *FakeRoleServiceClient) GetRoleCallCount() int {
	fake.getRoleMutex.RLock()
	defer fake.getRoleMutex.RUnlock()
	return len(fake.getRoleArgsForCall)
}

func (fake *FakeRoleServiceClient) GetRoleArgsForCall(i int) (context.Context, *perm_go.GetRoleRequest, []grpc.CallOption) {
	fake.getRoleMutex.RLock()
	defer fake.getRoleMutex.RUnlock()
	return fake.getRoleArgsForCall[i].ctx, fake.getRoleArgsForCall[i].in, fake.getRoleArgsForCall[i].opts
}

func (fake *FakeRoleServiceClient) GetRoleReturns(result1 *perm_go.GetRoleResponse, result2 error) {
	fake.GetRoleStub = nil
	fake.getRoleReturns = struct {
		result1 *perm_go.GetRoleResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) GetRoleReturnsOnCall(i int, result1 *perm_go.GetRoleResponse, result2 error) {
	fake.GetRoleStub = nil
	if fake.getRoleReturnsOnCall == nil {
		fake.getRoleReturnsOnCall = make(map[int]struct {
			result1 *perm_go.GetRoleResponse
			result2 error
		})
	}
	fake.getRoleReturnsOnCall[i] = struct {
		result1 *perm_go.GetRoleResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) DeleteRole(ctx context.Context, in *perm_go.DeleteRoleRequest, opts ...grpc.CallOption) (*perm_go.DeleteRoleResponse, error) {
	fake.deleteRoleMutex.Lock()
	ret, specificReturn := fake.deleteRoleReturnsOnCall[len(fake.deleteRoleArgsForCall)]
	fake.deleteRoleArgsForCall = append(fake.deleteRoleArgsForCall, struct {
		ctx  context.Context
		in   *perm_go.DeleteRoleRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("DeleteRole", []interface{}{ctx, in, opts})
	fake.deleteRoleMutex.Unlock()
	if fake.DeleteRoleStub != nil {
		return fake.DeleteRoleStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteRoleReturns.result1, fake.deleteRoleReturns.result2
}

func (fake *FakeRoleServiceClient) DeleteRoleCallCount() int {
	fake.deleteRoleMutex.RLock()
	defer fake.deleteRoleMutex.RUnlock()
	return len(fake.deleteRoleArgsForCall)
}

func (fake *FakeRoleServiceClient) DeleteRoleArgsForCall(i int) (context.Context, *perm_go.DeleteRoleRequest, []grpc.CallOption) {
	fake.deleteRoleMutex.RLock()
	defer fake.deleteRoleMutex.RUnlock()
	return fake.deleteRoleArgsForCall[i].ctx, fake.deleteRoleArgsForCall[i].in, fake.deleteRoleArgsForCall[i].opts
}

func (fake *FakeRoleServiceClient) DeleteRoleReturns(result1 *perm_go.DeleteRoleResponse, result2 error) {
	fake.DeleteRoleStub = nil
	fake.deleteRoleReturns = struct {
		result1 *perm_go.DeleteRoleResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) DeleteRoleReturnsOnCall(i int, result1 *perm_go.DeleteRoleResponse, result2 error) {
	fake.DeleteRoleStub = nil
	if fake.deleteRoleReturnsOnCall == nil {
		fake.deleteRoleReturnsOnCall = make(map[int]struct {
			result1 *perm_go.DeleteRoleResponse
			result2 error
		})
	}
	fake.deleteRoleReturnsOnCall[i] = struct {
		result1 *perm_go.DeleteRoleResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) AssignRole(ctx context.Context, in *perm_go.AssignRoleRequest, opts ...grpc.CallOption) (*perm_go.AssignRoleResponse, error) {
	fake.assignRoleMutex.Lock()
	ret, specificReturn := fake.assignRoleReturnsOnCall[len(fake.assignRoleArgsForCall)]
	fake.assignRoleArgsForCall = append(fake.assignRoleArgsForCall, struct {
		ctx  context.Context
		in   *perm_go.AssignRoleRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("AssignRole", []interface{}{ctx, in, opts})
	fake.assignRoleMutex.Unlock()
	if fake.AssignRoleStub != nil {
		return fake.AssignRoleStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.assignRoleReturns.result1, fake.assignRoleReturns.result2
}

func (fake *FakeRoleServiceClient) AssignRoleCallCount() int {
	fake.assignRoleMutex.RLock()
	defer fake.assignRoleMutex.RUnlock()
	return len(fake.assignRoleArgsForCall)
}

func (fake *FakeRoleServiceClient) AssignRoleArgsForCall(i int) (context.Context, *perm_go.AssignRoleRequest, []grpc.CallOption) {
	fake.assignRoleMutex.RLock()
	defer fake.assignRoleMutex.RUnlock()
	return fake.assignRoleArgsForCall[i].ctx, fake.assignRoleArgsForCall[i].in, fake.assignRoleArgsForCall[i].opts
}

func (fake *FakeRoleServiceClient) AssignRoleReturns(result1 *perm_go.AssignRoleResponse, result2 error) {
	fake.AssignRoleStub = nil
	fake.assignRoleReturns = struct {
		result1 *perm_go.AssignRoleResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) AssignRoleReturnsOnCall(i int, result1 *perm_go.AssignRoleResponse, result2 error) {
	fake.AssignRoleStub = nil
	if fake.assignRoleReturnsOnCall == nil {
		fake.assignRoleReturnsOnCall = make(map[int]struct {
			result1 *perm_go.AssignRoleResponse
			result2 error
		})
	}
	fake.assignRoleReturnsOnCall[i] = struct {
		result1 *perm_go.AssignRoleResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) UnassignRole(ctx context.Context, in *perm_go.UnassignRoleRequest, opts ...grpc.CallOption) (*perm_go.UnassignRoleResponse, error) {
	fake.unassignRoleMutex.Lock()
	ret, specificReturn := fake.unassignRoleReturnsOnCall[len(fake.unassignRoleArgsForCall)]
	fake.unassignRoleArgsForCall = append(fake.unassignRoleArgsForCall, struct {
		ctx  context.Context
		in   *perm_go.UnassignRoleRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("UnassignRole", []interface{}{ctx, in, opts})
	fake.unassignRoleMutex.Unlock()
	if fake.UnassignRoleStub != nil {
		return fake.UnassignRoleStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.unassignRoleReturns.result1, fake.unassignRoleReturns.result2
}

func (fake *FakeRoleServiceClient) UnassignRoleCallCount() int {
	fake.unassignRoleMutex.RLock()
	defer fake.unassignRoleMutex.RUnlock()
	return len(fake.unassignRoleArgsForCall)
}

func (fake *FakeRoleServiceClient) UnassignRoleArgsForCall(i int) (context.Context, *perm_go.UnassignRoleRequest, []grpc.CallOption) {
	fake.unassignRoleMutex.RLock()
	defer fake.unassignRoleMutex.RUnlock()
	return fake.unassignRoleArgsForCall[i].ctx, fake.unassignRoleArgsForCall[i].in, fake.unassignRoleArgsForCall[i].opts
}

func (fake *FakeRoleServiceClient) UnassignRoleReturns(result1 *perm_go.UnassignRoleResponse, result2 error) {
	fake.UnassignRoleStub = nil
	fake.unassignRoleReturns = struct {
		result1 *perm_go.UnassignRoleResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) UnassignRoleReturnsOnCall(i int, result1 *perm_go.UnassignRoleResponse, result2 error) {
	fake.UnassignRoleStub = nil
	if fake.unassignRoleReturnsOnCall == nil {
		fake.unassignRoleReturnsOnCall = make(map[int]struct {
			result1 *perm_go.UnassignRoleResponse
			result2 error
		})
	}
	fake.unassignRoleReturnsOnCall[i] = struct {
		result1 *perm_go.UnassignRoleResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) HasRole(ctx context.Context, in *perm_go.HasRoleRequest, opts ...grpc.CallOption) (*perm_go.HasRoleResponse, error) {
	fake.hasRoleMutex.Lock()
	ret, specificReturn := fake.hasRoleReturnsOnCall[len(fake.hasRoleArgsForCall)]
	fake.hasRoleArgsForCall = append(fake.hasRoleArgsForCall, struct {
		ctx  context.Context
		in   *perm_go.HasRoleRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("HasRole", []interface{}{ctx, in, opts})
	fake.hasRoleMutex.Unlock()
	if fake.HasRoleStub != nil {
		return fake.HasRoleStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.hasRoleReturns.result1, fake.hasRoleReturns.result2
}

func (fake *FakeRoleServiceClient) HasRoleCallCount() int {
	fake.hasRoleMutex.RLock()
	defer fake.hasRoleMutex.RUnlock()
	return len(fake.hasRoleArgsForCall)
}

func (fake *FakeRoleServiceClient) HasRoleArgsForCall(i int) (context.Context, *perm_go.HasRoleRequest, []grpc.CallOption) {
	fake.hasRoleMutex.RLock()
	defer fake.hasRoleMutex.RUnlock()
	return fake.hasRoleArgsForCall[i].ctx, fake.hasRoleArgsForCall[i].in, fake.hasRoleArgsForCall[i].opts
}

func (fake *FakeRoleServiceClient) HasRoleReturns(result1 *perm_go.HasRoleResponse, result2 error) {
	fake.HasRoleStub = nil
	fake.hasRoleReturns = struct {
		result1 *perm_go.HasRoleResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) HasRoleReturnsOnCall(i int, result1 *perm_go.HasRoleResponse, result2 error) {
	fake.HasRoleStub = nil
	if fake.hasRoleReturnsOnCall == nil {
		fake.hasRoleReturnsOnCall = make(map[int]struct {
			result1 *perm_go.HasRoleResponse
			result2 error
		})
	}
	fake.hasRoleReturnsOnCall[i] = struct {
		result1 *perm_go.HasRoleResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) ListActorRoles(ctx context.Context, in *perm_go.ListActorRolesRequest, opts ...grpc.CallOption) (*perm_go.ListActorRolesResponse, error) {
	fake.listActorRolesMutex.Lock()
	ret, specificReturn := fake.listActorRolesReturnsOnCall[len(fake.listActorRolesArgsForCall)]
	fake.listActorRolesArgsForCall = append(fake.listActorRolesArgsForCall, struct {
		ctx  context.Context
		in   *perm_go.ListActorRolesRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("ListActorRoles", []interface{}{ctx, in, opts})
	fake.listActorRolesMutex.Unlock()
	if fake.ListActorRolesStub != nil {
		return fake.ListActorRolesStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listActorRolesReturns.result1, fake.listActorRolesReturns.result2
}

func (fake *FakeRoleServiceClient) ListActorRolesCallCount() int {
	fake.listActorRolesMutex.RLock()
	defer fake.listActorRolesMutex.RUnlock()
	return len(fake.listActorRolesArgsForCall)
}

func (fake *FakeRoleServiceClient) ListActorRolesArgsForCall(i int) (context.Context, *perm_go.ListActorRolesRequest, []grpc.CallOption) {
	fake.listActorRolesMutex.RLock()
	defer fake.listActorRolesMutex.RUnlock()
	return fake.listActorRolesArgsForCall[i].ctx, fake.listActorRolesArgsForCall[i].in, fake.listActorRolesArgsForCall[i].opts
}

func (fake *FakeRoleServiceClient) ListActorRolesReturns(result1 *perm_go.ListActorRolesResponse, result2 error) {
	fake.ListActorRolesStub = nil
	fake.listActorRolesReturns = struct {
		result1 *perm_go.ListActorRolesResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) ListActorRolesReturnsOnCall(i int, result1 *perm_go.ListActorRolesResponse, result2 error) {
	fake.ListActorRolesStub = nil
	if fake.listActorRolesReturnsOnCall == nil {
		fake.listActorRolesReturnsOnCall = make(map[int]struct {
			result1 *perm_go.ListActorRolesResponse
			result2 error
		})
	}
	fake.listActorRolesReturnsOnCall[i] = struct {
		result1 *perm_go.ListActorRolesResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) ListRolePermissions(ctx context.Context, in *perm_go.ListRolePermissionsRequest, opts ...grpc.CallOption) (*perm_go.ListRolePermissionsResponse, error) {
	fake.listRolePermissionsMutex.Lock()
	ret, specificReturn := fake.listRolePermissionsReturnsOnCall[len(fake.listRolePermissionsArgsForCall)]
	fake.listRolePermissionsArgsForCall = append(fake.listRolePermissionsArgsForCall, struct {
		ctx  context.Context
		in   *perm_go.ListRolePermissionsRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("ListRolePermissions", []interface{}{ctx, in, opts})
	fake.listRolePermissionsMutex.Unlock()
	if fake.ListRolePermissionsStub != nil {
		return fake.ListRolePermissionsStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listRolePermissionsReturns.result1, fake.listRolePermissionsReturns.result2
}

func (fake *FakeRoleServiceClient) ListRolePermissionsCallCount() int {
	fake.listRolePermissionsMutex.RLock()
	defer fake.listRolePermissionsMutex.RUnlock()
	return len(fake.listRolePermissionsArgsForCall)
}

func (fake *FakeRoleServiceClient) ListRolePermissionsArgsForCall(i int) (context.Context, *perm_go.ListRolePermissionsRequest, []grpc.CallOption) {
	fake.listRolePermissionsMutex.RLock()
	defer fake.listRolePermissionsMutex.RUnlock()
	return fake.listRolePermissionsArgsForCall[i].ctx, fake.listRolePermissionsArgsForCall[i].in, fake.listRolePermissionsArgsForCall[i].opts
}

func (fake *FakeRoleServiceClient) ListRolePermissionsReturns(result1 *perm_go.ListRolePermissionsResponse, result2 error) {
	fake.ListRolePermissionsStub = nil
	fake.listRolePermissionsReturns = struct {
		result1 *perm_go.ListRolePermissionsResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) ListRolePermissionsReturnsOnCall(i int, result1 *perm_go.ListRolePermissionsResponse, result2 error) {
	fake.ListRolePermissionsStub = nil
	if fake.listRolePermissionsReturnsOnCall == nil {
		fake.listRolePermissionsReturnsOnCall = make(map[int]struct {
			result1 *perm_go.ListRolePermissionsResponse
			result2 error
		})
	}
	fake.listRolePermissionsReturnsOnCall[i] = struct {
		result1 *perm_go.ListRolePermissionsResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleServiceClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createRoleMutex.RLock()
	defer fake.createRoleMutex.RUnlock()
	fake.getRoleMutex.RLock()
	defer fake.getRoleMutex.RUnlock()
	fake.deleteRoleMutex.RLock()
	defer fake.deleteRoleMutex.RUnlock()
	fake.assignRoleMutex.RLock()
	defer fake.assignRoleMutex.RUnlock()
	fake.unassignRoleMutex.RLock()
	defer fake.unassignRoleMutex.RUnlock()
	fake.hasRoleMutex.RLock()
	defer fake.hasRoleMutex.RUnlock()
	fake.listActorRolesMutex.RLock()
	defer fake.listActorRolesMutex.RUnlock()
	fake.listRolePermissionsMutex.RLock()
	defer fake.listRolePermissionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRoleServiceClient) recordInvocation(key string, args []interface{}) {
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

var _ perm_go.RoleServiceClient = new(FakeRoleServiceClient)
