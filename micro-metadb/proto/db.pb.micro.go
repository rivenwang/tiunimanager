// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: db.proto

package db

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for TiEMDBService service

func NewTiEMDBServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TiEMDBService service

type TiEMDBService interface {
	// Auth Module
	FindTenant(ctx context.Context, in *DBFindTenantRequest, opts ...client.CallOption) (*DBFindTenantResponse, error)
	FindAccount(ctx context.Context, in *DBFindAccountRequest, opts ...client.CallOption) (*DBFindAccountResponse, error)
	SaveToken(ctx context.Context, in *DBSaveTokenRequest, opts ...client.CallOption) (*DBSaveTokenResponse, error)
	FindToken(ctx context.Context, in *DBFindTokenRequest, opts ...client.CallOption) (*DBFindTokenResponse, error)
	FindRolesByPermission(ctx context.Context, in *DBFindRolesByPermissionRequest, opts ...client.CallOption) (*DBFindRolesByPermissionResponse, error)
	// Host Module
	AddHost(ctx context.Context, in *DBAddHostRequest, opts ...client.CallOption) (*DBAddHostResponse, error)
	AddHostsInBatch(ctx context.Context, in *DBAddHostsInBatchRequest, opts ...client.CallOption) (*DBAddHostsInBatchResponse, error)
	RemoveHost(ctx context.Context, in *DBRemoveHostRequest, opts ...client.CallOption) (*DBRemoveHostResponse, error)
	RemoveHostsInBatch(ctx context.Context, in *DBRemoveHostsInBatchRequest, opts ...client.CallOption) (*DBRemoveHostsInBatchResponse, error)
	ListHost(ctx context.Context, in *DBListHostsRequest, opts ...client.CallOption) (*DBListHostsResponse, error)
	CheckDetails(ctx context.Context, in *DBCheckDetailsRequest, opts ...client.CallOption) (*DBCheckDetailsResponse, error)
	PreAllocHosts(ctx context.Context, in *DBPreAllocHostsRequest, opts ...client.CallOption) (*DBPreAllocHostsResponse, error)
	LockHosts(ctx context.Context, in *DBLockHostsRequest, opts ...client.CallOption) (*DBLockHostsResponse, error)
	GetFailureDomain(ctx context.Context, in *DBGetFailureDomainRequest, opts ...client.CallOption) (*DBGetFailureDomainResponse, error)
	// Cluster
	CreateCluster(ctx context.Context, in *DBCreateClusterRequest, opts ...client.CallOption) (*DBCreateClusterResponse, error)
	DeleteCluster(ctx context.Context, in *DBDeleteClusterRequest, opts ...client.CallOption) (*DBDeleteClusterResponse, error)
	UpdateClusterStatus(ctx context.Context, in *DBUpdateClusterStatusRequest, opts ...client.CallOption) (*DBUpdateClusterStatusResponse, error)
	UpdateClusterTiupConfig(ctx context.Context, in *DBUpdateTiupConfigRequest, opts ...client.CallOption) (*DBUpdateTiupConfigResponse, error)
	LoadCluster(ctx context.Context, in *DBLoadClusterRequest, opts ...client.CallOption) (*DBLoadClusterResponse, error)
	ListCluster(ctx context.Context, in *DBListClusterRequest, opts ...client.CallOption) (*DBListClusterResponse, error)
	// backup & recover & parameters
	SaveBackupRecord(ctx context.Context, in *DBSaveBackupRecordRequest, opts ...client.CallOption) (*DBSaveBackupRecordResponse, error)
	ListBackupRecords(ctx context.Context, in *DBListBackupRecordsRequest, opts ...client.CallOption) (*DBListBackupRecordsResponse, error)
	SaveRecoverRecord(ctx context.Context, in *DBSaveRecoverRecordRequest, opts ...client.CallOption) (*DBSaveRecoverRecordResponse, error)
	SaveParametersRecord(ctx context.Context, in *DBSaveParametersRequest, opts ...client.CallOption) (*DBSaveParametersResponse, error)
	GetCurrentParametersRecord(ctx context.Context, in *DBGetCurrentParametersRequest, opts ...client.CallOption) (*DBGetCurrentParametersResponse, error)
	// Tiup Task
	CreateTiupTask(ctx context.Context, in *CreateTiupTaskRequest, opts ...client.CallOption) (*CreateTiupTaskResponse, error)
	UpdateTiupTask(ctx context.Context, in *UpdateTiupTaskRequest, opts ...client.CallOption) (*UpdateTiupTaskResponse, error)
	FindTiupTaskByID(ctx context.Context, in *FindTiupTaskByIDRequest, opts ...client.CallOption) (*FindTiupTaskByIDResponse, error)
	GetTiupTaskStatusByBizID(ctx context.Context, in *GetTiupTaskStatusByBizIDRequest, opts ...client.CallOption) (*GetTiupTaskStatusByBizIDResponse, error)
	// Workflow and Task
	CreateFlow(ctx context.Context, in *DBCreateFlowRequest, opts ...client.CallOption) (*DBCreateFlowResponse, error)
	CreateTask(ctx context.Context, in *DBCreateTaskRequest, opts ...client.CallOption) (*DBCreateTaskResponse, error)
	UpdateFlow(ctx context.Context, in *DBUpdateFlowRequest, opts ...client.CallOption) (*DBUpdateFlowResponse, error)
	UpdateTask(ctx context.Context, in *DBUpdateTaskRequest, opts ...client.CallOption) (*DBUpdateTaskResponse, error)
	LoadFlow(ctx context.Context, in *DBLoadFlowRequest, opts ...client.CallOption) (*DBLoadFlowResponse, error)
	LoadTask(ctx context.Context, in *DBLoadTaskRequest, opts ...client.CallOption) (*DBLoadTaskResponse, error)
	// DataTransport
	CreateTransportRecord(ctx context.Context, in *DBCreateTransportRecordRequest, opts ...client.CallOption) (*DBCreateTransportRecordResponse, error)
	UpdateTransportRecord(ctx context.Context, in *DBUpdateTransportRecordRequest, opts ...client.CallOption) (*DBUpdateTransportRecordResponse, error)
	FindTrasnportRecordByID(ctx context.Context, in *DBFindTransportRecordByIDRequest, opts ...client.CallOption) (*DBFindTransportRecordByIDResponse, error)
}

type tiEMDBService struct {
	c    client.Client
	name string
}

func NewTiEMDBService(name string, c client.Client) TiEMDBService {
	return &tiEMDBService{
		c:    c,
		name: name,
	}
}

func (c *tiEMDBService) FindTenant(ctx context.Context, in *DBFindTenantRequest, opts ...client.CallOption) (*DBFindTenantResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.FindTenant", in)
	out := new(DBFindTenantResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) FindAccount(ctx context.Context, in *DBFindAccountRequest, opts ...client.CallOption) (*DBFindAccountResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.FindAccount", in)
	out := new(DBFindAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) SaveToken(ctx context.Context, in *DBSaveTokenRequest, opts ...client.CallOption) (*DBSaveTokenResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.SaveToken", in)
	out := new(DBSaveTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) FindToken(ctx context.Context, in *DBFindTokenRequest, opts ...client.CallOption) (*DBFindTokenResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.FindToken", in)
	out := new(DBFindTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) FindRolesByPermission(ctx context.Context, in *DBFindRolesByPermissionRequest, opts ...client.CallOption) (*DBFindRolesByPermissionResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.FindRolesByPermission", in)
	out := new(DBFindRolesByPermissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) AddHost(ctx context.Context, in *DBAddHostRequest, opts ...client.CallOption) (*DBAddHostResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.AddHost", in)
	out := new(DBAddHostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) AddHostsInBatch(ctx context.Context, in *DBAddHostsInBatchRequest, opts ...client.CallOption) (*DBAddHostsInBatchResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.AddHostsInBatch", in)
	out := new(DBAddHostsInBatchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) RemoveHost(ctx context.Context, in *DBRemoveHostRequest, opts ...client.CallOption) (*DBRemoveHostResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.RemoveHost", in)
	out := new(DBRemoveHostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) RemoveHostsInBatch(ctx context.Context, in *DBRemoveHostsInBatchRequest, opts ...client.CallOption) (*DBRemoveHostsInBatchResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.RemoveHostsInBatch", in)
	out := new(DBRemoveHostsInBatchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) ListHost(ctx context.Context, in *DBListHostsRequest, opts ...client.CallOption) (*DBListHostsResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.ListHost", in)
	out := new(DBListHostsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) CheckDetails(ctx context.Context, in *DBCheckDetailsRequest, opts ...client.CallOption) (*DBCheckDetailsResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.CheckDetails", in)
	out := new(DBCheckDetailsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) PreAllocHosts(ctx context.Context, in *DBPreAllocHostsRequest, opts ...client.CallOption) (*DBPreAllocHostsResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.PreAllocHosts", in)
	out := new(DBPreAllocHostsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) LockHosts(ctx context.Context, in *DBLockHostsRequest, opts ...client.CallOption) (*DBLockHostsResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.LockHosts", in)
	out := new(DBLockHostsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) GetFailureDomain(ctx context.Context, in *DBGetFailureDomainRequest, opts ...client.CallOption) (*DBGetFailureDomainResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.GetFailureDomain", in)
	out := new(DBGetFailureDomainResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) CreateCluster(ctx context.Context, in *DBCreateClusterRequest, opts ...client.CallOption) (*DBCreateClusterResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.CreateCluster", in)
	out := new(DBCreateClusterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) DeleteCluster(ctx context.Context, in *DBDeleteClusterRequest, opts ...client.CallOption) (*DBDeleteClusterResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.DeleteCluster", in)
	out := new(DBDeleteClusterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) UpdateClusterStatus(ctx context.Context, in *DBUpdateClusterStatusRequest, opts ...client.CallOption) (*DBUpdateClusterStatusResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.UpdateClusterStatus", in)
	out := new(DBUpdateClusterStatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) UpdateClusterTiupConfig(ctx context.Context, in *DBUpdateTiupConfigRequest, opts ...client.CallOption) (*DBUpdateTiupConfigResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.UpdateClusterTiupConfig", in)
	out := new(DBUpdateTiupConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) LoadCluster(ctx context.Context, in *DBLoadClusterRequest, opts ...client.CallOption) (*DBLoadClusterResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.LoadCluster", in)
	out := new(DBLoadClusterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) ListCluster(ctx context.Context, in *DBListClusterRequest, opts ...client.CallOption) (*DBListClusterResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.ListCluster", in)
	out := new(DBListClusterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) SaveBackupRecord(ctx context.Context, in *DBSaveBackupRecordRequest, opts ...client.CallOption) (*DBSaveBackupRecordResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.SaveBackupRecord", in)
	out := new(DBSaveBackupRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) ListBackupRecords(ctx context.Context, in *DBListBackupRecordsRequest, opts ...client.CallOption) (*DBListBackupRecordsResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.ListBackupRecords", in)
	out := new(DBListBackupRecordsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) SaveRecoverRecord(ctx context.Context, in *DBSaveRecoverRecordRequest, opts ...client.CallOption) (*DBSaveRecoverRecordResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.SaveRecoverRecord", in)
	out := new(DBSaveRecoverRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) SaveParametersRecord(ctx context.Context, in *DBSaveParametersRequest, opts ...client.CallOption) (*DBSaveParametersResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.SaveParametersRecord", in)
	out := new(DBSaveParametersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) GetCurrentParametersRecord(ctx context.Context, in *DBGetCurrentParametersRequest, opts ...client.CallOption) (*DBGetCurrentParametersResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.GetCurrentParametersRecord", in)
	out := new(DBGetCurrentParametersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) CreateTiupTask(ctx context.Context, in *CreateTiupTaskRequest, opts ...client.CallOption) (*CreateTiupTaskResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.CreateTiupTask", in)
	out := new(CreateTiupTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) UpdateTiupTask(ctx context.Context, in *UpdateTiupTaskRequest, opts ...client.CallOption) (*UpdateTiupTaskResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.UpdateTiupTask", in)
	out := new(UpdateTiupTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) FindTiupTaskByID(ctx context.Context, in *FindTiupTaskByIDRequest, opts ...client.CallOption) (*FindTiupTaskByIDResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.FindTiupTaskByID", in)
	out := new(FindTiupTaskByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) GetTiupTaskStatusByBizID(ctx context.Context, in *GetTiupTaskStatusByBizIDRequest, opts ...client.CallOption) (*GetTiupTaskStatusByBizIDResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.GetTiupTaskStatusByBizID", in)
	out := new(GetTiupTaskStatusByBizIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) CreateFlow(ctx context.Context, in *DBCreateFlowRequest, opts ...client.CallOption) (*DBCreateFlowResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.CreateFlow", in)
	out := new(DBCreateFlowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) CreateTask(ctx context.Context, in *DBCreateTaskRequest, opts ...client.CallOption) (*DBCreateTaskResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.CreateTask", in)
	out := new(DBCreateTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) UpdateFlow(ctx context.Context, in *DBUpdateFlowRequest, opts ...client.CallOption) (*DBUpdateFlowResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.UpdateFlow", in)
	out := new(DBUpdateFlowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) UpdateTask(ctx context.Context, in *DBUpdateTaskRequest, opts ...client.CallOption) (*DBUpdateTaskResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.UpdateTask", in)
	out := new(DBUpdateTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) LoadFlow(ctx context.Context, in *DBLoadFlowRequest, opts ...client.CallOption) (*DBLoadFlowResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.LoadFlow", in)
	out := new(DBLoadFlowResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) LoadTask(ctx context.Context, in *DBLoadTaskRequest, opts ...client.CallOption) (*DBLoadTaskResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.LoadTask", in)
	out := new(DBLoadTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) CreateTransportRecord(ctx context.Context, in *DBCreateTransportRecordRequest, opts ...client.CallOption) (*DBCreateTransportRecordResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.CreateTransportRecord", in)
	out := new(DBCreateTransportRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) UpdateTransportRecord(ctx context.Context, in *DBUpdateTransportRecordRequest, opts ...client.CallOption) (*DBUpdateTransportRecordResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.UpdateTransportRecord", in)
	out := new(DBUpdateTransportRecordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiEMDBService) FindTrasnportRecordByID(ctx context.Context, in *DBFindTransportRecordByIDRequest, opts ...client.CallOption) (*DBFindTransportRecordByIDResponse, error) {
	req := c.c.NewRequest(c.name, "TiEMDBService.FindTrasnportRecordByID", in)
	out := new(DBFindTransportRecordByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TiEMDBService service

type TiEMDBServiceHandler interface {
	// Auth Module
	FindTenant(context.Context, *DBFindTenantRequest, *DBFindTenantResponse) error
	FindAccount(context.Context, *DBFindAccountRequest, *DBFindAccountResponse) error
	SaveToken(context.Context, *DBSaveTokenRequest, *DBSaveTokenResponse) error
	FindToken(context.Context, *DBFindTokenRequest, *DBFindTokenResponse) error
	FindRolesByPermission(context.Context, *DBFindRolesByPermissionRequest, *DBFindRolesByPermissionResponse) error
	// Host Module
	AddHost(context.Context, *DBAddHostRequest, *DBAddHostResponse) error
	AddHostsInBatch(context.Context, *DBAddHostsInBatchRequest, *DBAddHostsInBatchResponse) error
	RemoveHost(context.Context, *DBRemoveHostRequest, *DBRemoveHostResponse) error
	RemoveHostsInBatch(context.Context, *DBRemoveHostsInBatchRequest, *DBRemoveHostsInBatchResponse) error
	ListHost(context.Context, *DBListHostsRequest, *DBListHostsResponse) error
	CheckDetails(context.Context, *DBCheckDetailsRequest, *DBCheckDetailsResponse) error
	PreAllocHosts(context.Context, *DBPreAllocHostsRequest, *DBPreAllocHostsResponse) error
	LockHosts(context.Context, *DBLockHostsRequest, *DBLockHostsResponse) error
	GetFailureDomain(context.Context, *DBGetFailureDomainRequest, *DBGetFailureDomainResponse) error
	// Cluster
	CreateCluster(context.Context, *DBCreateClusterRequest, *DBCreateClusterResponse) error
	DeleteCluster(context.Context, *DBDeleteClusterRequest, *DBDeleteClusterResponse) error
	UpdateClusterStatus(context.Context, *DBUpdateClusterStatusRequest, *DBUpdateClusterStatusResponse) error
	UpdateClusterTiupConfig(context.Context, *DBUpdateTiupConfigRequest, *DBUpdateTiupConfigResponse) error
	LoadCluster(context.Context, *DBLoadClusterRequest, *DBLoadClusterResponse) error
	ListCluster(context.Context, *DBListClusterRequest, *DBListClusterResponse) error
	// backup & recover & parameters
	SaveBackupRecord(context.Context, *DBSaveBackupRecordRequest, *DBSaveBackupRecordResponse) error
	ListBackupRecords(context.Context, *DBListBackupRecordsRequest, *DBListBackupRecordsResponse) error
	SaveRecoverRecord(context.Context, *DBSaveRecoverRecordRequest, *DBSaveRecoverRecordResponse) error
	SaveParametersRecord(context.Context, *DBSaveParametersRequest, *DBSaveParametersResponse) error
	GetCurrentParametersRecord(context.Context, *DBGetCurrentParametersRequest, *DBGetCurrentParametersResponse) error
	// Tiup Task
	CreateTiupTask(context.Context, *CreateTiupTaskRequest, *CreateTiupTaskResponse) error
	UpdateTiupTask(context.Context, *UpdateTiupTaskRequest, *UpdateTiupTaskResponse) error
	FindTiupTaskByID(context.Context, *FindTiupTaskByIDRequest, *FindTiupTaskByIDResponse) error
	GetTiupTaskStatusByBizID(context.Context, *GetTiupTaskStatusByBizIDRequest, *GetTiupTaskStatusByBizIDResponse) error
	// Workflow and Task
	CreateFlow(context.Context, *DBCreateFlowRequest, *DBCreateFlowResponse) error
	CreateTask(context.Context, *DBCreateTaskRequest, *DBCreateTaskResponse) error
	UpdateFlow(context.Context, *DBUpdateFlowRequest, *DBUpdateFlowResponse) error
	UpdateTask(context.Context, *DBUpdateTaskRequest, *DBUpdateTaskResponse) error
	LoadFlow(context.Context, *DBLoadFlowRequest, *DBLoadFlowResponse) error
	LoadTask(context.Context, *DBLoadTaskRequest, *DBLoadTaskResponse) error
	// DataTransport
	CreateTransportRecord(context.Context, *DBCreateTransportRecordRequest, *DBCreateTransportRecordResponse) error
	UpdateTransportRecord(context.Context, *DBUpdateTransportRecordRequest, *DBUpdateTransportRecordResponse) error
	FindTrasnportRecordByID(context.Context, *DBFindTransportRecordByIDRequest, *DBFindTransportRecordByIDResponse) error
}

func RegisterTiEMDBServiceHandler(s server.Server, hdlr TiEMDBServiceHandler, opts ...server.HandlerOption) error {
	type tiEMDBService interface {
		FindTenant(ctx context.Context, in *DBFindTenantRequest, out *DBFindTenantResponse) error
		FindAccount(ctx context.Context, in *DBFindAccountRequest, out *DBFindAccountResponse) error
		SaveToken(ctx context.Context, in *DBSaveTokenRequest, out *DBSaveTokenResponse) error
		FindToken(ctx context.Context, in *DBFindTokenRequest, out *DBFindTokenResponse) error
		FindRolesByPermission(ctx context.Context, in *DBFindRolesByPermissionRequest, out *DBFindRolesByPermissionResponse) error
		AddHost(ctx context.Context, in *DBAddHostRequest, out *DBAddHostResponse) error
		AddHostsInBatch(ctx context.Context, in *DBAddHostsInBatchRequest, out *DBAddHostsInBatchResponse) error
		RemoveHost(ctx context.Context, in *DBRemoveHostRequest, out *DBRemoveHostResponse) error
		RemoveHostsInBatch(ctx context.Context, in *DBRemoveHostsInBatchRequest, out *DBRemoveHostsInBatchResponse) error
		ListHost(ctx context.Context, in *DBListHostsRequest, out *DBListHostsResponse) error
		CheckDetails(ctx context.Context, in *DBCheckDetailsRequest, out *DBCheckDetailsResponse) error
		PreAllocHosts(ctx context.Context, in *DBPreAllocHostsRequest, out *DBPreAllocHostsResponse) error
		LockHosts(ctx context.Context, in *DBLockHostsRequest, out *DBLockHostsResponse) error
		GetFailureDomain(ctx context.Context, in *DBGetFailureDomainRequest, out *DBGetFailureDomainResponse) error
		CreateCluster(ctx context.Context, in *DBCreateClusterRequest, out *DBCreateClusterResponse) error
		DeleteCluster(ctx context.Context, in *DBDeleteClusterRequest, out *DBDeleteClusterResponse) error
		UpdateClusterStatus(ctx context.Context, in *DBUpdateClusterStatusRequest, out *DBUpdateClusterStatusResponse) error
		UpdateClusterTiupConfig(ctx context.Context, in *DBUpdateTiupConfigRequest, out *DBUpdateTiupConfigResponse) error
		LoadCluster(ctx context.Context, in *DBLoadClusterRequest, out *DBLoadClusterResponse) error
		ListCluster(ctx context.Context, in *DBListClusterRequest, out *DBListClusterResponse) error
		SaveBackupRecord(ctx context.Context, in *DBSaveBackupRecordRequest, out *DBSaveBackupRecordResponse) error
		ListBackupRecords(ctx context.Context, in *DBListBackupRecordsRequest, out *DBListBackupRecordsResponse) error
		SaveRecoverRecord(ctx context.Context, in *DBSaveRecoverRecordRequest, out *DBSaveRecoverRecordResponse) error
		SaveParametersRecord(ctx context.Context, in *DBSaveParametersRequest, out *DBSaveParametersResponse) error
		GetCurrentParametersRecord(ctx context.Context, in *DBGetCurrentParametersRequest, out *DBGetCurrentParametersResponse) error
		CreateTiupTask(ctx context.Context, in *CreateTiupTaskRequest, out *CreateTiupTaskResponse) error
		UpdateTiupTask(ctx context.Context, in *UpdateTiupTaskRequest, out *UpdateTiupTaskResponse) error
		FindTiupTaskByID(ctx context.Context, in *FindTiupTaskByIDRequest, out *FindTiupTaskByIDResponse) error
		GetTiupTaskStatusByBizID(ctx context.Context, in *GetTiupTaskStatusByBizIDRequest, out *GetTiupTaskStatusByBizIDResponse) error
		CreateFlow(ctx context.Context, in *DBCreateFlowRequest, out *DBCreateFlowResponse) error
		CreateTask(ctx context.Context, in *DBCreateTaskRequest, out *DBCreateTaskResponse) error
		UpdateFlow(ctx context.Context, in *DBUpdateFlowRequest, out *DBUpdateFlowResponse) error
		UpdateTask(ctx context.Context, in *DBUpdateTaskRequest, out *DBUpdateTaskResponse) error
		LoadFlow(ctx context.Context, in *DBLoadFlowRequest, out *DBLoadFlowResponse) error
		LoadTask(ctx context.Context, in *DBLoadTaskRequest, out *DBLoadTaskResponse) error
		CreateTransportRecord(ctx context.Context, in *DBCreateTransportRecordRequest, out *DBCreateTransportRecordResponse) error
		UpdateTransportRecord(ctx context.Context, in *DBUpdateTransportRecordRequest, out *DBUpdateTransportRecordResponse) error
		FindTrasnportRecordByID(ctx context.Context, in *DBFindTransportRecordByIDRequest, out *DBFindTransportRecordByIDResponse) error
	}
	type TiEMDBService struct {
		tiEMDBService
	}
	h := &tiEMDBServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TiEMDBService{h}, opts...))
}

type tiEMDBServiceHandler struct {
	TiEMDBServiceHandler
}

func (h *tiEMDBServiceHandler) FindTenant(ctx context.Context, in *DBFindTenantRequest, out *DBFindTenantResponse) error {
	return h.TiEMDBServiceHandler.FindTenant(ctx, in, out)
}

func (h *tiEMDBServiceHandler) FindAccount(ctx context.Context, in *DBFindAccountRequest, out *DBFindAccountResponse) error {
	return h.TiEMDBServiceHandler.FindAccount(ctx, in, out)
}

func (h *tiEMDBServiceHandler) SaveToken(ctx context.Context, in *DBSaveTokenRequest, out *DBSaveTokenResponse) error {
	return h.TiEMDBServiceHandler.SaveToken(ctx, in, out)
}

func (h *tiEMDBServiceHandler) FindToken(ctx context.Context, in *DBFindTokenRequest, out *DBFindTokenResponse) error {
	return h.TiEMDBServiceHandler.FindToken(ctx, in, out)
}

func (h *tiEMDBServiceHandler) FindRolesByPermission(ctx context.Context, in *DBFindRolesByPermissionRequest, out *DBFindRolesByPermissionResponse) error {
	return h.TiEMDBServiceHandler.FindRolesByPermission(ctx, in, out)
}

func (h *tiEMDBServiceHandler) AddHost(ctx context.Context, in *DBAddHostRequest, out *DBAddHostResponse) error {
	return h.TiEMDBServiceHandler.AddHost(ctx, in, out)
}

func (h *tiEMDBServiceHandler) AddHostsInBatch(ctx context.Context, in *DBAddHostsInBatchRequest, out *DBAddHostsInBatchResponse) error {
	return h.TiEMDBServiceHandler.AddHostsInBatch(ctx, in, out)
}

func (h *tiEMDBServiceHandler) RemoveHost(ctx context.Context, in *DBRemoveHostRequest, out *DBRemoveHostResponse) error {
	return h.TiEMDBServiceHandler.RemoveHost(ctx, in, out)
}

func (h *tiEMDBServiceHandler) RemoveHostsInBatch(ctx context.Context, in *DBRemoveHostsInBatchRequest, out *DBRemoveHostsInBatchResponse) error {
	return h.TiEMDBServiceHandler.RemoveHostsInBatch(ctx, in, out)
}

func (h *tiEMDBServiceHandler) ListHost(ctx context.Context, in *DBListHostsRequest, out *DBListHostsResponse) error {
	return h.TiEMDBServiceHandler.ListHost(ctx, in, out)
}

func (h *tiEMDBServiceHandler) CheckDetails(ctx context.Context, in *DBCheckDetailsRequest, out *DBCheckDetailsResponse) error {
	return h.TiEMDBServiceHandler.CheckDetails(ctx, in, out)
}

func (h *tiEMDBServiceHandler) PreAllocHosts(ctx context.Context, in *DBPreAllocHostsRequest, out *DBPreAllocHostsResponse) error {
	return h.TiEMDBServiceHandler.PreAllocHosts(ctx, in, out)
}

func (h *tiEMDBServiceHandler) LockHosts(ctx context.Context, in *DBLockHostsRequest, out *DBLockHostsResponse) error {
	return h.TiEMDBServiceHandler.LockHosts(ctx, in, out)
}

func (h *tiEMDBServiceHandler) GetFailureDomain(ctx context.Context, in *DBGetFailureDomainRequest, out *DBGetFailureDomainResponse) error {
	return h.TiEMDBServiceHandler.GetFailureDomain(ctx, in, out)
}

func (h *tiEMDBServiceHandler) CreateCluster(ctx context.Context, in *DBCreateClusterRequest, out *DBCreateClusterResponse) error {
	return h.TiEMDBServiceHandler.CreateCluster(ctx, in, out)
}

func (h *tiEMDBServiceHandler) DeleteCluster(ctx context.Context, in *DBDeleteClusterRequest, out *DBDeleteClusterResponse) error {
	return h.TiEMDBServiceHandler.DeleteCluster(ctx, in, out)
}

func (h *tiEMDBServiceHandler) UpdateClusterStatus(ctx context.Context, in *DBUpdateClusterStatusRequest, out *DBUpdateClusterStatusResponse) error {
	return h.TiEMDBServiceHandler.UpdateClusterStatus(ctx, in, out)
}

func (h *tiEMDBServiceHandler) UpdateClusterTiupConfig(ctx context.Context, in *DBUpdateTiupConfigRequest, out *DBUpdateTiupConfigResponse) error {
	return h.TiEMDBServiceHandler.UpdateClusterTiupConfig(ctx, in, out)
}

func (h *tiEMDBServiceHandler) LoadCluster(ctx context.Context, in *DBLoadClusterRequest, out *DBLoadClusterResponse) error {
	return h.TiEMDBServiceHandler.LoadCluster(ctx, in, out)
}

func (h *tiEMDBServiceHandler) ListCluster(ctx context.Context, in *DBListClusterRequest, out *DBListClusterResponse) error {
	return h.TiEMDBServiceHandler.ListCluster(ctx, in, out)
}

func (h *tiEMDBServiceHandler) SaveBackupRecord(ctx context.Context, in *DBSaveBackupRecordRequest, out *DBSaveBackupRecordResponse) error {
	return h.TiEMDBServiceHandler.SaveBackupRecord(ctx, in, out)
}

func (h *tiEMDBServiceHandler) ListBackupRecords(ctx context.Context, in *DBListBackupRecordsRequest, out *DBListBackupRecordsResponse) error {
	return h.TiEMDBServiceHandler.ListBackupRecords(ctx, in, out)
}

func (h *tiEMDBServiceHandler) SaveRecoverRecord(ctx context.Context, in *DBSaveRecoverRecordRequest, out *DBSaveRecoverRecordResponse) error {
	return h.TiEMDBServiceHandler.SaveRecoverRecord(ctx, in, out)
}

func (h *tiEMDBServiceHandler) SaveParametersRecord(ctx context.Context, in *DBSaveParametersRequest, out *DBSaveParametersResponse) error {
	return h.TiEMDBServiceHandler.SaveParametersRecord(ctx, in, out)
}

func (h *tiEMDBServiceHandler) GetCurrentParametersRecord(ctx context.Context, in *DBGetCurrentParametersRequest, out *DBGetCurrentParametersResponse) error {
	return h.TiEMDBServiceHandler.GetCurrentParametersRecord(ctx, in, out)
}

func (h *tiEMDBServiceHandler) CreateTiupTask(ctx context.Context, in *CreateTiupTaskRequest, out *CreateTiupTaskResponse) error {
	return h.TiEMDBServiceHandler.CreateTiupTask(ctx, in, out)
}

func (h *tiEMDBServiceHandler) UpdateTiupTask(ctx context.Context, in *UpdateTiupTaskRequest, out *UpdateTiupTaskResponse) error {
	return h.TiEMDBServiceHandler.UpdateTiupTask(ctx, in, out)
}

func (h *tiEMDBServiceHandler) FindTiupTaskByID(ctx context.Context, in *FindTiupTaskByIDRequest, out *FindTiupTaskByIDResponse) error {
	return h.TiEMDBServiceHandler.FindTiupTaskByID(ctx, in, out)
}

func (h *tiEMDBServiceHandler) GetTiupTaskStatusByBizID(ctx context.Context, in *GetTiupTaskStatusByBizIDRequest, out *GetTiupTaskStatusByBizIDResponse) error {
	return h.TiEMDBServiceHandler.GetTiupTaskStatusByBizID(ctx, in, out)
}

func (h *tiEMDBServiceHandler) CreateFlow(ctx context.Context, in *DBCreateFlowRequest, out *DBCreateFlowResponse) error {
	return h.TiEMDBServiceHandler.CreateFlow(ctx, in, out)
}

func (h *tiEMDBServiceHandler) CreateTask(ctx context.Context, in *DBCreateTaskRequest, out *DBCreateTaskResponse) error {
	return h.TiEMDBServiceHandler.CreateTask(ctx, in, out)
}

func (h *tiEMDBServiceHandler) UpdateFlow(ctx context.Context, in *DBUpdateFlowRequest, out *DBUpdateFlowResponse) error {
	return h.TiEMDBServiceHandler.UpdateFlow(ctx, in, out)
}

func (h *tiEMDBServiceHandler) UpdateTask(ctx context.Context, in *DBUpdateTaskRequest, out *DBUpdateTaskResponse) error {
	return h.TiEMDBServiceHandler.UpdateTask(ctx, in, out)
}

func (h *tiEMDBServiceHandler) LoadFlow(ctx context.Context, in *DBLoadFlowRequest, out *DBLoadFlowResponse) error {
	return h.TiEMDBServiceHandler.LoadFlow(ctx, in, out)
}

func (h *tiEMDBServiceHandler) LoadTask(ctx context.Context, in *DBLoadTaskRequest, out *DBLoadTaskResponse) error {
	return h.TiEMDBServiceHandler.LoadTask(ctx, in, out)
}

func (h *tiEMDBServiceHandler) CreateTransportRecord(ctx context.Context, in *DBCreateTransportRecordRequest, out *DBCreateTransportRecordResponse) error {
	return h.TiEMDBServiceHandler.CreateTransportRecord(ctx, in, out)
}

func (h *tiEMDBServiceHandler) UpdateTransportRecord(ctx context.Context, in *DBUpdateTransportRecordRequest, out *DBUpdateTransportRecordResponse) error {
	return h.TiEMDBServiceHandler.UpdateTransportRecord(ctx, in, out)
}

func (h *tiEMDBServiceHandler) FindTrasnportRecordByID(ctx context.Context, in *DBFindTransportRecordByIDRequest, out *DBFindTransportRecordByIDResponse) error {
	return h.TiEMDBServiceHandler.FindTrasnportRecordByID(ctx, in, out)
}
