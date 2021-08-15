package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	euniceBirthday = nook.Birthday{
		Day:   3,
		Month: time.April}
)

var (
	euniceCode = nook.Code{
		Value: "shp02"}
)

var (
	euniceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Eunice"}

	euniceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bérénice"}

	euniceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Eunice"}

	euniceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bérénice"}

	euniceGermanName = nook.Name{
		Language: language.German,
		Value:    "Edith"}

	euniceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tenerina"}

	euniceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モヘア"}

	euniceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곱슬이"}

	euniceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lanolina"}

	euniceRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Юнис"}

	euniceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毛海儿"}

	euniceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lanolina"}

	euniceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毛海兒"}
)

var (
	euniceName = nook.Languages{
		language.AmericanEnglish:      euniceAmericanEnglishName,
		language.CanadianFrench:       euniceCanadianFrenchName,
		language.Dutch:                euniceDutchName,
		language.French:               euniceFrenchName,
		language.German:               euniceGermanName,
		language.Italian:              euniceItalianName,
		language.Japanese:             euniceJapaneseName,
		language.Korean:               euniceKoreanName,
		language.LatinAmericanSpanish: euniceLatinAmericanSpanishName,
		language.Russian:              euniceRussianName,
		language.SimplifiedChinese:    euniceSimplifiedChineseName,
		language.Spanish:              euniceSpanishName,
		language.TraditionalChinese:   euniceTraditionalChineseName}
)

var (
	euniceCharacter = nook.Character{
		Animal:   Sheep,
		Birthday: euniceBirthday,
		Code:     euniceCode,
		Gender:   nook.Female,
		Name:     euniceName}
)

var (
	euniceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "lambchop"}

	euniceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bêbêêêêh"}

	euniceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "krulletje"}

	euniceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bêbêêêêh"}

	euniceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "la-a-amm"}

	euniceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "laaagna"}

	euniceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "メェー"}

	euniceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "메에에"}

	euniceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "yupitiii"}

	euniceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ребрышко"}

	euniceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咩"}

	euniceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "veeenga"}

	euniceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咩"}
)

var (
	eunicePhrase = nook.Languages{
		language.AmericanEnglish:      euniceAmericanEnglishName,
		language.CanadianFrench:       euniceCanadianFrenchName,
		language.Dutch:                euniceDutchName,
		language.French:               euniceFrenchName,
		language.German:               euniceGermanName,
		language.Italian:              euniceItalianName,
		language.Japanese:             euniceJapaneseName,
		language.Korean:               euniceKoreanName,
		language.LatinAmericanSpanish: euniceLatinAmericanSpanishName,
		language.Russian:              euniceRussianName,
		language.SimplifiedChinese:    euniceSimplifiedChineseName,
		language.Spanish:              euniceSpanishName,
		language.TraditionalChinese:   euniceTraditionalChineseName}
)

var (
	Eunice = nook.Villager{
		Character:   euniceCharacter,
		Personality: nook.Normal,
		Phrase:      eunicePhrase}
)
