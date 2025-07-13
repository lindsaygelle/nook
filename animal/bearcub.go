package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// bearcub is the common reference for the Bearcub animal type.
	bearcub = "Bearcub"
)

var (
	// bearcubNameAmericanEnglish represents the Bearcub animal type's name in American English.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bearcub",
	}

	// bearcubNameCanadianFrench represents the Bearcub animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bearcub",
	}

	// bearcubNameDutch represents the Bearcub animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Bearcub",
	}

	// bearcubNameFrench represents the Bearcub animal type's name in French.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameFrench = nook.Name{
		Language: language.French,
		Value:    "Bearcub",
	}

	// bearcubNameGerman represents the Bearcub animal type's name in German.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameGerman = nook.Name{
		Language: language.German,
		Value:    "Bearcub",
	}

	// bearcubNameItalian represents the Bearcub animal type's name in Italian.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Bearcub",
	}

	// bearcubNameJapanese represents the Bearcub animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Bearcub",
	}

	// bearcubNameKorean represents the Bearcub animal type's name in Korean.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Bearcub",
	}

	// bearcubNameLatinAmericanSpanish represents the Bearcub animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bearcub",
	}

	// bearcubNameRussian represents the Bearcub animal type's name in Russian.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Bearcub",
	}

	// bearcubNameSimplifiedChinese represents the Bearcub animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Bearcub",
	}

	// bearcubNameSpanish represents the Bearcub animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Bearcub",
	}

	// bearcubNameTraditionalChinese represents the Bearcub animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Bearcub"
	bearcubNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Bearcub",
	}
)

var (
	// bearcubName contains the localized names of the Bearcub animal type.
	bearcubName = nook.Languages{
		language.AmericanEnglish:      bearcubNameAmericanEnglish,
		language.CanadianFrench:       bearcubNameCanadianFrench,
		language.Dutch:                bearcubNameDutch,
		language.French:               bearcubNameFrench,
		language.German:               bearcubNameGerman,
		language.Italian:              bearcubNameItalian,
		language.Japanese:             bearcubNameJapanese,
		language.Korean:               bearcubNameKorean,
		language.LatinAmericanSpanish: bearcubNameLatinAmericanSpanish,
		language.Russian:              bearcubNameRussian,
		language.SimplifiedChinese:    bearcubNameSimplifiedChinese,
		language.Spanish:              bearcubNameSpanish,
		language.TraditionalChinese:   bearcubNameTraditionalChinese,
	}
)

var (
	// Bearcub represents the Bearcub animal type in the Animal Crossing series.
	Bearcub = nook.Animal{
		Key:  nook.Key(bearcub),
		Name: bearcubName,
	}
)
