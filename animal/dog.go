package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // dog is the common reference for the Dog animal type.
    dog = "Dog"
)

var (
	// dogNameAmericanEnglish represents the American English name for the Dog animal type.
	dogNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    dog,
	}
)

var (
	// dogName contains the names of the Dog animal type in different languages.
	dogName = nook.Languages{
		language.AmericanEnglish: dogNameAmericanEnglish,
	}
)

var (
	// Dog represents an dog animal in the Animal Crossing series.
	Dog = nook.Animal{
		Key:  nook.Key(dog),
		Name: dogName,
	}
)
