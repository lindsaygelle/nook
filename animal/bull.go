package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// bull is the common reference for the Bull animal type.
	bull = "Bull"
)

var (
	// bullNameAmericanEnglish represents the Bull animal type's name in American English.
	// TODO: Verify translation accuracy for "Bull"
	bullNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bull",
	}

	// bullNameCanadianFrench represents the Bull animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Bull"
	bullNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bull",
	}

	// bullNameDutch represents the Bull animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Bull"
	bullNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Bull",
	}

	// bullNameFrench represents the Bull animal type's name in French.
	// TODO: Verify translation accuracy for "Bull"
	bullNameFrench = nook.Name{
		Language: language.French,
		Value:    "Bull",
	}

	// bullNameGerman represents the Bull animal type's name in German.
	// TODO: Verify translation accuracy for "Bull"
	bullNameGerman = nook.Name{
		Language: language.German,
		Value:    "Bull",
	}

	// bullNameItalian represents the Bull animal type's name in Italian.
	// TODO: Verify translation accuracy for "Bull"
	bullNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Bull",
	}

	// bullNameJapanese represents the Bull animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Bull"
	bullNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Bull",
	}

	// bullNameKorean represents the Bull animal type's name in Korean.
	// TODO: Verify translation accuracy for "Bull"
	bullNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Bull",
	}

	// bullNameLatinAmericanSpanish represents the Bull animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Bull"
	bullNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bull",
	}

	// bullNameRussian represents the Bull animal type's name in Russian.
	// TODO: Verify translation accuracy for "Bull"
	bullNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Bull",
	}

	// bullNameSimplifiedChinese represents the Bull animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Bull"
	bullNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Bull",
	}

	// bullNameSpanish represents the Bull animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Bull"
	bullNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Bull",
	}

	// bullNameTraditionalChinese represents the Bull animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Bull"
	bullNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Bull",
	}
)

var (
	// bullName contains the localized names of the Bull animal type.
	bullName = nook.Languages{
		language.AmericanEnglish:      bullNameAmericanEnglish,
		language.CanadianFrench:       bullNameCanadianFrench,
		language.Dutch:                bullNameDutch,
		language.French:               bullNameFrench,
		language.German:               bullNameGerman,
		language.Italian:              bullNameItalian,
		language.Japanese:             bullNameJapanese,
		language.Korean:               bullNameKorean,
		language.LatinAmericanSpanish: bullNameLatinAmericanSpanish,
		language.Russian:              bullNameRussian,
		language.SimplifiedChinese:    bullNameSimplifiedChinese,
		language.Spanish:              bullNameSpanish,
		language.TraditionalChinese:   bullNameTraditionalChinese,
	}
)

var (
	// Bull represents the Bull animal type in the Animal Crossing series.
	Bull = nook.Animal{
		Key:  nook.Key(bull),
		Name: bullName,
	}
)
