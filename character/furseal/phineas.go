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
	phineasBirthday = nook.Birthday{
		Day:   20,
		Month: time.June}
)

var (
	phineasCode = nook.Code{
		Value: "bln/fsl"}
)

var (
	phineasAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Phineas"}

	phineasCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hélium"}

	phineasDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Phineas"}

	phineasFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hélium"}

	phineasGermanName = nook.Name{
		Language: language.German,
		Value:    "Helios"}

	phineasItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Elio"}

	phineasJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パロンチーノ"}

	phineasKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파론티노"}

	phineasLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Elio"}

	phineasRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Финеас"}

	phineasSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "博陆恩"}

	phineasSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Elio"}

	phineasTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "博陸恩"}
)

var (
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
	phineasCharacter = nook.Character{
		Animal:   animal.Furseal,
		Birthday: phineasBirthday,
		Code:     phineasCode,
		Key:      character.Phineas,
		Gender:   gender.Male,
		Name:     phineasName}
)

var (
	Phineas = nook.Resident{
		Character: phineasCharacter}
)
