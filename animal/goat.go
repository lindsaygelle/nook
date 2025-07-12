package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// goat is the common reference for the Goat animal type.
	goat = "Goat"
)

var (
	// goatNameAmericanEnglish represents the Goat animal type's name in American English.
	// TODO: Verify translation accuracy for "Goat"
	goatNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Goat",
	}

	// goatNameCanadianFrench represents the Goat animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Goat"
	goatNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Goat",
	}

	// goatNameDutch represents the Goat animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Goat"
	goatNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Goat",
	}

	// goatNameFrench represents the Goat animal type's name in French.
	// TODO: Verify translation accuracy for "Goat"
	goatNameFrench = nook.Name{
		Language: language.French,
		Value:    "Goat",
	}

	// goatNameGerman represents the Goat animal type's name in German.
	// TODO: Verify translation accuracy for "Goat"
	goatNameGerman = nook.Name{
		Language: language.German,
		Value:    "Goat",
	}

	// goatNameItalian represents the Goat animal type's name in Italian.
	// TODO: Verify translation accuracy for "Goat"
	goatNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Goat",
	}

	// goatNameJapanese represents the Goat animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Goat"
	goatNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Goat",
	}

	// goatNameKorean represents the Goat animal type's name in Korean.
	// TODO: Verify translation accuracy for "Goat"
	goatNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Goat",
	}

	// goatNameLatinAmericanSpanish represents the Goat animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Goat"
	goatNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Goat",
	}

	// goatNameRussian represents the Goat animal type's name in Russian.
	// TODO: Verify translation accuracy for "Goat"
	goatNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Goat",
	}

	// goatNameSimplifiedChinese represents the Goat animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Goat"
	goatNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Goat",
	}

	// goatNameSpanish represents the Goat animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Goat"
	goatNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Goat",
	}

	// goatNameTraditionalChinese represents the Goat animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Goat"
	goatNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Goat",
	}
)

var (
	// goatName contains the localized names of the Goat animal type.
	goatName = nook.Languages{
		language.AmericanEnglish:      goatNameAmericanEnglish,
		language.CanadianFrench:       goatNameCanadianFrench,
		language.Dutch:                goatNameDutch,
		language.French:               goatNameFrench,
		language.German:               goatNameGerman,
		language.Italian:              goatNameItalian,
		language.Japanese:             goatNameJapanese,
		language.Korean:               goatNameKorean,
		language.LatinAmericanSpanish: goatNameLatinAmericanSpanish,
		language.Russian:              goatNameRussian,
		language.SimplifiedChinese:    goatNameSimplifiedChinese,
		language.Spanish:              goatNameSpanish,
		language.TraditionalChinese:   goatNameTraditionalChinese,
	}
)

var (
	// Goat represents the Goat animal type in the Animal Crossing series.
	Goat = nook.Animal{
		Key:  nook.Key(goat),
		Name: goatName,
	}
)
