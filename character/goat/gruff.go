package goat

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
	gruffBirthday = nook.Birthday{
		Day:   29,
		Month: time.August}
)

var (
	gruffCode = nook.Code{
		Value: "goa04"}
)

var (
	gruffAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gruff"}

	gruffCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Grognon"}

	gruffDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gruff"}

	gruffFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grognon"}

	gruffGermanName = nook.Name{
		Language: language.German,
		Value:    "Gregor"}

	gruffItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Capriolé"}

	gruffJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビリー"}

	gruffKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "빌리"}

	gruffLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sátiro"}

	gruffRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Графф"}

	gruffSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "比利"}

	gruffSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sátiro"}

	gruffTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "比利"}
)

var (
	gruffName = nook.Languages{
		language.AmericanEnglish:      gruffAmericanEnglishName,
		language.CanadianFrench:       gruffCanadianFrenchName,
		language.Dutch:                gruffDutchName,
		language.French:               gruffFrenchName,
		language.German:               gruffGermanName,
		language.Italian:              gruffItalianName,
		language.Japanese:             gruffJapaneseName,
		language.Korean:               gruffKoreanName,
		language.LatinAmericanSpanish: gruffLatinAmericanSpanishName,
		language.Russian:              gruffRussianName,
		language.SimplifiedChinese:    gruffSimplifiedChineseName,
		language.Spanish:              gruffSpanishName,
		language.TraditionalChinese:   gruffTraditionalChineseName}
)

var (
	gruffCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: gruffBirthday,
		Code:     gruffCode,
		Key:      character.Gruff,
		Gender:   gender.Male,
		Name:     gruffName}
)

var (
	gruffAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bleh eh eh"}

	gruffCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hé bêêê"}

	gruffDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mekker"}

	gruffFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hé bêêê"}

	gruffGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mähmääh"}

	gruffItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bleh eh"}

	gruffJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "イェーイ"}

	gruffKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "예이예"}

	gruffLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "meeeh"}

	gruffRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ме-хе-хе"}

	gruffSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "耶"}

	gruffSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zampoña"}

	gruffTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "耶"}
)

var (
	gruffPhrase = nook.Languages{
		language.AmericanEnglish:      gruffAmericanEnglishName,
		language.CanadianFrench:       gruffCanadianFrenchName,
		language.Dutch:                gruffDutchName,
		language.French:               gruffFrenchName,
		language.German:               gruffGermanName,
		language.Italian:              gruffItalianName,
		language.Japanese:             gruffJapaneseName,
		language.Korean:               gruffKoreanName,
		language.LatinAmericanSpanish: gruffLatinAmericanSpanishName,
		language.Russian:              gruffRussianName,
		language.SimplifiedChinese:    gruffSimplifiedChineseName,
		language.Spanish:              gruffSpanishName,
		language.TraditionalChinese:   gruffTraditionalChineseName}
)

var (
	Gruff = nook.Villager{
		Character:   gruffCharacter,
		Personality: personality.Cranky,
		Phrase:      gruffPhrase}
)
