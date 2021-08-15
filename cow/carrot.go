package cow

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	carrotBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	carrotCode = nook.Code{
		Value: ""}
)

var (
	carrotAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Carrot"}

	carrotCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	carrotDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	carrotFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	carrotGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	carrotItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	carrotJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャロットきゃはっ"}

	carrotKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	carrotLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	carrotRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	carrotSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	carrotSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	carrotTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	carrotName = nook.Languages{
		language.AmericanEnglish:      carrotAmericanEnglishName,
		language.CanadianFrench:       carrotCanadianFrenchName,
		language.Dutch:                carrotDutchName,
		language.French:               carrotFrenchName,
		language.German:               carrotGermanName,
		language.Italian:              carrotItalianName,
		language.Japanese:             carrotJapaneseName,
		language.Korean:               carrotKoreanName,
		language.LatinAmericanSpanish: carrotLatinAmericanSpanishName,
		language.Russian:              carrotRussianName,
		language.SimplifiedChinese:    carrotSimplifiedChineseName,
		language.Spanish:              carrotSpanishName,
		language.TraditionalChinese:   carrotTraditionalChineseName}
)

var (
	carrotCharacter = nook.Character{
		Animal:   Cow,
		Birthday: carrotBirthday,
		Code:     carrotCode,
		Gender:   nook.Female,
		Name:     carrotName}
)

var (
	carrotAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	carrotCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	carrotDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	carrotFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	carrotGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	carrotItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	carrotJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "きゃはっ"}

	carrotKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	carrotLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	carrotRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	carrotSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	carrotSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	carrotTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	carrotPhrase = nook.Languages{
		language.AmericanEnglish:      carrotAmericanEnglishName,
		language.CanadianFrench:       carrotCanadianFrenchName,
		language.Dutch:                carrotDutchName,
		language.French:               carrotFrenchName,
		language.German:               carrotGermanName,
		language.Italian:              carrotItalianName,
		language.Japanese:             carrotJapaneseName,
		language.Korean:               carrotKoreanName,
		language.LatinAmericanSpanish: carrotLatinAmericanSpanishName,
		language.Russian:              carrotRussianName,
		language.SimplifiedChinese:    carrotSimplifiedChineseName,
		language.Spanish:              carrotSpanishName,
		language.TraditionalChinese:   carrotTraditionalChineseName}
)

var (
	Carrot = nook.Villager{
		Character:   carrotCharacter,
		Personality: nook.Normal,
		Phrase:      carrotPhrase}
)
