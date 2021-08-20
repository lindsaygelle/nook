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
	gramsBirthday = nook.Birthday{
		Day:   15,
		Month: time.April}
)

var (
	gramsCode = nook.Code{
		Value: "kpg"}
)

var (
	gramsAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Grams"}

	gramsCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mamiral"}

	gramsDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Grams"}

	gramsFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mamiral"}

	gramsGermanName = nook.Name{
		Language: language.German,
		Value:    "Bonnie"}

	gramsItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Remigia"}

	gramsJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゲコ"}

	gramsKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "귀녀"}

	gramsLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rosauria"}

	gramsRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Грэмс"}

	gramsSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿玉婆婆"}

	gramsSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosauria"}

	gramsTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿玉婆婆"}
)

var (
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
	Grams = nook.Resident{
		Character: gramsCharacter}
)
