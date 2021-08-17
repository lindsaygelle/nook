package cat

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
	kabukiBirthday = nook.Birthday{
		Day:   29,
		Month: time.November}
)

var (
	kabukiCode = nook.Code{
		Value: "cat09"}
)

var (
	kabukiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kabuki"}

	kabukiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kabuki"}

	kabukiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kabuki"}

	kabukiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kabuki"}

	kabukiGermanName = nook.Name{
		Language: language.German,
		Value:    "Kabuki"}

	kabukiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kabuki"}

	kabukiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "かぶきち"}

	kabukiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "가북희"}

	kabukiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kabuki"}

	kabukiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кабуки"}

	kabukiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "戈伍纪"}

	kabukiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kabuki"}

	kabukiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "戈伍紀"}
)

var (
	kabukiName = nook.Languages{
		language.AmericanEnglish:      kabukiAmericanEnglishName,
		language.CanadianFrench:       kabukiCanadianFrenchName,
		language.Dutch:                kabukiDutchName,
		language.French:               kabukiFrenchName,
		language.German:               kabukiGermanName,
		language.Italian:              kabukiItalianName,
		language.Japanese:             kabukiJapaneseName,
		language.Korean:               kabukiKoreanName,
		language.LatinAmericanSpanish: kabukiLatinAmericanSpanishName,
		language.Russian:              kabukiRussianName,
		language.SimplifiedChinese:    kabukiSimplifiedChineseName,
		language.Spanish:              kabukiSpanishName,
		language.TraditionalChinese:   kabukiTraditionalChineseName}
)

var (
	kabukiCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: kabukiBirthday,
		Code:     kabukiCode,
		Key:      character.Kabuki,
		Gender:   gender.Male,
		Name:     kabukiName}
)

var (
	kabukiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "meooo-OH"}

	kabukiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mia-ouh"}

	kabukiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mie-auw"}

	kabukiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mia-ouh"}

	kabukiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "miiiau"}

	kabukiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "meooo-OH"}

	kabukiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぃよぉー"}

	kabukiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꺄"}

	kabukiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "imikimono"}

	kabukiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мя-я-яу"}

	kabukiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咦唷"}

	kabukiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "imikimono"}

	kabukiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咦唷"}
)

var (
	kabukiPhrase = nook.Languages{
		language.AmericanEnglish:      kabukiAmericanEnglishPhrase,
		language.CanadianFrench:       kabukiCanadianFrenchPhrase,
		language.Dutch:                kabukiDutchPhrase,
		language.French:               kabukiFrenchPhrase,
		language.German:               kabukiGermanPhrase,
		language.Italian:              kabukiItalianPhrase,
		language.Japanese:             kabukiJapanesePhrase,
		language.Korean:               kabukiKoreanPhrase,
		language.LatinAmericanSpanish: kabukiLatinAmericanSpanishPhrase,
		language.Russian:              kabukiRussianPhrase,
		language.SimplifiedChinese:    kabukiSimplifiedChinesePhrase,
		language.Spanish:              kabukiSpanishPhrase,
		language.TraditionalChinese:   kabukiTraditionalChinesePhrase}
)

var (
	Kabuki = nook.Villager{
		Character:   kabukiCharacter,
		Personality: personality.Cranky,
		Phrase:      kabukiPhrase}
)
