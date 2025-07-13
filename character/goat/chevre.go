package goat

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
	// chevreBirthday represents chevre birthday.
	chevreBirthday = nook.Birthday{
		Day:   6,
		Month: time.March}
)

var (
	// chevreCode represents chevre code.
	chevreCode = nook.Code{
		Value: "goa00"}
)

var (
	// chevreAmericanEnglishName represents chevre american english name.
	chevreAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chevre"}

	// chevreCanadianFrenchName represents chevre canadian french name.
	chevreCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Biquette"}

	// chevreDutchName represents chevre dutch name.
	chevreDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chevre"}

	// chevreFrenchName represents chevre french name.
	chevreFrenchName = nook.Name{
		Language: language.French,
		Value:    "Biquette"}

	// chevreGermanName represents chevre german name.
	chevreGermanName = nook.Name{
		Language: language.German,
		Value:    "Anette"}

	// chevreItalianName represents chevre italian name.
	chevreItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Diletta"}

	// chevreJapaneseName represents chevre japanese name.
	chevreJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ユキ"}

	// chevreKoreanName represents chevre korean name.
	chevreKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "윤이"}

	// chevreLatinAmericanSpanishName represents chevre latin american spanish name.
	chevreLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cabriola"}

	// chevreRussianName represents chevre russian name.
	chevreRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шевр"}

	// chevreSimplifiedChineseName represents chevre simplified chinese name.
	chevreSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪儿"}

	// chevreSpanishName represents chevre spanish name.
	chevreSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cabriola"}

	// chevreTraditionalChineseName represents chevre traditional chinese name.
	chevreTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪兒"}
)

var (
	// chevreName represents chevre name.
	chevreName = nook.Languages{
		language.AmericanEnglish:      chevreAmericanEnglishName,
		language.CanadianFrench:       chevreCanadianFrenchName,
		language.Dutch:                chevreDutchName,
		language.French:               chevreFrenchName,
		language.German:               chevreGermanName,
		language.Italian:              chevreItalianName,
		language.Japanese:             chevreJapaneseName,
		language.Korean:               chevreKoreanName,
		language.LatinAmericanSpanish: chevreLatinAmericanSpanishName,
		language.Russian:              chevreRussianName,
		language.SimplifiedChinese:    chevreSimplifiedChineseName,
		language.Spanish:              chevreSpanishName,
		language.TraditionalChinese:   chevreTraditionalChineseName}
)

var (
	// chevreCharacter represents chevre character.
	chevreCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: chevreBirthday,
		Code:     chevreCode,
		Key:      character.Chevre,
		Gender:   gender.Female,
		Name:     chevreName,
		Special:  false}
)

var (
	// chevreAmericanEnglishPhrase represents chevre american english phrase.
	chevreAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "la baa"}

	// chevreCanadianFrenchPhrase represents chevre canadian french phrase.
	chevreCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "la baa"}

	// chevreDutchPhrase represents chevre dutch phrase.
	chevreDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "la mèè"}

	// chevreFrenchPhrase represents chevre french phrase.
	chevreFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "la baa"}

	// chevreGermanPhrase represents chevre german phrase.
	chevreGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "määääh"}

	// chevreItalianPhrase represents chevre italian phrase.
	chevreItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beeello"}

	// chevreJapanesePhrase represents chevre japanese phrase.
	chevreJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "っぺ"}

	// chevreKoreanPhrase represents chevre korean phrase.
	chevreKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "맞아유"}

	// chevreLatinAmericanSpanishPhrase represents chevre latin american spanish phrase.
	chevreLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "beee-beee"}

	// chevreRussianPhrase represents chevre russian phrase.
	chevreRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ле ме-е-е"}

	// chevreSimplifiedChinesePhrase represents chevre simplified chinese phrase.
	chevreSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呗"}

	// chevreSpanishPhrase represents chevre spanish phrase.
	chevreSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "beee-beee"}

	// chevreTraditionalChinesePhrase represents chevre traditional chinese phrase.
	chevreTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唄"}
)

var (
	// chevrePhrase represents chevre phrase.
	chevrePhrase = nook.Languages{
		language.AmericanEnglish:      chevreAmericanEnglishPhrase,
		language.CanadianFrench:       chevreCanadianFrenchPhrase,
		language.Dutch:                chevreDutchPhrase,
		language.French:               chevreFrenchPhrase,
		language.German:               chevreGermanPhrase,
		language.Italian:              chevreItalianPhrase,
		language.Japanese:             chevreJapanesePhrase,
		language.Korean:               chevreKoreanPhrase,
		language.LatinAmericanSpanish: chevreLatinAmericanSpanishPhrase,
		language.Russian:              chevreRussianPhrase,
		language.SimplifiedChinese:    chevreSimplifiedChinesePhrase,
		language.Spanish:              chevreSpanishPhrase,
		language.TraditionalChinese:   chevreTraditionalChinesePhrase}
)

var (
	// Chevre represents chevre.
	Chevre = nook.Villager{
		Character:   chevreCharacter,
		Personality: personality.Normal,
		Phrase:      chevrePhrase}
)
