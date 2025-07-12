package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// dog is the common reference for the Dog animal type.
	dog = "Dog"
)

var (
	// dogNameAmericanEnglish represents the Dog animal type's name in American English.
	// TODO: Verify translation accuracy for "Dog"
	dogNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dog",
	}

	// dogNameCanadianFrench represents the Dog animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Dog"
	dogNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dog",
	}

	// dogNameDutch represents the Dog animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Dog"
	dogNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Dog",
	}

	// dogNameFrench represents the Dog animal type's name in French.
	// TODO: Verify translation accuracy for "Dog"
	dogNameFrench = nook.Name{
		Language: language.French,
		Value:    "Dog",
	}

	// dogNameGerman represents the Dog animal type's name in German.
	// TODO: Verify translation accuracy for "Dog"
	dogNameGerman = nook.Name{
		Language: language.German,
		Value:    "Dog",
	}

	// dogNameItalian represents the Dog animal type's name in Italian.
	// TODO: Verify translation accuracy for "Dog"
	dogNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Dog",
	}

	// dogNameJapanese represents the Dog animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Dog"
	dogNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Dog",
	}

	// dogNameKorean represents the Dog animal type's name in Korean.
	// TODO: Verify translation accuracy for "Dog"
	dogNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Dog",
	}

	// dogNameLatinAmericanSpanish represents the Dog animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Dog"
	dogNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dog",
	}

	// dogNameRussian represents the Dog animal type's name in Russian.
	// TODO: Verify translation accuracy for "Dog"
	dogNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Dog",
	}

	// dogNameSimplifiedChinese represents the Dog animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Dog"
	dogNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Dog",
	}

	// dogNameSpanish represents the Dog animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Dog"
	dogNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Dog",
	}

	// dogNameTraditionalChinese represents the Dog animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Dog"
	dogNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Dog",
	}
)

var (
	// dogName contains the localized names of the Dog animal type.
	dogName = nook.Languages{
		language.AmericanEnglish:      dogNameAmericanEnglish,
		language.CanadianFrench:       dogNameCanadianFrench,
		language.Dutch:                dogNameDutch,
		language.French:               dogNameFrench,
		language.German:               dogNameGerman,
		language.Italian:              dogNameItalian,
		language.Japanese:             dogNameJapanese,
		language.Korean:               dogNameKorean,
		language.LatinAmericanSpanish: dogNameLatinAmericanSpanish,
		language.Russian:              dogNameRussian,
		language.SimplifiedChinese:    dogNameSimplifiedChinese,
		language.Spanish:              dogNameSpanish,
		language.TraditionalChinese:   dogNameTraditionalChinese,
	}
)

var (
	// Dog represents the Dog animal type in the Animal Crossing series.
	Dog = nook.Animal{
		Key:  nook.Key(dog),
		Name: dogName,
	}
)
