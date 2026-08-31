package region

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// worldwide is the common reference for the Worldwide region.
	worldwide = "Worldwide"
)

var (
	// worldwideNameAmericanEnglish represents the Worldwide region's name in American English.
	worldwideNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Worldwide",
	}

	// worldwideNameCanadianFrench represents the Worldwide region's name in Canadian French.
	worldwideNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Monde",
	}

	// worldwideNameDutch represents the Worldwide region's name in Dutch.
	worldwideNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Wereld",
	}

	// worldwideNameFrench represents the Worldwide region's name in French.
	worldwideNameFrench = nook.Name{
		Language: language.French,
		Value:    "Monde",
	}

	// worldwideNameGerman represents the Worldwide region's name in German.
	worldwideNameGerman = nook.Name{
		Language: language.German,
		Value:    "Welt",
	}

	// worldwideNameItalian represents the Worldwide region's name in Italian.
	worldwideNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Mondo",
	}

	// worldwideNameJapanese represents the Worldwide region's name in Japanese.
	worldwideNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "世界",
	}

	// worldwideNameKorean represents the Worldwide region's name in Korean.
	worldwideNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "세계",
	}

	// worldwideNameLatinAmericanSpanish represents the Worldwide region's name in Latin American Spanish.
	worldwideNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mundo",
	}

	// worldwideNameRussian represents the Worldwide region's name in Russian.
	worldwideNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Весь мир",
	}

	// worldwideNameSimplifiedChinese represents the Worldwide region's name in Simplified Chinese.
	worldwideNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "世界",
	}

	// worldwideNameSpanish represents the Worldwide region's name in Spanish.
	worldwideNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Mundo",
	}

	// worldwideNameTraditionalChinese represents the Worldwide region's name in Traditional Chinese.
	worldwideNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "世界",
	}
)

var (
	// worldwideName contains the localized names of the Worldwide region.
	worldwideName = nook.Languages{
		language.AmericanEnglish:      worldwideNameAmericanEnglish,
		language.CanadianFrench:       worldwideNameCanadianFrench,
		language.Dutch:                worldwideNameDutch,
		language.French:               worldwideNameFrench,
		language.German:               worldwideNameGerman,
		language.Italian:              worldwideNameItalian,
		language.Japanese:             worldwideNameJapanese,
		language.Korean:               worldwideNameKorean,
		language.LatinAmericanSpanish: worldwideNameLatinAmericanSpanish,
		language.Russian:              worldwideNameRussian,
		language.SimplifiedChinese:    worldwideNameSimplifiedChinese,
		language.Spanish:              worldwideNameSpanish,
		language.TraditionalChinese:   worldwideNameTraditionalChinese,
	}
)

var (
	// Worldwide represents the Worldwide region in the Animal Crossing series.
	Worldwide = nook.Region{
		Key:  nook.Key(worldwide),
		Name: worldwideName,
	}
)
