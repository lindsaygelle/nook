package deer

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Bibitoudoux"}

	faunaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Faunado ree mi"}

	faunaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bibitoudoux"}

	faunaGermanName = nook.Name{
		Language: language.German,
		Value:    "Fatimabimbam"}

	faunaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cervinavule vu"}

	faunaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドレミでしか"}

	faunaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "솔미파샵파샵"}

	faunaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Faunapuchu"}

	faunaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фаунаолешек"}

	faunaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "音音小鹿"}

	faunaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Faunapuchu"}

	faunaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "音音小鹿"}
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
		Animal:   Deer,
		Birthday: faunaBirthday,
		Code:     faunaCode,
		Gender:   nook.Female,
		Name:     faunaName}
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
	Fauna = nook.Villager{
		Character:   faunaCharacter,
		Personality: nook.Normal,
		Phrase:      faunaPhrase}
)
