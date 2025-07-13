package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// frog is the common reference for the Frog animal type.
	frog = "Frog"
)

var (
	// frogNameAmericanEnglish represents the Frog animal type's name in American English.
	// TODO: Verify translation accuracy for "Frog"
	frogNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Frog",
	}

	// frogNameCanadianFrench represents the Frog animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Frog"
	frogNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Frog",
	}

	// frogNameDutch represents the Frog animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Frog"
	frogNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Frog",
	}

	// frogNameFrench represents the Frog animal type's name in French.
	// TODO: Verify translation accuracy for "Frog"
	frogNameFrench = nook.Name{
		Language: language.French,
		Value:    "Frog",
	}

	// frogNameGerman represents the Frog animal type's name in German.
	// TODO: Verify translation accuracy for "Frog"
	frogNameGerman = nook.Name{
		Language: language.German,
		Value:    "Frog",
	}

	// frogNameItalian represents the Frog animal type's name in Italian.
	// TODO: Verify translation accuracy for "Frog"
	frogNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Frog",
	}

	// frogNameJapanese represents the Frog animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Frog"
	frogNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Frog",
	}

	// frogNameKorean represents the Frog animal type's name in Korean.
	// TODO: Verify translation accuracy for "Frog"
	frogNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Frog",
	}

	// frogNameLatinAmericanSpanish represents the Frog animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Frog"
	frogNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Frog",
	}

	// frogNameRussian represents the Frog animal type's name in Russian.
	// TODO: Verify translation accuracy for "Frog"
	frogNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Frog",
	}

	// frogNameSimplifiedChinese represents the Frog animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Frog"
	frogNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Frog",
	}

	// frogNameSpanish represents the Frog animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Frog"
	frogNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Frog",
	}

	// frogNameTraditionalChinese represents the Frog animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Frog"
	frogNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Frog",
	}
)

var (
	// frogName contains the localized names of the Frog animal type.
	frogName = nook.Languages{
		language.AmericanEnglish:      frogNameAmericanEnglish,
		language.CanadianFrench:       frogNameCanadianFrench,
		language.Dutch:                frogNameDutch,
		language.French:               frogNameFrench,
		language.German:               frogNameGerman,
		language.Italian:              frogNameItalian,
		language.Japanese:             frogNameJapanese,
		language.Korean:               frogNameKorean,
		language.LatinAmericanSpanish: frogNameLatinAmericanSpanish,
		language.Russian:              frogNameRussian,
		language.SimplifiedChinese:    frogNameSimplifiedChinese,
		language.Spanish:              frogNameSpanish,
		language.TraditionalChinese:   frogNameTraditionalChinese,
	}
)

var (
	// Frog represents the Frog animal type in the Animal Crossing series.
	Frog = nook.Animal{
		Key:  nook.Key(frog),
		Name: frogName,
	}
)
