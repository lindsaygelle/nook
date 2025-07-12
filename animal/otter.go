package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// otter is the common reference for the Otter animal type.
	otter = "Otter"
)

var (
	// otterNameAmericanEnglish represents the Otter animal type's name in American English.
	// TODO: Verify translation accuracy for "Otter"
	otterNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Otter",
	}

	// otterNameCanadianFrench represents the Otter animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Otter"
	otterNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Otter",
	}

	// otterNameDutch represents the Otter animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Otter"
	otterNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Otter",
	}

	// otterNameFrench represents the Otter animal type's name in French.
	// TODO: Verify translation accuracy for "Otter"
	otterNameFrench = nook.Name{
		Language: language.French,
		Value:    "Otter",
	}

	// otterNameGerman represents the Otter animal type's name in German.
	// TODO: Verify translation accuracy for "Otter"
	otterNameGerman = nook.Name{
		Language: language.German,
		Value:    "Otter",
	}

	// otterNameItalian represents the Otter animal type's name in Italian.
	// TODO: Verify translation accuracy for "Otter"
	otterNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Otter",
	}

	// otterNameJapanese represents the Otter animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Otter"
	otterNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Otter",
	}

	// otterNameKorean represents the Otter animal type's name in Korean.
	// TODO: Verify translation accuracy for "Otter"
	otterNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Otter",
	}

	// otterNameLatinAmericanSpanish represents the Otter animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Otter"
	otterNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Otter",
	}

	// otterNameRussian represents the Otter animal type's name in Russian.
	// TODO: Verify translation accuracy for "Otter"
	otterNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Otter",
	}

	// otterNameSimplifiedChinese represents the Otter animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Otter"
	otterNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Otter",
	}

	// otterNameSpanish represents the Otter animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Otter"
	otterNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Otter",
	}

	// otterNameTraditionalChinese represents the Otter animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Otter"
	otterNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Otter",
	}
)

var (
	// otterName contains the localized names of the Otter animal type.
	otterName = nook.Languages{
		language.AmericanEnglish:      otterNameAmericanEnglish,
		language.CanadianFrench:       otterNameCanadianFrench,
		language.Dutch:                otterNameDutch,
		language.French:               otterNameFrench,
		language.German:               otterNameGerman,
		language.Italian:              otterNameItalian,
		language.Japanese:             otterNameJapanese,
		language.Korean:               otterNameKorean,
		language.LatinAmericanSpanish: otterNameLatinAmericanSpanish,
		language.Russian:              otterNameRussian,
		language.SimplifiedChinese:    otterNameSimplifiedChinese,
		language.Spanish:              otterNameSpanish,
		language.TraditionalChinese:   otterNameTraditionalChinese,
	}
)

var (
	// Otter represents the Otter animal type in the Animal Crossing series.
	Otter = nook.Animal{
		Key:  nook.Key(otter),
		Name: otterName,
	}
)
