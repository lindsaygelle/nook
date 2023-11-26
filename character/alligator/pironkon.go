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
	pironkonBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

// pironkonCode represents Pironkon's unique code.
var (
	pironkonCode = nook.Code{
		Value: ""}
)

// Different names for Pironkon in various languages.
var (
	pironkonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pironkon"}

	pironkonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	pironkonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	pironkonFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	pironkonGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	pironkonItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	pironkonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピロンコン"}

	pironkonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	pironkonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	pironkonRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	pironkonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	pironkonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	pironkonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// pironkonName represents Pironkon's name in different languages.
var (
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
	pironkonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "パヨマケ"}

	pironkonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	pironkonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	pironkonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	pironkonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	pironkonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	pironkonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "パヨマケ"}

	pironkonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	pironkonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	pironkonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	pironkonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	pironkonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	pironkonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// pironkonPhrase represents Pironkon's phrases in different languages.
var (
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
	Pironkon = nook.Villager{
		Character:   pironkonCharacter,
		Personality: personality.Lazy,
		Phrase:      pironkonPhrase}
)
