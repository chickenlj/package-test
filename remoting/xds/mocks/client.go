/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mocks

import (
	"github.com/stretchr/testify/mock"
)

import (
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"dubbo.apache.org/dubbo-go/v3/registry"
	xdsCommon "dubbo.apache.org/dubbo-go/v3/remoting/xds/common"
	"dubbo.apache.org/dubbo-go/v3/xds/client/resource"
)

// WrappedClientMock is an autogenerated mock type for the WrappedClientMock type
type WrappedClientMock struct {
	mock.Mock
}

func (m *WrappedClientMock) GetDubboGoMetadata() (map[string]string, error) {
	args := m.Called()
	return args.Get(0).(map[string]string), args.Error(1)
}

func (m *WrappedClientMock) Subscribe(svcUniqueName, interfaceName, hostAddr string, lst registry.NotifyListener) error {
	args := m.Called(svcUniqueName, interfaceName, hostAddr, lst)
	return args.Error(0)
}

func (m *WrappedClientMock) UnSubscribe(svcUniqueName string) {

}

func (m *WrappedClientMock) GetRouterConfig(hostAddr string) resource.RouteConfigUpdate {
	args := m.Called(hostAddr)
	return args.Get(0).(resource.RouteConfigUpdate)
}
func (m *WrappedClientMock) GetHostAddrByServiceUniqueKey(serviceUniqueKey string) (string, error) {
	args := m.Called(serviceUniqueKey)
	return args.String(0), args.Error(1)
}
func (m *WrappedClientMock) ChangeInterfaceMap(serviceUniqueKey string, add bool) error {
	args := m.Called(serviceUniqueKey, add)
	return args.Error(0)
}

func (m *WrappedClientMock) GetClusterUpdateIgnoreVersion(hostAddr string) resource.ClusterUpdate {
	args := m.Called(hostAddr)
	return args.Get(0).(resource.ClusterUpdate)
}

func (m *WrappedClientMock) GetHostAddress() xdsCommon.HostAddr {
	args := m.Called()
	return args.Get(0).(xdsCommon.HostAddr)
}

func (m *WrappedClientMock) GetIstioPodIP() string {
	args := m.Called()
	return args.String(0)
}

func (m *WrappedClientMock) MatchRoute(routerConfig resource.RouteConfigUpdate, invocation protocol.Invocation) (*resource.Route, error) {
	args := m.Called(routerConfig, invocation)
	return args.Get(0).(*resource.Route), args.Error(1)
}
