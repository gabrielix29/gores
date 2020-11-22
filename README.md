# gores

![Travis (.org)](https://img.shields.io/travis/bitcubix/gores)
[![Go Report Card](https://goreportcard.com/badge/github.com/bitcubix/gores)](https://goreportcard.com/report/github.com/bitcubix/gores)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/bitcubix/gores)](https://pkg.go.dev/github.com/bitcubix/gores)

Go package that handles HTTP responses

## installation

`go get github.com/alioygur/gores`

## usage

```go
package main

import (
	"log"
	"net/http"

	"github.com/bitcubix/gores"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
    // JSON response
    http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
        user := User{Name: "user", Email: "user@email.com", Age: 30}
        gores.JSON(w, http.StatusOK, user)
    })
}
```
