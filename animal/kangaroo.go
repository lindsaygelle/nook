package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// kangaroo is the common reference for the Kangaroo animal type.
	kangaroo = "Kangaroo"
)

var (
	// kangarooNameAmericanEnglish represents the Kangaroo animal type's name in American English.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kangaroo",
	}

	// kangarooNameCanadianFrench represents the Kangaroo animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kangaroo",
	}

	// kangarooNameDutch represents the Kangaroo animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Kangaroo",
	}

	// kangarooNameFrench represents the Kangaroo animal type's name in French.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameFrench = nook.Name{
		Language: language.French,
		Value:    "Kangaroo",
	}

	// kangarooNameGerman represents the Kangaroo animal type's name in German.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameGerman = nook.Name{
		Language: language.German,
		Value:    "Kangaroo",
	}

	// kangarooNameItalian represents the Kangaroo animal type's name in Italian.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Kangaroo",
	}

	// kangarooNameJapanese represents the Kangaroo animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Kangaroo",
	}

	// kangarooNameKorean represents the Kangaroo animal type's name in Korean.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Kangaroo",
	}

	// kangarooNameLatinAmericanSpanish represents the Kangaroo animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kangaroo",
	}

	// kangarooNameRussian represents the Kangaroo animal type's name in Russian.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Kangaroo",
	}

	// kangarooNameSimplifiedChinese represents the Kangaroo animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Kangaroo",
	}

	// kangarooNameSpanish represents the Kangaroo animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Kangaroo",
	}

	// kangarooNameTraditionalChinese represents the Kangaroo animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Kangaroo"
	kangarooNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Kangaroo",
	}
)

var (
	// kangarooName contains the localized names of the Kangaroo animal type.
	kangarooName = nook.Languages{
		language.AmericanEnglish:      kangarooNameAmericanEnglish,
		language.CanadianFrench:       kangarooNameCanadianFrench,
		language.Dutch:                kangarooNameDutch,
		language.French:               kangarooNameFrench,
		language.German:               kangarooNameGerman,
		language.Italian:              kangarooNameItalian,
		language.Japanese:             kangarooNameJapanese,
		language.Korean:               kangarooNameKorean,
		language.LatinAmericanSpanish: kangarooNameLatinAmericanSpanish,
		language.Russian:              kangarooNameRussian,
		language.SimplifiedChinese:    kangarooNameSimplifiedChinese,
		language.Spanish:              kangarooNameSpanish,
		language.TraditionalChinese:   kangarooNameTraditionalChinese,
	}
)

var (
	// Kangaroo represents the Kangaroo animal type in the Animal Crossing series.
	Kangaroo = nook.Animal{
		Key:  nook.Key(kangaroo),
		Name: kangarooName,
	}
)
