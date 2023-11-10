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

// charliseBirthday represents Charlise's birthday (April 17th).
var (
	charliseBirthday = nook.Birthday{
		Day:   17,
		Month: time.April}
)

// charliseCode represents Charlise's unique code (bea12).
var (
	charliseCode = nook.Code{
		Value: "bea12"}
)

// Different names for Charlise in various languages.
var (
	charliseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Charlise"}

	charliseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Zabou"}

	charliseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Charlise"}

	charliseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Zabou"}

	charliseGermanName = nook.Name{
		Language: language.German,
		Value:    "Tabea"}

	charliseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Orsola"}

	charliseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャーミー"}

	charliseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "챠미"}

	charliseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Charo"}

	charliseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шарлиз"}

	charliseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "恰咪"}

	charliseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Charo"}

	charliseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "恰咪"}
)

// charliseName represents Charlise's name in different languages.
var (
	charliseName = nook.Languages{
		language.AmericanEnglish:      charliseAmericanEnglishName,
		language.CanadianFrench:       charliseCanadianFrenchName,
		language.Dutch:                charliseDutchName,
		language.French:               charliseFrenchName,
		language.German:               charliseGermanName,
		language.Italian:              charliseItalianName,
		language.Japanese:             charliseJapaneseName,
		language.Korean:               charliseKoreanName,
		language.LatinAmericanSpanish: charliseLatinAmericanSpanishName,
		language.Russian:              charliseRussianName,
		language.SimplifiedChinese:    charliseSimplifiedChineseName,
		language.Spanish:              charliseSpanishName,
		language.TraditionalChinese:   charliseTraditionalChineseName}
)

// charliseCharacter represents Charlise's character information.
var (
	charliseCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: charliseBirthday,
		Code:     charliseCode,
		Key:      character.Charlise,
		Gender:   gender.Female,
		Name:     charliseName,
		Special:  false}
)

// Different phrases spoken by Charlise in various languages.
var (
	charliseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "urgh"}

	charliseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chacha"}

	charliseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pff"}

	charliseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "verlue"}

	charliseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bärbär"}

	charliseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sbuff"}

	charliseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よいしょ"}

	charliseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "으랏차차"}

	charliseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "urrrgh"}

	charliseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ого-го"}

	charliseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唷咻"}

	charliseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "comolooyes"}

	charliseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唷咻"}
)

// charlisePhrase represents Charlise's phrases in different languages.
var (
	charlisePhrase = nook.Languages{
		language.AmericanEnglish:      charliseAmericanEnglishPhrase,
		language.CanadianFrench:       charliseCanadianFrenchPhrase,
		language.Dutch:                charliseDutchPhrase,
		language.French:               charliseFrenchPhrase,
		language.German:               charliseGermanPhrase,
		language.Italian:              charliseItalianPhrase,
		language.Japanese:             charliseJapanesePhrase,
		language.Korean:               charliseKoreanPhrase,
		language.LatinAmericanSpanish: charliseLatinAmericanSpanishPhrase,
		language.Russian:              charliseRussianPhrase,
		language.SimplifiedChinese:    charliseSimplifiedChinesePhrase,
		language.Spanish:              charliseSpanishPhrase,
		language.TraditionalChinese:   charliseTraditionalChinesePhrase}
)

// Charlise represents the character Charlise with her complete information.
var (
	Charlise = nook.Villager{
		Character:   charliseCharacter,
		Personality: personality.BigSister,
		Phrase:      charlisePhrase}
)
