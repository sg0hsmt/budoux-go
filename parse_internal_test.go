package budoux

import "testing"

func TestUnicodeBlockAndFeature(t *testing.T) {
	tbl := []struct {
		in      string
		idx     int
		char    string
		feature string
	}{
		{
			in:      "abc",
			idx:     0,
			char:    "a",
			feature: "001",
		},
		{
			in:      "xyz",
			idx:     2,
			char:    "z",
			feature: "001",
		},
		{
			in:      "out of index",
			idx:     12,
			char:    "",
			feature: "999",
		},
		{
			in:      "あいうえお",
			idx:     0,
			char:    "あ",
			feature: "108",
		},
		{
			in:      "わをん",
			idx:     2,
			char:    "ん",
			feature: "108",
		},
		{
			in:      "安",
			idx:     0,
			char:    "安",
			feature: "120",
		},
		{
			in:      "範囲外アクセス",
			idx:     7,
			char:    "",
			feature: "999",
		},
		{
			in:      "≠",
			idx:     0,
			char:    "≠",
			feature: "079",
		},
	}

	for _, v := range tbl {
		t.Run(v.in, func(t *testing.T) {
			char, feature := getUnicodeBlockAndFeature([]rune(v.in), v.idx)

			if char != v.char {
				t.Errorf("unmatch char, want %q, got %q", v.char, char)
			}

			if feature != v.feature {
				t.Errorf("unmatch feature, want %q, got %q", v.feature, feature)
			}
		})
	}
}
