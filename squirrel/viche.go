package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

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
		Value:    "N/A"}

	vicheSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	vicheSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ardelta"}

	vicheTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Squirrel,
		Birthday: vicheBirthday,
		Code:     vicheCode,
		Gender:   nook.Female,
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
		Value:    "N/A"}

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
		Value:    "N/A"}

	vicheSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	vicheSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "costina"}

	vicheTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Normal,
		Phrase:      vichePhrase}
)
