package squirrel

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
		Value:    "Pécan"}

	pecanDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pecan"}

	pecanFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pécan"}

	pecanGermanName = nook.Name{
		Language: language.German,
		Value:    "Noisette"}

	pecanItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nocina"}

	pecanJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レベッカ"}

	pecanKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레베카"}

	pecanLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Camila"}

	pecanRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пекан"}

	pecanSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雷贝嘉"}

	pecanSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Camila"}

	pecanTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雷貝嘉"}
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
		Animal:   animal.Squirrel,
		Birthday: pecanBirthday,
		Code:     pecanCode,
		Key:      character.Pecan,
		Gender:   gender.Female,
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
		language.AmericanEnglish:      pecanAmericanEnglishPhrase,
		language.CanadianFrench:       pecanCanadianFrenchPhrase,
		language.Dutch:                pecanDutchPhrase,
		language.French:               pecanFrenchPhrase,
		language.German:               pecanGermanPhrase,
		language.Italian:              pecanItalianPhrase,
		language.Japanese:             pecanJapanesePhrase,
		language.Korean:               pecanKoreanPhrase,
		language.LatinAmericanSpanish: pecanLatinAmericanSpanishPhrase,
		language.Russian:              pecanRussianPhrase,
		language.SimplifiedChinese:    pecanSimplifiedChinesePhrase,
		language.Spanish:              pecanSpanishPhrase,
		language.TraditionalChinese:   pecanTraditionalChinesePhrase}
)

var (
	Pecan = nook.Villager{
		Character:   pecanCharacter,
		Personality: personality.Snooty,
		Phrase:      pecanPhrase}
)
