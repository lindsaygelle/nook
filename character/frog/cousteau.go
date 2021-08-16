package frog

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
		Value:    "Figaro"}

	cousteauDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cousteau"}

	cousteauFrenchName = nook.Name{
		Language: language.French,
		Value:    "Figaro"}

	cousteauGermanName = nook.Name{
		Language: language.German,
		Value:    "Jacques"}

	cousteauItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jacques"}

	cousteauJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハルマキ"}

	cousteauKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "왕서방"}

	cousteauLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cousteau"}

	cousteauRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кусто"}

	cousteauSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "春卷"}

	cousteauSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cousteau"}

	cousteauTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "春卷"}
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
		Animal:   animal.Frog,
		Birthday: cousteauBirthday,
		Code:     cousteauCode,
		Key:      character.Cousteau,
		Gender:   gender.Male,
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
		Personality: personality.Jock,
		Phrase:      cousteauPhrase}
)
