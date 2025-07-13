package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// bookerBirthday represents booker birthday.
	bookerBirthday = nook.Birthday{
		Day:   23,
		Month: time.April}
)

var (
	// bookerCode represents booker code.
	bookerCode = nook.Code{
		Value: "pla/dgb"}
)

var (
	// bookerAmericanEnglishName represents booker american english name.
	bookerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Booker"}

	// bookerCanadianFrenchName represents booker canadian french name.
	bookerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chausset"}

	// bookerDutchName represents booker dutch name.
	bookerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Booker"}

	// bookerFrenchName represents booker french name.
	bookerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chausset"}

	// bookerGermanName represents booker german name.
	bookerGermanName = nook.Name{
		Language: language.German,
		Value:    "Wuff"}

	// bookerItalianName represents booker italian name.
	bookerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fido"}

	// bookerJapaneseName represents booker japanese name.
	bookerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "おまわりさんB"}

	// bookerKoreanName represents booker korean name.
	bookerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "경찰관B"}

	// bookerLatinAmericanSpanishName represents booker latin american spanish name.
	bookerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nocencio"}

	// bookerRussianName represents booker russian name.
	bookerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Букер"}

	// bookerSimplifiedChineseName represents booker simplified chinese name.
	bookerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "警察叔叔B"}

	// bookerSpanishName represents booker spanish name.
	bookerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nocencio"}

	// bookerTraditionalChineseName represents booker traditional chinese name.
	bookerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "警察叔叔B"}
)

var (
	// bookerName represents booker name.
	bookerName = nook.Languages{
		language.AmericanEnglish:      bookerAmericanEnglishName,
		language.CanadianFrench:       bookerCanadianFrenchName,
		language.Dutch:                bookerDutchName,
		language.French:               bookerFrenchName,
		language.German:               bookerGermanName,
		language.Italian:              bookerItalianName,
		language.Japanese:             bookerJapaneseName,
		language.Korean:               bookerKoreanName,
		language.LatinAmericanSpanish: bookerLatinAmericanSpanishName,
		language.Russian:              bookerRussianName,
		language.SimplifiedChinese:    bookerSimplifiedChineseName,
		language.Spanish:              bookerSpanishName,
		language.TraditionalChinese:   bookerTraditionalChineseName}
)

var (
	// bookerCharacter represents booker character.
	bookerCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: bookerBirthday,
		Code:     bookerCode,
		Key:      character.Booker,
		Gender:   gender.Male,
		Name:     bookerName,
		Special:  true}
)

var (
	// Booker represents booker.
	Booker = nook.Resident{
		Character: bookerCharacter}
)
