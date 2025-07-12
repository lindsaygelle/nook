package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// cat is the common reference for the Cat animal type.
	cat = "Cat"
)

var (
	// catNameAmericanEnglish represents the Cat animal type's name in American English.
	// TODO: Verify translation accuracy for "Cat"
	catNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cat",
	}

	// catNameCanadianFrench represents the Cat animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Cat"
	catNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cat",
	}

	// catNameDutch represents the Cat animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Cat"
	catNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Cat",
	}

	// catNameFrench represents the Cat animal type's name in French.
	// TODO: Verify translation accuracy for "Cat"
	catNameFrench = nook.Name{
		Language: language.French,
		Value:    "Cat",
	}

	// catNameGerman represents the Cat animal type's name in German.
	// TODO: Verify translation accuracy for "Cat"
	catNameGerman = nook.Name{
		Language: language.German,
		Value:    "Cat",
	}

	// catNameItalian represents the Cat animal type's name in Italian.
	// TODO: Verify translation accuracy for "Cat"
	catNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Cat",
	}

	// catNameJapanese represents the Cat animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Cat"
	catNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Cat",
	}

	// catNameKorean represents the Cat animal type's name in Korean.
	// TODO: Verify translation accuracy for "Cat"
	catNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Cat",
	}

	// catNameLatinAmericanSpanish represents the Cat animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Cat"
	catNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cat",
	}

	// catNameRussian represents the Cat animal type's name in Russian.
	// TODO: Verify translation accuracy for "Cat"
	catNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Cat",
	}

	// catNameSimplifiedChinese represents the Cat animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Cat"
	catNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Cat",
	}

	// catNameSpanish represents the Cat animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Cat"
	catNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Cat",
	}

	// catNameTraditionalChinese represents the Cat animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Cat"
	catNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Cat",
	}
)

var (
	// catName contains the localized names of the Cat animal type.
	catName = nook.Languages{
		language.AmericanEnglish:      catNameAmericanEnglish,
		language.CanadianFrench:       catNameCanadianFrench,
		language.Dutch:                catNameDutch,
		language.French:               catNameFrench,
		language.German:               catNameGerman,
		language.Italian:              catNameItalian,
		language.Japanese:             catNameJapanese,
		language.Korean:               catNameKorean,
		language.LatinAmericanSpanish: catNameLatinAmericanSpanish,
		language.Russian:              catNameRussian,
		language.SimplifiedChinese:    catNameSimplifiedChinese,
		language.Spanish:              catNameSpanish,
		language.TraditionalChinese:   catNameTraditionalChinese,
	}
)

var (
	// Cat represents the Cat animal type in the Animal Crossing series.
	Cat = nook.Animal{
		Key:  nook.Key(cat),
		Name: catName,
	}
)
