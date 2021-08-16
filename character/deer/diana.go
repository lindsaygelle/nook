package deer

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
	dianaBirthday = nook.Birthday{
		Day:   4,
		Month: time.January}
)

var (
	dianaCode = nook.Code{
		Value: "der08"}
)

var (
	dianaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Diana"}

	dianaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Didi"}

	dianaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Diana"}

	dianaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Didi"}

	dianaGermanName = nook.Name{
		Language: language.German,
		Value:    "Vroni"}

	dianaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Diana"}

	dianaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナタリー"}

	dianaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나탈리"}

	dianaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bambina"}

	dianaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Диана"}

	dianaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "倪家莉"}

	dianaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bambina"}

	dianaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "倪家莉"}
)

var (
	dianaName = nook.Languages{
		language.AmericanEnglish:      dianaAmericanEnglishName,
		language.CanadianFrench:       dianaCanadianFrenchName,
		language.Dutch:                dianaDutchName,
		language.French:               dianaFrenchName,
		language.German:               dianaGermanName,
		language.Italian:              dianaItalianName,
		language.Japanese:             dianaJapaneseName,
		language.Korean:               dianaKoreanName,
		language.LatinAmericanSpanish: dianaLatinAmericanSpanishName,
		language.Russian:              dianaRussianName,
		language.SimplifiedChinese:    dianaSimplifiedChineseName,
		language.Spanish:              dianaSpanishName,
		language.TraditionalChinese:   dianaTraditionalChineseName}
)

var (
	dianaCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: dianaBirthday,
		Code:     dianaCode,
		Key:      character.Diana,
		Gender:   gender.Female,
		Name:     dianaName}
)

var (
	dianaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "no doy"}

	dianaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bichouchou"}

	dianaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "joh"}

	dianaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bichouchou"}

	dianaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kitz"}

	dianaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mamy"}

	dianaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でしょ"}

	dianaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라니까"}

	dianaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nonuá"}

	dianaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "еще бы"}

	dianaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是吧"}

	dianaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "braaavo"}

	dianaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是吧"}
)

var (
	dianaPhrase = nook.Languages{
		language.AmericanEnglish:      dianaAmericanEnglishName,
		language.CanadianFrench:       dianaCanadianFrenchName,
		language.Dutch:                dianaDutchName,
		language.French:               dianaFrenchName,
		language.German:               dianaGermanName,
		language.Italian:              dianaItalianName,
		language.Japanese:             dianaJapaneseName,
		language.Korean:               dianaKoreanName,
		language.LatinAmericanSpanish: dianaLatinAmericanSpanishName,
		language.Russian:              dianaRussianName,
		language.SimplifiedChinese:    dianaSimplifiedChineseName,
		language.Spanish:              dianaSpanishName,
		language.TraditionalChinese:   dianaTraditionalChineseName}
)

var (
	Diana = nook.Villager{
		Character:   dianaCharacter,
		Personality: personality.Snooty,
		Phrase:      dianaPhrase}
)
