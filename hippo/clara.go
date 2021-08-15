package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	claraBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	claraCode = nook.Code{
		Value: ""}
)

var (
	claraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Clara"}

	claraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	claraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	claraFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	claraGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	claraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	claraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クララ"}

	claraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	claraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	claraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	claraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	claraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	claraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	claraName = nook.Languages{
		language.AmericanEnglish:      claraAmericanEnglishName,
		language.CanadianFrench:       claraCanadianFrenchName,
		language.Dutch:                claraDutchName,
		language.French:               claraFrenchName,
		language.German:               claraGermanName,
		language.Italian:              claraItalianName,
		language.Japanese:             claraJapaneseName,
		language.Korean:               claraKoreanName,
		language.LatinAmericanSpanish: claraLatinAmericanSpanishName,
		language.Russian:              claraRussianName,
		language.SimplifiedChinese:    claraSimplifiedChineseName,
		language.Spanish:              claraSpanishName,
		language.TraditionalChinese:   claraTraditionalChineseName}
)

var (
	claraCharacter = nook.Character{
		Animal:   Hippo,
		Birthday: claraBirthday,
		Code:     claraCode,
		Gender:   nook.Female,
		Name:     claraName}
)

var (
	claraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	claraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	claraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	claraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	claraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	claraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	claraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "うふふ"}

	claraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	claraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	claraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	claraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	claraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	claraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	claraPhrase = nook.Languages{
		language.AmericanEnglish:      claraAmericanEnglishName,
		language.CanadianFrench:       claraCanadianFrenchName,
		language.Dutch:                claraDutchName,
		language.French:               claraFrenchName,
		language.German:               claraGermanName,
		language.Italian:              claraItalianName,
		language.Japanese:             claraJapaneseName,
		language.Korean:               claraKoreanName,
		language.LatinAmericanSpanish: claraLatinAmericanSpanishName,
		language.Russian:              claraRussianName,
		language.SimplifiedChinese:    claraSimplifiedChineseName,
		language.Spanish:              claraSpanishName,
		language.TraditionalChinese:   claraTraditionalChineseName}
)

var (
	Clara = nook.Villager{
		Character:   claraCharacter,
		Personality: nook.Normal,
		Phrase:      claraPhrase}
)
