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
	// Define Cyrus' birthday.
	cyrusBirthday = nook.Birthday{
		Day:   26,
		Month: time.January}
)

var (
	// Define Cyrus' code.
	cyrusCode = nook.Code{
		Value: "alp"}
)

// Define Cyrus' names in various languages.
var (
	// cyrusAmericanEnglishName represents cyrus american english name.
	cyrusAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cyrus"}

	// cyrusCanadianFrenchName represents cyrus canadian french name.
	cyrusCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Serge"}

	// cyrusDutchName represents cyrus dutch name.
	cyrusDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cyrus"}

	// cyrusFrenchName represents cyrus french name.
	cyrusFrenchName = nook.Name{
		Language: language.French,
		Value:    "Serge"}

	// cyrusGermanName represents cyrus german name.
	cyrusGermanName = nook.Name{
		Language: language.German,
		Value:    "Björn"}

	// cyrusItalianName represents cyrus italian name.
	cyrusItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Merino"}

	// cyrusJapaneseName represents cyrus japanese name.
	cyrusJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カイゾー"}

	// cyrusKoreanName represents cyrus korean name.
	cyrusKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리포"}

	// cyrusLatinAmericanSpanishName represents cyrus latin american spanish name.
	cyrusLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Al"}

	// cyrusRussianName represents cyrus russian name.
	cyrusRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сайрус"}

	// cyrusSimplifiedChineseName represents cyrus simplified chinese name.
	cyrusSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "健兆"}

	// cyrusSpanishName represents cyrus spanish name.
	cyrusSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Al"}

	// cyrusTraditionalChineseName represents cyrus traditional chinese name.
	cyrusTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "健兆"}
)

var (
	// Define Cyrus' names in different languages as a map.
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
	// Define Cyrus' character attributes.
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
	// Cyrus as a resident.
	Cyrus = nook.Resident{
		Character: cyrusCharacter}
)
