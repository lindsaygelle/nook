package duck

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
	mallaryBirthday = nook.Birthday{
		Day:   17,
		Month: time.November}
)

var (
	mallaryCode = nook.Code{
		Value: "duk06"}
)

var (
	mallaryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mallary"}

	mallaryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mallory"}

	mallaryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mallary"}

	mallaryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mallory"}

	mallaryGermanName = nook.Name{
		Language: language.German,
		Value:    "Marina"}

	mallaryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sofia"}

	mallaryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スミモモ"}

	mallaryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스미모"}

	mallaryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mirta"}

	mallaryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мэллари"}

	mallarySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "苏米"}

	mallarySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mirta"}

	mallaryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蘇米"}
)

var (
	mallaryName = nook.Languages{
		language.AmericanEnglish:      mallaryAmericanEnglishName,
		language.CanadianFrench:       mallaryCanadianFrenchName,
		language.Dutch:                mallaryDutchName,
		language.French:               mallaryFrenchName,
		language.German:               mallaryGermanName,
		language.Italian:              mallaryItalianName,
		language.Japanese:             mallaryJapaneseName,
		language.Korean:               mallaryKoreanName,
		language.LatinAmericanSpanish: mallaryLatinAmericanSpanishName,
		language.Russian:              mallaryRussianName,
		language.SimplifiedChinese:    mallarySimplifiedChineseName,
		language.Spanish:              mallarySpanishName,
		language.TraditionalChinese:   mallaryTraditionalChineseName}
)

var (
	mallaryCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: mallaryBirthday,
		Code:     mallaryCode,
		Key:      character.Mallary,
		Gender:   gender.Female,
		Name:     mallaryName}
)

var (
	mallaryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quackpth"}

	mallaryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coinpff"}

	mallaryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwaks"}

	mallaryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "coinpff"}

	mallaryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quackpss"}

	mallaryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quackplà"}

	mallaryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヨネ"}

	mallaryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "모모"}

	mallaryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "recuac"}

	mallaryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кряк-фи"}

	mallarySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "八十八"}

	mallarySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "recuac"}

	mallaryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "八十八"}
)

var (
	mallaryPhrase = nook.Languages{
		language.AmericanEnglish:      mallaryAmericanEnglishName,
		language.CanadianFrench:       mallaryCanadianFrenchName,
		language.Dutch:                mallaryDutchName,
		language.French:               mallaryFrenchName,
		language.German:               mallaryGermanName,
		language.Italian:              mallaryItalianName,
		language.Japanese:             mallaryJapaneseName,
		language.Korean:               mallaryKoreanName,
		language.LatinAmericanSpanish: mallaryLatinAmericanSpanishName,
		language.Russian:              mallaryRussianName,
		language.SimplifiedChinese:    mallarySimplifiedChineseName,
		language.Spanish:              mallarySpanishName,
		language.TraditionalChinese:   mallaryTraditionalChineseName}
)

var (
	Mallary = nook.Villager{
		Character:   mallaryCharacter,
		Personality: personality.Snooty,
		Phrase:      mallaryPhrase}
)
