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
		Value:    "Violette"}

	divaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Diva"}

	divaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Violette"}

	divaGermanName = nook.Name{
		Language: language.German,
		Value:    "Dörte"}

	divaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Viola"}

	divaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイーダ"}

	divaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이다"}

	divaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Morania"}

	divaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дива"}

	divaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小睫"}

	divaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Morania"}

	divaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小睫"}
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
		Animal:   animal.Frog,
		Birthday: divaBirthday,
		Code:     divaCode,
		Key:      character.Diva,
		Gender:   gender.Female,
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
		Personality: personality.BigSister,
		Phrase:      divaPhrase}
)
