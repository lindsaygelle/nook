package lion

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Azizzigoto"}

	azizGermanName = nook.Name{
		Language: language.German,
		Value:    "DanielRRAAHH"}

	azizItalianName = nook.Name{
		Language: language.Italian,
		Value:    "LionelloROAR"}

	azizJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アッサムカンジー"}

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
		Value:    "阿三罗尔"}

	azizSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Leonardoaborigen"}

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
		Animal:   Lion,
		Birthday: azizBirthday,
		Code:     azizCode,
		Gender:   nook.Male,
		Name:     azizName}
)

var (
	azizAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	azizCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	azizDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	azizLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	azizRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	azizSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗尔"}

	azizSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aborigen"}

	azizTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	azizPhrase = nook.Languages{
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
	Aziz = nook.Villager{
		Character:   azizCharacter,
		Personality: nook.Jock,
		Phrase:      azizPhrase}
)
