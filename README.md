# go-flat

go package providing json flattening and unflattening

## 소개 (Introduction)

`go-flat`은 JSON을 평면화(flattening)하고 다시 원래 형태로 되돌리는 기능을 제공하는 Go 패키지입니다. 이 패키지를 사용하면 복잡한 JSON 데이터를 간단하게 다룰 수 있습니다.

`go-flat` is a Go package that provides functionality for flattening and unflattening JSON. With this package, you can easily handle complex JSON data.

## 사용 방법 (Usage)

1. `go-flat` 패키지를 프로젝트에 추가합니다. (Add the `go-flat` package to your project)
2. 필요한 패키지를 import 합니다. (Import the necessary packages)
   ```go
   import "github.com/brian-s-side-project/go-flat"
   ```
3. JSON을 평면화하려면 `go-flat`의 `Flatten` 함수를 사용합니다. (To flatten JSON, use the `Flatten` function from `go-flat`)
   ```go
   flattenedJSON, err := goflat.Flatten(originalJSON, options)
   if err != nil {
        // Error handling
   }
   ```
4. 평면화된 JSON을 다시 원래 형태로 되돌리려면 `go-flat`의 `Unflatten` 함수를 사용합니다. (To unflatten the flattened JSON, use the `Unflatten` function from `go-flat`)
   ```go
   unflattenedJSON, err := goflat.Unflatten(flattenedJSON, options)
   if err != nil {
        // Error handling
   }
   ```

## 예시 (Example)

### 사용 예시 (Usage Example)

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/brian-s-side-project/go-flat"
)

func main() {
    // Original JSON
    originalJSON := `{
        "name": "John Doe",
        "age": 30,
        "address": {
            "street": "123 Main St",
            "city": "New York",
            "state": "NY"
        }
    }`

    // Flatten the JSON
    options := goflat.Options{
        Separator: ".", // Separator to use for flattened keys (default is ".")
        MaxDepth:  0,   // Maximum depth to flatten (0 means flatten all levels, default is 0)
    }
    flattenedJSON, err := goflat.Flatten(originalJSON, options)
    if err != nil {
        fmt.Println("Error flattening JSON:", err)
        return
    }

    fmt.Println("Flattened JSON:", flattenedJSON)

    // Unflatten the JSON
    unflattenedJSON, err := goflat.Unflatten(flattenedJSON, options)
    if err != nil {
        fmt.Println("Error unflattening JSON:", err)
        return
    }

    fmt.Println("Unflattened JSON:", unflattenedJSON)

    // Output:
    // Flattened JSON: {"address.city":"New York","address.state":"NY","address.street":"123 Main St","age":30,"name":"John Doe"}
    // Unflattened JSON: {"address":{"city":"New York","state":"NY","street":"123 Main St"},"age":30,"name":"John Doe"}
}
```
