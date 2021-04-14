# go-base57

Base57 encoder/decoder implementation for using with `github.com/lithammer/shortuuid`. It supports lexicographic ordering and is compatible with the ruby/python implementations.

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
