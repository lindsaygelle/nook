package octopus

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Marina"}

	marinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marina"}

	marinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marina"}

	marinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Marianne"}

	marinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marina"}

	marinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タコリーナ"}

	marinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "문리나"}

	marinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marina"}

	marinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Марина"}

	marinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "章立娜"}

	marinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marina"}

	marinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "章立娜"}
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
		Animal:   animal.Octopus,
		Birthday: marinaBirthday,
		Code:     marinaCode,
		Key:      character.Marina,
		Gender:   gender.Female,
		Name:     marinaName,
		Special:  false}
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
		language.AmericanEnglish:      marinaAmericanEnglishPhrase,
		language.CanadianFrench:       marinaCanadianFrenchPhrase,
		language.Dutch:                marinaDutchPhrase,
		language.French:               marinaFrenchPhrase,
		language.German:               marinaGermanPhrase,
		language.Italian:              marinaItalianPhrase,
		language.Japanese:             marinaJapanesePhrase,
		language.Korean:               marinaKoreanPhrase,
		language.LatinAmericanSpanish: marinaLatinAmericanSpanishPhrase,
		language.Russian:              marinaRussianPhrase,
		language.SimplifiedChinese:    marinaSimplifiedChinesePhrase,
		language.Spanish:              marinaSpanishPhrase,
		language.TraditionalChinese:   marinaTraditionalChinesePhrase}
)

var (
	Marina = nook.Villager{
		Character:   marinaCharacter,
		Personality: personality.Normal,
		Phrase:      marinaPhrase}
)
