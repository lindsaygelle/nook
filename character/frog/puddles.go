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
	// puddlesBirthday represents puddles birthday.
	puddlesBirthday = nook.Birthday{
		Day:   13,
		Month: time.January}
)

var (
	// puddlesCode represents puddles code.
	puddlesCode = nook.Code{
		Value: "flg06"}
)

var (
	// puddlesAmericanEnglishName represents puddles american english name.
	puddlesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Puddles"}

	// puddlesCanadianFrenchName represents puddles canadian french name.
	puddlesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rénata"}

	// puddlesDutchName represents puddles dutch name.
	puddlesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Puddles"}

	// puddlesFrenchName represents puddles french name.
	puddlesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rénata"}

	// puddlesGermanName represents puddles german name.
	puddlesGermanName = nook.Name{
		Language: language.German,
		Value:    "Nele"}

	// puddlesItalianName represents puddles italian name.
	puddlesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grazia"}

	// puddlesJapaneseName represents puddles japanese name.
	puddlesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チョキ"}

	// puddlesKoreanName represents puddles korean name.
	puddlesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "가위"}

	// puddlesLatinAmericanSpanishName represents puddles latin american spanish name.
	puddlesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nenúfar"}

	// puddlesRussianName represents puddles russian name.
	puddlesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рената"}

	// puddlesSimplifiedChineseName represents puddles simplified chinese name.
	puddlesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "剪刀"}

	// puddlesSpanishName represents puddles spanish name.
	puddlesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nenúfar"}

	// puddlesTraditionalChineseName represents puddles traditional chinese name.
	puddlesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "剪刀"}
)

var (
	// puddlesName represents puddles name.
	puddlesName = nook.Languages{
		language.AmericanEnglish:      puddlesAmericanEnglishName,
		language.CanadianFrench:       puddlesCanadianFrenchName,
		language.Dutch:                puddlesDutchName,
		language.French:               puddlesFrenchName,
		language.German:               puddlesGermanName,
		language.Italian:              puddlesItalianName,
		language.Japanese:             puddlesJapaneseName,
		language.Korean:               puddlesKoreanName,
		language.LatinAmericanSpanish: puddlesLatinAmericanSpanishName,
		language.Russian:              puddlesRussianName,
		language.SimplifiedChinese:    puddlesSimplifiedChineseName,
		language.Spanish:              puddlesSpanishName,
		language.TraditionalChinese:   puddlesTraditionalChineseName}
)

var (
	// puddlesCharacter represents puddles character.
	puddlesCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: puddlesBirthday,
		Code:     puddlesCode,
		Key:      character.Puddles,
		Gender:   gender.Female,
		Name:     puddlesName,
		Special:  false}
)

var (
	// puddlesAmericanEnglishPhrase represents puddles american english phrase.
	puddlesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "splish"}

	// puddlesCanadianFrenchPhrase represents puddles canadian french phrase.
	puddlesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "splufff"}

	// puddlesDutchPhrase represents puddles dutch phrase.
	puddlesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "plons"}

	// puddlesFrenchPhrase represents puddles french phrase.
	puddlesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "splufff"}

	// puddlesGermanPhrase represents puddles german phrase.
	puddlesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "platschi"}

	// puddlesItalianPhrase represents puddles italian phrase.
	puddlesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "splash"}

	// puddlesJapanesePhrase represents puddles japanese phrase.
	puddlesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "っちゃ"}

	// puddlesKoreanPhrase represents puddles korean phrase.
	puddlesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그랬쪄"}

	// puddlesLatinAmericanSpanishPhrase represents puddles latin american spanish phrase.
	puddlesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "esplís"}

	// puddlesRussianPhrase represents puddles russian phrase.
	puddlesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "плюх"}

	// puddlesSimplifiedChinesePhrase represents puddles simplified chinese phrase.
	puddlesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "石头"}

	// puddlesSpanishPhrase represents puddles spanish phrase.
	puddlesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "esplís"}

	// puddlesTraditionalChinesePhrase represents puddles traditional chinese phrase.
	puddlesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "石頭"}
)

var (
	// puddlesPhrase represents puddles phrase.
	puddlesPhrase = nook.Languages{
		language.AmericanEnglish:      puddlesAmericanEnglishPhrase,
		language.CanadianFrench:       puddlesCanadianFrenchPhrase,
		language.Dutch:                puddlesDutchPhrase,
		language.French:               puddlesFrenchPhrase,
		language.German:               puddlesGermanPhrase,
		language.Italian:              puddlesItalianPhrase,
		language.Japanese:             puddlesJapanesePhrase,
		language.Korean:               puddlesKoreanPhrase,
		language.LatinAmericanSpanish: puddlesLatinAmericanSpanishPhrase,
		language.Russian:              puddlesRussianPhrase,
		language.SimplifiedChinese:    puddlesSimplifiedChinesePhrase,
		language.Spanish:              puddlesSpanishPhrase,
		language.TraditionalChinese:   puddlesTraditionalChinesePhrase}
)

var (
	// Puddles represents puddles.
	Puddles = nook.Villager{
		Character:   puddlesCharacter,
		Personality: personality.Peppy,
		Phrase:      puddlesPhrase}
)
