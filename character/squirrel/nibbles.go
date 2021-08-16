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
	nibblesBirthday = nook.Birthday{
		Day:   19,
		Month: time.July}
)

var (
	nibblesCode = nook.Code{
		Value: "squ04"}
)

var (
	nibblesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nibbles"}

	nibblesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lola"}

	nibblesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nibbles"}

	nibblesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lola"}

	nibblesGermanName = nook.Name{
		Language: language.German,
		Value:    "Knuspi"}

	nibblesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rododea"}

	nibblesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガリガリ"}

	nibblesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "그리미"}

	nibblesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dentina"}

	nibblesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ниблз"}

	nibblesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咬咬"}

	nibblesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Dentina"}

	nibblesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咬咬"}
)

var (
	nibblesName = nook.Languages{
		language.AmericanEnglish:      nibblesAmericanEnglishName,
		language.CanadianFrench:       nibblesCanadianFrenchName,
		language.Dutch:                nibblesDutchName,
		language.French:               nibblesFrenchName,
		language.German:               nibblesGermanName,
		language.Italian:              nibblesItalianName,
		language.Japanese:             nibblesJapaneseName,
		language.Korean:               nibblesKoreanName,
		language.LatinAmericanSpanish: nibblesLatinAmericanSpanishName,
		language.Russian:              nibblesRussianName,
		language.SimplifiedChinese:    nibblesSimplifiedChineseName,
		language.Spanish:              nibblesSpanishName,
		language.TraditionalChinese:   nibblesTraditionalChineseName}
)

var (
	nibblesCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: nibblesBirthday,
		Code:     nibblesCode,
		Key:      character.Nibbles,
		Gender:   gender.Female,
		Name:     nibblesName}
)

var (
	nibblesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "niblet"}

	nibblesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grignotte"}

	nibblesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knabbel"}

	nibblesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nunuche"}

	nibblesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knusper"}

	nibblesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "truciolo"}

	nibblesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ガジガジ"}

	nibblesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "송송"}

	nibblesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "trocitrí"}

	nibblesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "орешек"}

	nibblesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咯吱咯吱"}

	nibblesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bellotita"}

	nibblesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咯吱咯吱"}
)

var (
	nibblesPhrase = nook.Languages{
		language.AmericanEnglish:      nibblesAmericanEnglishName,
		language.CanadianFrench:       nibblesCanadianFrenchName,
		language.Dutch:                nibblesDutchName,
		language.French:               nibblesFrenchName,
		language.German:               nibblesGermanName,
		language.Italian:              nibblesItalianName,
		language.Japanese:             nibblesJapaneseName,
		language.Korean:               nibblesKoreanName,
		language.LatinAmericanSpanish: nibblesLatinAmericanSpanishName,
		language.Russian:              nibblesRussianName,
		language.SimplifiedChinese:    nibblesSimplifiedChineseName,
		language.Spanish:              nibblesSpanishName,
		language.TraditionalChinese:   nibblesTraditionalChineseName}
)

var (
	Nibbles = nook.Villager{
		Character:   nibblesCharacter,
		Personality: personality.Peppy,
		Phrase:      nibblesPhrase}
)
