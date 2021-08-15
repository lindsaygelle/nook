package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bellaBirthday = nook.Birthday{
		Day:   28,
		Month: time.December}
)

var (
	bellaCode = nook.Code{
		Value: "mus02"}
)

var (
	bellaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bella"}

	bellaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Belle"}

	bellaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bella"}

	bellaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Belle"}

	bellaGermanName = nook.Name{
		Language: language.German,
		Value:    "Susi"}

	bellaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bella"}

	bellaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "イザベラ"}

	bellaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "이자벨"}

	bellaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Prity"}

	bellaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Белла"}

	bellaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贝拉"}

	bellaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Prity"}

	bellaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "貝拉"}
)

var (
	bellaName = nook.Languages{
		language.AmericanEnglish:      bellaAmericanEnglishName,
		language.CanadianFrench:       bellaCanadianFrenchName,
		language.Dutch:                bellaDutchName,
		language.French:               bellaFrenchName,
		language.German:               bellaGermanName,
		language.Italian:              bellaItalianName,
		language.Japanese:             bellaJapaneseName,
		language.Korean:               bellaKoreanName,
		language.LatinAmericanSpanish: bellaLatinAmericanSpanishName,
		language.Russian:              bellaRussianName,
		language.SimplifiedChinese:    bellaSimplifiedChineseName,
		language.Spanish:              bellaSpanishName,
		language.TraditionalChinese:   bellaTraditionalChineseName}
)

var (
	bellaCharacter = nook.Character{
		Animal:   Mouse,
		Birthday: bellaBirthday,
		Code:     bellaCode,
		Gender:   nook.Female,
		Name:     bellaName}
)

var (
	bellaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "eeks"}

	bellaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hiiiii"}

	bellaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ieps"}

	bellaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hiiiii"}

	bellaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "visavisa"}

	bellaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "eeks"}

	bellaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ギャハッ"}

	bellaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "캬학"}

	bellaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "petisú"}

	bellaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пип"}

	bellaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘎哈"}

	bellaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "petisú"}

	bellaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘎哈"}
)

var (
	bellaPhrase = nook.Languages{
		language.AmericanEnglish:      bellaAmericanEnglishName,
		language.CanadianFrench:       bellaCanadianFrenchName,
		language.Dutch:                bellaDutchName,
		language.French:               bellaFrenchName,
		language.German:               bellaGermanName,
		language.Italian:              bellaItalianName,
		language.Japanese:             bellaJapaneseName,
		language.Korean:               bellaKoreanName,
		language.LatinAmericanSpanish: bellaLatinAmericanSpanishName,
		language.Russian:              bellaRussianName,
		language.SimplifiedChinese:    bellaSimplifiedChineseName,
		language.Spanish:              bellaSpanishName,
		language.TraditionalChinese:   bellaTraditionalChineseName}
)

var (
	Bella = nook.Villager{
		Character:   bellaCharacter,
		Personality: nook.Peppy,
		Phrase:      bellaPhrase}
)
