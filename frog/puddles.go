package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Rénatasplufff"}

	puddlesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Puddlesplons"}

	puddlesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rénatasplufff"}

	puddlesGermanName = nook.Name{
		Language: language.German,
		Value:    "Neleplatschi"}

	puddlesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Graziasplash"}

	puddlesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チョキっちゃ"}

	puddlesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "가위그랬쪄"}

	puddlesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nenúfaresplís"}

	puddlesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ренатаплюх"}

	puddlesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "剪刀石头"}

	puddlesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nenúfaresplís"}

	puddlesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "剪刀石頭"}
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
		Animal:   Frog,
		Birthday: puddlesBirthday,
		Code:     puddlesCode,
		Gender:   nook.Female,
		Name:     puddlesName}
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
	Puddles = nook.Villager{
		Character:   puddlesCharacter,
		Personality: nook.Peppy,
		Phrase:      puddlesPhrase}
)
