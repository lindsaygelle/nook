package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // turtle is the common reference for the Turtle animal type.
    turtle = "Turtle"
)

var (
	// turtleNameAmericanEnglish represents the American English name for the Turtle animal type.
	turtleNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    turtle,
	}
)

var (
	// turtleName contains the names of the Turtle animal type in different languages.
	turtleName = nook.Languages{
		language.AmericanEnglish: turtleNameAmericanEnglish,
	}
)

var (
	// Turtle represents an turtle animal in the Animal Crossing series.
	Turtle = nook.Animal{
		Key:  nook.Key(turtle),
		Name: turtleName,
	}
)
