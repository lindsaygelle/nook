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
	chipBirthday = nook.Birthday{
		Day:   9,
		Month: time.December}
)

var (
	chipCode = nook.Code{
		Value: "bev"}
)

var (
	chipAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chip"}

	chipCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Castor"}

	chipDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chip"}

	chipFrenchName = nook.Name{
		Language: language.French,
		Value:    "Castor"}

	chipGermanName = nook.Name{
		Language: language.German,
		Value:    "Bartholo"}

	chipItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Castore"}

	chipJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "うおまさ"}

	chipKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뚱달"}

	chipLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Martín"}

	chipRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чип"}

	chipSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "俞正"}

	chipSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Martín"}

	chipTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "俞正"}
)

var (
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
	chipCharacter = nook.Character{
		Animal:   animal.Beaver,
		Birthday: chipBirthday,
		Code:     chipCode,
		Key:      character.Chip,
		Gender:   gender.Male,
		Name:     chipName}
)

var (
	Chip = nook.Resident{
		Character: chipCharacter}
)
