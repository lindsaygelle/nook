package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// shakiBirthday represents shaki birthday.
	shakiBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// shakiCode represents shaki code.
	shakiCode = nook.Code{
		Value: "xsq"}
)

var (
	// shakiAmericanEnglishName represents shaki american english name.
	shakiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Shaki"}

	// shakiCanadianFrenchName represents shaki canadian french name.
	shakiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// shakiDutchName represents shaki dutch name.
	shakiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// shakiFrenchName represents shaki french name.
	shakiFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// shakiGermanName represents shaki german name.
	shakiGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// shakiItalianName represents shaki italian name.
	shakiItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// shakiJapaneseName represents shaki japanese name.
	shakiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シャキッ"}

	// shakiKoreanName represents shaki korean name.
	shakiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// shakiLatinAmericanSpanishName represents shaki latin american spanish name.
	shakiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// shakiRussianName represents shaki russian name.
	shakiRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// shakiSimplifiedChineseName represents shaki simplified chinese name.
	shakiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// shakiSpanishName represents shaki spanish name.
	shakiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// shakiTraditionalChineseName represents shaki traditional chinese name.
	shakiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// shakiName represents shaki name.
	shakiName = nook.Languages{
		language.AmericanEnglish:      shakiAmericanEnglishName,
		language.CanadianFrench:       shakiCanadianFrenchName,
		language.Dutch:                shakiDutchName,
		language.French:               shakiFrenchName,
		language.German:               shakiGermanName,
		language.Italian:              shakiItalianName,
		language.Japanese:             shakiJapaneseName,
		language.Korean:               shakiKoreanName,
		language.LatinAmericanSpanish: shakiLatinAmericanSpanishName,
		language.Russian:              shakiRussianName,
		language.SimplifiedChinese:    shakiSimplifiedChineseName,
		language.Spanish:              shakiSpanishName,
		language.TraditionalChinese:   shakiTraditionalChineseName}
)

var (
	// shakiCharacter represents shaki character.
	shakiCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: shakiBirthday,
		Code:     shakiCode,
		Key:      character.Shaki,
		Gender:   gender.Female,
		Name:     shakiName,
		Special:  true}
)

var (
	// Shaki represents shaki.
	Shaki = nook.Resident{
		Character: shakiCharacter}
)
