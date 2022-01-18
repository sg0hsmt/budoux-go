package budoux_test

import (
	"testing"

	"github.com/sg0hsmt/budoux-go"
	"github.com/sg0hsmt/budoux-go/models"
)

func BenchmarkParse(b *testing.B) {
	model := models.DefaultJapaneseModel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		budoux.Parse(model, "日本語の文章をいい感じに分割します。")
	}
}
