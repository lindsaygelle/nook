package owl

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	celesteBirthday = nook.Birthday{
		Day:   7,
		Month: time.September}
)

var (
	celesteCode = nook.Code{
		Value: "ows"}
)

var (
	celesteAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Celeste"}

	celesteCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Céleste"}

	celesteDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Celeste"}

	celesteFrenchName = nook.Name{
		Language: language.French,
		Value:    "Céleste"}

	celesteGermanName = nook.Name{
		Language: language.German,
		Value:    "Eufemia"}

	celesteItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Celeste"}

	celesteJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フーコ"}

	celesteKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "부옥"}

	celesteLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Estela"}

	celesteRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Селеста"}

	celesteSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "傅珂"}

	celesteSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Estela"}

	celesteTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "傅珂"}
)

var (
	celesteName = nook.Languages{
		language.AmericanEnglish:      celesteAmericanEnglishName,
		language.CanadianFrench:       celesteCanadianFrenchName,
		language.Dutch:                celesteDutchName,
		language.French:               celesteFrenchName,
		language.German:               celesteGermanName,
		language.Italian:              celesteItalianName,
		language.Japanese:             celesteJapaneseName,
		language.Korean:               celesteKoreanName,
		language.LatinAmericanSpanish: celesteLatinAmericanSpanishName,
		language.Russian:              celesteRussianName,
		language.SimplifiedChinese:    celesteSimplifiedChineseName,
		language.Spanish:              celesteSpanishName,
		language.TraditionalChinese:   celesteTraditionalChineseName}
)

var (
	celesteCharacter = nook.Character{
		Animal:   animal.Owl,
		Birthday: celesteBirthday,
		Code:     celesteCode,
		Gender:   gender.Female,
		Name:     celesteName}
)

var (
	Celeste = nook.Resident{
		Character: celesteCharacter}
)
