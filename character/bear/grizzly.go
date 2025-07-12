package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// grizzlyBirthday represents Grizzly's birthday.
var (
	// grizzlyBirthday represents grizzly birthday.
	grizzlyBirthday = nook.Birthday{
		Day:   31,
		Month: time.July}
)

// grizzlyCode represents Grizzly's unique code.
var (
	// grizzlyCode represents grizzly code.
	grizzlyCode = nook.Code{
		Value: "bea09"}
)

// Different names for Grizzly in various languages.
var (
	// grizzlyAmericanEnglishName represents grizzly american english name.
	grizzlyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Grizzly"}

	// grizzlyCanadianFrenchName represents grizzly canadian french name.
	grizzlyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Grizzly"}

	// grizzlyDutchName represents grizzly dutch name.
	grizzlyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Grizzly"}

	// grizzlyFrenchName represents grizzly french name.
	grizzlyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grizzly"}

	// grizzlyGermanName represents grizzly german name.
	grizzlyGermanName = nook.Name{
		Language: language.German,
		Value:    "Gerald"}

	// grizzlyItalianName represents grizzly italian name.
	grizzlyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grizzly"}

	// grizzlyJapaneseName represents grizzly japanese name.
	grizzlyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ムー"}

	// grizzlyKoreanName represents grizzly korean name.
	grizzlyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "무뚝"}

	// grizzlyLatinAmericanSpanishName represents grizzly latin american spanish name.
	grizzlyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gruñón"}

	// grizzlyRussianName represents grizzly russian name.
	grizzlyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гризли"}

	// grizzlySimplifiedChineseName represents grizzly simplified chinese name.
	grizzlySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "穆穆"}

	// grizzlySpanishName represents grizzly spanish name.
	grizzlySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gruñón"}

	// grizzlyTraditionalChineseName represents grizzly traditional chinese name.
	grizzlyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "穆穆"}
)

// grizzlyName represents Grizzly's name in different languages.
var (
	// grizzlyName represents grizzly name.
	grizzlyName = nook.Languages{
		language.AmericanEnglish:      grizzlyAmericanEnglishName,
		language.CanadianFrench:       grizzlyCanadianFrenchName,
		language.Dutch:                grizzlyDutchName,
		language.French:               grizzlyFrenchName,
		language.German:               grizzlyGermanName,
		language.Italian:              grizzlyItalianName,
		language.Japanese:             grizzlyJapaneseName,
		language.Korean:               grizzlyKoreanName,
		language.LatinAmericanSpanish: grizzlyLatinAmericanSpanishName,
		language.Russian:              grizzlyRussianName,
		language.SimplifiedChinese:    grizzlySimplifiedChineseName,
		language.Spanish:              grizzlySpanishName,
		language.TraditionalChinese:   grizzlyTraditionalChineseName}
)

// grizzlyCharacter represents Grizzly's character information.
var (
	// grizzlyCharacter represents grizzly character.
	grizzlyCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: grizzlyBirthday,
		Code:     grizzlyCode,
		Key:      character.Grizzly,
		Gender:   gender.Male,
		Name:     grizzlyName,
		Special:  false}
)

// Different phrases spoken by Grizzly in various languages.
var (
	// grizzlyAmericanEnglishPhrase represents grizzly american english phrase.
	grizzlyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grrr"}

	// grizzlyCanadianFrenchPhrase represents grizzly canadian french phrase.
	grizzlyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "canaillou"}

	// grizzlyDutchPhrase represents grizzly dutch phrase.
	grizzlyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "grrr"}

	// grizzlyFrenchPhrase represents grizzly french phrase.
	grizzlyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "canaillou"}

	// grizzlyGermanPhrase represents grizzly german phrase.
	grizzlyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grrr"}

	// grizzlyItalianPhrase represents grizzly italian phrase.
	grizzlyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrr"}

	// grizzlyJapanesePhrase represents grizzly japanese phrase.
	grizzlyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だクマ"}

	// grizzlyKoreanPhrase represents grizzly korean phrase.
	grizzlyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냅둬"}

	// grizzlyLatinAmericanSpanishPhrase represents grizzly latin american spanish phrase.
	grizzlyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrre"}

	// grizzlyRussianPhrase represents grizzly russian phrase.
	grizzlyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "грр"}

	// grizzlySimplifiedChinesePhrase represents grizzly simplified chinese phrase.
	grizzlySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "熊熊"}

	// grizzlySpanishPhrase represents grizzly spanish phrase.
	grizzlySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tostón"}

	// grizzlyTraditionalChinesePhrase represents grizzly traditional chinese phrase.
	grizzlyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "熊熊"}
)

// grizzlyPhrase represents Grizzly's phrases in different languages.
var (
	// grizzlyPhrase represents grizzly phrase.
	grizzlyPhrase = nook.Languages{
		language.AmericanEnglish:      grizzlyAmericanEnglishPhrase,
		language.CanadianFrench:       grizzlyCanadianFrenchPhrase,
		language.Dutch:                grizzlyDutchPhrase,
		language.French:               grizzlyFrenchPhrase,
		language.German:               grizzlyGermanPhrase,
		language.Italian:              grizzlyItalianPhrase,
		language.Japanese:             grizzlyJapanesePhrase,
		language.Korean:               grizzlyKoreanPhrase,
		language.LatinAmericanSpanish: grizzlyLatinAmericanSpanishPhrase,
		language.Russian:              grizzlyRussianPhrase,
		language.SimplifiedChinese:    grizzlySimplifiedChinesePhrase,
		language.Spanish:              grizzlySpanishPhrase,
		language.TraditionalChinese:   grizzlyTraditionalChinesePhrase}
)

// Grizzly represents the character Grizzly with his complete information.
var (
	// Grizzly represents grizzly.
	Grizzly = nook.Villager{
		Character:   grizzlyCharacter,
		Personality: personality.Cranky,
		Phrase:      grizzlyPhrase}
)
