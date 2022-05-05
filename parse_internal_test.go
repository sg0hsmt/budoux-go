package budoux

import "testing"

func TestUnicodeBlockAndFeature(t *testing.T) {
	tbl := []struct {
		in      string
		idx     int
		char    string
		size    int
		feature string
	}{
		{
			in:      "abc",
			idx:     0,
			char:    "a",
			size:    1,
			feature: "001",
		},
		{
			in:      "xyz",
			idx:     2,
			char:    "z",
			size:    1,
			feature: "001",
		},
		{
			in:      "out of index",
			idx:     12,
			char:    "",
			size:    0,
			feature: invalidFeature,
		},
		{
			in:      "あいうえお",
			idx:     0,
			char:    "あ",
			size:    3,
			feature: "108",
		},
		{
			in:      "わをん",
			idx:     6,
			char:    "ん",
			size:    3,
			feature: "108",
		},
		{
			in:      "安",
			idx:     0,
			char:    "安",
			size:    3,
			feature: "120",
		},
		{
			in:      "範囲外アクセス",
			idx:     21,
			char:    "",
			size:    0,
			feature: invalidFeature,
		},
		{
			in:      "≠",
			idx:     0,
			char:    "≠",
			size:    3,
			feature: "079",
		},
		{
			in:      "不正",
			idx:     1,
			char:    "",
			size:    1,
			feature: invalidFeature,
		},
	}

	for _, v := range tbl {
		t.Run(v.in, func(t *testing.T) {
			char, size, feature := getUnicodeBlockAndFeature(v.in, v.idx)

			if char != v.char {
				t.Errorf("unmatch char, want %q, got %q", v.char, char)
			}

			if size != v.size {
				t.Errorf("unmatch size, want %d, got %d", v.size, size)
			}

			if feature != v.feature {
				t.Errorf("unmatch feature, want %q, got %q", v.feature, feature)
			}
		})
	}
}
