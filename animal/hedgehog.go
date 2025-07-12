package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// hedgehog is the common reference for the Hedgehog animal type.
	hedgehog = "Hedgehog"
)

var (
	// hedgehogNameAmericanEnglish represents the Hedgehog animal type's name in American English.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hedgehog",
	}

	// hedgehogNameCanadianFrench represents the Hedgehog animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hedgehog",
	}

	// hedgehogNameDutch represents the Hedgehog animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Hedgehog",
	}

	// hedgehogNameFrench represents the Hedgehog animal type's name in French.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameFrench = nook.Name{
		Language: language.French,
		Value:    "Hedgehog",
	}

	// hedgehogNameGerman represents the Hedgehog animal type's name in German.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameGerman = nook.Name{
		Language: language.German,
		Value:    "Hedgehog",
	}

	// hedgehogNameItalian represents the Hedgehog animal type's name in Italian.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Hedgehog",
	}

	// hedgehogNameJapanese represents the Hedgehog animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Hedgehog",
	}

	// hedgehogNameKorean represents the Hedgehog animal type's name in Korean.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Hedgehog",
	}

	// hedgehogNameLatinAmericanSpanish represents the Hedgehog animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hedgehog",
	}

	// hedgehogNameRussian represents the Hedgehog animal type's name in Russian.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Hedgehog",
	}

	// hedgehogNameSimplifiedChinese represents the Hedgehog animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Hedgehog",
	}

	// hedgehogNameSpanish represents the Hedgehog animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Hedgehog",
	}

	// hedgehogNameTraditionalChinese represents the Hedgehog animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Hedgehog"
	hedgehogNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Hedgehog",
	}
)

var (
	// hedgehogName contains the localized names of the Hedgehog animal type.
	hedgehogName = nook.Languages{
		language.AmericanEnglish:      hedgehogNameAmericanEnglish,
		language.CanadianFrench:       hedgehogNameCanadianFrench,
		language.Dutch:                hedgehogNameDutch,
		language.French:               hedgehogNameFrench,
		language.German:               hedgehogNameGerman,
		language.Italian:              hedgehogNameItalian,
		language.Japanese:             hedgehogNameJapanese,
		language.Korean:               hedgehogNameKorean,
		language.LatinAmericanSpanish: hedgehogNameLatinAmericanSpanish,
		language.Russian:              hedgehogNameRussian,
		language.SimplifiedChinese:    hedgehogNameSimplifiedChinese,
		language.Spanish:              hedgehogNameSpanish,
		language.TraditionalChinese:   hedgehogNameTraditionalChinese,
	}
)

var (
	// Hedgehog represents the Hedgehog animal type in the Animal Crossing series.
	Hedgehog = nook.Animal{
		Key:  nook.Key(hedgehog),
		Name: hedgehogName,
	}
)
