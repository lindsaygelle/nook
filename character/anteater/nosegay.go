package anteater

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
	nosegayBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	nosegayCode = nook.Code{
		Value: ""}
)

var (
	nosegayAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nosegay"}

	nosegayCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	nosegayDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	nosegayFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tamara"}

	nosegayGermanName = nook.Name{
		Language: language.German,
		Value:    "Hortense"}

	nosegayItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Narica"}

	nosegayJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アリンコ"}

	nosegayKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	nosegayLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	nosegayRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	nosegaySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小蚁"}

	nosegaySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marisa"}

	nosegayTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	nosegayName = nook.Languages{
		language.AmericanEnglish:      nosegayAmericanEnglishName,
		language.CanadianFrench:       nosegayCanadianFrenchName,
		language.Dutch:                nosegayDutchName,
		language.French:               nosegayFrenchName,
		language.German:               nosegayGermanName,
		language.Italian:              nosegayItalianName,
		language.Japanese:             nosegayJapaneseName,
		language.Korean:               nosegayKoreanName,
		language.LatinAmericanSpanish: nosegayLatinAmericanSpanishName,
		language.Russian:              nosegayRussianName,
		language.SimplifiedChinese:    nosegaySimplifiedChineseName,
		language.Spanish:              nosegaySpanishName,
		language.TraditionalChinese:   nosegayTraditionalChineseName}
)

var (
	nosegayCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: nosegayBirthday,
		Code:     nosegayCode,
		Key:      character.Nosegay,
		Gender:   gender.Female,
		Name:     nosegayName,
		Special:  false}
)

var (
	nosegayAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hoooonk"}

	nosegayCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	nosegayDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	nosegayFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hoooonk"}

	nosegayGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnauf"}

	nosegayItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hoooonk"}

	nosegayJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でアリ"}

	nosegayKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	nosegayLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	nosegayRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	nosegaySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "确实"}

	nosegaySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "achís"}

	nosegayTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	nosegayPhrase = nook.Languages{
		language.AmericanEnglish:      nosegayAmericanEnglishPhrase,
		language.CanadianFrench:       nosegayCanadianFrenchPhrase,
		language.Dutch:                nosegayDutchPhrase,
		language.French:               nosegayFrenchPhrase,
		language.German:               nosegayGermanPhrase,
		language.Italian:              nosegayItalianPhrase,
		language.Japanese:             nosegayJapanesePhrase,
		language.Korean:               nosegayKoreanPhrase,
		language.LatinAmericanSpanish: nosegayLatinAmericanSpanishPhrase,
		language.Russian:              nosegayRussianPhrase,
		language.SimplifiedChinese:    nosegaySimplifiedChinesePhrase,
		language.Spanish:              nosegaySpanishPhrase,
		language.TraditionalChinese:   nosegayTraditionalChinesePhrase}
)

var (
	Nosegay = nook.Villager{
		Character:   nosegayCharacter,
		Personality: personality.Normal,
		Phrase:      nosegayPhrase}
)
