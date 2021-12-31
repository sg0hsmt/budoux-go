package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sg0hsmt/budoux-go"
	"github.com/sg0hsmt/budoux-go/models"
)

func main() {
	var in string
	var path string

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.StringVar(&in, "in", "", "input string")
	flags.StringVar(&path, "model", "", "path of model file (default: internal model)")

	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalln("parse flags:", err)
	}

	var model budoux.Model

	if path == "" {
		model = models.DefaultJapaneseModel()
	} else {
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalln("read model:", err)
		}

		if err := json.Unmarshal(buf, &model); err != nil {
			log.Fatalln("unmarshal model:", err)
		}
	}

	for _, v := range budoux.Parse(model, in) {
		fmt.Println(v)
	}
}
