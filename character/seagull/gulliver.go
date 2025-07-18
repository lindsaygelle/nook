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
	// gulliverBirthday represents gulliver birthday.
	gulliverBirthday = nook.Birthday{
		Day:   25,
		Month: time.May}
)

var (
	// gulliverCode represents gulliver code.
	gulliverCode = nook.Code{
		Value: "seg/gul"}
)

var (
	// gulliverAmericanEnglishName represents gulliver american english name.
	gulliverAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gulliver"}

	// gulliverCanadianFrenchName represents gulliver canadian french name.
	gulliverCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gulliver"}

	// gulliverDutchName represents gulliver dutch name.
	gulliverDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gulliver"}

	// gulliverFrenchName represents gulliver french name.
	gulliverFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gulliver"}

	// gulliverGermanName represents gulliver german name.
	gulliverGermanName = nook.Name{
		Language: language.German,
		Value:    "Gulliver"}

	// gulliverItalianName represents gulliver italian name.
	gulliverItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gulliver"}

	// gulliverJapaneseName represents gulliver japanese name.
	gulliverJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジョニー"}

	// gulliverKoreanName represents gulliver korean name.
	gulliverKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "죠니"}

	// gulliverLatinAmericanSpanishName represents gulliver latin american spanish name.
	gulliverLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gulliver"}

	// gulliverRussianName represents gulliver russian name.
	gulliverRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гулливер"}

	// gulliverSimplifiedChineseName represents gulliver simplified chinese name.
	gulliverSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吕游"}

	// gulliverSpanishName represents gulliver spanish name.
	gulliverSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gulliver"}

	// gulliverTraditionalChineseName represents gulliver traditional chinese name.
	gulliverTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呂遊"}
)

var (
	// gulliverName represents gulliver name.
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
	// gulliverCharacter represents gulliver character.
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
	// Gulliver represents gulliver.
	Gulliver = nook.Resident{
		Character: gulliverCharacter}
)
