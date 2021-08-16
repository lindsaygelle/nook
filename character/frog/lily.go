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
		Value:    "Raina"}

	lilyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lily"}

	lilyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Raina"}

	lilyGermanName = nook.Name{
		Language: language.German,
		Value:    "Liliane"}

	lilyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gigliola"}

	lilyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイニー"}

	lilyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레이니"}

	lilyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lili"}

	lilyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лили"}

	lilySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雷妮"}

	lilySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lili"}

	lilyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雷妮"}
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
		Animal:   animal.Frog,
		Birthday: lilyBirthday,
		Code:     lilyCode,
		Key:      character.Lily,
		Gender:   gender.Female,
		Name:     lilyName}
)

var (
	lilyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zzrrbbit"}

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
		Personality: personality.Normal,
		Phrase:      lilyPhrase}
)
