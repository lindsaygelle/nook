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
	// svenBirthday represents sven birthday.
	svenBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// svenCode represents sven code.
	svenCode = nook.Code{
		Value: ""}
)

var (
	// svenAmericanEnglishName represents sven american english name.
	svenAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sven"}

	// svenCanadianFrenchName represents sven canadian french name.
	svenCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// svenDutchName represents sven dutch name.
	svenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// svenFrenchName represents sven french name.
	svenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sven"}

	// svenGermanName represents sven german name.
	svenGermanName = nook.Name{
		Language: language.German,
		Value:    "Sven"}

	// svenItalianName represents sven italian name.
	svenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giacobbe"}

	// svenJapaneseName represents sven japanese name.
	svenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャマジ"}

	// svenKoreanName represents sven korean name.
	svenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// svenLatinAmericanSpanishName represents sven latin american spanish name.
	svenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// svenRussianName represents sven russian name.
	svenRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// svenSimplifiedChineseName represents sven simplified chinese name.
	svenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "伯伯"}

	// svenSpanishName represents sven spanish name.
	svenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sven"}

	// svenTraditionalChineseName represents sven traditional chinese name.
	svenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// svenName represents sven name.
	svenName = nook.Languages{
		language.AmericanEnglish:      svenAmericanEnglishName,
		language.CanadianFrench:       svenCanadianFrenchName,
		language.Dutch:                svenDutchName,
		language.French:               svenFrenchName,
		language.German:               svenGermanName,
		language.Italian:              svenItalianName,
		language.Japanese:             svenJapaneseName,
		language.Korean:               svenKoreanName,
		language.LatinAmericanSpanish: svenLatinAmericanSpanishName,
		language.Russian:              svenRussianName,
		language.SimplifiedChinese:    svenSimplifiedChineseName,
		language.Spanish:              svenSpanishName,
		language.TraditionalChinese:   svenTraditionalChineseName}
)

var (
	// svenCharacter represents sven character.
	svenCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: svenBirthday,
		Code:     svenCode,
		Key:      character.Sven,
		Gender:   gender.Male,
		Name:     svenName,
		Special:  false}
)

var (
	// svenAmericanEnglishPhrase represents sven american english phrase.
	svenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "buh-uh-ud"}

	// svenCanadianFrenchPhrase represents sven canadian french phrase.
	svenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// svenDutchPhrase represents sven dutch phrase.
	svenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// svenFrenchPhrase represents sven french phrase.
	svenFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "beuh-heu"}

	// svenGermanPhrase represents sven german phrase.
	svenGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "moooh"}

	// svenItalianPhrase represents sven italian phrase.
	svenItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bee eh"}

	// svenJapanesePhrase represents sven japanese phrase.
	svenJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とほほ"}

	// svenKoreanPhrase represents sven korean phrase.
	svenKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// svenLatinAmericanSpanishPhrase represents sven latin american spanish phrase.
	svenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// svenRussianPhrase represents sven russian phrase.
	svenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// svenSimplifiedChinesePhrase represents sven simplified chinese phrase.
	svenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "活活"}

	// svenSpanishPhrase represents sven spanish phrase.
	svenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "buuueeeeh"}

	// svenTraditionalChinesePhrase represents sven traditional chinese phrase.
	svenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// svenPhrase represents sven phrase.
	svenPhrase = nook.Languages{
		language.AmericanEnglish:      svenAmericanEnglishPhrase,
		language.CanadianFrench:       svenCanadianFrenchPhrase,
		language.Dutch:                svenDutchPhrase,
		language.French:               svenFrenchPhrase,
		language.German:               svenGermanPhrase,
		language.Italian:              svenItalianPhrase,
		language.Japanese:             svenJapanesePhrase,
		language.Korean:               svenKoreanPhrase,
		language.LatinAmericanSpanish: svenLatinAmericanSpanishPhrase,
		language.Russian:              svenRussianPhrase,
		language.SimplifiedChinese:    svenSimplifiedChinesePhrase,
		language.Spanish:              svenSpanishPhrase,
		language.TraditionalChinese:   svenTraditionalChinesePhrase}
)

var (
	// Sven represents sven.
	Sven = nook.Villager{
		Character:   svenCharacter,
		Personality: personality.Lazy,
		Phrase:      svenPhrase}
)
