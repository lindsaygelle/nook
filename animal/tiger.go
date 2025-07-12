package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// tiger is the common reference for the Tiger animal type.
	tiger = "Tiger"
)

var (
	// tigerNameAmericanEnglish represents the Tiger animal type's name in American English.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tiger",
	}

	// tigerNameCanadianFrench represents the Tiger animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tiger",
	}

	// tigerNameDutch represents the Tiger animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Tiger",
	}

	// tigerNameFrench represents the Tiger animal type's name in French.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameFrench = nook.Name{
		Language: language.French,
		Value:    "Tiger",
	}

	// tigerNameGerman represents the Tiger animal type's name in German.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameGerman = nook.Name{
		Language: language.German,
		Value:    "Tiger",
	}

	// tigerNameItalian represents the Tiger animal type's name in Italian.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Tiger",
	}

	// tigerNameJapanese represents the Tiger animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Tiger",
	}

	// tigerNameKorean represents the Tiger animal type's name in Korean.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Tiger",
	}

	// tigerNameLatinAmericanSpanish represents the Tiger animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tiger",
	}

	// tigerNameRussian represents the Tiger animal type's name in Russian.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Tiger",
	}

	// tigerNameSimplifiedChinese represents the Tiger animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Tiger",
	}

	// tigerNameSpanish represents the Tiger animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Tiger",
	}

	// tigerNameTraditionalChinese represents the Tiger animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Tiger"
	tigerNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Tiger",
	}
)

var (
	// tigerName contains the localized names of the Tiger animal type.
	tigerName = nook.Languages{
		language.AmericanEnglish:      tigerNameAmericanEnglish,
		language.CanadianFrench:       tigerNameCanadianFrench,
		language.Dutch:                tigerNameDutch,
		language.French:               tigerNameFrench,
		language.German:               tigerNameGerman,
		language.Italian:              tigerNameItalian,
		language.Japanese:             tigerNameJapanese,
		language.Korean:               tigerNameKorean,
		language.LatinAmericanSpanish: tigerNameLatinAmericanSpanish,
		language.Russian:              tigerNameRussian,
		language.SimplifiedChinese:    tigerNameSimplifiedChinese,
		language.Spanish:              tigerNameSpanish,
		language.TraditionalChinese:   tigerNameTraditionalChinese,
	}
)

var (
	// Tiger represents the Tiger animal type in the Animal Crossing series.
	Tiger = nook.Animal{
		Key:  nook.Key(tiger),
		Name: tigerName,
	}
)
