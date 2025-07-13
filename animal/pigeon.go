package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// pigeon is the common reference for the Pigeon animal type.
	pigeon = "Pigeon"
)

var (
	// pigeonNameAmericanEnglish represents the Pigeon animal type's name in American English.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pigeon",
	}

	// pigeonNameCanadianFrench represents the Pigeon animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pigeon",
	}

	// pigeonNameDutch represents the Pigeon animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Pigeon",
	}

	// pigeonNameFrench represents the Pigeon animal type's name in French.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameFrench = nook.Name{
		Language: language.French,
		Value:    "Pigeon",
	}

	// pigeonNameGerman represents the Pigeon animal type's name in German.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameGerman = nook.Name{
		Language: language.German,
		Value:    "Pigeon",
	}

	// pigeonNameItalian represents the Pigeon animal type's name in Italian.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Pigeon",
	}

	// pigeonNameJapanese represents the Pigeon animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Pigeon",
	}

	// pigeonNameKorean represents the Pigeon animal type's name in Korean.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Pigeon",
	}

	// pigeonNameLatinAmericanSpanish represents the Pigeon animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pigeon",
	}

	// pigeonNameRussian represents the Pigeon animal type's name in Russian.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Pigeon",
	}

	// pigeonNameSimplifiedChinese represents the Pigeon animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Pigeon",
	}

	// pigeonNameSpanish represents the Pigeon animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Pigeon",
	}

	// pigeonNameTraditionalChinese represents the Pigeon animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Pigeon"
	pigeonNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Pigeon",
	}
)

var (
	// pigeonName contains the localized names of the Pigeon animal type.
	pigeonName = nook.Languages{
		language.AmericanEnglish:      pigeonNameAmericanEnglish,
		language.CanadianFrench:       pigeonNameCanadianFrench,
		language.Dutch:                pigeonNameDutch,
		language.French:               pigeonNameFrench,
		language.German:               pigeonNameGerman,
		language.Italian:              pigeonNameItalian,
		language.Japanese:             pigeonNameJapanese,
		language.Korean:               pigeonNameKorean,
		language.LatinAmericanSpanish: pigeonNameLatinAmericanSpanish,
		language.Russian:              pigeonNameRussian,
		language.SimplifiedChinese:    pigeonNameSimplifiedChinese,
		language.Spanish:              pigeonNameSpanish,
		language.TraditionalChinese:   pigeonNameTraditionalChinese,
	}
)

var (
	// Pigeon represents the Pigeon animal type in the Animal Crossing series.
	Pigeon = nook.Animal{
		Key:  nook.Key(pigeon),
		Name: pigeonName,
	}
)
