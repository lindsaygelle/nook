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
	jambetteBirthday = nook.Birthday{
		Day:   27,
		Month: time.October}
)

var (
	jambetteCode = nook.Code{
		Value: "flg13"}
)

var (
	jambetteAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jambette"}

	jambetteCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gambette"}

	jambetteDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jambette"}

	jambetteFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gambette"}

	jambetteGermanName = nook.Name{
		Language: language.German,
		Value:    "Jeanette"}

	jambetteItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giada"}

	jambetteJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エスメラルダ"}

	jambetteKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "에스메랄다"}

	jambetteLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Anquita"}

	jambetteRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джембет"}

	jambetteSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "翡翠"}

	jambetteSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Anquita"}

	jambetteTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "翡翠"}
)

var (
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
	jambetteAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "croak-kay"}

	jambetteCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "quoi"}

	jambetteDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "inderkwaak"}

	jambetteFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "quoi"}

	jambetteGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quaacki"}

	jambetteItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "crok-ok"}

	jambetteJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですわよ"}

	jambetteKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쪼오옥"}

	jambetteLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "croac"}

	jambetteRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кво-кей"}

	jambetteSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是蛙"}

	jambetteSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "croac"}

	jambetteTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是蛙"}
)

var (
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
	Jambette = nook.Villager{
		Character:   jambetteCharacter,
		Personality: personality.Normal,
		Phrase:      jambettePhrase}
)
