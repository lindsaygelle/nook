package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Tamarahoooonk"}

	nosegayGermanName = nook.Name{
		Language: language.German,
		Value:    "Hortenseschnauf"}

	nosegayItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Naricahoooonk"}

	nosegayJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アリンコでアリ"}

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
		Value:    "小蚁确实"}

	nosegaySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marisaachís"}

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
		Animal:   Anteater,
		Birthday: nosegayBirthday,
		Code:     nosegayCode,
		Gender:   nook.Female,
		Name:     nosegayName}
)

var (
	nosegayAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	nosegayCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	nosegayDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	nosegayLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	nosegayRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	nosegaySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "确实"}

	nosegaySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "achís"}

	nosegayTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	nosegayPhrase = nook.Languages{
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
	Nosegay = nook.Villager{
		Character:   nosegayCharacter,
		Personality: nook.Normal,
		Phrase:      nosegayPhrase}
)
