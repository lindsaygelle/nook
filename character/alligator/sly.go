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

// slyBirthday represents Sly's birthday (November 15th).
var (
	slyBirthday = nook.Birthday{
		Day:   15,
		Month: time.November}
)

// slyCode represents Sly's unique code ("crd06").
var (
	slyCode = nook.Code{
		Value: "crd06"}
)

// Different names for Sly in various languages.
var (
	slyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sly"}

	slyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chuck"}

	slyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sly"}

	slyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chuck"}

	slyGermanName = nook.Name{
		Language: language.German,
		Value:    "Steve"}

	slyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Driberto"}

	slyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハイド"}

	slyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "하이드"}

	slyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Estallón"}

	slyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Слай"}

	slySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "海德"}

	slySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Estallón"}

	slyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "海德"}
)

// slyName represents Sly's name in different languages.
var (
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
	slyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hoo-rah"}

	slyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "repos"}

	slyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ingerukt"}

	slyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "repos"}

	slyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnapp"}

	slyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ciaf"}

	slyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "カサコソ"}

	slyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "부스럭"}

	slyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "smack"}

	slyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ура-а-а"}

	slySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唰唰"}

	slySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "capón"}

	slyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唰唰"}
)

// slyPhrase represents Sly's phrases in different languages.
var (
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
	Sly = nook.Villager{
		Character:   slyCharacter,
		Personality: personality.Jock,
		Phrase:      slyPhrase}
)
