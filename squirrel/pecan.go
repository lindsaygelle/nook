package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pecanBirthday = nook.Birthday{
		Day:   10,
		Month: time.September}
)

var (
	pecanCode = nook.Code{
		Value: "squ03"}
)

var (
	pecanAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pecan"}

	pecanCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pécanfouloulou"}

	pecanDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pecaneekhoorn"}

	pecanFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pécanfouloulou"}

	pecanGermanName = nook.Name{
		Language: language.German,
		Value:    "Noisettespatzl"}

	pecanItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nocinacipì"}

	pecanJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レベッカつんっ"}

	pecanKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레베카쯧쯧"}

	pecanLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Camilapiñoní"}

	pecanRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пеканбурундук"}

	pecanSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雷贝嘉正经"}

	pecanSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Camilapiñoncito"}

	pecanTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雷貝嘉正經"}
)

var (
	pecanName = nook.Languages{
		language.AmericanEnglish:      pecanAmericanEnglishName,
		language.CanadianFrench:       pecanCanadianFrenchName,
		language.Dutch:                pecanDutchName,
		language.French:               pecanFrenchName,
		language.German:               pecanGermanName,
		language.Italian:              pecanItalianName,
		language.Japanese:             pecanJapaneseName,
		language.Korean:               pecanKoreanName,
		language.LatinAmericanSpanish: pecanLatinAmericanSpanishName,
		language.Russian:              pecanRussianName,
		language.SimplifiedChinese:    pecanSimplifiedChineseName,
		language.Spanish:              pecanSpanishName,
		language.TraditionalChinese:   pecanTraditionalChineseName}
)

var (
	pecanCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: pecanBirthday,
		Code:     pecanCode,
		Gender:   nook.Female,
		Name:     pecanName}
)

var (
	pecanAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chipmunk"}

	pecanCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "fouloulou"}

	pecanDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "eekhoorn"}

	pecanFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "fouloulou"}

	pecanGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "spatzl"}

	pecanItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cipì"}

	pecanJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "つんっ"}

	pecanKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쯧쯧"}

	pecanLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "piñoní"}

	pecanRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бурундук"}

	pecanSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "正经"}

	pecanSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "piñoncito"}

	pecanTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "正經"}
)

var (
	pecanPhrase = nook.Languages{
		language.AmericanEnglish:      pecanAmericanEnglishName,
		language.CanadianFrench:       pecanCanadianFrenchName,
		language.Dutch:                pecanDutchName,
		language.French:               pecanFrenchName,
		language.German:               pecanGermanName,
		language.Italian:              pecanItalianName,
		language.Japanese:             pecanJapaneseName,
		language.Korean:               pecanKoreanName,
		language.LatinAmericanSpanish: pecanLatinAmericanSpanishName,
		language.Russian:              pecanRussianName,
		language.SimplifiedChinese:    pecanSimplifiedChineseName,
		language.Spanish:              pecanSpanishName,
		language.TraditionalChinese:   pecanTraditionalChineseName}
)

var (
	Pecan = nook.Villager{
		Character:   pecanCharacter,
		Personality: nook.Snooty,
		Phrase:      pecanPhrase}
)
