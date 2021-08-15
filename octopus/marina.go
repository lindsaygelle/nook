package octopus

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	marinaBirthday = nook.Birthday{
		Day:   26,
		Month: time.June}
)

var (
	marinaCode = nook.Code{
		Value: "ocp01"}
)

var (
	marinaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Marina"}

	marinaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marinabeurp"}

	marinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marinablorp"}

	marinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marinabeurp"}

	marinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Marianneblubblub"}

	marinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marinablurp"}

	marinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タコリーナきゃ"}

	marinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "문리나캬캬"}

	marinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marinabliup"}

	marinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Маринахлюп"}

	marinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "章立娜咔"}

	marinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marinadospatas"}

	marinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "章立娜咔"}
)

var (
	marinaName = nook.Languages{
		language.AmericanEnglish:      marinaAmericanEnglishName,
		language.CanadianFrench:       marinaCanadianFrenchName,
		language.Dutch:                marinaDutchName,
		language.French:               marinaFrenchName,
		language.German:               marinaGermanName,
		language.Italian:              marinaItalianName,
		language.Japanese:             marinaJapaneseName,
		language.Korean:               marinaKoreanName,
		language.LatinAmericanSpanish: marinaLatinAmericanSpanishName,
		language.Russian:              marinaRussianName,
		language.SimplifiedChinese:    marinaSimplifiedChineseName,
		language.Spanish:              marinaSpanishName,
		language.TraditionalChinese:   marinaTraditionalChineseName}
)

var (
	marinaCharacter = nook.Character{
		Animal:   Octopus,
		Birthday: marinaBirthday,
		Code:     marinaCode,
		Gender:   nook.Female,
		Name:     marinaName}
)

var (
	marinaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "blurp"}

	marinaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "beurp"}

	marinaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "blorp"}

	marinaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "beurp"}

	marinaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "blubblub"}

	marinaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blurp"}

	marinaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "きゃ"}

	marinaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "캬캬"}

	marinaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bliup"}

	marinaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хлюп"}

	marinaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咔"}

	marinaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "dospatas"}

	marinaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咔"}
)

var (
	marinaPhrase = nook.Languages{
		language.AmericanEnglish:      marinaAmericanEnglishName,
		language.CanadianFrench:       marinaCanadianFrenchName,
		language.Dutch:                marinaDutchName,
		language.French:               marinaFrenchName,
		language.German:               marinaGermanName,
		language.Italian:              marinaItalianName,
		language.Japanese:             marinaJapaneseName,
		language.Korean:               marinaKoreanName,
		language.LatinAmericanSpanish: marinaLatinAmericanSpanishName,
		language.Russian:              marinaRussianName,
		language.SimplifiedChinese:    marinaSimplifiedChineseName,
		language.Spanish:              marinaSpanishName,
		language.TraditionalChinese:   marinaTraditionalChineseName}
)

var (
	Marina = nook.Villager{
		Character:   marinaCharacter,
		Personality: nook.Normal,
		Phrase:      marinaPhrase}
)
