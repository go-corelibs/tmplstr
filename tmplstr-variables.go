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

// VariablesWalkFn is the function signature for Variable-walking methods. When
// these functions return true, the WalkVariables call is immediately stopped
type VariablesWalkFn func(variables *Variables) (stop bool)

type Variables []*Variable

// Render returns the source text represented by this list of Variables
func (vs Variables) Render() (output string) {
	for _, v := range vs {
		output += v.Render()
	}
	return
}

func (vs Variables) WalkVariables(fn func(variables *Variables) (stop bool)) (stopped bool) {
	if stopped = fn(&vs); stopped {
		return
	}
	for _, v := range vs {
		if v.Grouping != nil {
			if stopped = v.Grouping.Group.WalkVariables(fn); stopped {
				return
			}
		}
	}
	return
}
