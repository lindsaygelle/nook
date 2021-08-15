package goat

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Grognonhé bêêê"}

	gruffDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gruffmekker"}

	gruffFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grognonhé bêêê"}

	gruffGermanName = nook.Name{
		Language: language.German,
		Value:    "Gregormähmääh"}

	gruffItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Capriolébleh eh"}

	gruffJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビリーイェーイ"}

	gruffKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "빌리예이예"}

	gruffLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sátiromeeeh"}

	gruffRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Граффме-хе-хе"}

	gruffSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "比利耶"}

	gruffSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sátirozampoña"}

	gruffTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "比利耶"}
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
		Animal:   Goat,
		Birthday: gruffBirthday,
		Code:     gruffCode,
		Gender:   nook.Male,
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
		Personality: nook.Cranky,
		Phrase:      gruffPhrase}
)
