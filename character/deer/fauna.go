package deer

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
	faunaBirthday = nook.Birthday{
		Day:   26,
		Month: time.March}
)

var (
	faunaCode = nook.Code{
		Value: "der00"}
)

var (
	faunaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Fauna"}

	faunaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bibi"}

	faunaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Fauna"}

	faunaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bibi"}

	faunaGermanName = nook.Name{
		Language: language.German,
		Value:    "Fatima"}

	faunaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cervina"}

	faunaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドレミ"}

	faunaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "솔미"}

	faunaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fauna"}

	faunaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фауна"}

	faunaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "音音"}

	faunaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fauna"}

	faunaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "音音"}
)

var (
	faunaName = nook.Languages{
		language.AmericanEnglish:      faunaAmericanEnglishName,
		language.CanadianFrench:       faunaCanadianFrenchName,
		language.Dutch:                faunaDutchName,
		language.French:               faunaFrenchName,
		language.German:               faunaGermanName,
		language.Italian:              faunaItalianName,
		language.Japanese:             faunaJapaneseName,
		language.Korean:               faunaKoreanName,
		language.LatinAmericanSpanish: faunaLatinAmericanSpanishName,
		language.Russian:              faunaRussianName,
		language.SimplifiedChinese:    faunaSimplifiedChineseName,
		language.Spanish:              faunaSpanishName,
		language.TraditionalChinese:   faunaTraditionalChineseName}
)

var (
	faunaCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: faunaBirthday,
		Code:     faunaCode,
		Key:      character.Fauna,
		Gender:   gender.Female,
		Name:     faunaName,
		Special:  false}
)

var (
	faunaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "dearie"}

	faunaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "toudoux"}

	faunaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "do ree mi"}

	faunaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "toudoux"}

	faunaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bimbam"}

	faunaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "vule vu"}

	faunaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でしか"}

	faunaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "파샵파샵"}

	faunaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "puchu"}

	faunaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "олешек"}

	faunaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小鹿"}

	faunaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "puchu"}

	faunaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小鹿"}
)

var (
	faunaPhrase = nook.Languages{
		language.AmericanEnglish:      faunaAmericanEnglishPhrase,
		language.CanadianFrench:       faunaCanadianFrenchPhrase,
		language.Dutch:                faunaDutchPhrase,
		language.French:               faunaFrenchPhrase,
		language.German:               faunaGermanPhrase,
		language.Italian:              faunaItalianPhrase,
		language.Japanese:             faunaJapanesePhrase,
		language.Korean:               faunaKoreanPhrase,
		language.LatinAmericanSpanish: faunaLatinAmericanSpanishPhrase,
		language.Russian:              faunaRussianPhrase,
		language.SimplifiedChinese:    faunaSimplifiedChinesePhrase,
		language.Spanish:              faunaSpanishPhrase,
		language.TraditionalChinese:   faunaTraditionalChinesePhrase}
)

var (
	Fauna = nook.Villager{
		Character:   faunaCharacter,
		Personality: personality.Normal,
		Phrase:      faunaPhrase}
)
