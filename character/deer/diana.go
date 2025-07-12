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
	// dianaBirthday represents diana birthday.
	dianaBirthday = nook.Birthday{
		Day:   4,
		Month: time.January}
)

var (
	// dianaCode represents diana code.
	dianaCode = nook.Code{
		Value: "der08"}
)

var (
	// dianaAmericanEnglishName represents diana american english name.
	dianaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Diana"}

	// dianaCanadianFrenchName represents diana canadian french name.
	dianaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Didi"}

	// dianaDutchName represents diana dutch name.
	dianaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Diana"}

	// dianaFrenchName represents diana french name.
	dianaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Didi"}

	// dianaGermanName represents diana german name.
	dianaGermanName = nook.Name{
		Language: language.German,
		Value:    "Vroni"}

	// dianaItalianName represents diana italian name.
	dianaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Diana"}

	// dianaJapaneseName represents diana japanese name.
	dianaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナタリー"}

	// dianaKoreanName represents diana korean name.
	dianaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나탈리"}

	// dianaLatinAmericanSpanishName represents diana latin american spanish name.
	dianaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bambina"}

	// dianaRussianName represents diana russian name.
	dianaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Диана"}

	// dianaSimplifiedChineseName represents diana simplified chinese name.
	dianaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "倪家莉"}

	// dianaSpanishName represents diana spanish name.
	dianaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bambina"}

	// dianaTraditionalChineseName represents diana traditional chinese name.
	dianaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "倪家莉"}
)

var (
	// dianaName represents diana name.
	dianaName = nook.Languages{
		language.AmericanEnglish:      dianaAmericanEnglishName,
		language.CanadianFrench:       dianaCanadianFrenchName,
		language.Dutch:                dianaDutchName,
		language.French:               dianaFrenchName,
		language.German:               dianaGermanName,
		language.Italian:              dianaItalianName,
		language.Japanese:             dianaJapaneseName,
		language.Korean:               dianaKoreanName,
		language.LatinAmericanSpanish: dianaLatinAmericanSpanishName,
		language.Russian:              dianaRussianName,
		language.SimplifiedChinese:    dianaSimplifiedChineseName,
		language.Spanish:              dianaSpanishName,
		language.TraditionalChinese:   dianaTraditionalChineseName}
)

var (
	// dianaCharacter represents diana character.
	dianaCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: dianaBirthday,
		Code:     dianaCode,
		Key:      character.Diana,
		Gender:   gender.Female,
		Name:     dianaName,
		Special:  false}
)

var (
	// dianaAmericanEnglishPhrase represents diana american english phrase.
	dianaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "no doy"}

	// dianaCanadianFrenchPhrase represents diana canadian french phrase.
	dianaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bichouchou"}

	// dianaDutchPhrase represents diana dutch phrase.
	dianaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "joh"}

	// dianaFrenchPhrase represents diana french phrase.
	dianaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bichouchou"}

	// dianaGermanPhrase represents diana german phrase.
	dianaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kitz"}

	// dianaItalianPhrase represents diana italian phrase.
	dianaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mamy"}

	// dianaJapanesePhrase represents diana japanese phrase.
	dianaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でしょ"}

	// dianaKoreanPhrase represents diana korean phrase.
	dianaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라니까"}

	// dianaLatinAmericanSpanishPhrase represents diana latin american spanish phrase.
	dianaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nonuá"}

	// dianaRussianPhrase represents diana russian phrase.
	dianaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "еще бы"}

	// dianaSimplifiedChinesePhrase represents diana simplified chinese phrase.
	dianaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是吧"}

	// dianaSpanishPhrase represents diana spanish phrase.
	dianaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "braaavo"}

	// dianaTraditionalChinesePhrase represents diana traditional chinese phrase.
	dianaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是吧"}
)

var (
	// dianaPhrase represents diana phrase.
	dianaPhrase = nook.Languages{
		language.AmericanEnglish:      dianaAmericanEnglishPhrase,
		language.CanadianFrench:       dianaCanadianFrenchPhrase,
		language.Dutch:                dianaDutchPhrase,
		language.French:               dianaFrenchPhrase,
		language.German:               dianaGermanPhrase,
		language.Italian:              dianaItalianPhrase,
		language.Japanese:             dianaJapanesePhrase,
		language.Korean:               dianaKoreanPhrase,
		language.LatinAmericanSpanish: dianaLatinAmericanSpanishPhrase,
		language.Russian:              dianaRussianPhrase,
		language.SimplifiedChinese:    dianaSimplifiedChinesePhrase,
		language.Spanish:              dianaSpanishPhrase,
		language.TraditionalChinese:   dianaTraditionalChinesePhrase}
)

var (
	// Diana represents diana.
	Diana = nook.Villager{
		Character:   dianaCharacter,
		Personality: personality.Snooty,
		Phrase:      dianaPhrase}
)
