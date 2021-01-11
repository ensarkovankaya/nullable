[![Build Status](https://travis-ci.com/kernle32dll/nullable.svg?branch=master)](https://travis-ci.com/kernle32dll/nullable)
[![GoDoc](https://godoc.org/github.com/allaboutapps/nullable?status.svg)](http://godoc.org/github.com/allaboutapps/nullable)
[![Go Report Card](https://goreportcard.com/badge/github.com/allaboutapps/nullable)](https://goreportcard.com/report/github.com/allaboutapps/nullable)
[![codecov](https://codecov.io/gh/kernle32dll/nullable/branch/master/graph/badge.svg)](https://codecov.io/gh/kernle32dll/nullable)

# nullable

Provides ability to determine if a json key has been set to null or not provided.
Inspired by [How to determine if a JSON key has been set to null or not provided](https://www.calhoun.io/how-to-determine-if-a-json-key-has-been-set-to-null-or-not-provided/) article by **Jon Calhoun**

# Install

To get the package, execute:

```bash
go get github.com/allaboutapps/nullable
```

To import this package, add the following line to your code:

```bash
import "github.com/allaboutapps/nullable"
```

Refer to it as `nullable`.

For more details, see the API documentation.

# Example

```go
import "github.com/allaboutapps/nullable"

func main() {
	usr := struct {
		Name nullable.String `json:"name"`
	}{}

	data := []byte("{}")
	if err := json.Unmarshal(data, &usr); err != nil {
		log.Fatalf("unmarshaling failed: %s\n", err)
	}

	fmt.Println(usr.Name.Present) // false
	fmt.Println(usr.Name.Valid)   // false
	fmt.Println(usr.Name.Value)   // ""

	data = []byte("{\"name\":null}")
	if err := json.Unmarshal(data, &usr); err != nil {
		log.Fatalf("unmarshaling failed: %s\n", err)
	}

	fmt.Println(usr.Name.Present) // true
	fmt.Println(usr.Name.Valid)   // false
	fmt.Println(usr.Name.Value)   // ""

	data = []byte("{\"name\":\"John\"}")
	if err := json.Unmarshal(data, &usr); err != nil {
		log.Fatalf("unmarshaling failed: %s\n", err)
	}

	fmt.Println(usr.Name.Present) // true
	fmt.Println(usr.Name.Valid)   // true
	fmt.Println(usr.Name.Value)   // "John"
}
```

# go-swagger usage example

```yaml
definitions:
  NullableInt:
    type: integer
    example: 1234
    x-go-type:
      import:
        package: github.com/allaboutapps/nullable
      type: Int
  NullableInt16:
    type: integer
    example: 1234
    x-go-type:
      import:
        package: github.com/allaboutapps/nullable
      type: Int16
  NullableFloat:
    type: number
    format: float
    example: 1.5
    x-go-type:
      import:
        package: github.com/allaboutapps/nullable
      type: Float32
  NullableFloat64:
    type: number
    format: float
    example: 1.5
    x-go-type:
      import:
        package: github.com/allaboutapps/nullable
      type: Float64
  NullableString:
    type: string
    example: example
    x-go-type:
      import:
        package: github.com/allaboutapps/nullable
      type: String
```

# Contributing

Please feel free to submit issues, fork the repository and send pull requests!
