package budoux

import (
	"testing"
)

func BenchmarkGetUnicodeBlockAndFeature(b *testing.B) {
	in := []rune("日本語の文章をいい感じに分割します。")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getUnicodeBlockAndFeature(in, 9)
	}
}

func BenchmarkGetGetFeature(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getFeature("日", "本", "語", "の", "文", "章", "120", "120", "120", "108", "120", "120", "O", "O", "O")
	}
}
