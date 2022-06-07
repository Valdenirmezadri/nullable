# Nullable SQL Data Types for Golang (GO)

## Features
- 100% [GORM](https://gorm.io/) support
- Can be marshalled into JSON
- Can be unmarshal from JSON
- Convenient Set/Get operation
- Support MySQL, MariaDB, SQLite, and PostgreSQL
- Zero configuration, just use it as normal data type.
- Heavily tested! So you don't have to worry of many bugs :D

## Supported Data Types
- uint

# How to Use?

Very easy! first of all, let's install like normal Go packages

```bash
go get github.com/Valdenirmezadri/nullable
```

## Create a new variable

all you need to have is a basic variable and a nullable variable created with `nullable.NewUint(yourBasicVar)`.

```go
import (
    "fmt"
    "gorm.io/gorm"
    "github.com/Valdenirmezadri/nullable"
)

func main() {
    // Create new
    var theNumber uint64 = 70
    nullableNumber := nullable.NewUint(theNumber)
    fmt.Println(nullableNumber.Get()) // Output: 70

    // Change to another number
    var anotherNumber uint64 = 3306
    nullableNumber.Set(anotherNumber)
    fmt.Println(nullableNumber.Get()) // Output: 3306

    // Change to nil
    nullableNumber.Set(nil)
    fmt.Println(nullableNumber.Get()) // Output: 0
}
```