package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// pig is the common reference for the Pig animal type.
	pig = "Pig"
)

var (
	// pigNameAmericanEnglish represents the Pig animal type's name in American English.
	// TODO: Verify translation accuracy for "Pig"
	pigNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pig",
	}

	// pigNameCanadianFrench represents the Pig animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Pig"
	pigNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pig",
	}

	// pigNameDutch represents the Pig animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Pig"
	pigNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Pig",
	}

	// pigNameFrench represents the Pig animal type's name in French.
	// TODO: Verify translation accuracy for "Pig"
	pigNameFrench = nook.Name{
		Language: language.French,
		Value:    "Pig",
	}

	// pigNameGerman represents the Pig animal type's name in German.
	// TODO: Verify translation accuracy for "Pig"
	pigNameGerman = nook.Name{
		Language: language.German,
		Value:    "Pig",
	}

	// pigNameItalian represents the Pig animal type's name in Italian.
	// TODO: Verify translation accuracy for "Pig"
	pigNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Pig",
	}

	// pigNameJapanese represents the Pig animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Pig"
	pigNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Pig",
	}

	// pigNameKorean represents the Pig animal type's name in Korean.
	// TODO: Verify translation accuracy for "Pig"
	pigNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Pig",
	}

	// pigNameLatinAmericanSpanish represents the Pig animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Pig"
	pigNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pig",
	}

	// pigNameRussian represents the Pig animal type's name in Russian.
	// TODO: Verify translation accuracy for "Pig"
	pigNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Pig",
	}

	// pigNameSimplifiedChinese represents the Pig animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Pig"
	pigNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Pig",
	}

	// pigNameSpanish represents the Pig animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Pig"
	pigNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Pig",
	}

	// pigNameTraditionalChinese represents the Pig animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Pig"
	pigNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Pig",
	}
)

var (
	// pigName contains the localized names of the Pig animal type.
	pigName = nook.Languages{
		language.AmericanEnglish:      pigNameAmericanEnglish,
		language.CanadianFrench:       pigNameCanadianFrench,
		language.Dutch:                pigNameDutch,
		language.French:               pigNameFrench,
		language.German:               pigNameGerman,
		language.Italian:              pigNameItalian,
		language.Japanese:             pigNameJapanese,
		language.Korean:               pigNameKorean,
		language.LatinAmericanSpanish: pigNameLatinAmericanSpanish,
		language.Russian:              pigNameRussian,
		language.SimplifiedChinese:    pigNameSimplifiedChinese,
		language.Spanish:              pigNameSpanish,
		language.TraditionalChinese:   pigNameTraditionalChinese,
	}
)

var (
	// Pig represents the Pig animal type in the Animal Crossing series.
	Pig = nook.Animal{
		Key:  nook.Key(pig),
		Name: pigName,
	}
)
