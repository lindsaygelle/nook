package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// reindeer is the common reference for the Reindeer animal type.
	reindeer = "Reindeer"
)

var (
	// reindeerNameAmericanEnglish represents the Reindeer animal type's name in American English.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Reindeer",
	}

	// reindeerNameCanadianFrench represents the Reindeer animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Reindeer",
	}

	// reindeerNameDutch represents the Reindeer animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Reindeer",
	}

	// reindeerNameFrench represents the Reindeer animal type's name in French.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameFrench = nook.Name{
		Language: language.French,
		Value:    "Reindeer",
	}

	// reindeerNameGerman represents the Reindeer animal type's name in German.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameGerman = nook.Name{
		Language: language.German,
		Value:    "Reindeer",
	}

	// reindeerNameItalian represents the Reindeer animal type's name in Italian.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Reindeer",
	}

	// reindeerNameJapanese represents the Reindeer animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Reindeer",
	}

	// reindeerNameKorean represents the Reindeer animal type's name in Korean.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Reindeer",
	}

	// reindeerNameLatinAmericanSpanish represents the Reindeer animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Reindeer",
	}

	// reindeerNameRussian represents the Reindeer animal type's name in Russian.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Reindeer",
	}

	// reindeerNameSimplifiedChinese represents the Reindeer animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Reindeer",
	}

	// reindeerNameSpanish represents the Reindeer animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Reindeer",
	}

	// reindeerNameTraditionalChinese represents the Reindeer animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Reindeer"
	reindeerNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Reindeer",
	}
)

var (
	// reindeerName contains the localized names of the Reindeer animal type.
	reindeerName = nook.Languages{
		language.AmericanEnglish:      reindeerNameAmericanEnglish,
		language.CanadianFrench:       reindeerNameCanadianFrench,
		language.Dutch:                reindeerNameDutch,
		language.French:               reindeerNameFrench,
		language.German:               reindeerNameGerman,
		language.Italian:              reindeerNameItalian,
		language.Japanese:             reindeerNameJapanese,
		language.Korean:               reindeerNameKorean,
		language.LatinAmericanSpanish: reindeerNameLatinAmericanSpanish,
		language.Russian:              reindeerNameRussian,
		language.SimplifiedChinese:    reindeerNameSimplifiedChinese,
		language.Spanish:              reindeerNameSpanish,
		language.TraditionalChinese:   reindeerNameTraditionalChinese,
	}
)

var (
	// Reindeer represents the Reindeer animal type in the Animal Crossing series.
	Reindeer = nook.Animal{
		Key:  nook.Key(reindeer),
		Name: reindeerName,
	}
)
