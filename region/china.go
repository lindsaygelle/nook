package region

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// china is the common reference for the China region.
	china = "China"
)

var (
	// chinaNameAmericanEnglish represents the China region's name in American English.
	chinaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "China",
	}

	// chinaNameCanadianFrench represents the China region's name in Canadian French.
	chinaNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chine",
	}

	// chinaNameDutch represents the China region's name in Dutch.
	chinaNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "China",
	}

	// chinaNameFrench represents the China region's name in French.
	chinaNameFrench = nook.Name{
		Language: language.French,
		Value:    "Chine",
	}

	// chinaNameGerman represents the China region's name in German.
	chinaNameGerman = nook.Name{
		Language: language.German,
		Value:    "China",
	}

	// chinaNameItalian represents the China region's name in Italian.
	chinaNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Cina",
	}

	// chinaNameJapanese represents the China region's name in Japanese.
	chinaNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "中国",
	}

	// chinaNameKorean represents the China region's name in Korean.
	chinaNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "중국",
	}

	// chinaNameLatinAmericanSpanish represents the China region's name in Latin American Spanish.
	chinaNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "China",
	}

	// chinaNameRussian represents the China region's name in Russian.
	chinaNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Китай",
	}

	// chinaNameSimplifiedChinese represents the China region's name in Simplified Chinese.
	chinaNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "中国",
	}

	// chinaNameSpanish represents the China region's name in Spanish.
	chinaNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "China",
	}

	// chinaNameTraditionalChinese represents the China region's name in Traditional Chinese.
	chinaNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "中國",
	}
)

var (
	// chinaName contains the localized names of the China region.
	chinaName = nook.Languages{
		language.AmericanEnglish:      chinaNameAmericanEnglish,
		language.CanadianFrench:       chinaNameCanadianFrench,
		language.Dutch:                chinaNameDutch,
		language.French:               chinaNameFrench,
		language.German:               chinaNameGerman,
		language.Italian:              chinaNameItalian,
		language.Japanese:             chinaNameJapanese,
		language.Korean:               chinaNameKorean,
		language.LatinAmericanSpanish: chinaNameLatinAmericanSpanish,
		language.Russian:              chinaNameRussian,
		language.SimplifiedChinese:    chinaNameSimplifiedChinese,
		language.Spanish:              chinaNameSpanish,
		language.TraditionalChinese:   chinaNameTraditionalChinese,
	}
)

var (
	// China represents the China region in the Animal Crossing series.
	China = nook.Region{
		Key:  nook.Key(china),
		Name: chinaName,
	}
)
