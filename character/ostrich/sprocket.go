package ostrich

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
		Value:    "Laflèche"}

	sprocketDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sprocket"}

	sprocketFrenchName = nook.Name{
		Language: language.French,
		Value:    "Laflèche"}

	sprocketGermanName = nook.Name{
		Language: language.German,
		Value:    "Lutz"}

	sprocketItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frankie"}

	sprocketJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヘルツ"}

	sprocketKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "헤르츠"}

	sprocketLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ráfaga"}

	sprocketRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Спрокет"}

	sprocketSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鹤兹"}

	sprocketSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ráfaga"}

	sprocketTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鶴茲"}
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
		Animal:   animal.Ostrich,
		Birthday: sprocketBirthday,
		Code:     sprocketCode,
		Key:      character.Sprocket,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      sprocketAmericanEnglishPhrase,
		language.CanadianFrench:       sprocketCanadianFrenchPhrase,
		language.Dutch:                sprocketDutchPhrase,
		language.French:               sprocketFrenchPhrase,
		language.German:               sprocketGermanPhrase,
		language.Italian:              sprocketItalianPhrase,
		language.Japanese:             sprocketJapanesePhrase,
		language.Korean:               sprocketKoreanPhrase,
		language.LatinAmericanSpanish: sprocketLatinAmericanSpanishPhrase,
		language.Russian:              sprocketRussianPhrase,
		language.SimplifiedChinese:    sprocketSimplifiedChinesePhrase,
		language.Spanish:              sprocketSpanishPhrase,
		language.TraditionalChinese:   sprocketTraditionalChinesePhrase}
)

var (
	Sprocket = nook.Villager{
		Character:   sprocketCharacter,
		Personality: personality.Jock,
		Phrase:      sprocketPhrase}
)
