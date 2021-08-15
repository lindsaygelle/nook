package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rolloBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	rolloCode = nook.Code{
		Value: ""}
)

var (
	rolloAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rollo"}

	rolloCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	rolloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	rolloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hippo"}

	rolloGermanName = nook.Name{
		Language: language.German,
		Value:    "Winfried"}

	rolloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rollo"}

	rolloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヒポクラテス"}

	rolloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	rolloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	rolloRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	rolloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗素"}

	rolloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hipo"}

	rolloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	rolloName = nook.Languages{
		language.AmericanEnglish:      rolloAmericanEnglishName,
		language.CanadianFrench:       rolloCanadianFrenchName,
		language.Dutch:                rolloDutchName,
		language.French:               rolloFrenchName,
		language.German:               rolloGermanName,
		language.Italian:              rolloItalianName,
		language.Japanese:             rolloJapaneseName,
		language.Korean:               rolloKoreanName,
		language.LatinAmericanSpanish: rolloLatinAmericanSpanishName,
		language.Russian:              rolloRussianName,
		language.SimplifiedChinese:    rolloSimplifiedChineseName,
		language.Spanish:              rolloSpanishName,
		language.TraditionalChinese:   rolloTraditionalChineseName}
)

var (
	rolloCharacter = nook.Character{
		Animal:   Hippo,
		Birthday: rolloBirthday,
		Code:     rolloCode,
		Gender:   nook.Male,
		Name:     rolloName}
)

var (
	rolloAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	rolloCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	rolloDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	rolloFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "que diable"}

	rolloGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "rülps"}

	rolloItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "trippo"}

	rolloJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのな"}

	rolloKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	rolloLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	rolloRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	rolloSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈"}

	rolloSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "hip-hip"}

	rolloTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	rolloPhrase = nook.Languages{
		language.AmericanEnglish:      rolloAmericanEnglishName,
		language.CanadianFrench:       rolloCanadianFrenchName,
		language.Dutch:                rolloDutchName,
		language.French:               rolloFrenchName,
		language.German:               rolloGermanName,
		language.Italian:              rolloItalianName,
		language.Japanese:             rolloJapaneseName,
		language.Korean:               rolloKoreanName,
		language.LatinAmericanSpanish: rolloLatinAmericanSpanishName,
		language.Russian:              rolloRussianName,
		language.SimplifiedChinese:    rolloSimplifiedChineseName,
		language.Spanish:              rolloSpanishName,
		language.TraditionalChinese:   rolloTraditionalChineseName}
)

var (
	Rollo = nook.Villager{
		Character:   rolloCharacter,
		Personality: nook.Lazy,
		Phrase:      rolloPhrase}
)
