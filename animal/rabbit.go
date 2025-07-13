package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// rabbit is the common reference for the Rabbit animal type.
	rabbit = "Rabbit"
)

var (
	// rabbitNameAmericanEnglish represents the Rabbit animal type's name in American English.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rabbit",
	}

	// rabbitNameCanadianFrench represents the Rabbit animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rabbit",
	}

	// rabbitNameDutch represents the Rabbit animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Rabbit",
	}

	// rabbitNameFrench represents the Rabbit animal type's name in French.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameFrench = nook.Name{
		Language: language.French,
		Value:    "Rabbit",
	}

	// rabbitNameGerman represents the Rabbit animal type's name in German.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameGerman = nook.Name{
		Language: language.German,
		Value:    "Rabbit",
	}

	// rabbitNameItalian represents the Rabbit animal type's name in Italian.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Rabbit",
	}

	// rabbitNameJapanese represents the Rabbit animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Rabbit",
	}

	// rabbitNameKorean represents the Rabbit animal type's name in Korean.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Rabbit",
	}

	// rabbitNameLatinAmericanSpanish represents the Rabbit animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rabbit",
	}

	// rabbitNameRussian represents the Rabbit animal type's name in Russian.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Rabbit",
	}

	// rabbitNameSimplifiedChinese represents the Rabbit animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Rabbit",
	}

	// rabbitNameSpanish represents the Rabbit animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Rabbit",
	}

	// rabbitNameTraditionalChinese represents the Rabbit animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Rabbit"
	rabbitNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Rabbit",
	}
)

var (
	// rabbitName contains the localized names of the Rabbit animal type.
	rabbitName = nook.Languages{
		language.AmericanEnglish:      rabbitNameAmericanEnglish,
		language.CanadianFrench:       rabbitNameCanadianFrench,
		language.Dutch:                rabbitNameDutch,
		language.French:               rabbitNameFrench,
		language.German:               rabbitNameGerman,
		language.Italian:              rabbitNameItalian,
		language.Japanese:             rabbitNameJapanese,
		language.Korean:               rabbitNameKorean,
		language.LatinAmericanSpanish: rabbitNameLatinAmericanSpanish,
		language.Russian:              rabbitNameRussian,
		language.SimplifiedChinese:    rabbitNameSimplifiedChinese,
		language.Spanish:              rabbitNameSpanish,
		language.TraditionalChinese:   rabbitNameTraditionalChinese,
	}
)

var (
	// Rabbit represents the Rabbit animal type in the Animal Crossing series.
	Rabbit = nook.Animal{
		Key:  nook.Key(rabbit),
		Name: rabbitName,
	}
)
