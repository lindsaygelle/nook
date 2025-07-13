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

// anabelleBirthday represents Anabelle's birthday.
var (
	// anabelleBirthday represents anabelle birthday.
	anabelleBirthday = nook.Birthday{
		Day:   16,
		Month: time.February}
)

// anabelleCode represents Anabelle's unique code.
var (
	// anabelleCode represents anabelle code.
	anabelleCode = nook.Code{
		Value: "ant03"}
)

// Different names for Anabelle in various languages.
var (
	// anabelleAmericanEnglishName represents anabelle american english name.
	anabelleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Anabelle"}

	// anabelleCanadianFrenchName represents anabelle canadian french name.
	anabelleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Anabelle"}

	// anabelleDutchName represents anabelle dutch name.
	anabelleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Anabelle"}

	// anabelleFrenchName represents anabelle french name.
	anabelleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Anabelle"}

	// anabelleGermanName represents anabelle german name.
	anabelleGermanName = nook.Name{
		Language: language.German,
		Value:    "Annabell"}

	// anabelleItalianName represents anabelle italian name.
	anabelleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Annalisa"}

	// anabelleJapaneseName represents anabelle japanese name.
	anabelleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "あるみ"}

	// anabelleKoreanName represents anabelle korean name.
	anabelleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아롱이"}

	// anabelleLatinAmericanSpanishName represents anabelle latin american spanish name.
	anabelleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Anabel"}

	// anabelleRussianName represents anabelle russian name.
	anabelleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Анабель"}

	// anabelleSimplifiedChineseName represents anabelle simplified chinese name.
	anabelleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "有美"}

	// anabelleSpanishName represents anabelle spanish name.
	anabelleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Anabel"}

	// anabelleTraditionalChineseName represents anabelle traditional chinese name.
	anabelleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "有美"}
)

// anabelleName represents Anabelle's name in different languages.
var (
	// anabelleName represents anabelle name.
	anabelleName = nook.Languages{
		language.AmericanEnglish:      anabelleAmericanEnglishName,
		language.CanadianFrench:       anabelleCanadianFrenchName,
		language.Dutch:                anabelleDutchName,
		language.French:               anabelleFrenchName,
		language.German:               anabelleGermanName,
		language.Italian:              anabelleItalianName,
		language.Japanese:             anabelleJapaneseName,
		language.Korean:               anabelleKoreanName,
		language.LatinAmericanSpanish: anabelleLatinAmericanSpanishName,
		language.Russian:              anabelleRussianName,
		language.SimplifiedChinese:    anabelleSimplifiedChineseName,
		language.Spanish:              anabelleSpanishName,
		language.TraditionalChinese:   anabelleTraditionalChineseName}
)

// anabelleCharacter represents Anabelle's character information.
var (
	// anabelleCharacter represents anabelle character.
	anabelleCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: anabelleBirthday,
		Code:     anabelleCode,
		Key:      character.Anabelle,
		Gender:   gender.Female,
		Name:     anabelleName,
		Special:  false}
)

// Different phrases spoken by Anabelle in various languages.
var (
	// anabelleAmericanEnglishPhrase represents anabelle american english phrase.
	anabelleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snorty"}

	// anabelleCanadianFrenchPhrase represents anabelle canadian french phrase.
	anabelleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grognon"}

	// anabelleDutchPhrase represents anabelle dutch phrase.
	anabelleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snork"}

	// anabelleFrenchPhrase represents anabelle french phrase.
	anabelleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grognon"}

	// anabelleGermanPhrase represents anabelle german phrase.
	anabelleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "puuust"}

	// anabelleItalianPhrase represents anabelle italian phrase.
	anabelleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snorty"}

	// anabelleJapanesePhrase represents anabelle japanese phrase.
	anabelleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "マジでー"}

	// anabelleKoreanPhrase represents anabelle korean phrase.
	anabelleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "정말"}

	// anabelleLatinAmericanSpanishPhrase represents anabelle latin american spanish phrase.
	anabelleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fa-fiú"}

	// anabelleRussianPhrase represents anabelle russian phrase.
	anabelleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрум-хрум"}

	// anabelleSimplifiedChinesePhrase represents anabelle simplified chinese phrase.
	anabelleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "真的假的"}

	// anabelleSpanishPhrase represents anabelle spanish phrase.
	anabelleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "esnoink"}

	// anabelleTraditionalChinesePhrase represents anabelle traditional chinese phrase.
	anabelleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "真的假的"}
)

// anabellePhrase represents Anabelle's phrases in different languages.
var (
	// anabellePhrase represents anabelle phrase.
	anabellePhrase = nook.Languages{
		language.AmericanEnglish:      anabelleAmericanEnglishPhrase,
		language.CanadianFrench:       anabelleCanadianFrenchPhrase,
		language.Dutch:                anabelleDutchPhrase,
		language.French:               anabelleFrenchPhrase,
		language.German:               anabelleGermanPhrase,
		language.Italian:              anabelleItalianPhrase,
		language.Japanese:             anabelleJapanesePhrase,
		language.Korean:               anabelleKoreanPhrase,
		language.LatinAmericanSpanish: anabelleLatinAmericanSpanishPhrase,
		language.Russian:              anabelleRussianPhrase,
		language.SimplifiedChinese:    anabelleSimplifiedChinesePhrase,
		language.Spanish:              anabelleSpanishPhrase,
		language.TraditionalChinese:   anabelleTraditionalChinesePhrase}
)

// Anabelle represents the character Anabelle with her complete information.
var (
	// Anabelle represents anabelle.
	Anabelle = nook.Villager{
		Character:   anabelleCharacter,
		Personality: personality.Peppy,
		Phrase:      anabellePhrase}
)
