package monkey

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// porterBirthday represents porter birthday.
	porterBirthday = nook.Birthday{
		Day:   17,
		Month: time.April}
)

var (
	// porterCode represents porter code.
	porterCode = nook.Code{
		Value: "mnk"}
)

var (
	// porterAmericanEnglishName represents porter american english name.
	porterAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Porter"}

	// porterCanadianFrenchName represents porter canadian french name.
	porterCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lazare"}

	// porterDutchName represents porter dutch name.
	porterDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Porter"}

	// porterFrenchName represents porter french name.
	porterFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lazare"}

	// porterGermanName represents porter german name.
	porterGermanName = nook.Name{
		Language: language.German,
		Value:    "Flip"}

	// porterItalianName represents porter italian name.
	porterItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ciufciuf"}

	// porterJapaneseName represents porter japanese name.
	porterJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "えきいんさん"}

	// porterKoreanName represents porter korean name.
	porterKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "역무원"}

	// porterLatinAmericanSpanishName represents porter latin american spanish name.
	porterLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Estasio"}

	// porterRussianName represents porter russian name.
	porterRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Портер"}

	// porterSimplifiedChineseName represents porter simplified chinese name.
	porterSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "车长叔叔"}

	// porterSpanishName represents porter spanish name.
	porterSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Estasio"}

	// porterTraditionalChineseName represents porter traditional chinese name.
	porterTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "車長叔叔"}
)

var (
	// porterName represents porter name.
	porterName = nook.Languages{
		language.AmericanEnglish:      porterAmericanEnglishName,
		language.CanadianFrench:       porterCanadianFrenchName,
		language.Dutch:                porterDutchName,
		language.French:               porterFrenchName,
		language.German:               porterGermanName,
		language.Italian:              porterItalianName,
		language.Japanese:             porterJapaneseName,
		language.Korean:               porterKoreanName,
		language.LatinAmericanSpanish: porterLatinAmericanSpanishName,
		language.Russian:              porterRussianName,
		language.SimplifiedChinese:    porterSimplifiedChineseName,
		language.Spanish:              porterSpanishName,
		language.TraditionalChinese:   porterTraditionalChineseName}
)

var (
	// porterCharacter represents porter character.
	porterCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: porterBirthday,
		Code:     porterCode,
		Key:      character.Porter,
		Gender:   gender.Male,
		Name:     porterName,
		Special:  true}
)

var (
	// Porter represents porter.
	Porter = nook.Resident{
		Character: porterCharacter}
)
