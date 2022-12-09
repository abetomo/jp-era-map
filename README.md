# jp-era-map

## Example: Variables

```go
package main

import (
    "fmt"
    era "github.com/abetomo/jp-era-map"
)

func main() {
    // 1949
    fmt.Printf("%d\n", era.JpEraMap["S24"])
    // 2012
    fmt.Printf("%d\n", era.JpEraMap["H24"])
}
```

## Example: JpEra()

```go
fmt.Println(JpEra("大正3年"))
fmt.Println(JpEra("R01"))
// Output:
// 1914
// 2019
```
