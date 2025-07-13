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
	// masaBirthday represents masa birthday.
	masaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// masaCode represents masa code.
	masaCode = nook.Code{
		Value: ""}
)

var (
	// masaAmericanEnglishName represents masa american english name.
	masaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Masa"}

	// masaCanadianFrenchName represents masa canadian french name.
	masaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// masaDutchName represents masa dutch name.
	masaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// masaFrenchName represents masa french name.
	masaFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// masaGermanName represents masa german name.
	masaGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// masaItalianName represents masa italian name.
	masaItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// masaJapaneseName represents masa japanese name.
	masaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マサ"}

	// masaKoreanName represents masa korean name.
	masaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// masaLatinAmericanSpanishName represents masa latin american spanish name.
	masaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// masaRussianName represents masa russian name.
	masaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// masaSimplifiedChineseName represents masa simplified chinese name.
	masaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// masaSpanishName represents masa spanish name.
	masaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// masaTraditionalChineseName represents masa traditional chinese name.
	masaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// masaName represents masa name.
	masaName = nook.Languages{
		language.AmericanEnglish:      masaAmericanEnglishName,
		language.CanadianFrench:       masaCanadianFrenchName,
		language.Dutch:                masaDutchName,
		language.French:               masaFrenchName,
		language.German:               masaGermanName,
		language.Italian:              masaItalianName,
		language.Japanese:             masaJapaneseName,
		language.Korean:               masaKoreanName,
		language.LatinAmericanSpanish: masaLatinAmericanSpanishName,
		language.Russian:              masaRussianName,
		language.SimplifiedChinese:    masaSimplifiedChineseName,
		language.Spanish:              masaSpanishName,
		language.TraditionalChinese:   masaTraditionalChineseName}
)

var (
	// masaCharacter represents masa character.
	masaCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: masaBirthday,
		Code:     masaCode,
		Key:      character.Masa,
		Gender:   gender.Male,
		Name:     masaName,
		Special:  false}
)

var (
	// masaAmericanEnglishPhrase represents masa american english phrase.
	masaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "てもんだ"}

	// masaCanadianFrenchPhrase represents masa canadian french phrase.
	masaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// masaDutchPhrase represents masa dutch phrase.
	masaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// masaFrenchPhrase represents masa french phrase.
	masaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// masaGermanPhrase represents masa german phrase.
	masaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// masaItalianPhrase represents masa italian phrase.
	masaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// masaJapanesePhrase represents masa japanese phrase.
	masaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "てもんだ"}

	// masaKoreanPhrase represents masa korean phrase.
	masaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// masaLatinAmericanSpanishPhrase represents masa latin american spanish phrase.
	masaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// masaRussianPhrase represents masa russian phrase.
	masaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// masaSimplifiedChinesePhrase represents masa simplified chinese phrase.
	masaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// masaSpanishPhrase represents masa spanish phrase.
	masaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// masaTraditionalChinesePhrase represents masa traditional chinese phrase.
	masaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// masaPhrase represents masa phrase.
	masaPhrase = nook.Languages{
		language.AmericanEnglish:      masaAmericanEnglishPhrase,
		language.CanadianFrench:       masaCanadianFrenchPhrase,
		language.Dutch:                masaDutchPhrase,
		language.French:               masaFrenchPhrase,
		language.German:               masaGermanPhrase,
		language.Italian:              masaItalianPhrase,
		language.Japanese:             masaJapanesePhrase,
		language.Korean:               masaKoreanPhrase,
		language.LatinAmericanSpanish: masaLatinAmericanSpanishPhrase,
		language.Russian:              masaRussianPhrase,
		language.SimplifiedChinese:    masaSimplifiedChinesePhrase,
		language.Spanish:              masaSpanishPhrase,
		language.TraditionalChinese:   masaTraditionalChinesePhrase}
)

var (
	// Masa represents masa.
	Masa = nook.Villager{
		Character:   masaCharacter,
		Personality: personality.Jock,
		Phrase:      masaPhrase}
)
