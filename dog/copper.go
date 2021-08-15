package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	copperBirthday = nook.Birthday{
		Day:   28,
		Month: time.June}
)

var (
	copperCode = nook.Code{
		Value: "plc/dga"}
)

var (
	copperAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Copper"}

	copperCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maret"}

	copperDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Copper"}

	copperFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maret"}

	copperGermanName = nook.Name{
		Language: language.German,
		Value:    "Harry"}

	copperItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Birro"}

	copperJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "おまわりさんA"}

	copperKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "경찰관A"}

	copperLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vigilio"}

	copperRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Коппер"}

	copperSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "警察叔叔A"}

	copperSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vigilio"}

	copperTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "警察叔叔A"}
)

var (
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
	copperCharacter = nook.Character{
		Animal:   Dog,
		Birthday: copperBirthday,
		Code:     copperCode,
		Gender:   nook.Male,
		Name:     copperName}
)

var (
	Copper = nook.Resident{
		Character: copperCharacter}
)
