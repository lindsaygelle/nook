package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	trufflesBirthday = nook.Birthday{
		Day:   28,
		Month: time.July}
)

var (
	trufflesCode = nook.Code{
		Value: "pig01"}
)

var (
	trufflesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Truffles"}

	trufflesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Trufajambon"}

	trufflesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Trufflessnuiter"}

	trufflesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Trufajambon"}

	trufflesGermanName = nook.Name{
		Language: language.German,
		Value:    "Luzieschnauzie"}

	trufflesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grufettaficcanaso"}

	trufflesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トンコだわさ"}

	trufflesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "탱고거라예"}

	trufflesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Trufasoink-oink"}

	trufflesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Трафлсхрюша"}

	trufflesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘟嘟哇赛"}

	trufflesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Trufasoink-oink"}

	trufflesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘟嘟哇賽"}
)

var (
	trufflesName = nook.Languages{
		language.AmericanEnglish:      trufflesAmericanEnglishName,
		language.CanadianFrench:       trufflesCanadianFrenchName,
		language.Dutch:                trufflesDutchName,
		language.French:               trufflesFrenchName,
		language.German:               trufflesGermanName,
		language.Italian:              trufflesItalianName,
		language.Japanese:             trufflesJapaneseName,
		language.Korean:               trufflesKoreanName,
		language.LatinAmericanSpanish: trufflesLatinAmericanSpanishName,
		language.Russian:              trufflesRussianName,
		language.SimplifiedChinese:    trufflesSimplifiedChineseName,
		language.Spanish:              trufflesSpanishName,
		language.TraditionalChinese:   trufflesTraditionalChineseName}
)

var (
	trufflesCharacter = nook.Character{
		Animal:   Pig,
		Birthday: trufflesBirthday,
		Code:     trufflesCode,
		Gender:   nook.Female,
		Name:     trufflesName}
)

var (
	trufflesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snoutie"}

	trufflesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "jambon"}

	trufflesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuiter"}

	trufflesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "jambon"}

	trufflesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnauzie"}

	trufflesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ficcanaso"}

	trufflesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だわさ"}

	trufflesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "거라예"}

	trufflesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oink-oink"}

	trufflesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюша"}

	trufflesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哇赛"}

	trufflesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oink-oink"}

	trufflesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哇賽"}
)

var (
	trufflesPhrase = nook.Languages{
		language.AmericanEnglish:      trufflesAmericanEnglishName,
		language.CanadianFrench:       trufflesCanadianFrenchName,
		language.Dutch:                trufflesDutchName,
		language.French:               trufflesFrenchName,
		language.German:               trufflesGermanName,
		language.Italian:              trufflesItalianName,
		language.Japanese:             trufflesJapaneseName,
		language.Korean:               trufflesKoreanName,
		language.LatinAmericanSpanish: trufflesLatinAmericanSpanishName,
		language.Russian:              trufflesRussianName,
		language.SimplifiedChinese:    trufflesSimplifiedChineseName,
		language.Spanish:              trufflesSpanishName,
		language.TraditionalChinese:   trufflesTraditionalChineseName}
)

var (
	Truffles = nook.Villager{
		Character:   trufflesCharacter,
		Personality: nook.Peppy,
		Phrase:      trufflesPhrase}
)
