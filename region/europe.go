package region

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// europe is the common reference for the Europe region.
	europe = "Europe"
)

var (
	// europeNameAmericanEnglish represents the Europe region's name in American English.
	europeNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Europe",
	}

	// europeNameCanadianFrench represents the Europe region's name in Canadian French.
	europeNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Europe",
	}

	// europeNameDutch represents the Europe region's name in Dutch.
	europeNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Europa",
	}

	// europeNameFrench represents the Europe region's name in French.
	europeNameFrench = nook.Name{
		Language: language.French,
		Value:    "Europe",
	}

	// europeNameGerman represents the Europe region's name in German.
	europeNameGerman = nook.Name{
		Language: language.German,
		Value:    "Europa",
	}

	// europeNameItalian represents the Europe region's name in Italian.
	europeNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Europa",
	}

	// europeNameJapanese represents the Europe region's name in Japanese.
	europeNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "ヨーロッパ",
	}

	// europeNameKorean represents the Europe region's name in Korean.
	europeNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "유럽",
	}

	// europeNameLatinAmericanSpanish represents the Europe region's name in Latin American Spanish.
	europeNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Europa",
	}

	// europeNameRussian represents the Europe region's name in Russian.
	europeNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Европа",
	}

	// europeNameSimplifiedChinese represents the Europe region's name in Simplified Chinese.
	europeNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "欧洲",
	}

	// europeNameSpanish represents the Europe region's name in Spanish.
	europeNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Europa",
	}

	// europeNameTraditionalChinese represents the Europe region's name in Traditional Chinese.
	europeNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "歐洲",
	}
)

var (
	// europeName contains the localized names of the Europe region.
	europeName = nook.Languages{
		language.AmericanEnglish:      europeNameAmericanEnglish,
		language.CanadianFrench:       europeNameCanadianFrench,
		language.Dutch:                europeNameDutch,
		language.French:               europeNameFrench,
		language.German:               europeNameGerman,
		language.Italian:              europeNameItalian,
		language.Japanese:             europeNameJapanese,
		language.Korean:               europeNameKorean,
		language.LatinAmericanSpanish: europeNameLatinAmericanSpanish,
		language.Russian:              europeNameRussian,
		language.SimplifiedChinese:    europeNameSimplifiedChinese,
		language.Spanish:              europeNameSpanish,
		language.TraditionalChinese:   europeNameTraditionalChinese,
	}
)

var (
	// Europe represents the Europe region in the Animal Crossing series.
	Europe = nook.Region{
		Key:  nook.Key(europe),
		Name: europeName,
	}
)
