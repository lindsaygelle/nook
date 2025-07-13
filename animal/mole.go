package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// mole is the common reference for the Mole animal type.
	mole = "Mole"
)

var (
	// moleNameAmericanEnglish represents the Mole animal type's name in American English.
	// TODO: Verify translation accuracy for "Mole"
	moleNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mole",
	}

	// moleNameCanadianFrench represents the Mole animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Mole"
	moleNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mole",
	}

	// moleNameDutch represents the Mole animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Mole"
	moleNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Mole",
	}

	// moleNameFrench represents the Mole animal type's name in French.
	// TODO: Verify translation accuracy for "Mole"
	moleNameFrench = nook.Name{
		Language: language.French,
		Value:    "Mole",
	}

	// moleNameGerman represents the Mole animal type's name in German.
	// TODO: Verify translation accuracy for "Mole"
	moleNameGerman = nook.Name{
		Language: language.German,
		Value:    "Mole",
	}

	// moleNameItalian represents the Mole animal type's name in Italian.
	// TODO: Verify translation accuracy for "Mole"
	moleNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Mole",
	}

	// moleNameJapanese represents the Mole animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Mole"
	moleNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Mole",
	}

	// moleNameKorean represents the Mole animal type's name in Korean.
	// TODO: Verify translation accuracy for "Mole"
	moleNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Mole",
	}

	// moleNameLatinAmericanSpanish represents the Mole animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Mole"
	moleNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mole",
	}

	// moleNameRussian represents the Mole animal type's name in Russian.
	// TODO: Verify translation accuracy for "Mole"
	moleNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Mole",
	}

	// moleNameSimplifiedChinese represents the Mole animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Mole"
	moleNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Mole",
	}

	// moleNameSpanish represents the Mole animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Mole"
	moleNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Mole",
	}

	// moleNameTraditionalChinese represents the Mole animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Mole"
	moleNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Mole",
	}
)

var (
	// moleName contains the localized names of the Mole animal type.
	moleName = nook.Languages{
		language.AmericanEnglish:      moleNameAmericanEnglish,
		language.CanadianFrench:       moleNameCanadianFrench,
		language.Dutch:                moleNameDutch,
		language.French:               moleNameFrench,
		language.German:               moleNameGerman,
		language.Italian:              moleNameItalian,
		language.Japanese:             moleNameJapanese,
		language.Korean:               moleNameKorean,
		language.LatinAmericanSpanish: moleNameLatinAmericanSpanish,
		language.Russian:              moleNameRussian,
		language.SimplifiedChinese:    moleNameSimplifiedChinese,
		language.Spanish:              moleNameSpanish,
		language.TraditionalChinese:   moleNameTraditionalChinese,
	}
)

var (
	// Mole represents the Mole animal type in the Animal Crossing series.
	Mole = nook.Animal{
		Key:  nook.Key(mole),
		Name: moleName,
	}
)
