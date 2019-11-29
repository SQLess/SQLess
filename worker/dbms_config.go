/*
 * Copyright 2018 The CovenantSQL Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package worker

import (
	"time"

	"github.com/SQLess/SQLess/rpc"
	"github.com/SQLess/SQLess/rpc/mux"
)

var (
	// DefaultMaxReqTimeGap defines max time gap between request and server.
	DefaultMaxReqTimeGap = time.Minute
)

// DBMSConfig defines the local multi-database management system config.
type DBMSConfig struct {
	RootDir          string
	Server           *mux.Server
	DirectServer     *rpc.Server // optional server to provide DBMS service
	MaxReqTimeGap    time.Duration
	OnCreateDatabase func()
}
