package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// gyroid is the common reference for the Gyroid animal type.
	gyroid = "Gyroid"
)

var (
	// gyroidNameAmericanEnglish represents the Gyroid animal type's name in American English.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gyroid",
	}

	// gyroidNameCanadianFrench represents the Gyroid animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gyroid",
	}

	// gyroidNameDutch represents the Gyroid animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Gyroid",
	}

	// gyroidNameFrench represents the Gyroid animal type's name in French.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameFrench = nook.Name{
		Language: language.French,
		Value:    "Gyroid",
	}

	// gyroidNameGerman represents the Gyroid animal type's name in German.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameGerman = nook.Name{
		Language: language.German,
		Value:    "Gyroid",
	}

	// gyroidNameItalian represents the Gyroid animal type's name in Italian.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Gyroid",
	}

	// gyroidNameJapanese represents the Gyroid animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Gyroid",
	}

	// gyroidNameKorean represents the Gyroid animal type's name in Korean.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Gyroid",
	}

	// gyroidNameLatinAmericanSpanish represents the Gyroid animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gyroid",
	}

	// gyroidNameRussian represents the Gyroid animal type's name in Russian.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Gyroid",
	}

	// gyroidNameSimplifiedChinese represents the Gyroid animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Gyroid",
	}

	// gyroidNameSpanish represents the Gyroid animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Gyroid",
	}

	// gyroidNameTraditionalChinese represents the Gyroid animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Gyroid"
	gyroidNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Gyroid",
	}
)

var (
	// gyroidName contains the localized names of the Gyroid animal type.
	gyroidName = nook.Languages{
		language.AmericanEnglish:      gyroidNameAmericanEnglish,
		language.CanadianFrench:       gyroidNameCanadianFrench,
		language.Dutch:                gyroidNameDutch,
		language.French:               gyroidNameFrench,
		language.German:               gyroidNameGerman,
		language.Italian:              gyroidNameItalian,
		language.Japanese:             gyroidNameJapanese,
		language.Korean:               gyroidNameKorean,
		language.LatinAmericanSpanish: gyroidNameLatinAmericanSpanish,
		language.Russian:              gyroidNameRussian,
		language.SimplifiedChinese:    gyroidNameSimplifiedChinese,
		language.Spanish:              gyroidNameSpanish,
		language.TraditionalChinese:   gyroidNameTraditionalChinese,
	}
)

var (
	// Gyroid represents the Gyroid animal type in the Animal Crossing series.
	Gyroid = nook.Animal{
		Key:  nook.Key(gyroid),
		Name: gyroidName,
	}
)
