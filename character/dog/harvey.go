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
	// harveyBirthday represents harvey birthday.
	harveyBirthday = nook.Birthday{
		Day:   2,
		Month: time.August}
)

var (
	// harveyCode represents harvey code.
	harveyCode = nook.Code{
		Value: "spn"}
)

var (
	// harveyAmericanEnglishName represents harvey american english name.
	harveyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Harvey"}

	// harveyCanadianFrenchName represents harvey canadian french name.
	harveyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Joe"}

	// harveyDutchName represents harvey dutch name.
	harveyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Harvey"}

	// harveyFrenchName represents harvey french name.
	harveyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Joe"}

	// harveyGermanName represents harvey german name.
	harveyGermanName = nook.Name{
		Language: language.German,
		Value:    "Harvey"}

	// harveyItalianName represents harvey italian name.
	harveyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fiorilio"}

	// harveyJapaneseName represents harvey japanese name.
	harveyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パニエル"}

	// harveyKoreanName represents harvey korean name.
	harveyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파니엘"}

	// harveyLatinAmericanSpanishName represents harvey latin american spanish name.
	harveyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fauno"}

	// harveyRussianName represents harvey russian name.
	harveyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Харви"}

	// harveySimplifiedChineseName represents harvey simplified chinese name.
	harveySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巴猎"}

	// harveySpanishName represents harvey spanish name.
	harveySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fauno"}

	// harveyTraditionalChineseName represents harvey traditional chinese name.
	harveyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巴獵"}
)

var (
	// harveyName represents harvey name.
	harveyName = nook.Languages{
		language.AmericanEnglish:      harveyAmericanEnglishName,
		language.CanadianFrench:       harveyCanadianFrenchName,
		language.Dutch:                harveyDutchName,
		language.French:               harveyFrenchName,
		language.German:               harveyGermanName,
		language.Italian:              harveyItalianName,
		language.Japanese:             harveyJapaneseName,
		language.Korean:               harveyKoreanName,
		language.LatinAmericanSpanish: harveyLatinAmericanSpanishName,
		language.Russian:              harveyRussianName,
		language.SimplifiedChinese:    harveySimplifiedChineseName,
		language.Spanish:              harveySpanishName,
		language.TraditionalChinese:   harveyTraditionalChineseName}
)

var (
	// harveyCharacter represents harvey character.
	harveyCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: harveyBirthday,
		Code:     harveyCode,
		Key:      character.Harvey,
		Gender:   gender.Male,
		Name:     harveyName,
		Special:  true}
)

var (
	// Harvey represents harvey.
	Harvey = nook.Resident{
		Character: harveyCharacter}
)
