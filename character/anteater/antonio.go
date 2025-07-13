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

// antonioBirthday represents Antonio's birthday.
var (
	// antonioBirthday represents antonio birthday.
	antonioBirthday = nook.Birthday{
		Day:   20,
		Month: time.October}
)

// antonioCode represents Antonio's unique code.
var (
	// antonioCode represents antonio code.
	antonioCode = nook.Code{
		Value: "ant01"}
)

// Different names for Antonio in various languages.
var (
	// antonioAmericanEnglishName represents antonio american english name.
	antonioAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Antonio"}

	// antonioCanadianFrenchName represents antonio canadian french name.
	antonioCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Antonio"}

	// antonioDutchName represents antonio dutch name.
	antonioDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Antonio"}

	// antonioFrenchName represents antonio french name.
	antonioFrenchName = nook.Name{
		Language: language.French,
		Value:    "Antonio"}

	// antonioGermanName represents antonio german name.
	antonioGermanName = nook.Name{
		Language: language.German,
		Value:    "Siggi"}

	// antonioItalianName represents antonio italian name.
	antonioItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Antonio"}

	// antonioJapaneseName represents antonio japanese name.
	antonioJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マコト"}

	// antonioKoreanName represents antonio korean name.
	antonioKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "퍼머거"}

	// antonioLatinAmericanSpanishName represents antonio latin american spanish name.
	antonioLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Antonio"}

	// antonioRussianName represents antonio russian name.
	antonioRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Антонио"}

	// antonioSimplifiedChineseName represents antonio simplified chinese name.
	antonioSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿诚"}

	// antonioSpanishName represents antonio spanish name.
	antonioSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Antonio"}

	// antonioTraditionalChineseName represents antonio traditional chinese name.
	antonioTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿誠"}
)

// antonioName represents Antonio's name in different languages.
var (
	// antonioName represents antonio name.
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
	// antonioCharacter represents antonio character.
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
	// antonioAmericanEnglishPhrase represents antonio american english phrase.
	antonioAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "honk"}

	// antonioCanadianFrenchPhrase represents antonio canadian french phrase.
	antonioCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pouet"}

	// antonioDutchPhrase represents antonio dutch phrase.
	antonioDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuit"}

	// antonioFrenchPhrase represents antonio french phrase.
	antonioFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pouet"}

	// antonioGermanPhrase represents antonio german phrase.
	antonioGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schlürrrf"}

	// antonioItalianPhrase represents antonio italian phrase.
	antonioItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "honk"}

	// antonioJapanesePhrase represents antonio japanese phrase.
	antonioJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ホントに"}

	// antonioKoreanPhrase represents antonio korean phrase.
	antonioKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "진짜로"}

	// antonioLatinAmericanSpanishPhrase represents antonio latin american spanish phrase.
	antonioLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fufuf"}

	// antonioRussianPhrase represents antonio russian phrase.
	antonioRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "го-го-го"}

	// antonioSimplifiedChinesePhrase represents antonio simplified chinese phrase.
	antonioSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "真诚"}

	// antonioSpanishPhrase represents antonio spanish phrase.
	antonioSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fufuf"}

	// antonioTraditionalChinesePhrase represents antonio traditional chinese phrase.
	antonioTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "真誠"}
)

// antonioPhrase represents Antonio's phrases in different languages.
var (
	// antonioPhrase represents antonio phrase.
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
	// Antonio represents antonio.
	Antonio = nook.Villager{
		Character:   antonioCharacter,
		Personality: personality.Jock,
		Phrase:      antonioPhrase}
)
