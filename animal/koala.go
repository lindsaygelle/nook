package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// koala is the common reference for the Koala animal type.
	koala = "Koala"
)

var (
	// koalaNameAmericanEnglish represents the Koala animal type's name in American English.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Koala",
	}

	// koalaNameCanadianFrench represents the Koala animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Koala",
	}

	// koalaNameDutch represents the Koala animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Koala",
	}

	// koalaNameFrench represents the Koala animal type's name in French.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameFrench = nook.Name{
		Language: language.French,
		Value:    "Koala",
	}

	// koalaNameGerman represents the Koala animal type's name in German.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameGerman = nook.Name{
		Language: language.German,
		Value:    "Koala",
	}

	// koalaNameItalian represents the Koala animal type's name in Italian.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Koala",
	}

	// koalaNameJapanese represents the Koala animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Koala",
	}

	// koalaNameKorean represents the Koala animal type's name in Korean.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Koala",
	}

	// koalaNameLatinAmericanSpanish represents the Koala animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Koala",
	}

	// koalaNameRussian represents the Koala animal type's name in Russian.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Koala",
	}

	// koalaNameSimplifiedChinese represents the Koala animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Koala",
	}

	// koalaNameSpanish represents the Koala animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Koala",
	}

	// koalaNameTraditionalChinese represents the Koala animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Koala"
	koalaNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Koala",
	}
)

var (
	// koalaName contains the localized names of the Koala animal type.
	koalaName = nook.Languages{
		language.AmericanEnglish:      koalaNameAmericanEnglish,
		language.CanadianFrench:       koalaNameCanadianFrench,
		language.Dutch:                koalaNameDutch,
		language.French:               koalaNameFrench,
		language.German:               koalaNameGerman,
		language.Italian:              koalaNameItalian,
		language.Japanese:             koalaNameJapanese,
		language.Korean:               koalaNameKorean,
		language.LatinAmericanSpanish: koalaNameLatinAmericanSpanish,
		language.Russian:              koalaNameRussian,
		language.SimplifiedChinese:    koalaNameSimplifiedChinese,
		language.Spanish:              koalaNameSpanish,
		language.TraditionalChinese:   koalaNameTraditionalChinese,
	}
)

var (
	// Koala represents the Koala animal type in the Animal Crossing series.
	Koala = nook.Animal{
		Key:  nook.Key(koala),
		Name: koalaName,
	}
)
