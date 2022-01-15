# BudouX-Go Demo

Source code of BudouX-Go demo page.
The demo page is build with [go-app](https://github.com/maxence-charriere/go-app).

## Requirement

Go 1.17 or later.

## Test

```console
go test ./...
```

## Run Demo App

Build WebAssembly app before start server.

```console
GOARCH=wasm GOOS=js go build -o web/app.wasm
```

Run server.

```console
go run main.go
```

or build and run server.

```console
go build
./demo
```

## Static Site Generate

Build WebAssembly app before static site generate.

```console
GOARCH=wasm GOOS=js go build -o dist/web/app.wasm
```

Generate static website and output to dist directory.

```console
go run main.go -ssg=dist
```

If deploy to [GitHub Pages](https://pages.github.com/), set repository name to prefix option.

```console
go run main.go -ssg=dist -prefix=REPOSITORY_NAME
```
