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
	// digbyBirthday represents digby birthday.
	digbyBirthday = nook.Birthday{
		Day:   20,
		Month: time.December}
)

var (
	// digbyCode represents digby code.
	digbyCode = nook.Code{
		Value: "szo"}
)

var (
	// digbyAmericanEnglishName represents digby american english name.
	digbyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Digby"}

	// digbyCanadianFrenchName represents digby canadian french name.
	digbyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Max"}

	// digbyDutchName represents digby dutch name.
	digbyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Digby"}

	// digbyFrenchName represents digby french name.
	digbyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Max"}

	// digbyGermanName represents digby german name.
	digbyGermanName = nook.Name{
		Language: language.German,
		Value:    "Moritz"}

	// digbyItalianName represents digby italian name.
	digbyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fofò"}

	// digbyJapaneseName represents digby japanese name.
	digbyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ケント"}

	// digbyKoreanName represents digby korean name.
	digbyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "켄트"}

	// digbyLatinAmericanSpanishName represents digby latin american spanish name.
	digbyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Candrés"}

	// digbyRussianName represents digby russian name.
	digbyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дигби"}

	// digbySimplifiedChineseName represents digby simplified chinese name.
	digbySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "西施德"}

	// digbySpanishName represents digby spanish name.
	digbySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Candrés"}

	// digbyTraditionalChineseName represents digby traditional chinese name.
	digbyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "西施德"}
)

var (
	// digbyName represents digby name.
	digbyName = nook.Languages{
		language.AmericanEnglish:      digbyAmericanEnglishName,
		language.CanadianFrench:       digbyCanadianFrenchName,
		language.Dutch:                digbyDutchName,
		language.French:               digbyFrenchName,
		language.German:               digbyGermanName,
		language.Italian:              digbyItalianName,
		language.Japanese:             digbyJapaneseName,
		language.Korean:               digbyKoreanName,
		language.LatinAmericanSpanish: digbyLatinAmericanSpanishName,
		language.Russian:              digbyRussianName,
		language.SimplifiedChinese:    digbySimplifiedChineseName,
		language.Spanish:              digbySpanishName,
		language.TraditionalChinese:   digbyTraditionalChineseName}
)

var (
	// digbyCharacter represents digby character.
	digbyCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: digbyBirthday,
		Code:     digbyCode,
		Key:      character.Digby,
		Gender:   gender.Male,
		Name:     digbyName,
		Special:  true}
)

var (
	// Digby represents digby.
	Digby = nook.Resident{
		Character: digbyCharacter}
)
