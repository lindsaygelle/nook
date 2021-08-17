package dog

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
		Value:    "Nonos"}

	bonesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bones"}

	bonesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nonos"}

	bonesGermanName = nook.Name{
		Language: language.German,
		Value:    "Strolch"}

	bonesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tobia"}

	bonesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トミ"}

	bonesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토미"}

	bonesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cándido"}

	bonesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Боунс"}

	bonesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿富"}

	bonesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cándido"}

	bonesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿富"}
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
		Animal:   animal.Dog,
		Birthday: bonesBirthday,
		Code:     bonesCode,
		Key:      character.Bones,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      bonesAmericanEnglishPhrase,
		language.CanadianFrench:       bonesCanadianFrenchPhrase,
		language.Dutch:                bonesDutchPhrase,
		language.French:               bonesFrenchPhrase,
		language.German:               bonesGermanPhrase,
		language.Italian:              bonesItalianPhrase,
		language.Japanese:             bonesJapanesePhrase,
		language.Korean:               bonesKoreanPhrase,
		language.LatinAmericanSpanish: bonesLatinAmericanSpanishPhrase,
		language.Russian:              bonesRussianPhrase,
		language.SimplifiedChinese:    bonesSimplifiedChinesePhrase,
		language.Spanish:              bonesSpanishPhrase,
		language.TraditionalChinese:   bonesTraditionalChinesePhrase}
)

var (
	Bones = nook.Villager{
		Character:   bonesCharacter,
		Personality: personality.Lazy,
		Phrase:      bonesPhrase}
)
