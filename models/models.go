// Package models is trained machine learning model collections.
package models

import "github.com/sg0hsmt/budoux-go"

// DefaultJapaneseModel returns trained japanese model.
func DefaultJapaneseModel() budoux.Model {
	return jaKnbc
}

// DefaultSimplifiedChineseModel returns trained simplified chinese model.
func DefaultSimplifiedChineseModel() budoux.Model {
	return zhHans
}
