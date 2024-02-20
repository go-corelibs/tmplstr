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
	clStrings "github.com/go-corelibs/strings"
)

// PruneTemplateComments is like RemoveTemplateComments with a very
// fundamental difference, instead of using ScanCarve and ScanBothCarve,
// PruneTemplateComments uses ParseTemplate, walks the Variables in the
// Tree, setting all Variable.Comment values to nil and finally returning the
// rendered results of the modified Tree. If PruneTemplateComments fails to
// ParseTemplate, the error is returned and pruned is empty
//
// RemoveTemplateComments is much faster at this task and PruneTemplateComments
// is more-or-less an example of using ParseTemplate to modify the template
// source text programmatically
//
// Benchmark Comparison:
//
//	PruneTemplateComments-4   	1000000000	         0.031750 ns/op
//	RemoveTemplateComments-4   	1000000000	         0.004838 ns/op
//	(both implementations have 0 B/op and 0 allocs/op)
func PruneTemplateComments(input string) (pruned string, err error) {
	var tree Tree
	if tree, err = ParseTemplate("prune-template-comments.tmpl", input); err == nil {
		for _, branch := range tree {
			if branch.Text != nil {
				pruned += *branch.Text
				continue
			} else if branch.Action != nil {
				branch.Action.WalkVariables(func(variables *Variables) (stop bool) {
					if len(*variables) > 1 {
						// more than just a valid template comment, prune...
						for _, v := range *variables {
							if v.Comment != nil {
								v.Comment = nil
							}
						}
					}
					return
				})
				pruned += branch.Action.Render()
			}
		}
	}
	return
}

// RemoveTemplateComments removes all C-style block comments from within
// template pipelines, preserving escaped and quoted comments, using
// [github.com/go-corelibs/strings] ScanCarve and ScanBothCarve making
// RemoveTemplateComments very fast compared to PruneTemplateComments
func RemoveTemplateComments(input string) (cleaned string) {

	for temp := input[:]; ; {
		if beforePipelines, pipelines, afterPipelines, foundPipelines := clStrings.ScanCarve(temp, "{{", "}}"); foundPipelines {
			cleaned += beforePipelines
			cleaned += "{{"

			var endingDashed bool
			if last := len(pipelines) - 1; last > 0 {
				if endingDashed = pipelines[last] == '-'; endingDashed {
					// do this before the prefixing dash so that `last` is correct
					pipelines = pipelines[:last]
				}
				if pipelines[0] == '-' {
					cleaned += "-"
					pipelines = pipelines[1:]
				}
			}

			var inner string
			innerPipelines := pipelines[:]
			for {
				if beforeComment, _, afterComment, foundComment := clStrings.ScanBothCarve(innerPipelines, "/*", "*/"); foundComment {
					inner += beforeComment
					innerPipelines = afterComment
					continue
				} else {
					inner += beforeComment
				}
				break
			}
			if clStrings.Empty(inner) {
				cleaned += pipelines
			} else {
				cleaned += inner
			}

			if endingDashed {
				cleaned += "-"
			}
			cleaned += "}}"

			temp = afterPipelines
			continue
		} else {
			cleaned += beforePipelines
		}
		break
	}

	return
}
