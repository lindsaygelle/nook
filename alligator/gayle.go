package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	gayleBirthday = nook.Birthday{
		Day:   17,
		Month: time.May}
)

var (
	gayleCode = nook.Code{
		Value: "crd07"}
)

var (
	gayleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gayle"}

	gayleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Odile"}

	gayleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gayle"}

	gayleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Odile"}

	gayleGermanName = nook.Name{
		Language: language.German,
		Value:    "Rosa"}

	gayleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Codrilla"}

	gayleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アリゲッティ"}

	gayleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "앨리"}

	gayleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Boni"}

	gayleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гейл"}

	gayleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "爱莉"}

	gayleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Boni"}

	gayleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "愛莉"}
)

var (
	gayleName = nook.Languages{
		language.AmericanEnglish:      gayleAmericanEnglishName,
		language.CanadianFrench:       gayleCanadianFrenchName,
		language.Dutch:                gayleDutchName,
		language.French:               gayleFrenchName,
		language.German:               gayleGermanName,
		language.Italian:              gayleItalianName,
		language.Japanese:             gayleJapaneseName,
		language.Korean:               gayleKoreanName,
		language.LatinAmericanSpanish: gayleLatinAmericanSpanishName,
		language.Russian:              gayleRussianName,
		language.SimplifiedChinese:    gayleSimplifiedChineseName,
		language.Spanish:              gayleSpanishName,
		language.TraditionalChinese:   gayleTraditionalChineseName}
)

var (
	gayleCharacter = nook.Character{
		Animal:   Alligator,
		Birthday: gayleBirthday,
		Code:     gayleCode,
		Gender:   nook.Female,
		Name:     gayleName}
)

var (
	gayleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snacky"}

	gayleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "croc-croc"}

	gayleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snackje"}

	gayleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croquignou"}

	gayleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zähni"}

	gayleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "désolée"}

	gayleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワニャン"}

	gayleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아거얌"}

	gayleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "smuack"}

	gayleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "перекус"}

	gayleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鳄莉"}

	gayleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "corazón"}

	gayleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鱷莉"}
)

var (
	gaylePhrase = nook.Languages{
		language.AmericanEnglish:      gayleAmericanEnglishName,
		language.CanadianFrench:       gayleCanadianFrenchName,
		language.Dutch:                gayleDutchName,
		language.French:               gayleFrenchName,
		language.German:               gayleGermanName,
		language.Italian:              gayleItalianName,
		language.Japanese:             gayleJapaneseName,
		language.Korean:               gayleKoreanName,
		language.LatinAmericanSpanish: gayleLatinAmericanSpanishName,
		language.Russian:              gayleRussianName,
		language.SimplifiedChinese:    gayleSimplifiedChineseName,
		language.Spanish:              gayleSpanishName,
		language.TraditionalChinese:   gayleTraditionalChineseName}
)

var (
	Gayle = nook.Villager{
		Character:   gayleCharacter,
		Personality: nook.Normal,
		Phrase:      gaylePhrase}
)
