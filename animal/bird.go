package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// bird is the common reference for the Bird animal type.
	bird = "Bird"
)

var (
	// birdNameAmericanEnglish represents the Bird animal type's name in American English.
	// TODO: Verify translation accuracy for "Bird"
	birdNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bird",
	}

	// birdNameCanadianFrench represents the Bird animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Bird"
	birdNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bird",
	}

	// birdNameDutch represents the Bird animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Bird"
	birdNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Bird",
	}

	// birdNameFrench represents the Bird animal type's name in French.
	// TODO: Verify translation accuracy for "Bird"
	birdNameFrench = nook.Name{
		Language: language.French,
		Value:    "Bird",
	}

	// birdNameGerman represents the Bird animal type's name in German.
	// TODO: Verify translation accuracy for "Bird"
	birdNameGerman = nook.Name{
		Language: language.German,
		Value:    "Bird",
	}

	// birdNameItalian represents the Bird animal type's name in Italian.
	// TODO: Verify translation accuracy for "Bird"
	birdNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Bird",
	}

	// birdNameJapanese represents the Bird animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Bird"
	birdNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Bird",
	}

	// birdNameKorean represents the Bird animal type's name in Korean.
	// TODO: Verify translation accuracy for "Bird"
	birdNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Bird",
	}

	// birdNameLatinAmericanSpanish represents the Bird animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Bird"
	birdNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bird",
	}

	// birdNameRussian represents the Bird animal type's name in Russian.
	// TODO: Verify translation accuracy for "Bird"
	birdNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Bird",
	}

	// birdNameSimplifiedChinese represents the Bird animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Bird"
	birdNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Bird",
	}

	// birdNameSpanish represents the Bird animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Bird"
	birdNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Bird",
	}

	// birdNameTraditionalChinese represents the Bird animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Bird"
	birdNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Bird",
	}
)

var (
	// birdName contains the localized names of the Bird animal type.
	birdName = nook.Languages{
		language.AmericanEnglish:      birdNameAmericanEnglish,
		language.CanadianFrench:       birdNameCanadianFrench,
		language.Dutch:                birdNameDutch,
		language.French:               birdNameFrench,
		language.German:               birdNameGerman,
		language.Italian:              birdNameItalian,
		language.Japanese:             birdNameJapanese,
		language.Korean:               birdNameKorean,
		language.LatinAmericanSpanish: birdNameLatinAmericanSpanish,
		language.Russian:              birdNameRussian,
		language.SimplifiedChinese:    birdNameSimplifiedChinese,
		language.Spanish:              birdNameSpanish,
		language.TraditionalChinese:   birdNameTraditionalChinese,
	}
)

var (
	// Bird represents the Bird animal type in the Animal Crossing series.
	Bird = nook.Animal{
		Key:  nook.Key(bird),
		Name: birdName,
	}
)
