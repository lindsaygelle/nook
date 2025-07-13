package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// owl is the common reference for the Owl animal type.
	owl = "Owl"
)

var (
	// owlNameAmericanEnglish represents the Owl animal type's name in American English.
	// TODO: Verify translation accuracy for "Owl"
	owlNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Owl",
	}

	// owlNameCanadianFrench represents the Owl animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Owl"
	owlNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Owl",
	}

	// owlNameDutch represents the Owl animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Owl"
	owlNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Owl",
	}

	// owlNameFrench represents the Owl animal type's name in French.
	// TODO: Verify translation accuracy for "Owl"
	owlNameFrench = nook.Name{
		Language: language.French,
		Value:    "Owl",
	}

	// owlNameGerman represents the Owl animal type's name in German.
	// TODO: Verify translation accuracy for "Owl"
	owlNameGerman = nook.Name{
		Language: language.German,
		Value:    "Owl",
	}

	// owlNameItalian represents the Owl animal type's name in Italian.
	// TODO: Verify translation accuracy for "Owl"
	owlNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Owl",
	}

	// owlNameJapanese represents the Owl animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Owl"
	owlNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Owl",
	}

	// owlNameKorean represents the Owl animal type's name in Korean.
	// TODO: Verify translation accuracy for "Owl"
	owlNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Owl",
	}

	// owlNameLatinAmericanSpanish represents the Owl animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Owl"
	owlNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Owl",
	}

	// owlNameRussian represents the Owl animal type's name in Russian.
	// TODO: Verify translation accuracy for "Owl"
	owlNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Owl",
	}

	// owlNameSimplifiedChinese represents the Owl animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Owl"
	owlNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Owl",
	}

	// owlNameSpanish represents the Owl animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Owl"
	owlNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Owl",
	}

	// owlNameTraditionalChinese represents the Owl animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Owl"
	owlNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Owl",
	}
)

var (
	// owlName contains the localized names of the Owl animal type.
	owlName = nook.Languages{
		language.AmericanEnglish:      owlNameAmericanEnglish,
		language.CanadianFrench:       owlNameCanadianFrench,
		language.Dutch:                owlNameDutch,
		language.French:               owlNameFrench,
		language.German:               owlNameGerman,
		language.Italian:              owlNameItalian,
		language.Japanese:             owlNameJapanese,
		language.Korean:               owlNameKorean,
		language.LatinAmericanSpanish: owlNameLatinAmericanSpanish,
		language.Russian:              owlNameRussian,
		language.SimplifiedChinese:    owlNameSimplifiedChinese,
		language.Spanish:              owlNameSpanish,
		language.TraditionalChinese:   owlNameTraditionalChinese,
	}
)

var (
	// Owl represents the Owl animal type in the Animal Crossing series.
	Owl = nook.Animal{
		Key:  nook.Key(owl),
		Name: owlName,
	}
)
