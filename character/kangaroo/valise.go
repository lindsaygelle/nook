package kangaroo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	valiseBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	valiseCode = nook.Code{
		Value: ""}
)

var (
	valiseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Valise"}

	valiseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	valiseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	valiseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pochette"}

	valiseGermanName = nook.Name{
		Language: language.German,
		Value:    "Elke"}

	valiseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pochette"}

	valiseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エプロン"}

	valiseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	valiseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	valiseRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	valiseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "昭昭"}

	valiseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ana Rosa"}

	valiseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	valiseName = nook.Languages{
		language.AmericanEnglish:      valiseAmericanEnglishName,
		language.CanadianFrench:       valiseCanadianFrenchName,
		language.Dutch:                valiseDutchName,
		language.French:               valiseFrenchName,
		language.German:               valiseGermanName,
		language.Italian:              valiseItalianName,
		language.Japanese:             valiseJapaneseName,
		language.Korean:               valiseKoreanName,
		language.LatinAmericanSpanish: valiseLatinAmericanSpanishName,
		language.Russian:              valiseRussianName,
		language.SimplifiedChinese:    valiseSimplifiedChineseName,
		language.Spanish:              valiseSpanishName,
		language.TraditionalChinese:   valiseTraditionalChineseName}
)

var (
	valiseCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: valiseBirthday,
		Code:     valiseCode,
		Gender:   gender.Female,
		Name:     valiseName}
)

var (
	valiseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tadder"}

	valiseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	valiseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	valiseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "p'tit bout"}

	valiseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krabbel"}

	valiseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bonjour"}

	valiseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ピョン"}

	valiseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	valiseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	valiseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	valiseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噼哟"}

	valiseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "personilla"}

	valiseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	valisePhrase = nook.Languages{
		language.AmericanEnglish:      valiseAmericanEnglishName,
		language.CanadianFrench:       valiseCanadianFrenchName,
		language.Dutch:                valiseDutchName,
		language.French:               valiseFrenchName,
		language.German:               valiseGermanName,
		language.Italian:              valiseItalianName,
		language.Japanese:             valiseJapaneseName,
		language.Korean:               valiseKoreanName,
		language.LatinAmericanSpanish: valiseLatinAmericanSpanishName,
		language.Russian:              valiseRussianName,
		language.SimplifiedChinese:    valiseSimplifiedChineseName,
		language.Spanish:              valiseSpanishName,
		language.TraditionalChinese:   valiseTraditionalChineseName}
)

var (
	Valise = nook.Villager{
		Character:   valiseCharacter,
		Personality: personality.Snooty,
		Phrase:      valisePhrase}
)
