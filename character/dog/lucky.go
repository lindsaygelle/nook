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
	// luckyBirthday represents lucky birthday.
	luckyBirthday = nook.Birthday{
		Day:   4,
		Month: time.November}
)

var (
	// luckyCode represents lucky code.
	luckyCode = nook.Code{
		Value: "dog02"}
)

var (
	// luckyAmericanEnglishName represents lucky american english name.
	luckyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lucky"}

	// luckyCanadianFrenchName represents lucky canadian french name.
	luckyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ramsès"}

	// luckyDutchName represents lucky dutch name.
	luckyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lucky"}

	// luckyFrenchName represents lucky french name.
	luckyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ramsès"}

	// luckyGermanName represents lucky german name.
	luckyGermanName = nook.Name{
		Language: language.German,
		Value:    "Viktor"}

	// luckyItalianName represents lucky italian name.
	luckyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Felice"}

	// luckyJapaneseName represents lucky japanese name.
	luckyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラッキー"}

	// luckyKoreanName represents lucky korean name.
	luckyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "럭키"}

	// luckyLatinAmericanSpanishName represents lucky latin american spanish name.
	luckyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pupas"}

	// luckyRussianName represents lucky russian name.
	luckyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лаки"}

	// luckySimplifiedChineseName represents lucky simplified chinese name.
	luckySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大吉"}

	// luckySpanishName represents lucky spanish name.
	luckySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pupas"}

	// luckyTraditionalChineseName represents lucky traditional chinese name.
	luckyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大吉"}
)

var (
	// luckyName represents lucky name.
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
	// luckyCharacter represents lucky character.
	luckyCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: luckyBirthday,
		Code:     luckyCode,
		Key:      character.Lucky,
		Gender:   gender.Male,
		Name:     luckyName,
		Special:  false}
)

var (
	// luckyAmericanEnglishPhrase represents lucky american english phrase.
	luckyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "rrr-owch"}

	// luckyCanadianFrenchPhrase represents lucky canadian french phrase.
	luckyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "graaaouch"}

	// luckyDutchPhrase represents lucky dutch phrase.
	luckyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "www-⁠au"}

	// luckyFrenchPhrase represents lucky french phrase.
	luckyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "graaaouch"}

	// luckyGermanPhrase represents lucky german phrase.
	luckyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "rrr-awau"}

	// luckyItalianPhrase represents lucky italian phrase.
	luckyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ossorrr"}

	// luckyJapanesePhrase represents lucky japanese phrase.
	luckyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "らしいよ"}

	// luckyKoreanPhrase represents lucky korean phrase.
	luckyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇대"}

	// luckyLatinAmericanSpanishPhrase represents lucky latin american spanish phrase.
	luckyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guuuuf"}

	// luckyRussianPhrase represents lucky russian phrase.
	luckyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ррр-ой-ой"}

	// luckySimplifiedChinesePhrase represents lucky simplified chinese phrase.
	luckySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大概"}

	// luckySpanishPhrase represents lucky spanish phrase.
	luckySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guuuuAY"}

	// luckyTraditionalChinesePhrase represents lucky traditional chinese phrase.
	luckyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大概"}
)

var (
	// luckyPhrase represents lucky phrase.
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
	// Lucky represents lucky.
	Lucky = nook.Villager{
		Character:   luckyCharacter,
		Personality: personality.Lazy,
		Phrase:      luckyPhrase}
)
