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
	ctx "context"
	"github.com/golang/mock/gomock"
	"github.com/pingcap-inc/tiem/common/structs"
	"github.com/pingcap-inc/tiem/library/secondparty"
	"github.com/pingcap-inc/tiem/models/cluster/management"
	mock_secondparty_v2 "github.com/pingcap-inc/tiem/test/mocksecondparty_v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestContain(t *testing.T) {
	ports := []int{4000, 4001, 4002}
	got := Contain(ports, 4000)
	assert.Equal(t, got, true)

	got = Contain(ports, 4003)
	assert.Equal(t, got, false)

	hosts := []string{"127.0.0.1", "127.0.0.2"}
	got = Contain(hosts, "127.0.0.1")
	assert.Equal(t, got, true)

	got = Contain(hosts, "127.0.0.3")
	assert.Equal(t, got, false)
}

func TestScaleOutPreCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTiup := mock_secondparty_v2.NewMockSecondPartyService(ctrl)
	secondparty.Manager = mockTiup

	t.Run("normal", func(t *testing.T) {
		mockTiup.EXPECT().ClusterComponentCtl(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
			gomock.Any(), gomock.Any()).Return("{\"enable-placement-rules\": \"true\"}", nil)
		meta := &ClusterMeta{
			Cluster: &management.Cluster{
				Version: "v5.0.0",
			},
			Instances: map[string][]*management.ClusterInstance{
				"PD": {
					{
						HostIP: []string{"127.0.0.1"},
						Ports:  []int32{4000},
					},
				},
			}}
		computes := []structs.ClusterResourceParameterCompute{
			{
				Type: "TiFlash",
			},
		}

		err := ScaleOutPreCheck(ctx.TODO(), meta, computes)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockTiup.EXPECT().ClusterComponentCtl(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
			gomock.Any(), gomock.Any()).Return("{\"enable-placement-rules\": \"false\"}", nil)
		meta := &ClusterMeta{
			Cluster: &management.Cluster{
				Version: "v5.0.0",
			},
			Instances: map[string][]*management.ClusterInstance{
				"PD": {
					{
						HostIP: []string{"127.0.0.1"},
						Ports:  []int32{4000},
					},
				},
			}}
		computes := []structs.ClusterResourceParameterCompute{
			{
				Type: "TiFlash",
			},
		}

		err := ScaleOutPreCheck(ctx.TODO(), meta, computes)
		assert.Error(t, err)
	})

	t.Run("empty", func(t *testing.T) {
		computes := make([]structs.ClusterResourceParameterCompute, 0)
		meta := &ClusterMeta{}
		err := ScaleOutPreCheck(ctx.TODO(), meta, computes)
		assert.Error(t, err)
	})
}

func TestWaitWorkflow(t *testing.T) {
	err := WaitWorkflow("111", 10*time.Second)
	assert.Error(t, err)
	//TODO: test WaitWorkflow
}
