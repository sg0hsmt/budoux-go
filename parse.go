package budoux

import (
	"sort"
	"strconv"
	"strings"
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

	var buf strings.Builder
	buf.WriteString(string(runes[:3]))

	p1 := "U" // unknown
	p2 := "U" // unknown
	p3 := "U" // unknown

	w1, b1 := getUnicodeBlockAndFeature(runes, 0) // i-3
	w2, b2 := getUnicodeBlockAndFeature(runes, 1) // i-2
	w3, b3 := getUnicodeBlockAndFeature(runes, 2) // i-1
	w4, b4 := getUnicodeBlockAndFeature(runes, 3) // i
	w5, b5 := getUnicodeBlockAndFeature(runes, 4) // i+1

	for i := 3; i < len(runes); i++ {
		w6, b6 := getUnicodeBlockAndFeature(runes, i+2)

		score := getScore(model, w1, w2, w3, w4, w5, w6, b1, b2, b3, b4, b5, b6, p1, p2, p3)

		if score > threshold {
			out = append(out, buf.String())
			buf.Reset()
			buf.Grow(12) // Assume 4 characters
			buf.WriteString(w4)
		} else {
			buf.WriteString(w4)
		}

		p1 = p2
		p2 = p3

		if score > 0 {
			p3 = "B" // positive
		} else {
			p3 = "O" // negative
		}

		w1, b1 = w2, b2
		w2, b2 = w3, b3
		w3, b3 = w4, b4
		w4, b4 = w5, b5
		w5, b5 = w6, b6
	}

	if buf.Len() != 0 {
		out = append(out, buf.String())
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

	switch {
	case i < 10:
		return string(v), "00" + strconv.Itoa(i)
	case i < 100:
		return string(v), "0" + strconv.Itoa(i)
	default:
		return string(v), strconv.Itoa(i)
	}
}

// getScore from features.
func getScore(model Model, w1, w2, w3, w4, w5, w6, b1, b2, b3, b4, b5, b6, p1, p2, p3 string) int {
	score := 0

	// UP is means unigram of previous results.
	score += model["UP1:"+p1]
	score += model["UP2:"+p2]
	score += model["UP3:"+p3]

	// BP is means bigram of previous results.
	score += model["BP1:"+p1+p2]
	score += model["BP2:"+p2+p3]

	// UW is means unigram of words.
	score += model["UW1:"+w1]
	score += model["UW2:"+w2]
	score += model["UW3:"+w3]
	score += model["UW4:"+w4]
	score += model["UW5:"+w5]
	score += model["UW6:"+w6]

	// BW is means bigram of words.
	score += model["BW1:"+w2+w3]
	score += model["BW2:"+w3+w4]
	score += model["BW3:"+w4+w5]

	// TW is means trigram of words.
	score += model["TW1:"+w1+w2+w3]
	score += model["TW2:"+w2+w3+w4]
	score += model["TW3:"+w3+w4+w5]
	score += model["TW4:"+w4+w5+w6]

	// UB is means unigram of unicode blocks.
	score += model["UB1:"+b1]
	score += model["UB2:"+b2]
	score += model["UB3:"+b3]
	score += model["UB4:"+b4]
	score += model["UB5:"+b5]
	score += model["UB6:"+b6]

	// BB is means bigram of unicode blocks.
	score += model["BB1:"+b2+b3]
	score += model["BB2:"+b3+b4]
	score += model["BB3:"+b4+b5]

	// TB is means trigram of unicode blocks.
	score += model["TB1:"+b1+b2+b3]
	score += model["TB2:"+b2+b3+b4]
	score += model["TB3:"+b3+b4+b5]
	score += model["TB4:"+b4+b5+b6]

	// UQ is combination of UP and UB.
	score += model["UQ1:"+p1+b1]
	score += model["UQ2:"+p2+b2]
	score += model["UQ3:"+p3+b3]

	// BQ is combination of UP and BB.
	score += model["BQ1:"+p2+b2+b3]
	score += model["BQ2:"+p2+b3+b4]
	score += model["BQ3:"+p3+b2+b3]
	score += model["BQ4:"+p3+b3+b4]

	// TQ is combination of UP and TB.
	score += model["TQ1:"+p2+b1+b2+b3]
	score += model["TQ2:"+p2+b2+b3+b4]
	score += model["TQ3:"+p3+b1+b2+b3]
	score += model["TQ4:"+p3+b2+b3+b4]

	return score
}
