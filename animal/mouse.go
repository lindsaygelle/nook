package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // mouse is the common reference for the Mouse animal type.
    mouse = "Mouse"
)

var (
	// mouseNameAmericanEnglish represents the American English name for the Mouse animal type.
	mouseNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    mouse,
	}
)

var (
	// mouseName contains the names of the Mouse animal type in different languages.
	mouseName = nook.Languages{
		language.AmericanEnglish: mouseNameAmericanEnglish,
	}
)

var (
	// Mouse represents an mouse animal in the Animal Crossing series.
	Mouse = nook.Animal{
		Key:  nook.Key(mouse),
		Name: mouseName,
	}
)
