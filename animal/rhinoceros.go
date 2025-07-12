package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// rhinoceros is the common reference for the Rhinoceros animal type.
	rhinoceros = "Rhinoceros"
)

var (
	// rhinocerosNameAmericanEnglish represents the Rhinoceros animal type's name in American English.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rhinoceros",
	}

	// rhinocerosNameCanadianFrench represents the Rhinoceros animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rhinoceros",
	}

	// rhinocerosNameDutch represents the Rhinoceros animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Rhinoceros",
	}

	// rhinocerosNameFrench represents the Rhinoceros animal type's name in French.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameFrench = nook.Name{
		Language: language.French,
		Value:    "Rhinoceros",
	}

	// rhinocerosNameGerman represents the Rhinoceros animal type's name in German.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameGerman = nook.Name{
		Language: language.German,
		Value:    "Rhinoceros",
	}

	// rhinocerosNameItalian represents the Rhinoceros animal type's name in Italian.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Rhinoceros",
	}

	// rhinocerosNameJapanese represents the Rhinoceros animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Rhinoceros",
	}

	// rhinocerosNameKorean represents the Rhinoceros animal type's name in Korean.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Rhinoceros",
	}

	// rhinocerosNameLatinAmericanSpanish represents the Rhinoceros animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rhinoceros",
	}

	// rhinocerosNameRussian represents the Rhinoceros animal type's name in Russian.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Rhinoceros",
	}

	// rhinocerosNameSimplifiedChinese represents the Rhinoceros animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Rhinoceros",
	}

	// rhinocerosNameSpanish represents the Rhinoceros animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Rhinoceros",
	}

	// rhinocerosNameTraditionalChinese represents the Rhinoceros animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Rhinoceros"
	rhinocerosNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Rhinoceros",
	}
)

var (
	// rhinocerosName contains the localized names of the Rhinoceros animal type.
	rhinocerosName = nook.Languages{
		language.AmericanEnglish:      rhinocerosNameAmericanEnglish,
		language.CanadianFrench:       rhinocerosNameCanadianFrench,
		language.Dutch:                rhinocerosNameDutch,
		language.French:               rhinocerosNameFrench,
		language.German:               rhinocerosNameGerman,
		language.Italian:              rhinocerosNameItalian,
		language.Japanese:             rhinocerosNameJapanese,
		language.Korean:               rhinocerosNameKorean,
		language.LatinAmericanSpanish: rhinocerosNameLatinAmericanSpanish,
		language.Russian:              rhinocerosNameRussian,
		language.SimplifiedChinese:    rhinocerosNameSimplifiedChinese,
		language.Spanish:              rhinocerosNameSpanish,
		language.TraditionalChinese:   rhinocerosNameTraditionalChinese,
	}
)

var (
	// Rhinoceros represents the Rhinoceros animal type in the Animal Crossing series.
	Rhinoceros = nook.Animal{
		Key:  nook.Key(rhinoceros),
		Name: rhinocerosName,
	}
)
