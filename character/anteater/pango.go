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

// pangoBirthday represents Pango's birthday.
var (
	// pangoBirthday represents pango birthday.
	pangoBirthday = nook.Birthday{
		Day:   9,
		Month: time.November}
)

// pangoCode represents Pango's unique code.
var (
	// pangoCode represents pango code.
	pangoCode = nook.Code{
		Value: "ant02"}
)

// Different names for Pango in various languages.
var (
	// pangoAmericanEnglishName represents pango american english name.
	pangoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pango"}

	// pangoCanadianFrenchName represents pango canadian french name.
	pangoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mathilda"}

	// pangoDutchName represents pango dutch name.
	pangoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pango"}

	// pangoFrenchName represents pango french name.
	pangoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mathilda"}

	// pangoGermanName represents pango german name.
	pangoGermanName = nook.Name{
		Language: language.German,
		Value:    "Mathilda"}

	// pangoItalianName represents pango italian name.
	pangoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carlotta"}

	// pangoJapaneseName represents pango japanese name.
	pangoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パトラ"}

	// pangoKoreanName represents pango korean name.
	pangoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "패트라"}

	// pangoLatinAmericanSpanishName represents pango latin american spanish name.
	pangoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aspidora"}

	// pangoRussianName represents pango russian name.
	pangoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Панго"}

	// pangoSimplifiedChineseName represents pango simplified chinese name.
	pangoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "佩希"}

	// pangoSpanishName represents pango spanish name.
	pangoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aspidora"}

	// pangoTraditionalChineseName represents pango traditional chinese name.
	pangoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "佩希"}
)

// pangoName represents Pango's name in different languages.
var (
	// pangoName represents pango name.
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
	// pangoCharacter represents pango character.
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
	// pangoAmericanEnglishPhrase represents pango american english phrase.
	pangoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snooooof"}

	// pangoCanadianFrenchPhrase represents pango canadian french phrase.
	pangoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pouuuuuf"}

	// pangoDutchPhrase represents pango dutch phrase.
	pangoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snufffff"}

	// pangoFrenchPhrase represents pango french phrase.
	pangoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pouuuuuf"}

	// pangoGermanPhrase represents pango german phrase.
	pangoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnieef"}

	// pangoItalianPhrase represents pango italian phrase.
	pangoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snuuf"}

	// pangoJapanesePhrase represents pango japanese phrase.
	pangoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だっしー"}

	// pangoKoreanPhrase represents pango korean phrase.
	pangoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라지요"}

	// pangoLatinAmericanSpanishPhrase represents pango latin american spanish phrase.
	pangoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "snuf-snuf"}

	// pangoRussianPhrase represents pango russian phrase.
	pangoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "снуф-снуф"}

	// pangoSimplifiedChinesePhrase represents pango simplified chinese phrase.
	pangoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "希希"}

	// pangoSpanishPhrase represents pango spanish phrase.
	pangoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "snuf-snuf"}

	// pangoTraditionalChinesePhrase represents pango traditional chinese phrase.
	pangoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "希希"}
)

// pangoPhrase represents Pango's phrases in different languages.
var (
	// pangoPhrase represents pango phrase.
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
	// Pango represents pango.
	Pango = nook.Villager{
		Character:   pangoCharacter,
		Personality: personality.Peppy,
		Phrase:      pangoPhrase}
)
