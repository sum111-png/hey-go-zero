//  Copyright [2020] [hey-go-zero]
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"flag"
	"fmt"

	"hey-go-zero/common/errorx"
	"hey-go-zero/common/middleware"
	"hey-go-zero/service/course/api/internal/config"
	"hey-go-zero/service/course/api/internal/handler"
	"hey-go-zero/service/course/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/course-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	errHandler := errorx.Handler{}
	httpx.SetErrorHandler(errHandler.Handle())

	handler.RegisterHandlers(server, ctx)
	server.Use(middleware.UserCheck)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
