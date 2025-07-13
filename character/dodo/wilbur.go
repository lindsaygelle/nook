package dodo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// wilburBirthday represents wilbur birthday.
	wilburBirthday = nook.Birthday{
		Day:   17,
		Month: time.December}
)

var (
	// wilburCode represents wilbur code.
	wilburCode = nook.Code{
		Value: "doc"}
)

var (
	// wilburAmericanEnglishName represents wilbur american english name.
	wilburAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wilbur"}

	// wilburCanadianFrenchName represents wilbur canadian french name.
	wilburCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rodrigue"}

	// wilburDutchName represents wilbur dutch name.
	wilburDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wilbur"}

	// wilburFrenchName represents wilbur french name.
	wilburFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rodrigue"}

	// wilburGermanName represents wilbur german name.
	wilburGermanName = nook.Name{
		Language: language.German,
		Value:    "Udo"}

	// wilburItalianName represents wilbur italian name.
	wilburItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dodoardo"}

	// wilburJapaneseName represents wilbur japanese name.
	wilburJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロドリー"}

	// wilburKoreanName represents wilbur korean name.
	wilburKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "로드리"}

	// wilburLatinAmericanSpanishName represents wilbur latin american spanish name.
	wilburLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rodri"}

	// wilburRussianName represents wilbur russian name.
	wilburRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уилбур"}

	// wilburSimplifiedChineseName represents wilbur simplified chinese name.
	wilburSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "陆德里"}

	// wilburSpanishName represents wilbur spanish name.
	wilburSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rodri"}

	// wilburTraditionalChineseName represents wilbur traditional chinese name.
	wilburTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "陸德里"}
)

var (
	// wilburName represents wilbur name.
	wilburName = nook.Languages{
		language.AmericanEnglish:      wilburAmericanEnglishName,
		language.CanadianFrench:       wilburCanadianFrenchName,
		language.Dutch:                wilburDutchName,
		language.French:               wilburFrenchName,
		language.German:               wilburGermanName,
		language.Italian:              wilburItalianName,
		language.Japanese:             wilburJapaneseName,
		language.Korean:               wilburKoreanName,
		language.LatinAmericanSpanish: wilburLatinAmericanSpanishName,
		language.Russian:              wilburRussianName,
		language.SimplifiedChinese:    wilburSimplifiedChineseName,
		language.Spanish:              wilburSpanishName,
		language.TraditionalChinese:   wilburTraditionalChineseName}
)

var (
	// wilburCharacter represents wilbur character.
	wilburCharacter = nook.Character{
		Animal:   animal.Dodo,
		Birthday: wilburBirthday,
		Code:     wilburCode,
		Key:      character.Wilbur,
		Gender:   gender.Male,
		Name:     wilburName,
		Special:  true}
)

var (
	// Wilbur represents wilbur.
	Wilbur = nook.Resident{
		Character: wilburCharacter}
)
