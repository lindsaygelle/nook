package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// sloth is the common reference for the Sloth animal type.
	sloth = "Sloth"
)

var (
	// slothNameAmericanEnglish represents the Sloth animal type's name in American English.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sloth",
	}

	// slothNameCanadianFrench represents the Sloth animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sloth",
	}

	// slothNameDutch represents the Sloth animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Sloth",
	}

	// slothNameFrench represents the Sloth animal type's name in French.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameFrench = nook.Name{
		Language: language.French,
		Value:    "Sloth",
	}

	// slothNameGerman represents the Sloth animal type's name in German.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameGerman = nook.Name{
		Language: language.German,
		Value:    "Sloth",
	}

	// slothNameItalian represents the Sloth animal type's name in Italian.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Sloth",
	}

	// slothNameJapanese represents the Sloth animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Sloth",
	}

	// slothNameKorean represents the Sloth animal type's name in Korean.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Sloth",
	}

	// slothNameLatinAmericanSpanish represents the Sloth animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sloth",
	}

	// slothNameRussian represents the Sloth animal type's name in Russian.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Sloth",
	}

	// slothNameSimplifiedChinese represents the Sloth animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Sloth",
	}

	// slothNameSpanish represents the Sloth animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Sloth",
	}

	// slothNameTraditionalChinese represents the Sloth animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Sloth"
	slothNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Sloth",
	}
)

var (
	// slothName contains the localized names of the Sloth animal type.
	slothName = nook.Languages{
		language.AmericanEnglish:      slothNameAmericanEnglish,
		language.CanadianFrench:       slothNameCanadianFrench,
		language.Dutch:                slothNameDutch,
		language.French:               slothNameFrench,
		language.German:               slothNameGerman,
		language.Italian:              slothNameItalian,
		language.Japanese:             slothNameJapanese,
		language.Korean:               slothNameKorean,
		language.LatinAmericanSpanish: slothNameLatinAmericanSpanish,
		language.Russian:              slothNameRussian,
		language.SimplifiedChinese:    slothNameSimplifiedChinese,
		language.Spanish:              slothNameSpanish,
		language.TraditionalChinese:   slothNameTraditionalChinese,
	}
)

var (
	// Sloth represents the Sloth animal type in the Animal Crossing series.
	Sloth = nook.Animal{
		Key:  nook.Key(sloth),
		Name: slothName,
	}
)
