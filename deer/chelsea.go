package deer

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	chelseaBirthday = nook.Birthday{
		Day:   18,
		Month: time.January}
)

var (
	chelseaCode = nook.Code{
		Value: "der10"}
)

var (
	chelseaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chelsea"}

	chelseaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chelsea"}

	chelseaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chelsea"}

	chelseaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chelsea"}

	chelseaGermanName = nook.Name{
		Language: language.German,
		Value:    "Chelsea"}

	chelseaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chelsea"}

	chelseaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チェルシー"}

	chelseaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "첼시"}

	chelseaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chelsea"}

	chelseaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Челси"}

	chelseaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雀儿喜"}

	chelseaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chelsea"}

	chelseaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雀兒喜"}
)

var (
	chelseaName = nook.Languages{
		language.AmericanEnglish:      chelseaAmericanEnglishName,
		language.CanadianFrench:       chelseaCanadianFrenchName,
		language.Dutch:                chelseaDutchName,
		language.French:               chelseaFrenchName,
		language.German:               chelseaGermanName,
		language.Italian:              chelseaItalianName,
		language.Japanese:             chelseaJapaneseName,
		language.Korean:               chelseaKoreanName,
		language.LatinAmericanSpanish: chelseaLatinAmericanSpanishName,
		language.Russian:              chelseaRussianName,
		language.SimplifiedChinese:    chelseaSimplifiedChineseName,
		language.Spanish:              chelseaSpanishName,
		language.TraditionalChinese:   chelseaTraditionalChineseName}
)

var (
	chelseaCharacter = nook.Character{
		Animal:   Deer,
		Birthday: chelseaBirthday,
		Code:     chelseaCode,
		Gender:   nook.Female,
		Name:     chelseaName}
)

var (
	chelseaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pound cake"}

	chelseaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bibiche"}

	chelseaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "melodietje"}

	chelseaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bibiche"}

	chelseaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "denkpink"}

	chelseaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "melodì"}

	chelseaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "メロメロ"}

	chelseaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하트하트"}

	chelseaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "melomelo"}

	chelseaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бисквит"}

	chelseaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喜喜"}

	chelseaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "melomelo"}

	chelseaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喜喜"}
)

var (
	chelseaPhrase = nook.Languages{
		language.AmericanEnglish:      chelseaAmericanEnglishName,
		language.CanadianFrench:       chelseaCanadianFrenchName,
		language.Dutch:                chelseaDutchName,
		language.French:               chelseaFrenchName,
		language.German:               chelseaGermanName,
		language.Italian:              chelseaItalianName,
		language.Japanese:             chelseaJapaneseName,
		language.Korean:               chelseaKoreanName,
		language.LatinAmericanSpanish: chelseaLatinAmericanSpanishName,
		language.Russian:              chelseaRussianName,
		language.SimplifiedChinese:    chelseaSimplifiedChineseName,
		language.Spanish:              chelseaSpanishName,
		language.TraditionalChinese:   chelseaTraditionalChineseName}
)

var (
	Chelsea = nook.Villager{
		Character:   chelseaCharacter,
		Personality: nook.Normal,
		Phrase:      chelseaPhrase}
)
