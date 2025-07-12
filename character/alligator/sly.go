package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// slyBirthday represents Sly's birthday.
var (
	// slyBirthday represents sly birthday.
	slyBirthday = nook.Birthday{
		Day:   15,
		Month: time.November}
)

// slyCode represents Sly's unique code.
var (
	// slyCode represents sly code.
	slyCode = nook.Code{
		Value: "crd06"}
)

// Different names for Sly in various languages.
var (
	// slyAmericanEnglishName represents sly american english name.
	slyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sly"}

	// slyCanadianFrenchName represents sly canadian french name.
	slyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chuck"}

	// slyDutchName represents sly dutch name.
	slyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sly"}

	// slyFrenchName represents sly french name.
	slyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chuck"}

	// slyGermanName represents sly german name.
	slyGermanName = nook.Name{
		Language: language.German,
		Value:    "Steve"}

	// slyItalianName represents sly italian name.
	slyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Driberto"}

	// slyJapaneseName represents sly japanese name.
	slyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハイド"}

	// slyKoreanName represents sly korean name.
	slyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "하이드"}

	// slyLatinAmericanSpanishName represents sly latin american spanish name.
	slyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Estallón"}

	// slyRussianName represents sly russian name.
	slyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Слай"}

	// slySimplifiedChineseName represents sly simplified chinese name.
	slySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "海德"}

	// slySpanishName represents sly spanish name.
	slySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Estallón"}

	// slyTraditionalChineseName represents sly traditional chinese name.
	slyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "海德"}
)

// slyName represents Sly's name in different languages.
var (
	// slyName represents sly name.
	slyName = nook.Languages{
		language.AmericanEnglish:      slyAmericanEnglishName,
		language.CanadianFrench:       slyCanadianFrenchName,
		language.Dutch:                slyDutchName,
		language.French:               slyFrenchName,
		language.German:               slyGermanName,
		language.Italian:              slyItalianName,
		language.Japanese:             slyJapaneseName,
		language.Korean:               slyKoreanName,
		language.LatinAmericanSpanish: slyLatinAmericanSpanishName,
		language.Russian:              slyRussianName,
		language.SimplifiedChinese:    slySimplifiedChineseName,
		language.Spanish:              slySpanishName,
		language.TraditionalChinese:   slyTraditionalChineseName}
)

// slyCharacter represents Sly's character information.
var (
	// slyCharacter represents sly character.
	slyCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: slyBirthday,
		Code:     slyCode,
		Key:      character.Sly,
		Gender:   gender.Male,
		Name:     slyName,
		Special:  false}
)

// Different phrases spoken by Sly in various languages.
var (
	// slyAmericanEnglishPhrase represents sly american english phrase.
	slyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hoo-rah"}

	// slyCanadianFrenchPhrase represents sly canadian french phrase.
	slyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "repos"}

	// slyDutchPhrase represents sly dutch phrase.
	slyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ingerukt"}

	// slyFrenchPhrase represents sly french phrase.
	slyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "repos"}

	// slyGermanPhrase represents sly german phrase.
	slyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnapp"}

	// slyItalianPhrase represents sly italian phrase.
	slyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ciaf"}

	// slyJapanesePhrase represents sly japanese phrase.
	slyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "カサコソ"}

	// slyKoreanPhrase represents sly korean phrase.
	slyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "부스럭"}

	// slyLatinAmericanSpanishPhrase represents sly latin american spanish phrase.
	slyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "smack"}

	// slyRussianPhrase represents sly russian phrase.
	slyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ура-а-а"}

	// slySimplifiedChinesePhrase represents sly simplified chinese phrase.
	slySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唰唰"}

	// slySpanishPhrase represents sly spanish phrase.
	slySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "capón"}

	// slyTraditionalChinesePhrase represents sly traditional chinese phrase.
	slyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唰唰"}
)

// slyPhrase represents Sly's phrases in different languages.
var (
	// slyPhrase represents sly phrase.
	slyPhrase = nook.Languages{
		language.AmericanEnglish:      slyAmericanEnglishPhrase,
		language.CanadianFrench:       slyCanadianFrenchPhrase,
		language.Dutch:                slyDutchPhrase,
		language.French:               slyFrenchPhrase,
		language.German:               slyGermanPhrase,
		language.Italian:              slyItalianPhrase,
		language.Japanese:             slyJapanesePhrase,
		language.Korean:               slyKoreanPhrase,
		language.LatinAmericanSpanish: slyLatinAmericanSpanishPhrase,
		language.Russian:              slyRussianPhrase,
		language.SimplifiedChinese:    slySimplifiedChinesePhrase,
		language.Spanish:              slySpanishPhrase,
		language.TraditionalChinese:   slyTraditionalChinesePhrase}
)

// Sly represents the character Sly with his complete information.
var (
	// Sly represents sly.
	Sly = nook.Villager{
		Character:   slyCharacter,
		Personality: personality.Jock,
		Phrase:      slyPhrase}
)
