package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	cousteauBirthday = nook.Birthday{
		Day:   17,
		Month: time.December}
)

var (
	cousteauCode = nook.Code{
		Value: "flg10"}
)

var (
	cousteauAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cousteau"}

	cousteauCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Figarobidule"}

	cousteauDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cousteauoui oui"}

	cousteauFrenchName = nook.Name{
		Language: language.French,
		Value:    "Figarobidule"}

	cousteauGermanName = nook.Name{
		Language: language.German,
		Value:    "Jacquesoui, oui"}

	cousteauItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jacqueset voilà"}

	cousteauJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハルマキのことよ"}

	cousteauKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "왕서방라이라이"}

	cousteauLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cousteauoui oui"}

	cousteauRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кустоуи-уи"}

	cousteauSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "春卷不客气"}

	cousteauSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cousteauoui oui"}

	cousteauTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "春卷不客氣"}
)

var (
	cousteauName = nook.Languages{
		language.AmericanEnglish:      cousteauAmericanEnglishName,
		language.CanadianFrench:       cousteauCanadianFrenchName,
		language.Dutch:                cousteauDutchName,
		language.French:               cousteauFrenchName,
		language.German:               cousteauGermanName,
		language.Italian:              cousteauItalianName,
		language.Japanese:             cousteauJapaneseName,
		language.Korean:               cousteauKoreanName,
		language.LatinAmericanSpanish: cousteauLatinAmericanSpanishName,
		language.Russian:              cousteauRussianName,
		language.SimplifiedChinese:    cousteauSimplifiedChineseName,
		language.Spanish:              cousteauSpanishName,
		language.TraditionalChinese:   cousteauTraditionalChineseName}
)

var (
	cousteauCharacter = nook.Character{
		Animal:   Frog,
		Birthday: cousteauBirthday,
		Code:     cousteauCode,
		Gender:   nook.Male,
		Name:     cousteauName}
)

var (
	cousteauAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "oui oui"}

	cousteauCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bidule"}

	cousteauDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oui oui"}

	cousteauFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bidule"}

	cousteauGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "oui, oui"}

	cousteauItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "et voilà"}

	cousteauJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "のことよ"}

	cousteauKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라이라이"}

	cousteauLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oui oui"}

	cousteauRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "уи-уи"}

	cousteauSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不客气"}

	cousteauSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oui oui"}

	cousteauTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不客氣"}
)

var (
	cousteauPhrase = nook.Languages{
		language.AmericanEnglish:      cousteauAmericanEnglishName,
		language.CanadianFrench:       cousteauCanadianFrenchName,
		language.Dutch:                cousteauDutchName,
		language.French:               cousteauFrenchName,
		language.German:               cousteauGermanName,
		language.Italian:              cousteauItalianName,
		language.Japanese:             cousteauJapaneseName,
		language.Korean:               cousteauKoreanName,
		language.LatinAmericanSpanish: cousteauLatinAmericanSpanishName,
		language.Russian:              cousteauRussianName,
		language.SimplifiedChinese:    cousteauSimplifiedChineseName,
		language.Spanish:              cousteauSpanishName,
		language.TraditionalChinese:   cousteauTraditionalChineseName}
)

var (
	Cousteau = nook.Villager{
		Character:   cousteauCharacter,
		Personality: nook.Jock,
		Phrase:      cousteauPhrase}
)
