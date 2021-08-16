package bull

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	vicBirthday = nook.Birthday{
		Day:   29,
		Month: time.December}
)

var (
	vicCode = nook.Code{
		Value: "bul08"}
)

var (
	vicAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Vic"}

	vicCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Toto"}

	vicDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Vic"}

	vicFrenchName = nook.Name{
		Language: language.French,
		Value:    "Toto"}

	vicGermanName = nook.Name{
		Language: language.German,
		Value:    "Klaus"}

	vicItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Paco"}

	vicJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ノルマン"}

	vicKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "노르망"}

	vicLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Artorito"}

	vicRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Вик"}

	vicSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卢尔曼"}

	vicSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Artorito"}

	vicTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "盧爾曼"}
)

var (
	vicName = nook.Languages{
		language.AmericanEnglish:      vicAmericanEnglishName,
		language.CanadianFrench:       vicCanadianFrenchName,
		language.Dutch:                vicDutchName,
		language.French:               vicFrenchName,
		language.German:               vicGermanName,
		language.Italian:              vicItalianName,
		language.Japanese:             vicJapaneseName,
		language.Korean:               vicKoreanName,
		language.LatinAmericanSpanish: vicLatinAmericanSpanishName,
		language.Russian:              vicRussianName,
		language.SimplifiedChinese:    vicSimplifiedChineseName,
		language.Spanish:              vicSpanishName,
		language.TraditionalChinese:   vicTraditionalChineseName}
)

var (
	vicCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: vicBirthday,
		Code:     vicCode,
		Gender:   gender.Male,
		Name:     vicName}
)

var (
	vicAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cud"}

	vicCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "moumoul"}

	vicDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kauw"}

	vicFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "moumoul"}

	vicGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grumpf"}

	vicItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "trullalà"}

	vicJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "モイ"}

	vicKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뭐임"}

	vicLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "brufff"}

	vicRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "полундра"}

	vicSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哞"}

	vicSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muuubién"}

	vicTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哞"}
)

var (
	vicPhrase = nook.Languages{
		language.AmericanEnglish:      vicAmericanEnglishName,
		language.CanadianFrench:       vicCanadianFrenchName,
		language.Dutch:                vicDutchName,
		language.French:               vicFrenchName,
		language.German:               vicGermanName,
		language.Italian:              vicItalianName,
		language.Japanese:             vicJapaneseName,
		language.Korean:               vicKoreanName,
		language.LatinAmericanSpanish: vicLatinAmericanSpanishName,
		language.Russian:              vicRussianName,
		language.SimplifiedChinese:    vicSimplifiedChineseName,
		language.Spanish:              vicSpanishName,
		language.TraditionalChinese:   vicTraditionalChineseName}
)

var (
	Vic = nook.Villager{
		Character:   vicCharacter,
		Personality: personality.Cranky,
		Phrase:      vicPhrase}
)
