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

// teddyBirthday represents Teddy's birthday.
var (
	// teddyBirthday represents teddy birthday.
	teddyBirthday = nook.Birthday{
		Day:   26,
		Month: time.September}
)

// teddyCode represents Teddy's unique code.
var (
	// teddyCode represents teddy code.
	teddyCode = nook.Code{
		Value: "bea00"}
)

// Different names for Teddy in various languages.
var (
	// teddyAmericanEnglishName represents teddy american english name.
	teddyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Teddy"}

	// teddyCanadianFrenchName represents teddy canadian french name.
	teddyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Teddy"}

	// teddyDutchName represents teddy dutch name.
	teddyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Teddy"}

	// teddyFrenchName represents teddy french name.
	teddyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Teddy"}

	// teddyGermanName represents teddy german name.
	teddyGermanName = nook.Name{
		Language: language.German,
		Value:    "Torsten"}

	// teddyItalianName represents teddy italian name.
	teddyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Teddy"}

	// teddyJapaneseName represents teddy japanese name.
	teddyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たいへいた"}

	// teddyKoreanName represents teddy korean name.
	teddyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "병태"}

	// teddyLatinAmericanSpanishName represents teddy latin american spanish name.
	teddyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Teddy"}

	// teddyRussianName represents teddy russian name.
	teddyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тедди"}

	// teddySimplifiedChineseName represents teddy simplified chinese name.
	teddySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "太平"}

	// teddySpanishName represents teddy spanish name.
	teddySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Teddy"}

	// teddyTraditionalChineseName represents teddy traditional chinese name.
	teddyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "太平"}
)

// teddyName represents Teddy's name in different languages.
var (
	// teddyName represents teddy name.
	teddyName = nook.Languages{
		language.AmericanEnglish:      teddyAmericanEnglishName,
		language.CanadianFrench:       teddyCanadianFrenchName,
		language.Dutch:                teddyDutchName,
		language.French:               teddyFrenchName,
		language.German:               teddyGermanName,
		language.Italian:              teddyItalianName,
		language.Japanese:             teddyJapaneseName,
		language.Korean:               teddyKoreanName,
		language.LatinAmericanSpanish: teddyLatinAmericanSpanishName,
		language.Russian:              teddyRussianName,
		language.SimplifiedChinese:    teddySimplifiedChineseName,
		language.Spanish:              teddySpanishName,
		language.TraditionalChinese:   teddyTraditionalChineseName}
)

// teddyCharacter represents Teddy's character information.
var (
	// teddyCharacter represents teddy character.
	teddyCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: teddyBirthday,
		Code:     teddyCode,
		Key:      character.Teddy,
		Gender:   gender.Male,
		Name:     teddyName,
		Special:  false}
)

// Different phrases spoken by Teddy in various languages.
var (
	// teddyAmericanEnglishPhrase represents teddy american english phrase.
	teddyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grooof"}

	// teddyCanadianFrenchPhrase represents teddy canadian french phrase.
	teddyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grrrrrr"}

	// teddyDutchPhrase represents teddy dutch phrase.
	teddyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "grommm"}

	// teddyFrenchPhrase represents teddy french phrase.
	teddyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grrrrrr"}

	// teddyGermanPhrase represents teddy german phrase.
	teddyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gruff"}

	// teddyItalianPhrase represents teddy italian phrase.
	teddyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gruf"}

	// teddyJapanesePhrase represents teddy japanese phrase.
	teddyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですたい"}

	// teddyKoreanPhrase represents teddy korean phrase.
	teddyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "입네다"}

	// teddyLatinAmericanSpanishPhrase represents teddy latin american spanish phrase.
	teddyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gruuuf"}

	// teddyRussianPhrase represents teddy russian phrase.
	teddyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "быр-быр"}

	// teddySimplifiedChinesePhrase represents teddy simplified chinese phrase.
	teddySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "太好了"}

	// teddySpanishPhrase represents teddy spanish phrase.
	teddySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gruuuf"}

	// teddyTraditionalChinesePhrase represents teddy traditional chinese phrase.
	teddyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "太好了"}
)

// teddyPhrase represents Teddy's phrases in different languages.
var (
	// teddyPhrase represents teddy phrase.
	teddyPhrase = nook.Languages{
		language.AmericanEnglish:      teddyAmericanEnglishPhrase,
		language.CanadianFrench:       teddyCanadianFrenchPhrase,
		language.Dutch:                teddyDutchPhrase,
		language.French:               teddyFrenchPhrase,
		language.German:               teddyGermanPhrase,
		language.Italian:              teddyItalianPhrase,
		language.Japanese:             teddyJapanesePhrase,
		language.Korean:               teddyKoreanPhrase,
		language.LatinAmericanSpanish: teddyLatinAmericanSpanishPhrase,
		language.Russian:              teddyRussianPhrase,
		language.SimplifiedChinese:    teddySimplifiedChinesePhrase,
		language.Spanish:              teddySpanishPhrase,
		language.TraditionalChinese:   teddyTraditionalChinesePhrase}
)

// Teddy represents the character Teddy with his complete information.
var (
	// Teddy represents teddy.
	Teddy = nook.Villager{
		Character:   teddyCharacter,
		Personality: personality.Jock,
		Phrase:      teddyPhrase}
)
