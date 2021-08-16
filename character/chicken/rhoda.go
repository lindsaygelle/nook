package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	rhodaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	rhodaCode = nook.Code{
		Value: ""}
)

var (
	rhodaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rhoda"}

	rhodaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Arlette"}

	rhodaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	rhodaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Arlette"}

	rhodaGermanName = nook.Name{
		Language: language.German,
		Value:    "Helene"}

	rhodaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Petula"}

	rhodaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブイヨン"}

	rhodaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	rhodaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pochola"}

	rhodaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	rhodaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "布布"}

	rhodaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pochola"}

	rhodaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	rhodaName = nook.Languages{
		language.AmericanEnglish:      rhodaAmericanEnglishName,
		language.CanadianFrench:       rhodaCanadianFrenchName,
		language.Dutch:                rhodaDutchName,
		language.French:               rhodaFrenchName,
		language.German:               rhodaGermanName,
		language.Italian:              rhodaItalianName,
		language.Japanese:             rhodaJapaneseName,
		language.Korean:               rhodaKoreanName,
		language.LatinAmericanSpanish: rhodaLatinAmericanSpanishName,
		language.Russian:              rhodaRussianName,
		language.SimplifiedChinese:    rhodaSimplifiedChineseName,
		language.Spanish:              rhodaSpanishName,
		language.TraditionalChinese:   rhodaTraditionalChineseName}
)

var (
	rhodaCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: rhodaBirthday,
		Code:     rhodaCode,
		Gender:   gender.Female,
		Name:     rhodaName}
)

var (
	rhodaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "clucky"}

	rhodaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Unknown"}

	rhodaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	rhodaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "Unknown"}

	rhodaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "glucker"}

	rhodaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "skiok"}

	rhodaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "コッコ"}

	rhodaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	rhodaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Unknown"}

	rhodaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	rhodaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Unknown"}

	rhodaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "Unknown"}

	rhodaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	rhodaPhrase = nook.Languages{
		language.AmericanEnglish:      rhodaAmericanEnglishName,
		language.CanadianFrench:       rhodaCanadianFrenchName,
		language.Dutch:                rhodaDutchName,
		language.French:               rhodaFrenchName,
		language.German:               rhodaGermanName,
		language.Italian:              rhodaItalianName,
		language.Japanese:             rhodaJapaneseName,
		language.Korean:               rhodaKoreanName,
		language.LatinAmericanSpanish: rhodaLatinAmericanSpanishName,
		language.Russian:              rhodaRussianName,
		language.SimplifiedChinese:    rhodaSimplifiedChineseName,
		language.Spanish:              rhodaSpanishName,
		language.TraditionalChinese:   rhodaTraditionalChineseName}
)

var (
	Rhoda = nook.Villager{
		Character:   rhodaCharacter,
		Personality: personality.Snooty,
		Phrase:      rhodaPhrase}
)
