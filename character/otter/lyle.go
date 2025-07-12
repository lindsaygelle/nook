package otter

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// lyleBirthday represents lyle birthday.
	lyleBirthday = nook.Birthday{
		Day:   6,
		Month: time.June}
)

var (
	// lyleCode represents lyle code.
	lyleCode = nook.Code{
		Value: "ott"}
)

var (
	// lyleAmericanEnglishName represents lyle american english name.
	lyleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lyle"}

	// lyleCanadianFrenchName represents lyle canadian french name.
	lyleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lionel"}

	// lyleDutchName represents lyle dutch name.
	lyleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lyle"}

	// lyleFrenchName represents lyle french name.
	lyleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lionel"}

	// lyleGermanName represents lyle german name.
	lyleGermanName = nook.Name{
		Language: language.German,
		Value:    "Fred"}

	// lyleItalianName represents lyle italian name.
	lyleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frodolo"}

	// lyleJapaneseName represents lyle japanese name.
	lyleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホンマさん"}

	// lyleKoreanName represents lyle korean name.
	lyleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안심해씨"}

	// lyleLatinAmericanSpanishName represents lyle latin american spanish name.
	lyleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sisebuto"}

	// lyleRussianName represents lyle russian name.
	lyleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лайл"}

	// lyleSimplifiedChineseName represents lyle simplified chinese name.
	lyleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿本先生"}

	// lyleSpanishName represents lyle spanish name.
	lyleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sisebuto"}

	// lyleTraditionalChineseName represents lyle traditional chinese name.
	lyleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿本先生"}
)

var (
	// lyleName represents lyle name.
	lyleName = nook.Languages{
		language.AmericanEnglish:      lyleAmericanEnglishName,
		language.CanadianFrench:       lyleCanadianFrenchName,
		language.Dutch:                lyleDutchName,
		language.French:               lyleFrenchName,
		language.German:               lyleGermanName,
		language.Italian:              lyleItalianName,
		language.Japanese:             lyleJapaneseName,
		language.Korean:               lyleKoreanName,
		language.LatinAmericanSpanish: lyleLatinAmericanSpanishName,
		language.Russian:              lyleRussianName,
		language.SimplifiedChinese:    lyleSimplifiedChineseName,
		language.Spanish:              lyleSpanishName,
		language.TraditionalChinese:   lyleTraditionalChineseName}
)

var (
	// lyleCharacter represents lyle character.
	lyleCharacter = nook.Character{
		Animal:   animal.Otter,
		Birthday: lyleBirthday,
		Code:     lyleCode,
		Key:      character.Lyle,
		Gender:   gender.Male,
		Name:     lyleName,
		Special:  true}
)

var (
	// Lyle represents lyle.
	Lyle = nook.Resident{
		Character: lyleCharacter}
)
