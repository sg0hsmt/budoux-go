package budoux_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sg0hsmt/budoux-go"
	"github.com/sg0hsmt/budoux-go/models"
)

func Example() {
	model := models.DefaultJapaneseModel()

	out := budoux.Parse(model, "日本語の文章をいい感じに分割します。")

	fmt.Printf("%q", out)
	// Output:
	// ["日本語の" "文章を" "いい" "感じに" "分割します。"]
}

func ExampleParse() {
	model := models.DefaultJapaneseModel()

	out := budoux.Parse(model, "日本語の文章をいい感じに分割します。")

	fmt.Printf("%q", out)
	// Output:
	// ["日本語の" "文章を" "いい" "感じに" "分割します。"]
}

func ExampleParse_customModel() {
	// You can use your own custom model.
	customModel := []byte(`{"UW3:を":1001}`)

	var model budoux.Model

	if err := json.Unmarshal(customModel, &model); err != nil {
		log.Fatalln("unmarshal model:", err)
	}

	out := budoux.Parse(model, "日本語の文章をいい感じに分割します。")

	fmt.Printf("%q", out)
	// Output:
	// ["日本語の文章を" "いい感じに分割します。"]
}

func ExampleParseWithThreshold() {
	model := models.DefaultJapaneseModel()

	// Use default threshold (same as budoux.Parse).
	out := budoux.ParseWithThreshold(model, "日本語の文章をいい感じに分割します。", budoux.DefaultThreshold)

	fmt.Printf("%q", out)
	// Output:
	// ["日本語の" "文章を" "いい" "感じに" "分割します。"]
}

func ExampleParseWithThreshold_customThreshold() {
	model := models.DefaultJapaneseModel()

	// If you use a large threshold, will not be split.
	out := budoux.ParseWithThreshold(model, "日本語の文章をいい感じに分割します。", 100000000)

	fmt.Printf("%q", out)
	// Output:
	// ["日本語の文章をいい感じに分割します。"]
}

func ExampleParser() {
	model := models.DefaultJapaneseModel()
	parser := budoux.New(model, budoux.DefaultThreshold)

	out := parser.Parse("日本語の文章をいい感じに分割します。")

	fmt.Printf("%q", out)
	// Output:
	// ["日本語の" "文章を" "いい" "感じに" "分割します。"]
}

func ExampleParser_customModel() {
	// You can use your own custom model.
	customModel := []byte(`{"UW3:を":1001}`)

	var model budoux.Model

	if err := json.Unmarshal(customModel, &model); err != nil {
		log.Fatalln("unmarshal model:", err)
	}

	parser := budoux.New(model, budoux.DefaultThreshold)

	out := parser.Parse("日本語の文章をいい感じに分割します。")

	fmt.Printf("%q", out)
	// Output:
	// ["日本語の文章を" "いい感じに分割します。"]
}

func ExampleParser_customThreshold() {
	model := models.DefaultJapaneseModel()
	parser := budoux.New(model, 100000000)

	out := parser.Parse("日本語の文章をいい感じに分割します。")

	fmt.Printf("%q", out)
	// Output:
	// ["日本語の文章をいい感じに分割します。"]
}
