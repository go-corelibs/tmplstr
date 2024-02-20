// Copyright (c) 2024  The Go-CoreLibs Authors
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

// Package tmplstr provides text and html template utilities
package tmplstr

import (
	clStrings "github.com/go-corelibs/strings"
)

// ParseTemplate uses [github.com/alecthomas/participle/v2] to parse the given
// template file content into an abstract syntax tree. ParseTemplate is
// intended to facilitate extracting contextual information from text or html
// template source text
func ParseTemplate(filename, input string) (trees Tree, err error) {
	for tmp := input[:]; len(tmp) > 0; {
		if before, text, after, found := clStrings.ScanCarve(tmp, "{{", "}}"); found {
			if len(before) > 0 {
				// keep stuff before text
				trees = append(trees, &Branch{Text: &before})
			}
			tmp = after // setup next iteration
			var stmnt *Action
			if stmnt, err = gTemplateParser.ParseString(filename, "{{"+text+"}}"); err != nil {
				return nil, err
			}
			trees = append(trees, &Branch{Action: stmnt})
			continue // move to next iteration
		}
		// keep remainder
		trees = append(trees, &Branch{Text: &tmp})
		break
	}
	return
}
