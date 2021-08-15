package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Billgrih grih"}

	kodyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kodyokidoki"}

	kodyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Billgrih grih"}

	kodyGermanName = nook.Name{
		Language: language.German,
		Value:    "Berndokeydokey"}

	kodyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Salomoneoki-doki"}

	kodyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイダホのこころ"}

	kodyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이다호나는나야"}

	kodyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Orestesgraaau"}

	kodyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кодигррр"}

	kodySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艾德豪精神"}

	kodySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Orestesmola"}

	kodyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艾德豪精神"}
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
		Animal:   Bearcub,
		Birthday: kodyBirthday,
		Code:     kodyCode,
		Gender:   nook.Male,
		Name:     kodyName}
)

var (
	kodyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "okey-dokeygrah grah"}

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
	Kody = nook.Villager{
		Character:   kodyCharacter,
		Personality: nook.Jock,
		Phrase:      kodyPhrase}
)
