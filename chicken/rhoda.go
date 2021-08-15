package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

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
		Value:    "N/A"}

	rhodaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pochola"}

	rhodaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	rhodaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "布布"}

	rhodaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pochola"}

	rhodaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Chicken,
		Birthday: rhodaBirthday,
		Code:     rhodaCode,
		Gender:   nook.Female,
		Name:     rhodaName}
)

var (
	rhodaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	rhodaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Unknown"}

	rhodaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	rhodaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Unknown"}

	rhodaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	rhodaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Unknown"}

	rhodaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "Unknown"}

	rhodaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Snooty,
		Phrase:      rhodaPhrase}
)
