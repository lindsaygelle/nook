package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	kitDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	kitFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	kitGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	kitItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	kitJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キット"}

	kitKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	kitLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	kitRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	kitSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	kitSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	kitTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Squirrel,
		Birthday: kitBirthday,
		Code:     kitCode,
		Gender:   nook.Male,
		Name:     kitName}
)

var (
	kitAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	kitCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	kitDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	kitFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	kitGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	kitItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	kitJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だキョ"}

	kitKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	kitLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	kitRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	kitSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	kitSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	kitTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Jock,
		Phrase:      kitPhrase}
)
