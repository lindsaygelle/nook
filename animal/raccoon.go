package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// raccoon is the common reference for the Raccoon animal type.
	raccoon = "Raccoon"
)

var (
	// raccoonNameAmericanEnglish represents the Raccoon animal type's name in American English.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Raccoon",
	}

	// raccoonNameCanadianFrench represents the Raccoon animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Raccoon",
	}

	// raccoonNameDutch represents the Raccoon animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Raccoon",
	}

	// raccoonNameFrench represents the Raccoon animal type's name in French.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameFrench = nook.Name{
		Language: language.French,
		Value:    "Raccoon",
	}

	// raccoonNameGerman represents the Raccoon animal type's name in German.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameGerman = nook.Name{
		Language: language.German,
		Value:    "Raccoon",
	}

	// raccoonNameItalian represents the Raccoon animal type's name in Italian.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Raccoon",
	}

	// raccoonNameJapanese represents the Raccoon animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Raccoon",
	}

	// raccoonNameKorean represents the Raccoon animal type's name in Korean.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Raccoon",
	}

	// raccoonNameLatinAmericanSpanish represents the Raccoon animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Raccoon",
	}

	// raccoonNameRussian represents the Raccoon animal type's name in Russian.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Raccoon",
	}

	// raccoonNameSimplifiedChinese represents the Raccoon animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Raccoon",
	}

	// raccoonNameSpanish represents the Raccoon animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Raccoon",
	}

	// raccoonNameTraditionalChinese represents the Raccoon animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Raccoon"
	raccoonNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Raccoon",
	}
)

var (
	// raccoonName contains the localized names of the Raccoon animal type.
	raccoonName = nook.Languages{
		language.AmericanEnglish:      raccoonNameAmericanEnglish,
		language.CanadianFrench:       raccoonNameCanadianFrench,
		language.Dutch:                raccoonNameDutch,
		language.French:               raccoonNameFrench,
		language.German:               raccoonNameGerman,
		language.Italian:              raccoonNameItalian,
		language.Japanese:             raccoonNameJapanese,
		language.Korean:               raccoonNameKorean,
		language.LatinAmericanSpanish: raccoonNameLatinAmericanSpanish,
		language.Russian:              raccoonNameRussian,
		language.SimplifiedChinese:    raccoonNameSimplifiedChinese,
		language.Spanish:              raccoonNameSpanish,
		language.TraditionalChinese:   raccoonNameTraditionalChinese,
	}
)

var (
	// Raccoon represents the Raccoon animal type in the Animal Crossing series.
	Raccoon = nook.Animal{
		Key:  nook.Key(raccoon),
		Name: raccoonName,
	}
)
