package kangaroo

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Pochettep'tit bout"}

	valiseGermanName = nook.Name{
		Language: language.German,
		Value:    "Elkekrabbel"}

	valiseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pochettebonjour"}

	valiseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エプロンピョン"}

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
		Value:    "昭昭噼哟"}

	valiseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ana Rosapersonilla"}

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
		Animal:   Kangaroo,
		Birthday: valiseBirthday,
		Code:     valiseCode,
		Gender:   nook.Female,
		Name:     valiseName}
)

var (
	valiseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	valiseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	valiseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	valiseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	valiseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	valiseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噼哟"}

	valiseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "personilla"}

	valiseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Snooty,
		Phrase:      valisePhrase}
)
