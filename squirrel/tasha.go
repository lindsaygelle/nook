package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	tashaBirthday = nook.Birthday{
		Day:   30,
		Month: time.November}
)

var (
	tashaCode = nook.Code{
		Value: "squ13"}
)

var (
	tashaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tasha"}

	tashaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nadeigeptigri"}

	tashaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tashajaja"}

	tashaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nadeigeptigri"}

	tashaGermanName = nook.Name{
		Language: language.German,
		Value:    "Nataljatjaja"}

	tashaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Teresamunch"}

	tashaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナターシャやるわね"}

	tashaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나타샤굿굿"}

	tashaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Taniabi-buá"}

	tashaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ташаладно"}

	tashaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "娜塔丽做得好"}

	tashaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Taniapincel"}

	tashaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "娜塔麗做得好"}
)

var (
	tashaName = nook.Languages{
		language.AmericanEnglish:      tashaAmericanEnglishName,
		language.CanadianFrench:       tashaCanadianFrenchName,
		language.Dutch:                tashaDutchName,
		language.French:               tashaFrenchName,
		language.German:               tashaGermanName,
		language.Italian:              tashaItalianName,
		language.Japanese:             tashaJapaneseName,
		language.Korean:               tashaKoreanName,
		language.LatinAmericanSpanish: tashaLatinAmericanSpanishName,
		language.Russian:              tashaRussianName,
		language.SimplifiedChinese:    tashaSimplifiedChineseName,
		language.Spanish:              tashaSpanishName,
		language.TraditionalChinese:   tashaTraditionalChineseName}
)

var (
	tashaCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: tashaBirthday,
		Code:     tashaCode,
		Gender:   nook.Female,
		Name:     tashaName}
)

var (
	tashaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nice nice"}

	tashaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ptigri"}

	tashaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jaja"}

	tashaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ptigri"}

	tashaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tjaja"}

	tashaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "munch"}

	tashaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やるわね"}

	tashaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "굿굿"}

	tashaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bi-buá"}

	tashaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ладно"}

	tashaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "做得好"}

	tashaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pincel"}

	tashaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "做得好"}
)

var (
	tashaPhrase = nook.Languages{
		language.AmericanEnglish:      tashaAmericanEnglishName,
		language.CanadianFrench:       tashaCanadianFrenchName,
		language.Dutch:                tashaDutchName,
		language.French:               tashaFrenchName,
		language.German:               tashaGermanName,
		language.Italian:              tashaItalianName,
		language.Japanese:             tashaJapaneseName,
		language.Korean:               tashaKoreanName,
		language.LatinAmericanSpanish: tashaLatinAmericanSpanishName,
		language.Russian:              tashaRussianName,
		language.SimplifiedChinese:    tashaSimplifiedChineseName,
		language.Spanish:              tashaSpanishName,
		language.TraditionalChinese:   tashaTraditionalChineseName}
)

var (
	Tasha = nook.Villager{
		Character:   tashaCharacter,
		Personality: nook.Snooty,
		Phrase:      tashaPhrase}
)
