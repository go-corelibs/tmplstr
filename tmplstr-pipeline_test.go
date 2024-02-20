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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPipeline(t *testing.T) {
	Convey("WalkVariables", t, func() {

		Convey("stopped", func() {
			tree, err := ParseTemplate("walk-variables", `{{ one two }}{{ many more }}`)
			So(err, ShouldBeNil)
			stopped := tree.WalkVariables(func(variables *Variables) (stop bool) {
				return true
			})
			So(stopped, ShouldBeTrue)
		})

		Convey("piped stopped", func() {
			tree, err := ParseTemplate("walk-variables", `{{ one two|many more }}`)
			So(err, ShouldBeNil)
			stopped := tree.WalkVariables(func(variables *Variables) (stop bool) {
				if len(*variables) > 0 {
					if v := (*variables)[0]; v.Ident != nil {
						stop = *v.Ident == "many"
					}
				}
				return
			})
			So(stopped, ShouldBeTrue)
		})
	})
}
