package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// hippo is the common reference for the Hippo animal type.
	hippo = "Hippo"
)

var (
	// hippoNameAmericanEnglish represents the Hippo animal type's name in American English.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hippo",
	}

	// hippoNameCanadianFrench represents the Hippo animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hippo",
	}

	// hippoNameDutch represents the Hippo animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Hippo",
	}

	// hippoNameFrench represents the Hippo animal type's name in French.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameFrench = nook.Name{
		Language: language.French,
		Value:    "Hippo",
	}

	// hippoNameGerman represents the Hippo animal type's name in German.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameGerman = nook.Name{
		Language: language.German,
		Value:    "Hippo",
	}

	// hippoNameItalian represents the Hippo animal type's name in Italian.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Hippo",
	}

	// hippoNameJapanese represents the Hippo animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Hippo",
	}

	// hippoNameKorean represents the Hippo animal type's name in Korean.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Hippo",
	}

	// hippoNameLatinAmericanSpanish represents the Hippo animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hippo",
	}

	// hippoNameRussian represents the Hippo animal type's name in Russian.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Hippo",
	}

	// hippoNameSimplifiedChinese represents the Hippo animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Hippo",
	}

	// hippoNameSpanish represents the Hippo animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Hippo",
	}

	// hippoNameTraditionalChinese represents the Hippo animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Hippo"
	hippoNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Hippo",
	}
)

var (
	// hippoName contains the localized names of the Hippo animal type.
	hippoName = nook.Languages{
		language.AmericanEnglish:      hippoNameAmericanEnglish,
		language.CanadianFrench:       hippoNameCanadianFrench,
		language.Dutch:                hippoNameDutch,
		language.French:               hippoNameFrench,
		language.German:               hippoNameGerman,
		language.Italian:              hippoNameItalian,
		language.Japanese:             hippoNameJapanese,
		language.Korean:               hippoNameKorean,
		language.LatinAmericanSpanish: hippoNameLatinAmericanSpanish,
		language.Russian:              hippoNameRussian,
		language.SimplifiedChinese:    hippoNameSimplifiedChinese,
		language.Spanish:              hippoNameSpanish,
		language.TraditionalChinese:   hippoNameTraditionalChinese,
	}
)

var (
	// Hippo represents the Hippo animal type in the Animal Crossing series.
	Hippo = nook.Animal{
		Key:  nook.Key(hippo),
		Name: hippoName,
	}
)
