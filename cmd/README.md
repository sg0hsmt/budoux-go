# BudouX-Go (CLI)

CLI for BudouX-Go.

## Usage

```
$ go run main.go -h
Usage of main.exe:
  -in string
        input string
  -model string
        path of model file (default: internal model)
```

Split Japanese sentences with internal model.

```console
$ go run main.go -in "日本語の文章をいい感じに分割します。"
日本語の
文章を
いい
感じに
分割します。
```

Split Japanese sentences using the specified model.

```console
$ go run main.go -model ja-knbc.json -in "日本語の文章をいい感じに分割します。"
日本語の
文章を
いい
感じに
分割します。
```
