package pigeon

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// brewsterBirthday represents brewster birthday.
	brewsterBirthday = nook.Birthday{
		Day:   15,
		Month: time.October}
)

var (
	// brewsterCode represents brewster code.
	brewsterCode = nook.Code{
		Value: "pge"}
)

var (
	// brewsterAmericanEnglishName represents brewster american english name.
	brewsterAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Brewster"}

	// brewsterCanadianFrenchName represents brewster canadian french name.
	brewsterCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Robusto"}

	// brewsterDutchName represents brewster dutch name.
	brewsterDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Brewster"}

	// brewsterFrenchName represents brewster french name.
	brewsterFrenchName = nook.Name{
		Language: language.French,
		Value:    "Robusto"}

	// brewsterGermanName represents brewster german name.
	brewsterGermanName = nook.Name{
		Language: language.German,
		Value:    "Kofi"}

	// brewsterItalianName represents brewster italian name.
	brewsterItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bartolo"}

	// brewsterJapaneseName represents brewster japanese name.
	brewsterJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マスター"}

	// brewsterKoreanName represents brewster korean name.
	brewsterKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마스터"}

	// brewsterLatinAmericanSpanishName represents brewster latin american spanish name.
	brewsterLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fígaro"}

	// brewsterRussianName represents brewster russian name.
	brewsterRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Брюстер"}

	// brewsterSimplifiedChineseName represents brewster simplified chinese name.
	brewsterSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "老板"}

	// brewsterSpanishName represents brewster spanish name.
	brewsterSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fígaro"}

	// brewsterTraditionalChineseName represents brewster traditional chinese name.
	brewsterTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "老闆"}
)

var (
	// brewsterName represents brewster name.
	brewsterName = nook.Languages{
		language.AmericanEnglish:      brewsterAmericanEnglishName,
		language.CanadianFrench:       brewsterCanadianFrenchName,
		language.Dutch:                brewsterDutchName,
		language.French:               brewsterFrenchName,
		language.German:               brewsterGermanName,
		language.Italian:              brewsterItalianName,
		language.Japanese:             brewsterJapaneseName,
		language.Korean:               brewsterKoreanName,
		language.LatinAmericanSpanish: brewsterLatinAmericanSpanishName,
		language.Russian:              brewsterRussianName,
		language.SimplifiedChinese:    brewsterSimplifiedChineseName,
		language.Spanish:              brewsterSpanishName,
		language.TraditionalChinese:   brewsterTraditionalChineseName}
)

var (
	// brewsterCharacter represents brewster character.
	brewsterCharacter = nook.Character{
		Animal:   animal.Pigeon,
		Birthday: brewsterBirthday,
		Code:     brewsterCode,
		Key:      character.Brewster,
		Gender:   gender.Male,
		Name:     brewsterName,
		Special:  true}
)

var (
	// Brewster represents brewster.
	Brewster = nook.Resident{
		Character: brewsterCharacter}
)
