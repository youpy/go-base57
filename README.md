# go-base57 [![Go](https://github.com/youpy/go-base57/actions/workflows/go.yml/badge.svg)](https://github.com/youpy/go-base57/actions/workflows/go.yml)

Base57 encoder/decoder implementation for using with `github.com/lithammer/shortuuid`. It supports lexicographic ordering and is compatible with the [ruby](https://github.com/sudhirj/shortuuid.rb)/[python](https://github.com/skorokithakis/shortuuid) implementations.

# Usage

```go
package main

import (
    "fmt"

    "github.com/youpy/go-base57"
    "github.com/lithammer/shortuuid/v3"
)

func main() {
    u := shortuuid.NewWithEncoder(base57.New())
}
```

# See also

- https://github.com/skorokithakis/shortuuid/issues/35
