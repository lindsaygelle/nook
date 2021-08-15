package kangaroo

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Mathilderiquiqui"}

	mathildaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mathildaukkepuk"}

	mathildaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mathilderiquiqui"}

	mathildaGermanName = nook.Name{
		Language: language.German,
		Value:    "Margaknirps"}

	mathildaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Matildaoplà"}

	mathildaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アザラクッハ"}

	mathildaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안젤라응"}

	mathildaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pugildaala-rorro"}

	mathildaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Матильдамалютка"}

	mathildaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亚莎嘿"}

	mathildaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pugildapeso mosca"}

	mathildaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亞莎嘿"}
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
		Animal:   Kangaroo,
		Birthday: mathildaBirthday,
		Code:     mathildaCode,
		Gender:   nook.Female,
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
	Mathilda = nook.Villager{
		Character:   mathildaCharacter,
		Personality: nook.Snooty,
		Phrase:      mathildaPhrase}
)
