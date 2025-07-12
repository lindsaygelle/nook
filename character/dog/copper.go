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
	// copperBirthday represents copper birthday.
	copperBirthday = nook.Birthday{
		Day:   28,
		Month: time.June}
)

var (
	// copperCode represents copper code.
	copperCode = nook.Code{
		Value: "plc/dga"}
)

var (
	// copperAmericanEnglishName represents copper american english name.
	copperAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Copper"}

	// copperCanadianFrenchName represents copper canadian french name.
	copperCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maret"}

	// copperDutchName represents copper dutch name.
	copperDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Copper"}

	// copperFrenchName represents copper french name.
	copperFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maret"}

	// copperGermanName represents copper german name.
	copperGermanName = nook.Name{
		Language: language.German,
		Value:    "Harry"}

	// copperItalianName represents copper italian name.
	copperItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Birro"}

	// copperJapaneseName represents copper japanese name.
	copperJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "おまわりさんA"}

	// copperKoreanName represents copper korean name.
	copperKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "경찰관A"}

	// copperLatinAmericanSpanishName represents copper latin american spanish name.
	copperLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vigilio"}

	// copperRussianName represents copper russian name.
	copperRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Коппер"}

	// copperSimplifiedChineseName represents copper simplified chinese name.
	copperSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "警察叔叔A"}

	// copperSpanishName represents copper spanish name.
	copperSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vigilio"}

	// copperTraditionalChineseName represents copper traditional chinese name.
	copperTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "警察叔叔A"}
)

var (
	// copperName represents copper name.
	copperName = nook.Languages{
		language.AmericanEnglish:      copperAmericanEnglishName,
		language.CanadianFrench:       copperCanadianFrenchName,
		language.Dutch:                copperDutchName,
		language.French:               copperFrenchName,
		language.German:               copperGermanName,
		language.Italian:              copperItalianName,
		language.Japanese:             copperJapaneseName,
		language.Korean:               copperKoreanName,
		language.LatinAmericanSpanish: copperLatinAmericanSpanishName,
		language.Russian:              copperRussianName,
		language.SimplifiedChinese:    copperSimplifiedChineseName,
		language.Spanish:              copperSpanishName,
		language.TraditionalChinese:   copperTraditionalChineseName}
)

var (
	// copperCharacter represents copper character.
	copperCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: copperBirthday,
		Code:     copperCode,
		Key:      character.Copper,
		Gender:   gender.Male,
		Name:     copperName,
		Special:  true}
)

var (
	// Copper represents copper.
	Copper = nook.Resident{
		Character: copperCharacter}
)
