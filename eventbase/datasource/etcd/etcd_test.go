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

package etcd_test

import (
	"testing"
	"time"

	"github.com/go-chassis/cari/db"
	"github.com/stretchr/testify/assert"
	// support embedded etcd
	_ "github.com/little-cui/etcdadpt/embedded"
	_ "github.com/little-cui/etcdadpt/remote"

	"github.com/apache/servicecomb-service-center/eventbase/datasource/etcd"
	"github.com/apache/servicecomb-service-center/eventbase/test"
)

func TestNewDatasource(t *testing.T) {
	t.Run("create etcd datasource should pass with no error", func(t *testing.T) {
		cfg := &db.Config{
			Kind:    test.Etcd,
			URI:     test.EtcdURI,
			Timeout: 10 * time.Second,
		}
		etcdDatasource, err := etcd.NewDatasource(cfg)
		assert.NoError(t, err)
		assert.NotNil(t, etcdDatasource)
		assert.NotNil(t, etcdDatasource.TaskDao())
		assert.NotNil(t, etcdDatasource.TombstoneDao())
	})
}
