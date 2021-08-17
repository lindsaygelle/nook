package dog

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
	luckyBirthday = nook.Birthday{
		Day:   4,
		Month: time.November}
)

var (
	luckyCode = nook.Code{
		Value: "dog02"}
)

var (
	luckyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lucky"}

	luckyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ramsès"}

	luckyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lucky"}

	luckyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ramsès"}

	luckyGermanName = nook.Name{
		Language: language.German,
		Value:    "Viktor"}

	luckyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Felice"}

	luckyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラッキー"}

	luckyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "럭키"}

	luckyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pupas"}

	luckyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лаки"}

	luckySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大吉"}

	luckySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pupas"}

	luckyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大吉"}
)

var (
	luckyName = nook.Languages{
		language.AmericanEnglish:      luckyAmericanEnglishName,
		language.CanadianFrench:       luckyCanadianFrenchName,
		language.Dutch:                luckyDutchName,
		language.French:               luckyFrenchName,
		language.German:               luckyGermanName,
		language.Italian:              luckyItalianName,
		language.Japanese:             luckyJapaneseName,
		language.Korean:               luckyKoreanName,
		language.LatinAmericanSpanish: luckyLatinAmericanSpanishName,
		language.Russian:              luckyRussianName,
		language.SimplifiedChinese:    luckySimplifiedChineseName,
		language.Spanish:              luckySpanishName,
		language.TraditionalChinese:   luckyTraditionalChineseName}
)

var (
	luckyCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: luckyBirthday,
		Code:     luckyCode,
		Key:      character.Lucky,
		Gender:   gender.Male,
		Name:     luckyName}
)

var (
	luckyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "rrr-owch"}

	luckyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "graaaouch"}

	luckyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "www-⁠au"}

	luckyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "graaaouch"}

	luckyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "rrr-awau"}

	luckyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ossorrr"}

	luckyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "らしいよ"}

	luckyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇대"}

	luckyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guuuuf"}

	luckyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ррр-ой-ой"}

	luckySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大概"}

	luckySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guuuuAY"}

	luckyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大概"}
)

var (
	luckyPhrase = nook.Languages{
		language.AmericanEnglish:      luckyAmericanEnglishPhrase,
		language.CanadianFrench:       luckyCanadianFrenchPhrase,
		language.Dutch:                luckyDutchPhrase,
		language.French:               luckyFrenchPhrase,
		language.German:               luckyGermanPhrase,
		language.Italian:              luckyItalianPhrase,
		language.Japanese:             luckyJapanesePhrase,
		language.Korean:               luckyKoreanPhrase,
		language.LatinAmericanSpanish: luckyLatinAmericanSpanishPhrase,
		language.Russian:              luckyRussianPhrase,
		language.SimplifiedChinese:    luckySimplifiedChinesePhrase,
		language.Spanish:              luckySpanishPhrase,
		language.TraditionalChinese:   luckyTraditionalChinesePhrase}
)

var (
	Lucky = nook.Villager{
		Character:   luckyCharacter,
		Personality: personality.Lazy,
		Phrase:      luckyPhrase}
)
