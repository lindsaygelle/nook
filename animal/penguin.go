package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// penguin is the common reference for the Penguin animal type.
	penguin = "Penguin"
)

var (
	// penguinNameAmericanEnglish represents the Penguin animal type's name in American English.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Penguin",
	}

	// penguinNameCanadianFrench represents the Penguin animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Penguin",
	}

	// penguinNameDutch represents the Penguin animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Penguin",
	}

	// penguinNameFrench represents the Penguin animal type's name in French.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameFrench = nook.Name{
		Language: language.French,
		Value:    "Penguin",
	}

	// penguinNameGerman represents the Penguin animal type's name in German.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameGerman = nook.Name{
		Language: language.German,
		Value:    "Penguin",
	}

	// penguinNameItalian represents the Penguin animal type's name in Italian.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Penguin",
	}

	// penguinNameJapanese represents the Penguin animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Penguin",
	}

	// penguinNameKorean represents the Penguin animal type's name in Korean.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Penguin",
	}

	// penguinNameLatinAmericanSpanish represents the Penguin animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Penguin",
	}

	// penguinNameRussian represents the Penguin animal type's name in Russian.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Penguin",
	}

	// penguinNameSimplifiedChinese represents the Penguin animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Penguin",
	}

	// penguinNameSpanish represents the Penguin animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Penguin",
	}

	// penguinNameTraditionalChinese represents the Penguin animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Penguin"
	penguinNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Penguin",
	}
)

var (
	// penguinName contains the localized names of the Penguin animal type.
	penguinName = nook.Languages{
		language.AmericanEnglish:      penguinNameAmericanEnglish,
		language.CanadianFrench:       penguinNameCanadianFrench,
		language.Dutch:                penguinNameDutch,
		language.French:               penguinNameFrench,
		language.German:               penguinNameGerman,
		language.Italian:              penguinNameItalian,
		language.Japanese:             penguinNameJapanese,
		language.Korean:               penguinNameKorean,
		language.LatinAmericanSpanish: penguinNameLatinAmericanSpanish,
		language.Russian:              penguinNameRussian,
		language.SimplifiedChinese:    penguinNameSimplifiedChinese,
		language.Spanish:              penguinNameSpanish,
		language.TraditionalChinese:   penguinNameTraditionalChinese,
	}
)

var (
	// Penguin represents the Penguin animal type in the Animal Crossing series.
	Penguin = nook.Animal{
		Key:  nook.Key(penguin),
		Name: penguinName,
	}
)
