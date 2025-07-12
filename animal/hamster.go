package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// hamster is the common reference for the Hamster animal type.
	hamster = "Hamster"
)

var (
	// hamsterNameAmericanEnglish represents the Hamster animal type's name in American English.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hamster",
	}

	// hamsterNameCanadianFrench represents the Hamster animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hamster",
	}

	// hamsterNameDutch represents the Hamster animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Hamster",
	}

	// hamsterNameFrench represents the Hamster animal type's name in French.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameFrench = nook.Name{
		Language: language.French,
		Value:    "Hamster",
	}

	// hamsterNameGerman represents the Hamster animal type's name in German.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameGerman = nook.Name{
		Language: language.German,
		Value:    "Hamster",
	}

	// hamsterNameItalian represents the Hamster animal type's name in Italian.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Hamster",
	}

	// hamsterNameJapanese represents the Hamster animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Hamster",
	}

	// hamsterNameKorean represents the Hamster animal type's name in Korean.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Hamster",
	}

	// hamsterNameLatinAmericanSpanish represents the Hamster animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hamster",
	}

	// hamsterNameRussian represents the Hamster animal type's name in Russian.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Hamster",
	}

	// hamsterNameSimplifiedChinese represents the Hamster animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Hamster",
	}

	// hamsterNameSpanish represents the Hamster animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Hamster",
	}

	// hamsterNameTraditionalChinese represents the Hamster animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Hamster"
	hamsterNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Hamster",
	}
)

var (
	// hamsterName contains the localized names of the Hamster animal type.
	hamsterName = nook.Languages{
		language.AmericanEnglish:      hamsterNameAmericanEnglish,
		language.CanadianFrench:       hamsterNameCanadianFrench,
		language.Dutch:                hamsterNameDutch,
		language.French:               hamsterNameFrench,
		language.German:               hamsterNameGerman,
		language.Italian:              hamsterNameItalian,
		language.Japanese:             hamsterNameJapanese,
		language.Korean:               hamsterNameKorean,
		language.LatinAmericanSpanish: hamsterNameLatinAmericanSpanish,
		language.Russian:              hamsterNameRussian,
		language.SimplifiedChinese:    hamsterNameSimplifiedChinese,
		language.Spanish:              hamsterNameSpanish,
		language.TraditionalChinese:   hamsterNameTraditionalChinese,
	}
)

var (
	// Hamster represents the Hamster animal type in the Animal Crossing series.
	Hamster = nook.Animal{
		Key:  nook.Key(hamster),
		Name: hamsterName,
	}
)
