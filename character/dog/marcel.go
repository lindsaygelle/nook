package dog

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
	marcelBirthday = nook.Birthday{
		Day:   31,
		Month: time.December}
)

var (
	marcelCode = nook.Code{
		Value: "dog15"}
)

var (
	marcelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Marcel"}

	marcelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ismaël"}

	marcelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marcel"}

	marcelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ismaël"}

	marcelGermanName = nook.Name{
		Language: language.German,
		Value:    "Ronaldo"}

	marcelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giosuè"}

	marcelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "もんじゃ"}

	marcelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에드워드"}

	marcelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eto"}

	marcelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Марсель"}

	marcelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "文字烧"}

	marcelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eto"}

	marcelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "文字燒"}
)

var (
	marcelName = nook.Languages{
		language.AmericanEnglish:      marcelAmericanEnglishName,
		language.CanadianFrench:       marcelCanadianFrenchName,
		language.Dutch:                marcelDutchName,
		language.French:               marcelFrenchName,
		language.German:               marcelGermanName,
		language.Italian:              marcelItalianName,
		language.Japanese:             marcelJapaneseName,
		language.Korean:               marcelKoreanName,
		language.LatinAmericanSpanish: marcelLatinAmericanSpanishName,
		language.Russian:              marcelRussianName,
		language.SimplifiedChinese:    marcelSimplifiedChineseName,
		language.Spanish:              marcelSpanishName,
		language.TraditionalChinese:   marcelTraditionalChineseName}
)

var (
	marcelCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: marcelBirthday,
		Code:     marcelCode,
		Key:      character.Marcel,
		Gender:   gender.Male,
		Name:     marcelName}
)

var (
	marcelAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "non"}

	marcelCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miaouf"}

	marcelDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mais non"}

	marcelFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miaouf"}

	marcelGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tusch"}

	marcelItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snif snif"}

	marcelJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なんじゃ"}

	marcelKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하옵소서"}

	marcelLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nenonó"}

	marcelRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ау-уи"}

	marcelSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "什么字"}

	marcelSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "colmillos"}

	marcelTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "什麼字"}
)

var (
	marcelPhrase = nook.Languages{
		language.AmericanEnglish:      marcelAmericanEnglishPhrase,
		language.CanadianFrench:       marcelCanadianFrenchPhrase,
		language.Dutch:                marcelDutchPhrase,
		language.French:               marcelFrenchPhrase,
		language.German:               marcelGermanPhrase,
		language.Italian:              marcelItalianPhrase,
		language.Japanese:             marcelJapanesePhrase,
		language.Korean:               marcelKoreanPhrase,
		language.LatinAmericanSpanish: marcelLatinAmericanSpanishPhrase,
		language.Russian:              marcelRussianPhrase,
		language.SimplifiedChinese:    marcelSimplifiedChinesePhrase,
		language.Spanish:              marcelSpanishPhrase,
		language.TraditionalChinese:   marcelTraditionalChinesePhrase}
)

var (
	Marcel = nook.Villager{
		Character:   marcelCharacter,
		Personality: personality.Lazy,
		Phrase:      marcelPhrase}
)
