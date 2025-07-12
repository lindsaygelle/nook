package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// beaver is the common reference for the Beaver animal type.
	beaver = "Beaver"
)

var (
	// beaverNameAmericanEnglish represents the Beaver animal type's name in American English.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Beaver",
	}

	// beaverNameCanadianFrench represents the Beaver animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Beaver",
	}

	// beaverNameDutch represents the Beaver animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Beaver",
	}

	// beaverNameFrench represents the Beaver animal type's name in French.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameFrench = nook.Name{
		Language: language.French,
		Value:    "Beaver",
	}

	// beaverNameGerman represents the Beaver animal type's name in German.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameGerman = nook.Name{
		Language: language.German,
		Value:    "Beaver",
	}

	// beaverNameItalian represents the Beaver animal type's name in Italian.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Beaver",
	}

	// beaverNameJapanese represents the Beaver animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Beaver",
	}

	// beaverNameKorean represents the Beaver animal type's name in Korean.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Beaver",
	}

	// beaverNameLatinAmericanSpanish represents the Beaver animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Beaver",
	}

	// beaverNameRussian represents the Beaver animal type's name in Russian.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Beaver",
	}

	// beaverNameSimplifiedChinese represents the Beaver animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Beaver",
	}

	// beaverNameSpanish represents the Beaver animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Beaver",
	}

	// beaverNameTraditionalChinese represents the Beaver animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Beaver"
	beaverNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Beaver",
	}
)

var (
	// beaverName contains the localized names of the Beaver animal type.
	beaverName = nook.Languages{
		language.AmericanEnglish:      beaverNameAmericanEnglish,
		language.CanadianFrench:       beaverNameCanadianFrench,
		language.Dutch:                beaverNameDutch,
		language.French:               beaverNameFrench,
		language.German:               beaverNameGerman,
		language.Italian:              beaverNameItalian,
		language.Japanese:             beaverNameJapanese,
		language.Korean:               beaverNameKorean,
		language.LatinAmericanSpanish: beaverNameLatinAmericanSpanish,
		language.Russian:              beaverNameRussian,
		language.SimplifiedChinese:    beaverNameSimplifiedChinese,
		language.Spanish:              beaverNameSpanish,
		language.TraditionalChinese:   beaverNameTraditionalChinese,
	}
)

var (
	// Beaver represents the Beaver animal type in the Animal Crossing series.
	Beaver = nook.Animal{
		Key:  nook.Key(beaver),
		Name: beaverName,
	}
)
