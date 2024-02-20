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
	"math/rand"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type tSimpleInputOutput struct {
	in  string
	out string
}

type tTextCaseList struct {
	label string
	cases []tSimpleInputOutput
}

func TestTemplating(t *testing.T) {

	Convey("PruneTemplateComments", t, func() {
		checks := []struct {
			label  string
			err    bool
			input  string
			output string
		}{
			{
				label:  "empty",
				err:    false,
				input:  ``,
				output: ``,
			},
			{
				label:  "simple",
				err:    false,
				input:  `before {{ _ "thing" /* comment */ }} after`,
				output: `before {{ _ "thing"  }} after`,
			},
			{
				label:  "actual",
				err:    false,
				input:  `before {{/* comment */}} after`,
				output: `before {{/* comment */}} after`,
			},
		}

		for _, check := range checks {
			Convey(check.label, func() {
				pruned, err := PruneTemplateComments(check.input)
				if check.err {
					So(err, ShouldNotBeNil)
				} else {
					So(err, ShouldBeNil)
				}
				So(pruned, ShouldEqual, check.output)
			})
		}
	})

	Convey("RemoveTemplateComments", t, func() {
		checks := []tTextCaseList{
			{
				label: "not things",
				cases: []tSimpleInputOutput{
					{
						in:  ``,
						out: ``,
					},
					{
						in:  `    `,
						out: `    `,
					},
					{
						in:  `this is some text`,
						out: `this is some text`,
					},
					{
						in:  `{ { is not a thing } }`,
						out: `{ { is not a thing } }`,
					},
					{
						in:  `before {{/* comment */}} after`,
						out: `before {{/* comment */}} after`,
					},
					{
						in:  `before {{-/* comment */}} after`,
						out: `before {{-/* comment */}} after`,
					},
					{
						in:  `before {{/* comment */-}} after`,
						out: `before {{/* comment */-}} after`,
					},
					{
						in:  `before {{-/* comment */-}} after`,
						out: `before {{-/* comment */-}} after`,
					},
				},
			},
			{
				label: "supported things",
				cases: []tSimpleInputOutput{
					{
						in:  `before {{ _ "thing"/* comment */}} after`,
						out: `before {{ _ "thing"}} after`,
					},
					{
						in:  `before {{ _ "thing" /* comment */ $var /*comment*/ }} after`,
						out: `before {{ _ "thing"  $var  }} after`,
					},
					{
						in:  `before {{ _ "thing" /* comment */ $var /*comment*/ | other /*comment*/ .Stuff }} after`,
						out: `before {{ _ "thing"  $var  | other  .Stuff }} after`,
					},
					{
						in:  `before {{ _ "thing /* comment */" $var /*comment*/ | other /*comment*/ .Stuff }} after`,
						out: `before {{ _ "thing /* comment */" $var  | other  .Stuff }} after`,
					},
				},
			},
		}

		for _, check := range checks {
			Convey(check.label, func() {
				for idx, test := range check.cases {
					Convey("(test #"+strconv.Itoa(idx+1)+")", func() {
						So(RemoveTemplateComments(test.in), ShouldEqual, test.out)
					})
				}
			})
		}

	})

}

func BenchmarkPruneTemplateComments(b *testing.B) {
	for i := 0; i < 1000; i++ {
		end := rand.Intn(gPruneTemplateCommentsTestingParagraphLen)
		_, _ = PruneTemplateComments(gPruneTemplateCommentsTestingParagraph[:end])
	}
}

func BenchmarkRemoveTemplateComments(b *testing.B) {
	for i := 0; i < 1000; i++ {
		end := rand.Intn(gPruneTemplateCommentsTestingParagraphLen)
		_ = RemoveTemplateComments(gPruneTemplateCommentsTestingParagraph[:end])
	}
}

const (
	gPruneTemplateCommentsTestingParagraph = `
"quoted {{text}}" escaped \}} and actual }}
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
{{/* actual comment */}} {{ _ "multiline comment" /*
  this comment is indented
  and on multiple lines
*/ arguments... }}
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, {{ sunt in culpa qui officia deserunt mollit anim id est laborum.
/* this is really strange */
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod }}
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`
	gPruneTemplateCommentsTestingParagraphLen = len(gPruneTemplateCommentsTestingParagraph) - 1
)
