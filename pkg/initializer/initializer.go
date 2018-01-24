/*
 * Copyright 2018 the original author or authors.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *  
 *        http://www.apache.org/licenses/LICENSE-2.0
 *  
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package initializer

import (
	"path/filepath"
	"fmt"
	"errors"

	"github.com/projectriff/riff-cli/pkg/options"
	"github.com/projectriff/riff-cli/pkg/initializer/core"
	"github.com/projectriff/riff-cli/pkg/initializer/java"
	"github.com/projectriff/riff-cli/pkg/initializer/python"
	"github.com/projectriff/riff-cli/pkg/initializer/node"
	"github.com/projectriff/riff-cli/pkg/initializer/shell"
)

var supportedExtensions = []string{"js", "java", "py", "sh"}

var languageForFileExtension = map[string]string{
	"sh":   "shell",
	"java": "java",
	"js":   "node",
	"py":   "python",
}

func Java() core.Initializer {
	return core.Initializer{
		Initialize: java.Initialize,
	}
}

func Python() core.Initializer {
	return core.Initializer{
		Initialize: python.Initialize,
	}
}
func Node() core.Initializer {
	return core.Initializer{
		Initialize: node.Initialize,
	}
}
func Shell() core.Initializer {
	return core.Initializer{
		Initialize: shell.Initialize,
	}
}

func Initialize(opts options.InitOptions) error {
	functionPath, err := core.ResolveFunctionPath(opts, "")
	if err != nil {
		return err
	}

	language := languageForFileExtension[filepath.Ext(functionPath)[1:]]

	switch language {
	case "shell":
		Shell().Initialize(opts)
	case "node":
		Node().Initialize(opts)
	case "java":
		fmt.Println("Java resources detected. Use 'riff init java' to specify additional required options")
		return nil
	case "python":
		fmt.Println("Python resources detected. Use 'riff init python' to specify additional required options")
		return nil
	default:
		//TODO: Should never get here
		return errors.New(fmt.Sprintf("unsupported language %s\n", language))
	}
	return nil
}
