# Cake 🍰

[![Go Report Card](https://goreportcard.com/badge/github.com/prettyboiiii/cake)](https://goreportcard.com/report/github.com/prettyboiiii/cake)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE) [![Go Reference](https://pkg.go.dev/badge/github.com/prettyboiiii/cake.svg)](https://pkg.go.dev/github.com/prettyboiiii/cake/slice)

Cake is a library providing utility code aims to make the developer experience better and less duplicate code, such as slice methods similar to array methods in Javascript/Typescript and more to come in the future

## Installation

To install the library in your go module, use `go get`:

```shell
go get github.com/prettyboiiii/cake@latest
```

## Features

### Slice

utility methods for Go's slice

#### Synchronous methods

- [Find](slice\Find.go)
- [From](slice\From.go)
- [New](slice\New.go)
- [Push](slice\Push.go)

#### Asynchronous methods

_WIP_

## Usage

Import the slice method into your module

```go
import cake "github.com/prettyboiiii/cake/slice
```

Simple example of how to make a new cake.Slice and use its methods

```go
func main() {
  oldSlice := []int{1,2,3}
  cakes := cake.From(stringSlice)
  cakes = cake.New(1,2,3)
  cakes.Push(9000)

  tooMuchCake := cakes.Find(func(cake int) bool {
    return cake > 9000
  })
}

```

or you can do it like this

```go
func isTooMuchCake(cake int) bool {
  return cake > 9000 // 🍰😳
}


func main() {
  ...
  tooMuchCake := cakes.Find(isTooMuchCake)
}
```

## Contributing

If you'd like to contribute to this project, please follow these guidelines:

- Mention the issue/wanted feature in [New issue](https://github.com/prettyboiiii/cake/issues/new/choose)
- Fork the repository
- Create a new branch for your feature or fix
- (Optional) Bind the branch to the Issue to make it easier to track. Read more [Creating a branch to work on an issue](https://docs.github.com/en/issues/tracking-your-work-with-issues/creating-a-branch-for-an-issue)
- Make your changes
- Submit a pull request
