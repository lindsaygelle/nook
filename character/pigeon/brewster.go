package pigeon

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	brewsterBirthday = nook.Birthday{
		Day:   15,
		Month: time.October}
)

var (
	brewsterCode = nook.Code{
		Value: "pge"}
)

var (
	brewsterAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Brewster"}

	brewsterCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Robusto"}

	brewsterDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Brewster"}

	brewsterFrenchName = nook.Name{
		Language: language.French,
		Value:    "Robusto"}

	brewsterGermanName = nook.Name{
		Language: language.German,
		Value:    "Kofi"}

	brewsterItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bartolo"}

	brewsterJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マスター"}

	brewsterKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마스터"}

	brewsterLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fígaro"}

	brewsterRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Брюстер"}

	brewsterSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "老板"}

	brewsterSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fígaro"}

	brewsterTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "老闆"}
)

var (
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
	brewsterCharacter = nook.Character{
		Animal:   animal.Pigeon,
		Birthday: brewsterBirthday,
		Code:     brewsterCode,
		Gender:   gender.Male,
		Name:     brewsterName}
)

var (
	Brewster = nook.Resident{
		Character: brewsterCharacter}
)
