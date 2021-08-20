package cow

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
		Value:    "キャロット"}

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
		Animal:   animal.Cow,
		Birthday: carrotBirthday,
		Code:     carrotCode,
		Key:      character.Carrot,
		Gender:   gender.Female,
		Name:     carrotName,
		Special:  false}
)

var (
	carrotAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "きゃはっ"}

	carrotCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	carrotDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	carrotFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	carrotGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	carrotItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	carrotJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "きゃはっ"}

	carrotKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	carrotLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	carrotRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	carrotSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	carrotSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	carrotTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	carrotPhrase = nook.Languages{
		language.AmericanEnglish:      carrotAmericanEnglishPhrase,
		language.CanadianFrench:       carrotCanadianFrenchPhrase,
		language.Dutch:                carrotDutchPhrase,
		language.French:               carrotFrenchPhrase,
		language.German:               carrotGermanPhrase,
		language.Italian:              carrotItalianPhrase,
		language.Japanese:             carrotJapanesePhrase,
		language.Korean:               carrotKoreanPhrase,
		language.LatinAmericanSpanish: carrotLatinAmericanSpanishPhrase,
		language.Russian:              carrotRussianPhrase,
		language.SimplifiedChinese:    carrotSimplifiedChinesePhrase,
		language.Spanish:              carrotSpanishPhrase,
		language.TraditionalChinese:   carrotTraditionalChinesePhrase}
)

var (
	Carrot = nook.Villager{
		Character:   carrotCharacter,
		Personality: personality.Normal,
		Phrase:      carrotPhrase}
)
