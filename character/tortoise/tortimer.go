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
	// tortimerBirthday represents tortimer birthday.
	tortimerBirthday = nook.Birthday{
		Day:   31,
		Month: time.December}
)

var (
	// tortimerCode represents tortimer code.
	tortimerCode = nook.Code{
		Value: "ttlA"}
)

var (
	// tortimerAmericanEnglishName represents tortimer american english name.
	tortimerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tortimer"}

	// tortimerCanadianFrenchName represents tortimer canadian french name.
	tortimerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tortimer"}

	// tortimerDutchName represents tortimer dutch name.
	tortimerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tortimer"}

	// tortimerFrenchName represents tortimer french name.
	tortimerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tortimer"}

	// tortimerGermanName represents tortimer german name.
	tortimerGermanName = nook.Name{
		Language: language.German,
		Value:    "Törtel"}

	// tortimerItalianName represents tortimer italian name.
	tortimerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tortimer"}

	// tortimerJapaneseName represents tortimer japanese name.
	tortimerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "コトブキ"}

	// tortimerKoreanName represents tortimer korean name.
	tortimerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고북"}

	// tortimerLatinAmericanSpanishName represents tortimer latin american spanish name.
	tortimerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tórtimer"}

	// tortimerRussianName represents tortimer russian name.
	tortimerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тортимер"}

	// tortimerSimplifiedChineseName represents tortimer simplified chinese name.
	tortimerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "寿伯"}

	// tortimerSpanishName represents tortimer spanish name.
	tortimerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tórtimer"}

	// tortimerTraditionalChineseName represents tortimer traditional chinese name.
	tortimerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "壽伯"}
)

var (
	// tortimerName represents tortimer name.
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
	// tortimerCharacter represents tortimer character.
	tortimerCharacter = nook.Character{
		Animal:   animal.Tortoise,
		Birthday: tortimerBirthday,
		Code:     tortimerCode,
		Key:      character.Tortimer,
		Gender:   gender.Male,
		Name:     tortimerName,
		Special:  true}
)

var (
	// Tortimer represents tortimer.
	Tortimer = nook.Resident{
		Character: tortimerCharacter}
)
