package pig

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
		Value:    "Trufa"}

	trufflesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Truffles"}

	trufflesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Trufa"}

	trufflesGermanName = nook.Name{
		Language: language.German,
		Value:    "Luzie"}

	trufflesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grufetta"}

	trufflesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トンコ"}

	trufflesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "탱고"}

	trufflesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Trufas"}

	trufflesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Трафлс"}

	trufflesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘟嘟"}

	trufflesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Trufas"}

	trufflesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘟嘟"}
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
		Animal:   animal.Pig,
		Birthday: trufflesBirthday,
		Code:     trufflesCode,
		Key:      character.Truffles,
		Gender:   gender.Female,
		Name:     trufflesName,
		Special:  false}
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
		language.AmericanEnglish:      trufflesAmericanEnglishPhrase,
		language.CanadianFrench:       trufflesCanadianFrenchPhrase,
		language.Dutch:                trufflesDutchPhrase,
		language.French:               trufflesFrenchPhrase,
		language.German:               trufflesGermanPhrase,
		language.Italian:              trufflesItalianPhrase,
		language.Japanese:             trufflesJapanesePhrase,
		language.Korean:               trufflesKoreanPhrase,
		language.LatinAmericanSpanish: trufflesLatinAmericanSpanishPhrase,
		language.Russian:              trufflesRussianPhrase,
		language.SimplifiedChinese:    trufflesSimplifiedChinesePhrase,
		language.Spanish:              trufflesSpanishPhrase,
		language.TraditionalChinese:   trufflesTraditionalChinesePhrase}
)

var (
	Truffles = nook.Villager{
		Character:   trufflesCharacter,
		Personality: personality.Peppy,
		Phrase:      trufflesPhrase}
)
