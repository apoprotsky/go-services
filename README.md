# Go Service Container

[![Go Report Card](https://goreportcard.com/badge/github.com/apoprotsky/go-services)](https://goreportcard.com/report/github.com/apoprotsky/go-services)

`go-services` library implements `Service Container` pattern.
It creates instances of objects, resolving their dependencies using simple API.

To resolve services dependencies simply add a public Pointer type field in your service structure or
declare method `GoService` which arguments is pointers to other services.

## Usage example

Declaring `db` package

```go
package db

type Service struct {}

func (svc *Service) GetName() string {
    return "Name"
}
```

Declaring `accounts` package

```go
package accounts

import "app/db"

type Account struct {
    Id   uint32
    Name string 
}

type Service struct {
    DBService *db.Service
}

func (svc *Service) GetById(id uint32) *Account {
    return &Account{Id: id, Name: svc.DBService.GetName()}
}
```

Main application package

```go
package main

import (
    "app/accounts"
    "fmt"
    "github.com/apoprotsky/go-services"
)

func main() {
    accountsService := gs.Get((*accounts.Service)(nil)).(*accounts.Service)
    account := accountsService.GetById(1)
    fmt.Printf("%+v\n", account)
}
```

See example directory for more information.
