package dog

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
	// beaBirthday represents bea birthday.
	beaBirthday = nook.Birthday{
		Day:   15,
		Month: time.October}
)

var (
	// beaCode represents bea code.
	beaCode = nook.Code{
		Value: "dog10"}
)

var (
	// beaAmericanEnglishName represents bea american english name.
	beaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bea"}

	// beaCanadianFrenchName represents bea canadian french name.
	beaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Béa"}

	// beaDutchName represents bea dutch name.
	beaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bea"}

	// beaFrenchName represents bea french name.
	beaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Béa"}

	// beaGermanName represents bea german name.
	beaGermanName = nook.Name{
		Language: language.German,
		Value:    "Bea"}

	// beaItalianName represents bea italian name.
	beaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cucciola"}

	// beaJapaneseName represents bea japanese name.
	beaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ベーグル"}

	// beaKoreanName represents bea korean name.
	beaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "베이글"}

	// beaLatinAmericanSpanishName represents bea latin american spanish name.
	beaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bea"}

	// beaRussianName represents bea russian name.
	beaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Беа"}

	// beaSimplifiedChineseName represents bea simplified chinese name.
	beaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贝果"}

	// beaSpanishName represents bea spanish name.
	beaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bea"}

	// beaTraditionalChineseName represents bea traditional chinese name.
	beaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "貝果"}
)

var (
	// beaName represents bea name.
	beaName = nook.Languages{
		language.AmericanEnglish:      beaAmericanEnglishName,
		language.CanadianFrench:       beaCanadianFrenchName,
		language.Dutch:                beaDutchName,
		language.French:               beaFrenchName,
		language.German:               beaGermanName,
		language.Italian:              beaItalianName,
		language.Japanese:             beaJapaneseName,
		language.Korean:               beaKoreanName,
		language.LatinAmericanSpanish: beaLatinAmericanSpanishName,
		language.Russian:              beaRussianName,
		language.SimplifiedChinese:    beaSimplifiedChineseName,
		language.Spanish:              beaSpanishName,
		language.TraditionalChinese:   beaTraditionalChineseName}
)

var (
	// beaCharacter represents bea character.
	beaCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: beaBirthday,
		Code:     beaCode,
		Key:      character.Bea,
		Gender:   gender.Female,
		Name:     beaName,
		Special:  false}
)

var (
	// beaAmericanEnglishPhrase represents bea american english phrase.
	beaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bingo"}

	// beaCanadianFrenchPhrase represents bea canadian french phrase.
	beaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon chaton"}

	// beaDutchPhrase represents bea dutch phrase.
	beaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bingo"}

	// beaFrenchPhrase represents bea french phrase.
	beaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon chaton"}

	// beaGermanPhrase represents bea german phrase.
	beaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bingo"}

	// beaItalianPhrase represents bea italian phrase.
	beaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bingo"}

	// beaJapanesePhrase represents bea japanese phrase.
	beaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "グー"}

	// beaKoreanPhrase represents bea korean phrase.
	beaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쫀득"}

	// beaLatinAmericanSpanishPhrase represents bea latin american spanish phrase.
	beaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gufi-guf"}

	// beaRussianPhrase represents bea russian phrase.
	beaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "динго"}

	// beaSimplifiedChinesePhrase represents bea simplified chinese phrase.
	beaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宾果"}

	// beaSpanishPhrase represents bea spanish phrase.
	beaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "huesín"}

	// beaTraditionalChinesePhrase represents bea traditional chinese phrase.
	beaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "賓果"}
)

var (
	// beaPhrase represents bea phrase.
	beaPhrase = nook.Languages{
		language.AmericanEnglish:      beaAmericanEnglishPhrase,
		language.CanadianFrench:       beaCanadianFrenchPhrase,
		language.Dutch:                beaDutchPhrase,
		language.French:               beaFrenchPhrase,
		language.German:               beaGermanPhrase,
		language.Italian:              beaItalianPhrase,
		language.Japanese:             beaJapanesePhrase,
		language.Korean:               beaKoreanPhrase,
		language.LatinAmericanSpanish: beaLatinAmericanSpanishPhrase,
		language.Russian:              beaRussianPhrase,
		language.SimplifiedChinese:    beaSimplifiedChinesePhrase,
		language.Spanish:              beaSpanishPhrase,
		language.TraditionalChinese:   beaTraditionalChinesePhrase}
)

var (
	// Bea represents bea.
	Bea = nook.Villager{
		Character:   beaCharacter,
		Personality: personality.Normal,
		Phrase:      beaPhrase}
)
