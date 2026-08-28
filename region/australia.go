package region

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// australia is the common reference for the Australia region.
	australia = "Australia"
)

var (
	// australiaNameAmericanEnglish represents the Australia region's name in American English.
	australiaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Australia",
	}

	// australiaNameCanadianFrench represents the Australia region's name in Canadian French.
	australiaNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Australie",
	}

	// australiaNameDutch represents the Australia region's name in Dutch.
	australiaNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Australië",
	}

	// australiaNameFrench represents the Australia region's name in French.
	australiaNameFrench = nook.Name{
		Language: language.French,
		Value:    "Australie",
	}

	// australiaNameGerman represents the Australia region's name in German.
	australiaNameGerman = nook.Name{
		Language: language.German,
		Value:    "Australien",
	}

	// australiaNameItalian represents the Australia region's name in Italian.
	australiaNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Australia",
	}

	// australiaNameJapanese represents the Australia region's name in Japanese.
	australiaNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "オーストラリア",
	}

	// australiaNameKorean represents the Australia region's name in Korean.
	australiaNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "오스트레일리아",
	}

	// australiaNameLatinAmericanSpanish represents the Australia region's name in Latin American Spanish.
	australiaNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Australia",
	}

	// australiaNameRussian represents the Australia region's name in Russian.
	australiaNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Австралия",
	}

	// australiaNameSimplifiedChinese represents the Australia region's name in Simplified Chinese.
	australiaNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "澳大利亚",
	}

	// australiaNameSpanish represents the Australia region's name in Spanish.
	australiaNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Australia",
	}

	// australiaNameTraditionalChinese represents the Australia region's name in Traditional Chinese.
	australiaNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "澳洲",
	}
)

var (
	// australiaName contains the localized names of the Australia region.
	australiaName = nook.Languages{
		language.AmericanEnglish:      australiaNameAmericanEnglish,
		language.CanadianFrench:       australiaNameCanadianFrench,
		language.Dutch:                australiaNameDutch,
		language.French:               australiaNameFrench,
		language.German:               australiaNameGerman,
		language.Italian:              australiaNameItalian,
		language.Japanese:             australiaNameJapanese,
		language.Korean:               australiaNameKorean,
		language.LatinAmericanSpanish: australiaNameLatinAmericanSpanish,
		language.Russian:              australiaNameRussian,
		language.SimplifiedChinese:    australiaNameSimplifiedChinese,
		language.Spanish:              australiaNameSpanish,
		language.TraditionalChinese:   australiaNameTraditionalChinese,
	}
)

var (
	// Australia represents the Australia region in the Animal Crossing series.
	Australia = nook.Region{
		Key:  nook.Key(australia),
		Name: australiaName,
	}
)
