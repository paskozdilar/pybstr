# pybstr

Package pybstr provides functions to encode and decode Python-style byte
strings in Go.

## Encode

```go
package main

import (
    "fmt"

    "github.com/paskozdilar/pybstr"
)

func main() {
    data := []byte("Hello, World!")
    encoded := pybstr.Encode(data)
    fmt.Println(encoded) // Output: b'Hello, World!'
}
```

## Decode

```go
package main

import (
    "fmt"
    "log"

    "github.com/paskozdilar/pybstr"
)

func main() {
    s := "b'Hello, World!'"
    decoded, err := pybstr.Decode(s)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(decoded)) // Output: Hello, World!
}
```
