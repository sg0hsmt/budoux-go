# BudouX-Go

BudouX-Go is a golang port of [BudouX](https://github.com/google/budoux).

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
