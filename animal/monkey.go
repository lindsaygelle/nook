package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// monkey is the common reference for the Monkey animal type.
	monkey = "Monkey"
)

var (
	// monkeyNameAmericanEnglish represents the Monkey animal type's name in American English.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Monkey",
	}

	// monkeyNameCanadianFrench represents the Monkey animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Monkey",
	}

	// monkeyNameDutch represents the Monkey animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Monkey",
	}

	// monkeyNameFrench represents the Monkey animal type's name in French.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameFrench = nook.Name{
		Language: language.French,
		Value:    "Monkey",
	}

	// monkeyNameGerman represents the Monkey animal type's name in German.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameGerman = nook.Name{
		Language: language.German,
		Value:    "Monkey",
	}

	// monkeyNameItalian represents the Monkey animal type's name in Italian.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Monkey",
	}

	// monkeyNameJapanese represents the Monkey animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Monkey",
	}

	// monkeyNameKorean represents the Monkey animal type's name in Korean.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Monkey",
	}

	// monkeyNameLatinAmericanSpanish represents the Monkey animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Monkey",
	}

	// monkeyNameRussian represents the Monkey animal type's name in Russian.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Monkey",
	}

	// monkeyNameSimplifiedChinese represents the Monkey animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Monkey",
	}

	// monkeyNameSpanish represents the Monkey animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Monkey",
	}

	// monkeyNameTraditionalChinese represents the Monkey animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Monkey"
	monkeyNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Monkey",
	}
)

var (
	// monkeyName contains the localized names of the Monkey animal type.
	monkeyName = nook.Languages{
		language.AmericanEnglish:      monkeyNameAmericanEnglish,
		language.CanadianFrench:       monkeyNameCanadianFrench,
		language.Dutch:                monkeyNameDutch,
		language.French:               monkeyNameFrench,
		language.German:               monkeyNameGerman,
		language.Italian:              monkeyNameItalian,
		language.Japanese:             monkeyNameJapanese,
		language.Korean:               monkeyNameKorean,
		language.LatinAmericanSpanish: monkeyNameLatinAmericanSpanish,
		language.Russian:              monkeyNameRussian,
		language.SimplifiedChinese:    monkeyNameSimplifiedChinese,
		language.Spanish:              monkeyNameSpanish,
		language.TraditionalChinese:   monkeyNameTraditionalChinese,
	}
)

var (
	// Monkey represents the Monkey animal type in the Animal Crossing series.
	Monkey = nook.Animal{
		Key:  nook.Key(monkey),
		Name: monkeyName,
	}
)
