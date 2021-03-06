/*******************************************************************************
 * Copyright 2017 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/
package device

import (
	"commons/errors"
	"commons/logger"
	"commons/url"
	"commons/util"
	"messenger"
	"os"
)

const (
	GET             = "GET"
	DELETE          = "DELETE"
	POST            = "POST"
	PUT             = "PUT"
	HTTP_TAG        = "http://"
	SYSTEMCONTAINER = "SYSTEMCONTAINER"
)

type Command interface {
	Restore() error
	Reboot() error
}

type Executor struct{}

var httpExecutor messenger.Command
var systemContainerIP string

func init() {
	httpExecutor = messenger.NewExecutor()
	systemContainerIP = os.Getenv(SYSTEMCONTAINER)
}

func (Executor) Restore() error {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	if len(systemContainerIP) == 0 {
		logger.Logging(logger.ERROR, "sysetm container ip is not found")
		return errors.NotFound{"system container ip"}
	}

	reqUrl := util.MakeSCRequestUrl(systemContainerIP, url.Restore())
	_, _, err := httpExecutor.SendHttpRequest(POST, reqUrl)
	if err != nil {
		logger.Logging(logger.DEBUG, err.Error())
	}
	return err
}

func (Executor) Reboot() error {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	if len(systemContainerIP) == 0 {
		logger.Logging(logger.ERROR, "system container ip is not found")
		return errors.NotFound{"system container ip"}
	}

	reqUrl := util.MakeSCRequestUrl(systemContainerIP, url.Reboot())
	_, _, err := httpExecutor.SendHttpRequest(POST, reqUrl)
	if err != nil {
		logger.Logging(logger.DEBUG, err.Error())
	}
	return err
}
