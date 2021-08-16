package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	robinBirthday = nook.Birthday{
		Day:   4,
		Month: time.December}
)

var (
	robinCode = nook.Code{
		Value: "brd01"}
)

var (
	robinAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Robin"}

	robinCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Robie"}

	robinDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Robin"}

	robinFrenchName = nook.Name{
		Language: language.French,
		Value:    "Robie"}

	robinGermanName = nook.Name{
		Language: language.German,
		Value:    "Jule"}

	robinItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rossana"}

	robinJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パーチク"}

	robinKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파틱"}

	robinLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aria"}

	robinRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Робин"}

	robinSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喳喳"}

	robinSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aria"}

	robinTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喳喳"}
)

var (
	robinName = nook.Languages{
		language.AmericanEnglish:      robinAmericanEnglishName,
		language.CanadianFrench:       robinCanadianFrenchName,
		language.Dutch:                robinDutchName,
		language.French:               robinFrenchName,
		language.German:               robinGermanName,
		language.Italian:              robinItalianName,
		language.Japanese:             robinJapaneseName,
		language.Korean:               robinKoreanName,
		language.LatinAmericanSpanish: robinLatinAmericanSpanishName,
		language.Russian:              robinRussianName,
		language.SimplifiedChinese:    robinSimplifiedChineseName,
		language.Spanish:              robinSpanishName,
		language.TraditionalChinese:   robinTraditionalChineseName}
)

var (
	robinCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: robinBirthday,
		Code:     robinCode,
		Gender:   gender.Female,
		Name:     robinName}
)

var (
	robinAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "la-di-da"}

	robinCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ta-di-da"}

	robinDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "lalala"}

	robinFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ta-di-da"}

	robinGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ladida"}

	robinItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yuppidù"}

	robinJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "さ"}

	robinKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "글쎄"}

	robinLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "larará"}

	robinRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ля-ля-ля"}

	robinSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喳"}

	robinSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "picatoste"}

	robinTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喳"}
)

var (
	robinPhrase = nook.Languages{
		language.AmericanEnglish:      robinAmericanEnglishName,
		language.CanadianFrench:       robinCanadianFrenchName,
		language.Dutch:                robinDutchName,
		language.French:               robinFrenchName,
		language.German:               robinGermanName,
		language.Italian:              robinItalianName,
		language.Japanese:             robinJapaneseName,
		language.Korean:               robinKoreanName,
		language.LatinAmericanSpanish: robinLatinAmericanSpanishName,
		language.Russian:              robinRussianName,
		language.SimplifiedChinese:    robinSimplifiedChineseName,
		language.Spanish:              robinSpanishName,
		language.TraditionalChinese:   robinTraditionalChineseName}
)

var (
	Robin = nook.Villager{
		Character:   robinCharacter,
		Personality: personality.Snooty,
		Phrase:      robinPhrase}
)
