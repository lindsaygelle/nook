package owl

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// blathersBirthday represents blathers birthday.
	blathersBirthday = nook.Birthday{
		Day:   24,
		Month: time.September}
)

var (
	// blathersCode represents blathers code.
	blathersCode = nook.Code{
		Value: "owl"}
)

var (
	// blathersAmericanEnglishName represents blathers american english name.
	blathersAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Blathers"}

	// blathersCanadianFrenchName represents blathers canadian french name.
	blathersCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Thibou"}

	// blathersDutchName represents blathers dutch name.
	blathersDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Blathers"}

	// blathersFrenchName represents blathers french name.
	blathersFrenchName = nook.Name{
		Language: language.French,
		Value:    "Thibou"}

	// blathersGermanName represents blathers german name.
	blathersGermanName = nook.Name{
		Language: language.German,
		Value:    "Eugen"}

	// blathersItalianName represents blathers italian name.
	blathersItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Blatero"}

	// blathersJapaneseName represents blathers japanese name.
	blathersJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フータ"}

	// blathersKoreanName represents blathers korean name.
	blathersKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "부엉"}

	// blathersLatinAmericanSpanishName represents blathers latin american spanish name.
	blathersLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sócrates"}

	// blathersRussianName represents blathers russian name.
	blathersRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Блезерс"}

	// blathersSimplifiedChineseName represents blathers simplified chinese name.
	blathersSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "傅达"}

	// blathersSpanishName represents blathers spanish name.
	blathersSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sócrates"}

	// blathersTraditionalChineseName represents blathers traditional chinese name.
	blathersTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "傅達"}
)

var (
	// blathersName represents blathers name.
	blathersName = nook.Languages{
		language.AmericanEnglish:      blathersAmericanEnglishName,
		language.CanadianFrench:       blathersCanadianFrenchName,
		language.Dutch:                blathersDutchName,
		language.French:               blathersFrenchName,
		language.German:               blathersGermanName,
		language.Italian:              blathersItalianName,
		language.Japanese:             blathersJapaneseName,
		language.Korean:               blathersKoreanName,
		language.LatinAmericanSpanish: blathersLatinAmericanSpanishName,
		language.Russian:              blathersRussianName,
		language.SimplifiedChinese:    blathersSimplifiedChineseName,
		language.Spanish:              blathersSpanishName,
		language.TraditionalChinese:   blathersTraditionalChineseName}
)

var (
	// blathersCharacter represents blathers character.
	blathersCharacter = nook.Character{
		Animal:   animal.Owl,
		Birthday: blathersBirthday,
		Code:     blathersCode,
		Key:      character.Blathers,
		Gender:   gender.Male,
		Name:     blathersName,
		Special:  true}
)

var (
	// Blathers represents blathers.
	Blathers = nook.Resident{
		Character: blathersCharacter}
)
