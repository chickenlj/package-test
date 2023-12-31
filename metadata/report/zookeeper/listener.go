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

package zookeeper

import (
	"strings"
	"sync"
)

import (
	gxset "github.com/dubbogo/gost/container/set"
	"github.com/dubbogo/gost/log/logger"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/registry"
	"dubbo.apache.org/dubbo-go/v3/remoting"
	"dubbo.apache.org/dubbo-go/v3/remoting/zookeeper"
)

// CacheListener defines keyListeners and rootPath
type CacheListener struct {
	// key is zkNode Path and value is set of listeners
	keyListeners    sync.Map
	zkEventListener *zookeeper.ZkEventListener
	rootPath        string
}

// NewCacheListener creates a new CacheListener
func NewCacheListener(rootPath string, listener *zookeeper.ZkEventListener) *CacheListener {
	return &CacheListener{zkEventListener: listener, rootPath: rootPath}
}

// AddListener will add a listener if loaded
func (l *CacheListener) AddListener(key string, listener registry.MappingListener) {
	// FIXME do not use Client.ExistW, cause it has a bug(can not watch zk node that do not exist)
	_, _, _, err := l.zkEventListener.Client.Conn.ExistsW(key)
	// reference from https://stackoverflow.com/questions/34018908/golang-why-dont-we-have-a-set-datastructure
	// make a map[your type]struct{} like set in java
	if err != nil {
		return
	}
	listeners, loaded := l.keyListeners.LoadOrStore(key, map[registry.MappingListener]struct{}{listener: {}})
	if loaded {
		listeners.(map[registry.MappingListener]struct{})[listener] = struct{}{}
		l.keyListeners.Store(key, listeners)
	}
}

// RemoveListener will delete a listener if loaded
func (l *CacheListener) RemoveListener(key string, listener registry.MappingListener) {
	listeners, loaded := l.keyListeners.Load(key)
	if loaded {
		delete(listeners.(map[registry.MappingListener]struct{}), listener)
	}
}

// DataChange changes all listeners' event
func (l *CacheListener) DataChange(event remoting.Event) bool {
	if listeners, ok := l.keyListeners.Load(event.Path); ok {
		for listener := range listeners.(map[registry.MappingListener]struct{}) {
			appNames := strings.Split(event.Content, constant.CommaSeparator)
			set := gxset.NewSet()
			for _, e := range appNames {
				set.Add(e)
			}
			err := listener.OnEvent(registry.NewServiceMappingChangedEvent(l.pathToKey(event.Path), set))
			if err != nil {
				logger.Error("Error notify mapping change event.", err)
				return false
			}
		}
		return true
	}
	return false
}

func (l *CacheListener) pathToKey(path string) string {
	if len(path) == 0 {
		return path
	}
	groupKey := strings.Replace(strings.Replace(path, l.rootPath+constant.PathSeparator, "", -1), constant.PathSeparator, constant.DotSeparator, -1)
	return groupKey[strings.Index(groupKey, constant.DotSeparator)+1:]
}
