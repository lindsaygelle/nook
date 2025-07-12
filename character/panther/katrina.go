package panther

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// katrinaBirthday represents katrina birthday.
	katrinaBirthday = nook.Birthday{
		Day:   28,
		Month: time.October}
)

var (
	// katrinaCode represents katrina code.
	katrinaCode = nook.Code{
		Value: "bpt"}
)

var (
	// katrinaAmericanEnglishName represents katrina american english name.
	katrinaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Katrina"}

	// katrinaCanadianFrenchName represents katrina canadian french name.
	katrinaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Astrid"}

	// katrinaDutchName represents katrina dutch name.
	katrinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Katrina"}

	// katrinaFrenchName represents katrina french name.
	katrinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Astrid"}

	// katrinaGermanName represents katrina german name.
	katrinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Smeralda"}

	// katrinaItalianName represents katrina italian name.
	katrinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Vanda"}

	// katrinaJapaneseName represents katrina japanese name.
	katrinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハッケミィ"}

	// katrinaKoreanName represents katrina korean name.
	katrinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마추릴라"}

	// katrinaLatinAmericanSpanishName represents katrina latin american spanish name.
	katrinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Katrina"}

	// katrinaRussianName represents katrina russian name.
	katrinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Катрина"}

	// katrinaSimplifiedChineseName represents katrina simplified chinese name.
	katrinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "星薇"}

	// katrinaSpanishName represents katrina spanish name.
	katrinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Katrina"}

	// katrinaTraditionalChineseName represents katrina traditional chinese name.
	katrinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "星薇"}
)

var (
	// katrinaName represents katrina name.
	katrinaName = nook.Languages{
		language.AmericanEnglish:      katrinaAmericanEnglishName,
		language.CanadianFrench:       katrinaCanadianFrenchName,
		language.Dutch:                katrinaDutchName,
		language.French:               katrinaFrenchName,
		language.German:               katrinaGermanName,
		language.Italian:              katrinaItalianName,
		language.Japanese:             katrinaJapaneseName,
		language.Korean:               katrinaKoreanName,
		language.LatinAmericanSpanish: katrinaLatinAmericanSpanishName,
		language.Russian:              katrinaRussianName,
		language.SimplifiedChinese:    katrinaSimplifiedChineseName,
		language.Spanish:              katrinaSpanishName,
		language.TraditionalChinese:   katrinaTraditionalChineseName}
)

var (
	// katrinaCharacter represents katrina character.
	katrinaCharacter = nook.Character{
		Animal:   animal.Panther,
		Birthday: katrinaBirthday,
		Code:     katrinaCode,
		Key:      character.Katrina,
		Gender:   gender.Female,
		Name:     katrinaName,
		Special:  true}
)

var (
	// Katrina represents katrina.
	Katrina = nook.Resident{
		Character: katrinaCharacter}
)
