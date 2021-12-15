/******************************************************************************
 * Copyright (c)  2021 PingCAP, Inc.                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");            *
 * you may not use this file except in compliance with the License.           *
 * You may obtain a copy of the License at                                    *
 *                                                                            *
 * http://www.apache.org/licenses/LICENSE-2.0                                 *
 *                                                                            *
 * Unless required by applicable law or agreed to in writing, software        *
 * distributed under the License is distributed on an "AS IS" BASIS,          *
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.   *
 * See the License for the specific language governing permissions and        *
 * limitations under the License.                                             *
 ******************************************************************************/

package handler

import (
	"bytes"
	"context"
	"github.com/pingcap-inc/tiem/common/constants"
	"github.com/pingcap-inc/tiem/common/structs"
	"github.com/pingcap-inc/tiem/library/common"
	"github.com/pingcap-inc/tiem/library/common/resource-type"
	"github.com/pingcap-inc/tiem/library/framework"
	"github.com/pingcap-inc/tiem/library/knowledge"
	"github.com/pingcap-inc/tiem/library/util/uuidutil"
	resourceManagement "github.com/pingcap-inc/tiem/micro-cluster/resourcemanager/management"
	"github.com/pingcap-inc/tiem/models"
	"github.com/pingcap-inc/tiem/models/cluster/management"
	dbCommon "github.com/pingcap-inc/tiem/models/common"
	"strconv"
	"strings"
	"text/template"
)

type ClusterMeta struct {
	Cluster              *management.Cluster
	Instances            map[string][]*management.ClusterInstance
	NodeExporterPort     int32
	BlackboxExporterPort int32
}

// BuildCluster
// @Description: build cluster from structs.CreateClusterParameter
// @Receiver p
// @Parameter ctx
// @Parameter param
// @return error
func (p *ClusterMeta) BuildCluster(ctx context.Context, param structs.CreateClusterParameter) error {
	p.Cluster = &management.Cluster{
		Entity: dbCommon.Entity{
			TenantId: framework.GetTenantIDFromContext(ctx),
			Status:   string(constants.ClusterInitializing),
		},
		Name:              param.Name,
		DBUser:            param.DBUser,
		DBPassword:        param.DBPassword,
		Type:              param.Type,
		Version:           param.Version,
		TLS:               param.TLS,
		Tags:              param.Tags,
		OwnerId:           framework.GetUserIDFromContext(ctx),
		ParameterGroupID:  param.ParameterGroupID,
		Copies:            param.Copies,
		Exclusive:         param.Exclusive,
		Region:            param.Region,
		CpuArchitecture:   constants.ArchType(param.CpuArchitecture),
		MaintenanceStatus: constants.ClusterMaintenanceNone,
		MaintainWindow:    "",
	}
	_, err := models.GetClusterReaderWriter().Create(ctx, p.Cluster)
	if err == nil {
		framework.LogWithContext(ctx).Infof("create cluster [%s] succeed", p.Cluster.Name)
	} else {
		framework.LogWithContext(ctx).Errorf("create cluster [%s] failed, err : %s", p.Cluster.Name, err.Error())
	}
	return err
}

// AddInstances
// @Description add new instances into cluster topology, then alloc host ip, port and disk for these instances
// @Parameter	computes
// @Return		error
func (p *ClusterMeta) AddInstances(ctx context.Context, computes []structs.ClusterResourceParameterCompute) error {
	if len(computes) <= 0 {
		return framework.NewTiEMError(common.TIEM_PARAMETER_INVALID, "cluster resource parameter is empty!")
	}

	if p.Cluster == nil {
		return framework.NewTiEMError(common.TIEM_UNRECOGNIZED_ERROR, "cluster in meta is nil!")
	}

	if len(p.Instances) == 0 {
		p.Instances = make(map[string][]*management.ClusterInstance)
	}

	for _, compute := range computes {
		for _, item := range compute.Resource {
			for i := 0; i < item.Count; i++ {
				instance := &management.ClusterInstance{
					Entity: dbCommon.Entity{
						TenantId: p.Cluster.TenantId,
						Status:   string(constants.ClusterInstanceInitializing),
					},
					Type:         compute.Type,
					Version:      p.Cluster.Version,
					ClusterID:    p.Cluster.ID,
					CpuCores:     int8(knowledge.ParseCpu(item.Spec)),
					Memory:       int8(knowledge.ParseMemory(item.Spec)),
					Zone:         resource.GetDomainNameFromCode(item.Zone),
					DiskType:     item.DiskType,
					DiskCapacity: int32(item.DiskCapacity),
				}
				p.Instances[compute.Type] = append(p.Instances[compute.Type], instance)
			}
		}
	}
	framework.LogWithContext(ctx).Infof("add new instances into cluster[%s] topology", p.Cluster.Name)
	return nil
}

func (p *ClusterMeta) GenerateInstanceResourceRequirements(instances []*management.ClusterInstance) []resource.AllocRequirement {
	requirements := make([]resource.AllocRequirement, 0)
	for _, instance := range instances {
		portRange := knowledge.GetComponentPortRange(p.Cluster.Type, p.Cluster.Version, instance.Type)
		requirements = append(requirements, resource.AllocRequirement{
			Location: resource.Location{
				Region: p.Cluster.Region,
				Zone:   instance.Zone,
			},
			Require: resource.Requirement{
				Exclusive: p.Cluster.Exclusive,
				PortReq: []resource.PortRequirement{
					{
						Start:   int32(portRange.Start),
						End:     int32(portRange.End),
						PortCnt: int32(portRange.Count),
					},
				},
				DiskReq: resource.DiskRequirement{
					NeedDisk: true,
					Capacity: instance.DiskCapacity,
					DiskType: instance.DiskType,
				},
				ComputeReq: resource.ComputeRequirement{
					CpuCores: int32(instance.CpuCores),
					Memory:   int32(instance.Memory),
				},
			},
			Count: 1,
			HostFilter: resource.Filter{
				Arch: string(p.Cluster.CpuArchitecture),
			},
			Strategy: resource.RandomRack,
		})
	}
	return requirements
}

func (p *ClusterMeta) GenerateGlobalMonitoredPortRequirements() []resource.AllocRequirement {
	requirements := make([]resource.AllocRequirement, 0)

	portRange := knowledge.GetClusterPortRange(p.Cluster.Type, p.Cluster.Version)
	requirements = append(requirements, resource.AllocRequirement{
		Location: resource.Location{Region: p.Cluster.Region},
		Require: resource.Requirement{
			PortReq: []resource.PortRequirement{
				{
					Start:   int32(portRange.Start),
					End:     int32(portRange.End),
					PortCnt: int32(portRange.Count),
				},
			},
			DiskReq:    resource.DiskRequirement{NeedDisk: false},
			ComputeReq: resource.ComputeRequirement{CpuCores: 0, Memory: 0},
		},
		Count: 1,
		HostFilter: resource.Filter{
			Arch: string(p.Cluster.CpuArchitecture),
		},
		Strategy: resource.ClusterPorts,
	})

	return requirements
}

func (p *ClusterMeta) ApplyGlobalPortResource(nodeExporterPort, blackboxExporterPort int32) {
	p.NodeExporterPort = nodeExporterPort
	p.BlackboxExporterPort = blackboxExporterPort
}

func (p *ClusterMeta) ApplyInstanceResource(resource *resource.BatchAllocResponse) {

}

// AllocInstanceResource
// @Description alloc host ip, port and disk for all new instances
// @Return		alloc request id
// @Return		error
func (p *ClusterMeta) AllocInstanceResource(ctx context.Context) (string, error) {
	requestID := uuidutil.GenerateID()
	instances := make([]*management.ClusterInstance, 0)
	for _, components := range p.Instances {
		for _, instance := range components {
			if instance.Status != string(constants.ClusterInstanceInitializing) {
				continue
			}
			instances = append(instances, instance)
		}
	}

	// Alloc instances resource
	resourceManager := resourceManagement.NewResourceManager()
	request := &resource.BatchAllocRequest{
		BatchRequests: []resource.AllocReq{
			{
				Applicant: resource.Applicant{
					HolderId:  p.Cluster.ID,
					RequestId: requestID,
				},
				Requires: p.GenerateInstanceResourceRequirements(instances),
			},
		},
	}
	response, err := resourceManager.AllocResources(ctx, request)
	if err != nil {
		framework.LogWithContext(ctx).Errorf(
			"cluster[%s] alloc instances resource error: %s", p.Cluster.Name, err.Error())
		return "", err
	}
	for i, instance := range instances {
		instance.HostID = response.BatchResults[0].Results[i].HostId
		instance.HostIP = append(instance.HostIP, response.BatchResults[0].Results[i].HostIp)
		instance.Ports = response.BatchResults[0].Results[i].PortRes[0].Ports
		instance.DiskID = response.BatchResults[0].Results[i].DiskRes.DiskId
		instance.DiskPath = response.BatchResults[0].Results[i].DiskRes.Path
	}

	// Alloc monitored ports resource
	request = &resource.BatchAllocRequest{
		BatchRequests: []resource.AllocReq{
			{
				Applicant: resource.Applicant{
					HolderId:  p.Cluster.ID,
					RequestId: requestID,
				},
				Requires: p.GenerateGlobalMonitoredPortRequirements(),
			},
		},
	}
	response, err = resourceManager.AllocResources(ctx, request)
	if err != nil {
		framework.LogWithContext(ctx).Errorf(
			"cluster[%s] alloc monitored ports resource error: %s", p.Cluster.Name, err.Error())
		return "", err
	}

	p.NodeExporterPort = response.BatchResults[0].Results[0].PortRes[0].Ports[0]
	p.BlackboxExporterPort = response.BatchResults[0].Results[0].PortRes[0].Ports[1]

	return requestID, nil
}

// FreedInstanceResource
// @Description return host ip, port and disk for all existing instance
// @Return
// @Return		error
func (p *ClusterMeta) FreedInstanceResource(ctx context.Context) error {
	// todo
	return nil
}

// GenerateTopologyConfig
// @Description generate yaml config based on cluster topology
// @Return		yaml config
// @Return		error
func (p *ClusterMeta) GenerateTopologyConfig(ctx context.Context) (string, error) {
	if p.Cluster == nil || len(p.Instances) == 0 {
		return "", framework.NewTiEMError(common.TIEM_PARAMETER_INVALID, "cluster topology is empty, please check it!")
	}

	t, err := template.New("cluster_topology.yaml").ParseFiles("template/cluster_topology.yaml")
	if err != nil {
		return "", framework.NewTiEMError(common.TIEM_PARAMETER_INVALID, err.Error())
	}

	topology := new(bytes.Buffer)
	if err = t.Execute(topology, p); err != nil {
		return "", framework.NewTiEMError(common.TIEM_UNRECOGNIZED_ERROR, err.Error())
	}
	framework.LogWithContext(ctx).Infof("generate topology config: %s", topology.String())

	return topology.String(), nil
}

// UpdateClusterStatus
// @Description update cluster status
// @Return		error
func (p *ClusterMeta) UpdateClusterStatus(ctx context.Context, status constants.ClusterRunningStatus) error {
	p.Cluster.Status = string(status)
	err := models.GetClusterReaderWriter().UpdateStatus(ctx, p.Cluster.ID, status)

	if err != nil {
		framework.LogWithContext(ctx).Infof("update cluster[%s] status into %s failed", p.Cluster.Name, status)
	} else {
		framework.LogWithContext(ctx).Errorf("update cluster[%s] status into %s succeed", p.Cluster.Name, status)
	}
	return err
}

// GetInstance
// @Description get instance based on instanceID
// @Parameter	instance id (format: ip:port)
// @Return		instance
// @Return		error
func (p *ClusterMeta) GetInstance(ctx context.Context, instanceAddress string) (*management.ClusterInstance, error) {
	host := strings.Split(instanceAddress, ":")
	if len(host) != 2 {
		return nil, framework.NewTiEMError(common.TIEM_PARAMETER_INVALID, "parameter format is wrong")
	}
	port, err := strconv.ParseInt(host[1], 10, 32)
	if err != nil {
		return nil, framework.NewTiEMError(common.TIEM_PARAMETER_INVALID, "parameter format is wrong")
	}

	for _, components := range p.Instances {
		for _, instance := range components {
			if Contain(instance.HostIP, host[0]) && Contain(instance.Ports, int32(port)) {
				return instance, nil
			}
		}
	}
	return nil, framework.NewTiEMError(common.TIEM_INSTANCE_NOT_FOUND, "instance not found")
}

// IsComponentRequired
// @Description judge whether component is required
// @Parameter	component type
// @Return		bool
func (p *ClusterMeta) IsComponentRequired(ctx context.Context, componentType string) bool {
	return knowledge.GetComponentSpec(p.Cluster.Type,
		p.Cluster.Version, componentType).ComponentConstraint.ComponentRequired
}

// DeleteInstance
// @Description delete instance from cluster topology based on instance id
// @Parameter	instance id (format: ip:port)
// @Return		error
func (p *ClusterMeta) DeleteInstance(ctx context.Context, instanceAddress string) error {
	instance, err := p.GetInstance(ctx, instanceAddress)
	if err != nil {
		return err
	}
	// recycle instance resource
	request := &resource.RecycleRequest{
		RecycleReqs: []resource.RecycleRequire{
			{
				RecycleType: resource.RecycleHost,
				HolderID:    instance.ClusterID,
				HostID:      instance.HostID,
				ComputeReq: resource.ComputeRequirement{
					CpuCores: int32(instance.CpuCores),
					Memory:   int32(instance.Memory),
				},
				DiskReq: resource.DiskResource{
					DiskId: instance.DiskID,
				},
				PortReq: []resource.PortResource{
					{
						Ports: instance.Ports,
					},
				},
			},
		},
	}
	resourceManager := resourceManagement.NewResourceManager()
	err = resourceManager.RecycleResources(ctx, request)
	if err != nil {
		return err
	}

	//TODO: delete from db
	return nil
}

// CloneMeta
// @Description: clone meta info from cluster based on create cluster parameter
// @Receiver p
// @Parameter ctx
// @return *ClusterMeta
func (p *ClusterMeta) CloneMeta(ctx context.Context, parameter structs.CreateClusterParameter) (*ClusterMeta, error) {
	// clone cluster info
	cluster := &management.Cluster{
		Name:            parameter.Name,            // user specify (required)
		DBUser:          parameter.DBUser,          // user specify (required)
		DBPassword:      parameter.DBPassword,      // user specify (required)
		Region:          parameter.Region,          // user specify (required)
		Type:            p.Cluster.Type,            // user not specify
		Version:         p.Cluster.Version,         // user specify (option)
		Tags:            p.Cluster.Tags,            // user specify (option)
		TLS:             p.Cluster.TLS,             // user specify (option)
		Copies:          p.Cluster.Copies,          // user specify (option)
		Exclusive:       p.Cluster.Exclusive,       // user specify (option)
		CpuArchitecture: p.Cluster.CpuArchitecture, // user specify (option)
	}
	// if user specify cluster version
	if len(parameter.Version) > 0 {
		if parameter.Version < p.Cluster.Version {
			return nil, framework.NewTiEMError(common.TIEM_CHECK_CLUSTER_VERSION_ERROR,
				"the specified cluster version is less than source cluster version")
		}
		cluster.Version = parameter.Version
	}
	// if user specify cluster tags
	if len(parameter.Tags) > 0 {
		cluster.Tags = parameter.Tags
	}
	// if user specify tls
	if parameter.TLS != p.Cluster.TLS {
		cluster.TLS = parameter.TLS
	}
	// if user specify copies
	if parameter.Copies > 0 {
		cluster.Copies = parameter.Copies
	}
	// if user specify exclusive
	if parameter.Exclusive != p.Cluster.Exclusive {
		cluster.Exclusive = parameter.Exclusive
	}
	// if user specify cpu arch
	if len(parameter.CpuArchitecture) > 0 {
		cluster.CpuArchitecture = constants.ArchType(parameter.CpuArchitecture)
	}
	cluster.Status = string(constants.ClusterInitializing)

	meta := &ClusterMeta{
		Cluster: cluster,
	}
	meta.Instances = make(map[string][]*management.ClusterInstance)
	for componentType, components := range p.Instances {
		for _, instance := range components {
			newInstance := &management.ClusterInstance{
				Type:         instance.Type,
				Zone:         instance.Zone,
				Version:      cluster.Version,
				CpuCores:     instance.CpuCores,
				Memory:       instance.Memory,
				DiskType:     instance.DiskType,
				DiskCapacity: instance.DiskCapacity,
			}
			meta.Instances[componentType] = append(meta.Instances[componentType], newInstance)
		}
	}

	return meta, nil
}

// Save
// @Description save cluster meta into db
// @Return		error
func (p *ClusterMeta) Save(ctx context.Context) error {
	//TODO: write cluster meta into db
	return nil
}

// StartMaintenance
// @Description: try to start a maintenance
// @Receiver p
// @Parameter ctx
// @Parameter maintenanceStatus
// @return error
func (p *ClusterMeta) StartMaintenance(ctx context.Context, maintenanceStatus constants.ClusterMaintenanceStatus) error {
	// deleting will end all maintenance
	err := models.GetClusterReaderWriter().SetMaintenanceStatus(ctx, p.Cluster.ID, maintenanceStatus)

	if err == nil {
		p.Cluster.MaintenanceStatus = maintenanceStatus
	}
	return err
}

// EndMaintenance
// @Description: clear maintenance status after maintenance finished or failed
// @Receiver p
// @Parameter ctx
// @Parameter maintenanceStatus
// @return error
func (p *ClusterMeta) EndMaintenance(ctx context.Context, originStatus constants.ClusterMaintenanceStatus) error {
	err := models.GetClusterReaderWriter().ClearMaintenanceStatus(ctx, p.Cluster.ID, originStatus)

	if err == nil {
		p.Cluster.MaintenanceStatus = constants.ClusterMaintenanceNone
	}
	return err
}

type ComponentAddress struct {
	IP   string
	Port int
}

// GetClusterConnectAddresses
// @Description: Access the TiDB cluster
// @Receiver p
// @return []ComponentAddress
func (p *ClusterMeta) GetClusterConnectAddresses() []ComponentAddress {
	// got all tidb instances, then get connect addresses
	return nil
}

// GetClusterStatusAddress
// @Description: TiDB Server status information reporting.
// @Receiver p
// @return []ComponentAddress
func (p *ClusterMeta) GetClusterStatusAddress() []ComponentAddress {
	//
	return nil
}

// GetPDClientAddresses
// @Description: communication address for PD Servers to connect.
// @Receiver p
// @return []ComponentAddress
func (p *ClusterMeta) GetPDClientAddresses() []ComponentAddress {
	// todo
	return nil
}

// GetMonitorAddresses
// @Description: Prometheus Service communication port
// @Receiver p
// @return []ComponentAddress
func (p *ClusterMeta) GetMonitorAddresses() []ComponentAddress {
	// todo
	return nil
}

// UpdateMeta
// @Description: update cluster meta, include cluster and all instances
// @Receiver p
// @Parameter ctx
// @return error
func (p *ClusterMeta) UpdateMeta(ctx context.Context) error {
	instances := make([]*management.ClusterInstance, 0)
	if p.Instances != nil {
		for _, v := range p.Instances {
			instances = append(instances, v...)
		}
	}
	return models.GetClusterReaderWriter().UpdateMeta(ctx, p.Cluster, instances)
}

// Delete
// @Description: delete cluster
// @Receiver p
// @Parameter ctx
// @return error
func (p *ClusterMeta) Delete(ctx context.Context) error {
	return models.GetClusterReaderWriter().Delete(ctx, p.Cluster.ID)
}

func Get(ctx context.Context, clusterID string) (*ClusterMeta, error) {
	cluster, instances, err := models.GetClusterReaderWriter().GetMeta(ctx, clusterID)

	if err != nil {
		return nil, err
	}

	instancesMap := make(map[string][]*management.ClusterInstance)

	if instances != nil && len(instances) > 0 {
		for _, instance := range instances {
			if existed, ok := instancesMap[instance.Type]; ok {
				instancesMap[instance.Type] = append(existed, instance)
			} else {
				instancesMap[instance.Type] = append(make([]*management.ClusterInstance, 0), instance)
			}
		}
	}
	return &ClusterMeta{
		Cluster:   cluster,
		Instances: instancesMap,
	}, nil
}
