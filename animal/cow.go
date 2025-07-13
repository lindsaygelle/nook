package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// cow is the common reference for the Cow animal type.
	cow = "Cow"
)

var (
	// cowNameAmericanEnglish represents the Cow animal type's name in American English.
	// TODO: Verify translation accuracy for "Cow"
	cowNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cow",
	}

	// cowNameCanadianFrench represents the Cow animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Cow"
	cowNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cow",
	}

	// cowNameDutch represents the Cow animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Cow"
	cowNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Cow",
	}

	// cowNameFrench represents the Cow animal type's name in French.
	// TODO: Verify translation accuracy for "Cow"
	cowNameFrench = nook.Name{
		Language: language.French,
		Value:    "Cow",
	}

	// cowNameGerman represents the Cow animal type's name in German.
	// TODO: Verify translation accuracy for "Cow"
	cowNameGerman = nook.Name{
		Language: language.German,
		Value:    "Cow",
	}

	// cowNameItalian represents the Cow animal type's name in Italian.
	// TODO: Verify translation accuracy for "Cow"
	cowNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Cow",
	}

	// cowNameJapanese represents the Cow animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Cow"
	cowNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Cow",
	}

	// cowNameKorean represents the Cow animal type's name in Korean.
	// TODO: Verify translation accuracy for "Cow"
	cowNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Cow",
	}

	// cowNameLatinAmericanSpanish represents the Cow animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Cow"
	cowNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cow",
	}

	// cowNameRussian represents the Cow animal type's name in Russian.
	// TODO: Verify translation accuracy for "Cow"
	cowNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Cow",
	}

	// cowNameSimplifiedChinese represents the Cow animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Cow"
	cowNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Cow",
	}

	// cowNameSpanish represents the Cow animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Cow"
	cowNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Cow",
	}

	// cowNameTraditionalChinese represents the Cow animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Cow"
	cowNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Cow",
	}
)

var (
	// cowName contains the localized names of the Cow animal type.
	cowName = nook.Languages{
		language.AmericanEnglish:      cowNameAmericanEnglish,
		language.CanadianFrench:       cowNameCanadianFrench,
		language.Dutch:                cowNameDutch,
		language.French:               cowNameFrench,
		language.German:               cowNameGerman,
		language.Italian:              cowNameItalian,
		language.Japanese:             cowNameJapanese,
		language.Korean:               cowNameKorean,
		language.LatinAmericanSpanish: cowNameLatinAmericanSpanish,
		language.Russian:              cowNameRussian,
		language.SimplifiedChinese:    cowNameSimplifiedChinese,
		language.Spanish:              cowNameSpanish,
		language.TraditionalChinese:   cowNameTraditionalChinese,
	}
)

var (
	// Cow represents the Cow animal type in the Animal Crossing series.
	Cow = nook.Animal{
		Key:  nook.Key(cow),
		Name: cowName,
	}
)
