package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// lion is the common reference for the Lion animal type.
	lion = "Lion"
)

var (
	// lionNameAmericanEnglish represents the Lion animal type's name in American English.
	// TODO: Verify translation accuracy for "Lion"
	lionNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lion",
	}

	// lionNameCanadianFrench represents the Lion animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Lion"
	lionNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lion",
	}

	// lionNameDutch represents the Lion animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Lion"
	lionNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Lion",
	}

	// lionNameFrench represents the Lion animal type's name in French.
	// TODO: Verify translation accuracy for "Lion"
	lionNameFrench = nook.Name{
		Language: language.French,
		Value:    "Lion",
	}

	// lionNameGerman represents the Lion animal type's name in German.
	// TODO: Verify translation accuracy for "Lion"
	lionNameGerman = nook.Name{
		Language: language.German,
		Value:    "Lion",
	}

	// lionNameItalian represents the Lion animal type's name in Italian.
	// TODO: Verify translation accuracy for "Lion"
	lionNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Lion",
	}

	// lionNameJapanese represents the Lion animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Lion"
	lionNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Lion",
	}

	// lionNameKorean represents the Lion animal type's name in Korean.
	// TODO: Verify translation accuracy for "Lion"
	lionNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Lion",
	}

	// lionNameLatinAmericanSpanish represents the Lion animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Lion"
	lionNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lion",
	}

	// lionNameRussian represents the Lion animal type's name in Russian.
	// TODO: Verify translation accuracy for "Lion"
	lionNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Lion",
	}

	// lionNameSimplifiedChinese represents the Lion animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Lion"
	lionNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Lion",
	}

	// lionNameSpanish represents the Lion animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Lion"
	lionNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Lion",
	}

	// lionNameTraditionalChinese represents the Lion animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Lion"
	lionNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Lion",
	}
)

var (
	// lionName contains the localized names of the Lion animal type.
	lionName = nook.Languages{
		language.AmericanEnglish:      lionNameAmericanEnglish,
		language.CanadianFrench:       lionNameCanadianFrench,
		language.Dutch:                lionNameDutch,
		language.French:               lionNameFrench,
		language.German:               lionNameGerman,
		language.Italian:              lionNameItalian,
		language.Japanese:             lionNameJapanese,
		language.Korean:               lionNameKorean,
		language.LatinAmericanSpanish: lionNameLatinAmericanSpanish,
		language.Russian:              lionNameRussian,
		language.SimplifiedChinese:    lionNameSimplifiedChinese,
		language.Spanish:              lionNameSpanish,
		language.TraditionalChinese:   lionNameTraditionalChinese,
	}
)

var (
	// Lion represents the Lion animal type in the Animal Crossing series.
	Lion = nook.Animal{
		Key:  nook.Key(lion),
		Name: lionName,
	}
)
