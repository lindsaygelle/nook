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
		Value:    "Damia"}

	sallyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sally"}

	sallyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Damia"}

	sallyGermanName = nook.Name{
		Language: language.German,
		Value:    "Hanne"}

	sallyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rossella"}

	sallyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ララミー"}

	sallyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "라라미"}

	sallyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Praliné"}

	sallyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Салли"}

	sallySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "拉拉米"}

	sallySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Praliné"}

	sallyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "拉拉米"}
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
		Animal:   animal.Squirrel,
		Birthday: sallyBirthday,
		Code:     sallyCode,
		Key:      character.Sally,
		Gender:   gender.Female,
		Name:     sallyName,
		Special:  false}
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
		language.AmericanEnglish:      sallyAmericanEnglishPhrase,
		language.CanadianFrench:       sallyCanadianFrenchPhrase,
		language.Dutch:                sallyDutchPhrase,
		language.French:               sallyFrenchPhrase,
		language.German:               sallyGermanPhrase,
		language.Italian:              sallyItalianPhrase,
		language.Japanese:             sallyJapanesePhrase,
		language.Korean:               sallyKoreanPhrase,
		language.LatinAmericanSpanish: sallyLatinAmericanSpanishPhrase,
		language.Russian:              sallyRussianPhrase,
		language.SimplifiedChinese:    sallySimplifiedChinesePhrase,
		language.Spanish:              sallySpanishPhrase,
		language.TraditionalChinese:   sallyTraditionalChinesePhrase}
)

var (
	Sally = nook.Villager{
		Character:   sallyCharacter,
		Personality: personality.Normal,
		Phrase:      sallyPhrase}
)
