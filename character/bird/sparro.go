package bird

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
	sparroBirthday = nook.Birthday{
		Day:   20,
		Month: time.November}
)

var (
	sparroCode = nook.Code{
		Value: "brd18"}
)

var (
	sparroAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sparro"}

	sparroCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Darius"}

	sparroDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sparro"}

	sparroFrenchName = nook.Name{
		Language: language.French,
		Value:    "Darius"}

	sparroGermanName = nook.Name{
		Language: language.German,
		Value:    "Nestor"}

	sparroItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Piumardo"}

	sparroJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちゅんのすけ"}

	sparroKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "춘섭"}

	sparroLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jaime"}

	sparroRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Спарро"}

	sparroSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "周之翔"}

	sparroSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jaime"}

	sparroTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "周之翔"}
)

var (
	sparroName = nook.Languages{
		language.AmericanEnglish:      sparroAmericanEnglishName,
		language.CanadianFrench:       sparroCanadianFrenchName,
		language.Dutch:                sparroDutchName,
		language.French:               sparroFrenchName,
		language.German:               sparroGermanName,
		language.Italian:              sparroItalianName,
		language.Japanese:             sparroJapaneseName,
		language.Korean:               sparroKoreanName,
		language.LatinAmericanSpanish: sparroLatinAmericanSpanishName,
		language.Russian:              sparroRussianName,
		language.SimplifiedChinese:    sparroSimplifiedChineseName,
		language.Spanish:              sparroSpanishName,
		language.TraditionalChinese:   sparroTraditionalChineseName}
)

var (
	sparroCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: sparroBirthday,
		Code:     sparroCode,
		Key:      character.Sparro,
		Gender:   gender.Male,
		Name:     sparroName,
		Special:  false}
)

var (
	sparroAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "like whoa"}

	sparroCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "piaaaaf"}

	sparroDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wreed"}

	sparroFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "piaaaaf"}

	sparroGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "birtz"}

	sparroItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "adiós"}

	sparroJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だチュン"}

	sparroKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "섭섭"}

	sparroLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chirpi"}

	sparroRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "так-так"}

	sparroSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啁啁"}

	sparroSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gorrión"}

	sparroTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啁啁"}
)

var (
	sparroPhrase = nook.Languages{
		language.AmericanEnglish:      sparroAmericanEnglishPhrase,
		language.CanadianFrench:       sparroCanadianFrenchPhrase,
		language.Dutch:                sparroDutchPhrase,
		language.French:               sparroFrenchPhrase,
		language.German:               sparroGermanPhrase,
		language.Italian:              sparroItalianPhrase,
		language.Japanese:             sparroJapanesePhrase,
		language.Korean:               sparroKoreanPhrase,
		language.LatinAmericanSpanish: sparroLatinAmericanSpanishPhrase,
		language.Russian:              sparroRussianPhrase,
		language.SimplifiedChinese:    sparroSimplifiedChinesePhrase,
		language.Spanish:              sparroSpanishPhrase,
		language.TraditionalChinese:   sparroTraditionalChinesePhrase}
)

var (
	Sparro = nook.Villager{
		Character:   sparroCharacter,
		Personality: personality.Jock,
		Phrase:      sparroPhrase}
)
