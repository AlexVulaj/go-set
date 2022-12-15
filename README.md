# go-set

A simple set implementation for use in Go.

## Usage

```go
import (
	...
	set "github.com/AlexVulaj/go-set"
	...
)

...
// Create a new empty Set of type int
set.NewSet[int]()

// Create a new non-empty Set of type string
set.NewSet[string]("hello", "world")
```