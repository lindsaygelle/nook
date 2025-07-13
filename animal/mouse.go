package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// mouse is the common reference for the Mouse animal type.
	mouse = "Mouse"
)

var (
	// mouseNameAmericanEnglish represents the Mouse animal type's name in American English.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mouse",
	}

	// mouseNameCanadianFrench represents the Mouse animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mouse",
	}

	// mouseNameDutch represents the Mouse animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Mouse",
	}

	// mouseNameFrench represents the Mouse animal type's name in French.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameFrench = nook.Name{
		Language: language.French,
		Value:    "Mouse",
	}

	// mouseNameGerman represents the Mouse animal type's name in German.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameGerman = nook.Name{
		Language: language.German,
		Value:    "Mouse",
	}

	// mouseNameItalian represents the Mouse animal type's name in Italian.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Mouse",
	}

	// mouseNameJapanese represents the Mouse animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Mouse",
	}

	// mouseNameKorean represents the Mouse animal type's name in Korean.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Mouse",
	}

	// mouseNameLatinAmericanSpanish represents the Mouse animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mouse",
	}

	// mouseNameRussian represents the Mouse animal type's name in Russian.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Mouse",
	}

	// mouseNameSimplifiedChinese represents the Mouse animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Mouse",
	}

	// mouseNameSpanish represents the Mouse animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Mouse",
	}

	// mouseNameTraditionalChinese represents the Mouse animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Mouse"
	mouseNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Mouse",
	}
)

var (
	// mouseName contains the localized names of the Mouse animal type.
	mouseName = nook.Languages{
		language.AmericanEnglish:      mouseNameAmericanEnglish,
		language.CanadianFrench:       mouseNameCanadianFrench,
		language.Dutch:                mouseNameDutch,
		language.French:               mouseNameFrench,
		language.German:               mouseNameGerman,
		language.Italian:              mouseNameItalian,
		language.Japanese:             mouseNameJapanese,
		language.Korean:               mouseNameKorean,
		language.LatinAmericanSpanish: mouseNameLatinAmericanSpanish,
		language.Russian:              mouseNameRussian,
		language.SimplifiedChinese:    mouseNameSimplifiedChinese,
		language.Spanish:              mouseNameSpanish,
		language.TraditionalChinese:   mouseNameTraditionalChinese,
	}
)

var (
	// Mouse represents the Mouse animal type in the Animal Crossing series.
	Mouse = nook.Animal{
		Key:  nook.Key(mouse),
		Name: mouseName,
	}
)
