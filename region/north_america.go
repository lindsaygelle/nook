package region

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// northAmerica is the common reference for the North America region.
	northAmerica = "NorthAmerica"
)

var (
	// northAmericaNameAmericanEnglish represents the North America region's name in American English.
	northAmericaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "North America",
	}

	// northAmericaNameCanadianFrench represents the North America region's name in Canadian French.
	northAmericaNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Amérique du Nord",
	}

	// northAmericaNameDutch represents the North America region's name in Dutch.
	northAmericaNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Noord-Amerika",
	}

	// northAmericaNameFrench represents the North America region's name in French.
	northAmericaNameFrench = nook.Name{
		Language: language.French,
		Value:    "Amérique du Nord",
	}

	// northAmericaNameGerman represents the North America region's name in German.
	northAmericaNameGerman = nook.Name{
		Language: language.German,
		Value:    "Nordamerika",
	}

	// northAmericaNameItalian represents the North America region's name in Italian.
	northAmericaNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Nord America",
	}

	// northAmericaNameJapanese represents the North America region's name in Japanese.
	northAmericaNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "北アメリカ大陸",
	}

	// northAmericaNameKorean represents the North America region's name in Korean.
	northAmericaNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "북아메리카",
	}

	// northAmericaNameLatinAmericanSpanish represents the North America region's name in Latin American Spanish.
	northAmericaNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "América del Norte",
	}

	// northAmericaNameRussian represents the North America region's name in Russian.
	northAmericaNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Северная Америка",
	}

	// northAmericaNameSimplifiedChinese represents the North America region's name in Simplified Chinese.
	northAmericaNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "北美洲",
	}

	// northAmericaNameSpanish represents the North America region's name in Spanish.
	northAmericaNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "América del Norte",
	}

	// northAmericaNameTraditionalChinese represents the North America region's name in Traditional Chinese.
	northAmericaNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "北美洲",
	}
)

var (
	// northAmericaName contains the localized names of the North America region.
	northAmericaName = nook.Languages{
		language.AmericanEnglish:      northAmericaNameAmericanEnglish,
		language.CanadianFrench:       northAmericaNameCanadianFrench,
		language.Dutch:                northAmericaNameDutch,
		language.French:               northAmericaNameFrench,
		language.German:               northAmericaNameGerman,
		language.Italian:              northAmericaNameItalian,
		language.Japanese:             northAmericaNameJapanese,
		language.Korean:               northAmericaNameKorean,
		language.LatinAmericanSpanish: northAmericaNameLatinAmericanSpanish,
		language.Russian:              northAmericaNameRussian,
		language.SimplifiedChinese:    northAmericaNameSimplifiedChinese,
		language.Spanish:              northAmericaNameSpanish,
		language.TraditionalChinese:   northAmericaNameTraditionalChinese,
	}
)

var (
	// NorthAmerica represents the North America region in the Animal Crossing series.
	NorthAmerica = nook.Region{
		Key:  nook.Key(northAmerica),
		Name: northAmericaName,
	}
)
