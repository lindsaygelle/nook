package turtle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// gramsBirthday represents grams birthday.
	gramsBirthday = nook.Birthday{
		Day:   15,
		Month: time.April}
)

var (
	// gramsCode represents grams code.
	gramsCode = nook.Code{
		Value: "kpg"}
)

var (
	// gramsAmericanEnglishName represents grams american english name.
	gramsAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Grams"}

	// gramsCanadianFrenchName represents grams canadian french name.
	gramsCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mamiral"}

	// gramsDutchName represents grams dutch name.
	gramsDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Grams"}

	// gramsFrenchName represents grams french name.
	gramsFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mamiral"}

	// gramsGermanName represents grams german name.
	gramsGermanName = nook.Name{
		Language: language.German,
		Value:    "Bonnie"}

	// gramsItalianName represents grams italian name.
	gramsItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Remigia"}

	// gramsJapaneseName represents grams japanese name.
	gramsJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゲコ"}

	// gramsKoreanName represents grams korean name.
	gramsKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "귀녀"}

	// gramsLatinAmericanSpanishName represents grams latin american spanish name.
	gramsLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rosauria"}

	// gramsRussianName represents grams russian name.
	gramsRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Грэмс"}

	// gramsSimplifiedChineseName represents grams simplified chinese name.
	gramsSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿玉婆婆"}

	// gramsSpanishName represents grams spanish name.
	gramsSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosauria"}

	// gramsTraditionalChineseName represents grams traditional chinese name.
	gramsTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿玉婆婆"}
)

var (
	// gramsName represents grams name.
	gramsName = nook.Languages{
		language.AmericanEnglish:      gramsAmericanEnglishName,
		language.CanadianFrench:       gramsCanadianFrenchName,
		language.Dutch:                gramsDutchName,
		language.French:               gramsFrenchName,
		language.German:               gramsGermanName,
		language.Italian:              gramsItalianName,
		language.Japanese:             gramsJapaneseName,
		language.Korean:               gramsKoreanName,
		language.LatinAmericanSpanish: gramsLatinAmericanSpanishName,
		language.Russian:              gramsRussianName,
		language.SimplifiedChinese:    gramsSimplifiedChineseName,
		language.Spanish:              gramsSpanishName,
		language.TraditionalChinese:   gramsTraditionalChineseName}
)

var (
	// gramsCharacter represents grams character.
	gramsCharacter = nook.Character{
		Animal:   animal.Turtle,
		Birthday: gramsBirthday,
		Code:     gramsCode,
		Key:      character.Grams,
		Gender:   gender.Female,
		Name:     gramsName,
		Special:  true}
)

var (
	// Grams represents grams.
	Grams = nook.Resident{
		Character: gramsCharacter}
)
