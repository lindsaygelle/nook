package frog

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
		Value:    "Bajoue"}

	huckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Huck"}

	huckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bajoue"}

	huckGermanName = nook.Name{
		Language: language.German,
		Value:    "Knuth"}

	huckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Berto"}

	huckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ストロー"}

	huckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스트로"}

	huckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ricardo"}

	huckRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хак"}

	huckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吸管"}

	huckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ricardo"}

	huckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吸管"}
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
		Animal:   animal.Frog,
		Birthday: huckBirthday,
		Code:     huckCode,
		Key:      character.Huck,
		Gender:   gender.Male,
		Name:     huckName,
		Special:  false}
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
		language.AmericanEnglish:      huckAmericanEnglishPhrase,
		language.CanadianFrench:       huckCanadianFrenchPhrase,
		language.Dutch:                huckDutchPhrase,
		language.French:               huckFrenchPhrase,
		language.German:               huckGermanPhrase,
		language.Italian:              huckItalianPhrase,
		language.Japanese:             huckJapanesePhrase,
		language.Korean:               huckKoreanPhrase,
		language.LatinAmericanSpanish: huckLatinAmericanSpanishPhrase,
		language.Russian:              huckRussianPhrase,
		language.SimplifiedChinese:    huckSimplifiedChinesePhrase,
		language.Spanish:              huckSpanishPhrase,
		language.TraditionalChinese:   huckTraditionalChinesePhrase}
)

var (
	Huck = nook.Villager{
		Character:   huckCharacter,
		Personality: personality.Smug,
		Phrase:      huckPhrase}
)
