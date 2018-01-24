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

package core

import (
	"github.com/projectriff/riff-cli/pkg/options"
	"path/filepath"
	"fmt"
	"github.com/projectriff/riff-cli/pkg/generate"
	"github.com/projectriff/riff-cli/pkg/osutils"
	"errors"
	"github.com/projectriff/riff-cli/pkg/functions"
)

var supportedExtensions = []string{"js", "java", "py", "sh"}

var languageForFileExtension = map[string]string{
	"sh"	:  	"shell",
	"java"	: 	"java",
	"js"	:   "node",
	"py"	: 	"python",
}

type Initializer struct {
	Initialize func(options.InitOptions) error
}

func Initialize(language string, ext string, opts options.InitOptions) error {
	functionPath, err := ResolveFunctionPath(opts, ext)
	if err != nil {
		return err
	}

	if opts.Artifact != "" && languageForFileExtension[filepath.Ext(opts.Artifact)[1:]] != language {
		return errors.New(fmt.Sprintf("language %s conflicts with artifact file extension %s",language, opts.Artifact))
	}


	// Create function resources in function Path
	opts.FunctionName, _ = functions.FunctionNameFromPath(opts.FunctionPath)

	if opts.Input == "" {
		opts.Input = opts.FunctionName
	}

	if opts.Artifact =="" {
		opts.Artifact = filepath.Base(functionPath)
	}

	var protocolForLanguage = map[string]string{
		"shell"	:  	"stdio",
		"java"	: 	"http",
		"js"	:   "http",
		"node"	:   "http",
		"py"	: 	"stdio",
	}

	if opts.Protocol == "" {
		opts.Protocol = protocolForLanguage[language]
	}

	workdir := filepath.Dir(functionPath)

	err = generate.CreateFunction(workdir,language, opts)
	return err
}


//Assumes given file paths have been sanity checked and are valid
func ResolveFunctionPath(options options.InitOptions, ext string) (string, error) {


	absFilePath, err := filepath.Abs(options.FunctionPath)
	if err != nil {
		return "", err
	}

	var resolvedFunctionPath string
	var functionDir string
	var functionFile string
	if osutils.IsDirectory(absFilePath) {
		if options.Artifact == "" {
			functionFile = options.FunctionName
			functionDir = absFilePath
			if ext != "" {
				resolvedFunctionPath = filepath.Join(functionDir, fmt.Sprintf("%s.%s", functionFile, ext))
			} else {
				functionFile, err = searchForFunctionResource(functionDir, options.FunctionName)
				if err != nil {
					return "", err
				}
				resolvedFunctionPath = functionFile
			}
		} else {
			resolvedFunctionPath = filepath.Join(absFilePath, options.Artifact)
		}
	} else {
		resolvedFunctionPath = absFilePath
	}
	if !osutils.FileExists(resolvedFunctionPath) {
		return "", errors.New(fmt.Sprintf("function path %s does not exist", resolvedFunctionPath))
	}
	return resolvedFunctionPath, nil

}

func searchForFunctionResource(dir string, functionName string) (string, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return "", err
	}

	foundFile := ""
	for _, f := range (files) {
		if b := filepath.Base(f); b[0:len(b)-len(filepath.Ext(f))] == functionName {
			for _, e := range supportedExtensions {
				if filepath.Ext(f) == "."+e {
					if foundFile == "" {
						foundFile = f
					} else {
						return "", errors.New(fmt.Sprintf("function file is not unique %s, %s", filepath.Base(foundFile), filepath.Base(f)))
					}
				}
			}
		}

	}

	if foundFile == "" {
		return "", errors.New(fmt.Sprintf("no function file found in path %s", dir))
	}
	return foundFile, nil
}

