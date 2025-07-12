package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// furSeal is the common reference for the Furseal animal type.
	furSeal = "Furseal"
)

var (
	// furSealNameAmericanEnglish represents the Furseal animal type's name in American English.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Furseal",
	}

	// furSealNameCanadianFrench represents the Furseal animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Furseal",
	}

	// furSealNameDutch represents the Furseal animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Furseal",
	}

	// furSealNameFrench represents the Furseal animal type's name in French.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameFrench = nook.Name{
		Language: language.French,
		Value:    "Furseal",
	}

	// furSealNameGerman represents the Furseal animal type's name in German.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameGerman = nook.Name{
		Language: language.German,
		Value:    "Furseal",
	}

	// furSealNameItalian represents the Furseal animal type's name in Italian.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Furseal",
	}

	// furSealNameJapanese represents the Furseal animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Furseal",
	}

	// furSealNameKorean represents the Furseal animal type's name in Korean.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Furseal",
	}

	// furSealNameLatinAmericanSpanish represents the Furseal animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Furseal",
	}

	// furSealNameRussian represents the Furseal animal type's name in Russian.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Furseal",
	}

	// furSealNameSimplifiedChinese represents the Furseal animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Furseal",
	}

	// furSealNameSpanish represents the Furseal animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Furseal",
	}

	// furSealNameTraditionalChinese represents the Furseal animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Furseal"
	furSealNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Furseal",
	}
)

var (
	// furSealName contains the localized names of the Furseal animal type.
	furSealName = nook.Languages{
		language.AmericanEnglish:      furSealNameAmericanEnglish,
		language.CanadianFrench:       furSealNameCanadianFrench,
		language.Dutch:                furSealNameDutch,
		language.French:               furSealNameFrench,
		language.German:               furSealNameGerman,
		language.Italian:              furSealNameItalian,
		language.Japanese:             furSealNameJapanese,
		language.Korean:               furSealNameKorean,
		language.LatinAmericanSpanish: furSealNameLatinAmericanSpanish,
		language.Russian:              furSealNameRussian,
		language.SimplifiedChinese:    furSealNameSimplifiedChinese,
		language.Spanish:              furSealNameSpanish,
		language.TraditionalChinese:   furSealNameTraditionalChinese,
	}
)

var (
	// FurSeal represents the Furseal animal type in the Animal Crossing series.
	FurSeal = nook.Animal{
		Key:  nook.Key(furSeal),
		Name: furSealName,
	}
)
