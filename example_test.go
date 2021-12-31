package budoux_test

import (
	"fmt"

	"github.com/sg0hsmt/budoux-go"
	"github.com/sg0hsmt/budoux-go/models"
)

func Example() {
	model := models.DefaultJapaneseModel()
	parser := budoux.New(model, budoux.DefaultThreshold)

	out := parser.Parse("日本語の文章をいい感じに分割します。")

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
