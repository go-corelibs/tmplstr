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

// Grouping represents a grouping Pipeline
//
// Example: in `{{ ident (inner pipeline) }}` the Grouping is the
// `(inner pipeline)` portion
type Grouping struct {
	Open  *string   `parser:"( @GroupOpen"    json:"open,omitempty"`
	Group *Pipeline `parser:"  @@"            json:"group,omitempty"`
	Close *string   `parser:"  @GroupClose )" json:"close,omitempty"`
}

// Render returns the source text represented by this Grouping
func (g Grouping) Render() (source string) {
	source += *g.Open
	source += g.Group.Render()
	source += *g.Close
	return
}
