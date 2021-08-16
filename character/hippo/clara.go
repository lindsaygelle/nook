package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	claraBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	claraCode = nook.Code{
		Value: ""}
)

var (
	claraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Clara"}

	claraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	claraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	claraFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	claraGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	claraItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	claraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クララ"}

	claraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	claraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	claraRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	claraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	claraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	claraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	claraName = nook.Languages{
		language.AmericanEnglish:      claraAmericanEnglishName,
		language.CanadianFrench:       claraCanadianFrenchName,
		language.Dutch:                claraDutchName,
		language.French:               claraFrenchName,
		language.German:               claraGermanName,
		language.Italian:              claraItalianName,
		language.Japanese:             claraJapaneseName,
		language.Korean:               claraKoreanName,
		language.LatinAmericanSpanish: claraLatinAmericanSpanishName,
		language.Russian:              claraRussianName,
		language.SimplifiedChinese:    claraSimplifiedChineseName,
		language.Spanish:              claraSpanishName,
		language.TraditionalChinese:   claraTraditionalChineseName}
)

var (
	claraCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: claraBirthday,
		Code:     claraCode,
		Gender:   gender.Female,
		Name:     claraName}
)

var (
	claraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "うふふ"}

	claraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	claraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	claraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	claraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	claraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	claraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "うふふ"}

	claraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	claraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	claraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	claraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	claraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	claraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	claraPhrase = nook.Languages{
		language.AmericanEnglish:      claraAmericanEnglishName,
		language.CanadianFrench:       claraCanadianFrenchName,
		language.Dutch:                claraDutchName,
		language.French:               claraFrenchName,
		language.German:               claraGermanName,
		language.Italian:              claraItalianName,
		language.Japanese:             claraJapaneseName,
		language.Korean:               claraKoreanName,
		language.LatinAmericanSpanish: claraLatinAmericanSpanishName,
		language.Russian:              claraRussianName,
		language.SimplifiedChinese:    claraSimplifiedChineseName,
		language.Spanish:              claraSpanishName,
		language.TraditionalChinese:   claraTraditionalChineseName}
)

var (
	Clara = nook.Villager{
		Character:   claraCharacter,
		Personality: personality.Normal,
		Phrase:      claraPhrase}
)
