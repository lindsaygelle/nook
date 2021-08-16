package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	chadderBirthday = nook.Birthday{
		Day:   15,
		Month: time.December}
)

var (
	chadderCode = nook.Code{
		Value: "mus18"}
)

var (
	chadderAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chadder"}

	chadderCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mozzar"}

	chadderDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chadder"}

	chadderFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mozzar"}

	chadderGermanName = nook.Name{
		Language: language.German,
		Value:    "Charlie"}

	chadderItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gruviero"}

	chadderJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チーズ"}

	chadderKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "치즈"}

	chadderLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Roque"}

	chadderRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чаддер"}

	chadderSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "起司"}

	chadderSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Roque"}

	chadderTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "起司"}
)

var (
	chadderName = nook.Languages{
		language.AmericanEnglish:      chadderAmericanEnglishName,
		language.CanadianFrench:       chadderCanadianFrenchName,
		language.Dutch:                chadderDutchName,
		language.French:               chadderFrenchName,
		language.German:               chadderGermanName,
		language.Italian:              chadderItalianName,
		language.Japanese:             chadderJapaneseName,
		language.Korean:               chadderKoreanName,
		language.LatinAmericanSpanish: chadderLatinAmericanSpanishName,
		language.Russian:              chadderRussianName,
		language.SimplifiedChinese:    chadderSimplifiedChineseName,
		language.Spanish:              chadderSpanishName,
		language.TraditionalChinese:   chadderTraditionalChineseName}
)

var (
	chadderCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: chadderBirthday,
		Code:     chadderCode,
		Gender:   gender.Male,
		Name:     chadderName}
)

var (
	chadderAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "fromage"}

	chadderCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cheese"}

	chadderDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "cheese"}

	chadderFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "frometon"}

	chadderGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hoho"}

	chadderItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "iiik"}

	chadderJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まあね"}

	chadderKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꼬리꼬리"}

	chadderLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fundifún"}

	chadderRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сырок"}

	chadderSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "也是啦"}

	chadderSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fundifún"}

	chadderTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "也是啦"}
)

var (
	chadderPhrase = nook.Languages{
		language.AmericanEnglish:      chadderAmericanEnglishName,
		language.CanadianFrench:       chadderCanadianFrenchName,
		language.Dutch:                chadderDutchName,
		language.French:               chadderFrenchName,
		language.German:               chadderGermanName,
		language.Italian:              chadderItalianName,
		language.Japanese:             chadderJapaneseName,
		language.Korean:               chadderKoreanName,
		language.LatinAmericanSpanish: chadderLatinAmericanSpanishName,
		language.Russian:              chadderRussianName,
		language.SimplifiedChinese:    chadderSimplifiedChineseName,
		language.Spanish:              chadderSpanishName,
		language.TraditionalChinese:   chadderTraditionalChineseName}
)

var (
	Chadder = nook.Villager{
		Character:   chadderCharacter,
		Personality: personality.Smug,
		Phrase:      chadderPhrase}
)
