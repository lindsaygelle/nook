package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	vicheBirthday = nook.Birthday{
		Day:   7,
		Month: time.July}
)

var (
	vicheCode = nook.Code{
		Value: "squ20"}
)

var (
	vicheAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Viché"}

	vicheCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mara"}

	vicheDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	vicheFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mara"}

	vicheGermanName = nook.Name{
		Language: language.German,
		Value:    "L-Pyon"}

	vicheItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Vicky"}

	vicheJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みさき"}

	vicheKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미사키"}

	vicheLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ardelta"}

	vicheRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	vicheSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	vicheSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ardelta"}

	vicheTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	vicheName = nook.Languages{
		language.AmericanEnglish:      vicheAmericanEnglishName,
		language.CanadianFrench:       vicheCanadianFrenchName,
		language.Dutch:                vicheDutchName,
		language.French:               vicheFrenchName,
		language.German:               vicheGermanName,
		language.Italian:              vicheItalianName,
		language.Japanese:             vicheJapaneseName,
		language.Korean:               vicheKoreanName,
		language.LatinAmericanSpanish: vicheLatinAmericanSpanishName,
		language.Russian:              vicheRussianName,
		language.SimplifiedChinese:    vicheSimplifiedChineseName,
		language.Spanish:              vicheSpanishName,
		language.TraditionalChinese:   vicheTraditionalChineseName}
)

var (
	vicheCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: vicheBirthday,
		Code:     vicheCode,
		Gender:   gender.Female,
		Name:     vicheName}
)

var (
	vicheAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "splatastic"}

	vicheCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "marébasse"}

	vicheDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	vicheFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "marébasse"}

	vicheGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "limone"}

	vicheItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "maromaro"}

	vicheJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なんね"}

	vicheKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그거좋아"}

	vicheLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "costina"}

	vicheRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	vicheSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	vicheSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "costina"}

	vicheTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	vichePhrase = nook.Languages{
		language.AmericanEnglish:      vicheAmericanEnglishName,
		language.CanadianFrench:       vicheCanadianFrenchName,
		language.Dutch:                vicheDutchName,
		language.French:               vicheFrenchName,
		language.German:               vicheGermanName,
		language.Italian:              vicheItalianName,
		language.Japanese:             vicheJapaneseName,
		language.Korean:               vicheKoreanName,
		language.LatinAmericanSpanish: vicheLatinAmericanSpanishName,
		language.Russian:              vicheRussianName,
		language.SimplifiedChinese:    vicheSimplifiedChineseName,
		language.Spanish:              vicheSpanishName,
		language.TraditionalChinese:   vicheTraditionalChineseName}
)

var (
	Viche = nook.Villager{
		Character:   vicheCharacter,
		Personality: personality.Normal,
		Phrase:      vichePhrase}
)
