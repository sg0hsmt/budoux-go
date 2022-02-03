# BudouX-Go

[![Go Reference](https://pkg.go.dev/badge/github.com/sg0hsmt/budoux-go.svg)](https://pkg.go.dev/github.com/sg0hsmt/budoux-go)
[![Test](https://github.com/sg0hsmt/budoux-go/actions/workflows/test.yaml/badge.svg)](https://github.com/sg0hsmt/budoux-go/actions/workflows/test.yaml)
[![TinyGo](https://github.com/sg0hsmt/budoux-go/actions/workflows/tinygo.yaml/badge.svg)](https://github.com/sg0hsmt/budoux-go/actions/workflows/tinygo.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sg0hsmt/budoux-go)](https://goreportcard.com/report/github.com/sg0hsmt/budoux-go)
[![License](https://img.shields.io/github/license/sg0hsmt/budoux-go.svg)](https://github.com/sg0hsmt/budoux-go/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/sg0hsmt/budoux-go.svg)](https://github.com/sg0hsmt/budoux-go/releases/latest)

BudouX-Go is a golang port of [BudouX](https://github.com/google/budoux) (machine learning powered line break organizer tool).

Note:
This project contains the deliverables of the [BudouX](https://github.com/google/budoux) project.

Note:
BudouX-Go supported plain text only, not supports html inputs.

## Demo

https://sg0hsmt.github.io/budoux-go/

https://go.dev/play/p/PWcZH3eULL6

## Requirement

Go 1.11 or later.

## Usage

Import budoux and models.

```go
import (
    "github.com/sg0hsmt/budoux-go"
    "github.com/sg0hsmt/budoux-go/models"
)
```

Split sentences with internal model.

```go
func Example() {
    model := models.DefaultJapaneseModel()
    words := budoux.Parse(model, "これはテストです。")

    fmt.Printf("%q", words)
    // Output:
    // ["これは" "テストです。"]
}
```

Load model from json file and split sentences using the loaded model.

```go
func Example() {
    var model budoux.Model

    // You can use your own custom model.
    buf, err := ioutil.ReadFile(path_to_json)
    if err != nil {
      log.Fatalln("read model:", err)
    }

    if err := json.Unmarshal(buf, &model); err != nil {
      log.Fatalln("unmarshal model:", err)
    }

    words := budoux.Parse(model, "これはテストです。")

    fmt.Printf("%q", words)
    // Output:
    // ["これは" "テストです。"]
}
```

## Test

```console
go test ./...
```

You can use GitHub Actions locally by [act](https://github.com/nektos/act).

```console
act -j test
```

## Generate model from original BudouX

```console
go generate ./...
```

Note:
Generate model is require Go 1.13 or later.
