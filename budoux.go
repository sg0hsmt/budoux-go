// Package budoux is a golang port of [BudouX](https://github.com/google/budoux).
package budoux

// DefaultThreshold default threshold for splitting a sentence.
const DefaultThreshold = 1000

// Model trained machine learning model.
// key (string) is feature of character, value (int) is score of feature.
type Model map[string]int

// Parser machine learning based sentence splitter.
type Parser struct {
	model     Model
	threshold int
}

// New returns budoux-go parser instance.
func New(model Model, threshold int) *Parser {
	return &Parser{
		model:     model,
		threshold: threshold,
	}
}

// Parse returns splitted string slice from input.
func (s *Parser) Parse(in string) []string {
	return ParseWithThreshold(s.model, in, s.threshold)
}
