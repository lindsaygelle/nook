package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// antonioBirthday represents Antonio's birthday (October 20th).
var (
	antonioBirthday = nook.Birthday{
		Day:   20,
		Month: time.October}
)

// antonioCode represents Antonio's unique code ("ant01").
var (
	antonioCode = nook.Code{
		Value: "ant01"}
)

// Different names for Antonio in various languages.
var (
	antonioAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Antonio"}

	antonioCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Antonio"}

	antonioDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Antonio"}

	antonioFrenchName = nook.Name{
		Language: language.French,
		Value:    "Antonio"}

	antonioGermanName = nook.Name{
		Language: language.German,
		Value:    "Siggi"}

	antonioItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Antonio"}

	antonioJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マコト"}

	antonioKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "퍼머거"}

	antonioLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Antonio"}

	antonioRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Антонио"}

	antonioSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿诚"}

	antonioSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Antonio"}

	antonioTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿誠"}
)

// antonioName represents Antonio's name in different languages.
var (
	antonioName = nook.Languages{
		language.AmericanEnglish:      antonioAmericanEnglishName,
		language.CanadianFrench:       antonioCanadianFrenchName,
		language.Dutch:                antonioDutchName,
		language.French:               antonioFrenchName,
		language.German:               antonioGermanName,
		language.Italian:              antonioItalianName,
		language.Japanese:             antonioJapaneseName,
		language.Korean:               antonioKoreanName,
		language.LatinAmericanSpanish: antonioLatinAmericanSpanishName,
		language.Russian:              antonioRussianName,
		language.SimplifiedChinese:    antonioSimplifiedChineseName,
		language.Spanish:              antonioSpanishName,
		language.TraditionalChinese:   antonioTraditionalChineseName}
)

// antonioCharacter represents Antonio's character information.
var (
	antonioCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: antonioBirthday,
		Code:     antonioCode,
		Key:      character.Antonio,
		Gender:   gender.Male,
		Name:     antonioName,
		Special:  false}
)

// Different phrases spoken by Antonio in various languages.
var (
	antonioAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "honk"}

	antonioCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pouet"}

	antonioDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuit"}

	antonioFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pouet"}

	antonioGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schlürrrf"}

	antonioItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "honk"}

	antonioJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ホントに"}

	antonioKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "진짜로"}

	antonioLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fufuf"}

	antonioRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "го-го-го"}

	antonioSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "真诚"}

	antonioSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fufuf"}

	antonioTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "真誠"}
)

// antonioPhrase represents Antonio's phrases in different languages.
var (
	antonioPhrase = nook.Languages{
		language.AmericanEnglish:      antonioAmericanEnglishPhrase,
		language.CanadianFrench:       antonioCanadianFrenchPhrase,
		language.Dutch:                antonioDutchPhrase,
		language.French:               antonioFrenchPhrase,
		language.German:               antonioGermanPhrase,
		language.Italian:              antonioItalianPhrase,
		language.Japanese:             antonioJapanesePhrase,
		language.Korean:               antonioKoreanPhrase,
		language.LatinAmericanSpanish: antonioLatinAmericanSpanishPhrase,
		language.Russian:              antonioRussianPhrase,
		language.SimplifiedChinese:    antonioSimplifiedChinesePhrase,
		language.Spanish:              antonioSpanishPhrase,
		language.TraditionalChinese:   antonioTraditionalChinesePhrase}
)

// Antonio represents the character Antonio with his complete information.
var (
	Antonio = nook.Villager{
		Character:   antonioCharacter,
		Personality: personality.Jock,
		Phrase:      antonioPhrase}
)
