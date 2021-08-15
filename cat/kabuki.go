package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Kabukimia-ouh"}

	kabukiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kabukimie-auw"}

	kabukiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kabukimia-ouh"}

	kabukiGermanName = nook.Name{
		Language: language.German,
		Value:    "Kabukimiiiau"}

	kabukiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kabukimeooo-OH"}

	kabukiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "かぶきちぃよぉー"}

	kabukiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "가북희꺄"}

	kabukiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kabukiimikimono"}

	kabukiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кабукимя-я-яу"}

	kabukiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "戈伍纪咦唷"}

	kabukiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kabukiimikimono"}

	kabukiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "戈伍紀咦唷"}
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
		Animal:   Cat,
		Birthday: kabukiBirthday,
		Code:     kabukiCode,
		Gender:   nook.Male,
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
	Kabuki = nook.Villager{
		Character:   kabukiCharacter,
		Personality: nook.Cranky,
		Phrase:      kabukiPhrase}
)
