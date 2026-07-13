package seagull

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/game"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// gullivarrrGames represents gullivarrr game appearances.
	gullivarrrGames = []nook.Game{
		game.NewHorizons,
	}
)

var (
	// gullivarrrBirthday represents gullivarrr birthday.
	gullivarrrBirthday = nook.Birthday{
		Day:   25,
		Month: time.May}
)

var (
	// gullivarrrCode represents gullivarrr code.
	gullivarrrCode = nook.Code{
		Value: "gulB"}
)

var (
	// gullivarrrAmericanEnglishName represents gullivarrr american english name.
	gullivarrrAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gullivarrr"}

	// gullivarrrCanadianFrenchName represents gullivarrr canadian french name.
	gullivarrrCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gullivarrr"}

	// gullivarrrDutchName represents gullivarrr dutch name.
	gullivarrrDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gullivarrr"}

	// gullivarrrFrenchName represents gullivarrr french name.
	gullivarrrFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gullivarrr"}

	// gullivarrrGermanName represents gullivarrr german name.
	gullivarrrGermanName = nook.Name{
		Language: language.German,
		Value:    "Gullivarrr"}

	// gullivarrrItalianName represents gullivarrr italian name.
	gullivarrrItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gullivarrr"}

	// gullivarrrJapaneseName represents gullivarrr japanese name.
	gullivarrrJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "かいぞく"}

	// gullivarrrKoreanName represents gullivarrr korean name.
	gullivarrrKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "해적"}

	// gullivarrrLatinAmericanSpanishName represents gullivarrr latin american spanish name.
	gullivarrrLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gullivarrr"}

	// gullivarrrRussianName represents gullivarrr russian name.
	gullivarrrRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гулливарр"}

	// gullivarrrSimplifiedChineseName represents gullivarrr simplified chinese name.
	gullivarrrSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "海盗"}

	// gullivarrrSpanishName represents gullivarrr spanish name.
	gullivarrrSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gullivarrr"}

	// gullivarrrTraditionalChineseName represents gullivarrr traditional chinese name.
	gullivarrrTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "海盜"}
)

var (
	// gullivarrrName represents gullivarrr name.
	gullivarrrName = nook.Languages{
		language.AmericanEnglish:      gullivarrrAmericanEnglishName,
		language.CanadianFrench:       gullivarrrCanadianFrenchName,
		language.Dutch:                gullivarrrDutchName,
		language.French:               gullivarrrFrenchName,
		language.German:               gullivarrrGermanName,
		language.Italian:              gullivarrrItalianName,
		language.Japanese:             gullivarrrJapaneseName,
		language.Korean:               gullivarrrKoreanName,
		language.LatinAmericanSpanish: gullivarrrLatinAmericanSpanishName,
		language.Russian:              gullivarrrRussianName,
		language.SimplifiedChinese:    gullivarrrSimplifiedChineseName,
		language.Spanish:              gullivarrrSpanishName,
		language.TraditionalChinese:   gullivarrrTraditionalChineseName}
)

var (
	// gullivarrrCharacter represents gullivarrr character.
	gullivarrrCharacter = nook.Character{
		Animal:   animal.Seagull,
		Birthday: gullivarrrBirthday,
		Code:     gullivarrrCode,
		Games:    gullivarrrGames,
		Gender:   gender.Male,
		Key:      character.Gullivarrr,
		Name:     gullivarrrName,
		Special:  true}
)

var (
	// Gullivarrr represents gullivarrr.
	Gullivarrr = nook.Resident{
		Character: gullivarrrCharacter}
)
