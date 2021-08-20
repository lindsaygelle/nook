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
	kodyBirthday = nook.Birthday{
		Day:   28,
		Month: time.September}
)

var (
	kodyCode = nook.Code{
		Value: "cbr04"}
)

var (
	kodyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kody"}

	kodyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bill"}

	kodyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kody"}

	kodyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bill"}

	kodyGermanName = nook.Name{
		Language: language.German,
		Value:    "Bernd"}

	kodyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Salomone"}

	kodyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイダホ"}

	kodyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이다호"}

	kodyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Orestes"}

	kodyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Коди"}

	kodySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艾德豪"}

	kodySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Orestes"}

	kodyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艾德豪"}
)

var (
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
	kodyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "okey-dokey"}

	kodyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grih grih"}

	kodyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "okidoki"}

	kodyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grih grih"}

	kodyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "okeydokey"}

	kodyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oki-doki"}

	kodyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "のこころ"}

	kodyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "나는나야"}

	kodyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "graaau"}

	kodyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гррр"}

	kodySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "精神"}

	kodySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mola"}

	kodyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "精神"}
)

var (
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
	Kody = nook.Villager{
		Character:   kodyCharacter,
		Personality: personality.Jock,
		Phrase:      kodyPhrase}
)
