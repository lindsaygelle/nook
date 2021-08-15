package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	timbraBirthday = nook.Birthday{
		Day:   21,
		Month: time.October}
)

var (
	timbraCode = nook.Code{
		Value: "shp10"}
)

var (
	timbraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Timbra"}

	timbraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sélène"}

	timbraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Timbra"}

	timbraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sélène"}

	timbraGermanName = nook.Name{
		Language: language.German,
		Value:    "Tippsi"}

	timbraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ambra"}

	timbraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "つかさ"}

	timbraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "잔디"}

	timbraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hebra"}

	timbraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тимбра"}

	timbraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿司"}

	timbraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hebra"}

	timbraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿司"}
)

var (
	timbraName = nook.Languages{
		language.AmericanEnglish:      timbraAmericanEnglishName,
		language.CanadianFrench:       timbraCanadianFrenchName,
		language.Dutch:                timbraDutchName,
		language.French:               timbraFrenchName,
		language.German:               timbraGermanName,
		language.Italian:              timbraItalianName,
		language.Japanese:             timbraJapaneseName,
		language.Korean:               timbraKoreanName,
		language.LatinAmericanSpanish: timbraLatinAmericanSpanishName,
		language.Russian:              timbraRussianName,
		language.SimplifiedChinese:    timbraSimplifiedChineseName,
		language.Spanish:              timbraSpanishName,
		language.TraditionalChinese:   timbraTraditionalChineseName}
)

var (
	timbraCharacter = nook.Character{
		Animal:   Sheep,
		Birthday: timbraBirthday,
		Code:     timbraCode,
		Gender:   nook.Female,
		Name:     timbraName}
)

var (
	timbraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pine nut"}

	timbraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "barbiche"}

	timbraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "denappel"}

	timbraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "barbiche"}

	timbraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bwäääh"}

	timbraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pecoramè"}

	timbraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まつわ"}

	timbraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "달려"}

	timbraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "yalovééés"}

	timbraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сосенка"}

	timbraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "等啊"}

	timbraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "yalovéees"}

	timbraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "等啊"}
)

var (
	timbraPhrase = nook.Languages{
		language.AmericanEnglish:      timbraAmericanEnglishName,
		language.CanadianFrench:       timbraCanadianFrenchName,
		language.Dutch:                timbraDutchName,
		language.French:               timbraFrenchName,
		language.German:               timbraGermanName,
		language.Italian:              timbraItalianName,
		language.Japanese:             timbraJapaneseName,
		language.Korean:               timbraKoreanName,
		language.LatinAmericanSpanish: timbraLatinAmericanSpanishName,
		language.Russian:              timbraRussianName,
		language.SimplifiedChinese:    timbraSimplifiedChineseName,
		language.Spanish:              timbraSpanishName,
		language.TraditionalChinese:   timbraTraditionalChineseName}
)

var (
	Timbra = nook.Villager{
		Character:   timbraCharacter,
		Personality: nook.Snooty,
		Phrase:      timbraPhrase}
)
