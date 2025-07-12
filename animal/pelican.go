package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// pelican is the common reference for the Pelican animal type.
	pelican = "Pelican"
)

var (
	// pelicanNameAmericanEnglish represents the Pelican animal type's name in American English.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pelican",
	}

	// pelicanNameCanadianFrench represents the Pelican animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pelican",
	}

	// pelicanNameDutch represents the Pelican animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Pelican",
	}

	// pelicanNameFrench represents the Pelican animal type's name in French.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameFrench = nook.Name{
		Language: language.French,
		Value:    "Pelican",
	}

	// pelicanNameGerman represents the Pelican animal type's name in German.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameGerman = nook.Name{
		Language: language.German,
		Value:    "Pelican",
	}

	// pelicanNameItalian represents the Pelican animal type's name in Italian.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Pelican",
	}

	// pelicanNameJapanese represents the Pelican animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Pelican",
	}

	// pelicanNameKorean represents the Pelican animal type's name in Korean.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Pelican",
	}

	// pelicanNameLatinAmericanSpanish represents the Pelican animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pelican",
	}

	// pelicanNameRussian represents the Pelican animal type's name in Russian.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Pelican",
	}

	// pelicanNameSimplifiedChinese represents the Pelican animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Pelican",
	}

	// pelicanNameSpanish represents the Pelican animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Pelican",
	}

	// pelicanNameTraditionalChinese represents the Pelican animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Pelican"
	pelicanNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Pelican",
	}
)

var (
	// pelicanName contains the localized names of the Pelican animal type.
	pelicanName = nook.Languages{
		language.AmericanEnglish:      pelicanNameAmericanEnglish,
		language.CanadianFrench:       pelicanNameCanadianFrench,
		language.Dutch:                pelicanNameDutch,
		language.French:               pelicanNameFrench,
		language.German:               pelicanNameGerman,
		language.Italian:              pelicanNameItalian,
		language.Japanese:             pelicanNameJapanese,
		language.Korean:               pelicanNameKorean,
		language.LatinAmericanSpanish: pelicanNameLatinAmericanSpanish,
		language.Russian:              pelicanNameRussian,
		language.SimplifiedChinese:    pelicanNameSimplifiedChinese,
		language.Spanish:              pelicanNameSpanish,
		language.TraditionalChinese:   pelicanNameTraditionalChinese,
	}
)

var (
	// Pelican represents the Pelican animal type in the Animal Crossing series.
	Pelican = nook.Animal{
		Key:  nook.Key(pelican),
		Name: pelicanName,
	}
)
