package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sallyBirthday = nook.Birthday{
		Day:   19,
		Month: time.June}
)

var (
	sallyCode = nook.Code{
		Value: "squ07"}
)

var (
	sallyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sally"}

	sallyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Damiatac"}

	sallyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sallymuskaatje"}

	sallyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Damiatac"}

	sallyGermanName = nook.Name{
		Language: language.German,
		Value:    "Hannemuskat"}

	sallyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rossellanocciolina"}

	sallyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ララミーったら"}

	sallyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "라라미까꿍"}

	sallyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pralinéraminí"}

	sallyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Саллимускатик"}

	sallySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "拉拉米要是"}

	sallySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pralinéramita"}

	sallyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "拉拉米要是"}
)

var (
	sallyName = nook.Languages{
		language.AmericanEnglish:      sallyAmericanEnglishName,
		language.CanadianFrench:       sallyCanadianFrenchName,
		language.Dutch:                sallyDutchName,
		language.French:               sallyFrenchName,
		language.German:               sallyGermanName,
		language.Italian:              sallyItalianName,
		language.Japanese:             sallyJapaneseName,
		language.Korean:               sallyKoreanName,
		language.LatinAmericanSpanish: sallyLatinAmericanSpanishName,
		language.Russian:              sallyRussianName,
		language.SimplifiedChinese:    sallySimplifiedChineseName,
		language.Spanish:              sallySpanishName,
		language.TraditionalChinese:   sallyTraditionalChineseName}
)

var (
	sallyCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: sallyBirthday,
		Code:     sallyCode,
		Gender:   nook.Female,
		Name:     sallyName}
)

var (
	sallyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nutmeg"}

	sallyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tac"}

	sallyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "muskaatje"}

	sallyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tac"}

	sallyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "muskat"}

	sallyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "nocciolina"}

	sallyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ったら"}

	sallyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "까꿍"}

	sallyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "raminí"}

	sallyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мускатик"}

	sallySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "要是"}

	sallySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ramita"}

	sallyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "要是"}
)

var (
	sallyPhrase = nook.Languages{
		language.AmericanEnglish:      sallyAmericanEnglishName,
		language.CanadianFrench:       sallyCanadianFrenchName,
		language.Dutch:                sallyDutchName,
		language.French:               sallyFrenchName,
		language.German:               sallyGermanName,
		language.Italian:              sallyItalianName,
		language.Japanese:             sallyJapaneseName,
		language.Korean:               sallyKoreanName,
		language.LatinAmericanSpanish: sallyLatinAmericanSpanishName,
		language.Russian:              sallyRussianName,
		language.SimplifiedChinese:    sallySimplifiedChineseName,
		language.Spanish:              sallySpanishName,
		language.TraditionalChinese:   sallyTraditionalChineseName}
)

var (
	Sally = nook.Villager{
		Character:   sallyCharacter,
		Personality: nook.Normal,
		Phrase:      sallyPhrase}
)
