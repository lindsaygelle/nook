package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// chameleon is the common reference for the Chameleon animal type.
	chameleon = "Chameleon"
)

var (
	// chameleonNameAmericanEnglish represents the Chameleon animal type's name in American English.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chameleon",
	}

	// chameleonNameCanadianFrench represents the Chameleon animal type's name in Canadian French.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameCanadianFrench = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chameleon",
	}

	// chameleonNameDutch represents the Chameleon animal type's name in Dutch.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameDutch = nook.Name{
		Language: language.Dutch,
		Value:    "Chameleon",
	}

	// chameleonNameFrench represents the Chameleon animal type's name in French.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameFrench = nook.Name{
		Language: language.French,
		Value:    "Chameleon",
	}

	// chameleonNameGerman represents the Chameleon animal type's name in German.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameGerman = nook.Name{
		Language: language.German,
		Value:    "Chameleon",
	}

	// chameleonNameItalian represents the Chameleon animal type's name in Italian.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameItalian = nook.Name{
		Language: language.Italian,
		Value:    "Chameleon",
	}

	// chameleonNameJapanese represents the Chameleon animal type's name in Japanese.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameJapanese = nook.Name{
		Language: language.Japanese,
		Value:    "Chameleon",
	}

	// chameleonNameKorean represents the Chameleon animal type's name in Korean.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameKorean = nook.Name{
		Language: language.Korean,
		Value:    "Chameleon",
	}

	// chameleonNameLatinAmericanSpanish represents the Chameleon animal type's name in Latin American Spanish.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameLatinAmericanSpanish = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chameleon",
	}

	// chameleonNameRussian represents the Chameleon animal type's name in Russian.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameRussian = nook.Name{
		Language: language.Russian,
		Value:    "Chameleon",
	}

	// chameleonNameSimplifiedChinese represents the Chameleon animal type's name in Simplified Chinese.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameSimplifiedChinese = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Chameleon",
	}

	// chameleonNameSpanish represents the Chameleon animal type's name in Spanish.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameSpanish = nook.Name{
		Language: language.Spanish,
		Value:    "Chameleon",
	}

	// chameleonNameTraditionalChinese represents the Chameleon animal type's name in Traditional Chinese.
	// TODO: Verify translation accuracy for "Chameleon"
	chameleonNameTraditionalChinese = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Chameleon",
	}
)

var (
	// chameleonName contains the localized names of the Chameleon animal type.
	chameleonName = nook.Languages{
		language.AmericanEnglish:      chameleonNameAmericanEnglish,
		language.CanadianFrench:       chameleonNameCanadianFrench,
		language.Dutch:                chameleonNameDutch,
		language.French:               chameleonNameFrench,
		language.German:               chameleonNameGerman,
		language.Italian:              chameleonNameItalian,
		language.Japanese:             chameleonNameJapanese,
		language.Korean:               chameleonNameKorean,
		language.LatinAmericanSpanish: chameleonNameLatinAmericanSpanish,
		language.Russian:              chameleonNameRussian,
		language.SimplifiedChinese:    chameleonNameSimplifiedChinese,
		language.Spanish:              chameleonNameSpanish,
		language.TraditionalChinese:   chameleonNameTraditionalChinese,
	}
)

var (
	// Chameleon represents the Chameleon animal type in the Animal Crossing series.
	Chameleon = nook.Animal{
		Key:  nook.Key(chameleon),
		Name: chameleonName,
	}
)
