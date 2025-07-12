package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// axolotl is the common reference for the Axolotl animal type.
	axolotl = "Axolotl"
)

var (
	// axolotlNameAmericanEnglish represents the Axolotl animal type's name in American English.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Axolotl",
	}

	// axolotlNameCanadianFrench represents the Axolotl animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Axolotl",
	}

	// axolotlNameDutch represents the Axolotl animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Axolotl",
	}

	// axolotlNameFrench represents the Axolotl animal type's name in French.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameFrench = nook.Name{
		Language: language.French,
		Value:    "Axolotl",
	}

	// axolotlNameGerman represents the Axolotl animal type's name in German.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameGerman = nook.Name{
		Language: language.German,
		Value:    "Axolotl",
	}

	// axolotlNameItalian represents the Axolotl animal type's name in Italian.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Axolotl",
	}

	// axolotlNameJapanese represents the Axolotl animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Axolotl",
	}

	// axolotlNameKorean represents the Axolotl animal type's name in Korean.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Axolotl",
	}

	// axolotlNameLatinAmericanSpanish represents the Axolotl animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Axolotl",
	}

	// axolotlNameRussian represents the Axolotl animal type's name in Russian.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Axolotl",
	}

	// axolotlNameSimplifiedChinese represents the Axolotl animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Axolotl",
	}

	// axolotlNameSpanish represents the Axolotl animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Axolotl",
	}

	// axolotlNameTraditionalChinese represents the Axolotl animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Axolotl"
	axolotlNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Axolotl",
	}
)

var (
	// axolotlName contains the localized names of the Axolotl animal type.
	axolotlName = nook.Languages{
		language.AmericanEnglish:      axolotlNameAmericanEnglish,
		language.CanadianFrench:       axolotlNameCanadianFrench,
		language.Dutch:                axolotlNameDutch,
		language.French:               axolotlNameFrench,
		language.German:               axolotlNameGerman,
		language.Italian:              axolotlNameItalian,
		language.Japanese:             axolotlNameJapanese,
		language.Korean:               axolotlNameKorean,
		language.LatinAmericanSpanish: axolotlNameLatinAmericanSpanish,
		language.Russian:              axolotlNameRussian,
		language.SimplifiedChinese:    axolotlNameSimplifiedChinese,
		language.Spanish:              axolotlNameSpanish,
		language.TraditionalChinese:   axolotlNameTraditionalChinese,
	}
)

var (
	// Axolotl represents the Axolotl animal type in the Animal Crossing series.
	Axolotl = nook.Animal{
		Key:  nook.Key(axolotl),
		Name: axolotlName,
	}
)
