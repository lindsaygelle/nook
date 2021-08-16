package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	isabelleBirthday = nook.Birthday{
		Day:   20,
		Month: time.December}
)

var (
	isabelleCode = nook.Code{
		Value: "sza"}
)

var (
	isabelleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Isabelle"}

	isabelleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marie"}

	isabelleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Isabelle"}

	isabelleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marie"}

	isabelleGermanName = nook.Name{
		Language: language.German,
		Value:    "Melinda"}

	isabelleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fuffi"}

	isabelleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "しずえ"}

	isabelleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "여울"}

	isabelleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Canela"}

	isabelleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Изабель"}

	isabelleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "西施惠"}

	isabelleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Canela"}

	isabelleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "西施惠"}
)

var (
	isabelleName = nook.Languages{
		language.AmericanEnglish:      isabelleAmericanEnglishName,
		language.CanadianFrench:       isabelleCanadianFrenchName,
		language.Dutch:                isabelleDutchName,
		language.French:               isabelleFrenchName,
		language.German:               isabelleGermanName,
		language.Italian:              isabelleItalianName,
		language.Japanese:             isabelleJapaneseName,
		language.Korean:               isabelleKoreanName,
		language.LatinAmericanSpanish: isabelleLatinAmericanSpanishName,
		language.Russian:              isabelleRussianName,
		language.SimplifiedChinese:    isabelleSimplifiedChineseName,
		language.Spanish:              isabelleSpanishName,
		language.TraditionalChinese:   isabelleTraditionalChineseName}
)

var (
	isabelleCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: isabelleBirthday,
		Code:     isabelleCode,
		Key:      character.Isabelle,
		Gender:   gender.Female,
		Name:     isabelleName}
)

var (
	Isabelle = nook.Resident{
		Character: isabelleCharacter}
)
