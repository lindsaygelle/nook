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
	kksliderBirthday = nook.Birthday{
		Day:   23,
		Month: time.August}
)

var (
	kksliderCode = nook.Code{
		Value: "end/tkk"}
)

var (
	kksliderAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "K.K. Slider"}

	kksliderCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kéké"}

	kksliderDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "K.K."}

	kksliderFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kéké"}

	kksliderGermanName = nook.Name{
		Language: language.German,
		Value:    "K.K."}

	kksliderItalianName = nook.Name{
		Language: language.Italian,
		Value:    "K.K."}

	kksliderJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "とたけけ"}

	kksliderKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "K.K."}

	kksliderLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Totakeke"}

	kksliderRussianName = nook.Name{
		Language: language.Russian,
		Value:    "К. К."}

	kksliderSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "K.K."}

	kksliderSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Totakeke"}

	kksliderTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "K.K."}
)

var (
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
	kksliderCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: kksliderBirthday,
		Code:     kksliderCode,
		Key:      character.KKSlider,
		Gender:   gender.Male,
		Name:     kksliderName}
)

var (
	KKSlider = nook.Resident{
		Character: kksliderCharacter}
)
