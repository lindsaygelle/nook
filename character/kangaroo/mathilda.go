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
	mathildaBirthday = nook.Birthday{
		Day:   12,
		Month: time.November}
)

var (
	mathildaCode = nook.Code{
		Value: "kgr01"}
)

var (
	mathildaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mathilda"}

	mathildaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mathilde"}

	mathildaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mathilda"}

	mathildaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mathilde"}

	mathildaGermanName = nook.Name{
		Language: language.German,
		Value:    "Marga"}

	mathildaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Matilda"}

	mathildaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アザラク"}

	mathildaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안젤라"}

	mathildaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pugilda"}

	mathildaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Матильда"}

	mathildaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亚莎"}

	mathildaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pugilda"}

	mathildaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亞莎"}
)

var (
	mathildaName = nook.Languages{
		language.AmericanEnglish:      mathildaAmericanEnglishName,
		language.CanadianFrench:       mathildaCanadianFrenchName,
		language.Dutch:                mathildaDutchName,
		language.French:               mathildaFrenchName,
		language.German:               mathildaGermanName,
		language.Italian:              mathildaItalianName,
		language.Japanese:             mathildaJapaneseName,
		language.Korean:               mathildaKoreanName,
		language.LatinAmericanSpanish: mathildaLatinAmericanSpanishName,
		language.Russian:              mathildaRussianName,
		language.SimplifiedChinese:    mathildaSimplifiedChineseName,
		language.Spanish:              mathildaSpanishName,
		language.TraditionalChinese:   mathildaTraditionalChineseName}
)

var (
	mathildaCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: mathildaBirthday,
		Code:     mathildaCode,
		Key:      character.Mathilda,
		Gender:   gender.Female,
		Name:     mathildaName}
)

var (
	mathildaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wee baby"}

	mathildaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "riquiqui"}

	mathildaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ukkepuk"}

	mathildaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "riquiqui"}

	mathildaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knirps"}

	mathildaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oplà"}

	mathildaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ッハ"}

	mathildaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "응"}

	mathildaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ala-rorro"}

	mathildaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "малютка"}

	mathildaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘿"}

	mathildaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "peso mosca"}

	mathildaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘿"}
)

var (
	mathildaPhrase = nook.Languages{
		language.AmericanEnglish:      mathildaAmericanEnglishPhrase,
		language.CanadianFrench:       mathildaCanadianFrenchPhrase,
		language.Dutch:                mathildaDutchPhrase,
		language.French:               mathildaFrenchPhrase,
		language.German:               mathildaGermanPhrase,
		language.Italian:              mathildaItalianPhrase,
		language.Japanese:             mathildaJapanesePhrase,
		language.Korean:               mathildaKoreanPhrase,
		language.LatinAmericanSpanish: mathildaLatinAmericanSpanishPhrase,
		language.Russian:              mathildaRussianPhrase,
		language.SimplifiedChinese:    mathildaSimplifiedChinesePhrase,
		language.Spanish:              mathildaSpanishPhrase,
		language.TraditionalChinese:   mathildaTraditionalChinesePhrase}
)

var (
	Mathilda = nook.Villager{
		Character:   mathildaCharacter,
		Personality: personality.Snooty,
		Phrase:      mathildaPhrase}
)
