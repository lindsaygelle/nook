package bull

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	weldonBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	weldonCode = nook.Code{
		Value: ""}
)

var (
	weldonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Weldon"}

	weldonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	weldonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	weldonFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	weldonGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	weldonItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	weldonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ウェルダンムーチョ"}

	weldonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	weldonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	weldonRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	weldonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	weldonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	weldonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	weldonName = nook.Languages{
		language.AmericanEnglish:      weldonAmericanEnglishName,
		language.CanadianFrench:       weldonCanadianFrenchName,
		language.Dutch:                weldonDutchName,
		language.French:               weldonFrenchName,
		language.German:               weldonGermanName,
		language.Italian:              weldonItalianName,
		language.Japanese:             weldonJapaneseName,
		language.Korean:               weldonKoreanName,
		language.LatinAmericanSpanish: weldonLatinAmericanSpanishName,
		language.Russian:              weldonRussianName,
		language.SimplifiedChinese:    weldonSimplifiedChineseName,
		language.Spanish:              weldonSpanishName,
		language.TraditionalChinese:   weldonTraditionalChineseName}
)

var (
	weldonCharacter = nook.Character{
		Animal:   Bull,
		Birthday: weldonBirthday,
		Code:     weldonCode,
		Gender:   nook.Male,
		Name:     weldonName}
)

var (
	weldonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	weldonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	weldonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	weldonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	weldonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	weldonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	weldonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ムーチョ"}

	weldonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	weldonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	weldonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	weldonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	weldonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	weldonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	weldonPhrase = nook.Languages{
		language.AmericanEnglish:      weldonAmericanEnglishName,
		language.CanadianFrench:       weldonCanadianFrenchName,
		language.Dutch:                weldonDutchName,
		language.French:               weldonFrenchName,
		language.German:               weldonGermanName,
		language.Italian:              weldonItalianName,
		language.Japanese:             weldonJapaneseName,
		language.Korean:               weldonKoreanName,
		language.LatinAmericanSpanish: weldonLatinAmericanSpanishName,
		language.Russian:              weldonRussianName,
		language.SimplifiedChinese:    weldonSimplifiedChineseName,
		language.Spanish:              weldonSpanishName,
		language.TraditionalChinese:   weldonTraditionalChineseName}
)

var (
	Weldon = nook.Villager{
		Character:   weldonCharacter,
		Personality: nook.Cranky,
		Phrase:      weldonPhrase}
)
