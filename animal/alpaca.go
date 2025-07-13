package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// alpaca is the common reference for the Alpaca animal type.
	alpaca = "Alpaca"
)

var (
	// alpacaNameAmericanEnglish represents the Alpaca animal type's name in American English.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Alpaca",
	}

	// alpacaNameCanadianFrench represents the Alpaca animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Alpaca",
	}

	// alpacaNameDutch represents the Alpaca animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Alpaca",
	}

	// alpacaNameFrench represents the Alpaca animal type's name in French.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameFrench = nook.Name{
		Language: language.French,
		Value:    "Alpaca",
	}

	// alpacaNameGerman represents the Alpaca animal type's name in German.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameGerman = nook.Name{
		Language: language.German,
		Value:    "Alpaca",
	}

	// alpacaNameItalian represents the Alpaca animal type's name in Italian.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Alpaca",
	}

	// alpacaNameJapanese represents the Alpaca animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Alpaca",
	}

	// alpacaNameKorean represents the Alpaca animal type's name in Korean.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Alpaca",
	}

	// alpacaNameLatinAmericanSpanish represents the Alpaca animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Alpaca",
	}

	// alpacaNameRussian represents the Alpaca animal type's name in Russian.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Alpaca",
	}

	// alpacaNameSimplifiedChinese represents the Alpaca animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Alpaca",
	}

	// alpacaNameSpanish represents the Alpaca animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Alpaca",
	}

	// alpacaNameTraditionalChinese represents the Alpaca animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Alpaca"
	alpacaNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Alpaca",
	}
)

var (
	// alpacaName contains the localized names of the Alpaca animal type.
	alpacaName = nook.Languages{
		language.AmericanEnglish:      alpacaNameAmericanEnglish,
		language.CanadianFrench:       alpacaNameCanadianFrench,
		language.Dutch:                alpacaNameDutch,
		language.French:               alpacaNameFrench,
		language.German:               alpacaNameGerman,
		language.Italian:              alpacaNameItalian,
		language.Japanese:             alpacaNameJapanese,
		language.Korean:               alpacaNameKorean,
		language.LatinAmericanSpanish: alpacaNameLatinAmericanSpanish,
		language.Russian:              alpacaNameRussian,
		language.SimplifiedChinese:    alpacaNameSimplifiedChinese,
		language.Spanish:              alpacaNameSpanish,
		language.TraditionalChinese:   alpacaNameTraditionalChinese,
	}
)

var (
	// Alpaca represents the Alpaca animal type in the Animal Crossing series.
	Alpaca = nook.Animal{
		Key:  nook.Key(alpaca),
		Name: alpacaName,
	}
)
