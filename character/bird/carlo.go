package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// carloBirthday represents carlo birthday.
	carloBirthday = nook.Birthday{
		Day:   3,
		Month: time.May}
)

var (
	// carloCode represents carlo code.
	carloCode = nook.Code{
		Value: "cwb"}
)

var (
	// carloAmericanEnglishName represents carlo american english name.
	carloAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Carlo"}

	// carloCanadianFrenchName represents carlo canadian french name.
	carloCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Plumac"}

	// carloDutchName represents carlo dutch name.
	carloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// carloFrenchName represents carlo french name.
	carloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Plumac"}

	// carloGermanName represents carlo german name.
	carloGermanName = nook.Name{
		Language: language.German,
		Value:    "Ralle"}

	// carloItalianName represents carlo italian name.
	carloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giorvo"}

	// carloJapaneseName represents carlo japanese name.
	carloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グーさん"}

	// carloKoreanName represents carlo korean name.
	carloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "카를로"}

	// carloLatinAmericanSpanishName represents carlo latin american spanish name.
	carloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Graznel"}

	// carloRussianName represents carlo russian name.
	carloRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// carloSimplifiedChineseName represents carlo simplified chinese name.
	carloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// carloSpanishName represents carlo spanish name.
	carloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Graznel"}

	// carloTraditionalChineseName represents carlo traditional chinese name.
	carloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "營長"}
)

var (
	// carloName represents carlo name.
	carloName = nook.Languages{
		language.AmericanEnglish:      carloAmericanEnglishName,
		language.CanadianFrench:       carloCanadianFrenchName,
		language.Dutch:                carloDutchName,
		language.French:               carloFrenchName,
		language.German:               carloGermanName,
		language.Italian:              carloItalianName,
		language.Japanese:             carloJapaneseName,
		language.Korean:               carloKoreanName,
		language.LatinAmericanSpanish: carloLatinAmericanSpanishName,
		language.Russian:              carloRussianName,
		language.SimplifiedChinese:    carloSimplifiedChineseName,
		language.Spanish:              carloSpanishName,
		language.TraditionalChinese:   carloTraditionalChineseName}
)

var (
	// carloCharacter represents carlo character.
	carloCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: carloBirthday,
		Code:     carloCode,
		Key:      character.Carlo,
		Gender:   gender.Male,
		Name:     carloName,
		Special:  true}
)

var (
	// Carlo represents carlo.
	Carlo = nook.Resident{
		Character: carloCharacter}
)
