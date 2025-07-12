package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// walrus is the common reference for the Walrus animal type.
	walrus = "Walrus"
)

var (
	// walrusNameAmericanEnglish represents the Walrus animal type's name in American English.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Walrus",
	}

	// walrusNameCanadianFrench represents the Walrus animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Walrus",
	}

	// walrusNameDutch represents the Walrus animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Walrus",
	}

	// walrusNameFrench represents the Walrus animal type's name in French.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameFrench = nook.Name{
		Language: language.French,
		Value:    "Walrus",
	}

	// walrusNameGerman represents the Walrus animal type's name in German.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameGerman = nook.Name{
		Language: language.German,
		Value:    "Walrus",
	}

	// walrusNameItalian represents the Walrus animal type's name in Italian.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Walrus",
	}

	// walrusNameJapanese represents the Walrus animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Walrus",
	}

	// walrusNameKorean represents the Walrus animal type's name in Korean.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Walrus",
	}

	// walrusNameLatinAmericanSpanish represents the Walrus animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Walrus",
	}

	// walrusNameRussian represents the Walrus animal type's name in Russian.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Walrus",
	}

	// walrusNameSimplifiedChinese represents the Walrus animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Walrus",
	}

	// walrusNameSpanish represents the Walrus animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Walrus",
	}

	// walrusNameTraditionalChinese represents the Walrus animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Walrus"
	walrusNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Walrus",
	}
)

var (
	// walrusName contains the localized names of the Walrus animal type.
	walrusName = nook.Languages{
		language.AmericanEnglish:      walrusNameAmericanEnglish,
		language.CanadianFrench:       walrusNameCanadianFrench,
		language.Dutch:                walrusNameDutch,
		language.French:               walrusNameFrench,
		language.German:               walrusNameGerman,
		language.Italian:              walrusNameItalian,
		language.Japanese:             walrusNameJapanese,
		language.Korean:               walrusNameKorean,
		language.LatinAmericanSpanish: walrusNameLatinAmericanSpanish,
		language.Russian:              walrusNameRussian,
		language.SimplifiedChinese:    walrusNameSimplifiedChinese,
		language.Spanish:              walrusNameSpanish,
		language.TraditionalChinese:   walrusNameTraditionalChinese,
	}
)

var (
	// Walrus represents the Walrus animal type in the Animal Crossing series.
	Walrus = nook.Animal{
		Key:  nook.Key(walrus),
		Name: walrusName,
	}
)
