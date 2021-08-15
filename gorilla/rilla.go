package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rillaBirthday = nook.Birthday{
		Day:   1,
		Month: time.November}
)

var (
	rillaCode = nook.Code{
		Value: "gor11"}
)

var (
	rillaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rilla"}

	rillaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rillarorille"}

	rillaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rillahello"}

	rillaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rillarorille"}

	rillaGermanName = nook.Name{
		Language: language.German,
		Value:    "Rillakittykong"}

	rillaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rillaarrruuu"}

	rillaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リラハロー"}

	rillaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "릴라헬로"}

	rillaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rilakittybún"}

	rillaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Риллаалло"}

	rillaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莉拉哈罗"}

	rillaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rilakittybún"}

	rillaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莉拉哈囉"}
)

var (
	rillaName = nook.Languages{
		language.AmericanEnglish:      rillaAmericanEnglishName,
		language.CanadianFrench:       rillaCanadianFrenchName,
		language.Dutch:                rillaDutchName,
		language.French:               rillaFrenchName,
		language.German:               rillaGermanName,
		language.Italian:              rillaItalianName,
		language.Japanese:             rillaJapaneseName,
		language.Korean:               rillaKoreanName,
		language.LatinAmericanSpanish: rillaLatinAmericanSpanishName,
		language.Russian:              rillaRussianName,
		language.SimplifiedChinese:    rillaSimplifiedChineseName,
		language.Spanish:              rillaSpanishName,
		language.TraditionalChinese:   rillaTraditionalChineseName}
)

var (
	rillaCharacter = nook.Character{
		Animal:   Gorilla,
		Birthday: rillaBirthday,
		Code:     rillaCode,
		Gender:   nook.Female,
		Name:     rillaName}
)

var (
	rillaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hello"}

	rillaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "rorille"}

	rillaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hello"}

	rillaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "rorille"}

	rillaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kittykong"}

	rillaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "arrruuu"}

	rillaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ハロー"}

	rillaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "헬로"}

	rillaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kittybún"}

	rillaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "алло"}

	rillaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈罗"}

	rillaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "kittybún"}

	rillaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈囉"}
)

var (
	rillaPhrase = nook.Languages{
		language.AmericanEnglish:      rillaAmericanEnglishName,
		language.CanadianFrench:       rillaCanadianFrenchName,
		language.Dutch:                rillaDutchName,
		language.French:               rillaFrenchName,
		language.German:               rillaGermanName,
		language.Italian:              rillaItalianName,
		language.Japanese:             rillaJapaneseName,
		language.Korean:               rillaKoreanName,
		language.LatinAmericanSpanish: rillaLatinAmericanSpanishName,
		language.Russian:              rillaRussianName,
		language.SimplifiedChinese:    rillaSimplifiedChineseName,
		language.Spanish:              rillaSpanishName,
		language.TraditionalChinese:   rillaTraditionalChineseName}
)

var (
	Rilla = nook.Villager{
		Character:   rillaCharacter,
		Personality: nook.Peppy,
		Phrase:      rillaPhrase}
)
