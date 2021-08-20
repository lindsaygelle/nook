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
		Value:    "Linette"}

	lollyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lolly"}

	lollyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Linette"}

	lollyGermanName = nook.Name{
		Language: language.German,
		Value:    "Feline"}

	lollyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Maty"}

	lollyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラムネ"}

	lollyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사이다"}

	lollyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Feli"}

	lollyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лолли"}

	lollySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "柠檬娜"}

	lollySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Feli"}

	lollyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "檸檬娜"}
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
		Animal:   animal.Cat,
		Birthday: lollyBirthday,
		Code:     lollyCode,
		Key:      character.Lolly,
		Gender:   gender.Female,
		Name:     lollyName,
		Special:  false}
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
		language.AmericanEnglish:      lollyAmericanEnglishPhrase,
		language.CanadianFrench:       lollyCanadianFrenchPhrase,
		language.Dutch:                lollyDutchPhrase,
		language.French:               lollyFrenchPhrase,
		language.German:               lollyGermanPhrase,
		language.Italian:              lollyItalianPhrase,
		language.Japanese:             lollyJapanesePhrase,
		language.Korean:               lollyKoreanPhrase,
		language.LatinAmericanSpanish: lollyLatinAmericanSpanishPhrase,
		language.Russian:              lollyRussianPhrase,
		language.SimplifiedChinese:    lollySimplifiedChinesePhrase,
		language.Spanish:              lollySpanishPhrase,
		language.TraditionalChinese:   lollyTraditionalChinesePhrase}
)

var (
	Lolly = nook.Villager{
		Character:   lollyCharacter,
		Personality: personality.Normal,
		Phrase:      lollyPhrase}
)
