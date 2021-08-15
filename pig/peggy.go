package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	peggyBirthday = nook.Birthday{
		Day:   23,
		Month: time.May}
)

var (
	peggyCode = nook.Code{
		Value: "pig11"}
)

var (
	peggyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peggy"}

	peggyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rosefleur"}

	peggyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peggyknorrelief"}

	peggyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosefleur"}

	peggyGermanName = nook.Name{
		Language: language.German,
		Value:    "Quiekieglotz"}

	peggyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sallyoink"}

	peggyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちえりぷるる"}

	peggyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "체리아앙"}

	peggyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Peggytrufits"}

	peggyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пеггихрюмпатяга"}

	peggySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "千惠弹弹"}

	peggySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Peggytrufita"}

	peggyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "千惠彈彈"}
)

var (
	peggyName = nook.Languages{
		language.AmericanEnglish:      peggyAmericanEnglishName,
		language.CanadianFrench:       peggyCanadianFrenchName,
		language.Dutch:                peggyDutchName,
		language.French:               peggyFrenchName,
		language.German:               peggyGermanName,
		language.Italian:              peggyItalianName,
		language.Japanese:             peggyJapaneseName,
		language.Korean:               peggyKoreanName,
		language.LatinAmericanSpanish: peggyLatinAmericanSpanishName,
		language.Russian:              peggyRussianName,
		language.SimplifiedChinese:    peggySimplifiedChineseName,
		language.Spanish:              peggySpanishName,
		language.TraditionalChinese:   peggyTraditionalChineseName}
)

var (
	peggyCharacter = nook.Character{
		Animal:   Pig,
		Birthday: peggyBirthday,
		Code:     peggyCode,
		Gender:   nook.Female,
		Name:     peggyName}
)

var (
	peggyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "shweetie"}

	peggyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "fleur"}

	peggyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knorrelief"}

	peggyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "fleur"}

	peggyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "glotz"}

	peggyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oink"}

	peggyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぷるる"}

	peggyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아앙"}

	peggyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "trufits"}

	peggyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюмпатяга"}

	peggySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "弹弹"}

	peggySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trufita"}

	peggyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "彈彈"}
)

var (
	peggyPhrase = nook.Languages{
		language.AmericanEnglish:      peggyAmericanEnglishName,
		language.CanadianFrench:       peggyCanadianFrenchName,
		language.Dutch:                peggyDutchName,
		language.French:               peggyFrenchName,
		language.German:               peggyGermanName,
		language.Italian:              peggyItalianName,
		language.Japanese:             peggyJapaneseName,
		language.Korean:               peggyKoreanName,
		language.LatinAmericanSpanish: peggyLatinAmericanSpanishName,
		language.Russian:              peggyRussianName,
		language.SimplifiedChinese:    peggySimplifiedChineseName,
		language.Spanish:              peggySpanishName,
		language.TraditionalChinese:   peggyTraditionalChineseName}
)

var (
	Peggy = nook.Villager{
		Character:   peggyCharacter,
		Personality: nook.Peppy,
		Phrase:      peggyPhrase}
)
