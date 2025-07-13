package elephant

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
	// eloiseBirthday represents eloise birthday.
	eloiseBirthday = nook.Birthday{
		Day:   8,
		Month: time.December}
)

var (
	// eloiseCode represents eloise code.
	eloiseCode = nook.Code{
		Value: "elp03"}
)

var (
	// eloiseAmericanEnglishName represents eloise american english name.
	eloiseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Eloise"}

	// eloiseCanadianFrenchName represents eloise canadian french name.
	eloiseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Éloïse"}

	// eloiseDutchName represents eloise dutch name.
	eloiseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Eloise"}

	// eloiseFrenchName represents eloise french name.
	eloiseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Éloïse"}

	// eloiseGermanName represents eloise german name.
	eloiseGermanName = nook.Name{
		Language: language.German,
		Value:    "Frauke"}

	// eloiseItalianName represents eloise italian name.
	eloiseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Eloisa"}

	// eloiseJapaneseName represents eloise japanese name.
	eloiseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エレフィン"}

	// eloiseKoreanName represents eloise korean name.
	eloiseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "엘레핀"}

	// eloiseLatinAmericanSpanishName represents eloise latin american spanish name.
	eloiseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eloísa"}

	// eloiseRussianName represents eloise russian name.
	eloiseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элоиза"}

	// eloiseSimplifiedChineseName represents eloise simplified chinese name.
	eloiseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艾勒芬"}

	// eloiseSpanishName represents eloise spanish name.
	eloiseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eloísa"}

	// eloiseTraditionalChineseName represents eloise traditional chinese name.
	eloiseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艾勒芬"}
)

var (
	// eloiseName represents eloise name.
	eloiseName = nook.Languages{
		language.AmericanEnglish:      eloiseAmericanEnglishName,
		language.CanadianFrench:       eloiseCanadianFrenchName,
		language.Dutch:                eloiseDutchName,
		language.French:               eloiseFrenchName,
		language.German:               eloiseGermanName,
		language.Italian:              eloiseItalianName,
		language.Japanese:             eloiseJapaneseName,
		language.Korean:               eloiseKoreanName,
		language.LatinAmericanSpanish: eloiseLatinAmericanSpanishName,
		language.Russian:              eloiseRussianName,
		language.SimplifiedChinese:    eloiseSimplifiedChineseName,
		language.Spanish:              eloiseSpanishName,
		language.TraditionalChinese:   eloiseTraditionalChineseName}
)

var (
	// eloiseCharacter represents eloise character.
	eloiseCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: eloiseBirthday,
		Code:     eloiseCode,
		Key:      character.Eloise,
		Gender:   gender.Female,
		Name:     eloiseName,
		Special:  false}
)

var (
	// eloiseAmericanEnglishPhrase represents eloise american english phrase.
	eloiseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tooooot"}

	// eloiseCanadianFrenchPhrase represents eloise canadian french phrase.
	eloiseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tooouut"}

	// eloiseDutchPhrase represents eloise dutch phrase.
	eloiseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pwèèèèèp"}

	// eloiseFrenchPhrase represents eloise french phrase.
	eloiseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tooouut"}

	// eloiseGermanPhrase represents eloise german phrase.
	eloiseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tuutuuuu"}

	// eloiseItalianPhrase represents eloise italian phrase.
	eloiseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "nappole"}

	// eloiseJapanesePhrase represents eloise japanese phrase.
	eloiseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ルン"}

	// eloiseKoreanPhrase represents eloise korean phrase.
	eloiseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "훨훨"}

	// eloiseLatinAmericanSpanishPhrase represents eloise latin american spanish phrase.
	eloiseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fruuuf"}

	// eloiseRussianPhrase represents eloise russian phrase.
	eloiseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тру-тут"}

	// eloiseSimplifiedChinesePhrase represents eloise simplified chinese phrase.
	eloiseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鲁鲁"}

	// eloiseSpanishPhrase represents eloise spanish phrase.
	eloiseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "orejillas"}

	// eloiseTraditionalChinesePhrase represents eloise traditional chinese phrase.
	eloiseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "魯魯"}
)

var (
	// eloisePhrase represents eloise phrase.
	eloisePhrase = nook.Languages{
		language.AmericanEnglish:      eloiseAmericanEnglishPhrase,
		language.CanadianFrench:       eloiseCanadianFrenchPhrase,
		language.Dutch:                eloiseDutchPhrase,
		language.French:               eloiseFrenchPhrase,
		language.German:               eloiseGermanPhrase,
		language.Italian:              eloiseItalianPhrase,
		language.Japanese:             eloiseJapanesePhrase,
		language.Korean:               eloiseKoreanPhrase,
		language.LatinAmericanSpanish: eloiseLatinAmericanSpanishPhrase,
		language.Russian:              eloiseRussianPhrase,
		language.SimplifiedChinese:    eloiseSimplifiedChinesePhrase,
		language.Spanish:              eloiseSpanishPhrase,
		language.TraditionalChinese:   eloiseTraditionalChinesePhrase}
)

var (
	// Eloise represents eloise.
	Eloise = nook.Villager{
		Character:   eloiseCharacter,
		Personality: personality.Snooty,
		Phrase:      eloisePhrase}
)
