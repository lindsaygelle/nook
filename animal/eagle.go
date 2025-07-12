package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// eagle is the common reference for the Eagle animal type.
	eagle = "Eagle"
)

var (
	// eagleNameAmericanEnglish represents the Eagle animal type's name in American English.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Eagle",
	}

	// eagleNameCanadianFrench represents the Eagle animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Eagle",
	}

	// eagleNameDutch represents the Eagle animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Eagle",
	}

	// eagleNameFrench represents the Eagle animal type's name in French.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameFrench = nook.Name{
		Language: language.French,
		Value:    "Eagle",
	}

	// eagleNameGerman represents the Eagle animal type's name in German.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameGerman = nook.Name{
		Language: language.German,
		Value:    "Eagle",
	}

	// eagleNameItalian represents the Eagle animal type's name in Italian.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Eagle",
	}

	// eagleNameJapanese represents the Eagle animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Eagle",
	}

	// eagleNameKorean represents the Eagle animal type's name in Korean.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Eagle",
	}

	// eagleNameLatinAmericanSpanish represents the Eagle animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eagle",
	}

	// eagleNameRussian represents the Eagle animal type's name in Russian.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Eagle",
	}

	// eagleNameSimplifiedChinese represents the Eagle animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Eagle",
	}

	// eagleNameSpanish represents the Eagle animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Eagle",
	}

	// eagleNameTraditionalChinese represents the Eagle animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Eagle"
	eagleNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Eagle",
	}
)

var (
	// eagleName contains the localized names of the Eagle animal type.
	eagleName = nook.Languages{
		language.AmericanEnglish:      eagleNameAmericanEnglish,
		language.CanadianFrench:       eagleNameCanadianFrench,
		language.Dutch:                eagleNameDutch,
		language.French:               eagleNameFrench,
		language.German:               eagleNameGerman,
		language.Italian:              eagleNameItalian,
		language.Japanese:             eagleNameJapanese,
		language.Korean:               eagleNameKorean,
		language.LatinAmericanSpanish: eagleNameLatinAmericanSpanish,
		language.Russian:              eagleNameRussian,
		language.SimplifiedChinese:    eagleNameSimplifiedChinese,
		language.Spanish:              eagleNameSpanish,
		language.TraditionalChinese:   eagleNameTraditionalChinese,
	}
)

var (
	// Eagle represents the Eagle animal type in the Animal Crossing series.
	Eagle = nook.Animal{
		Key:  nook.Key(eagle),
		Name: eagleName,
	}
)
