package rabbit

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
	tobyBirthday = nook.Birthday{
		Day:   10,
		Month: time.July}
)

var (
	tobyCode = nook.Code{
		Value: "rbt20"}
)

var (
	tobyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Toby"}

	tobyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Toby"}

	tobyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Toby"}

	tobyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Toby"}

	tobyGermanName = nook.Name{
		Language: language.German,
		Value:    "Toby"}

	tobyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Toby"}

	tobyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トビー"}

	tobyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토비"}

	tobyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Toby"}

	tobyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тоби"}

	tobySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咚比"}

	tobySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Toby"}

	tobyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咚比"}
)

var (
	tobyName = nook.Languages{
		language.AmericanEnglish:      tobyAmericanEnglishName,
		language.CanadianFrench:       tobyCanadianFrenchName,
		language.Dutch:                tobyDutchName,
		language.French:               tobyFrenchName,
		language.German:               tobyGermanName,
		language.Italian:              tobyItalianName,
		language.Japanese:             tobyJapaneseName,
		language.Korean:               tobyKoreanName,
		language.LatinAmericanSpanish: tobyLatinAmericanSpanishName,
		language.Russian:              tobyRussianName,
		language.SimplifiedChinese:    tobySimplifiedChineseName,
		language.Spanish:              tobySpanishName,
		language.TraditionalChinese:   tobyTraditionalChineseName}
)

var (
	tobyCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: tobyBirthday,
		Code:     tobyCode,
		Key:      character.Toby,
		Gender:   gender.Male,
		Name:     tobyName}
)

var (
	tobyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ribbit"}

	tobyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kérocarot"}

	tobyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kikk'r"}

	tobyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "kérocarot"}

	tobyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kerokero"}

	tobyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "kerokero"}

	tobyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "けろけろ"}

	tobyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "케로케로"}

	tobyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "keroppi"}

	tobyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайцеквак"}

	tobySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蛙蛙"}

	tobySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "keroppi"}

	tobyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蛙蛙"}
)

var (
	tobyPhrase = nook.Languages{
		language.AmericanEnglish:      tobyAmericanEnglishName,
		language.CanadianFrench:       tobyCanadianFrenchName,
		language.Dutch:                tobyDutchName,
		language.French:               tobyFrenchName,
		language.German:               tobyGermanName,
		language.Italian:              tobyItalianName,
		language.Japanese:             tobyJapaneseName,
		language.Korean:               tobyKoreanName,
		language.LatinAmericanSpanish: tobyLatinAmericanSpanishName,
		language.Russian:              tobyRussianName,
		language.SimplifiedChinese:    tobySimplifiedChineseName,
		language.Spanish:              tobySpanishName,
		language.TraditionalChinese:   tobyTraditionalChineseName}
)

var (
	Toby = nook.Villager{
		Character:   tobyCharacter,
		Personality: personality.Smug,
		Phrase:      tobyPhrase}
)
