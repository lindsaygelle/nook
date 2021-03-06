package pig

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
		Value:    "Bonno"}

	hughDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hugh"}

	hughFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bonno"}

	hughGermanName = nook.Name{
		Language: language.German,
		Value:    "Hugo"}

	hughItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jambon"}

	hughJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クッチャネ"}

	hughKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "먹고파"}

	hughLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jacobo"}

	hughRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хью"}

	hughSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿猪"}

	hughSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jacobo"}

	hughTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿豬"}
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
		Animal:   animal.Pig,
		Birthday: hughBirthday,
		Code:     hughCode,
		Key:      character.Hugh,
		Gender:   gender.Male,
		Name:     hughName,
		Special:  false}
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
		language.AmericanEnglish:      hughAmericanEnglishPhrase,
		language.CanadianFrench:       hughCanadianFrenchPhrase,
		language.Dutch:                hughDutchPhrase,
		language.French:               hughFrenchPhrase,
		language.German:               hughGermanPhrase,
		language.Italian:              hughItalianPhrase,
		language.Japanese:             hughJapanesePhrase,
		language.Korean:               hughKoreanPhrase,
		language.LatinAmericanSpanish: hughLatinAmericanSpanishPhrase,
		language.Russian:              hughRussianPhrase,
		language.SimplifiedChinese:    hughSimplifiedChinesePhrase,
		language.Spanish:              hughSpanishPhrase,
		language.TraditionalChinese:   hughTraditionalChinesePhrase}
)

var (
	Hugh = nook.Villager{
		Character:   hughCharacter,
		Personality: personality.Lazy,
		Phrase:      hughPhrase}
)
