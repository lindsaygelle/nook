package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// seagull is the common reference for the Seagull animal type.
	seagull = "Seagull"
)

var (
	// seagullNameAmericanEnglish represents the Seagull animal type's name in American English.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Seagull",
	}

	// seagullNameCanadianFrench represents the Seagull animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Seagull",
	}

	// seagullNameDutch represents the Seagull animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Seagull",
	}

	// seagullNameFrench represents the Seagull animal type's name in French.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameFrench = nook.Name{
		Language: language.French,
		Value:    "Seagull",
	}

	// seagullNameGerman represents the Seagull animal type's name in German.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameGerman = nook.Name{
		Language: language.German,
		Value:    "Seagull",
	}

	// seagullNameItalian represents the Seagull animal type's name in Italian.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Seagull",
	}

	// seagullNameJapanese represents the Seagull animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Seagull",
	}

	// seagullNameKorean represents the Seagull animal type's name in Korean.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Seagull",
	}

	// seagullNameLatinAmericanSpanish represents the Seagull animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Seagull",
	}

	// seagullNameRussian represents the Seagull animal type's name in Russian.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Seagull",
	}

	// seagullNameSimplifiedChinese represents the Seagull animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Seagull",
	}

	// seagullNameSpanish represents the Seagull animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Seagull",
	}

	// seagullNameTraditionalChinese represents the Seagull animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Seagull"
	seagullNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Seagull",
	}
)

var (
	// seagullName contains the localized names of the Seagull animal type.
	seagullName = nook.Languages{
		language.AmericanEnglish:      seagullNameAmericanEnglish,
		language.CanadianFrench:       seagullNameCanadianFrench,
		language.Dutch:                seagullNameDutch,
		language.French:               seagullNameFrench,
		language.German:               seagullNameGerman,
		language.Italian:              seagullNameItalian,
		language.Japanese:             seagullNameJapanese,
		language.Korean:               seagullNameKorean,
		language.LatinAmericanSpanish: seagullNameLatinAmericanSpanish,
		language.Russian:              seagullNameRussian,
		language.SimplifiedChinese:    seagullNameSimplifiedChinese,
		language.Spanish:              seagullNameSpanish,
		language.TraditionalChinese:   seagullNameTraditionalChinese,
	}
)

var (
	// Seagull represents the Seagull animal type in the Animal Crossing series.
	Seagull = nook.Animal{
		Key:  nook.Key(seagull),
		Name: seagullName,
	}
)
