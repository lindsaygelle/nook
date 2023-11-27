package sheep

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
	// timbraBirthday represents Timbra's birthday.
	timbraBirthday = nook.Birthday{
		Day:   21,
		Month: time.October}
)

var (
	// timbraCode represents Timbra's unique code.
	timbraCode = nook.Code{
		Value: "shp10"}
)

var (
	// timbraAmericanEnglishName represents Timbra's name in American English.
	timbraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Timbra"}

	// timbraCanadianFrenchName represents Timbra's name in Canadian French.
	timbraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sélène"}

	// timbraDutchName represents Timbra's name in Dutch.
	timbraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Timbra"}

	// timbraFrenchName represents Timbra's name in French.
	timbraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sélène"}

	// timbraGermanName represents Timbra's name in German.
	timbraGermanName = nook.Name{
		Language: language.German,
		Value:    "Tippsi"}

	// timbraItalianName represents Timbra's name in Italian.
	timbraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ambra"}

	// timbraJapaneseName represents Timbra's name in Japanese.
	timbraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "つかさ"}

	// timbraKoreanName represents Timbra's name in Korean.
	timbraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "잔디"}

	// timbraLatinAmericanSpanishName represents Timbra's name in Latin American Spanish.
	timbraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hebra"}

	// timbraRussianName represents Timbra's name in Russian.
	timbraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тимбра"}

	// timbraSimplifiedChineseName represents Timbra's name in Simplified Chinese.
	timbraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿司"}

	// timbraSpanishName represents Timbra's name in Spanish.
	timbraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hebra"}

	// timbraTraditionalChineseName represents Timbra's name in Traditional Chinese.
	timbraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿司"}
)

var (
	// timbraName represents Timbra's name in different languages.
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
	// timbraCharacter represents Timbra's character information.
	timbraCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: timbraBirthday,
		Code:     timbraCode,
		Key:      character.Timbra,
		Gender:   gender.Female,
		Name:     timbraName,
		Special:  false}
)

var (
	// timbraAmericanEnglishPhrase represents Timbra's phrase in American English.
	timbraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pine nut"}

	// timbraCanadianFrenchPhrase represents Timbra's phrase in Canadian French.
	timbraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "barbiche"}

	// timbraDutchPhrase represents Timbra's phrase in Dutch.
	timbraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "denappel"}

	// timbraFrenchPhrase represents Timbra's phrase in French.
	timbraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "barbiche"}

	// timbraGermanPhrase represents Timbra's phrase in German.
	timbraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bwäääh"}

	// timbraItalianPhrase represents Timbra's phrase in Italian.
	timbraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pecoramè"}

	// timbraJapanesePhrase represents Timbra's phrase in Japanese.
	timbraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まつわ"}

	// timbraKoreanPhrase represents Timbra's phrase in Korean.
	timbraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "달려"}

	// timbraLatinAmericanSpanishPhrase represents Timbra's phrase in Latin American Spanish.
	timbraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "yalovééés"}

	// timbraRussianPhrase represents Timbra's phrase in Russian.
	timbraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сосенка"}

	// timbraSimplifiedChinesePhrase represents Timbra's phrase in Simplified Chinese.
	timbraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "等啊"}

	// timbraSpanishPhrase represents Timbra's phrase in Spanish.
	timbraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "yalovéees"}

	// timbraTraditionalChinesePhrase represents Timbra's phrase in Traditional Chinese.
	timbraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "等啊"}
)

var (
	// timbraPhrase represents Timbra's phrases in different languages.
	timbraPhrase = nook.Languages{
		language.AmericanEnglish:      timbraAmericanEnglishPhrase,
		language.CanadianFrench:       timbraCanadianFrenchPhrase,
		language.Dutch:                timbraDutchPhrase,
		language.French:               timbraFrenchPhrase,
		language.German:               timbraGermanPhrase,
		language.Italian:              timbraItalianPhrase,
		language.Japanese:             timbraJapanesePhrase,
		language.Korean:               timbraKoreanPhrase,
		language.LatinAmericanSpanish: timbraLatinAmericanSpanishPhrase,
		language.Russian:              timbraRussianPhrase,
		language.SimplifiedChinese:    timbraSimplifiedChinesePhrase,
		language.Spanish:              timbraSpanishPhrase,
		language.TraditionalChinese:   timbraTraditionalChinesePhrase}
)

var (
	// Timbra represents the character Timbra with her complete information.
	Timbra = nook.Villager{
		Character:   timbraCharacter,
		Personality: personality.Snooty,
		Phrase:      timbraPhrase}
)
