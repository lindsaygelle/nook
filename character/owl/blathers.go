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
	blathersBirthday = nook.Birthday{
		Day:   24,
		Month: time.September}
)

var (
	blathersCode = nook.Code{
		Value: "owl"}
)

var (
	blathersAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Blathers"}

	blathersCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Thibou"}

	blathersDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Blathers"}

	blathersFrenchName = nook.Name{
		Language: language.French,
		Value:    "Thibou"}

	blathersGermanName = nook.Name{
		Language: language.German,
		Value:    "Eugen"}

	blathersItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Blatero"}

	blathersJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フータ"}

	blathersKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "부엉"}

	blathersLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sócrates"}

	blathersRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Блезерс"}

	blathersSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "傅达"}

	blathersSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sócrates"}

	blathersTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "傅達"}
)

var (
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
	Blathers = nook.Resident{
		Character: blathersCharacter}
)
