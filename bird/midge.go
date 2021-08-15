package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	midgeBirthday = nook.Birthday{
		Day:   12,
		Month: time.March}
)

var (
	midgeCode = nook.Code{
		Value: "brd08"}
)

var (
	midgeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Midge"}

	midgeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Michèleloulou"}

	midgeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Midgekwetter"}

	midgeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Michèleloulou"}

	midgeGermanName = nook.Name{
		Language: language.German,
		Value:    "Annatwitiriti"}

	midgeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Minùfirulì"}

	midgeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "うずまきキョン"}

	midgeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핑글이핑그르르"}

	midgeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Palomaarrurrú"}

	midgeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Миджтви-тви"}

	midgeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小酒窝脸红"}

	midgeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Palomaarrurrú"}

	midgeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小酒窩臉紅"}
)

var (
	midgeName = nook.Languages{
		language.AmericanEnglish:      midgeAmericanEnglishName,
		language.CanadianFrench:       midgeCanadianFrenchName,
		language.Dutch:                midgeDutchName,
		language.French:               midgeFrenchName,
		language.German:               midgeGermanName,
		language.Italian:              midgeItalianName,
		language.Japanese:             midgeJapaneseName,
		language.Korean:               midgeKoreanName,
		language.LatinAmericanSpanish: midgeLatinAmericanSpanishName,
		language.Russian:              midgeRussianName,
		language.SimplifiedChinese:    midgeSimplifiedChineseName,
		language.Spanish:              midgeSpanishName,
		language.TraditionalChinese:   midgeTraditionalChineseName}
)

var (
	midgeCharacter = nook.Character{
		Animal:   Bird,
		Birthday: midgeBirthday,
		Code:     midgeCode,
		Gender:   nook.Female,
		Name:     midgeName}
)

var (
	midgeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tweedledee"}

	midgeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "loulou"}

	midgeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwetter"}

	midgeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "loulou"}

	midgeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "twitiriti"}

	midgeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "firulì"}

	midgeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キョン"}

	midgeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "핑그르르"}

	midgeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "arrurrú"}

	midgeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тви-тви"}

	midgeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "脸红"}

	midgeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "arrurrú"}

	midgeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "臉紅"}
)

var (
	midgePhrase = nook.Languages{
		language.AmericanEnglish:      midgeAmericanEnglishName,
		language.CanadianFrench:       midgeCanadianFrenchName,
		language.Dutch:                midgeDutchName,
		language.French:               midgeFrenchName,
		language.German:               midgeGermanName,
		language.Italian:              midgeItalianName,
		language.Japanese:             midgeJapaneseName,
		language.Korean:               midgeKoreanName,
		language.LatinAmericanSpanish: midgeLatinAmericanSpanishName,
		language.Russian:              midgeRussianName,
		language.SimplifiedChinese:    midgeSimplifiedChineseName,
		language.Spanish:              midgeSpanishName,
		language.TraditionalChinese:   midgeTraditionalChineseName}
)

var (
	Midge = nook.Villager{
		Character:   midgeCharacter,
		Personality: nook.Normal,
		Phrase:      midgePhrase}
)
