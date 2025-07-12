package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// skunk is the common reference for the Skunk animal type.
	skunk = "Skunk"
)

var (
	// skunkNameAmericanEnglish represents the Skunk animal type's name in American English.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Skunk",
	}

	// skunkNameCanadianFrench represents the Skunk animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Skunk",
	}

	// skunkNameDutch represents the Skunk animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Skunk",
	}

	// skunkNameFrench represents the Skunk animal type's name in French.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameFrench = nook.Name{
		Language: language.French,
		Value:    "Skunk",
	}

	// skunkNameGerman represents the Skunk animal type's name in German.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameGerman = nook.Name{
		Language: language.German,
		Value:    "Skunk",
	}

	// skunkNameItalian represents the Skunk animal type's name in Italian.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Skunk",
	}

	// skunkNameJapanese represents the Skunk animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Skunk",
	}

	// skunkNameKorean represents the Skunk animal type's name in Korean.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Skunk",
	}

	// skunkNameLatinAmericanSpanish represents the Skunk animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Skunk",
	}

	// skunkNameRussian represents the Skunk animal type's name in Russian.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Skunk",
	}

	// skunkNameSimplifiedChinese represents the Skunk animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Skunk",
	}

	// skunkNameSpanish represents the Skunk animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Skunk",
	}

	// skunkNameTraditionalChinese represents the Skunk animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Skunk"
	skunkNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Skunk",
	}
)

var (
	// skunkName contains the localized names of the Skunk animal type.
	skunkName = nook.Languages{
		language.AmericanEnglish:      skunkNameAmericanEnglish,
		language.CanadianFrench:       skunkNameCanadianFrench,
		language.Dutch:                skunkNameDutch,
		language.French:               skunkNameFrench,
		language.German:               skunkNameGerman,
		language.Italian:              skunkNameItalian,
		language.Japanese:             skunkNameJapanese,
		language.Korean:               skunkNameKorean,
		language.LatinAmericanSpanish: skunkNameLatinAmericanSpanish,
		language.Russian:              skunkNameRussian,
		language.SimplifiedChinese:    skunkNameSimplifiedChinese,
		language.Spanish:              skunkNameSpanish,
		language.TraditionalChinese:   skunkNameTraditionalChinese,
	}
)

var (
	// Skunk represents the Skunk animal type in the Animal Crossing series.
	Skunk = nook.Animal{
		Key:  nook.Key(skunk),
		Name: skunkName,
	}
)
