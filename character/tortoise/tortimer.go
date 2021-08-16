package tortoise

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	tortimerBirthday = nook.Birthday{
		Day:   31,
		Month: time.December}
)

var (
	tortimerCode = nook.Code{
		Value: "ttlA"}
)

var (
	tortimerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tortimer"}

	tortimerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tortimer"}

	tortimerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tortimer"}

	tortimerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tortimer"}

	tortimerGermanName = nook.Name{
		Language: language.German,
		Value:    "Törtel"}

	tortimerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tortimer"}

	tortimerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "コトブキ"}

	tortimerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고북"}

	tortimerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tórtimer"}

	tortimerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тортимер"}

	tortimerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "寿伯"}

	tortimerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tórtimer"}

	tortimerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "壽伯"}
)

var (
	tortimerName = nook.Languages{
		language.AmericanEnglish:      tortimerAmericanEnglishName,
		language.CanadianFrench:       tortimerCanadianFrenchName,
		language.Dutch:                tortimerDutchName,
		language.French:               tortimerFrenchName,
		language.German:               tortimerGermanName,
		language.Italian:              tortimerItalianName,
		language.Japanese:             tortimerJapaneseName,
		language.Korean:               tortimerKoreanName,
		language.LatinAmericanSpanish: tortimerLatinAmericanSpanishName,
		language.Russian:              tortimerRussianName,
		language.SimplifiedChinese:    tortimerSimplifiedChineseName,
		language.Spanish:              tortimerSpanishName,
		language.TraditionalChinese:   tortimerTraditionalChineseName}
)

var (
	tortimerCharacter = nook.Character{
		Animal:   animal.Tortoise,
		Birthday: tortimerBirthday,
		Code:     tortimerCode,
		Key:      character.Tortimer,
		Gender:   gender.Male,
		Name:     tortimerName}
)

var (
	Tortimer = nook.Resident{
		Character: tortimerCharacter}
)
