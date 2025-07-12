package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// panther is the common reference for the Panther animal type.
	panther = "Panther"
)

var (
	// pantherNameAmericanEnglish represents the Panther animal type's name in American English.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Panther",
	}

	// pantherNameCanadianFrench represents the Panther animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Panther",
	}

	// pantherNameDutch represents the Panther animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Panther",
	}

	// pantherNameFrench represents the Panther animal type's name in French.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameFrench = nook.Name{
		Language: language.French,
		Value:    "Panther",
	}

	// pantherNameGerman represents the Panther animal type's name in German.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameGerman = nook.Name{
		Language: language.German,
		Value:    "Panther",
	}

	// pantherNameItalian represents the Panther animal type's name in Italian.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Panther",
	}

	// pantherNameJapanese represents the Panther animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Panther",
	}

	// pantherNameKorean represents the Panther animal type's name in Korean.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Panther",
	}

	// pantherNameLatinAmericanSpanish represents the Panther animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Panther",
	}

	// pantherNameRussian represents the Panther animal type's name in Russian.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Panther",
	}

	// pantherNameSimplifiedChinese represents the Panther animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Panther",
	}

	// pantherNameSpanish represents the Panther animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Panther",
	}

	// pantherNameTraditionalChinese represents the Panther animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Panther"
	pantherNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Panther",
	}
)

var (
	// pantherName contains the localized names of the Panther animal type.
	pantherName = nook.Languages{
		language.AmericanEnglish:      pantherNameAmericanEnglish,
		language.CanadianFrench:       pantherNameCanadianFrench,
		language.Dutch:                pantherNameDutch,
		language.French:               pantherNameFrench,
		language.German:               pantherNameGerman,
		language.Italian:              pantherNameItalian,
		language.Japanese:             pantherNameJapanese,
		language.Korean:               pantherNameKorean,
		language.LatinAmericanSpanish: pantherNameLatinAmericanSpanish,
		language.Russian:              pantherNameRussian,
		language.SimplifiedChinese:    pantherNameSimplifiedChinese,
		language.Spanish:              pantherNameSpanish,
		language.TraditionalChinese:   pantherNameTraditionalChinese,
	}
)

var (
	// Panther represents the Panther animal type in the Animal Crossing series.
	Panther = nook.Animal{
		Key:  nook.Key(panther),
		Name: pantherName,
	}
)
