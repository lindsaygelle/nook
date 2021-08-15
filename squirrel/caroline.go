package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	carolineBirthday = nook.Birthday{
		Day:   15,
		Month: time.July}
)

var (
	carolineCode = nook.Code{
		Value: "squ06"}
)

var (
	carolineAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Caroline"}

	carolineCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Isabellehoulaaaa"}

	carolineDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Carolinehoelaaa"}

	carolineFrenchName = nook.Name{
		Language: language.French,
		Value:    "Isabellehoulaaaa"}

	carolineGermanName = nook.Name{
		Language: language.German,
		Value:    "Ricardahulaaaa"}

	carolineItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carolinaehilaaà"}

	carolineJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャロラインユー"}

	carolineKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캐롤라인휴"}

	carolineLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marilóulaaaa"}

	carolineRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Каролинауа-а-а"}

	carolineSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贾萝琳你"}

	carolineSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marilóulaaaa"}

	carolineTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "賈蘿琳你"}
)

var (
	carolineName = nook.Languages{
		language.AmericanEnglish:      carolineAmericanEnglishName,
		language.CanadianFrench:       carolineCanadianFrenchName,
		language.Dutch:                carolineDutchName,
		language.French:               carolineFrenchName,
		language.German:               carolineGermanName,
		language.Italian:              carolineItalianName,
		language.Japanese:             carolineJapaneseName,
		language.Korean:               carolineKoreanName,
		language.LatinAmericanSpanish: carolineLatinAmericanSpanishName,
		language.Russian:              carolineRussianName,
		language.SimplifiedChinese:    carolineSimplifiedChineseName,
		language.Spanish:              carolineSpanishName,
		language.TraditionalChinese:   carolineTraditionalChineseName}
)

var (
	carolineCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: carolineBirthday,
		Code:     carolineCode,
		Gender:   nook.Female,
		Name:     carolineName}
)

var (
	carolineAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hulaaaa"}

	carolineCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "houlaaaa"}

	carolineDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hoelaaa"}

	carolineFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "houlaaaa"}

	carolineGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hulaaaa"}

	carolineItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ehilaaà"}

	carolineJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ユー"}

	carolineKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "휴"}

	carolineLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ulaaaa"}

	carolineRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "уа-а-а"}

	carolineSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "你"}

	carolineSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ulaaaa"}

	carolineTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "你"}
)

var (
	carolinePhrase = nook.Languages{
		language.AmericanEnglish:      carolineAmericanEnglishName,
		language.CanadianFrench:       carolineCanadianFrenchName,
		language.Dutch:                carolineDutchName,
		language.French:               carolineFrenchName,
		language.German:               carolineGermanName,
		language.Italian:              carolineItalianName,
		language.Japanese:             carolineJapaneseName,
		language.Korean:               carolineKoreanName,
		language.LatinAmericanSpanish: carolineLatinAmericanSpanishName,
		language.Russian:              carolineRussianName,
		language.SimplifiedChinese:    carolineSimplifiedChineseName,
		language.Spanish:              carolineSpanishName,
		language.TraditionalChinese:   carolineTraditionalChineseName}
)

var (
	Caroline = nook.Villager{
		Character:   carolineCharacter,
		Personality: nook.Normal,
		Phrase:      carolinePhrase}
)
