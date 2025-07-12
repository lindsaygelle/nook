package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// deer is the common reference for the Deer animal type.
	deer = "Deer"
)

var (
	// deerNameAmericanEnglish represents the Deer animal type's name in American English.
	// TODO: Verify translation accuracy for "Deer"
	deerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Deer",
	}

	// deerNameCanadianFrench represents the Deer animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Deer"
	deerNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Deer",
	}

	// deerNameDutch represents the Deer animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Deer"
	deerNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Deer",
	}

	// deerNameFrench represents the Deer animal type's name in French.
	// TODO: Verify translation accuracy for "Deer"
	deerNameFrench = nook.Name{
		Language: language.French,
		Value:    "Deer",
	}

	// deerNameGerman represents the Deer animal type's name in German.
	// TODO: Verify translation accuracy for "Deer"
	deerNameGerman = nook.Name{
		Language: language.German,
		Value:    "Deer",
	}

	// deerNameItalian represents the Deer animal type's name in Italian.
	// TODO: Verify translation accuracy for "Deer"
	deerNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Deer",
	}

	// deerNameJapanese represents the Deer animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Deer"
	deerNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Deer",
	}

	// deerNameKorean represents the Deer animal type's name in Korean.
	// TODO: Verify translation accuracy for "Deer"
	deerNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Deer",
	}

	// deerNameLatinAmericanSpanish represents the Deer animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Deer"
	deerNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Deer",
	}

	// deerNameRussian represents the Deer animal type's name in Russian.
	// TODO: Verify translation accuracy for "Deer"
	deerNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Deer",
	}

	// deerNameSimplifiedChinese represents the Deer animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Deer"
	deerNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Deer",
	}

	// deerNameSpanish represents the Deer animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Deer"
	deerNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Deer",
	}

	// deerNameTraditionalChinese represents the Deer animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Deer"
	deerNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Deer",
	}
)

var (
	// deerName contains the localized names of the Deer animal type.
	deerName = nook.Languages{
		language.AmericanEnglish:      deerNameAmericanEnglish,
		language.CanadianFrench:       deerNameCanadianFrench,
		language.Dutch:                deerNameDutch,
		language.French:               deerNameFrench,
		language.German:               deerNameGerman,
		language.Italian:              deerNameItalian,
		language.Japanese:             deerNameJapanese,
		language.Korean:               deerNameKorean,
		language.LatinAmericanSpanish: deerNameLatinAmericanSpanish,
		language.Russian:              deerNameRussian,
		language.SimplifiedChinese:    deerNameSimplifiedChinese,
		language.Spanish:              deerNameSpanish,
		language.TraditionalChinese:   deerNameTraditionalChinese,
	}
)

var (
	// Deer represents the Deer animal type in the Animal Crossing series.
	Deer = nook.Animal{
		Key:  nook.Key(deer),
		Name: deerName,
	}
)
