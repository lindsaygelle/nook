package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// sheep is the common reference for the Sheep animal type.
	sheep = "Sheep"
)

var (
	// sheepNameAmericanEnglish represents the Sheep animal type's name in American English.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sheep",
	}

	// sheepNameCanadianFrench represents the Sheep animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sheep",
	}

	// sheepNameDutch represents the Sheep animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Sheep",
	}

	// sheepNameFrench represents the Sheep animal type's name in French.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameFrench = nook.Name{
		Language: language.French,
		Value:    "Sheep",
	}

	// sheepNameGerman represents the Sheep animal type's name in German.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameGerman = nook.Name{
		Language: language.German,
		Value:    "Sheep",
	}

	// sheepNameItalian represents the Sheep animal type's name in Italian.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Sheep",
	}

	// sheepNameJapanese represents the Sheep animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Sheep",
	}

	// sheepNameKorean represents the Sheep animal type's name in Korean.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Sheep",
	}

	// sheepNameLatinAmericanSpanish represents the Sheep animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sheep",
	}

	// sheepNameRussian represents the Sheep animal type's name in Russian.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Sheep",
	}

	// sheepNameSimplifiedChinese represents the Sheep animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Sheep",
	}

	// sheepNameSpanish represents the Sheep animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Sheep",
	}

	// sheepNameTraditionalChinese represents the Sheep animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Sheep"
	sheepNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Sheep",
	}
)

var (
	// sheepName contains the localized names of the Sheep animal type.
	sheepName = nook.Languages{
		language.AmericanEnglish:      sheepNameAmericanEnglish,
		language.CanadianFrench:       sheepNameCanadianFrench,
		language.Dutch:                sheepNameDutch,
		language.French:               sheepNameFrench,
		language.German:               sheepNameGerman,
		language.Italian:              sheepNameItalian,
		language.Japanese:             sheepNameJapanese,
		language.Korean:               sheepNameKorean,
		language.LatinAmericanSpanish: sheepNameLatinAmericanSpanish,
		language.Russian:              sheepNameRussian,
		language.SimplifiedChinese:    sheepNameSimplifiedChinese,
		language.Spanish:              sheepNameSpanish,
		language.TraditionalChinese:   sheepNameTraditionalChinese,
	}
)

var (
	// Sheep represents the Sheep animal type in the Animal Crossing series.
	Sheep = nook.Animal{
		Key:  nook.Key(sheep),
		Name: sheepName,
	}
)
