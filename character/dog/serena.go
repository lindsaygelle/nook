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
	// serenaBirthday represents serena birthday.
	serenaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// serenaCode represents serena code.
	serenaCode = nook.Code{
		Value: "gds"}
)

var (
	// serenaAmericanEnglishName represents serena american english name.
	serenaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Serena"}

	// serenaCanadianFrenchName represents serena canadian french name.
	serenaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Unknown"}

	// serenaDutchName represents serena dutch name.
	serenaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// serenaFrenchName represents serena french name.
	serenaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Unknown"}

	// serenaGermanName represents serena german name.
	serenaGermanName = nook.Name{
		Language: language.German,
		Value:    "Divahua"}

	// serenaItalianName represents serena italian name.
	serenaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Candea"}

	// serenaJapaneseName represents serena japanese name.
	serenaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "めがみさま"}

	// serenaKoreanName represents serena korean name.
	serenaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "Unknown"}

	// serenaLatinAmericanSpanishName represents serena latin american spanish name.
	serenaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Divahua"}

	// serenaRussianName represents serena russian name.
	serenaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// serenaSimplifiedChineseName represents serena simplified chinese name.
	serenaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// serenaSpanishName represents serena spanish name.
	serenaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Divahua"}

	// serenaTraditionalChineseName represents serena traditional chinese name.
	serenaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// serenaName represents serena name.
	serenaName = nook.Languages{
		language.AmericanEnglish:      serenaAmericanEnglishName,
		language.CanadianFrench:       serenaCanadianFrenchName,
		language.Dutch:                serenaDutchName,
		language.French:               serenaFrenchName,
		language.German:               serenaGermanName,
		language.Italian:              serenaItalianName,
		language.Japanese:             serenaJapaneseName,
		language.Korean:               serenaKoreanName,
		language.LatinAmericanSpanish: serenaLatinAmericanSpanishName,
		language.Russian:              serenaRussianName,
		language.SimplifiedChinese:    serenaSimplifiedChineseName,
		language.Spanish:              serenaSpanishName,
		language.TraditionalChinese:   serenaTraditionalChineseName}
)

var (
	// serenaCharacter represents serena character.
	serenaCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: serenaBirthday,
		Code:     serenaCode,
		Key:      character.Serena,
		Gender:   gender.Female,
		Name:     serenaName,
		Special:  true}
)

var (
	// Serena represents serena.
	Serena = nook.Resident{
		Character: serenaCharacter}
)
