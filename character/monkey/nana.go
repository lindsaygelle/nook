package monkey

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
	// nanaBirthday represents nana birthday.
	nanaBirthday = nook.Birthday{
		Day:   23,
		Month: time.August}
)

var (
	// nanaCode represents nana code.
	nanaCode = nook.Code{
		Value: "mnk01"}
)

var (
	// nanaAmericanEnglishName represents nana american english name.
	nanaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nana"}

	// nanaCanadianFrenchName represents nana canadian french name.
	nanaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mireille"}

	// nanaDutchName represents nana dutch name.
	nanaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nana"}

	// nanaFrenchName represents nana french name.
	nanaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mireille"}

	// nanaGermanName represents nana german name.
	nanaGermanName = nook.Name{
		Language: language.German,
		Value:    "Dorothea"}

	// nanaItalianName represents nana italian name.
	nanaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nanà"}

	// nanaJapaneseName represents nana japanese name.
	nanaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チッチ"}

	// nanaKoreanName represents nana korean name.
	nanaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "키키"}

	// nanaLatinAmericanSpanishName represents nana latin american spanish name.
	nanaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nana"}

	// nanaRussianName represents nana russian name.
	nanaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нана"}

	// nanaSimplifiedChineseName represents nana simplified chinese name.
	nanaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "七七"}

	// nanaSpanishName represents nana spanish name.
	nanaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nana"}

	// nanaTraditionalChineseName represents nana traditional chinese name.
	nanaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "七七"}
)

var (
	// nanaName represents nana name.
	nanaName = nook.Languages{
		language.AmericanEnglish:      nanaAmericanEnglishName,
		language.CanadianFrench:       nanaCanadianFrenchName,
		language.Dutch:                nanaDutchName,
		language.French:               nanaFrenchName,
		language.German:               nanaGermanName,
		language.Italian:              nanaItalianName,
		language.Japanese:             nanaJapaneseName,
		language.Korean:               nanaKoreanName,
		language.LatinAmericanSpanish: nanaLatinAmericanSpanishName,
		language.Russian:              nanaRussianName,
		language.SimplifiedChinese:    nanaSimplifiedChineseName,
		language.Spanish:              nanaSpanishName,
		language.TraditionalChinese:   nanaTraditionalChineseName}
)

var (
	// nanaCharacter represents nana character.
	nanaCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: nanaBirthday,
		Code:     nanaCode,
		Key:      character.Nana,
		Gender:   gender.Female,
		Name:     nanaName,
		Special:  false}
)

var (
	// nanaAmericanEnglishPhrase represents nana american english phrase.
	nanaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "po po"}

	// nanaCanadianFrenchPhrase represents nana canadian french phrase.
	nanaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "po po"}

	// nanaDutchPhrase represents nana dutch phrase.
	nanaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "banaan"}

	// nanaFrenchPhrase represents nana french phrase.
	nanaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "po po"}

	// nanaGermanPhrase represents nana german phrase.
	nanaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "spitzi"}

	// nanaItalianPhrase represents nana italian phrase.
	nanaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "po po"}

	// nanaJapanesePhrase represents nana japanese phrase.
	nanaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウキャ"}

	// nanaKoreanPhrase represents nana korean phrase.
	nanaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "끼끼"}

	// nanaLatinAmericanSpanishPhrase represents nana latin american spanish phrase.
	nanaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ukiú"}

	// nanaRussianPhrase represents nana russian phrase.
	nanaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пу-пу"}

	// nanaSimplifiedChinesePhrase represents nana simplified chinese phrase.
	nanaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唔唔"}

	// nanaSpanishPhrase represents nana spanish phrase.
	nanaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "monada"}

	// nanaTraditionalChinesePhrase represents nana traditional chinese phrase.
	nanaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唔唔"}
)

var (
	// nanaPhrase represents nana phrase.
	nanaPhrase = nook.Languages{
		language.AmericanEnglish:      nanaAmericanEnglishPhrase,
		language.CanadianFrench:       nanaCanadianFrenchPhrase,
		language.Dutch:                nanaDutchPhrase,
		language.French:               nanaFrenchPhrase,
		language.German:               nanaGermanPhrase,
		language.Italian:              nanaItalianPhrase,
		language.Japanese:             nanaJapanesePhrase,
		language.Korean:               nanaKoreanPhrase,
		language.LatinAmericanSpanish: nanaLatinAmericanSpanishPhrase,
		language.Russian:              nanaRussianPhrase,
		language.SimplifiedChinese:    nanaSimplifiedChinesePhrase,
		language.Spanish:              nanaSpanishPhrase,
		language.TraditionalChinese:   nanaTraditionalChinesePhrase}
)

var (
	// Nana represents nana.
	Nana = nook.Villager{
		Character:   nanaCharacter,
		Personality: personality.Normal,
		Phrase:      nanaPhrase}
)
