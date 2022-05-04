//go:build go1.13
// +build go1.13

//go:generate go run .
//go:generate gofmt -s -w ../unicode_blocks.go
//go:generate gofmt -s -w ../models/ja_knbc.go
//go:generate gofmt -s -w ../models/zh_hans.go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// defaultSrc original BudouX repository.
const defaultSrc = "https://raw.githubusercontent.com/google/budoux/v0.1.0"

var unicodeBlocksTemplate = template.Must(template.New("unicode_blocks").Parse(`
// Code generated by gen/generate.go. DO NOT EDIT.
// Generate from {{.Base}}
// This file is contains the deliverables of the [BudouX](https://github.com/google/budoux) project.
//
// BudouX | Apache License 2.0 | https://github.com/google/budoux/blob/main/LICENSE

package budoux

// unicodeBlocks range of code points block.
var unicodeBlocks = []rune{ {{range .Data}} {{.}}, {{end}} }
`))

var modelsTemplate = template.Must(template.New("models").Parse(`
// Code generated by gen/generate.go. DO NOT EDIT.
// Generate from {{.Base}}
// This file is contains the deliverables of the [BudouX](https://github.com/google/budoux) project.
//
// BudouX | Apache License 2.0 | https://github.com/google/budoux/blob/main/LICENSE

package models

import "github.com/sg0hsmt/budoux-go"

// {{.Name}} trained machine learning model.
var {{.Name}} = budoux.Model{ {{range $key, $value := .Data}} "{{$key}}": {{$value}}, {{end}} }
`))

func main() {
	src := ""
	out := ""

	flag.StringVar(&src, "src", defaultSrc, "original BudouX tree URL.")
	flag.StringVar(&out, "out", "../", "output directory.")
	flag.Parse()

	if _, err := url.Parse(src); err != nil {
		fmt.Printf("invalid src url: %v\n", err)
		os.Exit(1)
	}

	if err := genUnicodeBlocks(src, out); err != nil {
		fmt.Printf("generate unicode blocks: %v\n", err)
		os.Exit(1)
	}

	if err := genModels(src, out); err != nil {
		fmt.Printf("generate models: %v\n", err)
		os.Exit(1)
	}
}

func genUnicodeBlocks(src, out string) error {
	srcURL := src + "/budoux/unicode_blocks.json"

	body, dlErr := download(srcURL)
	if dlErr != nil {
		return fmt.Errorf("download: %w", dlErr)
	}

	inventory := struct {
		Base string
		Data []int
	}{
		Base: srcURL,
		Data: []int{},
	}

	if err := json.Unmarshal(body, &inventory.Data); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	file, fileErr := os.Create(filepath.Join(out, "unicode_blocks.go"))
	if fileErr != nil {
		return fmt.Errorf("open file: %w", fileErr)
	}
	defer file.Close()

	if err := unicodeBlocksTemplate.Execute(file, inventory); err != nil {
		return fmt.Errorf("execute: %w", err)
	}

	return nil
}

func genModels(src, out string) error {
	tbl := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "jaKnbc",
			in:   src + "/budoux/models/ja-knbc.json",
			out:  filepath.Join(out, "models", "ja_knbc.go"),
		},
		{
			name: "zhHans",
			in:   src + "/budoux/models/zh-hans.json",
			out:  filepath.Join(out, "models", "zh_hans.go"),
		},
	}

	for _, v := range tbl {
		if err := genLangModel(v.name, v.in, v.out); err != nil {
			return fmt.Errorf("generate %s model: %w", v.name, err)
		}
	}

	return nil
}

func genLangModel(name, srcURL, outPath string) error {
	body, dlErr := download(srcURL)
	if dlErr != nil {
		return fmt.Errorf("download: %w", dlErr)
	}

	inventory := struct {
		Name string
		Base string
		Data map[string]int
	}{
		Name: name,
		Base: srcURL,
		Data: map[string]int{},
	}

	if err := json.Unmarshal(body, &inventory.Data); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	file, fileErr := os.Create(outPath)
	if fileErr != nil {
		return fmt.Errorf("open file: %w", fileErr)
	}
	defer file.Close()

	if err := modelsTemplate.Execute(file, inventory); err != nil {
		return fmt.Errorf("execute: %w", err)
	}

	return nil
}

func download(src string) ([]byte, error) {
	resp, err := http.Get(src)
	if err != nil {
		return nil, fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status: %d %s",
			resp.StatusCode, strings.ToLower(http.StatusText(resp.StatusCode)))
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}

	return buf, nil
}
