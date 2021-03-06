package cat

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
	kattBirthday = nook.Birthday{
		Day:   27,
		Month: time.April}
)

var (
	kattCode = nook.Code{
		Value: "cat21"}
)

var (
	kattAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Katt"}

	kattCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kat"}

	kattDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Katt"}

	kattFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kat"}

	kattGermanName = nook.Name{
		Language: language.German,
		Value:    "Janine"}

	kattItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ofelia"}

	kattJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちょい"}

	kattKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쵸이"}

	kattLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Melina"}

	kattRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэт"}

	kattSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巧巧"}

	kattSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Melina"}

	kattTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巧巧"}
)

var (
	kattName = nook.Languages{
		language.AmericanEnglish:      kattAmericanEnglishName,
		language.CanadianFrench:       kattCanadianFrenchName,
		language.Dutch:                kattDutchName,
		language.French:               kattFrenchName,
		language.German:               kattGermanName,
		language.Italian:              kattItalianName,
		language.Japanese:             kattJapaneseName,
		language.Korean:               kattKoreanName,
		language.LatinAmericanSpanish: kattLatinAmericanSpanishName,
		language.Russian:              kattRussianName,
		language.SimplifiedChinese:    kattSimplifiedChineseName,
		language.Spanish:              kattSpanishName,
		language.TraditionalChinese:   kattTraditionalChineseName}
)

var (
	kattCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: kattBirthday,
		Code:     kattCode,
		Key:      character.Katt,
		Gender:   gender.Female,
		Name:     kattName,
		Special:  false}
)

var (
	kattAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "purrty"}

	kattCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miaouille"}

	kattDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "prrrachtig"}

	kattFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miaouille"}

	kattGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zischel"}

	kattItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fffz"}

	kattJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃね"}

	kattKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "휘리릭"}

	kattLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "purrrs"}

	kattRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хор-р-рошо"}

	kattSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "掰掰"}

	kattSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "salmonete"}

	kattTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "掰掰"}
)

var (
	kattPhrase = nook.Languages{
		language.AmericanEnglish:      kattAmericanEnglishPhrase,
		language.CanadianFrench:       kattCanadianFrenchPhrase,
		language.Dutch:                kattDutchPhrase,
		language.French:               kattFrenchPhrase,
		language.German:               kattGermanPhrase,
		language.Italian:              kattItalianPhrase,
		language.Japanese:             kattJapanesePhrase,
		language.Korean:               kattKoreanPhrase,
		language.LatinAmericanSpanish: kattLatinAmericanSpanishPhrase,
		language.Russian:              kattRussianPhrase,
		language.SimplifiedChinese:    kattSimplifiedChinesePhrase,
		language.Spanish:              kattSpanishPhrase,
		language.TraditionalChinese:   kattTraditionalChinesePhrase}
)

var (
	Katt = nook.Villager{
		Character:   kattCharacter,
		Personality: personality.BigSister,
		Phrase:      kattPhrase}
)
