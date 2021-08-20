package bird

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
		Animal:   animal.Bird,
		Birthday: luchaBirthday,
		Code:     luchaCode,
		Key:      character.Lucha,
		Gender:   gender.Male,
		Name:     luchaName,
		Special:  false}
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
		language.AmericanEnglish:      luchaAmericanEnglishPhrase,
		language.CanadianFrench:       luchaCanadianFrenchPhrase,
		language.Dutch:                luchaDutchPhrase,
		language.French:               luchaFrenchPhrase,
		language.German:               luchaGermanPhrase,
		language.Italian:              luchaItalianPhrase,
		language.Japanese:             luchaJapanesePhrase,
		language.Korean:               luchaKoreanPhrase,
		language.LatinAmericanSpanish: luchaLatinAmericanSpanishPhrase,
		language.Russian:              luchaRussianPhrase,
		language.SimplifiedChinese:    luchaSimplifiedChinesePhrase,
		language.Spanish:              luchaSpanishPhrase,
		language.TraditionalChinese:   luchaTraditionalChinesePhrase}
)

var (
	Lucha = nook.Villager{
		Character:   luchaCharacter,
		Personality: personality.Smug,
		Phrase:      luchaPhrase}
)
