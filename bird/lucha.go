package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	luchaBirthday = nook.Birthday{
		Day:   12,
		Month: time.December}
)

var (
	luchaCode = nook.Code{
		Value: "brd15"}
)

var (
	luchaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lucha"}

	luchaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Condor"}

	luchaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lucha"}

	luchaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Condor"}

	luchaGermanName = nook.Name{
		Language: language.German,
		Value:    "Lukas"}

	luchaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Loreto"}

	luchaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マスカラス"}

	luchaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마스카라스"}

	luchaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Plumerio"}

	luchaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Луча"}

	luchaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "摔角鸦"}

	luchaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Plumerio"}

	luchaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "摔角鴉"}
)

var (
	luchaName = nook.Languages{
		language.AmericanEnglish:      luchaAmericanEnglishName,
		language.CanadianFrench:       luchaCanadianFrenchName,
		language.Dutch:                luchaDutchName,
		language.French:               luchaFrenchName,
		language.German:               luchaGermanName,
		language.Italian:              luchaItalianName,
		language.Japanese:             luchaJapaneseName,
		language.Korean:               luchaKoreanName,
		language.LatinAmericanSpanish: luchaLatinAmericanSpanishName,
		language.Russian:              luchaRussianName,
		language.SimplifiedChinese:    luchaSimplifiedChineseName,
		language.Spanish:              luchaSpanishName,
		language.TraditionalChinese:   luchaTraditionalChineseName}
)

var (
	luchaCharacter = nook.Character{
		Animal:   Bird,
		Birthday: luchaBirthday,
		Code:     luchaCode,
		Gender:   nook.Male,
		Name:     luchaName}
)

var (
	luchaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cacaw"}

	luchaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "yaaaah"}

	luchaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kakauw"}

	luchaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "yaaaah"}

	luchaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mauser"}

	luchaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "lorito"}

	luchaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "カァー"}

	luchaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "까아악"}

	luchaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pío"}

	luchaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кар-ко"}

	luchaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘎"}

	luchaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "picopico"}

	luchaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘎"}
)

var (
	luchaPhrase = nook.Languages{
		language.AmericanEnglish:      luchaAmericanEnglishName,
		language.CanadianFrench:       luchaCanadianFrenchName,
		language.Dutch:                luchaDutchName,
		language.French:               luchaFrenchName,
		language.German:               luchaGermanName,
		language.Italian:              luchaItalianName,
		language.Japanese:             luchaJapaneseName,
		language.Korean:               luchaKoreanName,
		language.LatinAmericanSpanish: luchaLatinAmericanSpanishName,
		language.Russian:              luchaRussianName,
		language.SimplifiedChinese:    luchaSimplifiedChineseName,
		language.Spanish:              luchaSpanishName,
		language.TraditionalChinese:   luchaTraditionalChineseName}
)

var (
	Lucha = nook.Villager{
		Character:   luchaCharacter,
		Personality: nook.Smug,
		Phrase:      luchaPhrase}
)
