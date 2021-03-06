package frog

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
	puddlesBirthday = nook.Birthday{
		Day:   13,
		Month: time.January}
)

var (
	puddlesCode = nook.Code{
		Value: "flg06"}
)

var (
	puddlesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Puddles"}

	puddlesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rénata"}

	puddlesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Puddles"}

	puddlesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rénata"}

	puddlesGermanName = nook.Name{
		Language: language.German,
		Value:    "Nele"}

	puddlesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grazia"}

	puddlesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チョキ"}

	puddlesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "가위"}

	puddlesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nenúfar"}

	puddlesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рената"}

	puddlesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "剪刀"}

	puddlesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nenúfar"}

	puddlesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "剪刀"}
)

var (
	puddlesName = nook.Languages{
		language.AmericanEnglish:      puddlesAmericanEnglishName,
		language.CanadianFrench:       puddlesCanadianFrenchName,
		language.Dutch:                puddlesDutchName,
		language.French:               puddlesFrenchName,
		language.German:               puddlesGermanName,
		language.Italian:              puddlesItalianName,
		language.Japanese:             puddlesJapaneseName,
		language.Korean:               puddlesKoreanName,
		language.LatinAmericanSpanish: puddlesLatinAmericanSpanishName,
		language.Russian:              puddlesRussianName,
		language.SimplifiedChinese:    puddlesSimplifiedChineseName,
		language.Spanish:              puddlesSpanishName,
		language.TraditionalChinese:   puddlesTraditionalChineseName}
)

var (
	puddlesCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: puddlesBirthday,
		Code:     puddlesCode,
		Key:      character.Puddles,
		Gender:   gender.Female,
		Name:     puddlesName,
		Special:  false}
)

var (
	puddlesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "splish"}

	puddlesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "splufff"}

	puddlesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "plons"}

	puddlesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "splufff"}

	puddlesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "platschi"}

	puddlesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "splash"}

	puddlesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "っちゃ"}

	puddlesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그랬쪄"}

	puddlesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "esplís"}

	puddlesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "плюх"}

	puddlesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "石头"}

	puddlesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "esplís"}

	puddlesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "石頭"}
)

var (
	puddlesPhrase = nook.Languages{
		language.AmericanEnglish:      puddlesAmericanEnglishPhrase,
		language.CanadianFrench:       puddlesCanadianFrenchPhrase,
		language.Dutch:                puddlesDutchPhrase,
		language.French:               puddlesFrenchPhrase,
		language.German:               puddlesGermanPhrase,
		language.Italian:              puddlesItalianPhrase,
		language.Japanese:             puddlesJapanesePhrase,
		language.Korean:               puddlesKoreanPhrase,
		language.LatinAmericanSpanish: puddlesLatinAmericanSpanishPhrase,
		language.Russian:              puddlesRussianPhrase,
		language.SimplifiedChinese:    puddlesSimplifiedChinesePhrase,
		language.Spanish:              puddlesSpanishPhrase,
		language.TraditionalChinese:   puddlesTraditionalChinesePhrase}
)

var (
	Puddles = nook.Villager{
		Character:   puddlesCharacter,
		Personality: personality.Peppy,
		Phrase:      puddlesPhrase}
)
