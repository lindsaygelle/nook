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

// pironkonBirthday represents Pironkon's birthday.
var (
	// pironkonBirthday represents pironkon birthday.
	pironkonBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

// pironkonCode represents Pironkon's unique code.
var (
	// pironkonCode represents pironkon code.
	pironkonCode = nook.Code{
		Value: ""}
)

// Different names for Pironkon in various languages.
var (
	// pironkonAmericanEnglishName represents pironkon american english name.
	pironkonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pironkon"}

	// pironkonCanadianFrenchName represents pironkon canadian french name.
	pironkonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// pironkonDutchName represents pironkon dutch name.
	pironkonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// pironkonFrenchName represents pironkon french name.
	pironkonFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// pironkonGermanName represents pironkon german name.
	pironkonGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// pironkonItalianName represents pironkon italian name.
	pironkonItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// pironkonJapaneseName represents pironkon japanese name.
	pironkonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピロンコン"}

	// pironkonKoreanName represents pironkon korean name.
	pironkonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// pironkonLatinAmericanSpanishName represents pironkon latin american spanish name.
	pironkonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// pironkonRussianName represents pironkon russian name.
	pironkonRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// pironkonSimplifiedChineseName represents pironkon simplified chinese name.
	pironkonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// pironkonSpanishName represents pironkon spanish name.
	pironkonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// pironkonTraditionalChineseName represents pironkon traditional chinese name.
	pironkonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// pironkonName represents Pironkon's name in different languages.
var (
	// pironkonName represents pironkon name.
	pironkonName = nook.Languages{
		language.AmericanEnglish:      pironkonAmericanEnglishName,
		language.CanadianFrench:       pironkonCanadianFrenchName,
		language.Dutch:                pironkonDutchName,
		language.French:               pironkonFrenchName,
		language.German:               pironkonGermanName,
		language.Italian:              pironkonItalianName,
		language.Japanese:             pironkonJapaneseName,
		language.Korean:               pironkonKoreanName,
		language.LatinAmericanSpanish: pironkonLatinAmericanSpanishName,
		language.Russian:              pironkonRussianName,
		language.SimplifiedChinese:    pironkonSimplifiedChineseName,
		language.Spanish:              pironkonSpanishName,
		language.TraditionalChinese:   pironkonTraditionalChineseName}
)

// pironkonCharacter represents Pironkon's character information.
var (
	// pironkonCharacter represents pironkon character.
	pironkonCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: pironkonBirthday,
		Code:     pironkonCode,
		Key:      character.Pironkon,
		Gender:   gender.Male,
		Name:     pironkonName,
		Special:  false}
)

// Different phrases spoken by Pironkon in various languages.
var (
	// pironkonAmericanEnglishPhrase represents pironkon american english phrase.
	pironkonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "パヨマケ"}

	// pironkonCanadianFrenchPhrase represents pironkon canadian french phrase.
	pironkonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// pironkonDutchPhrase represents pironkon dutch phrase.
	pironkonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// pironkonFrenchPhrase represents pironkon french phrase.
	pironkonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// pironkonGermanPhrase represents pironkon german phrase.
	pironkonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// pironkonItalianPhrase represents pironkon italian phrase.
	pironkonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// pironkonJapanesePhrase represents pironkon japanese phrase.
	pironkonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "パヨマケ"}

	// pironkonKoreanPhrase represents pironkon korean phrase.
	pironkonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// pironkonLatinAmericanSpanishPhrase represents pironkon latin american spanish phrase.
	pironkonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// pironkonRussianPhrase represents pironkon russian phrase.
	pironkonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// pironkonSimplifiedChinesePhrase represents pironkon simplified chinese phrase.
	pironkonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// pironkonSpanishPhrase represents pironkon spanish phrase.
	pironkonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// pironkonTraditionalChinesePhrase represents pironkon traditional chinese phrase.
	pironkonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// pironkonPhrase represents Pironkon's phrases in different languages.
var (
	// pironkonPhrase represents pironkon phrase.
	pironkonPhrase = nook.Languages{
		language.AmericanEnglish:      pironkonAmericanEnglishPhrase,
		language.CanadianFrench:       pironkonCanadianFrenchPhrase,
		language.Dutch:                pironkonDutchPhrase,
		language.French:               pironkonFrenchPhrase,
		language.German:               pironkonGermanPhrase,
		language.Italian:              pironkonItalianPhrase,
		language.Japanese:             pironkonJapanesePhrase,
		language.Korean:               pironkonKoreanPhrase,
		language.LatinAmericanSpanish: pironkonLatinAmericanSpanishPhrase,
		language.Russian:              pironkonRussianPhrase,
		language.SimplifiedChinese:    pironkonSimplifiedChinesePhrase,
		language.Spanish:              pironkonSpanishPhrase,
		language.TraditionalChinese:   pironkonTraditionalChinesePhrase}
)

// Pironkon represents the character Pironkon with his complete information.
var (
	// Pironkon represents pironkon.
	Pironkon = nook.Villager{
		Character:   pironkonCharacter,
		Personality: personality.Lazy,
		Phrase:      pironkonPhrase}
)
