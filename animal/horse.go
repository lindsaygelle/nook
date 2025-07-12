package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// horse is the common reference for the Horse animal type.
	horse = "Horse"
)

var (
	// horseNameAmericanEnglish represents the Horse animal type's name in American English.
	// TODO: Verify translation accuracy for "Horse"
	horseNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Horse",
	}

	// horseNameCanadianFrench represents the Horse animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Horse"
	horseNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Horse",
	}

	// horseNameDutch represents the Horse animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Horse"
	horseNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Horse",
	}

	// horseNameFrench represents the Horse animal type's name in French.
	// TODO: Verify translation accuracy for "Horse"
	horseNameFrench = nook.Name{
		Language: language.French,
		Value:    "Horse",
	}

	// horseNameGerman represents the Horse animal type's name in German.
	// TODO: Verify translation accuracy for "Horse"
	horseNameGerman = nook.Name{
		Language: language.German,
		Value:    "Horse",
	}

	// horseNameItalian represents the Horse animal type's name in Italian.
	// TODO: Verify translation accuracy for "Horse"
	horseNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Horse",
	}

	// horseNameJapanese represents the Horse animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Horse"
	horseNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Horse",
	}

	// horseNameKorean represents the Horse animal type's name in Korean.
	// TODO: Verify translation accuracy for "Horse"
	horseNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Horse",
	}

	// horseNameLatinAmericanSpanish represents the Horse animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Horse"
	horseNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Horse",
	}

	// horseNameRussian represents the Horse animal type's name in Russian.
	// TODO: Verify translation accuracy for "Horse"
	horseNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Horse",
	}

	// horseNameSimplifiedChinese represents the Horse animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Horse"
	horseNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Horse",
	}

	// horseNameSpanish represents the Horse animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Horse"
	horseNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Horse",
	}

	// horseNameTraditionalChinese represents the Horse animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Horse"
	horseNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Horse",
	}
)

var (
	// horseName contains the localized names of the Horse animal type.
	horseName = nook.Languages{
		language.AmericanEnglish:      horseNameAmericanEnglish,
		language.CanadianFrench:       horseNameCanadianFrench,
		language.Dutch:                horseNameDutch,
		language.French:               horseNameFrench,
		language.German:               horseNameGerman,
		language.Italian:              horseNameItalian,
		language.Japanese:             horseNameJapanese,
		language.Korean:               horseNameKorean,
		language.LatinAmericanSpanish: horseNameLatinAmericanSpanish,
		language.Russian:              horseNameRussian,
		language.SimplifiedChinese:    horseNameSimplifiedChinese,
		language.Spanish:              horseNameSpanish,
		language.TraditionalChinese:   horseNameTraditionalChinese,
	}
)

var (
	// Horse represents the Horse animal type in the Animal Crossing series.
	Horse = nook.Animal{
		Key:  nook.Key(horse),
		Name: horseName,
	}
)
