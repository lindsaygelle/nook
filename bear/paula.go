package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	paulaBirthday = nook.Birthday{
		Day:   22,
		Month: time.March}
)

var (
	paulaCode = nook.Code{
		Value: "bea10"}
)

var (
	paulaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Paula"}

	paulaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Wendytourlou"}

	paulaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Paulajodelie"}

	paulaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Wendyléchouille"}

	paulaGermanName = nook.Name{
		Language: language.German,
		Value:    "Paulabralala"}

	paulaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Brunabah"}

	paulaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイチェルヤッホー"}

	paulaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레이첼야호야호"}

	paulaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Liziborobó"}

	paulaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Паулайодль"}

	paulaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "瑞秋耶呼"}

	paulaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Liziborobó"}

	paulaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑞秋耶呼"}
)

var (
	paulaName = nook.Languages{
		language.AmericanEnglish:      paulaAmericanEnglishName,
		language.CanadianFrench:       paulaCanadianFrenchName,
		language.Dutch:                paulaDutchName,
		language.French:               paulaFrenchName,
		language.German:               paulaGermanName,
		language.Italian:              paulaItalianName,
		language.Japanese:             paulaJapaneseName,
		language.Korean:               paulaKoreanName,
		language.LatinAmericanSpanish: paulaLatinAmericanSpanishName,
		language.Russian:              paulaRussianName,
		language.SimplifiedChinese:    paulaSimplifiedChineseName,
		language.Spanish:              paulaSpanishName,
		language.TraditionalChinese:   paulaTraditionalChineseName}
)

var (
	paulaCharacter = nook.Character{
		Animal:   Bear,
		Birthday: paulaBirthday,
		Code:     paulaCode,
		Gender:   nook.Female,
		Name:     paulaName}
)

var (
	paulaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yodelay"}

	paulaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tourlou"}

	paulaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jodelie"}

	paulaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "léchouille"}

	paulaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bralala"}

	paulaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bah"}

	paulaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヤッホー"}

	paulaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "야호야호"}

	paulaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "borobó"}

	paulaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "йодль"}

	paulaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "耶呼"}

	paulaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "borobó"}

	paulaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "耶呼"}
)

var (
	paulaPhrase = nook.Languages{
		language.AmericanEnglish:      paulaAmericanEnglishName,
		language.CanadianFrench:       paulaCanadianFrenchName,
		language.Dutch:                paulaDutchName,
		language.French:               paulaFrenchName,
		language.German:               paulaGermanName,
		language.Italian:              paulaItalianName,
		language.Japanese:             paulaJapaneseName,
		language.Korean:               paulaKoreanName,
		language.LatinAmericanSpanish: paulaLatinAmericanSpanishName,
		language.Russian:              paulaRussianName,
		language.SimplifiedChinese:    paulaSimplifiedChineseName,
		language.Spanish:              paulaSpanishName,
		language.TraditionalChinese:   paulaTraditionalChineseName}
)

var (
	Paula = nook.Villager{
		Character:   paulaCharacter,
		Personality: nook.BigSister,
		Phrase:      paulaPhrase}
)
