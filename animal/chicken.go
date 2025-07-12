package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// chicken is the common reference for the Chicken animal type.
	chicken = "Chicken"
)

var (
	// chickenNameAmericanEnglish represents the Chicken animal type's name in American English.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chicken",
	}

	// chickenNameCanadianFrench represents the Chicken animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chicken",
	}

	// chickenNameDutch represents the Chicken animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Chicken",
	}

	// chickenNameFrench represents the Chicken animal type's name in French.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameFrench = nook.Name{
		Language: language.French,
		Value:    "Chicken",
	}

	// chickenNameGerman represents the Chicken animal type's name in German.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameGerman = nook.Name{
		Language: language.German,
		Value:    "Chicken",
	}

	// chickenNameItalian represents the Chicken animal type's name in Italian.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Chicken",
	}

	// chickenNameJapanese represents the Chicken animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Chicken",
	}

	// chickenNameKorean represents the Chicken animal type's name in Korean.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Chicken",
	}

	// chickenNameLatinAmericanSpanish represents the Chicken animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chicken",
	}

	// chickenNameRussian represents the Chicken animal type's name in Russian.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Chicken",
	}

	// chickenNameSimplifiedChinese represents the Chicken animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Chicken",
	}

	// chickenNameSpanish represents the Chicken animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Chicken",
	}

	// chickenNameTraditionalChinese represents the Chicken animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Chicken"
	chickenNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Chicken",
	}
)

var (
	// chickenName contains the localized names of the Chicken animal type.
	chickenName = nook.Languages{
		language.AmericanEnglish:      chickenNameAmericanEnglish,
		language.CanadianFrench:       chickenNameCanadianFrench,
		language.Dutch:                chickenNameDutch,
		language.French:               chickenNameFrench,
		language.German:               chickenNameGerman,
		language.Italian:              chickenNameItalian,
		language.Japanese:             chickenNameJapanese,
		language.Korean:               chickenNameKorean,
		language.LatinAmericanSpanish: chickenNameLatinAmericanSpanish,
		language.Russian:              chickenNameRussian,
		language.SimplifiedChinese:    chickenNameSimplifiedChinese,
		language.Spanish:              chickenNameSpanish,
		language.TraditionalChinese:   chickenNameTraditionalChinese,
	}
)

var (
	// Chicken represents the Chicken animal type in the Animal Crossing series.
	Chicken = nook.Animal{
		Key:  nook.Key(chicken),
		Name: chickenName,
	}
)
