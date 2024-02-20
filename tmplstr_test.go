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
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//`{{- _ "key" (argv) /* comment */ -}}`,
//`{{- with _ "key" (argv) /* comment */ -}}`,
//`{{- with _ "key" (argv) /* comment */ | other -}}`,
//`{{- with $var := _ "key" (argv) /* comment */ -}}`,
//`{{- with $var := _ "key" (argv) /* comment */ | other -}}`,
//`{{- $var := _ "key" (argv) /* comment */ -}}`,
//`{{- $var = _ "key" (argv) /* comment */ -}}`,
//`{{- _ "key" (argv) /* comment */ | other -}}`,
//`{{- other (_ "key" (argv) /* comment */) -}}`,
//`{{- other (_ "key" (argv) /* comment */) | another (_ "key" (argv) /* comment */) -}}`,
//`{{- other (another (_ "key" (argv) /* comment */)) (_ "key" (argv) /* comment */) -}}`,

func mkStr(input string) *string {
	return &input
}

func mkFloat(input float64) *float64 {
	return &input
}

func mkInt(input int) *int {
	return &input
}

func TestTmplStr(t *testing.T) {

	cases := []struct {
		label string
		input string
		err   bool
		trees Tree
		print string
	}{
		{
			label: "invalid input",
			input: `{{ [  }}`,
			err:   true,
			trees: Tree(nil),
		},

		{
			label: "after invalid input",
			input: `{{ [  }} after`,
			err:   true,
			trees: Tree(nil),
		},

		{
			label: "complete standard template syntax",
			input: `
"{{23 -}} < {{- 45}}"
{{/* a comment */}}
{{pipeline}}
{{if pipeline}} T1 {{end}}
{{if pipeline}} T1 {{else}} T0 {{end}}
{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
{{if pipeline}} T1 {{else}}{{if pipeline}} T0 {{end}}{{end}}
{{range pipeline}} T1 {{end}}
{{range pipeline}} T1 {{else}} T0 {{end}}
{{break}}
{{continue}}
{{template "name"}}
{{template "name" pipeline}}
{{block "name" pipeline}} T1 {{end}}
{{define "name"}} T1 {{end}}
{{with pipeline}} T1 {{end}}
{{with pipeline}} T1 {{else}} T0 {{end}}
{{$variable := pipeline}}
{{$variable = pipeline}}
{{range $index, $element := pipeline}}
`,
			err: false,
			trees: Tree{

				// "{{23 -}} < {{- 45}}"
				{Text: mkStr("\n\"")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Int: mkInt(23)},
						{Space: mkStr(" ")},
					}}},
					Close: mkStr("-}}"),
				}},
				{Text: mkStr(" < ")},
				{Action: &Action{
					Open: mkStr("{{-"),
					Pipelines: Pipelines{{Root: Variables{
						{Space: mkStr(" ")},
						{Int: mkInt(45)},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\"\n")},

				// {{/* a comment */}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Comment: mkStr("/* a comment */")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{pipeline}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{if pipeline}} T1 {{end}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("if")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T1 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("end")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{if pipeline}} T1 {{else}} T0 {{end}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("if")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T1 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("else")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T0 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("end")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("if")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T1 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("else")},
						{Space: mkStr(" ")},
						{Ident: mkStr("if")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T0 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("end")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{if pipeline}} T1 {{else}}{{if pipeline}} T0 {{end}}{{end}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("if")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T1 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("else")},
					}}},
					Close: mkStr("}}"),
				}},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("if")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T0 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("end")},
					}}},
					Close: mkStr("}}"),
				}},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("end")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{range pipeline}} T1 {{end}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("range")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T1 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("end")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{range pipeline}} T1 {{else}} T0 {{end}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("range")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T1 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("else")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T0 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("end")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{break}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("break")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{continue}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("continue")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{template "name"}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("template")},
						{Space: mkStr(" ")},
						{String: mkStr("name")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{template "name" pipeline}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("template")},
						{Space: mkStr(" ")},
						{String: mkStr("name")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{block "name" pipeline}} T1 {{end}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("block")},
						{Space: mkStr(" ")},
						{String: mkStr("name")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T1 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("end")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{define "name"}} T1 {{end}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("define")},
						{Space: mkStr(" ")},
						{String: mkStr("name")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T1 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("end")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{with pipeline}} T1 {{end}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("with")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T1 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("end")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{with pipeline}} T1 {{else}} T0 {{end}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("with")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T1 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("else")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" T0 ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Ident: mkStr("end")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{$variable := pipeline}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Assign: mkStr("$variable :=")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{$variable = pipeline}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Assign: mkStr("$variable =")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},

				// {{range $index, $element := pipeline}}
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{{Root: Variables{
						{Range: mkStr("range $index, $element :=")},
						{Space: mkStr(" ")},
						{Ident: mkStr("pipeline")},
					}}},
					Close: mkStr("}}"),
				}},
				{Text: mkStr("\n")},
			},
		},

		{
			label: "float input",
			input: `{{ 0.1 }}`,
			print: `{{ 0.1 }}`,
			err:   false,
			trees: Tree{
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: Variables{
								{Space: mkStr(" ")},
								{Float: mkFloat(0.1)},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
			},
		},

		{
			label: "int input",
			input: `{{ 10 }}`,
			print: `{{ 10 }}`,
			err:   false,
			trees: Tree{
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: Variables{
								{Space: mkStr(" ")},
								{Int: mkInt(10)},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
			},
		},

		{
			label: "dash input (left)",
			input: `{{- with 10 }}`,
			print: `{{- with 10 }}`,
			err:   false,
			trees: Tree{
				{Action: &Action{
					Open: mkStr("{{-"),
					Pipelines: Pipelines{
						{
							Root: Variables{
								{Space: mkStr(" ")},
								{Ident: mkStr("with")},
								{Space: mkStr(" ")},
								{Int: mkInt(10)},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
			},
		},

		{
			label: "dash input (both)",
			input: `{{- with 10 -}}`,
			print: `{{- with 10 -}}`,
			err:   false,
			trees: Tree{
				{Action: &Action{
					Open: mkStr("{{-"),
					Pipelines: Pipelines{
						{
							Root: Variables{
								{Space: mkStr(" ")},
								{Ident: mkStr("with")},
								{Space: mkStr(" ")},
								{Int: mkInt(10)},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("-}}"),
				}},
			},
		},

		{
			label: "dash input (right)",
			input: `{{ with 10 -}}`,
			print: `{{ with 10 -}}`,
			err:   false,
			trees: Tree{
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: Variables{
								{Space: mkStr(" ")},
								{Ident: mkStr("with")},
								{Space: mkStr(" ")},
								{Int: mkInt(10)},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("-}}"),
				}},
			},
		},

		{
			label: "with input",
			input: `{{ with 10 }}`,
			print: `{{ with 10 }}`,
			err:   false,
			trees: Tree{
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: Variables{
								{Space: mkStr(" ")},
								{Ident: mkStr("with")},
								{Space: mkStr(" ")},
								{Int: mkInt(10)},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
			},
		},

		{
			label: "assign input",
			input: `{{ $name := 10 }}`,
			print: `{{ $name := 10 }}`,
			err:   false,
			trees: Tree{
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: Variables{
								{Space: mkStr(" ")},
								{Assign: mkStr("$name :=")},
								{Space: mkStr(" ")},
								{Int: mkInt(10)},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
			},
		},

		{
			label: "multiple statements",
			input: `{{ $name := 10 }} nonsense {{ other $name }}`,
			print: `{{ $name := 10 }} nonsense {{ other $name }}`,
			err:   false,
			trees: Tree{
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: Variables{
								{Space: mkStr(" ")},
								{Assign: mkStr("$name :=")},
								{Space: mkStr(" ")},
								{Int: mkInt(10)},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(" nonsense ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: Variables{
								{Space: mkStr(" ")},
								{Ident: mkStr("other")},
								{Space: mkStr(" ")},
								{Keyword: mkStr("$name")},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
			},
		},

		{
			label: "prefixing input",
			input: `This is leading text. {{ $name := 10 }}`,
			print: `This is leading text. {{ $name := 10 }}`,
			err:   false,
			trees: Tree{
				{Text: mkStr("This is leading text. ")},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: Variables{
								{Space: mkStr(" ")},
								{Assign: mkStr("$name :=")},
								{Space: mkStr(" ")},
								{Int: mkInt(10)},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
			},
		},

		{
			label: "escaped quoted string",
			input: `{{ "inner \"quoted\"" }}`,
			print: `{{ "inner \"quoted\"" }}`,
			err:   false,
			trees: Tree{
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: []*Variable{
								{Space: mkStr(" ")},
								{String: mkStr("inner \"quoted\"")},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
			},
		},

		{
			label: "literal string",
			input: "{{ `inner \"quoted\"` }}",
			print: "{{ `inner \"quoted\"` }}",
			err:   false,
			trees: Tree{
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: []*Variable{
								{Space: mkStr(" ")},
								{Literal: mkStr("inner \"quoted\"")},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
			},
		},

		{
			label: "rune string",
			input: `{{ '\'' }}`,
			print: `{{ '\'' }}`,
			err:   false,
			trees: Tree{
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: []*Variable{
								{Space: mkStr(" ")},
								{Rune: mkStr("'")},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
			},
		},

		{
			label: "basic",
			input: `This is the first line
		{{ _ "quoted text %q" $variable /* comment */ .Keyword }}
		This is the last line`,
			print: `{{ _ "quoted text %q" $variable /* comment */ .Keyword }}`,
			err:   false,
			trees: Tree{
				{Text: mkStr(`This is the first line
		`)},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: []*Variable{
								{Space: mkStr(" ")},
								{Ident: mkStr("_")},
								{Space: mkStr(" ")},
								{String: mkStr("quoted text %q")},
								{Space: mkStr(" ")},
								{Keyword: mkStr("$variable")},
								{Space: mkStr(" ")},
								{Comment: mkStr("/* comment */")},
								{Space: mkStr(" ")},
								{Keyword: mkStr(".Keyword")},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(`
		This is the last line`)},
			},
		},

		{
			label: "group test",
			input: `This is the first line
		{{ ( _ ) }}
		This is the last line`,
			print: `{{ ( _ ) }}`,
			err:   false,
			trees: Tree{
				{Text: mkStr(`This is the first line
		`)},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: []*Variable{
								{Space: mkStr(" ")},
								{Grouping: &Grouping{
									Open: mkStr("("),
									Group: &Pipeline{
										Root: []*Variable{
											{Space: mkStr(" ")},
											{Ident: mkStr("_")},
											{Space: mkStr(" ")},
										},
									},
									Close: mkStr(")"),
								}},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(`
		This is the last line`)},
			},
		},

		{
			label: "grouped",
			input: `This is the first line
		{{ other (_ "quoted text %q" $variable /* comment */ .Keyword) }}
		This is the last line`,
			print: `{{ other (_ "quoted text %q" $variable /* comment */ .Keyword) }}`,
			err:   false,
			trees: Tree{
				{Text: mkStr(`This is the first line
		`)},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: []*Variable{
								{Space: mkStr(" ")},
								{Ident: mkStr("other")},
								{Space: mkStr(" ")},
								{Grouping: &Grouping{
									Open: mkStr("("),
									Group: &Pipeline{
										Root: []*Variable{
											{Ident: mkStr("_")},
											{Space: mkStr(" ")},
											{String: mkStr("quoted text %q")},
											{Space: mkStr(" ")},
											{Keyword: mkStr("$variable")},
											{Space: mkStr(" ")},
											{Comment: mkStr("/* comment */")},
											{Space: mkStr(" ")},
											{Keyword: mkStr(".Keyword")},
										},
									},
									Close: mkStr(")"),
								}},
								{Space: mkStr(" ")},
							},
						},
					},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(`
		This is the last line`)},
			},
		},

		{
			label: "pipe test",
			input: `This is the first line
		{{ _ "quoted text %q" $variable /* comment */ .Keyword | other }}
		This is the last line`,
			print: `{{ _ "quoted text %q" $variable /* comment */ .Keyword | other }}`,
			err:   false,
			trees: Tree{
				{Text: mkStr(`This is the first line
		`)},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: []*Variable{
								{Space: mkStr(" ")},
								{Ident: mkStr("_")},
								{Space: mkStr(" ")},
								{String: mkStr("quoted text %q")},
								{Space: mkStr(" ")},
								{Keyword: mkStr("$variable")},
								{Space: mkStr(" ")},
								{Comment: mkStr("/* comment */")},
								{Space: mkStr(" ")},
								{Keyword: mkStr(".Keyword")},
								{Space: mkStr(" ")},
							},
							Pipe: &Pipeline{
								Root: []*Variable{
									{Space: mkStr(" ")},
									{Ident: mkStr("other")},
									{Space: mkStr(" ")},
								},
							},
						},
					},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(`
		This is the last line`)},
			},
		},

		{
			label: "pipe to group test",
			input: `This is the first line
		{{ _ "quoted text %q" $variable /* comment */ .Keyword | other ( _ "quoted text %q" $variable /* comment */ .Keyword) }}
		This is the last line`,
			print: `{{ _ "quoted text %q" $variable /* comment */ .Keyword | other ( _ "quoted text %q" $variable /* comment */ .Keyword) }}`,
			err:   false,
			trees: Tree{
				{Text: mkStr(`This is the first line
		`)},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: []*Variable{
								{Space: mkStr(" ")},
								{Ident: mkStr("_")},
								{Space: mkStr(" ")},
								{String: mkStr("quoted text %q")},
								{Space: mkStr(" ")},
								{Keyword: mkStr("$variable")},
								{Space: mkStr(" ")},
								{Comment: mkStr("/* comment */")},
								{Space: mkStr(" ")},
								{Keyword: mkStr(".Keyword")},
								{Space: mkStr(" ")},
							},
							Pipe: &Pipeline{
								Root: []*Variable{
									{Space: mkStr(" ")},
									{Ident: mkStr("other")},
									{Space: mkStr(" ")},
									{Grouping: &Grouping{
										Open: mkStr("("),
										Group: &Pipeline{
											Root: []*Variable{
												{Space: mkStr(" ")},
												{Ident: mkStr("_")},
												{Space: mkStr(" ")},
												{String: mkStr("quoted text %q")},
												{Space: mkStr(" ")},
												{Keyword: mkStr("$variable")},
												{Space: mkStr(" ")},
												{Comment: mkStr("/* comment */")},
												{Space: mkStr(" ")},
												{Keyword: mkStr(".Keyword")},
											},
										},
										Close: mkStr(")"),
									}},
									{Space: mkStr(" ")},
								},
							},
						},
					},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(`
		This is the last line`)},
			},
		},

		{
			label: "pipe within test",
			input: `This is the first line
			{{ _ "quoted text %q" $variable /* comment */ .Keyword | other ( _ "quoted text %q" $variable /* comment */ .Keyword | other) }}
			This is the last line`,
			print: `{{ _ "quoted text %q" $variable /* comment */ .Keyword | other ( _ "quoted text %q" $variable /* comment */ .Keyword | other) }}`,
			err:   false,
			trees: Tree{
				{Text: mkStr(`This is the first line
			`)},
				{Action: &Action{
					Open: mkStr("{{"),
					Pipelines: Pipelines{
						{
							Root: []*Variable{
								{Space: mkStr(" ")},
								{Ident: mkStr("_")},
								{Space: mkStr(" ")},
								{String: mkStr("quoted text %q")},
								{Space: mkStr(" ")},
								{Keyword: mkStr("$variable")},
								{Space: mkStr(" ")},
								{Comment: mkStr("/* comment */")},
								{Space: mkStr(" ")},
								{Keyword: mkStr(".Keyword")},
								{Space: mkStr(" ")},
							},
							Pipe: &Pipeline{
								Root: []*Variable{
									{Space: mkStr(" ")},
									{Ident: mkStr("other")},
									{Space: mkStr(" ")},
									{Grouping: &Grouping{
										Open: mkStr("("),
										Group: &Pipeline{
											Root: []*Variable{
												{Space: mkStr(" ")},
												{Ident: mkStr("_")},
												{Space: mkStr(" ")},
												{String: mkStr("quoted text %q")},
												{Space: mkStr(" ")},
												{Keyword: mkStr("$variable")},
												{Space: mkStr(" ")},
												{Comment: mkStr("/* comment */")},
												{Space: mkStr(" ")},
												{Keyword: mkStr(".Keyword")},
												{Space: mkStr(" ")},
											},
											Pipe: &Pipeline{
												Root: []*Variable{
													{Space: mkStr(" ")},
													{Ident: mkStr("other")},
												},
											},
										},
										Close: mkStr(")"),
									}},
									{Space: mkStr(" ")},
								},
							},
						},
					},
					Close: mkStr("}}"),
				}},
				{Text: mkStr(`
			This is the last line`)},
			},
		},
	}

	Convey("ParseTemplate", t, func() {
		for idx, test := range cases {
			filename := fmt.Sprintf("testing.%d.tmpl", idx+1)
			Convey(filename+" ("+test.label+")", func() {
				trees, err := ParseTemplate(filename, test.input)
				if test.err {
					So(err, ShouldNotBeNil)
				} else {
					So(err, ShouldBeNil)
					So(trees.Render(), ShouldEqual, test.input)
				}
				So(trees, ShouldEqual, test.trees)
			})
		}
	})

	Convey("Branch.Render", t, func() {
		c := &Branch{}
		So(c.Render(), ShouldEqual, "")
	})
}
