package seagull

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	gulliverBirthday = nook.Birthday{
		Day:   25,
		Month: time.May}
)

var (
	gulliverCode = nook.Code{
		Value: "seg/gul"}
)

var (
	gulliverAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gulliver"}

	gulliverCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gulliver"}

	gulliverDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gulliver"}

	gulliverFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gulliver"}

	gulliverGermanName = nook.Name{
		Language: language.German,
		Value:    "Gulliver"}

	gulliverItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gulliver"}

	gulliverJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジョニー"}

	gulliverKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "죠니"}

	gulliverLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gulliver"}

	gulliverRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гулливер"}

	gulliverSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吕游"}

	gulliverSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gulliver"}

	gulliverTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呂遊"}
)

var (
	gulliverName = nook.Languages{
		language.AmericanEnglish:      gulliverAmericanEnglishName,
		language.CanadianFrench:       gulliverCanadianFrenchName,
		language.Dutch:                gulliverDutchName,
		language.French:               gulliverFrenchName,
		language.German:               gulliverGermanName,
		language.Italian:              gulliverItalianName,
		language.Japanese:             gulliverJapaneseName,
		language.Korean:               gulliverKoreanName,
		language.LatinAmericanSpanish: gulliverLatinAmericanSpanishName,
		language.Russian:              gulliverRussianName,
		language.SimplifiedChinese:    gulliverSimplifiedChineseName,
		language.Spanish:              gulliverSpanishName,
		language.TraditionalChinese:   gulliverTraditionalChineseName}
)

var (
	gulliverCharacter = nook.Character{
		Animal:   animal.Seagull,
		Birthday: gulliverBirthday,
		Code:     gulliverCode,
		Key:      character.Gulliver,
		Gender:   gender.Male,
		Name:     gulliverName,
		Special:  true}
)

var (
	Gulliver = nook.Resident{
		Character: gulliverCharacter}
)
