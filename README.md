# BudouX-Go

[![Go Reference](https://pkg.go.dev/badge/github.com/sg0hsmt/budoux-go.svg)](https://pkg.go.dev/github.com/sg0hsmt/budoux-go)
[![Test](https://github.com/sg0hsmt/budoux-go/actions/workflows/test.yaml/badge.svg)](https://github.com/sg0hsmt/budoux-go/actions/workflows/test.yaml)
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

## Requirement

Go 1.11 or later.

## Usage

See [example_test.go](./example_test.go) and [cli](./cmd/).

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
