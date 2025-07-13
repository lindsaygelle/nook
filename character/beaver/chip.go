package beaver

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// chipBirthday represents chip birthday.
	chipBirthday = nook.Birthday{
		Day:   9,
		Month: time.December}
)

var (
	// chipCode represents chip code.
	chipCode = nook.Code{
		Value: "bev"}
)

var (
	// chipAmericanEnglishName represents chip american english name.
	chipAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chip"}

	// chipCanadianFrenchName represents chip canadian french name.
	chipCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Castor"}

	// chipDutchName represents chip dutch name.
	chipDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chip"}

	// chipFrenchName represents chip french name.
	chipFrenchName = nook.Name{
		Language: language.French,
		Value:    "Castor"}

	// chipGermanName represents chip german name.
	chipGermanName = nook.Name{
		Language: language.German,
		Value:    "Bartholo"}

	// chipItalianName represents chip italian name.
	chipItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Castore"}

	// chipJapaneseName represents chip japanese name.
	chipJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "うおまさ"}

	// chipKoreanName represents chip korean name.
	chipKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뚱달"}

	// chipLatinAmericanSpanishName represents chip latin american spanish name.
	chipLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Martín"}

	// chipRussianName represents chip russian name.
	chipRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чип"}

	// chipSimplifiedChineseName represents chip simplified chinese name.
	chipSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "俞正"}

	// chipSpanishName represents chip spanish name.
	chipSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Martín"}

	// chipTraditionalChineseName represents chip traditional chinese name.
	chipTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "俞正"}
)

var (
	// chipName represents chip name.
	chipName = nook.Languages{
		language.AmericanEnglish:      chipAmericanEnglishName,
		language.CanadianFrench:       chipCanadianFrenchName,
		language.Dutch:                chipDutchName,
		language.French:               chipFrenchName,
		language.German:               chipGermanName,
		language.Italian:              chipItalianName,
		language.Japanese:             chipJapaneseName,
		language.Korean:               chipKoreanName,
		language.LatinAmericanSpanish: chipLatinAmericanSpanishName,
		language.Russian:              chipRussianName,
		language.SimplifiedChinese:    chipSimplifiedChineseName,
		language.Spanish:              chipSpanishName,
		language.TraditionalChinese:   chipTraditionalChineseName}
)

var (
	// chipCharacter represents chip character.
	chipCharacter = nook.Character{
		Animal:   animal.Beaver,
		Birthday: chipBirthday,
		Code:     chipCode,
		Key:      character.Chip,
		Gender:   gender.Male,
		Name:     chipName,
		Special:  true}
)

var (
	// Chip represents chip.
	Chip = nook.Resident{
		Character: chipCharacter}
)
