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
	// huckBirthday represents huck birthday.
	huckBirthday = nook.Birthday{
		Day:   9,
		Month: time.July}
)

var (
	// huckCode represents huck code.
	huckCode = nook.Code{
		Value: "flg11"}
)

var (
	// huckAmericanEnglishName represents huck american english name.
	huckAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Huck"}

	// huckCanadianFrenchName represents huck canadian french name.
	huckCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bajoue"}

	// huckDutchName represents huck dutch name.
	huckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Huck"}

	// huckFrenchName represents huck french name.
	huckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bajoue"}

	// huckGermanName represents huck german name.
	huckGermanName = nook.Name{
		Language: language.German,
		Value:    "Knuth"}

	// huckItalianName represents huck italian name.
	huckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Berto"}

	// huckJapaneseName represents huck japanese name.
	huckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ストロー"}

	// huckKoreanName represents huck korean name.
	huckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스트로"}

	// huckLatinAmericanSpanishName represents huck latin american spanish name.
	huckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ricardo"}

	// huckRussianName represents huck russian name.
	huckRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хак"}

	// huckSimplifiedChineseName represents huck simplified chinese name.
	huckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吸管"}

	// huckSpanishName represents huck spanish name.
	huckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ricardo"}

	// huckTraditionalChineseName represents huck traditional chinese name.
	huckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吸管"}
)

var (
	// huckName represents huck name.
	huckName = nook.Languages{
		language.AmericanEnglish:      huckAmericanEnglishName,
		language.CanadianFrench:       huckCanadianFrenchName,
		language.Dutch:                huckDutchName,
		language.French:               huckFrenchName,
		language.German:               huckGermanName,
		language.Italian:              huckItalianName,
		language.Japanese:             huckJapaneseName,
		language.Korean:               huckKoreanName,
		language.LatinAmericanSpanish: huckLatinAmericanSpanishName,
		language.Russian:              huckRussianName,
		language.SimplifiedChinese:    huckSimplifiedChineseName,
		language.Spanish:              huckSpanishName,
		language.TraditionalChinese:   huckTraditionalChineseName}
)

var (
	// huckCharacter represents huck character.
	huckCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: huckBirthday,
		Code:     huckCode,
		Key:      character.Huck,
		Gender:   gender.Male,
		Name:     huckName,
		Special:  false}
)

var (
	// huckAmericanEnglishPhrase represents huck american english phrase.
	huckAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hopper"}

	// huckCanadianFrenchPhrase represents huck canadian french phrase.
	huckCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nouille"}

	// huckDutchPhrase represents huck dutch phrase.
	huckDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hophop"}

	// huckFrenchPhrase represents huck french phrase.
	huckFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nouille"}

	// huckGermanPhrase represents huck german phrase.
	huckGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hops"}

	// huckItalianPhrase represents huck italian phrase.
	huckItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hophop"}

	// huckJapanesePhrase represents huck japanese phrase.
	huckJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とかー"}

	// huckKoreanPhrase represents huck korean phrase.
	huckKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "싸우자"}

	// huckLatinAmericanSpanishPhrase represents huck latin american spanish phrase.
	huckLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ajop"}

	// huckRussianPhrase represents huck russian phrase.
	huckRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "скок"}

	// huckSimplifiedChinesePhrase represents huck simplified chinese phrase.
	huckSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喝喝"}

	// huckSpanishPhrase represents huck spanish phrase.
	huckSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grillito"}

	// huckTraditionalChinesePhrase represents huck traditional chinese phrase.
	huckTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喝喝"}
)

var (
	// huckPhrase represents huck phrase.
	huckPhrase = nook.Languages{
		language.AmericanEnglish:      huckAmericanEnglishPhrase,
		language.CanadianFrench:       huckCanadianFrenchPhrase,
		language.Dutch:                huckDutchPhrase,
		language.French:               huckFrenchPhrase,
		language.German:               huckGermanPhrase,
		language.Italian:              huckItalianPhrase,
		language.Japanese:             huckJapanesePhrase,
		language.Korean:               huckKoreanPhrase,
		language.LatinAmericanSpanish: huckLatinAmericanSpanishPhrase,
		language.Russian:              huckRussianPhrase,
		language.SimplifiedChinese:    huckSimplifiedChinesePhrase,
		language.Spanish:              huckSpanishPhrase,
		language.TraditionalChinese:   huckTraditionalChinesePhrase}
)

var (
	// Huck represents huck.
	Huck = nook.Villager{
		Character:   huckCharacter,
		Personality: personality.Smug,
		Phrase:      huckPhrase}
)
