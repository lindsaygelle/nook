package duck

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
	// deenaBirthday represents deena birthday.
	deenaBirthday = nook.Birthday{
		Day:   27,
		Month: time.June}
)

var (
	// deenaCode represents deena code.
	deenaCode = nook.Code{
		Value: "duk04"}
)

var (
	// deenaAmericanEnglishName represents deena american english name.
	deenaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Deena"}

	// deenaCanadianFrenchName represents deena canadian french name.
	deenaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mina"}

	// deenaDutchName represents deena dutch name.
	deenaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Deena"}

	// deenaFrenchName represents deena french name.
	deenaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mina"}

	// deenaGermanName represents deena german name.
	deenaGermanName = nook.Name{
		Language: language.German,
		Value:    "Olivia"}

	// deenaItalianName represents deena italian name.
	deenaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dina"}

	// deenaJapaneseName represents deena japanese name.
	deenaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "まりも"}

	// deenaKoreanName represents deena korean name.
	deenaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마리모"}

	// deenaLatinAmericanSpanishName represents deena latin american spanish name.
	deenaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Martita"}

	// deenaRussianName represents deena russian name.
	deenaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дина"}

	// deenaSimplifiedChineseName represents deena simplified chinese name.
	deenaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "球藻"}

	// deenaSpanishName represents deena spanish name.
	deenaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Martita"}

	// deenaTraditionalChineseName represents deena traditional chinese name.
	deenaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毬藻"}
)

var (
	// deenaName represents deena name.
	deenaName = nook.Languages{
		language.AmericanEnglish:      deenaAmericanEnglishName,
		language.CanadianFrench:       deenaCanadianFrenchName,
		language.Dutch:                deenaDutchName,
		language.French:               deenaFrenchName,
		language.German:               deenaGermanName,
		language.Italian:              deenaItalianName,
		language.Japanese:             deenaJapaneseName,
		language.Korean:               deenaKoreanName,
		language.LatinAmericanSpanish: deenaLatinAmericanSpanishName,
		language.Russian:              deenaRussianName,
		language.SimplifiedChinese:    deenaSimplifiedChineseName,
		language.Spanish:              deenaSpanishName,
		language.TraditionalChinese:   deenaTraditionalChineseName}
)

var (
	// deenaCharacter represents deena character.
	deenaCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: deenaBirthday,
		Code:     deenaCode,
		Key:      character.Deena,
		Gender:   gender.Female,
		Name:     deenaName,
		Special:  false}
)

var (
	// deenaAmericanEnglishPhrase represents deena american english phrase.
	deenaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "woowoo"}

	// deenaCanadianFrenchPhrase represents deena canadian french phrase.
	deenaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "youhou"}

	// deenaDutchPhrase represents deena dutch phrase.
	deenaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "joehoe"}

	// deenaFrenchPhrase represents deena french phrase.
	deenaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "youhou"}

	// deenaGermanPhrase represents deena german phrase.
	deenaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "juhuu"}

	// deenaItalianPhrase represents deena italian phrase.
	deenaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quiproquò"}

	// deenaJapanesePhrase represents deena japanese phrase.
	deenaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "マル"}

	// deenaKoreanPhrase represents deena korean phrase.
	deenaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우우"}

	// deenaLatinAmericanSpanishPhrase represents deena latin american spanish phrase.
	deenaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "yiuju"}

	// deenaRussianPhrase represents deena russian phrase.
	deenaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кря-кря"}

	// deenaSimplifiedChinesePhrase represents deena simplified chinese phrase.
	deenaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "球球"}

	// deenaSpanishPhrase represents deena spanish phrase.
	deenaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "miguita"}

	// deenaTraditionalChinesePhrase represents deena traditional chinese phrase.
	deenaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毬毬"}
)

var (
	// deenaPhrase represents deena phrase.
	deenaPhrase = nook.Languages{
		language.AmericanEnglish:      deenaAmericanEnglishPhrase,
		language.CanadianFrench:       deenaCanadianFrenchPhrase,
		language.Dutch:                deenaDutchPhrase,
		language.French:               deenaFrenchPhrase,
		language.German:               deenaGermanPhrase,
		language.Italian:              deenaItalianPhrase,
		language.Japanese:             deenaJapanesePhrase,
		language.Korean:               deenaKoreanPhrase,
		language.LatinAmericanSpanish: deenaLatinAmericanSpanishPhrase,
		language.Russian:              deenaRussianPhrase,
		language.SimplifiedChinese:    deenaSimplifiedChinesePhrase,
		language.Spanish:              deenaSpanishPhrase,
		language.TraditionalChinese:   deenaTraditionalChinesePhrase}
)

var (
	// Deena represents deena.
	Deena = nook.Villager{
		Character:   deenaCharacter,
		Personality: personality.Normal,
		Phrase:      deenaPhrase}
)
