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
	grizzlyBirthday = nook.Birthday{
		Day:   31,
		Month: time.July}
)

// grizzlyCode represents Grizzly's unique code.
var (
	grizzlyCode = nook.Code{
		Value: "bea09"}
)

// Different names for Grizzly in various languages.
var (
	grizzlyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Grizzly"}

	grizzlyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Grizzly"}

	grizzlyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Grizzly"}

	grizzlyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grizzly"}

	grizzlyGermanName = nook.Name{
		Language: language.German,
		Value:    "Gerald"}

	grizzlyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grizzly"}

	grizzlyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ムー"}

	grizzlyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "무뚝"}

	grizzlyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gruñón"}

	grizzlyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гризли"}

	grizzlySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "穆穆"}

	grizzlySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gruñón"}

	grizzlyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "穆穆"}
)

// grizzlyName represents Grizzly's name in different languages.
var (
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
	grizzlyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grrr"}

	grizzlyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "canaillou"}

	grizzlyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "grrr"}

	grizzlyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "canaillou"}

	grizzlyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grrr"}

	grizzlyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrr"}

	grizzlyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だクマ"}

	grizzlyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냅둬"}

	grizzlyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrre"}

	grizzlyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "грр"}

	grizzlySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "熊熊"}

	grizzlySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tostón"}

	grizzlyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "熊熊"}
)

// grizzlyPhrase represents Grizzly's phrases in different languages.
var (
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
	Grizzly = nook.Villager{
		Character:   grizzlyCharacter,
		Personality: personality.Cranky,
		Phrase:      grizzlyPhrase}
)
