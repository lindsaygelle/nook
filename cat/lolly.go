package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	lollyBirthday = nook.Birthday{
		Day:   27,
		Month: time.March}
)

var (
	lollyCode = nook.Code{
		Value: "cat18"}
)

var (
	lollyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lolly"}

	lollyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Linetteron-ron"}

	lollyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lollyzuurtje"}

	lollyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Linetteron-ron"}

	lollyGermanName = nook.Name{
		Language: language.German,
		Value:    "Felinekratz"}

	lollyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Matypurrrr"}

	lollyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラムネあのね"}

	lollyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사이다퐁퐁"}

	lollyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Felipurrrr"}

	lollyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лолликонфетка"}

	lollySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "柠檬娜娜个"}

	lollySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Felipurrrr"}

	lollyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "檸檬娜娜個"}
)

var (
	lollyName = nook.Languages{
		language.AmericanEnglish:      lollyAmericanEnglishName,
		language.CanadianFrench:       lollyCanadianFrenchName,
		language.Dutch:                lollyDutchName,
		language.French:               lollyFrenchName,
		language.German:               lollyGermanName,
		language.Italian:              lollyItalianName,
		language.Japanese:             lollyJapaneseName,
		language.Korean:               lollyKoreanName,
		language.LatinAmericanSpanish: lollyLatinAmericanSpanishName,
		language.Russian:              lollyRussianName,
		language.SimplifiedChinese:    lollySimplifiedChineseName,
		language.Spanish:              lollySpanishName,
		language.TraditionalChinese:   lollyTraditionalChineseName}
)

var (
	lollyCharacter = nook.Character{
		Animal:   Cat,
		Birthday: lollyBirthday,
		Code:     lollyCode,
		Gender:   nook.Female,
		Name:     lollyName}
)

var (
	lollyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bonbon"}

	lollyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ron-ron"}

	lollyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zuurtje"}

	lollyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ron-ron"}

	lollyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kratz"}

	lollyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "purrrr"}

	lollyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あのね"}

	lollyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "퐁퐁"}

	lollyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "purrrr"}

	lollyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "конфетка"}

	lollySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "娜个"}

	lollySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "purrrr"}

	lollyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "娜個"}
)

var (
	lollyPhrase = nook.Languages{
		language.AmericanEnglish:      lollyAmericanEnglishName,
		language.CanadianFrench:       lollyCanadianFrenchName,
		language.Dutch:                lollyDutchName,
		language.French:               lollyFrenchName,
		language.German:               lollyGermanName,
		language.Italian:              lollyItalianName,
		language.Japanese:             lollyJapaneseName,
		language.Korean:               lollyKoreanName,
		language.LatinAmericanSpanish: lollyLatinAmericanSpanishName,
		language.Russian:              lollyRussianName,
		language.SimplifiedChinese:    lollySimplifiedChineseName,
		language.Spanish:              lollySpanishName,
		language.TraditionalChinese:   lollyTraditionalChineseName}
)

var (
	Lolly = nook.Villager{
		Character:   lollyCharacter,
		Personality: nook.Normal,
		Phrase:      lollyPhrase}
)
