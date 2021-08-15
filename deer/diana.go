package deer

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	dianaBirthday = nook.Birthday{
		Day:   4,
		Month: time.January}
)

var (
	dianaCode = nook.Code{
		Value: "der08"}
)

var (
	dianaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Diana"}

	dianaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Didibichouchou"}

	dianaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dianajoh"}

	dianaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Didibichouchou"}

	dianaGermanName = nook.Name{
		Language: language.German,
		Value:    "Vronikitz"}

	dianaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dianamamy"}

	dianaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナタリーでしょ"}

	dianaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나탈리라니까"}

	dianaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bambinanonuá"}

	dianaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дианаеще бы"}

	dianaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "倪家莉是吧"}

	dianaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bambinabraaavo"}

	dianaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "倪家莉是吧"}
)

var (
	dianaName = nook.Languages{
		language.AmericanEnglish:      dianaAmericanEnglishName,
		language.CanadianFrench:       dianaCanadianFrenchName,
		language.Dutch:                dianaDutchName,
		language.French:               dianaFrenchName,
		language.German:               dianaGermanName,
		language.Italian:              dianaItalianName,
		language.Japanese:             dianaJapaneseName,
		language.Korean:               dianaKoreanName,
		language.LatinAmericanSpanish: dianaLatinAmericanSpanishName,
		language.Russian:              dianaRussianName,
		language.SimplifiedChinese:    dianaSimplifiedChineseName,
		language.Spanish:              dianaSpanishName,
		language.TraditionalChinese:   dianaTraditionalChineseName}
)

var (
	dianaCharacter = nook.Character{
		Animal:   Deer,
		Birthday: dianaBirthday,
		Code:     dianaCode,
		Gender:   nook.Female,
		Name:     dianaName}
)

var (
	dianaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "no doy"}

	dianaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bichouchou"}

	dianaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "joh"}

	dianaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bichouchou"}

	dianaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kitz"}

	dianaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mamy"}

	dianaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でしょ"}

	dianaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라니까"}

	dianaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nonuá"}

	dianaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "еще бы"}

	dianaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是吧"}

	dianaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "braaavo"}

	dianaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是吧"}
)

var (
	dianaPhrase = nook.Languages{
		language.AmericanEnglish:      dianaAmericanEnglishName,
		language.CanadianFrench:       dianaCanadianFrenchName,
		language.Dutch:                dianaDutchName,
		language.French:               dianaFrenchName,
		language.German:               dianaGermanName,
		language.Italian:              dianaItalianName,
		language.Japanese:             dianaJapaneseName,
		language.Korean:               dianaKoreanName,
		language.LatinAmericanSpanish: dianaLatinAmericanSpanishName,
		language.Russian:              dianaRussianName,
		language.SimplifiedChinese:    dianaSimplifiedChineseName,
		language.Spanish:              dianaSpanishName,
		language.TraditionalChinese:   dianaTraditionalChineseName}
)

var (
	Diana = nook.Villager{
		Character:   dianaCharacter,
		Personality: nook.Snooty,
		Phrase:      dianaPhrase}
)
