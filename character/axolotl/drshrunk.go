package axolotl

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// drshrunkBirthday represents drshrunk birthday.
	drshrunkBirthday = nook.Birthday{
		Day:   2,
		Month: time.January}
)

var (
	// drshrunkCode represents drshrunk code.
	drshrunkCode = nook.Code{
		Value: "upa"}
)

var (
	// drshrunkAmericanEnglishName represents drshrunk american english name.
	drshrunkAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dr. Shrunk"}

	// drshrunkCanadianFrenchName represents drshrunk canadian french name.
	drshrunkCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ciboulot"}

	// drshrunkDutchName represents drshrunk dutch name.
	drshrunkDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Shrunk"}

	// drshrunkFrenchName represents drshrunk french name.
	drshrunkFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ciboulot"}

	// drshrunkGermanName represents drshrunk german name.
	drshrunkGermanName = nook.Name{
		Language: language.German,
		Value:    "Samselt"}

	// drshrunkItalianName represents drshrunk italian name.
	drshrunkItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Strizzo"}

	// drshrunkJapaneseName represents drshrunk japanese name.
	drshrunkJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ししょー"}

	// drshrunkKoreanName represents drshrunk korean name.
	drshrunkKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스승"}

	// drshrunkLatinAmericanSpanishName represents drshrunk latin american spanish name.
	drshrunkLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dr. Sito"}

	// drshrunkRussianName represents drshrunk russian name.
	drshrunkRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шранк"}

	// drshrunkSimplifiedChineseName represents drshrunk simplified chinese name.
	drshrunkSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "笑匠"}

	// drshrunkSpanishName represents drshrunk spanish name.
	drshrunkSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Dr. Sito"}

	// drshrunkTraditionalChineseName represents drshrunk traditional chinese name.
	drshrunkTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "笑匠"}
)

var (
	// drshrunkName represents drshrunk name.
	drshrunkName = nook.Languages{
		language.AmericanEnglish:      drshrunkAmericanEnglishName,
		language.CanadianFrench:       drshrunkCanadianFrenchName,
		language.Dutch:                drshrunkDutchName,
		language.French:               drshrunkFrenchName,
		language.German:               drshrunkGermanName,
		language.Italian:              drshrunkItalianName,
		language.Japanese:             drshrunkJapaneseName,
		language.Korean:               drshrunkKoreanName,
		language.LatinAmericanSpanish: drshrunkLatinAmericanSpanishName,
		language.Russian:              drshrunkRussianName,
		language.SimplifiedChinese:    drshrunkSimplifiedChineseName,
		language.Spanish:              drshrunkSpanishName,
		language.TraditionalChinese:   drshrunkTraditionalChineseName}
)

var (
	// drshrunkCharacter represents drshrunk character.
	drshrunkCharacter = nook.Character{
		Animal:   animal.Axolotl,
		Birthday: drshrunkBirthday,
		Code:     drshrunkCode,
		Key:      character.DrShrunk,
		Gender:   gender.Male,
		Name:     drshrunkName,
		Special:  true}
)

var (
	// DrShrunk represents dr shrunk.
	DrShrunk = nook.Resident{
		Character: drshrunkCharacter}
)
