[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/tmplstr)
[![codecov](https://codecov.io/gh/go-corelibs/tmplstr/graph/badge.svg?token=Tsz0LabXbC)](https://codecov.io/gh/go-corelibs/tmplstr)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/tmplstr)](https://goreportcard.com/report/github.com/go-corelibs/tmplstr)

# tmplstr - text and html template utilities

A collection of utilities for working with text and html template source
files.

# Installation

``` shell
> go get github.com/go-corelibs/tmplstr@latest
```

# Examples

## RemoveTemplateComments

``` go
func main() {
    cleaned := tmplstr.RemoveTemplateComments(`{{ _ "message" /* comment */ }}`)
    // cleaned == `{{ _ "message"  }}`
}
```

## ParseTemplate

``` go
func main() {
    tree, _ := tmplstr.ParseTemplate(`before{{pipeline}}after`)
}
```

From the above example, calling `tree.Format()` would produce the following
JSON text:

``` json
[
  {
    "text": "before"
  },
  {
    "action": {
      "open": "{{",
      "pipelines": [
        {
          "variables": [
            {
              "ident": "pipeline"
            }
          ]
        }
      ],
      "close": "}}"
    }
  },
  {
    "text": "after"
  }
]
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
