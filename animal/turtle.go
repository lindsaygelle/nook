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
	// turtleNameAmericanEnglish represents the Turtle animal type's name in American English.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Turtle",
	}

	// turtleNameCanadianFrench represents the Turtle animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Turtle",
	}

	// turtleNameDutch represents the Turtle animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Turtle",
	}

	// turtleNameFrench represents the Turtle animal type's name in French.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameFrench = nook.Name{
		Language: language.French,
		Value:    "Turtle",
	}

	// turtleNameGerman represents the Turtle animal type's name in German.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameGerman = nook.Name{
		Language: language.German,
		Value:    "Turtle",
	}

	// turtleNameItalian represents the Turtle animal type's name in Italian.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Turtle",
	}

	// turtleNameJapanese represents the Turtle animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Turtle",
	}

	// turtleNameKorean represents the Turtle animal type's name in Korean.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Turtle",
	}

	// turtleNameLatinAmericanSpanish represents the Turtle animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Turtle",
	}

	// turtleNameRussian represents the Turtle animal type's name in Russian.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Turtle",
	}

	// turtleNameSimplifiedChinese represents the Turtle animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Turtle",
	}

	// turtleNameSpanish represents the Turtle animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Turtle",
	}

	// turtleNameTraditionalChinese represents the Turtle animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Turtle"
	turtleNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Turtle",
	}
)

var (
	// turtleName contains the localized names of the Turtle animal type.
	turtleName = nook.Languages{
		language.AmericanEnglish:      turtleNameAmericanEnglish,
		language.CanadianFrench:       turtleNameCanadianFrench,
		language.Dutch:                turtleNameDutch,
		language.French:               turtleNameFrench,
		language.German:               turtleNameGerman,
		language.Italian:              turtleNameItalian,
		language.Japanese:             turtleNameJapanese,
		language.Korean:               turtleNameKorean,
		language.LatinAmericanSpanish: turtleNameLatinAmericanSpanish,
		language.Russian:              turtleNameRussian,
		language.SimplifiedChinese:    turtleNameSimplifiedChinese,
		language.Spanish:              turtleNameSpanish,
		language.TraditionalChinese:   turtleNameTraditionalChinese,
	}
)

var (
	// Turtle represents the Turtle animal type in the Animal Crossing series.
	Turtle = nook.Animal{
		Key:  nook.Key(turtle),
		Name: turtleName,
	}
)
