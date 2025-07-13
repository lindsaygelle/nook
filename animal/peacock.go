package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// peacock is the common reference for the Peacock animal type.
	peacock = "Peacock"
)

var (
	// peacockNameAmericanEnglish represents the Peacock animal type's name in American English.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peacock",
	}

	// peacockNameCanadianFrench represents the Peacock animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Peacock",
	}

	// peacockNameDutch represents the Peacock animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Peacock",
	}

	// peacockNameFrench represents the Peacock animal type's name in French.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameFrench = nook.Name{
		Language: language.French,
		Value:    "Peacock",
	}

	// peacockNameGerman represents the Peacock animal type's name in German.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameGerman = nook.Name{
		Language: language.German,
		Value:    "Peacock",
	}

	// peacockNameItalian represents the Peacock animal type's name in Italian.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Peacock",
	}

	// peacockNameJapanese represents the Peacock animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Peacock",
	}

	// peacockNameKorean represents the Peacock animal type's name in Korean.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Peacock",
	}

	// peacockNameLatinAmericanSpanish represents the Peacock animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Peacock",
	}

	// peacockNameRussian represents the Peacock animal type's name in Russian.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Peacock",
	}

	// peacockNameSimplifiedChinese represents the Peacock animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Peacock",
	}

	// peacockNameSpanish represents the Peacock animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Peacock",
	}

	// peacockNameTraditionalChinese represents the Peacock animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Peacock"
	peacockNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Peacock",
	}
)

var (
	// peacockName contains the localized names of the Peacock animal type.
	peacockName = nook.Languages{
		language.AmericanEnglish:      peacockNameAmericanEnglish,
		language.CanadianFrench:       peacockNameCanadianFrench,
		language.Dutch:                peacockNameDutch,
		language.French:               peacockNameFrench,
		language.German:               peacockNameGerman,
		language.Italian:              peacockNameItalian,
		language.Japanese:             peacockNameJapanese,
		language.Korean:               peacockNameKorean,
		language.LatinAmericanSpanish: peacockNameLatinAmericanSpanish,
		language.Russian:              peacockNameRussian,
		language.SimplifiedChinese:    peacockNameSimplifiedChinese,
		language.Spanish:              peacockNameSpanish,
		language.TraditionalChinese:   peacockNameTraditionalChinese,
	}
)

var (
	// Peacock represents the Peacock animal type in the Animal Crossing series.
	Peacock = nook.Animal{
		Key:  nook.Key(peacock),
		Name: peacockName,
	}
)
