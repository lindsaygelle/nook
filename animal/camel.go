package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// camel is the common reference for the Camel animal type.
	camel = "Camel"
)

var (
	// camelNameAmericanEnglish represents the Camel animal type's name in American English.
	// TODO: Verify translation accuracy for "Camel"
	camelNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Camel",
	}

	// camelNameCanadianFrench represents the Camel animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Camel"
	camelNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Camel",
	}

	// camelNameDutch represents the Camel animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Camel"
	camelNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Camel",
	}

	// camelNameFrench represents the Camel animal type's name in French.
	// TODO: Verify translation accuracy for "Camel"
	camelNameFrench = nook.Name{
		Language: language.French,
		Value:    "Camel",
	}

	// camelNameGerman represents the Camel animal type's name in German.
	// TODO: Verify translation accuracy for "Camel"
	camelNameGerman = nook.Name{
		Language: language.German,
		Value:    "Camel",
	}

	// camelNameItalian represents the Camel animal type's name in Italian.
	// TODO: Verify translation accuracy for "Camel"
	camelNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Camel",
	}

	// camelNameJapanese represents the Camel animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Camel"
	camelNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Camel",
	}

	// camelNameKorean represents the Camel animal type's name in Korean.
	// TODO: Verify translation accuracy for "Camel"
	camelNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Camel",
	}

	// camelNameLatinAmericanSpanish represents the Camel animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Camel"
	camelNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Camel",
	}

	// camelNameRussian represents the Camel animal type's name in Russian.
	// TODO: Verify translation accuracy for "Camel"
	camelNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Camel",
	}

	// camelNameSimplifiedChinese represents the Camel animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Camel"
	camelNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Camel",
	}

	// camelNameSpanish represents the Camel animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Camel"
	camelNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Camel",
	}

	// camelNameTraditionalChinese represents the Camel animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Camel"
	camelNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Camel",
	}
)

var (
	// camelName contains the localized names of the Camel animal type.
	camelName = nook.Languages{
		language.AmericanEnglish:      camelNameAmericanEnglish,
		language.CanadianFrench:       camelNameCanadianFrench,
		language.Dutch:                camelNameDutch,
		language.French:               camelNameFrench,
		language.German:               camelNameGerman,
		language.Italian:              camelNameItalian,
		language.Japanese:             camelNameJapanese,
		language.Korean:               camelNameKorean,
		language.LatinAmericanSpanish: camelNameLatinAmericanSpanish,
		language.Russian:              camelNameRussian,
		language.SimplifiedChinese:    camelNameSimplifiedChinese,
		language.Spanish:              camelNameSpanish,
		language.TraditionalChinese:   camelNameTraditionalChinese,
	}
)

var (
	// Camel represents the Camel animal type in the Animal Crossing series.
	Camel = nook.Animal{
		Key:  nook.Key(camel),
		Name: camelName,
	}
)
