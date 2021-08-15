package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	violetBirthday = nook.Birthday{
		Day:   1,
		Month: time.September}
)

var (
	violetCode = nook.Code{
		Value: "gor07"}
)

var (
	violetAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Violet"}

	violetCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gaëllegong"}

	violetDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Violetlieverd"}

	violetFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gaëllegong"}

	violetGermanName = nook.Name{
		Language: language.German,
		Value:    "Kongaauauauu"}

	violetItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kongatump tump"}

	violetJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ウズメアイヤ"}

	violetKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "줌마에그머닛"}

	violetLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Magentauh-uh"}

	violetRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Конгасластена"}

	violetSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吴紫眉唉呀"}

	violetSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Magentacanapé"}

	violetTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吳紫眉唉呀"}
)

var (
	violetName = nook.Languages{
		language.AmericanEnglish:      violetAmericanEnglishName,
		language.CanadianFrench:       violetCanadianFrenchName,
		language.Dutch:                violetDutchName,
		language.French:               violetFrenchName,
		language.German:               violetGermanName,
		language.Italian:              violetItalianName,
		language.Japanese:             violetJapaneseName,
		language.Korean:               violetKoreanName,
		language.LatinAmericanSpanish: violetLatinAmericanSpanishName,
		language.Russian:              violetRussianName,
		language.SimplifiedChinese:    violetSimplifiedChineseName,
		language.Spanish:              violetSpanishName,
		language.TraditionalChinese:   violetTraditionalChineseName}
)

var (
	violetCharacter = nook.Character{
		Animal:   Gorilla,
		Birthday: violetBirthday,
		Code:     violetCode,
		Gender:   nook.Female,
		Name:     violetName}
)

var (
	violetAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "faboomsweetie"}

	violetCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gong"}

	violetDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "lieverd"}

	violetFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gong"}

	violetGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "auauauu"}

	violetItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tump tump"}

	violetJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アイヤ"}

	violetKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "에그머닛"}

	violetLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uh-uh"}

	violetRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сластена"}

	violetSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唉呀"}

	violetSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "canapé"}

	violetTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唉呀"}
)

var (
	violetPhrase = nook.Languages{
		language.AmericanEnglish:      violetAmericanEnglishName,
		language.CanadianFrench:       violetCanadianFrenchName,
		language.Dutch:                violetDutchName,
		language.French:               violetFrenchName,
		language.German:               violetGermanName,
		language.Italian:              violetItalianName,
		language.Japanese:             violetJapaneseName,
		language.Korean:               violetKoreanName,
		language.LatinAmericanSpanish: violetLatinAmericanSpanishName,
		language.Russian:              violetRussianName,
		language.SimplifiedChinese:    violetSimplifiedChineseName,
		language.Spanish:              violetSpanishName,
		language.TraditionalChinese:   violetTraditionalChineseName}
)

var (
	Violet = nook.Villager{
		Character:   violetCharacter,
		Personality: nook.Snooty,
		Phrase:      violetPhrase}
)
