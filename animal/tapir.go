package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// tapir is the common reference for the Tapir animal type.
	tapir = "Tapir"
)

var (
	// tapirNameAmericanEnglish represents the Tapir animal type's name in American English.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tapir",
	}

	// tapirNameCanadianFrench represents the Tapir animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tapir",
	}

	// tapirNameDutch represents the Tapir animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Tapir",
	}

	// tapirNameFrench represents the Tapir animal type's name in French.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameFrench = nook.Name{
		Language: language.French,
		Value:    "Tapir",
	}

	// tapirNameGerman represents the Tapir animal type's name in German.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameGerman = nook.Name{
		Language: language.German,
		Value:    "Tapir",
	}

	// tapirNameItalian represents the Tapir animal type's name in Italian.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Tapir",
	}

	// tapirNameJapanese represents the Tapir animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Tapir",
	}

	// tapirNameKorean represents the Tapir animal type's name in Korean.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Tapir",
	}

	// tapirNameLatinAmericanSpanish represents the Tapir animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tapir",
	}

	// tapirNameRussian represents the Tapir animal type's name in Russian.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Tapir",
	}

	// tapirNameSimplifiedChinese represents the Tapir animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Tapir",
	}

	// tapirNameSpanish represents the Tapir animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Tapir",
	}

	// tapirNameTraditionalChinese represents the Tapir animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Tapir"
	tapirNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Tapir",
	}
)

var (
	// tapirName contains the localized names of the Tapir animal type.
	tapirName = nook.Languages{
		language.AmericanEnglish:      tapirNameAmericanEnglish,
		language.CanadianFrench:       tapirNameCanadianFrench,
		language.Dutch:                tapirNameDutch,
		language.French:               tapirNameFrench,
		language.German:               tapirNameGerman,
		language.Italian:              tapirNameItalian,
		language.Japanese:             tapirNameJapanese,
		language.Korean:               tapirNameKorean,
		language.LatinAmericanSpanish: tapirNameLatinAmericanSpanish,
		language.Russian:              tapirNameRussian,
		language.SimplifiedChinese:    tapirNameSimplifiedChinese,
		language.Spanish:              tapirNameSpanish,
		language.TraditionalChinese:   tapirNameTraditionalChinese,
	}
)

var (
	// Tapir represents the Tapir animal type in the Animal Crossing series.
	Tapir = nook.Animal{
		Key:  nook.Key(tapir),
		Name: tapirName,
	}
)
