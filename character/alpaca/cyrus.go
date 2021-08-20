package alpaca

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	cyrusBirthday = nook.Birthday{
		Day:   26,
		Month: time.January}
)

var (
	cyrusCode = nook.Code{
		Value: "alp"}
)

var (
	cyrusAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cyrus"}

	cyrusCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Serge"}

	cyrusDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cyrus"}

	cyrusFrenchName = nook.Name{
		Language: language.French,
		Value:    "Serge"}

	cyrusGermanName = nook.Name{
		Language: language.German,
		Value:    "Björn"}

	cyrusItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Merino"}

	cyrusJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カイゾー"}

	cyrusKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리포"}

	cyrusLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Al"}

	cyrusRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сайрус"}

	cyrusSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "健兆"}

	cyrusSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Al"}

	cyrusTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "健兆"}
)

var (
	cyrusName = nook.Languages{
		language.AmericanEnglish:      cyrusAmericanEnglishName,
		language.CanadianFrench:       cyrusCanadianFrenchName,
		language.Dutch:                cyrusDutchName,
		language.French:               cyrusFrenchName,
		language.German:               cyrusGermanName,
		language.Italian:              cyrusItalianName,
		language.Japanese:             cyrusJapaneseName,
		language.Korean:               cyrusKoreanName,
		language.LatinAmericanSpanish: cyrusLatinAmericanSpanishName,
		language.Russian:              cyrusRussianName,
		language.SimplifiedChinese:    cyrusSimplifiedChineseName,
		language.Spanish:              cyrusSpanishName,
		language.TraditionalChinese:   cyrusTraditionalChineseName}
)

var (
	cyrusCharacter = nook.Character{
		Animal:   animal.Alpaca,
		Birthday: cyrusBirthday,
		Code:     cyrusCode,
		Key:      character.Cyrus,
		Gender:   gender.Male,
		Name:     cyrusName,
		Special:  true}
)

var (
	Cyrus = nook.Resident{
		Character: cyrusCharacter}
)
