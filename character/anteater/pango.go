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

// pangoBirthday represents Pango's birthday (November 9th).
var (
	pangoBirthday = nook.Birthday{
		Day:   9,
		Month: time.November}
)

// pangoCode represents Pango's unique code.
var (
	pangoCode = nook.Code{
		Value: "ant02"}
)

// Different names for Pango in various languages.
var (
	pangoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pango"}

	pangoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mathilda"}

	pangoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pango"}

	pangoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mathilda"}

	pangoGermanName = nook.Name{
		Language: language.German,
		Value:    "Mathilda"}

	pangoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carlotta"}

	pangoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パトラ"}

	pangoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "패트라"}

	pangoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aspidora"}

	pangoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Панго"}

	pangoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "佩希"}

	pangoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aspidora"}

	pangoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "佩希"}
)

// pangoName represents Pango's name in different languages.
var (
	pangoName = nook.Languages{
		language.AmericanEnglish:      pangoAmericanEnglishName,
		language.CanadianFrench:       pangoCanadianFrenchName,
		language.Dutch:                pangoDutchName,
		language.French:               pangoFrenchName,
		language.German:               pangoGermanName,
		language.Italian:              pangoItalianName,
		language.Japanese:             pangoJapaneseName,
		language.Korean:               pangoKoreanName,
		language.LatinAmericanSpanish: pangoLatinAmericanSpanishName,
		language.Russian:              pangoRussianName,
		language.SimplifiedChinese:    pangoSimplifiedChineseName,
		language.Spanish:              pangoSpanishName,
		language.TraditionalChinese:   pangoTraditionalChineseName}
)

// pangoCharacter represents Pango's character information.
var (
	pangoCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: pangoBirthday,
		Code:     pangoCode,
		Key:      character.Pango,
		Gender:   gender.Female,
		Name:     pangoName,
		Special:  false}
)

// Different phrases spoken by Pango in various languages.
var (
	pangoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snooooof"}

	pangoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pouuuuuf"}

	pangoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snufffff"}

	pangoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pouuuuuf"}

	pangoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnieef"}

	pangoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snuuf"}

	pangoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だっしー"}

	pangoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라지요"}

	pangoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "snuf-snuf"}

	pangoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "снуф-снуф"}

	pangoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "希希"}

	pangoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "snuf-snuf"}

	pangoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "希希"}
)

// pangoPhrase represents Pango's phrases in different languages.
var (
	pangoPhrase = nook.Languages{
		language.AmericanEnglish:      pangoAmericanEnglishPhrase,
		language.CanadianFrench:       pangoCanadianFrenchPhrase,
		language.Dutch:                pangoDutchPhrase,
		language.French:               pangoFrenchPhrase,
		language.German:               pangoGermanPhrase,
		language.Italian:              pangoItalianPhrase,
		language.Japanese:             pangoJapanesePhrase,
		language.Korean:               pangoKoreanPhrase,
		language.LatinAmericanSpanish: pangoLatinAmericanSpanishPhrase,
		language.Russian:              pangoRussianPhrase,
		language.SimplifiedChinese:    pangoSimplifiedChinesePhrase,
		language.Spanish:              pangoSpanishPhrase,
		language.TraditionalChinese:   pangoTraditionalChinesePhrase}
)

// Pango represents the character Pango with her complete information.
var (
	Pango = nook.Villager{
		Character:   pangoCharacter,
		Personality: personality.Peppy,
		Phrase:      pangoPhrase}
)
