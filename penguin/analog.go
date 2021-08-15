package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	analogBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	analogCode = nook.Code{
		Value: ""}
)

var (
	analogAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Analog"}

	analogCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	analogDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	analogFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	analogGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	analogItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	analogJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アナログ"}

	analogKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	analogLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	analogRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	analogSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	analogSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	analogTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	analogName = nook.Languages{
		language.AmericanEnglish:      analogAmericanEnglishName,
		language.CanadianFrench:       analogCanadianFrenchName,
		language.Dutch:                analogDutchName,
		language.French:               analogFrenchName,
		language.German:               analogGermanName,
		language.Italian:              analogItalianName,
		language.Japanese:             analogJapaneseName,
		language.Korean:               analogKoreanName,
		language.LatinAmericanSpanish: analogLatinAmericanSpanishName,
		language.Russian:              analogRussianName,
		language.SimplifiedChinese:    analogSimplifiedChineseName,
		language.Spanish:              analogSpanishName,
		language.TraditionalChinese:   analogTraditionalChineseName}
)

var (
	analogCharacter = nook.Character{
		Animal:   Penguin,
		Birthday: analogBirthday,
		Code:     analogCode,
		Gender:   nook.Male,
		Name:     analogName}
)

var (
	analogAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "だもんで"}

	analogCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	analogDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	analogFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	analogGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	analogItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	analogJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だもんで"}

	analogKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	analogLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	analogRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	analogSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	analogSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	analogTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	analogPhrase = nook.Languages{
		language.AmericanEnglish:      analogAmericanEnglishName,
		language.CanadianFrench:       analogCanadianFrenchName,
		language.Dutch:                analogDutchName,
		language.French:               analogFrenchName,
		language.German:               analogGermanName,
		language.Italian:              analogItalianName,
		language.Japanese:             analogJapaneseName,
		language.Korean:               analogKoreanName,
		language.LatinAmericanSpanish: analogLatinAmericanSpanishName,
		language.Russian:              analogRussianName,
		language.SimplifiedChinese:    analogSimplifiedChineseName,
		language.Spanish:              analogSpanishName,
		language.TraditionalChinese:   analogTraditionalChineseName}
)

var (
	Analog = nook.Villager{
		Character:   analogCharacter,
		Personality: nook.Jock,
		Phrase:      analogPhrase}
)
