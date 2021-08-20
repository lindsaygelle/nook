package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	carloBirthday = nook.Birthday{
		Day:   3,
		Month: time.May}
)

var (
	carloCode = nook.Code{
		Value: "cwb"}
)

var (
	carloAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Carlo"}

	carloCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Plumac"}

	carloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	carloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Plumac"}

	carloGermanName = nook.Name{
		Language: language.German,
		Value:    "Ralle"}

	carloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giorvo"}

	carloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グーさん"}

	carloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "카를로"}

	carloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Graznel"}

	carloRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	carloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	carloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Graznel"}

	carloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "營長"}
)

var (
	carloName = nook.Languages{
		language.AmericanEnglish:      carloAmericanEnglishName,
		language.CanadianFrench:       carloCanadianFrenchName,
		language.Dutch:                carloDutchName,
		language.French:               carloFrenchName,
		language.German:               carloGermanName,
		language.Italian:              carloItalianName,
		language.Japanese:             carloJapaneseName,
		language.Korean:               carloKoreanName,
		language.LatinAmericanSpanish: carloLatinAmericanSpanishName,
		language.Russian:              carloRussianName,
		language.SimplifiedChinese:    carloSimplifiedChineseName,
		language.Spanish:              carloSpanishName,
		language.TraditionalChinese:   carloTraditionalChineseName}
)

var (
	carloCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: carloBirthday,
		Code:     carloCode,
		Key:      character.Carlo,
		Gender:   gender.Male,
		Name:     carloName,
		Special:  true}
)

var (
	Carlo = nook.Resident{
		Character: carloCharacter}
)
