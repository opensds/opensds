// Copyright (c) 2017 Huawei Technologies Co., Ltd. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
This module implements a entry into the OpenSDS REST service.

*/

package main

import (
	"github.com/opensds/opensds/pkg/db"
	"github.com/opensds/opensds/pkg/dock"
	. "github.com/opensds/opensds/pkg/utils/config"
	"github.com/opensds/opensds/pkg/utils/constants"
	"github.com/opensds/opensds/pkg/utils/daemon"
	"github.com/opensds/opensds/pkg/utils/logs"
)

func init() {
	// Get the default global configuration.
	def := GetDefaultConfig()

	// Parse some configuration fields from command line.
	flag := &CONF.Flag
	flag.StringVar(&CONF.OsdsDock.ApiEndpoint, "api-endpoint", def.OsdsDock.ApiEndpoint, "Listen endpoint of dock service")
	flag.StringVar(&CONF.OsdsDock.DockType, "dock-type", def.OsdsDock.DockType, "Type of dock service")
	flag.BoolVar(&CONF.OsdsDock.Daemon, "daemon", def.OsdsDock.Daemon, "Run app as a daemon with -daemon=true")
	flag.DurationVar(&CONF.OsdsDock.LogFlushFrequency, "log-flush-frequency", def.OsdsLet.LogFlushFrequency, "Maximum number of seconds between log flushes")

	// Load global configuration from specified config file.
	CONF.Load(constants.OpensdsConfigPath)

	daemon.CheckAndRunDaemon(CONF.OsdsDock.Daemon)
}

func main() {
	// Open OpenSDS dock service log file.
	logs.InitLogs(CONF.OsdsDock.LogFlushFrequency)
	defer logs.FlushLogs()

	// Set up database session.
	db.Init(&CONF.Database)

	// Construct dock module grpc server struct and run dock server process.
	ds := dock.NewDockServer(CONF.OsdsDock.DockType, CONF.OsdsDock.ApiEndpoint)
	if err := ds.Run(); err != nil {
		panic(err)
	}
}
