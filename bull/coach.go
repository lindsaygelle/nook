package bull

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	coachBirthday = nook.Birthday{
		Day:   29,
		Month: time.April}
)

var (
	coachCode = nook.Code{
		Value: "bul07"}
)

var (
	coachAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Coach"}

	coachCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Arnold"}

	coachDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Coach"}

	coachFrenchName = nook.Name{
		Language: language.French,
		Value:    "Arnold"}

	coachGermanName = nook.Name{
		Language: language.German,
		Value:    "Arnold"}

	coachItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ercole"}

	coachJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "テッチャン"}

	coachKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "철소"}

	coachLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cacho"}

	coachRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Коуч"}

	coachSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大常"}

	coachSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cacho"}

	coachTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大常"}
)

var (
	coachName = nook.Languages{
		language.AmericanEnglish:      coachAmericanEnglishName,
		language.CanadianFrench:       coachCanadianFrenchName,
		language.Dutch:                coachDutchName,
		language.French:               coachFrenchName,
		language.German:               coachGermanName,
		language.Italian:              coachItalianName,
		language.Japanese:             coachJapaneseName,
		language.Korean:               coachKoreanName,
		language.LatinAmericanSpanish: coachLatinAmericanSpanishName,
		language.Russian:              coachRussianName,
		language.SimplifiedChinese:    coachSimplifiedChineseName,
		language.Spanish:              coachSpanishName,
		language.TraditionalChinese:   coachTraditionalChineseName}
)

var (
	coachCharacter = nook.Character{
		Animal:   Bull,
		Birthday: coachBirthday,
		Code:     coachCode,
		Gender:   nook.Male,
		Name:     coachName}
)

var (
	coachAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "stubble"}

	coachCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bétail"}

	coachDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "stoppel"}

	coachFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon cou"}

	coachGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnoff"}

	coachItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oooh issa"}

	coachJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ジョリッ"}

	coachKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땅땅"}

	coachLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chacho"}

	coachRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "раз-два"}

	coachSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "胡渣"}

	coachSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chacho"}

	coachTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鬍渣"}
)

var (
	coachPhrase = nook.Languages{
		language.AmericanEnglish:      coachAmericanEnglishName,
		language.CanadianFrench:       coachCanadianFrenchName,
		language.Dutch:                coachDutchName,
		language.French:               coachFrenchName,
		language.German:               coachGermanName,
		language.Italian:              coachItalianName,
		language.Japanese:             coachJapaneseName,
		language.Korean:               coachKoreanName,
		language.LatinAmericanSpanish: coachLatinAmericanSpanishName,
		language.Russian:              coachRussianName,
		language.SimplifiedChinese:    coachSimplifiedChineseName,
		language.Spanish:              coachSpanishName,
		language.TraditionalChinese:   coachTraditionalChineseName}
)

var (
	Coach = nook.Villager{
		Character:   coachCharacter,
		Personality: nook.Jock,
		Phrase:      coachPhrase}
)
