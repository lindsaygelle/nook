package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bonesBirthday = nook.Birthday{
		Day:   4,
		Month: time.August}
)

var (
	bonesCode = nook.Code{
		Value: "dog04"}
)

var (
	bonesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bones"}

	bonesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nonosyip yip"}

	bonesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bonessnuf snuf"}

	bonesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nonosyip yip"}

	bonesGermanName = nook.Name{
		Language: language.German,
		Value:    "Strolchhechel"}

	bonesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tobiayip yip"}

	bonesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トミネッ"}

	bonesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토미옙"}

	bonesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cándidoyip yip"}

	bonesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Боунстяв-тяв"}

	bonesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿富对吧"}

	bonesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cándidoyip yip"}

	bonesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿富對吧"}
)

var (
	bonesName = nook.Languages{
		language.AmericanEnglish:      bonesAmericanEnglishName,
		language.CanadianFrench:       bonesCanadianFrenchName,
		language.Dutch:                bonesDutchName,
		language.French:               bonesFrenchName,
		language.German:               bonesGermanName,
		language.Italian:              bonesItalianName,
		language.Japanese:             bonesJapaneseName,
		language.Korean:               bonesKoreanName,
		language.LatinAmericanSpanish: bonesLatinAmericanSpanishName,
		language.Russian:              bonesRussianName,
		language.SimplifiedChinese:    bonesSimplifiedChineseName,
		language.Spanish:              bonesSpanishName,
		language.TraditionalChinese:   bonesTraditionalChineseName}
)

var (
	bonesCharacter = nook.Character{
		Animal:   Dog,
		Birthday: bonesBirthday,
		Code:     bonesCode,
		Gender:   nook.Male,
		Name:     bonesName}
)

var (
	bonesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yip yip"}

	bonesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "yip yip"}

	bonesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuf snuf"}

	bonesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "yip yip"}

	bonesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hechel"}

	bonesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yip yip"}

	bonesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ネッ"}

	bonesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "옙"}

	bonesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "yip yip"}

	bonesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тяв-тяв"}

	bonesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "对吧"}

	bonesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "yip yip"}

	bonesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "對吧"}
)

var (
	bonesPhrase = nook.Languages{
		language.AmericanEnglish:      bonesAmericanEnglishName,
		language.CanadianFrench:       bonesCanadianFrenchName,
		language.Dutch:                bonesDutchName,
		language.French:               bonesFrenchName,
		language.German:               bonesGermanName,
		language.Italian:              bonesItalianName,
		language.Japanese:             bonesJapaneseName,
		language.Korean:               bonesKoreanName,
		language.LatinAmericanSpanish: bonesLatinAmericanSpanishName,
		language.Russian:              bonesRussianName,
		language.SimplifiedChinese:    bonesSimplifiedChineseName,
		language.Spanish:              bonesSpanishName,
		language.TraditionalChinese:   bonesTraditionalChineseName}
)

var (
	Bones = nook.Villager{
		Character:   bonesCharacter,
		Personality: nook.Lazy,
		Phrase:      bonesPhrase}
)
