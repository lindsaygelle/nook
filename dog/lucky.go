package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Ramsèsgraaaouch"}

	luckyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Luckywww-⁠au"}

	luckyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ramsèsgraaaouch"}

	luckyGermanName = nook.Name{
		Language: language.German,
		Value:    "Viktorrrr-awau"}

	luckyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Feliceossorrr"}

	luckyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラッキーらしいよ"}

	luckyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "럭키그렇대"}

	luckyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pupasguuuuf"}

	luckyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лакиррр-ой-ой"}

	luckySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大吉大概"}

	luckySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "PupasguuuuAY"}

	luckyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大吉大概"}
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
		Animal:   Dog,
		Birthday: luckyBirthday,
		Code:     luckyCode,
		Gender:   nook.Male,
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
	Lucky = nook.Villager{
		Character:   luckyCharacter,
		Personality: nook.Lazy,
		Phrase:      luckyPhrase}
)
