package ostrich

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
	cranstonBirthday = nook.Birthday{
		Day:   23,
		Month: time.September}
)

var (
	cranstonCode = nook.Code{
		Value: "ost06"}
)

var (
	cranstonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cranston"}

	cranstonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gabin"}

	cranstonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cranston"}

	cranstonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gabin"}

	cranstonGermanName = nook.Name{
		Language: language.German,
		Value:    "Guido"}

	cranstonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carmine"}

	cranstonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トキオ"}

	cranstonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "타키"}

	cranstonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Carmelo"}

	cranstonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Крэнстон"}

	cranstonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱禄"}

	cranstonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carmelo"}

	cranstonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱祿"}
)

var (
	cranstonName = nook.Languages{
		language.AmericanEnglish:      cranstonAmericanEnglishName,
		language.CanadianFrench:       cranstonCanadianFrenchName,
		language.Dutch:                cranstonDutchName,
		language.French:               cranstonFrenchName,
		language.German:               cranstonGermanName,
		language.Italian:              cranstonItalianName,
		language.Japanese:             cranstonJapaneseName,
		language.Korean:               cranstonKoreanName,
		language.LatinAmericanSpanish: cranstonLatinAmericanSpanishName,
		language.Russian:              cranstonRussianName,
		language.SimplifiedChinese:    cranstonSimplifiedChineseName,
		language.Spanish:              cranstonSpanishName,
		language.TraditionalChinese:   cranstonTraditionalChineseName}
)

var (
	cranstonCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: cranstonBirthday,
		Code:     cranstonCode,
		Key:      character.Cranston,
		Gender:   gender.Male,
		Name:     cranstonName}
)

var (
	cranstonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sweatband"}

	cranstonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "omelette"}

	cranstonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zweetband"}

	cranstonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "omelette"}

	cranstonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "happahappa"}

	cranstonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "boa"}

	cranstonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "およよ"}

	cranstonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오요용"}

	cranstonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gruqui"}

	cranstonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "не потей"}

	cranstonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哎呀呀"}

	cranstonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gruqui"}

	cranstonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哎呀呀"}
)

var (
	cranstonPhrase = nook.Languages{
		language.AmericanEnglish:      cranstonAmericanEnglishPhrase,
		language.CanadianFrench:       cranstonCanadianFrenchPhrase,
		language.Dutch:                cranstonDutchPhrase,
		language.French:               cranstonFrenchPhrase,
		language.German:               cranstonGermanPhrase,
		language.Italian:              cranstonItalianPhrase,
		language.Japanese:             cranstonJapanesePhrase,
		language.Korean:               cranstonKoreanPhrase,
		language.LatinAmericanSpanish: cranstonLatinAmericanSpanishPhrase,
		language.Russian:              cranstonRussianPhrase,
		language.SimplifiedChinese:    cranstonSimplifiedChinesePhrase,
		language.Spanish:              cranstonSpanishPhrase,
		language.TraditionalChinese:   cranstonTraditionalChinesePhrase}
)

var (
	Cranston = nook.Villager{
		Character:   cranstonCharacter,
		Personality: personality.Lazy,
		Phrase:      cranstonPhrase}
)
