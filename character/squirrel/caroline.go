package squirrel

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
		Value:    "Isabelle"}

	carolineDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Caroline"}

	carolineFrenchName = nook.Name{
		Language: language.French,
		Value:    "Isabelle"}

	carolineGermanName = nook.Name{
		Language: language.German,
		Value:    "Ricarda"}

	carolineItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carolina"}

	carolineJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャロライン"}

	carolineKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캐롤라인"}

	carolineLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mariló"}

	carolineRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Каролина"}

	carolineSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贾萝琳"}

	carolineSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mariló"}

	carolineTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "賈蘿琳"}
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
		Animal:   animal.Squirrel,
		Birthday: carolineBirthday,
		Code:     carolineCode,
		Key:      character.Caroline,
		Gender:   gender.Female,
		Name:     carolineName,
		Special:  false}
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
		language.AmericanEnglish:      carolineAmericanEnglishPhrase,
		language.CanadianFrench:       carolineCanadianFrenchPhrase,
		language.Dutch:                carolineDutchPhrase,
		language.French:               carolineFrenchPhrase,
		language.German:               carolineGermanPhrase,
		language.Italian:              carolineItalianPhrase,
		language.Japanese:             carolineJapanesePhrase,
		language.Korean:               carolineKoreanPhrase,
		language.LatinAmericanSpanish: carolineLatinAmericanSpanishPhrase,
		language.Russian:              carolineRussianPhrase,
		language.SimplifiedChinese:    carolineSimplifiedChinesePhrase,
		language.Spanish:              carolineSpanishPhrase,
		language.TraditionalChinese:   carolineTraditionalChinesePhrase}
)

var (
	Caroline = nook.Villager{
		Character:   carolineCharacter,
		Personality: personality.Normal,
		Phrase:      carolinePhrase}
)
