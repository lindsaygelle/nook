package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	divaBirthday = nook.Birthday{
		Day:   2,
		Month: time.October}
)

var (
	divaCode = nook.Code{
		Value: "flg18"}
)

var (
	divaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Diva"}

	divaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Violettecuicuisse"}

	divaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Divajeweetwel"}

	divaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Violettecuicuisse"}

	divaGermanName = nook.Name{
		Language: language.German,
		Value:    "Dörteplatsch"}

	divaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Violacrabum"}

	divaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイーダハーン"}

	divaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이다흐응"}

	divaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Moraniacrocró"}

	divaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дивазнаешь"}

	divaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小睫蛤"}

	divaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Moraniacrocró"}

	divaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小睫蛤"}
)

var (
	divaName = nook.Languages{
		language.AmericanEnglish:      divaAmericanEnglishName,
		language.CanadianFrench:       divaCanadianFrenchName,
		language.Dutch:                divaDutchName,
		language.French:               divaFrenchName,
		language.German:               divaGermanName,
		language.Italian:              divaItalianName,
		language.Japanese:             divaJapaneseName,
		language.Korean:               divaKoreanName,
		language.LatinAmericanSpanish: divaLatinAmericanSpanishName,
		language.Russian:              divaRussianName,
		language.SimplifiedChinese:    divaSimplifiedChineseName,
		language.Spanish:              divaSpanishName,
		language.TraditionalChinese:   divaTraditionalChineseName}
)

var (
	divaCharacter = nook.Character{
		Animal:   Frog,
		Birthday: divaBirthday,
		Code:     divaCode,
		Gender:   nook.Female,
		Name:     divaName}
)

var (
	divaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ya know"}

	divaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cuicuisse"}

	divaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jeweetwel"}

	divaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cuicuisse"}

	divaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "platsch"}

	divaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "crabum"}

	divaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ハーン"}

	divaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "흐응"}

	divaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "crocró"}

	divaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "знаешь"}

	divaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蛤"}

	divaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "crocró"}

	divaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蛤"}
)

var (
	divaPhrase = nook.Languages{
		language.AmericanEnglish:      divaAmericanEnglishName,
		language.CanadianFrench:       divaCanadianFrenchName,
		language.Dutch:                divaDutchName,
		language.French:               divaFrenchName,
		language.German:               divaGermanName,
		language.Italian:              divaItalianName,
		language.Japanese:             divaJapaneseName,
		language.Korean:               divaKoreanName,
		language.LatinAmericanSpanish: divaLatinAmericanSpanishName,
		language.Russian:              divaRussianName,
		language.SimplifiedChinese:    divaSimplifiedChineseName,
		language.Spanish:              divaSpanishName,
		language.TraditionalChinese:   divaTraditionalChineseName}
)

var (
	Diva = nook.Villager{
		Character:   divaCharacter,
		Personality: nook.BigSister,
		Phrase:      divaPhrase}
)
