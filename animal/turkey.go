package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// turkey is the common reference for the Turkey animal type.
	turkey = "Turkey"
)

var (
	// turkeyNameAmericanEnglish represents the Turkey animal type's name in American English.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Turkey",
	}

	// turkeyNameCanadianFrench represents the Turkey animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Turkey",
	}

	// turkeyNameDutch represents the Turkey animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Turkey",
	}

	// turkeyNameFrench represents the Turkey animal type's name in French.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameFrench = nook.Name{
		Language: language.French,
		Value:    "Turkey",
	}

	// turkeyNameGerman represents the Turkey animal type's name in German.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameGerman = nook.Name{
		Language: language.German,
		Value:    "Turkey",
	}

	// turkeyNameItalian represents the Turkey animal type's name in Italian.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Turkey",
	}

	// turkeyNameJapanese represents the Turkey animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Turkey",
	}

	// turkeyNameKorean represents the Turkey animal type's name in Korean.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Turkey",
	}

	// turkeyNameLatinAmericanSpanish represents the Turkey animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Turkey",
	}

	// turkeyNameRussian represents the Turkey animal type's name in Russian.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Turkey",
	}

	// turkeyNameSimplifiedChinese represents the Turkey animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Turkey",
	}

	// turkeyNameSpanish represents the Turkey animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Turkey",
	}

	// turkeyNameTraditionalChinese represents the Turkey animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Turkey"
	turkeyNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Turkey",
	}
)

var (
	// turkeyName contains the localized names of the Turkey animal type.
	turkeyName = nook.Languages{
		language.AmericanEnglish:      turkeyNameAmericanEnglish,
		language.CanadianFrench:       turkeyNameCanadianFrench,
		language.Dutch:                turkeyNameDutch,
		language.French:               turkeyNameFrench,
		language.German:               turkeyNameGerman,
		language.Italian:              turkeyNameItalian,
		language.Japanese:             turkeyNameJapanese,
		language.Korean:               turkeyNameKorean,
		language.LatinAmericanSpanish: turkeyNameLatinAmericanSpanish,
		language.Russian:              turkeyNameRussian,
		language.SimplifiedChinese:    turkeyNameSimplifiedChinese,
		language.Spanish:              turkeyNameSpanish,
		language.TraditionalChinese:   turkeyNameTraditionalChinese,
	}
)

var (
	// Turkey represents the Turkey animal type in the Animal Crossing series.
	Turkey = nook.Animal{
		Key:  nook.Key(turkey),
		Name: turkeyName,
	}
)
