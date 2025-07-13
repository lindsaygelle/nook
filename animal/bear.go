package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// bear is the common reference for the Bear animal type.
	bear = "Bear"
)

var (
	// bearNameAmericanEnglish represents the Bear animal type's name in American English.
	// TODO: Verify translation accuracy for "Bear"
	bearNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bear",
	}

	// bearNameCanadianFrench represents the Bear animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Bear"
	bearNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bear",
	}

	// bearNameDutch represents the Bear animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Bear"
	bearNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Bear",
	}

	// bearNameFrench represents the Bear animal type's name in French.
	// TODO: Verify translation accuracy for "Bear"
	bearNameFrench = nook.Name{
		Language: language.French,
		Value:    "Bear",
	}

	// bearNameGerman represents the Bear animal type's name in German.
	// TODO: Verify translation accuracy for "Bear"
	bearNameGerman = nook.Name{
		Language: language.German,
		Value:    "Bear",
	}

	// bearNameItalian represents the Bear animal type's name in Italian.
	// TODO: Verify translation accuracy for "Bear"
	bearNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Bear",
	}

	// bearNameJapanese represents the Bear animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Bear"
	bearNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Bear",
	}

	// bearNameKorean represents the Bear animal type's name in Korean.
	// TODO: Verify translation accuracy for "Bear"
	bearNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Bear",
	}

	// bearNameLatinAmericanSpanish represents the Bear animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Bear"
	bearNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bear",
	}

	// bearNameRussian represents the Bear animal type's name in Russian.
	// TODO: Verify translation accuracy for "Bear"
	bearNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Bear",
	}

	// bearNameSimplifiedChinese represents the Bear animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Bear"
	bearNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Bear",
	}

	// bearNameSpanish represents the Bear animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Bear"
	bearNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Bear",
	}

	// bearNameTraditionalChinese represents the Bear animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Bear"
	bearNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Bear",
	}
)

var (
	// bearName contains the localized names of the Bear animal type.
	bearName = nook.Languages{
		language.AmericanEnglish:      bearNameAmericanEnglish,
		language.CanadianFrench:       bearNameCanadianFrench,
		language.Dutch:                bearNameDutch,
		language.French:               bearNameFrench,
		language.German:               bearNameGerman,
		language.Italian:              bearNameItalian,
		language.Japanese:             bearNameJapanese,
		language.Korean:               bearNameKorean,
		language.LatinAmericanSpanish: bearNameLatinAmericanSpanish,
		language.Russian:              bearNameRussian,
		language.SimplifiedChinese:    bearNameSimplifiedChinese,
		language.Spanish:              bearNameSpanish,
		language.TraditionalChinese:   bearNameTraditionalChinese,
	}
)

var (
	// Bear represents the Bear animal type in the Animal Crossing series.
	Bear = nook.Animal{
		Key:  nook.Key(bear),
		Name: bearName,
	}
)
