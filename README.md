# strcgo

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dalikewara/strcgo)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dalikewara/strcgo)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dalikewara/strcgo)
![GitHub license](https://img.shields.io/github/license/dalikewara/strcgo)

**strcgo** provides custom helpers to handle tasks related to `struct`.

## Getting started

### Installation

You can use the `go get` method:

```bash
go get github.com/dalikewara/strcgo
```

### Usage

You can check all available helper functions through the related files above (`get.go`, etc).
For example, if you want to get tag value from e struct field, you can do like this:

```go
type s struct {
    ID   uint64 `json:"id"`
    Name string `json:"name"`
}

tagValue := strcgo.GetTagValueFromField(s{}, "Name", "json")

fmt.Println(tagValue) // output: name
```

## Release

### Changelog

Read at [CHANGELOG.md](https://github.com/dalikewara/strcgo/blob/master/CHANGELOG.md)

### Credits

Copyright &copy; 2024 [Dali Kewara](https://www.dalikewara.com)

### License

[MIT License](https://github.com/dalikewara/strcgo/blob/master/LICENSE)
