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

// Package commons/logger implements log stream which is used
// on Pharos Node.
package logger

import (
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

var loggers [3]*log.Logger
var logFlag int

const (
	INFO = iota
	DEBUG
	ERROR
)

// Initialize package global value.
func init() {
	logFlag = log.Ldate | log.Ltime
	loggers[INFO] = log.New(os.Stdout, "[INFO][NODE]", logFlag)
	loggers[DEBUG] = log.New(os.Stdout, "[DEBUG][NODE]", logFlag)
	loggers[ERROR] = log.New(os.Stdout, "[ERROR][NODE]", logFlag)
}

// Print log stream oh standard output with file name and function name, line.
func Logging(level int, msgs ...string) {
	pc, file, line, _ := runtime.Caller(1)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	funcName := parts[pl-1]

	packageName := ""
	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	loggers[level].Println(packageName, fileName, funcName, ":", strconv.Itoa(line), msgs)
}
