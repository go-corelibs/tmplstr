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

package tmplstr

import (
	"encoding/json"
)

// Tree is a list of Branch instances and is the top of the ParseTemplate
// abstract syntax tree
type Tree []*Branch

// Render returns the source text represented by this Tree
func (t Tree) Render() (source string) {
	for _, content := range t {
		source += content.Render()
	}
	return
}

// Format returns indented JSON output representing this Tree
func (t Tree) Format() (output string) {
	if data, err := json.MarshalIndent(t, "", "  "); err == nil {
		output = string(data)
	}
	return
}

// WalkVariables walks this Tree, calling the given VariablesWalkFn for all
// Variables present
func (t Tree) WalkVariables(fn VariablesWalkFn) (stopped bool) {
	for _, branch := range t {
		if branch.Text == nil && branch.Action != nil {
			if stopped = branch.Action.WalkVariables(fn); stopped {
				return
			}
		}
	}
	return
}
