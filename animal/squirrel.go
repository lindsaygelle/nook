package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// squirrel is the common reference for the Squirrel animal type.
	squirrel = "Squirrel"
)

var (
	// squirrelNameAmericanEnglish represents the Squirrel animal type's name in American English.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Squirrel",
	}

	// squirrelNameCanadianFrench represents the Squirrel animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Squirrel",
	}

	// squirrelNameDutch represents the Squirrel animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Squirrel",
	}

	// squirrelNameFrench represents the Squirrel animal type's name in French.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameFrench = nook.Name{
		Language: language.French,
		Value:    "Squirrel",
	}

	// squirrelNameGerman represents the Squirrel animal type's name in German.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameGerman = nook.Name{
		Language: language.German,
		Value:    "Squirrel",
	}

	// squirrelNameItalian represents the Squirrel animal type's name in Italian.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Squirrel",
	}

	// squirrelNameJapanese represents the Squirrel animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Squirrel",
	}

	// squirrelNameKorean represents the Squirrel animal type's name in Korean.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Squirrel",
	}

	// squirrelNameLatinAmericanSpanish represents the Squirrel animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Squirrel",
	}

	// squirrelNameRussian represents the Squirrel animal type's name in Russian.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Squirrel",
	}

	// squirrelNameSimplifiedChinese represents the Squirrel animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Squirrel",
	}

	// squirrelNameSpanish represents the Squirrel animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Squirrel",
	}

	// squirrelNameTraditionalChinese represents the Squirrel animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Squirrel"
	squirrelNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Squirrel",
	}
)

var (
	// squirrelName contains the localized names of the Squirrel animal type.
	squirrelName = nook.Languages{
		language.AmericanEnglish:      squirrelNameAmericanEnglish,
		language.CanadianFrench:       squirrelNameCanadianFrench,
		language.Dutch:                squirrelNameDutch,
		language.French:               squirrelNameFrench,
		language.German:               squirrelNameGerman,
		language.Italian:              squirrelNameItalian,
		language.Japanese:             squirrelNameJapanese,
		language.Korean:               squirrelNameKorean,
		language.LatinAmericanSpanish: squirrelNameLatinAmericanSpanish,
		language.Russian:              squirrelNameRussian,
		language.SimplifiedChinese:    squirrelNameSimplifiedChinese,
		language.Spanish:              squirrelNameSpanish,
		language.TraditionalChinese:   squirrelNameTraditionalChinese,
	}
)

var (
	// Squirrel represents the Squirrel animal type in the Animal Crossing series.
	Squirrel = nook.Animal{
		Key:  nook.Key(squirrel),
		Name: squirrelName,
	}
)
