package frillneckedlizard

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// frillardBirthday represents frillard birthday.
	frillardBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// frillardCode represents frillard code.
	frillardCode = nook.Code{
		Value: "liz"}
)

var (
	// frillardAmericanEnglishName represents frillard american english name.
	frillardAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Frillard"}

	// frillardCanadianFrenchName represents frillard canadian french name.
	frillardCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Unknown"}

	// frillardDutchName represents frillard dutch name.
	frillardDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// frillardFrenchName represents frillard french name.
	frillardFrenchName = nook.Name{
		Language: language.French,
		Value:    "Unknown"}

	// frillardGermanName represents frillard german name.
	frillardGermanName = nook.Name{
		Language: language.German,
		Value:    "Eqsos"}

	// frillardItalianName represents frillard italian name.
	frillardItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Oscar"}

	// frillardJapaneseName represents frillard japanese name.
	frillardJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "きょしょー"}

	// frillardKoreanName represents frillard korean name.
	frillardKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "Unknown"}

	// frillardLatinAmericanSpanishName represents frillard latin american spanish name.
	frillardLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Unknown"}

	// frillardRussianName represents frillard russian name.
	frillardRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// frillardSimplifiedChineseName represents frillard simplified chinese name.
	frillardSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// frillardSpanishName represents frillard spanish name.
	frillardSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Unknown"}

	// frillardTraditionalChineseName represents frillard traditional chinese name.
	frillardTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// frillardName represents frillard name.
	frillardName = nook.Languages{
		language.AmericanEnglish:      frillardAmericanEnglishName,
		language.CanadianFrench:       frillardCanadianFrenchName,
		language.Dutch:                frillardDutchName,
		language.French:               frillardFrenchName,
		language.German:               frillardGermanName,
		language.Italian:              frillardItalianName,
		language.Japanese:             frillardJapaneseName,
		language.Korean:               frillardKoreanName,
		language.LatinAmericanSpanish: frillardLatinAmericanSpanishName,
		language.Russian:              frillardRussianName,
		language.SimplifiedChinese:    frillardSimplifiedChineseName,
		language.Spanish:              frillardSpanishName,
		language.TraditionalChinese:   frillardTraditionalChineseName}
)

var (
	// frillardCharacter represents frillard character.
	frillardCharacter = nook.Character{
		Animal:   animal.FrillneckedLizard,
		Birthday: frillardBirthday,
		Code:     frillardCode,
		Key:      character.Frillard,
		Gender:   gender.Male,
		Name:     frillardName,
		Special:  true}
)

var (
	// Frillard represents frillard.
	Frillard = nook.Resident{
		Character: frillardCharacter}
)
