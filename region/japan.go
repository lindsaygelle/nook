package region

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// japan is the common reference for the Japan region.
	japan = "Japan"
)

var (
	// japanNameAmericanEnglish represents the Japan region's name in American English.
	japanNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Japan",
	}

	// japanNameCanadianFrench represents the Japan region's name in Canadian French.
	japanNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Japon",
	}

	// japanNameDutch represents the Japan region's name in Dutch.
	japanNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Japan",
	}

	// japanNameFrench represents the Japan region's name in French.
	japanNameFrench = nook.Name{
		Language: language.French,
		Value:    "Japon",
	}

	// japanNameGerman represents the Japan region's name in German.
	japanNameGerman = nook.Name{
		Language: language.German,
		Value:    "Japan",
	}

	// japanNameItalian represents the Japan region's name in Italian.
	japanNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Giappone",
	}

	// japanNameJapanese represents the Japan region's name in Japanese.
	japanNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "日本",
	}

	// japanNameKorean represents the Japan region's name in Korean.
	japanNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "일본",
	}

	// japanNameLatinAmericanSpanish represents the Japan region's name in Latin American Spanish.
	japanNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Japón",
	}

	// japanNameRussian represents the Japan region's name in Russian.
	japanNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Япония",
	}

	// japanNameSimplifiedChinese represents the Japan region's name in Simplified Chinese.
	japanNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "日本",
	}

	// japanNameSpanish represents the Japan region's name in Spanish.
	japanNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Japón",
	}

	// japanNameTraditionalChinese represents the Japan region's name in Traditional Chinese.
	japanNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "日本",
	}
)

var (
	// japanName contains the localized names of the Japan region.
	japanName = nook.Languages{
		language.AmericanEnglish:      japanNameAmericanEnglish,
		language.CanadianFrench:       japanNameCanadianFrench,
		language.Dutch:                japanNameDutch,
		language.French:               japanNameFrench,
		language.German:               japanNameGerman,
		language.Italian:              japanNameItalian,
		language.Japanese:             japanNameJapanese,
		language.Korean:               japanNameKorean,
		language.LatinAmericanSpanish: japanNameLatinAmericanSpanish,
		language.Russian:              japanNameRussian,
		language.SimplifiedChinese:    japanNameSimplifiedChinese,
		language.Spanish:              japanNameSpanish,
		language.TraditionalChinese:   japanNameTraditionalChinese,
	}
)

var (
	// Japan represents the Japan region in the Animal Crossing series.
	Japan = nook.Region{
		Key:  nook.Key(japan),
		Name: japanName,
	}
)
