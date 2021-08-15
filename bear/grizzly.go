package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	grizzlyBirthday = nook.Birthday{
		Day:   31,
		Month: time.July}
)

var (
	grizzlyCode = nook.Code{
		Value: "bea09"}
)

var (
	grizzlyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Grizzly"}

	grizzlyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Grizzlycanaillou"}

	grizzlyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Grizzlygrrr"}

	grizzlyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grizzlycanaillou"}

	grizzlyGermanName = nook.Name{
		Language: language.German,
		Value:    "Geraldgrrr"}

	grizzlyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grizzlygrrr"}

	grizzlyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ムーだクマ"}

	grizzlyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "무뚝냅둬"}

	grizzlyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gruñóngrrre"}

	grizzlyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гризлигрр"}

	grizzlySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "穆穆熊熊"}

	grizzlySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gruñóntostón"}

	grizzlyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "穆穆熊熊"}
)

var (
	grizzlyName = nook.Languages{
		language.AmericanEnglish:      grizzlyAmericanEnglishName,
		language.CanadianFrench:       grizzlyCanadianFrenchName,
		language.Dutch:                grizzlyDutchName,
		language.French:               grizzlyFrenchName,
		language.German:               grizzlyGermanName,
		language.Italian:              grizzlyItalianName,
		language.Japanese:             grizzlyJapaneseName,
		language.Korean:               grizzlyKoreanName,
		language.LatinAmericanSpanish: grizzlyLatinAmericanSpanishName,
		language.Russian:              grizzlyRussianName,
		language.SimplifiedChinese:    grizzlySimplifiedChineseName,
		language.Spanish:              grizzlySpanishName,
		language.TraditionalChinese:   grizzlyTraditionalChineseName}
)

var (
	grizzlyCharacter = nook.Character{
		Animal:   Bear,
		Birthday: grizzlyBirthday,
		Code:     grizzlyCode,
		Gender:   nook.Male,
		Name:     grizzlyName}
)

var (
	grizzlyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grrr"}

	grizzlyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "canaillou"}

	grizzlyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "grrr"}

	grizzlyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "canaillou"}

	grizzlyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grrr"}

	grizzlyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrr"}

	grizzlyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だクマ"}

	grizzlyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냅둬"}

	grizzlyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrre"}

	grizzlyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "грр"}

	grizzlySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "熊熊"}

	grizzlySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tostón"}

	grizzlyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "熊熊"}
)

var (
	grizzlyPhrase = nook.Languages{
		language.AmericanEnglish:      grizzlyAmericanEnglishName,
		language.CanadianFrench:       grizzlyCanadianFrenchName,
		language.Dutch:                grizzlyDutchName,
		language.French:               grizzlyFrenchName,
		language.German:               grizzlyGermanName,
		language.Italian:              grizzlyItalianName,
		language.Japanese:             grizzlyJapaneseName,
		language.Korean:               grizzlyKoreanName,
		language.LatinAmericanSpanish: grizzlyLatinAmericanSpanishName,
		language.Russian:              grizzlyRussianName,
		language.SimplifiedChinese:    grizzlySimplifiedChineseName,
		language.Spanish:              grizzlySpanishName,
		language.TraditionalChinese:   grizzlyTraditionalChineseName}
)

var (
	Grizzly = nook.Villager{
		Character:   grizzlyCharacter,
		Personality: nook.Cranky,
		Phrase:      grizzlyPhrase}
)
