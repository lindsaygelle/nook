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
	// kksliderBirthday represents kkslider birthday.
	kksliderBirthday = nook.Birthday{
		Day:   23,
		Month: time.August}
)

var (
	// kksliderCode represents kkslider code.
	kksliderCode = nook.Code{
		Value: "end/tkk"}
)

var (
	// kksliderAmericanEnglishName represents kkslider american english name.
	kksliderAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "K.K. Slider"}

	// kksliderCanadianFrenchName represents kkslider canadian french name.
	kksliderCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kéké"}

	// kksliderDutchName represents kkslider dutch name.
	kksliderDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "K.K."}

	// kksliderFrenchName represents kkslider french name.
	kksliderFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kéké"}

	// kksliderGermanName represents kkslider german name.
	kksliderGermanName = nook.Name{
		Language: language.German,
		Value:    "K.K."}

	// kksliderItalianName represents kkslider italian name.
	kksliderItalianName = nook.Name{
		Language: language.Italian,
		Value:    "K.K."}

	// kksliderJapaneseName represents kkslider japanese name.
	kksliderJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "とたけけ"}

	// kksliderKoreanName represents kkslider korean name.
	kksliderKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "K.K."}

	// kksliderLatinAmericanSpanishName represents kkslider latin american spanish name.
	kksliderLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Totakeke"}

	// kksliderRussianName represents kkslider russian name.
	kksliderRussianName = nook.Name{
		Language: language.Russian,
		Value:    "К. К."}

	// kksliderSimplifiedChineseName represents kkslider simplified chinese name.
	kksliderSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "K.K."}

	// kksliderSpanishName represents kkslider spanish name.
	kksliderSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Totakeke"}

	// kksliderTraditionalChineseName represents kkslider traditional chinese name.
	kksliderTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "K.K."}
)

var (
	// kksliderName represents kkslider name.
	kksliderName = nook.Languages{
		language.AmericanEnglish:      kksliderAmericanEnglishName,
		language.CanadianFrench:       kksliderCanadianFrenchName,
		language.Dutch:                kksliderDutchName,
		language.French:               kksliderFrenchName,
		language.German:               kksliderGermanName,
		language.Italian:              kksliderItalianName,
		language.Japanese:             kksliderJapaneseName,
		language.Korean:               kksliderKoreanName,
		language.LatinAmericanSpanish: kksliderLatinAmericanSpanishName,
		language.Russian:              kksliderRussianName,
		language.SimplifiedChinese:    kksliderSimplifiedChineseName,
		language.Spanish:              kksliderSpanishName,
		language.TraditionalChinese:   kksliderTraditionalChineseName}
)

var (
	// kksliderCharacter represents kkslider character.
	kksliderCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: kksliderBirthday,
		Code:     kksliderCode,
		Key:      character.KKSlider,
		Gender:   gender.Male,
		Name:     kksliderName,
		Special:  true}
)

var (
	// KKSlider represents k k slider.
	KKSlider = nook.Resident{
		Character: kksliderCharacter}
)
