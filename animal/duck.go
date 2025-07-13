package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// duck is the common reference for the Duck animal type.
	duck = "Duck"
)

var (
	// duckNameAmericanEnglish represents the Duck animal type's name in American English.
	// TODO: Verify translation accuracy for "Duck"
	duckNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Duck",
	}

	// duckNameCanadianFrench represents the Duck animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Duck"
	duckNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Duck",
	}

	// duckNameDutch represents the Duck animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Duck"
	duckNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Duck",
	}

	// duckNameFrench represents the Duck animal type's name in French.
	// TODO: Verify translation accuracy for "Duck"
	duckNameFrench = nook.Name{
		Language: language.French,
		Value:    "Duck",
	}

	// duckNameGerman represents the Duck animal type's name in German.
	// TODO: Verify translation accuracy for "Duck"
	duckNameGerman = nook.Name{
		Language: language.German,
		Value:    "Duck",
	}

	// duckNameItalian represents the Duck animal type's name in Italian.
	// TODO: Verify translation accuracy for "Duck"
	duckNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Duck",
	}

	// duckNameJapanese represents the Duck animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Duck"
	duckNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Duck",
	}

	// duckNameKorean represents the Duck animal type's name in Korean.
	// TODO: Verify translation accuracy for "Duck"
	duckNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Duck",
	}

	// duckNameLatinAmericanSpanish represents the Duck animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Duck"
	duckNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Duck",
	}

	// duckNameRussian represents the Duck animal type's name in Russian.
	// TODO: Verify translation accuracy for "Duck"
	duckNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Duck",
	}

	// duckNameSimplifiedChinese represents the Duck animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Duck"
	duckNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Duck",
	}

	// duckNameSpanish represents the Duck animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Duck"
	duckNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Duck",
	}

	// duckNameTraditionalChinese represents the Duck animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Duck"
	duckNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Duck",
	}
)

var (
	// duckName contains the localized names of the Duck animal type.
	duckName = nook.Languages{
		language.AmericanEnglish:      duckNameAmericanEnglish,
		language.CanadianFrench:       duckNameCanadianFrench,
		language.Dutch:                duckNameDutch,
		language.French:               duckNameFrench,
		language.German:               duckNameGerman,
		language.Italian:              duckNameItalian,
		language.Japanese:             duckNameJapanese,
		language.Korean:               duckNameKorean,
		language.LatinAmericanSpanish: duckNameLatinAmericanSpanish,
		language.Russian:              duckNameRussian,
		language.SimplifiedChinese:    duckNameSimplifiedChinese,
		language.Spanish:              duckNameSpanish,
		language.TraditionalChinese:   duckNameTraditionalChinese,
	}
)

var (
	// Duck represents the Duck animal type in the Animal Crossing series.
	Duck = nook.Animal{
		Key:  nook.Key(duck),
		Name: duckName,
	}
)
