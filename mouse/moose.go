package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	mooseBirthday = nook.Birthday{
		Day:   13,
		Month: time.September}
)

var (
	mooseCode = nook.Code{
		Value: "mus14"}
)

var (
	mooseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Moose"}

	mooseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Joachimbééléé"}

	mooseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Moosekleintje"}

	mooseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Joachimbééléé"}

	mooseGermanName = nook.Name{
		Language: language.German,
		Value:    "Mausbertkäääse"}

	mooseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Aldoquink"}

	mooseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピンにゃろ"}

	mooseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핑짜샤"}

	mooseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Satoatiza"}

	mooseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Музкоротышка"}

	mooseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿聘可恶"}

	mooseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Satoatiza"}

	mooseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿聘可惡"}
)

var (
	mooseName = nook.Languages{
		language.AmericanEnglish:      mooseAmericanEnglishName,
		language.CanadianFrench:       mooseCanadianFrenchName,
		language.Dutch:                mooseDutchName,
		language.French:               mooseFrenchName,
		language.German:               mooseGermanName,
		language.Italian:              mooseItalianName,
		language.Japanese:             mooseJapaneseName,
		language.Korean:               mooseKoreanName,
		language.LatinAmericanSpanish: mooseLatinAmericanSpanishName,
		language.Russian:              mooseRussianName,
		language.SimplifiedChinese:    mooseSimplifiedChineseName,
		language.Spanish:              mooseSpanishName,
		language.TraditionalChinese:   mooseTraditionalChineseName}
)

var (
	mooseCharacter = nook.Character{
		Animal:   Mouse,
		Birthday: mooseBirthday,
		Code:     mooseCode,
		Gender:   nook.Male,
		Name:     mooseName}
)

var (
	mooseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "shorty"}

	mooseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bééléé"}

	mooseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kleintje"}

	mooseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bééléé"}

	mooseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "käääse"}

	mooseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quink"}

	mooseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "にゃろ"}

	mooseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "짜샤"}

	mooseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "atiza"}

	mooseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "коротышка"}

	mooseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "可恶"}

	mooseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "atiza"}

	mooseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "可惡"}
)

var (
	moosePhrase = nook.Languages{
		language.AmericanEnglish:      mooseAmericanEnglishName,
		language.CanadianFrench:       mooseCanadianFrenchName,
		language.Dutch:                mooseDutchName,
		language.French:               mooseFrenchName,
		language.German:               mooseGermanName,
		language.Italian:              mooseItalianName,
		language.Japanese:             mooseJapaneseName,
		language.Korean:               mooseKoreanName,
		language.LatinAmericanSpanish: mooseLatinAmericanSpanishName,
		language.Russian:              mooseRussianName,
		language.SimplifiedChinese:    mooseSimplifiedChineseName,
		language.Spanish:              mooseSpanishName,
		language.TraditionalChinese:   mooseTraditionalChineseName}
)

var (
	Moose = nook.Villager{
		Character:   mooseCharacter,
		Personality: nook.Jock,
		Phrase:      moosePhrase}
)
