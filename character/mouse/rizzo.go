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
	rizzoBirthday = nook.Birthday{
		Day:   17,
		Month: time.January}
)

var (
	rizzoCode = nook.Code{
		Value: "mus09"}
)

var (
	rizzoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rizzo"}

	rizzoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sourizzi"}

	rizzoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rizzo"}

	rizzoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sourizzi"}

	rizzoGermanName = nook.Name{
		Language: language.German,
		Value:    "Ricky"}

	rizzoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rizzo"}

	rizzoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちょろきち"}

	rizzoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "조르쥐"}

	rizzoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ratolón"}

	rizzoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Риццо"}

	rizzoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "悄俊"}

	rizzoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ratolón"}

	rizzoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "悄俊"}
)

var (
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
	rizzoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "squee"}

	rizzoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nouik"}

	rizzoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "piephoi"}

	rizzoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nouik"}

	rizzoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "fieps"}

	rizzoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squiii"}

	rizzoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "がってん"}

	rizzoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "힐끔힐끔"}

	rizzoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñiiiik"}

	rizzoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пи-и-и"}

	rizzoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蹑蹑"}

	rizzoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñiiiik"}

	rizzoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "躡躡"}
)

var (
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
	Rizzo = nook.Villager{
		Character:   rizzoCharacter,
		Personality: personality.Cranky,
		Phrase:      rizzoPhrase}
)
