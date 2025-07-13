package bearcub

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
	// kodyBirthday represents kody birthday.
	kodyBirthday = nook.Birthday{
		Day:   28,
		Month: time.September}
)

var (
	// kodyCode represents kody code.
	kodyCode = nook.Code{
		Value: "cbr04"}
)

var (
	// kodyAmericanEnglishName represents kody american english name.
	kodyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kody"}

	// kodyCanadianFrenchName represents kody canadian french name.
	kodyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bill"}

	// kodyDutchName represents kody dutch name.
	kodyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kody"}

	// kodyFrenchName represents kody french name.
	kodyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bill"}

	// kodyGermanName represents kody german name.
	kodyGermanName = nook.Name{
		Language: language.German,
		Value:    "Bernd"}

	// kodyItalianName represents kody italian name.
	kodyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Salomone"}

	// kodyJapaneseName represents kody japanese name.
	kodyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイダホ"}

	// kodyKoreanName represents kody korean name.
	kodyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이다호"}

	// kodyLatinAmericanSpanishName represents kody latin american spanish name.
	kodyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Orestes"}

	// kodyRussianName represents kody russian name.
	kodyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Коди"}

	// kodySimplifiedChineseName represents kody simplified chinese name.
	kodySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艾德豪"}

	// kodySpanishName represents kody spanish name.
	kodySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Orestes"}

	// kodyTraditionalChineseName represents kody traditional chinese name.
	kodyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艾德豪"}
)

var (
	// kodyName represents kody name.
	kodyName = nook.Languages{
		language.AmericanEnglish:      kodyAmericanEnglishName,
		language.CanadianFrench:       kodyCanadianFrenchName,
		language.Dutch:                kodyDutchName,
		language.French:               kodyFrenchName,
		language.German:               kodyGermanName,
		language.Italian:              kodyItalianName,
		language.Japanese:             kodyJapaneseName,
		language.Korean:               kodyKoreanName,
		language.LatinAmericanSpanish: kodyLatinAmericanSpanishName,
		language.Russian:              kodyRussianName,
		language.SimplifiedChinese:    kodySimplifiedChineseName,
		language.Spanish:              kodySpanishName,
		language.TraditionalChinese:   kodyTraditionalChineseName}
)

var (
	// kodyCharacter represents kody character.
	kodyCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: kodyBirthday,
		Code:     kodyCode,
		Key:      character.Kody,
		Gender:   gender.Male,
		Name:     kodyName,
		Special:  false}
)

var (
	// kodyAmericanEnglishPhrase represents kody american english phrase.
	kodyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "okey-dokey"}

	// kodyCanadianFrenchPhrase represents kody canadian french phrase.
	kodyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grih grih"}

	// kodyDutchPhrase represents kody dutch phrase.
	kodyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "okidoki"}

	// kodyFrenchPhrase represents kody french phrase.
	kodyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grih grih"}

	// kodyGermanPhrase represents kody german phrase.
	kodyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "okeydokey"}

	// kodyItalianPhrase represents kody italian phrase.
	kodyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oki-doki"}

	// kodyJapanesePhrase represents kody japanese phrase.
	kodyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "のこころ"}

	// kodyKoreanPhrase represents kody korean phrase.
	kodyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "나는나야"}

	// kodyLatinAmericanSpanishPhrase represents kody latin american spanish phrase.
	kodyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "graaau"}

	// kodyRussianPhrase represents kody russian phrase.
	kodyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гррр"}

	// kodySimplifiedChinesePhrase represents kody simplified chinese phrase.
	kodySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "精神"}

	// kodySpanishPhrase represents kody spanish phrase.
	kodySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mola"}

	// kodyTraditionalChinesePhrase represents kody traditional chinese phrase.
	kodyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "精神"}
)

var (
	// kodyPhrase represents kody phrase.
	kodyPhrase = nook.Languages{
		language.AmericanEnglish:      kodyAmericanEnglishPhrase,
		language.CanadianFrench:       kodyCanadianFrenchPhrase,
		language.Dutch:                kodyDutchPhrase,
		language.French:               kodyFrenchPhrase,
		language.German:               kodyGermanPhrase,
		language.Italian:              kodyItalianPhrase,
		language.Japanese:             kodyJapanesePhrase,
		language.Korean:               kodyKoreanPhrase,
		language.LatinAmericanSpanish: kodyLatinAmericanSpanishPhrase,
		language.Russian:              kodyRussianPhrase,
		language.SimplifiedChinese:    kodySimplifiedChinesePhrase,
		language.Spanish:              kodySpanishPhrase,
		language.TraditionalChinese:   kodyTraditionalChinesePhrase}
)

var (
	// Kody represents kody.
	Kody = nook.Villager{
		Character:   kodyCharacter,
		Personality: personality.Jock,
		Phrase:      kodyPhrase}
)
