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
	// faunaBirthday represents fauna birthday.
	faunaBirthday = nook.Birthday{
		Day:   26,
		Month: time.March}
)

var (
	// faunaCode represents fauna code.
	faunaCode = nook.Code{
		Value: "der00"}
)

var (
	// faunaAmericanEnglishName represents fauna american english name.
	faunaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Fauna"}

	// faunaCanadianFrenchName represents fauna canadian french name.
	faunaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bibi"}

	// faunaDutchName represents fauna dutch name.
	faunaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Fauna"}

	// faunaFrenchName represents fauna french name.
	faunaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bibi"}

	// faunaGermanName represents fauna german name.
	faunaGermanName = nook.Name{
		Language: language.German,
		Value:    "Fatima"}

	// faunaItalianName represents fauna italian name.
	faunaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cervina"}

	// faunaJapaneseName represents fauna japanese name.
	faunaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドレミ"}

	// faunaKoreanName represents fauna korean name.
	faunaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "솔미"}

	// faunaLatinAmericanSpanishName represents fauna latin american spanish name.
	faunaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fauna"}

	// faunaRussianName represents fauna russian name.
	faunaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фауна"}

	// faunaSimplifiedChineseName represents fauna simplified chinese name.
	faunaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "音音"}

	// faunaSpanishName represents fauna spanish name.
	faunaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fauna"}

	// faunaTraditionalChineseName represents fauna traditional chinese name.
	faunaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "音音"}
)

var (
	// faunaName represents fauna name.
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
	// faunaCharacter represents fauna character.
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
	// faunaAmericanEnglishPhrase represents fauna american english phrase.
	faunaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "dearie"}

	// faunaCanadianFrenchPhrase represents fauna canadian french phrase.
	faunaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "toudoux"}

	// faunaDutchPhrase represents fauna dutch phrase.
	faunaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "do ree mi"}

	// faunaFrenchPhrase represents fauna french phrase.
	faunaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "toudoux"}

	// faunaGermanPhrase represents fauna german phrase.
	faunaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bimbam"}

	// faunaItalianPhrase represents fauna italian phrase.
	faunaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "vule vu"}

	// faunaJapanesePhrase represents fauna japanese phrase.
	faunaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でしか"}

	// faunaKoreanPhrase represents fauna korean phrase.
	faunaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "파샵파샵"}

	// faunaLatinAmericanSpanishPhrase represents fauna latin american spanish phrase.
	faunaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "puchu"}

	// faunaRussianPhrase represents fauna russian phrase.
	faunaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "олешек"}

	// faunaSimplifiedChinesePhrase represents fauna simplified chinese phrase.
	faunaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小鹿"}

	// faunaSpanishPhrase represents fauna spanish phrase.
	faunaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "puchu"}

	// faunaTraditionalChinesePhrase represents fauna traditional chinese phrase.
	faunaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小鹿"}
)

var (
	// faunaPhrase represents fauna phrase.
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
	// Fauna represents fauna.
	Fauna = nook.Villager{
		Character:   faunaCharacter,
		Personality: personality.Normal,
		Phrase:      faunaPhrase}
)
