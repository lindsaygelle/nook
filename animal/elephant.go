package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// elephant is the common reference for the Elephant animal type.
	elephant = "Elephant"
)

var (
	// elephantNameAmericanEnglish represents the Elephant animal type's name in American English.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Elephant",
	}

	// elephantNameCanadianFrench represents the Elephant animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Elephant",
	}

	// elephantNameDutch represents the Elephant animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Elephant",
	}

	// elephantNameFrench represents the Elephant animal type's name in French.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameFrench = nook.Name{
		Language: language.French,
		Value:    "Elephant",
	}

	// elephantNameGerman represents the Elephant animal type's name in German.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameGerman = nook.Name{
		Language: language.German,
		Value:    "Elephant",
	}

	// elephantNameItalian represents the Elephant animal type's name in Italian.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Elephant",
	}

	// elephantNameJapanese represents the Elephant animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Elephant",
	}

	// elephantNameKorean represents the Elephant animal type's name in Korean.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Elephant",
	}

	// elephantNameLatinAmericanSpanish represents the Elephant animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Elephant",
	}

	// elephantNameRussian represents the Elephant animal type's name in Russian.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Elephant",
	}

	// elephantNameSimplifiedChinese represents the Elephant animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Elephant",
	}

	// elephantNameSpanish represents the Elephant animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Elephant",
	}

	// elephantNameTraditionalChinese represents the Elephant animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Elephant"
	elephantNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Elephant",
	}
)

var (
	// elephantName contains the localized names of the Elephant animal type.
	elephantName = nook.Languages{
		language.AmericanEnglish:      elephantNameAmericanEnglish,
		language.CanadianFrench:       elephantNameCanadianFrench,
		language.Dutch:                elephantNameDutch,
		language.French:               elephantNameFrench,
		language.German:               elephantNameGerman,
		language.Italian:              elephantNameItalian,
		language.Japanese:             elephantNameJapanese,
		language.Korean:               elephantNameKorean,
		language.LatinAmericanSpanish: elephantNameLatinAmericanSpanish,
		language.Russian:              elephantNameRussian,
		language.SimplifiedChinese:    elephantNameSimplifiedChinese,
		language.Spanish:              elephantNameSpanish,
		language.TraditionalChinese:   elephantNameTraditionalChinese,
	}
)

var (
	// Elephant represents the Elephant animal type in the Animal Crossing series.
	Elephant = nook.Animal{
		Key:  nook.Key(elephant),
		Name: elephantName,
	}
)
