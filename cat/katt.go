package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Katmiaouille"}

	kattDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kattprrrachtig"}

	kattFrenchName = nook.Name{
		Language: language.French,
		Value:    "Katmiaouille"}

	kattGermanName = nook.Name{
		Language: language.German,
		Value:    "Janinezischel"}

	kattItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ofeliafffz"}

	kattJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちょいじゃね"}

	kattKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쵸이휘리릭"}

	kattLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Melinapurrrs"}

	kattRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэтхор-р-рошо"}

	kattSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巧巧掰掰"}

	kattSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Melinasalmonete"}

	kattTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巧巧掰掰"}
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
		Animal:   Cat,
		Birthday: kattBirthday,
		Code:     kattCode,
		Gender:   nook.Female,
		Name:     kattName}
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
	Katt = nook.Villager{
		Character:   kattCharacter,
		Personality: nook.BigSister,
		Phrase:      kattPhrase}
)
