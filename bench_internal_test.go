package budoux

import (
	"fmt"
	"testing"
)

func BenchmarkGetUnicodeBlockAndFeature(b *testing.B) {
	in := "日本語の文章をいい感じに分割します。"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getUnicodeBlockAndFeature(in, 27)
	}
}

func BenchmarkGetScore(b *testing.B) {
	model := Model{}
	for i := 0; i < 300; i++ {
		model[fmt.Sprintf("TEST:%03d", i)] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getScore(model, "日", "本", "語", "の", "文", "章", "120", "120", "120", "108", "120", "120", "O", "O", "O")
	}
}
