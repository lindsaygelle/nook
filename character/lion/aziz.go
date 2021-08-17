package lion

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
	azizBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	azizCode = nook.Code{
		Value: ""}
)

var (
	azizAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Aziz"}

	azizCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	azizDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	azizFrenchName = nook.Name{
		Language: language.French,
		Value:    "Aziz"}

	azizGermanName = nook.Name{
		Language: language.German,
		Value:    "Daniel"}

	azizItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lionello"}

	azizJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アッサム"}

	azizKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	azizLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	azizRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	azizSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿三"}

	azizSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Leonardo"}

	azizTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	azizName = nook.Languages{
		language.AmericanEnglish:      azizAmericanEnglishName,
		language.CanadianFrench:       azizCanadianFrenchName,
		language.Dutch:                azizDutchName,
		language.French:               azizFrenchName,
		language.German:               azizGermanName,
		language.Italian:              azizItalianName,
		language.Japanese:             azizJapaneseName,
		language.Korean:               azizKoreanName,
		language.LatinAmericanSpanish: azizLatinAmericanSpanishName,
		language.Russian:              azizRussianName,
		language.SimplifiedChinese:    azizSimplifiedChineseName,
		language.Spanish:              azizSpanishName,
		language.TraditionalChinese:   azizTraditionalChineseName}
)

var (
	azizCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: azizBirthday,
		Code:     azizCode,
		Key:      character.Aziz,
		Gender:   gender.Male,
		Name:     azizName}
)

var (
	azizAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "RAWR"}

	azizCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	azizDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	azizFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zigoto"}

	azizGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "RRAAHH"}

	azizItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ROAR"}

	azizJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "カンジー"}

	azizKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	azizLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	azizRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	azizSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗尔"}

	azizSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aborigen"}

	azizTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	azizPhrase = nook.Languages{
		language.AmericanEnglish:      azizAmericanEnglishPhrase,
		language.CanadianFrench:       azizCanadianFrenchPhrase,
		language.Dutch:                azizDutchPhrase,
		language.French:               azizFrenchPhrase,
		language.German:               azizGermanPhrase,
		language.Italian:              azizItalianPhrase,
		language.Japanese:             azizJapanesePhrase,
		language.Korean:               azizKoreanPhrase,
		language.LatinAmericanSpanish: azizLatinAmericanSpanishPhrase,
		language.Russian:              azizRussianPhrase,
		language.SimplifiedChinese:    azizSimplifiedChinesePhrase,
		language.Spanish:              azizSpanishPhrase,
		language.TraditionalChinese:   azizTraditionalChinesePhrase}
)

var (
	Aziz = nook.Villager{
		Character:   azizCharacter,
		Personality: personality.Jock,
		Phrase:      azizPhrase}
)
