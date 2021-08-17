package duck

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
	shinabiruBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	shinabiruCode = nook.Code{
		Value: ""}
)

var (
	shinabiruAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Shinabiru"}

	shinabiruCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	shinabiruDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	shinabiruFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	shinabiruGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	shinabiruItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	shinabiruJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シナビル"}

	shinabiruKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	shinabiruLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	shinabiruRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	shinabiruSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	shinabiruSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	shinabiruTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	shinabiruName = nook.Languages{
		language.AmericanEnglish:      shinabiruAmericanEnglishName,
		language.CanadianFrench:       shinabiruCanadianFrenchName,
		language.Dutch:                shinabiruDutchName,
		language.French:               shinabiruFrenchName,
		language.German:               shinabiruGermanName,
		language.Italian:              shinabiruItalianName,
		language.Japanese:             shinabiruJapaneseName,
		language.Korean:               shinabiruKoreanName,
		language.LatinAmericanSpanish: shinabiruLatinAmericanSpanishName,
		language.Russian:              shinabiruRussianName,
		language.SimplifiedChinese:    shinabiruSimplifiedChineseName,
		language.Spanish:              shinabiruSpanishName,
		language.TraditionalChinese:   shinabiruTraditionalChineseName}
)

var (
	shinabiruCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: shinabiruBirthday,
		Code:     shinabiruCode,
		Key:      character.Shinabiru,
		Gender:   gender.Male,
		Name:     shinabiruName}
)

var (
	shinabiruAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "シナ"}

	shinabiruCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	shinabiruDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	shinabiruFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	shinabiruGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	shinabiruItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	shinabiruJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "シナ"}

	shinabiruKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	shinabiruLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	shinabiruRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	shinabiruSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	shinabiruSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	shinabiruTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	shinabiruPhrase = nook.Languages{
		language.AmericanEnglish:      shinabiruAmericanEnglishPhrase,
		language.CanadianFrench:       shinabiruCanadianFrenchPhrase,
		language.Dutch:                shinabiruDutchPhrase,
		language.French:               shinabiruFrenchPhrase,
		language.German:               shinabiruGermanPhrase,
		language.Italian:              shinabiruItalianPhrase,
		language.Japanese:             shinabiruJapanesePhrase,
		language.Korean:               shinabiruKoreanPhrase,
		language.LatinAmericanSpanish: shinabiruLatinAmericanSpanishPhrase,
		language.Russian:              shinabiruRussianPhrase,
		language.SimplifiedChinese:    shinabiruSimplifiedChinesePhrase,
		language.Spanish:              shinabiruSpanishPhrase,
		language.TraditionalChinese:   shinabiruTraditionalChinesePhrase}
)

var (
	Shinabiru = nook.Villager{
		Character:   shinabiruCharacter,
		Personality: personality.Jock,
		Phrase:      shinabiruPhrase}
)
