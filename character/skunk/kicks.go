package skunk

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// kicksBirthday represents kicks birthday.
	kicksBirthday = nook.Birthday{
		Day:   30,
		Month: time.November}
)

var (
	// kicksCode represents kicks code.
	kicksCode = nook.Code{
		Value: "skk"}
)

var (
	// kicksAmericanEnglishName represents kicks american english name.
	kicksAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kicks"}

	// kicksCanadianFrenchName represents kicks canadian french name.
	kicksCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Blaise"}

	// kicksDutchName represents kicks dutch name.
	kicksDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kicks"}

	// kicksFrenchName represents kicks french name.
	kicksFrenchName = nook.Name{
		Language: language.French,
		Value:    "Blaise"}

	// kicksGermanName represents kicks german name.
	kicksGermanName = nook.Name{
		Language: language.German,
		Value:    "Schubert"}

	// kicksItalianName represents kicks italian name.
	kicksItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sciuscià"}

	// kicksJapaneseName represents kicks japanese name.
	kicksJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シャンク"}

	// kicksKoreanName represents kicks korean name.
	kicksKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "패트릭"}

	// kicksLatinAmericanSpanishName represents kicks latin american spanish name.
	kicksLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Betunio"}

	// kicksRussianName represents kicks russian name.
	kicksRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кикс"}

	// kicksSimplifiedChineseName represents kicks simplified chinese name.
	kicksSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "薛革"}

	// kicksSpanishName represents kicks spanish name.
	kicksSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Betunio"}

	// kicksTraditionalChineseName represents kicks traditional chinese name.
	kicksTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "薛革"}
)

var (
	// kicksName represents kicks name.
	kicksName = nook.Languages{
		language.AmericanEnglish:      kicksAmericanEnglishName,
		language.CanadianFrench:       kicksCanadianFrenchName,
		language.Dutch:                kicksDutchName,
		language.French:               kicksFrenchName,
		language.German:               kicksGermanName,
		language.Italian:              kicksItalianName,
		language.Japanese:             kicksJapaneseName,
		language.Korean:               kicksKoreanName,
		language.LatinAmericanSpanish: kicksLatinAmericanSpanishName,
		language.Russian:              kicksRussianName,
		language.SimplifiedChinese:    kicksSimplifiedChineseName,
		language.Spanish:              kicksSpanishName,
		language.TraditionalChinese:   kicksTraditionalChineseName}
)

var (
	// kicksCharacter represents kicks character.
	kicksCharacter = nook.Character{
		Animal:   animal.Skunk,
		Birthday: kicksBirthday,
		Code:     kicksCode,
		Key:      character.Kicks,
		Gender:   gender.Male,
		Name:     kicksName,
		Special:  true}
)

var (
	// Kicks represents kicks.
	Kicks = nook.Resident{
		Character: kicksCharacter}
)
