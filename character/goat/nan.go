package goat

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
	// nanBirthday represents nan birthday.
	nanBirthday = nook.Birthday{
		Day:   24,
		Month: time.August}
)

var (
	// nanCode represents nan code.
	nanCode = nook.Code{
		Value: "goa01"}
)

var (
	// nanAmericanEnglishName represents nan american english name.
	nanAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nan"}

	// nanCanadianFrenchName represents nan canadian french name.
	nanCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nana"}

	// nanDutchName represents nan dutch name.
	nanDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nan"}

	// nanFrenchName represents nan french name.
	nanFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nana"}

	// nanGermanName represents nan german name.
	nanGermanName = nook.Name{
		Language: language.German,
		Value:    "Zenobi"}

	// nanItalianName represents nan italian name.
	nanItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nan"}

	// nanJapaneseName represents nan japanese name.
	nanJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スミ"}

	// nanKoreanName represents nan korean name.
	nanKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "순이"}

	// nanLatinAmericanSpanishName represents nan latin american spanish name.
	nanLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pécora"}

	// nanRussianName represents nan russian name.
	nanRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нань"}

	// nanSimplifiedChineseName represents nan simplified chinese name.
	nanSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "墨墨"}

	// nanSpanishName represents nan spanish name.
	nanSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pécora"}

	// nanTraditionalChineseName represents nan traditional chinese name.
	nanTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "墨墨"}
)

var (
	// nanName represents nan name.
	nanName = nook.Languages{
		language.AmericanEnglish:      nanAmericanEnglishName,
		language.CanadianFrench:       nanCanadianFrenchName,
		language.Dutch:                nanDutchName,
		language.French:               nanFrenchName,
		language.German:               nanGermanName,
		language.Italian:              nanItalianName,
		language.Japanese:             nanJapaneseName,
		language.Korean:               nanKoreanName,
		language.LatinAmericanSpanish: nanLatinAmericanSpanishName,
		language.Russian:              nanRussianName,
		language.SimplifiedChinese:    nanSimplifiedChineseName,
		language.Spanish:              nanSpanishName,
		language.TraditionalChinese:   nanTraditionalChineseName}
)

var (
	// nanCharacter represents nan character.
	nanCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: nanBirthday,
		Code:     nanCode,
		Key:      character.Nan,
		Gender:   gender.Female,
		Name:     nanName,
		Special:  false}
)

var (
	// nanAmericanEnglishPhrase represents nan american english phrase.
	nanAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "kid"}

	// nanCanadianFrenchPhrase represents nan canadian french phrase.
	nanCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "fromage"}

	// nanDutchPhrase represents nan dutch phrase.
	nanDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "geitje"}

	// nanFrenchPhrase represents nan french phrase.
	nanFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "fromage"}

	// nanGermanPhrase represents nan german phrase.
	nanGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mekker"}

	// nanItalianPhrase represents nan italian phrase.
	nanItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bimbi"}

	// nanJapanesePhrase represents nan japanese phrase.
	nanJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "っしょ"}

	// nanKoreanPhrase represents nan korean phrase.
	nanKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그랬슈"}

	// nanLatinAmericanSpanishPhrase represents nan latin american spanish phrase.
	nanLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bibi-bee"}

	// nanRussianPhrase represents nan russian phrase.
	nanRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "козленок"}

	// nanSimplifiedChinesePhrase represents nan simplified chinese phrase.
	nanSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "书"}

	// nanSpanishPhrase represents nan spanish phrase.
	nanSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "peque"}

	// nanTraditionalChinesePhrase represents nan traditional chinese phrase.
	nanTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "書"}
)

var (
	// nanPhrase represents nan phrase.
	nanPhrase = nook.Languages{
		language.AmericanEnglish:      nanAmericanEnglishPhrase,
		language.CanadianFrench:       nanCanadianFrenchPhrase,
		language.Dutch:                nanDutchPhrase,
		language.French:               nanFrenchPhrase,
		language.German:               nanGermanPhrase,
		language.Italian:              nanItalianPhrase,
		language.Japanese:             nanJapanesePhrase,
		language.Korean:               nanKoreanPhrase,
		language.LatinAmericanSpanish: nanLatinAmericanSpanishPhrase,
		language.Russian:              nanRussianPhrase,
		language.SimplifiedChinese:    nanSimplifiedChinesePhrase,
		language.Spanish:              nanSpanishPhrase,
		language.TraditionalChinese:   nanTraditionalChinesePhrase}
)

var (
	// Nan represents nan.
	Nan = nook.Villager{
		Character:   nanCharacter,
		Personality: personality.Normal,
		Phrase:      nanPhrase}
)
