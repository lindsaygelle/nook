package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// dodo is the common reference for the Dodo animal type.
	dodo = "Dodo"
)

var (
	// dodoNameAmericanEnglish represents the Dodo animal type's name in American English.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dodo",
	}

	// dodoNameCanadianFrench represents the Dodo animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dodo",
	}

	// dodoNameDutch represents the Dodo animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Dodo",
	}

	// dodoNameFrench represents the Dodo animal type's name in French.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameFrench = nook.Name{
		Language: language.French,
		Value:    "Dodo",
	}

	// dodoNameGerman represents the Dodo animal type's name in German.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameGerman = nook.Name{
		Language: language.German,
		Value:    "Dodo",
	}

	// dodoNameItalian represents the Dodo animal type's name in Italian.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Dodo",
	}

	// dodoNameJapanese represents the Dodo animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Dodo",
	}

	// dodoNameKorean represents the Dodo animal type's name in Korean.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Dodo",
	}

	// dodoNameLatinAmericanSpanish represents the Dodo animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dodo",
	}

	// dodoNameRussian represents the Dodo animal type's name in Russian.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Dodo",
	}

	// dodoNameSimplifiedChinese represents the Dodo animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Dodo",
	}

	// dodoNameSpanish represents the Dodo animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Dodo",
	}

	// dodoNameTraditionalChinese represents the Dodo animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Dodo"
	dodoNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Dodo",
	}
)

var (
	// dodoName contains the localized names of the Dodo animal type.
	dodoName = nook.Languages{
		language.AmericanEnglish:      dodoNameAmericanEnglish,
		language.CanadianFrench:       dodoNameCanadianFrench,
		language.Dutch:                dodoNameDutch,
		language.French:               dodoNameFrench,
		language.German:               dodoNameGerman,
		language.Italian:              dodoNameItalian,
		language.Japanese:             dodoNameJapanese,
		language.Korean:               dodoNameKorean,
		language.LatinAmericanSpanish: dodoNameLatinAmericanSpanish,
		language.Russian:              dodoNameRussian,
		language.SimplifiedChinese:    dodoNameSimplifiedChinese,
		language.Spanish:              dodoNameSpanish,
		language.TraditionalChinese:   dodoNameTraditionalChinese,
	}
)

var (
	// Dodo represents the Dodo animal type in the Animal Crossing series.
	Dodo = nook.Animal{
		Key:  nook.Key(dodo),
		Name: dodoName,
	}
)
