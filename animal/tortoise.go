package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// tortoise is the common reference for the Tortoise animal type.
	tortoise = "Tortoise"
)

var (
	// tortoiseNameAmericanEnglish represents the Tortoise animal type's name in American English.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tortoise",
	}

	// tortoiseNameCanadianFrench represents the Tortoise animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tortoise",
	}

	// tortoiseNameDutch represents the Tortoise animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Tortoise",
	}

	// tortoiseNameFrench represents the Tortoise animal type's name in French.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameFrench = nook.Name{
		Language: language.French,
		Value:    "Tortoise",
	}

	// tortoiseNameGerman represents the Tortoise animal type's name in German.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameGerman = nook.Name{
		Language: language.German,
		Value:    "Tortoise",
	}

	// tortoiseNameItalian represents the Tortoise animal type's name in Italian.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Tortoise",
	}

	// tortoiseNameJapanese represents the Tortoise animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Tortoise",
	}

	// tortoiseNameKorean represents the Tortoise animal type's name in Korean.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Tortoise",
	}

	// tortoiseNameLatinAmericanSpanish represents the Tortoise animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tortoise",
	}

	// tortoiseNameRussian represents the Tortoise animal type's name in Russian.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Tortoise",
	}

	// tortoiseNameSimplifiedChinese represents the Tortoise animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Tortoise",
	}

	// tortoiseNameSpanish represents the Tortoise animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Tortoise",
	}

	// tortoiseNameTraditionalChinese represents the Tortoise animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Tortoise"
	tortoiseNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Tortoise",
	}
)

var (
	// tortoiseName contains the localized names of the Tortoise animal type.
	tortoiseName = nook.Languages{
		language.AmericanEnglish:      tortoiseNameAmericanEnglish,
		language.CanadianFrench:       tortoiseNameCanadianFrench,
		language.Dutch:                tortoiseNameDutch,
		language.French:               tortoiseNameFrench,
		language.German:               tortoiseNameGerman,
		language.Italian:              tortoiseNameItalian,
		language.Japanese:             tortoiseNameJapanese,
		language.Korean:               tortoiseNameKorean,
		language.LatinAmericanSpanish: tortoiseNameLatinAmericanSpanish,
		language.Russian:              tortoiseNameRussian,
		language.SimplifiedChinese:    tortoiseNameSimplifiedChinese,
		language.Spanish:              tortoiseNameSpanish,
		language.TraditionalChinese:   tortoiseNameTraditionalChinese,
	}
)

var (
	// Tortoise represents the Tortoise animal type in the Animal Crossing series.
	Tortoise = nook.Animal{
		Key:  nook.Key(tortoise),
		Name: tortoiseName,
	}
)
