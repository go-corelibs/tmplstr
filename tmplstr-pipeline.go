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

// Pipeline defines the source text representing a list of Variables which may
// also be piped into another Pipeline instance
type Pipeline struct {
	Root Variables `parser:"@@ ( @@ )*"    json:"variables,omitempty"`
	Pipe *Pipeline `parser:"( Pipe @@ )?"  json:"piped,omitempty"`
}

// Render returns the source text represented by this Pipeline
func (p *Pipeline) Render() (source string) {
	source += p.Root.Render()
	if p.Pipe != nil {
		source += "|"
		source += p.Pipe.Render()
	}
	return
}

func (p *Pipeline) WalkVariables(fn func(variables *Variables) (stop bool)) (stopped bool) {
	if stopped = p.Root.WalkVariables(fn); stopped {
		return
	} else if p.Pipe != nil {
		stopped = p.Pipe.WalkVariables(fn)
	}
	return
}
