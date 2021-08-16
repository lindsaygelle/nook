package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Gaëlle"}

	violetDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Violet"}

	violetFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gaëlle"}

	violetGermanName = nook.Name{
		Language: language.German,
		Value:    "Konga"}

	violetItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Konga"}

	violetJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ウズメ"}

	violetKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "줌마"}

	violetLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Magenta"}

	violetRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Конга"}

	violetSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吴紫眉"}

	violetSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Magenta"}

	violetTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吳紫眉"}
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
		Animal:   animal.Gorilla,
		Birthday: violetBirthday,
		Code:     violetCode,
		Gender:   gender.Female,
		Name:     violetName}
)

var (
	violetAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "faboom"}

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
		Personality: personality.Snooty,
		Phrase:      violetPhrase}
)
