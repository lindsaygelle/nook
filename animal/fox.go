package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// fox is the common reference for the Fox animal type.
	fox = "Fox"
)

var (
	// foxNameAmericanEnglish represents the Fox animal type's name in American English.
	// TODO: Verify translation accuracy for "Fox"
	foxNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Fox",
	}

	// foxNameCanadianFrench represents the Fox animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Fox"
	foxNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Fox",
	}

	// foxNameDutch represents the Fox animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Fox"
	foxNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Fox",
	}

	// foxNameFrench represents the Fox animal type's name in French.
	// TODO: Verify translation accuracy for "Fox"
	foxNameFrench = nook.Name{
		Language: language.French,
		Value:    "Fox",
	}

	// foxNameGerman represents the Fox animal type's name in German.
	// TODO: Verify translation accuracy for "Fox"
	foxNameGerman = nook.Name{
		Language: language.German,
		Value:    "Fox",
	}

	// foxNameItalian represents the Fox animal type's name in Italian.
	// TODO: Verify translation accuracy for "Fox"
	foxNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Fox",
	}

	// foxNameJapanese represents the Fox animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Fox"
	foxNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Fox",
	}

	// foxNameKorean represents the Fox animal type's name in Korean.
	// TODO: Verify translation accuracy for "Fox"
	foxNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Fox",
	}

	// foxNameLatinAmericanSpanish represents the Fox animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Fox"
	foxNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fox",
	}

	// foxNameRussian represents the Fox animal type's name in Russian.
	// TODO: Verify translation accuracy for "Fox"
	foxNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Fox",
	}

	// foxNameSimplifiedChinese represents the Fox animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Fox"
	foxNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Fox",
	}

	// foxNameSpanish represents the Fox animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Fox"
	foxNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Fox",
	}

	// foxNameTraditionalChinese represents the Fox animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Fox"
	foxNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Fox",
	}
)

var (
	// foxName contains the localized names of the Fox animal type.
	foxName = nook.Languages{
		language.AmericanEnglish:      foxNameAmericanEnglish,
		language.CanadianFrench:       foxNameCanadianFrench,
		language.Dutch:                foxNameDutch,
		language.French:               foxNameFrench,
		language.German:               foxNameGerman,
		language.Italian:              foxNameItalian,
		language.Japanese:             foxNameJapanese,
		language.Korean:               foxNameKorean,
		language.LatinAmericanSpanish: foxNameLatinAmericanSpanish,
		language.Russian:              foxNameRussian,
		language.SimplifiedChinese:    foxNameSimplifiedChinese,
		language.Spanish:              foxNameSpanish,
		language.TraditionalChinese:   foxNameTraditionalChinese,
	}
)

var (
	// Fox represents the Fox animal type in the Animal Crossing series.
	Fox = nook.Animal{
		Key:  nook.Key(fox),
		Name: foxName,
	}
)
