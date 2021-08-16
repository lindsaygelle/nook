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
	kitBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	kitCode = nook.Code{
		Value: ""}
)

var (
	kitAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kit"}

	kitCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	kitDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	kitFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	kitGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	kitItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	kitJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キット"}

	kitKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	kitLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	kitRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	kitSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	kitSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	kitTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	kitName = nook.Languages{
		language.AmericanEnglish:      kitAmericanEnglishName,
		language.CanadianFrench:       kitCanadianFrenchName,
		language.Dutch:                kitDutchName,
		language.French:               kitFrenchName,
		language.German:               kitGermanName,
		language.Italian:              kitItalianName,
		language.Japanese:             kitJapaneseName,
		language.Korean:               kitKoreanName,
		language.LatinAmericanSpanish: kitLatinAmericanSpanishName,
		language.Russian:              kitRussianName,
		language.SimplifiedChinese:    kitSimplifiedChineseName,
		language.Spanish:              kitSpanishName,
		language.TraditionalChinese:   kitTraditionalChineseName}
)

var (
	kitCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: kitBirthday,
		Code:     kitCode,
		Key:      character.Kit,
		Gender:   gender.Male,
		Name:     kitName}
)

var (
	kitAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "だキョ"}

	kitCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	kitDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	kitFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	kitGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	kitItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	kitJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だキョ"}

	kitKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	kitLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	kitRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	kitSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	kitSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	kitTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	kitPhrase = nook.Languages{
		language.AmericanEnglish:      kitAmericanEnglishName,
		language.CanadianFrench:       kitCanadianFrenchName,
		language.Dutch:                kitDutchName,
		language.French:               kitFrenchName,
		language.German:               kitGermanName,
		language.Italian:              kitItalianName,
		language.Japanese:             kitJapaneseName,
		language.Korean:               kitKoreanName,
		language.LatinAmericanSpanish: kitLatinAmericanSpanishName,
		language.Russian:              kitRussianName,
		language.SimplifiedChinese:    kitSimplifiedChineseName,
		language.Spanish:              kitSpanishName,
		language.TraditionalChinese:   kitTraditionalChineseName}
)

var (
	Kit = nook.Villager{
		Character:   kitCharacter,
		Personality: personality.Jock,
		Phrase:      kitPhrase}
)
