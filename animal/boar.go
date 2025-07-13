package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// boar is the common reference for the Boar animal type.
	boar = "Boar"
)

var (
	// boarNameAmericanEnglish represents the Boar animal type's name in American English.
	// TODO: Verify translation accuracy for "Boar"
	boarNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Boar",
	}

	// boarNameCanadianFrench represents the Boar animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Boar"
	boarNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Boar",
	}

	// boarNameDutch represents the Boar animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Boar"
	boarNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Boar",
	}

	// boarNameFrench represents the Boar animal type's name in French.
	// TODO: Verify translation accuracy for "Boar"
	boarNameFrench = nook.Name{
		Language: language.French,
		Value:    "Boar",
	}

	// boarNameGerman represents the Boar animal type's name in German.
	// TODO: Verify translation accuracy for "Boar"
	boarNameGerman = nook.Name{
		Language: language.German,
		Value:    "Boar",
	}

	// boarNameItalian represents the Boar animal type's name in Italian.
	// TODO: Verify translation accuracy for "Boar"
	boarNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Boar",
	}

	// boarNameJapanese represents the Boar animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Boar"
	boarNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Boar",
	}

	// boarNameKorean represents the Boar animal type's name in Korean.
	// TODO: Verify translation accuracy for "Boar"
	boarNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Boar",
	}

	// boarNameLatinAmericanSpanish represents the Boar animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Boar"
	boarNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Boar",
	}

	// boarNameRussian represents the Boar animal type's name in Russian.
	// TODO: Verify translation accuracy for "Boar"
	boarNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Boar",
	}

	// boarNameSimplifiedChinese represents the Boar animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Boar"
	boarNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Boar",
	}

	// boarNameSpanish represents the Boar animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Boar"
	boarNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Boar",
	}

	// boarNameTraditionalChinese represents the Boar animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Boar"
	boarNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Boar",
	}
)

var (
	// boarName contains the localized names of the Boar animal type.
	boarName = nook.Languages{
		language.AmericanEnglish:      boarNameAmericanEnglish,
		language.CanadianFrench:       boarNameCanadianFrench,
		language.Dutch:                boarNameDutch,
		language.French:               boarNameFrench,
		language.German:               boarNameGerman,
		language.Italian:              boarNameItalian,
		language.Japanese:             boarNameJapanese,
		language.Korean:               boarNameKorean,
		language.LatinAmericanSpanish: boarNameLatinAmericanSpanish,
		language.Russian:              boarNameRussian,
		language.SimplifiedChinese:    boarNameSimplifiedChinese,
		language.Spanish:              boarNameSpanish,
		language.TraditionalChinese:   boarNameTraditionalChinese,
	}
)

var (
	// Boar represents the Boar animal type in the Animal Crossing series.
	Boar = nook.Animal{
		Key:  nook.Key(boar),
		Name: boarName,
	}
)
