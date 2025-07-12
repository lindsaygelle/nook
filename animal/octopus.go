package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// octopus is the common reference for the Octopus animal type.
	octopus = "Octopus"
)

var (
	// octopusNameAmericanEnglish represents the Octopus animal type's name in American English.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Octopus",
	}

	// octopusNameCanadianFrench represents the Octopus animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Octopus",
	}

	// octopusNameDutch represents the Octopus animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Octopus",
	}

	// octopusNameFrench represents the Octopus animal type's name in French.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameFrench = nook.Name{
		Language: language.French,
		Value:    "Octopus",
	}

	// octopusNameGerman represents the Octopus animal type's name in German.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameGerman = nook.Name{
		Language: language.German,
		Value:    "Octopus",
	}

	// octopusNameItalian represents the Octopus animal type's name in Italian.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Octopus",
	}

	// octopusNameJapanese represents the Octopus animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Octopus",
	}

	// octopusNameKorean represents the Octopus animal type's name in Korean.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Octopus",
	}

	// octopusNameLatinAmericanSpanish represents the Octopus animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Octopus",
	}

	// octopusNameRussian represents the Octopus animal type's name in Russian.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Octopus",
	}

	// octopusNameSimplifiedChinese represents the Octopus animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Octopus",
	}

	// octopusNameSpanish represents the Octopus animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Octopus",
	}

	// octopusNameTraditionalChinese represents the Octopus animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Octopus"
	octopusNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Octopus",
	}
)

var (
	// octopusName contains the localized names of the Octopus animal type.
	octopusName = nook.Languages{
		language.AmericanEnglish:      octopusNameAmericanEnglish,
		language.CanadianFrench:       octopusNameCanadianFrench,
		language.Dutch:                octopusNameDutch,
		language.French:               octopusNameFrench,
		language.German:               octopusNameGerman,
		language.Italian:              octopusNameItalian,
		language.Japanese:             octopusNameJapanese,
		language.Korean:               octopusNameKorean,
		language.LatinAmericanSpanish: octopusNameLatinAmericanSpanish,
		language.Russian:              octopusNameRussian,
		language.SimplifiedChinese:    octopusNameSimplifiedChinese,
		language.Spanish:              octopusNameSpanish,
		language.TraditionalChinese:   octopusNameTraditionalChinese,
	}
)

var (
	// Octopus represents the Octopus animal type in the Animal Crossing series.
	Octopus = nook.Animal{
		Key:  nook.Key(octopus),
		Name: octopusName,
	}
)
