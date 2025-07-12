package dodo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// orvilleBirthday represents orville birthday.
	orvilleBirthday = nook.Birthday{
		Day:   2,
		Month: time.October}
)

var (
	// orvilleCode represents orville code.
	orvilleCode = nook.Code{
		Value: "dod"}
)

var (
	// orvilleAmericanEnglishName represents orville american english name.
	orvilleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Orville"}

	// orvilleCanadianFrenchName represents orville canadian french name.
	orvilleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Morris"}

	// orvilleDutchName represents orville dutch name.
	orvilleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Orville"}

	// orvilleFrenchName represents orville french name.
	orvilleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Morris"}

	// orvilleGermanName represents orville german name.
	orvilleGermanName = nook.Name{
		Language: language.German,
		Value:    "Bodo"}

	// orvilleItalianName represents orville italian name.
	orvilleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dodobaldo"}

	// orvilleJapaneseName represents orville japanese name.
	orvilleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モーリー"}

	// orvilleKoreanName represents orville korean name.
	orvilleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "모리"}

	// orvilleLatinAmericanSpanishName represents orville latin american spanish name.
	orvilleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rafa"}

	// orvilleRussianName represents orville russian name.
	orvilleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Орвилл"}

	// orvilleSimplifiedChineseName represents orville simplified chinese name.
	orvilleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莫里"}

	// orvilleSpanishName represents orville spanish name.
	orvilleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rafa"}

	// orvilleTraditionalChineseName represents orville traditional chinese name.
	orvilleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莫里"}
)

var (
	// orvilleName represents orville name.
	orvilleName = nook.Languages{
		language.AmericanEnglish:      orvilleAmericanEnglishName,
		language.CanadianFrench:       orvilleCanadianFrenchName,
		language.Dutch:                orvilleDutchName,
		language.French:               orvilleFrenchName,
		language.German:               orvilleGermanName,
		language.Italian:              orvilleItalianName,
		language.Japanese:             orvilleJapaneseName,
		language.Korean:               orvilleKoreanName,
		language.LatinAmericanSpanish: orvilleLatinAmericanSpanishName,
		language.Russian:              orvilleRussianName,
		language.SimplifiedChinese:    orvilleSimplifiedChineseName,
		language.Spanish:              orvilleSpanishName,
		language.TraditionalChinese:   orvilleTraditionalChineseName}
)

var (
	// orvilleCharacter represents orville character.
	orvilleCharacter = nook.Character{
		Animal:   animal.Dodo,
		Birthday: orvilleBirthday,
		Code:     orvilleCode,
		Key:      character.Orville,
		Gender:   gender.Male,
		Name:     orvilleName,
		Special:  true}
)

var (
	// Orville represents orville.
	Orville = nook.Resident{
		Character: orvilleCharacter}
)
