package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	huckBirthday = nook.Birthday{
		Day:   9,
		Month: time.July}
)

var (
	huckCode = nook.Code{
		Value: "flg11"}
)

var (
	huckAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Huck"}

	huckCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bajouenouille"}

	huckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Huckhophop"}

	huckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bajouenouille"}

	huckGermanName = nook.Name{
		Language: language.German,
		Value:    "Knuthhops"}

	huckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bertohophop"}

	huckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ストローとかー"}

	huckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스트로싸우자"}

	huckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ricardoajop"}

	huckRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хакскок"}

	huckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吸管喝喝"}

	huckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ricardogrillito"}

	huckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吸管喝喝"}
)

var (
	huckName = nook.Languages{
		language.AmericanEnglish:      huckAmericanEnglishName,
		language.CanadianFrench:       huckCanadianFrenchName,
		language.Dutch:                huckDutchName,
		language.French:               huckFrenchName,
		language.German:               huckGermanName,
		language.Italian:              huckItalianName,
		language.Japanese:             huckJapaneseName,
		language.Korean:               huckKoreanName,
		language.LatinAmericanSpanish: huckLatinAmericanSpanishName,
		language.Russian:              huckRussianName,
		language.SimplifiedChinese:    huckSimplifiedChineseName,
		language.Spanish:              huckSpanishName,
		language.TraditionalChinese:   huckTraditionalChineseName}
)

var (
	huckCharacter = nook.Character{
		Animal:   Frog,
		Birthday: huckBirthday,
		Code:     huckCode,
		Gender:   nook.Male,
		Name:     huckName}
)

var (
	huckAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hopper"}

	huckCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nouille"}

	huckDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hophop"}

	huckFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nouille"}

	huckGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hops"}

	huckItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hophop"}

	huckJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とかー"}

	huckKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "싸우자"}

	huckLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ajop"}

	huckRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "скок"}

	huckSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喝喝"}

	huckSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grillito"}

	huckTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喝喝"}
)

var (
	huckPhrase = nook.Languages{
		language.AmericanEnglish:      huckAmericanEnglishName,
		language.CanadianFrench:       huckCanadianFrenchName,
		language.Dutch:                huckDutchName,
		language.French:               huckFrenchName,
		language.German:               huckGermanName,
		language.Italian:              huckItalianName,
		language.Japanese:             huckJapaneseName,
		language.Korean:               huckKoreanName,
		language.LatinAmericanSpanish: huckLatinAmericanSpanishName,
		language.Russian:              huckRussianName,
		language.SimplifiedChinese:    huckSimplifiedChineseName,
		language.Spanish:              huckSpanishName,
		language.TraditionalChinese:   huckTraditionalChineseName}
)

var (
	Huck = nook.Villager{
		Character:   huckCharacter,
		Personality: nook.Smug,
		Phrase:      huckPhrase}
)
