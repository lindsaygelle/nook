package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Mallorycoinpff"}

	mallaryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mallarykwaks"}

	mallaryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mallorycoinpff"}

	mallaryGermanName = nook.Name{
		Language: language.German,
		Value:    "Marinaquackpss"}

	mallaryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sofiaquackplà"}

	mallaryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スミモモヨネ"}

	mallaryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스미모모모"}

	mallaryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mirtarecuac"}

	mallaryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мэлларикряк-фи"}

	mallarySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "苏米八十八"}

	mallarySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mirtarecuac"}

	mallaryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蘇米八十八"}
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
		Animal:   Duck,
		Birthday: mallaryBirthday,
		Code:     mallaryCode,
		Gender:   nook.Female,
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
		Personality: nook.Snooty,
		Phrase:      mallaryPhrase}
)
