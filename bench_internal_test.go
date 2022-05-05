package budoux

import (
	"testing"
)

func BenchmarkGetUnicodeBlockAndFeature(b *testing.B) {
	in := "日本語の文章をいい感じに分割します。"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getUnicodeBlockAndFeature(in, 27)
	}
}
