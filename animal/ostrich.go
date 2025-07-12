package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// ostrich is the common reference for the Ostrich animal type.
	ostrich = "Ostrich"
)

var (
	// ostrichNameAmericanEnglish represents the Ostrich animal type's name in American English.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ostrich",
	}

	// ostrichNameCanadianFrench represents the Ostrich animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ostrich",
	}

	// ostrichNameDutch represents the Ostrich animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Ostrich",
	}

	// ostrichNameFrench represents the Ostrich animal type's name in French.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameFrench = nook.Name{
		Language: language.French,
		Value:    "Ostrich",
	}

	// ostrichNameGerman represents the Ostrich animal type's name in German.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameGerman = nook.Name{
		Language: language.German,
		Value:    "Ostrich",
	}

	// ostrichNameItalian represents the Ostrich animal type's name in Italian.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Ostrich",
	}

	// ostrichNameJapanese represents the Ostrich animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Ostrich",
	}

	// ostrichNameKorean represents the Ostrich animal type's name in Korean.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Ostrich",
	}

	// ostrichNameLatinAmericanSpanish represents the Ostrich animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ostrich",
	}

	// ostrichNameRussian represents the Ostrich animal type's name in Russian.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Ostrich",
	}

	// ostrichNameSimplifiedChinese represents the Ostrich animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Ostrich",
	}

	// ostrichNameSpanish represents the Ostrich animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Ostrich",
	}

	// ostrichNameTraditionalChinese represents the Ostrich animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Ostrich"
	ostrichNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Ostrich",
	}
)

var (
	// ostrichName contains the localized names of the Ostrich animal type.
	ostrichName = nook.Languages{
		language.AmericanEnglish:      ostrichNameAmericanEnglish,
		language.CanadianFrench:       ostrichNameCanadianFrench,
		language.Dutch:                ostrichNameDutch,
		language.French:               ostrichNameFrench,
		language.German:               ostrichNameGerman,
		language.Italian:              ostrichNameItalian,
		language.Japanese:             ostrichNameJapanese,
		language.Korean:               ostrichNameKorean,
		language.LatinAmericanSpanish: ostrichNameLatinAmericanSpanish,
		language.Russian:              ostrichNameRussian,
		language.SimplifiedChinese:    ostrichNameSimplifiedChinese,
		language.Spanish:              ostrichNameSpanish,
		language.TraditionalChinese:   ostrichNameTraditionalChinese,
	}
)

var (
	// Ostrich represents the Ostrich animal type in the Animal Crossing series.
	Ostrich = nook.Animal{
		Key:  nook.Key(ostrich),
		Name: ostrichName,
	}
)
