package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// anteater is the common reference for the Anteater animal type.
	anteater = "Anteater"
)

var (
	// anteaterNameAmericanEnglish represents the Anteater animal type's name in American English.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Anteater",
	}

	// anteaterNameCanadianFrench represents the Anteater animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Anteater",
	}

	// anteaterNameDutch represents the Anteater animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Anteater",
	}

	// anteaterNameFrench represents the Anteater animal type's name in French.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameFrench = nook.Name{
		Language: language.French,
		Value:    "Anteater",
	}

	// anteaterNameGerman represents the Anteater animal type's name in German.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameGerman = nook.Name{
		Language: language.German,
		Value:    "Anteater",
	}

	// anteaterNameItalian represents the Anteater animal type's name in Italian.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Anteater",
	}

	// anteaterNameJapanese represents the Anteater animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Anteater",
	}

	// anteaterNameKorean represents the Anteater animal type's name in Korean.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Anteater",
	}

	// anteaterNameLatinAmericanSpanish represents the Anteater animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Anteater",
	}

	// anteaterNameRussian represents the Anteater animal type's name in Russian.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Anteater",
	}

	// anteaterNameSimplifiedChinese represents the Anteater animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Anteater",
	}

	// anteaterNameSpanish represents the Anteater animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Anteater",
	}

	// anteaterNameTraditionalChinese represents the Anteater animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Anteater"
	anteaterNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Anteater",
	}
)

var (
	// anteaterName contains the localized names of the Anteater animal type.
	anteaterName = nook.Languages{
		language.AmericanEnglish:      anteaterNameAmericanEnglish,
		language.CanadianFrench:       anteaterNameCanadianFrench,
		language.Dutch:                anteaterNameDutch,
		language.French:               anteaterNameFrench,
		language.German:               anteaterNameGerman,
		language.Italian:              anteaterNameItalian,
		language.Japanese:             anteaterNameJapanese,
		language.Korean:               anteaterNameKorean,
		language.LatinAmericanSpanish: anteaterNameLatinAmericanSpanish,
		language.Russian:              anteaterNameRussian,
		language.SimplifiedChinese:    anteaterNameSimplifiedChinese,
		language.Spanish:              anteaterNameSpanish,
		language.TraditionalChinese:   anteaterNameTraditionalChinese,
	}
)

var (
	// Anteater represents the Anteater animal type in the Animal Crossing series.
	Anteater = nook.Animal{
		Key:  nook.Key(anteater),
		Name: anteaterName,
	}
)
