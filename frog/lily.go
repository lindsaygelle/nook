package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	lilyBirthday = nook.Birthday{
		Day:   4,
		Month: time.February}
)

var (
	lilyCode = nook.Code{
		Value: "flg00"}
)

var (
	lilyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lily"}

	lilyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rainazzrrbbitt"}

	lilyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lilypadje"}

	lilyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rainawatich'"}

	lilyGermanName = nook.Name{
		Language: language.German,
		Value:    "Lilianekröti"}

	lilyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gigliolagragragra"}

	lilyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイニーでちゅ"}

	lilyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레이니그래요"}

	lilyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Liliriiibit"}

	lilyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лиликва"}

	lilySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雷妮雨雨"}

	lilySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lilimosquita"}

	lilyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雷妮雨雨"}
)

var (
	lilyName = nook.Languages{
		language.AmericanEnglish:      lilyAmericanEnglishName,
		language.CanadianFrench:       lilyCanadianFrenchName,
		language.Dutch:                lilyDutchName,
		language.French:               lilyFrenchName,
		language.German:               lilyGermanName,
		language.Italian:              lilyItalianName,
		language.Japanese:             lilyJapaneseName,
		language.Korean:               lilyKoreanName,
		language.LatinAmericanSpanish: lilyLatinAmericanSpanishName,
		language.Russian:              lilyRussianName,
		language.SimplifiedChinese:    lilySimplifiedChineseName,
		language.Spanish:              lilySpanishName,
		language.TraditionalChinese:   lilyTraditionalChineseName}
)

var (
	lilyCharacter = nook.Character{
		Animal:   Frog,
		Birthday: lilyBirthday,
		Code:     lilyCode,
		Gender:   nook.Female,
		Name:     lilyName}
)

var (
	lilyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zzrrbbittoady"}

	lilyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "zzrrbbitt"}

	lilyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "padje"}

	lilyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "watich'"}

	lilyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kröti"}

	lilyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gragragra"}

	lilyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でちゅ"}

	lilyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그래요"}

	lilyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "riiibit"}

	lilyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ква"}

	lilySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雨雨"}

	lilySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mosquita"}

	lilyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雨雨"}
)

var (
	lilyPhrase = nook.Languages{
		language.AmericanEnglish:      lilyAmericanEnglishName,
		language.CanadianFrench:       lilyCanadianFrenchName,
		language.Dutch:                lilyDutchName,
		language.French:               lilyFrenchName,
		language.German:               lilyGermanName,
		language.Italian:              lilyItalianName,
		language.Japanese:             lilyJapaneseName,
		language.Korean:               lilyKoreanName,
		language.LatinAmericanSpanish: lilyLatinAmericanSpanishName,
		language.Russian:              lilyRussianName,
		language.SimplifiedChinese:    lilySimplifiedChineseName,
		language.Spanish:              lilySpanishName,
		language.TraditionalChinese:   lilyTraditionalChineseName}
)

var (
	Lily = nook.Villager{
		Character:   lilyCharacter,
		Personality: nook.Normal,
		Phrase:      lilyPhrase}
)
