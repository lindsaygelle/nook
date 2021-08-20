package bull

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
		Value:    "ウェルダン"}

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
		Animal:   animal.Bull,
		Birthday: weldonBirthday,
		Code:     weldonCode,
		Key:      character.Weldon,
		Gender:   gender.Male,
		Name:     weldonName,
		Special:  false}
)

var (
	weldonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ムーチョ"}

	weldonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	weldonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	weldonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	weldonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	weldonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	weldonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ムーチョ"}

	weldonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	weldonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	weldonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	weldonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	weldonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	weldonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	weldonPhrase = nook.Languages{
		language.AmericanEnglish:      weldonAmericanEnglishPhrase,
		language.CanadianFrench:       weldonCanadianFrenchPhrase,
		language.Dutch:                weldonDutchPhrase,
		language.French:               weldonFrenchPhrase,
		language.German:               weldonGermanPhrase,
		language.Italian:              weldonItalianPhrase,
		language.Japanese:             weldonJapanesePhrase,
		language.Korean:               weldonKoreanPhrase,
		language.LatinAmericanSpanish: weldonLatinAmericanSpanishPhrase,
		language.Russian:              weldonRussianPhrase,
		language.SimplifiedChinese:    weldonSimplifiedChinesePhrase,
		language.Spanish:              weldonSpanishPhrase,
		language.TraditionalChinese:   weldonTraditionalChinesePhrase}
)

var (
	Weldon = nook.Villager{
		Character:   weldonCharacter,
		Personality: personality.Cranky,
		Phrase:      weldonPhrase}
)
