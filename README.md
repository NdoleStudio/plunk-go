# plunk-go

[![Build](https://github.com/NdoleStudio/plunk-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/plunk-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/plunk-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/plunk-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/plunk-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/plunk-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/plunk-go)](https://goreportcard.com/report/github.com/NdoleStudio/plunk-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/plunk-go)](https://github.com/NdoleStudio/plunk-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/plunk-go?color=brightgreen)](https://github.com/NdoleStudio/plunk-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/plunk-go)](https://pkg.go.dev/github.com/NdoleStudio/plunk-go)


This is an unofficial client library for interacting with the [Plunk API](https://next-wiki.useplunk.com/) using Go.

![Card](https://www.useplunk.com/assets/card.png)


## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its root if it already is):

```sh
go mod init
```

Then, reference plunk-go in a Go program with `import`:

```go
import (
	"github.com/NdoleStudio/plunk-go"
)
```

Run any of the normal `go` commands (`build`/`install`). The Go toolchain will resolve and fetch the plunk-go module automatically. Alternatively, you can also explicitly `go get` the package into a project:

```bash
go get github.com/NdoleStudio/plunk-go
```


## Implemented

- Contacts
    - `POST /contacts`: Create a new contact or update existing (upsert by email)
    - `DELETE /contacts/{id}`: Permanently delete a contact
    - `GET /contacts`: Get a paginated list of contacts with cursor-based pagination
- Track
    - `POST /track`: Track an event for a contact

## Usage

### Initializing the Client

An instance of the client can be created using `plunk.New()`.

```go
package main

import (
	"github.com/NdoleStudio/plunk-go"
)

func main()  {
	plunkClient := plunk.New(plunk.WithSecretKey(/* Secret Key from https://next-app.useplunk.com/settings?tab=general*/))
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
request := &plunk.CreateContactRequest{
    Email: "user@example.com"
    Subscribed: true
    Data: map[string]any{
        "firstName": "John",
        "lastName": "Doe",
        "plan": "premium",
    }
}

contact, response, err := plunkClient.Contacts.Create(context.Background(), request)
if err != nil {
    //handle error
}
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
