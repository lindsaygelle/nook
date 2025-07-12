package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// gorilla is the common reference for the Gorilla animal type.
	gorilla = "Gorilla"
)

var (
	// gorillaNameAmericanEnglish represents the Gorilla animal type's name in American English.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gorilla",
	}

	// gorillaNameCanadianFrench represents the Gorilla animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gorilla",
	}

	// gorillaNameDutch represents the Gorilla animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Gorilla",
	}

	// gorillaNameFrench represents the Gorilla animal type's name in French.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameFrench = nook.Name{
		Language: language.French,
		Value:    "Gorilla",
	}

	// gorillaNameGerman represents the Gorilla animal type's name in German.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameGerman = nook.Name{
		Language: language.German,
		Value:    "Gorilla",
	}

	// gorillaNameItalian represents the Gorilla animal type's name in Italian.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Gorilla",
	}

	// gorillaNameJapanese represents the Gorilla animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Gorilla",
	}

	// gorillaNameKorean represents the Gorilla animal type's name in Korean.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Gorilla",
	}

	// gorillaNameLatinAmericanSpanish represents the Gorilla animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gorilla",
	}

	// gorillaNameRussian represents the Gorilla animal type's name in Russian.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Gorilla",
	}

	// gorillaNameSimplifiedChinese represents the Gorilla animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Gorilla",
	}

	// gorillaNameSpanish represents the Gorilla animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Gorilla",
	}

	// gorillaNameTraditionalChinese represents the Gorilla animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Gorilla"
	gorillaNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Gorilla",
	}
)

var (
	// gorillaName contains the localized names of the Gorilla animal type.
	gorillaName = nook.Languages{
		language.AmericanEnglish:      gorillaNameAmericanEnglish,
		language.CanadianFrench:       gorillaNameCanadianFrench,
		language.Dutch:                gorillaNameDutch,
		language.French:               gorillaNameFrench,
		language.German:               gorillaNameGerman,
		language.Italian:              gorillaNameItalian,
		language.Japanese:             gorillaNameJapanese,
		language.Korean:               gorillaNameKorean,
		language.LatinAmericanSpanish: gorillaNameLatinAmericanSpanish,
		language.Russian:              gorillaNameRussian,
		language.SimplifiedChinese:    gorillaNameSimplifiedChinese,
		language.Spanish:              gorillaNameSpanish,
		language.TraditionalChinese:   gorillaNameTraditionalChinese,
	}
)

var (
	// Gorilla represents the Gorilla animal type in the Animal Crossing series.
	Gorilla = nook.Animal{
		Key:  nook.Key(gorilla),
		Name: gorillaName,
	}
)
