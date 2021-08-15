package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bowBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	bowCode = nook.Code{
		Value: ""}
)

var (
	bowAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bow"}

	bowCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	bowDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	bowFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	bowGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	bowItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	bowJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バウバウ"}

	bowKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	bowLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	bowRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	bowSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	bowSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	bowTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	bowName = nook.Languages{
		language.AmericanEnglish:      bowAmericanEnglishName,
		language.CanadianFrench:       bowCanadianFrenchName,
		language.Dutch:                bowDutchName,
		language.French:               bowFrenchName,
		language.German:               bowGermanName,
		language.Italian:              bowItalianName,
		language.Japanese:             bowJapaneseName,
		language.Korean:               bowKoreanName,
		language.LatinAmericanSpanish: bowLatinAmericanSpanishName,
		language.Russian:              bowRussianName,
		language.SimplifiedChinese:    bowSimplifiedChineseName,
		language.Spanish:              bowSpanishName,
		language.TraditionalChinese:   bowTraditionalChineseName}
)

var (
	bowCharacter = nook.Character{
		Animal:   Dog,
		Birthday: bowBirthday,
		Code:     bowCode,
		Gender:   nook.Male,
		Name:     bowName}
)

var (
	bowAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	bowCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	bowDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	bowFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	bowGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	bowItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	bowJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "バウ"}

	bowKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	bowLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	bowRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	bowSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	bowSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	bowTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	bowPhrase = nook.Languages{
		language.AmericanEnglish:      bowAmericanEnglishName,
		language.CanadianFrench:       bowCanadianFrenchName,
		language.Dutch:                bowDutchName,
		language.French:               bowFrenchName,
		language.German:               bowGermanName,
		language.Italian:              bowItalianName,
		language.Japanese:             bowJapaneseName,
		language.Korean:               bowKoreanName,
		language.LatinAmericanSpanish: bowLatinAmericanSpanishName,
		language.Russian:              bowRussianName,
		language.SimplifiedChinese:    bowSimplifiedChineseName,
		language.Spanish:              bowSpanishName,
		language.TraditionalChinese:   bowTraditionalChineseName}
)

var (
	Bow = nook.Villager{
		Character:   bowCharacter,
		Personality: nook.Lazy,
		Phrase:      bowPhrase}
)
