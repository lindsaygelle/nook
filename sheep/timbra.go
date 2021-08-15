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
		Value:    "Sélènebarbiche"}

	timbraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Timbradenappel"}

	timbraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sélènebarbiche"}

	timbraGermanName = nook.Name{
		Language: language.German,
		Value:    "Tippsibwäääh"}

	timbraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ambrapecoramè"}

	timbraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "つかさまつわ"}

	timbraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "잔디달려"}

	timbraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hebrayalovééés"}

	timbraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тимбрасосенка"}

	timbraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿司等啊"}

	timbraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hebrayalovéees"}

	timbraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿司等啊"}
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
