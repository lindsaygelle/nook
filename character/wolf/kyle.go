package wolf

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
	// kyleBirthday represents Kyle's birthday (December 6th).
	kyleBirthday = nook.Birthday{
		Day:   6,
		Month: time.December}
)

var (
	// kyleCode represents Kyle's unique code ("wol10").
	kyleCode = nook.Code{
		Value: "wol10"}
)

var (
	// kyleAmericanEnglishName represents Kyle's name in American English.
	kyleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kyle"}

	// kyleCanadianFrenchName represents Kyle's name in Canadian French.
	kyleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gary"}

	// kyleDutchName represents Kyle's name in Dutch.
	kyleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kyle"}

	// kyleFrenchName represents Kyle's name in French.
	kyleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gary"}

	// kyleGermanName represents Kyle's name in German.
	kyleGermanName = nook.Name{
		Language: language.German,
		Value:    "Wolfgang"}

	// kyleItalianName represents Kyle's name in Italian.
	kyleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ululo"}

	// kyleJapaneseName represents Kyle's name in Japanese.
	kyleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リカルド"}

	// kyleKoreanName represents Kyle's name in Korean.
	kyleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리카르도"}

	// kyleLatinAmericanSpanishName represents Kyle's name in Latin American Spanish.
	kyleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ataúlfo"}

	// kyleRussianName represents Kyle's name in Russian.
	kyleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кайл"}

	// kyleSimplifiedChineseName represents Kyle's name in Simplified Chinese.
	kyleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "李可"}

	// kyleSpanishName represents Kyle's name in Spanish.
	kyleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ataúlfo"}

	// kyleTraditionalChineseName represents Kyle's name in Traditional Chinese.
	kyleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "李可"}
)

var (
	// kyleName represents Kyle's name in different languages.
	kyleName = nook.Languages{
		language.AmericanEnglish:      kyleAmericanEnglishName,
		language.CanadianFrench:       kyleCanadianFrenchName,
		language.Dutch:                kyleDutchName,
		language.French:               kyleFrenchName,
		language.German:               kyleGermanName,
		language.Italian:              kyleItalianName,
		language.Japanese:             kyleJapaneseName,
		language.Korean:               kyleKoreanName,
		language.LatinAmericanSpanish: kyleLatinAmericanSpanishName,
		language.Russian:              kyleRussianName,
		language.SimplifiedChinese:    kyleSimplifiedChineseName,
		language.Spanish:              kyleSpanishName,
		language.TraditionalChinese:   kyleTraditionalChineseName}
)

var (
	// kyleCharacter represents Kyle's character information.
	kyleCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: kyleBirthday,
		Code:     kyleCode,
		Key:      character.Kyle,
		Gender:   gender.Male,
		Name:     kyleName,
		Special:  false}
)

var (
	// kyleAmericanEnglishPhrase represents Kyle's phrase in American English.
	kyleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "alpha"}

	// kyleCanadianFrenchPhrase represents Kyle's phrase in Canadian French.
	kyleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "garooouuu"}

	// kyleDutchPhrase represents Kyle's phrase in Dutch.
	kyleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "alfa"}

	// kyleFrenchPhrase represents Kyle's phrase in French.
	kyleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "garooouuu"}

	// kyleGermanPhrase represents Kyle's phrase in German.
	kyleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ahuuu"}

	// kyleItalianPhrase represents Kyle's phrase in Italian.
	kyleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sgnack"}

	// kyleJapanesePhrase represents Kyle's phrase in Japanese.
	kyleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "オゥイェ"}

	// kyleKoreanPhrase represents Kyle's phrase in Korean.
	kyleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오우예"}

	// kyleLatinAmericanSpanishPhrase represents Kyle's phrase in Latin American Spanish.
	kyleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "aujujú"}

	// kyleRussianPhrase represents Kyle's phrase in Russian.
	kyleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вожак"}

	// kyleSimplifiedChinesePhrase represents Kyle's phrase in Simplified Chinese.
	kyleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喔耶"}

	// kyleSpanishPhrase represents Kyle's phrase in Spanish.
	kyleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aujujú"}

	// kyleTraditionalChinesePhrase represents Kyle's phrase in Traditional Chinese.
	kyleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喔耶"}
)

var (
	// kylePhrase represents Kyle's phrase in different languages.
	kylePhrase = nook.Languages{
		language.AmericanEnglish:      kyleAmericanEnglishPhrase,
		language.CanadianFrench:       kyleCanadianFrenchPhrase,
		language.Dutch:                kyleDutchPhrase,
		language.French:               kyleFrenchPhrase,
		language.German:               kyleGermanPhrase,
		language.Italian:              kyleItalianPhrase,
		language.Japanese:             kyleJapanesePhrase,
		language.Korean:               kyleKoreanPhrase,
		language.LatinAmericanSpanish: kyleLatinAmericanSpanishPhrase,
		language.Russian:              kyleRussianPhrase,
		language.SimplifiedChinese:    kyleSimplifiedChinesePhrase,
		language.Spanish:              kyleSpanishPhrase,
		language.TraditionalChinese:   kyleTraditionalChinesePhrase}
)

var (
	// Kyle represents the character Kyle with his complete information.
	Kyle = nook.Villager{
		Character:   kyleCharacter,
		Personality: personality.Smug,
		Phrase:      kylePhrase}
)
