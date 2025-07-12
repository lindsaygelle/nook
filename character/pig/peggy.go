package pig

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
	// peggyBirthday represents peggy birthday.
	peggyBirthday = nook.Birthday{
		Day:   23,
		Month: time.May}
)

var (
	// peggyCode represents peggy code.
	peggyCode = nook.Code{
		Value: "pig11"}
)

var (
	// peggyAmericanEnglishName represents peggy american english name.
	peggyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peggy"}

	// peggyCanadianFrenchName represents peggy canadian french name.
	peggyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rose"}

	// peggyDutchName represents peggy dutch name.
	peggyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peggy"}

	// peggyFrenchName represents peggy french name.
	peggyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rose"}

	// peggyGermanName represents peggy german name.
	peggyGermanName = nook.Name{
		Language: language.German,
		Value:    "Quiekie"}

	// peggyItalianName represents peggy italian name.
	peggyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sally"}

	// peggyJapaneseName represents peggy japanese name.
	peggyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちえり"}

	// peggyKoreanName represents peggy korean name.
	peggyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "체리"}

	// peggyLatinAmericanSpanishName represents peggy latin american spanish name.
	peggyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Peggy"}

	// peggyRussianName represents peggy russian name.
	peggyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пегги"}

	// peggySimplifiedChineseName represents peggy simplified chinese name.
	peggySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "千惠"}

	// peggySpanishName represents peggy spanish name.
	peggySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Peggy"}

	// peggyTraditionalChineseName represents peggy traditional chinese name.
	peggyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "千惠"}
)

var (
	// peggyName represents peggy name.
	peggyName = nook.Languages{
		language.AmericanEnglish:      peggyAmericanEnglishName,
		language.CanadianFrench:       peggyCanadianFrenchName,
		language.Dutch:                peggyDutchName,
		language.French:               peggyFrenchName,
		language.German:               peggyGermanName,
		language.Italian:              peggyItalianName,
		language.Japanese:             peggyJapaneseName,
		language.Korean:               peggyKoreanName,
		language.LatinAmericanSpanish: peggyLatinAmericanSpanishName,
		language.Russian:              peggyRussianName,
		language.SimplifiedChinese:    peggySimplifiedChineseName,
		language.Spanish:              peggySpanishName,
		language.TraditionalChinese:   peggyTraditionalChineseName}
)

var (
	// peggyCharacter represents peggy character.
	peggyCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: peggyBirthday,
		Code:     peggyCode,
		Key:      character.Peggy,
		Gender:   gender.Female,
		Name:     peggyName,
		Special:  false}
)

var (
	// peggyAmericanEnglishPhrase represents peggy american english phrase.
	peggyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "shweetie"}

	// peggyCanadianFrenchPhrase represents peggy canadian french phrase.
	peggyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "fleur"}

	// peggyDutchPhrase represents peggy dutch phrase.
	peggyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knorrelief"}

	// peggyFrenchPhrase represents peggy french phrase.
	peggyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "fleur"}

	// peggyGermanPhrase represents peggy german phrase.
	peggyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "glotz"}

	// peggyItalianPhrase represents peggy italian phrase.
	peggyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oink"}

	// peggyJapanesePhrase represents peggy japanese phrase.
	peggyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぷるる"}

	// peggyKoreanPhrase represents peggy korean phrase.
	peggyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아앙"}

	// peggyLatinAmericanSpanishPhrase represents peggy latin american spanish phrase.
	peggyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "trufits"}

	// peggyRussianPhrase represents peggy russian phrase.
	peggyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюмпатяга"}

	// peggySimplifiedChinesePhrase represents peggy simplified chinese phrase.
	peggySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "弹弹"}

	// peggySpanishPhrase represents peggy spanish phrase.
	peggySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trufita"}

	// peggyTraditionalChinesePhrase represents peggy traditional chinese phrase.
	peggyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "彈彈"}
)

var (
	// peggyPhrase represents peggy phrase.
	peggyPhrase = nook.Languages{
		language.AmericanEnglish:      peggyAmericanEnglishPhrase,
		language.CanadianFrench:       peggyCanadianFrenchPhrase,
		language.Dutch:                peggyDutchPhrase,
		language.French:               peggyFrenchPhrase,
		language.German:               peggyGermanPhrase,
		language.Italian:              peggyItalianPhrase,
		language.Japanese:             peggyJapanesePhrase,
		language.Korean:               peggyKoreanPhrase,
		language.LatinAmericanSpanish: peggyLatinAmericanSpanishPhrase,
		language.Russian:              peggyRussianPhrase,
		language.SimplifiedChinese:    peggySimplifiedChinesePhrase,
		language.Spanish:              peggySpanishPhrase,
		language.TraditionalChinese:   peggyTraditionalChinesePhrase}
)

var (
	// Peggy represents peggy.
	Peggy = nook.Villager{
		Character:   peggyCharacter,
		Personality: personality.Peppy,
		Phrase:      peggyPhrase}
)
