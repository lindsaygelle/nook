package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	// jambetteBirthday represents jambette birthday.
	jambetteBirthday = nook.Birthday{
		Day:   27,
		Month: time.October}
)

var (
	// jambetteCode represents jambette code.
	jambetteCode = nook.Code{
		Value: "flg13"}
)

var (
	// jambetteAmericanEnglishName represents jambette american english name.
	jambetteAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jambette"}

	// jambetteCanadianFrenchName represents jambette canadian french name.
	jambetteCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gambette"}

	// jambetteDutchName represents jambette dutch name.
	jambetteDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jambette"}

	// jambetteFrenchName represents jambette french name.
	jambetteFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gambette"}

	// jambetteGermanName represents jambette german name.
	jambetteGermanName = nook.Name{
		Language: language.German,
		Value:    "Jeanette"}

	// jambetteItalianName represents jambette italian name.
	jambetteItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giada"}

	// jambetteJapaneseName represents jambette japanese name.
	jambetteJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エスメラルダ"}

	// jambetteKoreanName represents jambette korean name.
	jambetteKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에스메랄다"}

	// jambetteLatinAmericanSpanishName represents jambette latin american spanish name.
	jambetteLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Anquita"}

	// jambetteRussianName represents jambette russian name.
	jambetteRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джембет"}

	// jambetteSimplifiedChineseName represents jambette simplified chinese name.
	jambetteSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "翡翠"}

	// jambetteSpanishName represents jambette spanish name.
	jambetteSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Anquita"}

	// jambetteTraditionalChineseName represents jambette traditional chinese name.
	jambetteTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "翡翠"}
)

var (
	// jambetteName represents jambette name.
	jambetteName = nook.Languages{
		language.AmericanEnglish:      jambetteAmericanEnglishName,
		language.CanadianFrench:       jambetteCanadianFrenchName,
		language.Dutch:                jambetteDutchName,
		language.French:               jambetteFrenchName,
		language.German:               jambetteGermanName,
		language.Italian:              jambetteItalianName,
		language.Japanese:             jambetteJapaneseName,
		language.Korean:               jambetteKoreanName,
		language.LatinAmericanSpanish: jambetteLatinAmericanSpanishName,
		language.Russian:              jambetteRussianName,
		language.SimplifiedChinese:    jambetteSimplifiedChineseName,
		language.Spanish:              jambetteSpanishName,
		language.TraditionalChinese:   jambetteTraditionalChineseName}
)

var (
	// jambetteCharacter represents jambette character.
	jambetteCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: jambetteBirthday,
		Code:     jambetteCode,
		Key:      character.Jambette,
		Gender:   gender.Female,
		Name:     jambetteName,
		Special:  false}
)

var (
	// jambetteAmericanEnglishPhrase represents jambette american english phrase.
	jambetteAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "croak-kay"}

	// jambetteCanadianFrenchPhrase represents jambette canadian french phrase.
	jambetteCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "quoi"}

	// jambetteDutchPhrase represents jambette dutch phrase.
	jambetteDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "inderkwaak"}

	// jambetteFrenchPhrase represents jambette french phrase.
	jambetteFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "quoi"}

	// jambetteGermanPhrase represents jambette german phrase.
	jambetteGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quaacki"}

	// jambetteItalianPhrase represents jambette italian phrase.
	jambetteItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "crok-ok"}

	// jambetteJapanesePhrase represents jambette japanese phrase.
	jambetteJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですわよ"}

	// jambetteKoreanPhrase represents jambette korean phrase.
	jambetteKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쪼오옥"}

	// jambetteLatinAmericanSpanishPhrase represents jambette latin american spanish phrase.
	jambetteLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "croac"}

	// jambetteRussianPhrase represents jambette russian phrase.
	jambetteRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кво-кей"}

	// jambetteSimplifiedChinesePhrase represents jambette simplified chinese phrase.
	jambetteSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是蛙"}

	// jambetteSpanishPhrase represents jambette spanish phrase.
	jambetteSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "croac"}

	// jambetteTraditionalChinesePhrase represents jambette traditional chinese phrase.
	jambetteTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是蛙"}
)

var (
	// jambettePhrase represents jambette phrase.
	jambettePhrase = nook.Languages{
		language.AmericanEnglish:      jambetteAmericanEnglishPhrase,
		language.CanadianFrench:       jambetteCanadianFrenchPhrase,
		language.Dutch:                jambetteDutchPhrase,
		language.French:               jambetteFrenchPhrase,
		language.German:               jambetteGermanPhrase,
		language.Italian:              jambetteItalianPhrase,
		language.Japanese:             jambetteJapanesePhrase,
		language.Korean:               jambetteKoreanPhrase,
		language.LatinAmericanSpanish: jambetteLatinAmericanSpanishPhrase,
		language.Russian:              jambetteRussianPhrase,
		language.SimplifiedChinese:    jambetteSimplifiedChinesePhrase,
		language.Spanish:              jambetteSpanishPhrase,
		language.TraditionalChinese:   jambetteTraditionalChinesePhrase}
)

var (
	// Jambette represents jambette.
	Jambette = nook.Villager{
		Character:   jambetteCharacter,
		Personality: personality.Normal,
		Phrase:      jambettePhrase}
)
