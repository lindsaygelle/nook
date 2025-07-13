package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// wolf is the common reference for the Wolf animal type.
	wolf = "Wolf"
)

var (
	// wolfNameAmericanEnglish represents the Wolf animal type's name in American English.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wolf",
	}

	// wolfNameCanadianFrench represents the Wolf animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Wolf",
	}

	// wolfNameDutch represents the Wolf animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Wolf",
	}

	// wolfNameFrench represents the Wolf animal type's name in French.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameFrench = nook.Name{
		Language: language.French,
		Value:    "Wolf",
	}

	// wolfNameGerman represents the Wolf animal type's name in German.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameGerman = nook.Name{
		Language: language.German,
		Value:    "Wolf",
	}

	// wolfNameItalian represents the Wolf animal type's name in Italian.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Wolf",
	}

	// wolfNameJapanese represents the Wolf animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Wolf",
	}

	// wolfNameKorean represents the Wolf animal type's name in Korean.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Wolf",
	}

	// wolfNameLatinAmericanSpanish represents the Wolf animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Wolf",
	}

	// wolfNameRussian represents the Wolf animal type's name in Russian.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Wolf",
	}

	// wolfNameSimplifiedChinese represents the Wolf animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Wolf",
	}

	// wolfNameSpanish represents the Wolf animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Wolf",
	}

	// wolfNameTraditionalChinese represents the Wolf animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Wolf"
	wolfNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Wolf",
	}
)

var (
	// wolfName contains the localized names of the Wolf animal type.
	wolfName = nook.Languages{
		language.AmericanEnglish:      wolfNameAmericanEnglish,
		language.CanadianFrench:       wolfNameCanadianFrench,
		language.Dutch:                wolfNameDutch,
		language.French:               wolfNameFrench,
		language.German:               wolfNameGerman,
		language.Italian:              wolfNameItalian,
		language.Japanese:             wolfNameJapanese,
		language.Korean:               wolfNameKorean,
		language.LatinAmericanSpanish: wolfNameLatinAmericanSpanish,
		language.Russian:              wolfNameRussian,
		language.SimplifiedChinese:    wolfNameSimplifiedChinese,
		language.Spanish:              wolfNameSpanish,
		language.TraditionalChinese:   wolfNameTraditionalChinese,
	}
)

var (
	// Wolf represents the Wolf animal type in the Animal Crossing series.
	Wolf = nook.Animal{
		Key:  nook.Key(wolf),
		Name: wolfName,
	}
)
