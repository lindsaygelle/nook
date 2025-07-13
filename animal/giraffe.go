package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// giraffe is the common reference for the Giraffe animal type.
	giraffe = "Giraffe"
)

var (
	// giraffeNameAmericanEnglish represents the Giraffe animal type's name in American English.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Giraffe",
	}

	// giraffeNameCanadianFrench represents the Giraffe animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Giraffe",
	}

	// giraffeNameDutch represents the Giraffe animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Giraffe",
	}

	// giraffeNameFrench represents the Giraffe animal type's name in French.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameFrench = nook.Name{
		Language: language.French,
		Value:    "Giraffe",
	}

	// giraffeNameGerman represents the Giraffe animal type's name in German.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameGerman = nook.Name{
		Language: language.German,
		Value:    "Giraffe",
	}

	// giraffeNameItalian represents the Giraffe animal type's name in Italian.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Giraffe",
	}

	// giraffeNameJapanese represents the Giraffe animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Giraffe",
	}

	// giraffeNameKorean represents the Giraffe animal type's name in Korean.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Giraffe",
	}

	// giraffeNameLatinAmericanSpanish represents the Giraffe animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Giraffe",
	}

	// giraffeNameRussian represents the Giraffe animal type's name in Russian.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Giraffe",
	}

	// giraffeNameSimplifiedChinese represents the Giraffe animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Giraffe",
	}

	// giraffeNameSpanish represents the Giraffe animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Giraffe",
	}

	// giraffeNameTraditionalChinese represents the Giraffe animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Giraffe"
	giraffeNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Giraffe",
	}
)

var (
	// giraffeName contains the localized names of the Giraffe animal type.
	giraffeName = nook.Languages{
		language.AmericanEnglish:      giraffeNameAmericanEnglish,
		language.CanadianFrench:       giraffeNameCanadianFrench,
		language.Dutch:                giraffeNameDutch,
		language.French:               giraffeNameFrench,
		language.German:               giraffeNameGerman,
		language.Italian:              giraffeNameItalian,
		language.Japanese:             giraffeNameJapanese,
		language.Korean:               giraffeNameKorean,
		language.LatinAmericanSpanish: giraffeNameLatinAmericanSpanish,
		language.Russian:              giraffeNameRussian,
		language.SimplifiedChinese:    giraffeNameSimplifiedChinese,
		language.Spanish:              giraffeNameSpanish,
		language.TraditionalChinese:   giraffeNameTraditionalChinese,
	}
)

var (
	// Giraffe represents the Giraffe animal type in the Animal Crossing series.
	Giraffe = nook.Animal{
		Key:  nook.Key(giraffe),
		Name: giraffeName,
	}
)
