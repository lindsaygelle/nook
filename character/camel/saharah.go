package camel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// saharahBirthday represents saharah birthday.
	saharahBirthday = nook.Birthday{
		Day:   10,
		Month: time.November}
)

var (
	// saharahCode represents saharah code.
	saharahCode = nook.Code{
		Value: "cml"}
)

var (
	// saharahAmericanEnglishName represents saharah american english name.
	saharahAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Saharah"}

	// saharahCanadianFrenchName represents saharah canadian french name.
	saharahCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sarah"}

	// saharahDutchName represents saharah dutch name.
	saharahDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Saharah"}

	// saharahFrenchName represents saharah french name.
	saharahFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sarah"}

	// saharahGermanName represents saharah german name.
	saharahGermanName = nook.Name{
		Language: language.German,
		Value:    "Aziza"}

	// saharahItalianName represents saharah italian name.
	saharahItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sahara"}

	// saharahJapaneseName represents saharah japanese name.
	saharahJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ローラン"}

	// saharahKoreanName represents saharah korean name.
	saharahKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사하라"}

	// saharahLatinAmericanSpanishName represents saharah latin american spanish name.
	saharahLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Alcatifa"}

	// saharahRussianName represents saharah russian name.
	saharahRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сахара"}

	// saharahSimplifiedChineseName represents saharah simplified chinese name.
	saharahSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "骆岚"}

	// saharahSpanishName represents saharah spanish name.
	saharahSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Alcatifa"}

	// saharahTraditionalChineseName represents saharah traditional chinese name.
	saharahTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "駱嵐"}
)

var (
	// saharahName represents saharah name.
	saharahName = nook.Languages{
		language.AmericanEnglish:      saharahAmericanEnglishName,
		language.CanadianFrench:       saharahCanadianFrenchName,
		language.Dutch:                saharahDutchName,
		language.French:               saharahFrenchName,
		language.German:               saharahGermanName,
		language.Italian:              saharahItalianName,
		language.Japanese:             saharahJapaneseName,
		language.Korean:               saharahKoreanName,
		language.LatinAmericanSpanish: saharahLatinAmericanSpanishName,
		language.Russian:              saharahRussianName,
		language.SimplifiedChinese:    saharahSimplifiedChineseName,
		language.Spanish:              saharahSpanishName,
		language.TraditionalChinese:   saharahTraditionalChineseName}
)

var (
	// saharahCharacter represents saharah character.
	saharahCharacter = nook.Character{
		Animal:   animal.Camel,
		Birthday: saharahBirthday,
		Code:     saharahCode,
		Key:      character.Saharah,
		Gender:   gender.Female,
		Name:     saharahName,
		Special:  true}
)

var (
	// Saharah represents saharah.
	Saharah = nook.Resident{
		Character: saharahCharacter}
)
