package kangaroo

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
	// valiseBirthday represents valise birthday.
	valiseBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// valiseCode represents valise code.
	valiseCode = nook.Code{
		Value: ""}
)

var (
	// valiseAmericanEnglishName represents valise american english name.
	valiseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Valise"}

	// valiseCanadianFrenchName represents valise canadian french name.
	valiseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// valiseDutchName represents valise dutch name.
	valiseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// valiseFrenchName represents valise french name.
	valiseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pochette"}

	// valiseGermanName represents valise german name.
	valiseGermanName = nook.Name{
		Language: language.German,
		Value:    "Elke"}

	// valiseItalianName represents valise italian name.
	valiseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pochette"}

	// valiseJapaneseName represents valise japanese name.
	valiseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エプロン"}

	// valiseKoreanName represents valise korean name.
	valiseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// valiseLatinAmericanSpanishName represents valise latin american spanish name.
	valiseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// valiseRussianName represents valise russian name.
	valiseRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// valiseSimplifiedChineseName represents valise simplified chinese name.
	valiseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "昭昭"}

	// valiseSpanishName represents valise spanish name.
	valiseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ana Rosa"}

	// valiseTraditionalChineseName represents valise traditional chinese name.
	valiseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// valiseName represents valise name.
	valiseName = nook.Languages{
		language.AmericanEnglish:      valiseAmericanEnglishName,
		language.CanadianFrench:       valiseCanadianFrenchName,
		language.Dutch:                valiseDutchName,
		language.French:               valiseFrenchName,
		language.German:               valiseGermanName,
		language.Italian:              valiseItalianName,
		language.Japanese:             valiseJapaneseName,
		language.Korean:               valiseKoreanName,
		language.LatinAmericanSpanish: valiseLatinAmericanSpanishName,
		language.Russian:              valiseRussianName,
		language.SimplifiedChinese:    valiseSimplifiedChineseName,
		language.Spanish:              valiseSpanishName,
		language.TraditionalChinese:   valiseTraditionalChineseName}
)

var (
	// valiseCharacter represents valise character.
	valiseCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: valiseBirthday,
		Code:     valiseCode,
		Key:      character.Valise,
		Gender:   gender.Female,
		Name:     valiseName,
		Special:  false}
)

var (
	// valiseAmericanEnglishPhrase represents valise american english phrase.
	valiseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tadder"}

	// valiseCanadianFrenchPhrase represents valise canadian french phrase.
	valiseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// valiseDutchPhrase represents valise dutch phrase.
	valiseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// valiseFrenchPhrase represents valise french phrase.
	valiseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "p'tit bout"}

	// valiseGermanPhrase represents valise german phrase.
	valiseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krabbel"}

	// valiseItalianPhrase represents valise italian phrase.
	valiseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bonjour"}

	// valiseJapanesePhrase represents valise japanese phrase.
	valiseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ピョン"}

	// valiseKoreanPhrase represents valise korean phrase.
	valiseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// valiseLatinAmericanSpanishPhrase represents valise latin american spanish phrase.
	valiseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// valiseRussianPhrase represents valise russian phrase.
	valiseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// valiseSimplifiedChinesePhrase represents valise simplified chinese phrase.
	valiseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噼哟"}

	// valiseSpanishPhrase represents valise spanish phrase.
	valiseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "personilla"}

	// valiseTraditionalChinesePhrase represents valise traditional chinese phrase.
	valiseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// valisePhrase represents valise phrase.
	valisePhrase = nook.Languages{
		language.AmericanEnglish:      valiseAmericanEnglishPhrase,
		language.CanadianFrench:       valiseCanadianFrenchPhrase,
		language.Dutch:                valiseDutchPhrase,
		language.French:               valiseFrenchPhrase,
		language.German:               valiseGermanPhrase,
		language.Italian:              valiseItalianPhrase,
		language.Japanese:             valiseJapanesePhrase,
		language.Korean:               valiseKoreanPhrase,
		language.LatinAmericanSpanish: valiseLatinAmericanSpanishPhrase,
		language.Russian:              valiseRussianPhrase,
		language.SimplifiedChinese:    valiseSimplifiedChinesePhrase,
		language.Spanish:              valiseSpanishPhrase,
		language.TraditionalChinese:   valiseTraditionalChinesePhrase}
)

var (
	// Valise represents valise.
	Valise = nook.Villager{
		Character:   valiseCharacter,
		Personality: personality.Snooty,
		Phrase:      valisePhrase}
)
