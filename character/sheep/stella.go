package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	stellaBirthday = nook.Birthday{
		Day:   9,
		Month: time.April}
)

var (
	stellaCode = nook.Code{
		Value: "shp03"}
)

var (
	stellaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Stella"}

	stellaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bigoudi"}

	stellaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Stella"}

	stellaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bigoudi"}

	stellaGermanName = nook.Name{
		Language: language.German,
		Value:    "Stella"}

	stellaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Merina"}

	stellaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アクリル"}

	stellaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아크릴"}

	stellaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Merina"}

	stellaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стелла"}

	stellaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毡毡"}

	stellaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Merina"}

	stellaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "氈氈"}
)

var (
	stellaName = nook.Languages{
		language.AmericanEnglish:      stellaAmericanEnglishName,
		language.CanadianFrench:       stellaCanadianFrenchName,
		language.Dutch:                stellaDutchName,
		language.French:               stellaFrenchName,
		language.German:               stellaGermanName,
		language.Italian:              stellaItalianName,
		language.Japanese:             stellaJapaneseName,
		language.Korean:               stellaKoreanName,
		language.LatinAmericanSpanish: stellaLatinAmericanSpanishName,
		language.Russian:              stellaRussianName,
		language.SimplifiedChinese:    stellaSimplifiedChineseName,
		language.Spanish:              stellaSpanishName,
		language.TraditionalChinese:   stellaTraditionalChineseName}
)

var (
	stellaCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: stellaBirthday,
		Code:     stellaCode,
		Gender:   gender.Female,
		Name:     stellaName}
)

var (
	stellaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "baa-dabing"}

	stellaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon agneau"}

	stellaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mèèèh"}

	stellaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon agneau"}

	stellaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "badabing"}

	stellaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beeene"}

	stellaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウール"}

	stellaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "울"}

	stellaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "baa-bii"}

	stellaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бэ-дабинь"}

	stellaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "羊毛"}

	stellaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "no veeeas"}

	stellaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羊毛"}
)

var (
	stellaPhrase = nook.Languages{
		language.AmericanEnglish:      stellaAmericanEnglishName,
		language.CanadianFrench:       stellaCanadianFrenchName,
		language.Dutch:                stellaDutchName,
		language.French:               stellaFrenchName,
		language.German:               stellaGermanName,
		language.Italian:              stellaItalianName,
		language.Japanese:             stellaJapaneseName,
		language.Korean:               stellaKoreanName,
		language.LatinAmericanSpanish: stellaLatinAmericanSpanishName,
		language.Russian:              stellaRussianName,
		language.SimplifiedChinese:    stellaSimplifiedChineseName,
		language.Spanish:              stellaSpanishName,
		language.TraditionalChinese:   stellaTraditionalChineseName}
)

var (
	Stella = nook.Villager{
		Character:   stellaCharacter,
		Personality: personality.Normal,
		Phrase:      stellaPhrase}
)
