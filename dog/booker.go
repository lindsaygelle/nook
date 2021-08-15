package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bookerBirthday = nook.Birthday{
		Day:   23,
		Month: time.April}
)

var (
	bookerCode = nook.Code{
		Value: "pla/dgb"}
)

var (
	bookerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Booker"}

	bookerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chausset"}

	bookerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Booker"}

	bookerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chausset"}

	bookerGermanName = nook.Name{
		Language: language.German,
		Value:    "Wuff"}

	bookerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fido"}

	bookerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "おまわりさんB"}

	bookerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "경찰관B"}

	bookerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nocencio"}

	bookerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Букер"}

	bookerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "警察叔叔B"}

	bookerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nocencio"}

	bookerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "警察叔叔B"}
)

var (
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
	bookerCharacter = nook.Character{
		Animal:   Dog,
		Birthday: bookerBirthday,
		Code:     bookerCode,
		Gender:   nook.Male,
		Name:     bookerName}
)

var (
	Booker = nook.Resident{
		Character: bookerCharacter}
)
