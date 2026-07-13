package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// newLeaf is the common reference for Animal Crossing: New Leaf.
	newLeaf = "NewLeaf"
)

var (
	// newLeafNameAmericanEnglish represents Animal Crossing: New Leaf's name in American English.
	newLeafNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: New Leaf",
	}
)

var (
	// newLeafName contains the localized names of Animal Crossing: New Leaf.
	newLeafName = nook.Languages{
		language.AmericanEnglish: newLeafNameAmericanEnglish,
	}
)

var (
	// NewLeaf represents Animal Crossing: New Leaf.
	NewLeaf = nook.Game{
		Key:  nook.Key(newLeaf),
		Name: newLeafName,
	}
)
