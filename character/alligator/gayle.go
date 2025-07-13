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

// gayleBirthday represents Gayle's birthday.
var (
	// gayleBirthday represents gayle birthday.
	gayleBirthday = nook.Birthday{
		Day:   17,
		Month: time.May}
)

// gayleCode represents Gayle's unique code.
var (
	// gayleCode represents gayle code.
	gayleCode = nook.Code{
		Value: "crd07"}
)

// Different names for Gayle in various languages.
var (
	// gayleAmericanEnglishName represents gayle american english name.
	gayleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gayle"}

	// gayleCanadianFrenchName represents gayle canadian french name.
	gayleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Odile"}

	// gayleDutchName represents gayle dutch name.
	gayleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gayle"}

	// gayleFrenchName represents gayle french name.
	gayleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Odile"}

	// gayleGermanName represents gayle german name.
	gayleGermanName = nook.Name{
		Language: language.German,
		Value:    "Rosa"}

	// gayleItalianName represents gayle italian name.
	gayleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Codrilla"}

	// gayleJapaneseName represents gayle japanese name.
	gayleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アリゲッティ"}

	// gayleKoreanName represents gayle korean name.
	gayleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "앨리"}

	// gayleLatinAmericanSpanishName represents gayle latin american spanish name.
	gayleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Boni"}

	// gayleRussianName represents gayle russian name.
	gayleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гейл"}

	// gayleSimplifiedChineseName represents gayle simplified chinese name.
	gayleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "爱莉"}

	// gayleSpanishName represents gayle spanish name.
	gayleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Boni"}

	// gayleTraditionalChineseName represents gayle traditional chinese name.
	gayleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "愛莉"}
)

// gayleName represents Gayle's name in different languages.
var (
	// gayleName represents gayle name.
	gayleName = nook.Languages{
		language.AmericanEnglish:      gayleAmericanEnglishName,
		language.CanadianFrench:       gayleCanadianFrenchName,
		language.Dutch:                gayleDutchName,
		language.French:               gayleFrenchName,
		language.German:               gayleGermanName,
		language.Italian:              gayleItalianName,
		language.Japanese:             gayleJapaneseName,
		language.Korean:               gayleKoreanName,
		language.LatinAmericanSpanish: gayleLatinAmericanSpanishName,
		language.Russian:              gayleRussianName,
		language.SimplifiedChinese:    gayleSimplifiedChineseName,
		language.Spanish:              gayleSpanishName,
		language.TraditionalChinese:   gayleTraditionalChineseName}
)

// gayleCharacter represents Gayle's character information.
var (
	// gayleCharacter represents gayle character.
	gayleCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: gayleBirthday,
		Code:     gayleCode,
		Key:      character.Gayle,
		Gender:   gender.Female,
		Name:     gayleName,
		Special:  false}
)

// Different phrases spoken by Gayle in various languages.
var (
	// gayleAmericanEnglishPhrase represents gayle american english phrase.
	gayleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snacky"}

	// gayleCanadianFrenchPhrase represents gayle canadian french phrase.
	gayleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "croc-croc"}

	// gayleDutchPhrase represents gayle dutch phrase.
	gayleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snackje"}

	// gayleFrenchPhrase represents gayle french phrase.
	gayleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croquignou"}

	// gayleGermanPhrase represents gayle german phrase.
	gayleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zähni"}

	// gayleItalianPhrase represents gayle italian phrase.
	gayleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "désolée"}

	// gayleJapanesePhrase represents gayle japanese phrase.
	gayleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワニャン"}

	// gayleKoreanPhrase represents gayle korean phrase.
	gayleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아거얌"}

	// gayleLatinAmericanSpanishPhrase represents gayle latin american spanish phrase.
	gayleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "smuack"}

	// gayleRussianPhrase represents gayle russian phrase.
	gayleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "перекус"}

	// gayleSimplifiedChinesePhrase represents gayle simplified chinese phrase.
	gayleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鳄莉"}

	// gayleSpanishPhrase represents gayle spanish phrase.
	gayleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "corazón"}

	// gayleTraditionalChinesePhrase represents gayle traditional chinese phrase.
	gayleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鱷莉"}
)

// gaylePhrase represents Gayle's phrases in different languages.
var (
	// gaylePhrase represents gayle phrase.
	gaylePhrase = nook.Languages{
		language.AmericanEnglish:      gayleAmericanEnglishPhrase,
		language.CanadianFrench:       gayleCanadianFrenchPhrase,
		language.Dutch:                gayleDutchPhrase,
		language.French:               gayleFrenchPhrase,
		language.German:               gayleGermanPhrase,
		language.Italian:              gayleItalianPhrase,
		language.Japanese:             gayleJapanesePhrase,
		language.Korean:               gayleKoreanPhrase,
		language.LatinAmericanSpanish: gayleLatinAmericanSpanishPhrase,
		language.Russian:              gayleRussianPhrase,
		language.SimplifiedChinese:    gayleSimplifiedChinesePhrase,
		language.Spanish:              gayleSpanishPhrase,
		language.TraditionalChinese:   gayleTraditionalChinesePhrase}
)

// Gayle represents the character Gayle with her complete information.
var (
	// Gayle represents gayle.
	Gayle = nook.Villager{
		Character:   gayleCharacter,
		Personality: personality.Normal,
		Phrase:      gaylePhrase}
)
