package furseal

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// phineasBirthday represents phineas birthday.
	phineasBirthday = nook.Birthday{
		Day:   20,
		Month: time.June}
)

var (
	// phineasCode represents phineas code.
	phineasCode = nook.Code{
		Value: "bln/fsl"}
)

var (
	// phineasAmericanEnglishName represents phineas american english name.
	phineasAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Phineas"}

	// phineasCanadianFrenchName represents phineas canadian french name.
	phineasCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hélium"}

	// phineasDutchName represents phineas dutch name.
	phineasDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Phineas"}

	// phineasFrenchName represents phineas french name.
	phineasFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hélium"}

	// phineasGermanName represents phineas german name.
	phineasGermanName = nook.Name{
		Language: language.German,
		Value:    "Helios"}

	// phineasItalianName represents phineas italian name.
	phineasItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Elio"}

	// phineasJapaneseName represents phineas japanese name.
	phineasJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パロンチーノ"}

	// phineasKoreanName represents phineas korean name.
	phineasKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파론티노"}

	// phineasLatinAmericanSpanishName represents phineas latin american spanish name.
	phineasLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Elio"}

	// phineasRussianName represents phineas russian name.
	phineasRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Финеас"}

	// phineasSimplifiedChineseName represents phineas simplified chinese name.
	phineasSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "博陆恩"}

	// phineasSpanishName represents phineas spanish name.
	phineasSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Elio"}

	// phineasTraditionalChineseName represents phineas traditional chinese name.
	phineasTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "博陸恩"}
)

var (
	// phineasName represents phineas name.
	phineasName = nook.Languages{
		language.AmericanEnglish:      phineasAmericanEnglishName,
		language.CanadianFrench:       phineasCanadianFrenchName,
		language.Dutch:                phineasDutchName,
		language.French:               phineasFrenchName,
		language.German:               phineasGermanName,
		language.Italian:              phineasItalianName,
		language.Japanese:             phineasJapaneseName,
		language.Korean:               phineasKoreanName,
		language.LatinAmericanSpanish: phineasLatinAmericanSpanishName,
		language.Russian:              phineasRussianName,
		language.SimplifiedChinese:    phineasSimplifiedChineseName,
		language.Spanish:              phineasSpanishName,
		language.TraditionalChinese:   phineasTraditionalChineseName}
)

var (
	// phineasCharacter represents phineas character.
	phineasCharacter = nook.Character{
		Animal:   animal.FurSeal,
		Birthday: phineasBirthday,
		Code:     phineasCode,
		Key:      character.Phineas,
		Gender:   gender.Male,
		Name:     phineasName,
		Special:  true}
)

var (
	// Phineas represents phineas.
	Phineas = nook.Resident{
		Character: phineasCharacter}
)
