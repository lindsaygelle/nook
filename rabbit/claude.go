package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	claudeBirthday = nook.Birthday{
		Day:   3,
		Month: time.December}
)

var (
	claudeCode = nook.Code{
		Value: "rbt11"}
)

var (
	claudeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Claude"}

	claudeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Claudesans rire"}

	claudeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Claudehopsala"}

	claudeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Claudesans rire"}

	claudeGermanName = nook.Name{
		Language: language.German,
		Value:    "Claudehüpfauf"}

	claudeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Claudehoppela"}

	claudeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビネガーぶいぶい"}

	claudeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "비니거아으셔"}

	claudeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pablohópala"}

	claudeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клодскок-поскок"}

	claudeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿醋酸酸"}

	claudeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pablohópala"}

	claudeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿醋酸酸"}
)

var (
	claudeName = nook.Languages{
		language.AmericanEnglish:      claudeAmericanEnglishName,
		language.CanadianFrench:       claudeCanadianFrenchName,
		language.Dutch:                claudeDutchName,
		language.French:               claudeFrenchName,
		language.German:               claudeGermanName,
		language.Italian:              claudeItalianName,
		language.Japanese:             claudeJapaneseName,
		language.Korean:               claudeKoreanName,
		language.LatinAmericanSpanish: claudeLatinAmericanSpanishName,
		language.Russian:              claudeRussianName,
		language.SimplifiedChinese:    claudeSimplifiedChineseName,
		language.Spanish:              claudeSpanishName,
		language.TraditionalChinese:   claudeTraditionalChineseName}
)

var (
	claudeCharacter = nook.Character{
		Animal:   Rabbit,
		Birthday: claudeBirthday,
		Code:     claudeCode,
		Gender:   nook.Male,
		Name:     claudeName}
)

var (
	claudeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hopalong"}

	claudeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sans rire"}

	claudeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hopsala"}

	claudeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sans rire"}

	claudeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hüpfauf"}

	claudeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hoppela"}

	claudeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぶいぶい"}

	claudeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아으셔"}

	claudeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hópala"}

	claudeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "скок-поскок"}

	claudeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "酸酸"}

	claudeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "hópala"}

	claudeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "酸酸"}
)

var (
	claudePhrase = nook.Languages{
		language.AmericanEnglish:      claudeAmericanEnglishName,
		language.CanadianFrench:       claudeCanadianFrenchName,
		language.Dutch:                claudeDutchName,
		language.French:               claudeFrenchName,
		language.German:               claudeGermanName,
		language.Italian:              claudeItalianName,
		language.Japanese:             claudeJapaneseName,
		language.Korean:               claudeKoreanName,
		language.LatinAmericanSpanish: claudeLatinAmericanSpanishName,
		language.Russian:              claudeRussianName,
		language.SimplifiedChinese:    claudeSimplifiedChineseName,
		language.Spanish:              claudeSpanishName,
		language.TraditionalChinese:   claudeTraditionalChineseName}
)

var (
	Claude = nook.Villager{
		Character:   claudeCharacter,
		Personality: nook.Lazy,
		Phrase:      claudePhrase}
)
