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
	"strconv"
	"unicode/utf8"
)

type Variable struct {
	Assign   *string   `parser:"(  @Assignment" json:"assign,omitempty"`
	Range    *string   `parser:" | @Range"      json:"range,omitempty"`
	Ident    *string   `parser:" | @Ident"      json:"ident,omitempty"`
	Keyword  *string   `parser:" | @Keyword"    json:"keyword,omitempty"`
	Literal  *string   `parser:" | @Literal"    json:"literal,omitempty"`
	String   *string   `parser:" | @String"     json:"string,omitempty"`
	Rune     *string   `parser:" | @Rune"       json:"rune,omitempty"`
	Float    *float64  `parser:" | @Float"      json:"float,omitempty"`
	Int      *int      `parser:" | @Int"        json:"int,omitempty"`
	Space    *string   `parser:" | ( @Space )+" json:"space,omitempty"`
	Comment  *string   `parser:" | @Comment"    json:"comment,omitempty"`
	Grouping *Grouping `parser:" | @@ )"        json:"grouping,omitempty"`
}

// Render returns the source text represented by this Variable
func (v *Variable) Render() (source string) {
	switch {
	case v.Ident != nil:
		return *v.Ident
	case v.Keyword != nil:
		return *v.Keyword
	case v.Literal != nil:
		return "`" + *v.Literal + "`"
	case v.String != nil:
		return strconv.Quote(*v.String)
	case v.Rune != nil:
		r, _ := utf8.DecodeRuneInString(*v.Rune)
		return strconv.QuoteRune(r)
	case v.Float != nil:
		return fmt.Sprintf("%v", *v.Float)
	case v.Int != nil:
		return strconv.Itoa(*v.Int)
	case v.Space != nil:
		return *v.Space
	case v.Comment != nil:
		return *v.Comment
	case v.Assign != nil:
		return *v.Assign
	case v.Range != nil:
		return *v.Range
	case v.Grouping != nil:
		return v.Grouping.Render()
	}
	return
}
