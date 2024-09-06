// Copyright 2020-2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pluginrpcutil

import (
	"github.com/bufbuild/buf/private/pkg/command"
	"pluginrpc.com/pluginrpc"
)

// NewRunner returns a new pluginrpc.Runner for the command.Runner and program name.
func NewRunner(delegate command.Runner, programName string, programArgs ...string) pluginrpc.Runner {
	return newRunner(delegate, programName, programArgs...)
}

// RunnerProvider provides pluginrpc.Runners for program names and args.
type RunnerProvider interface {
	NewRunner(programName string, programArgs ...string) pluginrpc.Runner
}

// RunnerProviderFunc is a function that implements RunnerProvider.
type RunnerProviderFunc func(programName string, programArgs ...string) pluginrpc.Runner

// NewRunner implements RunnerProvider.
func (r RunnerProviderFunc) NewRunner(programName string, programArgs ...string) pluginrpc.Runner {
	return r(programName, programArgs...)
}

// NewRunnerProvider returns a new RunnerProvider for the command.Runner.
func NewRunnerProvider(delegate command.Runner) RunnerProvider {
	return RunnerProviderFunc(
		func(programName string, programArgs ...string) pluginrpc.Runner {
			return NewRunner(
				delegate,
				programName,
				programArgs...,
			)
		},
	)
}