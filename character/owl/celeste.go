package owl

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// celesteBirthday represents celeste birthday.
	celesteBirthday = nook.Birthday{
		Day:   7,
		Month: time.September}
)

var (
	// celesteCode represents celeste code.
	celesteCode = nook.Code{
		Value: "ows"}
)

var (
	// celesteAmericanEnglishName represents celeste american english name.
	celesteAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Celeste"}

	// celesteCanadianFrenchName represents celeste canadian french name.
	celesteCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Céleste"}

	// celesteDutchName represents celeste dutch name.
	celesteDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Celeste"}

	// celesteFrenchName represents celeste french name.
	celesteFrenchName = nook.Name{
		Language: language.French,
		Value:    "Céleste"}

	// celesteGermanName represents celeste german name.
	celesteGermanName = nook.Name{
		Language: language.German,
		Value:    "Eufemia"}

	// celesteItalianName represents celeste italian name.
	celesteItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Celeste"}

	// celesteJapaneseName represents celeste japanese name.
	celesteJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フーコ"}

	// celesteKoreanName represents celeste korean name.
	celesteKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "부옥"}

	// celesteLatinAmericanSpanishName represents celeste latin american spanish name.
	celesteLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Estela"}

	// celesteRussianName represents celeste russian name.
	celesteRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Селеста"}

	// celesteSimplifiedChineseName represents celeste simplified chinese name.
	celesteSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "傅珂"}

	// celesteSpanishName represents celeste spanish name.
	celesteSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Estela"}

	// celesteTraditionalChineseName represents celeste traditional chinese name.
	celesteTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "傅珂"}
)

var (
	// celesteName represents celeste name.
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
	// celesteCharacter represents celeste character.
	celesteCharacter = nook.Character{
		Animal:   animal.Owl,
		Birthday: celesteBirthday,
		Code:     celesteCode,
		Key:      character.Celeste,
		Gender:   gender.Female,
		Name:     celesteName,
		Special:  true}
)

var (
	// Celeste represents celeste.
	Celeste = nook.Resident{
		Character: celesteCharacter}
)
