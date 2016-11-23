# guess-path

Guess the path containing the desired packaged resources.

# Usage

Suppose you have a package with a resource folder located at `your/resources/`.

`guess-path` will help you to get the path to those resources given those inputs:

- a static value registered via ldflags
- the path to the resources relative to the package root `your/resources/`

The resolution is,
- test the provided static value, if non empty, return it without further check.
- test for a workspace path with runtime values `/runtimepath/github.com/you/yourpackage/`.
- test for a vendored path with cwd `vendor/github.com/you/yourpackage/`.

The return value is
- exactly `TheStaticBuildValue`, if non empty
- concatenation of the matched path and the relative path to the resources `okpath/your/resources/`
- empty when no path matched

```go
package yourpackage

import (
  "github.com/mh-cbon/guess-path"
)

var TheStaticBuildValue = ""

func GuessPath() string {
  p := guesspath.Path(
    TheStaticBuildValue,
    "your/pkg/",
    "your/resources/",
  )
  if p == "" {
  	panic("resources not found")
  }
  // p = somewhere/your/resources
  return p
}

func GuessGlob() string {
  p := guesspath.Glob(
    TheStaticBuildValue,
    "your/pkg/",
    "your/resources/",
    "*.tmpl",
  )
  if p==""{
  	panic("resources not found")
  }
  // p = somewhere/your/resources/*.tmpl
  return p
}
```
