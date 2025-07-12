package mouse

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
	// penelopeBirthday represents penelope birthday.
	penelopeBirthday = nook.Birthday{
		Day:   5,
		Month: time.February}
)

var (
	// penelopeCode represents penelope code.
	penelopeCode = nook.Code{
		Value: "mus17"}
)

var (
	// penelopeAmericanEnglishName represents penelope american english name.
	penelopeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Penelope"}

	// penelopeCanadianFrenchName represents penelope canadian french name.
	penelopeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Missy"}

	// penelopeDutchName represents penelope dutch name.
	penelopeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Penelope"}

	// penelopeFrenchName represents penelope french name.
	penelopeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Missy"}

	// penelopeGermanName represents penelope german name.
	penelopeGermanName = nook.Name{
		Language: language.German,
		Value:    "Penelope"}

	// penelopeItalianName represents penelope italian name.
	penelopeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Penelope"}

	// penelopeJapaneseName represents penelope japanese name.
	penelopeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チューこ"}

	// penelopeKoreanName represents penelope korean name.
	penelopeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "찍순이"}

	// penelopeLatinAmericanSpanishName represents penelope latin american spanish name.
	penelopeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Adela"}

	// penelopeRussianName represents penelope russian name.
	penelopeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Адела"}

	// penelopeSimplifiedChineseName represents penelope simplified chinese name.
	penelopeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小啾"}

	// penelopeSpanishName represents penelope spanish name.
	penelopeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Adela"}

	// penelopeTraditionalChineseName represents penelope traditional chinese name.
	penelopeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小啾"}
)

var (
	// penelopeName represents penelope name.
	penelopeName = nook.Languages{
		language.AmericanEnglish:      penelopeAmericanEnglishName,
		language.CanadianFrench:       penelopeCanadianFrenchName,
		language.Dutch:                penelopeDutchName,
		language.French:               penelopeFrenchName,
		language.German:               penelopeGermanName,
		language.Italian:              penelopeItalianName,
		language.Japanese:             penelopeJapaneseName,
		language.Korean:               penelopeKoreanName,
		language.LatinAmericanSpanish: penelopeLatinAmericanSpanishName,
		language.Russian:              penelopeRussianName,
		language.SimplifiedChinese:    penelopeSimplifiedChineseName,
		language.Spanish:              penelopeSpanishName,
		language.TraditionalChinese:   penelopeTraditionalChineseName}
)

var (
	// penelopeCharacter represents penelope character.
	penelopeCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: penelopeBirthday,
		Code:     penelopeCode,
		Key:      character.Penelope,
		Gender:   gender.Female,
		Name:     penelopeName,
		Special:  false}
)

var (
	// penelopeAmericanEnglishPhrase represents penelope american english phrase.
	penelopeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "oh bow"}

	// penelopeCanadianFrenchPhrase represents penelope canadian french phrase.
	penelopeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grugru"}

	// penelopeDutchPhrase represents penelope dutch phrase.
	penelopeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "strikje"}

	// penelopeFrenchPhrase represents penelope french phrase.
	penelopeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grugru"}

	// penelopeGermanPhrase represents penelope german phrase.
	penelopeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bling"}

	// penelopeItalianPhrase represents penelope italian phrase.
	penelopeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cheeese"}

	// penelopeJapanesePhrase represents penelope japanese phrase.
	penelopeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でちゅの"}

	// penelopeKoreanPhrase represents penelope korean phrase.
	penelopeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그러찍"}

	// penelopeLatinAmericanSpanishPhrase represents penelope latin american spanish phrase.
	penelopeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "brillilí"}

	// penelopeRussianPhrase represents penelope russian phrase.
	penelopeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бантик"}

	// penelopeSimplifiedChinesePhrase represents penelope simplified chinese phrase.
	penelopeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啾啾"}

	// penelopeSpanishPhrase represents penelope spanish phrase.
	penelopeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "migas"}

	// penelopeTraditionalChinesePhrase represents penelope traditional chinese phrase.
	penelopeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啾啾"}
)

var (
	// penelopePhrase represents penelope phrase.
	penelopePhrase = nook.Languages{
		language.AmericanEnglish:      penelopeAmericanEnglishPhrase,
		language.CanadianFrench:       penelopeCanadianFrenchPhrase,
		language.Dutch:                penelopeDutchPhrase,
		language.French:               penelopeFrenchPhrase,
		language.German:               penelopeGermanPhrase,
		language.Italian:              penelopeItalianPhrase,
		language.Japanese:             penelopeJapanesePhrase,
		language.Korean:               penelopeKoreanPhrase,
		language.LatinAmericanSpanish: penelopeLatinAmericanSpanishPhrase,
		language.Russian:              penelopeRussianPhrase,
		language.SimplifiedChinese:    penelopeSimplifiedChinesePhrase,
		language.Spanish:              penelopeSpanishPhrase,
		language.TraditionalChinese:   penelopeTraditionalChinesePhrase}
)

var (
	// Penelope represents penelope.
	Penelope = nook.Villager{
		Character:   penelopeCharacter,
		Personality: personality.Peppy,
		Phrase:      penelopePhrase}
)
