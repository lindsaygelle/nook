package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// frillneckedLizard is the common reference for the FrillneckedLizard animal type.
	frillneckedLizard = "FrillneckedLizard"
)

var (
	// frillneckedLizardNameAmericanEnglish represents the FrillneckedLizard animal type's name in American English.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "FrillneckedLizard",
	}

	// frillneckedLizardNameCanadianFrench represents the FrillneckedLizard animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "FrillneckedLizard",
	}

	// frillneckedLizardNameDutch represents the FrillneckedLizard animal type's name in Dutch.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "FrillneckedLizard",
	}

	// frillneckedLizardNameFrench represents the FrillneckedLizard animal type's name in French.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameFrench = nook.Name{
		Language: language.French,
		Value:    "FrillneckedLizard",
	}

	// frillneckedLizardNameGerman represents the FrillneckedLizard animal type's name in German.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameGerman = nook.Name{
		Language: language.German,
		Value:    "FrillneckedLizard",
	}

	// frillneckedLizardNameItalian represents the FrillneckedLizard animal type's name in Italian.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "FrillneckedLizard",
	}

	// frillneckedLizardNameJapanese represents the FrillneckedLizard animal type's name in Japanese.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "FrillneckedLizard",
	}

	// frillneckedLizardNameKorean represents the FrillneckedLizard animal type's name in Korean.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "FrillneckedLizard",
	}

	// frillneckedLizardNameLatinAmericanSpanish represents the FrillneckedLizard animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "FrillneckedLizard",
	}

	// frillneckedLizardNameRussian represents the FrillneckedLizard animal type's name in Russian.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "FrillneckedLizard",
	}

	// frillneckedLizardNameSimplifiedChinese represents the FrillneckedLizard animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "FrillneckedLizard",
	}

	// frillneckedLizardNameSpanish represents the FrillneckedLizard animal type's name in Spanish.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "FrillneckedLizard",
	}

	// frillneckedLizardNameTraditionalChinese represents the FrillneckedLizard animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "FrillneckedLizard"
	frillneckedLizardNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "FrillneckedLizard",
	}
)

var (
	// frillneckedLizardName contains the localized names of the FrillneckedLizard animal type.
	frillneckedLizardName = nook.Languages{
		language.AmericanEnglish:      frillneckedLizardNameAmericanEnglish,
		language.CanadianFrench:       frillneckedLizardNameCanadianFrench,
		language.Dutch:                frillneckedLizardNameDutch,
		language.French:               frillneckedLizardNameFrench,
		language.German:               frillneckedLizardNameGerman,
		language.Italian:              frillneckedLizardNameItalian,
		language.Japanese:             frillneckedLizardNameJapanese,
		language.Korean:               frillneckedLizardNameKorean,
		language.LatinAmericanSpanish: frillneckedLizardNameLatinAmericanSpanish,
		language.Russian:              frillneckedLizardNameRussian,
		language.SimplifiedChinese:    frillneckedLizardNameSimplifiedChinese,
		language.Spanish:              frillneckedLizardNameSpanish,
		language.TraditionalChinese:   frillneckedLizardNameTraditionalChinese,
	}
)

var (
	// FrillneckedLizard represents the FrillneckedLizard animal type in the Animal Crossing series.
	FrillneckedLizard = nook.Animal{
		Key:  nook.Key(frillneckedLizard),
		Name: frillneckedLizardName,
	}
)
