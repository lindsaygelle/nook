package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// alligator is the common reference for the Alligator animal type.
	alligator = "Alligator"
)

var (
	// alligatorNameAmericanEnglish represents the Alligator animal type's name in American English.
	alligatorNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Alligator",
	}

	// alligatorNameCanadianFrench represents the Alligator animal type's name in Canadian French.
	alligatorNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Crocodile",
	}

	// alligatorNameDutch represents the Alligator animal type's name in Dutch.
	alligatorNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Krokodil",
	}

	// alligatorNameFrench represents the Alligator animal type's name in French.
	alligatorNameFrench = nook.Name{
		Language: language.French,
		Value:    "Crocodile",
	}

	// alligatorNameGerman represents the Alligator animal type's name in German.
	alligatorNameGerman = nook.Name{
		Language: language.German,
		Value:    "Krokodil",
	}

	// alligatorNameItalian represents the Alligator animal type's name in Italian.
	alligatorNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Coccodrillo",
	}

	// alligatorNameJapanese represents the Alligator animal type's name in Japanese.
	alligatorNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "ワニ",
	}

	// alligatorNameKorean represents the Alligator animal type's name in Korean.
	alligatorNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "악어",
	}

	// alligatorNameLatinAmericanSpanish represents the Alligator animal type's name in Latin American Spanish.
	alligatorNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cocodrilo",
	}

	// alligatorNameRussian represents the Alligator animal type's name in Russian.
	alligatorNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Крокодил",
	}

	// alligatorNameSimplifiedChinese represents the Alligator animal type's name in Simplified Chinese.
	alligatorNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鳄鱼",
	}

	// alligatorNameSpanish represents the Alligator animal type's name in Spanish.
	alligatorNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Cocodrilo",
	}

	// alligatorNameTraditionalChinese represents the Alligator animal type's name in Traditional Chinese.
	alligatorNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鱷魚",
	}
)

var (
	// alligatorName contains the localized names of the Alligator animal type.
	alligatorName = nook.Languages{
		language.AmericanEnglish:      alligatorNameAmericanEnglish,
		language.CanadianFrench:       alligatorNameCanadianFrench,
		language.Dutch:                alligatorNameDutch,
		language.French:               alligatorNameFrench,
		language.German:               alligatorNameGerman,
		language.Italian:              alligatorNameItalian,
		language.Japanese:             alligatorNameJapanese,
		language.Korean:               alligatorNameKorean,
		language.LatinAmericanSpanish: alligatorNameLatinAmericanSpanish,
		language.Russian:              alligatorNameRussian,
		language.SimplifiedChinese:    alligatorNameSimplifiedChinese,
		language.Spanish:              alligatorNameSpanish,
		language.TraditionalChinese:   alligatorNameTraditionalChinese,
	}
)

var (
	// Alligator represents the Alligator animal type in the Animal Crossing series.
	Alligator = nook.Animal{
		Key:  nook.Key(alligator),
		Name: alligatorName,
	}
)
