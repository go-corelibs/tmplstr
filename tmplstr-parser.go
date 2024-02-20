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
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

const (
	gScalarPattern = `[a-zA-Z][_a-zA-Z0-9]*(\.[a-zA-Z][_a-zA-Z0-9]*)*`
)

var (
	gTemplateLexer = lexer.MustSimple([]lexer.SimpleRule{
		{Name: `Range`, Pattern: `range\s+(?:\$` + gScalarPattern + `)(?:\s*,\s*\$` + gScalarPattern + `)?\s+\:=`},
		{Name: `Assignment`, Pattern: `\$` + gScalarPattern + `\s+:??=`},
		{Name: `Ident`, Pattern: `[_a-zA-Z][_a-zA-Z0-9]*`},
		{Name: `Keyword`, Pattern: `[.$]` + gScalarPattern},
		{Name: `Literal`, Pattern: "`[^`]+?`"},
		{Name: `String`, Pattern: `"(?:(\\"|[^"])+?)*"`},
		{Name: `Rune`, Pattern: `'(?:(\\'|[^'])+?)+'`},
		{Name: `Float`, Pattern: `\d+\.\d+`},
		{Name: `Int`, Pattern: `\d+`},
		{Name: `Comment`, Pattern: `/\*(?:.+?)\*/`},
		{Name: `Space`, Pattern: `\s+`},
		{Name: `Pipe`, Pattern: `\|`},
		{Name: `GroupOpen`, Pattern: `\(`},
		{Name: `GroupClose`, Pattern: `\)`},
		{Name: `StatementOpen`, Pattern: `\{\{\-?`},
		{Name: `StatementClose`, Pattern: `\-?\}\}`},
		{Name: `Text`, Pattern: `.+`},
	})
	gTemplateParser = participle.MustBuild[Action](
		participle.Lexer(gTemplateLexer),
		participle.Unquote("Literal", "String", "Rune"),
		participle.UseLookahead(1024),
	)
)
