package ostrich

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sprocketBirthday = nook.Birthday{
		Day:   1,
		Month: time.December}
)

var (
	sprocketCode = nook.Code{
		Value: "ost03"}
)

var (
	sprocketAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sprocket"}

	sprocketCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Laflèchepik pik"}

	sprocketDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sprocketklik-kluk"}

	sprocketFrenchName = nook.Name{
		Language: language.French,
		Value:    "Laflèchepik pik"}

	sprocketGermanName = nook.Name{
		Language: language.German,
		Value:    "Lutzaffenzahn"}

	sprocketItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frankieniiun"}

	sprocketJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヘルツだメカ"}

	sprocketKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "헤르츠치지직"}

	sprocketLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ráfagachiuuun"}

	sprocketRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Спрокетдзынь"}

	sprocketSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鹤兹机械"}

	sprocketSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ráfagachiuuun"}

	sprocketTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鶴茲機械"}
)

var (
	sprocketName = nook.Languages{
		language.AmericanEnglish:      sprocketAmericanEnglishName,
		language.CanadianFrench:       sprocketCanadianFrenchName,
		language.Dutch:                sprocketDutchName,
		language.French:               sprocketFrenchName,
		language.German:               sprocketGermanName,
		language.Italian:              sprocketItalianName,
		language.Japanese:             sprocketJapaneseName,
		language.Korean:               sprocketKoreanName,
		language.LatinAmericanSpanish: sprocketLatinAmericanSpanishName,
		language.Russian:              sprocketRussianName,
		language.SimplifiedChinese:    sprocketSimplifiedChineseName,
		language.Spanish:              sprocketSpanishName,
		language.TraditionalChinese:   sprocketTraditionalChineseName}
)

var (
	sprocketCharacter = nook.Character{
		Animal:   Ostrich,
		Birthday: sprocketBirthday,
		Code:     sprocketCode,
		Gender:   nook.Male,
		Name:     sprocketName}
)

var (
	sprocketAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zort"}

	sprocketCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pik pik"}

	sprocketDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "klik-kluk"}

	sprocketFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pik pik"}

	sprocketGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "affenzahn"}

	sprocketItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "niiun"}

	sprocketJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だメカ"}

	sprocketKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "치지직"}

	sprocketLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chiuuun"}

	sprocketRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дзынь"}

	sprocketSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "机械"}

	sprocketSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chiuuun"}

	sprocketTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "機械"}
)

var (
	sprocketPhrase = nook.Languages{
		language.AmericanEnglish:      sprocketAmericanEnglishName,
		language.CanadianFrench:       sprocketCanadianFrenchName,
		language.Dutch:                sprocketDutchName,
		language.French:               sprocketFrenchName,
		language.German:               sprocketGermanName,
		language.Italian:              sprocketItalianName,
		language.Japanese:             sprocketJapaneseName,
		language.Korean:               sprocketKoreanName,
		language.LatinAmericanSpanish: sprocketLatinAmericanSpanishName,
		language.Russian:              sprocketRussianName,
		language.SimplifiedChinese:    sprocketSimplifiedChineseName,
		language.Spanish:              sprocketSpanishName,
		language.TraditionalChinese:   sprocketTraditionalChineseName}
)

var (
	Sprocket = nook.Villager{
		Character:   sprocketCharacter,
		Personality: nook.Jock,
		Phrase:      sprocketPhrase}
)
