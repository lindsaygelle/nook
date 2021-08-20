package frillneckedlizard

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	frillardBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	frillardCode = nook.Code{
		Value: "liz"}
)

var (
	frillardAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Frillard"}

	frillardCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Unknown"}

	frillardDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	frillardFrenchName = nook.Name{
		Language: language.French,
		Value:    "Unknown"}

	frillardGermanName = nook.Name{
		Language: language.German,
		Value:    "Eqsos"}

	frillardItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Oscar"}

	frillardJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "きょしょー"}

	frillardKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "Unknown"}

	frillardLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Unknown"}

	frillardRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	frillardSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	frillardSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Unknown"}

	frillardTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	frillardName = nook.Languages{
		language.AmericanEnglish:      frillardAmericanEnglishName,
		language.CanadianFrench:       frillardCanadianFrenchName,
		language.Dutch:                frillardDutchName,
		language.French:               frillardFrenchName,
		language.German:               frillardGermanName,
		language.Italian:              frillardItalianName,
		language.Japanese:             frillardJapaneseName,
		language.Korean:               frillardKoreanName,
		language.LatinAmericanSpanish: frillardLatinAmericanSpanishName,
		language.Russian:              frillardRussianName,
		language.SimplifiedChinese:    frillardSimplifiedChineseName,
		language.Spanish:              frillardSpanishName,
		language.TraditionalChinese:   frillardTraditionalChineseName}
)

var (
	frillardCharacter = nook.Character{
		Animal:   animal.Frillneckedlizard,
		Birthday: frillardBirthday,
		Code:     frillardCode,
		Key:      character.Frillard,
		Gender:   gender.Male,
		Name:     frillardName,
		Special:  true}
)

var (
	Frillard = nook.Resident{
		Character: frillardCharacter}
)
