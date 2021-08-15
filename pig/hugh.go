package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	hughBirthday = nook.Birthday{
		Day:   30,
		Month: time.December}
)

var (
	hughCode = nook.Code{
		Value: "pig03"}
)

var (
	hughAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hugh"}

	hughCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bonnocuick"}

	hughDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hughsnuif"}

	hughFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bonnocuick"}

	hughGermanName = nook.Name{
		Language: language.German,
		Value:    "Hugoschniff"}

	hughItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jambonsnort"}

	hughJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クッチャネとかね"}

	hughKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "먹고파아님말구"}

	hughLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jacobogroink"}

	hughRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хьюхрюк-хрюк"}

	hughSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿猪懒懒"}

	hughSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jacobogroink"}

	hughTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿豬懶懶"}
)

var (
	hughName = nook.Languages{
		language.AmericanEnglish:      hughAmericanEnglishName,
		language.CanadianFrench:       hughCanadianFrenchName,
		language.Dutch:                hughDutchName,
		language.French:               hughFrenchName,
		language.German:               hughGermanName,
		language.Italian:              hughItalianName,
		language.Japanese:             hughJapaneseName,
		language.Korean:               hughKoreanName,
		language.LatinAmericanSpanish: hughLatinAmericanSpanishName,
		language.Russian:              hughRussianName,
		language.SimplifiedChinese:    hughSimplifiedChineseName,
		language.Spanish:              hughSpanishName,
		language.TraditionalChinese:   hughTraditionalChineseName}
)

var (
	hughCharacter = nook.Character{
		Animal:   Pig,
		Birthday: hughBirthday,
		Code:     hughCode,
		Gender:   nook.Male,
		Name:     hughName}
)

var (
	hughAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snortle"}

	hughCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cuick"}

	hughDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuif"}

	hughFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cuick"}

	hughGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schniff"}

	hughItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snort"}

	hughJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とかね"}

	hughKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아님말구"}

	hughLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "groink"}

	hughRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюк-хрюк"}

	hughSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "懒懒"}

	hughSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "groink"}

	hughTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "懶懶"}
)

var (
	hughPhrase = nook.Languages{
		language.AmericanEnglish:      hughAmericanEnglishName,
		language.CanadianFrench:       hughCanadianFrenchName,
		language.Dutch:                hughDutchName,
		language.French:               hughFrenchName,
		language.German:               hughGermanName,
		language.Italian:              hughItalianName,
		language.Japanese:             hughJapaneseName,
		language.Korean:               hughKoreanName,
		language.LatinAmericanSpanish: hughLatinAmericanSpanishName,
		language.Russian:              hughRussianName,
		language.SimplifiedChinese:    hughSimplifiedChineseName,
		language.Spanish:              hughSpanishName,
		language.TraditionalChinese:   hughTraditionalChineseName}
)

var (
	Hugh = nook.Villager{
		Character:   hughCharacter,
		Personality: nook.Lazy,
		Phrase:      hughPhrase}
)
