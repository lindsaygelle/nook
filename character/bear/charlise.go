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

// charliseBirthday represents Charlise's birthday.
var (
	// charliseBirthday represents charlise birthday.
	charliseBirthday = nook.Birthday{
		Day:   17,
		Month: time.April}
)

// charliseCode represents Charlise's unique code.
var (
	// charliseCode represents charlise code.
	charliseCode = nook.Code{
		Value: "bea12"}
)

// Different names for Charlise in various languages.
var (
	// charliseAmericanEnglishName represents charlise american english name.
	charliseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Charlise"}

	// charliseCanadianFrenchName represents charlise canadian french name.
	charliseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Zabou"}

	// charliseDutchName represents charlise dutch name.
	charliseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Charlise"}

	// charliseFrenchName represents charlise french name.
	charliseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Zabou"}

	// charliseGermanName represents charlise german name.
	charliseGermanName = nook.Name{
		Language: language.German,
		Value:    "Tabea"}

	// charliseItalianName represents charlise italian name.
	charliseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Orsola"}

	// charliseJapaneseName represents charlise japanese name.
	charliseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャーミー"}

	// charliseKoreanName represents charlise korean name.
	charliseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "챠미"}

	// charliseLatinAmericanSpanishName represents charlise latin american spanish name.
	charliseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Charo"}

	// charliseRussianName represents charlise russian name.
	charliseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шарлиз"}

	// charliseSimplifiedChineseName represents charlise simplified chinese name.
	charliseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "恰咪"}

	// charliseSpanishName represents charlise spanish name.
	charliseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Charo"}

	// charliseTraditionalChineseName represents charlise traditional chinese name.
	charliseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "恰咪"}
)

// charliseName represents Charlise's name in different languages.
var (
	// charliseName represents charlise name.
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
	// charliseCharacter represents charlise character.
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
	// charliseAmericanEnglishPhrase represents charlise american english phrase.
	charliseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "urgh"}

	// charliseCanadianFrenchPhrase represents charlise canadian french phrase.
	charliseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chacha"}

	// charliseDutchPhrase represents charlise dutch phrase.
	charliseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pff"}

	// charliseFrenchPhrase represents charlise french phrase.
	charliseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "verlue"}

	// charliseGermanPhrase represents charlise german phrase.
	charliseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bärbär"}

	// charliseItalianPhrase represents charlise italian phrase.
	charliseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sbuff"}

	// charliseJapanesePhrase represents charlise japanese phrase.
	charliseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よいしょ"}

	// charliseKoreanPhrase represents charlise korean phrase.
	charliseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "으랏차차"}

	// charliseLatinAmericanSpanishPhrase represents charlise latin american spanish phrase.
	charliseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "urrrgh"}

	// charliseRussianPhrase represents charlise russian phrase.
	charliseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ого-го"}

	// charliseSimplifiedChinesePhrase represents charlise simplified chinese phrase.
	charliseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唷咻"}

	// charliseSpanishPhrase represents charlise spanish phrase.
	charliseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "comolooyes"}

	// charliseTraditionalChinesePhrase represents charlise traditional chinese phrase.
	charliseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唷咻"}
)

// charlisePhrase represents Charlise's phrases in different languages.
var (
	// charlisePhrase represents charlise phrase.
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
	// Charlise represents charlise.
	Charlise = nook.Villager{
		Character:   charliseCharacter,
		Personality: personality.BigSister,
		Phrase:      charlisePhrase}
)
