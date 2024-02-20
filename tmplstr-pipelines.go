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

// Pipelines is a list of Pipeline instances
type Pipelines []*Pipeline

// Render returns the source text represented by this list of Pipelines
func (p Pipelines) Render() (source string) {
	for _, pipe := range p {
		source += pipe.Render()
	}
	return
}
