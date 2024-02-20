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

// Action represents a single text or html template action
//
// See: https://pkg.go.dev/text/template#hdr-Actions
type Action struct {
	Open      *string   `parser:"( @StatementOpen "   json:"open,omitempty"`
	Pipelines Pipelines `parser:"  @@"                json:"pipelines,omitempty"`
	Close     *string   `parser:"  @StatementClose )" json:"close,omitempty"`
}

// Render returns the source text represented by this Branch
func (a *Action) Render() (source string) {
	source += *a.Open
	source += a.Pipelines.Render()
	source += *a.Close
	return
}

// WalkVariables walks this Action, calling the given VariablesWalkFn for all
// Variables present
func (a *Action) WalkVariables(fn VariablesWalkFn) (stopped bool) {
	for _, pipeline := range a.Pipelines {
		if stopped = pipeline.WalkVariables(fn); !stopped {
			return
		}
	}
	return
}
