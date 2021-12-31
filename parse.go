package budoux

import (
	"fmt"
	"sort"
)

// Parse returns splitted string slice from input.
// It is shorthand for budoux.ParseWithThreshold(model, in, budoux.DefaultThreshold).
func Parse(model Model, in string) []string {
	return ParseWithThreshold(model, in, DefaultThreshold)
}

// ParseWithThreshold returns splitted string slice from input.
func ParseWithThreshold(model Model, in string, threshold int) []string {
	runes := []rune(in)
	if len(runes) <= 3 {
		return []string{in}
	}

	out := []string{}
	buf := string(runes[:3])

	p1 := "U" // unknown
	p2 := "U" // unknown
	p3 := "U" // unknown

	for i := 3; i < len(runes); i++ {
		w1, b1 := getUnicodeBlockAndFeature(runes, i-3)
		w2, b2 := getUnicodeBlockAndFeature(runes, i-2)
		w3, b3 := getUnicodeBlockAndFeature(runes, i-1)
		w4, b4 := getUnicodeBlockAndFeature(runes, i)
		w5, b5 := getUnicodeBlockAndFeature(runes, i+1)
		w6, b6 := getUnicodeBlockAndFeature(runes, i+2)

		feature := getFeature(w1, w2, w3, w4, w5, w6, b1, b2, b3, b4, b5, b6, p1, p2, p3)
		score := 0

		for _, f := range feature {
			if v, ok := model[f]; ok {
				score += v
			}
		}

		if score > threshold {
			out = append(out, buf)
			buf = w4
		} else {
			buf += w4
		}

		p1 = p2
		p2 = p3

		if score > 0 {
			p3 = "B" // positive
		} else {
			p3 = "O" // negative
		}
	}

	if buf != "" {
		out = append(out, buf)
	}

	return out
}

// getUnicodeBlockAndFeature returns unicode character and block feature from rune slice.
func getUnicodeBlockAndFeature(in []rune, index int) (string, string) {
	if len(in) <= index {
		return "", "999" // out of index.
	}

	v := in[index]
	i := sort.Search(len(unicodeBlocks), func(i int) bool {
		return v < unicodeBlocks[i]
	})

	return string(v), fmt.Sprintf("%03d", i)
}

// getFeature returns feature list.
func getFeature(w1, w2, w3, w4, w5, w6, b1, b2, b3, b4, b5, b6, p1, p2, p3 string) []string {
	return []string{
		// UP is means unigram of previous results.
		"UP1:" + p1,
		"UP2:" + p2,
		"UP3:" + p3,
		// BP is means bigram of previous results.
		"BP1:" + p1 + p2,
		"BP2:" + p2 + p3,
		// UW is means unigram of words.
		"UW1:" + w1,
		"UW2:" + w2,
		"UW3:" + w3,
		"UW4:" + w4,
		"UW5:" + w5,
		"UW6:" + w6,
		// BW is means bigram of words.
		"BW1:" + w2 + w3,
		"BW2:" + w3 + w4,
		"BW3:" + w4 + w5,
		// TW is means trigram of words.
		"TW1:" + w1 + w2 + w3,
		"TW2:" + w2 + w3 + w4,
		"TW3:" + w3 + w4 + w5,
		"TW4:" + w4 + w5 + w6,
		// UB is means unigram of unicode blocks.
		"UB1:" + b1,
		"UB2:" + b2,
		"UB3:" + b3,
		"UB4:" + b4,
		"UB5:" + b5,
		"UB6:" + b6,
		// BB is means bigram of unicode blocks.
		"BB1:" + b2 + b3,
		"BB2:" + b3 + b4,
		"BB3:" + b4 + b5,
		// TB is means trigram of unicode blocks.
		"TB1:" + b1 + b2 + b3,
		"TB2:" + b2 + b3 + b4,
		"TB3:" + b3 + b4 + b5,
		"TB4:" + b4 + b5 + b6,
		// UQ is combination of UP and UB.
		"UQ1:" + p1 + b1,
		"UQ2:" + p2 + b2,
		"UQ3:" + p3 + b3,
		// BQ is combination of UP and BB.
		"BQ1:" + p2 + b2 + b3,
		"BQ2:" + p2 + b3 + b4,
		"BQ3:" + p3 + b2 + b3,
		"BQ4:" + p3 + b3 + b4,
		// TQ is combination of UP and TB.
		"TQ1:" + p2 + b1 + b2 + b3,
		"TQ2:" + p2 + b2 + b3 + b4,
		"TQ3:" + p3 + b1 + b2 + b3,
		"TQ4:" + p3 + b2 + b3 + b4,
	}
}
