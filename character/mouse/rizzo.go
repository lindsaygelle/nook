package mouse

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
	// rizzoBirthday represents rizzo birthday.
	rizzoBirthday = nook.Birthday{
		Day:   17,
		Month: time.January}
)

var (
	// rizzoCode represents rizzo code.
	rizzoCode = nook.Code{
		Value: "mus09"}
)

var (
	// rizzoAmericanEnglishName represents rizzo american english name.
	rizzoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rizzo"}

	// rizzoCanadianFrenchName represents rizzo canadian french name.
	rizzoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sourizzi"}

	// rizzoDutchName represents rizzo dutch name.
	rizzoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rizzo"}

	// rizzoFrenchName represents rizzo french name.
	rizzoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sourizzi"}

	// rizzoGermanName represents rizzo german name.
	rizzoGermanName = nook.Name{
		Language: language.German,
		Value:    "Ricky"}

	// rizzoItalianName represents rizzo italian name.
	rizzoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rizzo"}

	// rizzoJapaneseName represents rizzo japanese name.
	rizzoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちょろきち"}

	// rizzoKoreanName represents rizzo korean name.
	rizzoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "조르쥐"}

	// rizzoLatinAmericanSpanishName represents rizzo latin american spanish name.
	rizzoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ratolón"}

	// rizzoRussianName represents rizzo russian name.
	rizzoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Риццо"}

	// rizzoSimplifiedChineseName represents rizzo simplified chinese name.
	rizzoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "悄俊"}

	// rizzoSpanishName represents rizzo spanish name.
	rizzoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ratolón"}

	// rizzoTraditionalChineseName represents rizzo traditional chinese name.
	rizzoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "悄俊"}
)

var (
	// rizzoName represents rizzo name.
	rizzoName = nook.Languages{
		language.AmericanEnglish:      rizzoAmericanEnglishName,
		language.CanadianFrench:       rizzoCanadianFrenchName,
		language.Dutch:                rizzoDutchName,
		language.French:               rizzoFrenchName,
		language.German:               rizzoGermanName,
		language.Italian:              rizzoItalianName,
		language.Japanese:             rizzoJapaneseName,
		language.Korean:               rizzoKoreanName,
		language.LatinAmericanSpanish: rizzoLatinAmericanSpanishName,
		language.Russian:              rizzoRussianName,
		language.SimplifiedChinese:    rizzoSimplifiedChineseName,
		language.Spanish:              rizzoSpanishName,
		language.TraditionalChinese:   rizzoTraditionalChineseName}
)

var (
	// rizzoCharacter represents rizzo character.
	rizzoCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: rizzoBirthday,
		Code:     rizzoCode,
		Key:      character.Rizzo,
		Gender:   gender.Male,
		Name:     rizzoName,
		Special:  false}
)

var (
	// rizzoAmericanEnglishPhrase represents rizzo american english phrase.
	rizzoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "squee"}

	// rizzoCanadianFrenchPhrase represents rizzo canadian french phrase.
	rizzoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nouik"}

	// rizzoDutchPhrase represents rizzo dutch phrase.
	rizzoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "piephoi"}

	// rizzoFrenchPhrase represents rizzo french phrase.
	rizzoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nouik"}

	// rizzoGermanPhrase represents rizzo german phrase.
	rizzoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "fieps"}

	// rizzoItalianPhrase represents rizzo italian phrase.
	rizzoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squiii"}

	// rizzoJapanesePhrase represents rizzo japanese phrase.
	rizzoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "がってん"}

	// rizzoKoreanPhrase represents rizzo korean phrase.
	rizzoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "힐끔힐끔"}

	// rizzoLatinAmericanSpanishPhrase represents rizzo latin american spanish phrase.
	rizzoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñiiiik"}

	// rizzoRussianPhrase represents rizzo russian phrase.
	rizzoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пи-и-и"}

	// rizzoSimplifiedChinesePhrase represents rizzo simplified chinese phrase.
	rizzoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蹑蹑"}

	// rizzoSpanishPhrase represents rizzo spanish phrase.
	rizzoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñiiiik"}

	// rizzoTraditionalChinesePhrase represents rizzo traditional chinese phrase.
	rizzoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "躡躡"}
)

var (
	// rizzoPhrase represents rizzo phrase.
	rizzoPhrase = nook.Languages{
		language.AmericanEnglish:      rizzoAmericanEnglishPhrase,
		language.CanadianFrench:       rizzoCanadianFrenchPhrase,
		language.Dutch:                rizzoDutchPhrase,
		language.French:               rizzoFrenchPhrase,
		language.German:               rizzoGermanPhrase,
		language.Italian:              rizzoItalianPhrase,
		language.Japanese:             rizzoJapanesePhrase,
		language.Korean:               rizzoKoreanPhrase,
		language.LatinAmericanSpanish: rizzoLatinAmericanSpanishPhrase,
		language.Russian:              rizzoRussianPhrase,
		language.SimplifiedChinese:    rizzoSimplifiedChinesePhrase,
		language.Spanish:              rizzoSpanishPhrase,
		language.TraditionalChinese:   rizzoTraditionalChinesePhrase}
)

var (
	// Rizzo represents rizzo.
	Rizzo = nook.Villager{
		Character:   rizzoCharacter,
		Personality: personality.Cranky,
		Phrase:      rizzoPhrase}
)
