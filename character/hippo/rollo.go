package hippo

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
		Value:    ""}

	rolloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

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
		Value:    ""}

	rolloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	rolloRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	rolloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗素"}

	rolloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hipo"}

	rolloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
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
		Animal:   animal.Hippo,
		Birthday: rolloBirthday,
		Code:     rolloCode,
		Key:      character.Rollo,
		Gender:   gender.Male,
		Name:     rolloName,
		Special:  false}
)

var (
	rolloAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "beaulch"}

	rolloCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	rolloDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

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
		Value:    ""}

	rolloLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	rolloRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	rolloSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈"}

	rolloSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "hip-hip"}

	rolloTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	rolloPhrase = nook.Languages{
		language.AmericanEnglish:      rolloAmericanEnglishPhrase,
		language.CanadianFrench:       rolloCanadianFrenchPhrase,
		language.Dutch:                rolloDutchPhrase,
		language.French:               rolloFrenchPhrase,
		language.German:               rolloGermanPhrase,
		language.Italian:              rolloItalianPhrase,
		language.Japanese:             rolloJapanesePhrase,
		language.Korean:               rolloKoreanPhrase,
		language.LatinAmericanSpanish: rolloLatinAmericanSpanishPhrase,
		language.Russian:              rolloRussianPhrase,
		language.SimplifiedChinese:    rolloSimplifiedChinesePhrase,
		language.Spanish:              rolloSpanishPhrase,
		language.TraditionalChinese:   rolloTraditionalChinesePhrase}
)

var (
	Rollo = nook.Villager{
		Character:   rolloCharacter,
		Personality: personality.Lazy,
		Phrase:      rolloPhrase}
)
