package sheep

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
	// stellaBirthday represents Stella's birthday.
	stellaBirthday = nook.Birthday{
		Day:   9,
		Month: time.April}
)

var (
	// stellaCode represents Stella's unique code.
	stellaCode = nook.Code{
		Value: "shp03"}
)

var (
	// stellaAmericanEnglishName represents Stella's name in American English.
	stellaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Stella"}

	// stellaCanadianFrenchName represents Stella's name in Canadian French.
	stellaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bigoudi"}

	// stellaDutchName represents Stella's name in Dutch.
	stellaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Stella"}

	// stellaFrenchName represents Stella's name in French.
	stellaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bigoudi"}

	// stellaGermanName represents Stella's name in German.
	stellaGermanName = nook.Name{
		Language: language.German,
		Value:    "Stella"}

	// stellaItalianName represents Stella's name in Italian.
	stellaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Merina"}

	// stellaJapaneseName represents Stella's name in Japanese.
	stellaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アクリル"}

	// stellaKoreanName represents Stella's name in Korean.
	stellaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아크릴"}

	// stellaLatinAmericanSpanishName represents Stella's name in Latin American Spanish.
	stellaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Merina"}

	// stellaRussianName represents Stella's name in Russian.
	stellaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стелла"}

	// stellaSimplifiedChineseName represents Stella's name in Simplified Chinese.
	stellaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毡毡"}

	// stellaSpanishName represents Stella's name in Spanish.
	stellaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Merina"}

	// stellaTraditionalChineseName represents Stella's name in Traditional Chinese.
	stellaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "氈氈"}
)

var (
	// stellaName represents Stella's name in different languages.
	stellaName = nook.Languages{
		language.AmericanEnglish:      stellaAmericanEnglishName,
		language.CanadianFrench:       stellaCanadianFrenchName,
		language.Dutch:                stellaDutchName,
		language.French:               stellaFrenchName,
		language.German:               stellaGermanName,
		language.Italian:              stellaItalianName,
		language.Japanese:             stellaJapaneseName,
		language.Korean:               stellaKoreanName,
		language.LatinAmericanSpanish: stellaLatinAmericanSpanishName,
		language.Russian:              stellaRussianName,
		language.SimplifiedChinese:    stellaSimplifiedChineseName,
		language.Spanish:              stellaSpanishName,
		language.TraditionalChinese:   stellaTraditionalChineseName}
)

var (
	// stellaCharacter represents Stella's character information.
	stellaCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: stellaBirthday,
		Code:     stellaCode,
		Key:      character.Stella,
		Gender:   gender.Female,
		Name:     stellaName,
		Special:  false}
)

var (
	// stellaAmericanEnglishPhrase represents Stella's phrase in American English.
	stellaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "baa-dabing"}

	// stellaCanadianFrenchPhrase represents Stella's phrase in Canadian French.
	stellaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon agneau"}

	// stellaDutchPhrase represents Stella's phrase in Dutch.
	stellaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mèèèh"}

	// stellaFrenchPhrase represents Stella's phrase in French.
	stellaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon agneau"}

	// stellaGermanPhrase represents Stella's phrase in German.
	stellaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "badabing"}

	// stellaItalianPhrase represents Stella's phrase in Italian.
	stellaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beeene"}

	// stellaJapanesePhrase represents Stella's phrase in Japanese.
	stellaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウール"}

	// stellaKoreanPhrase represents Stella's phrase in Korean.
	stellaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "울"}

	// stellaLatinAmericanSpanishPhrase represents Stella's phrase in Latin American Spanish.
	stellaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "baa-bii"}

	// stellaRussianPhrase represents Stella's phrase in Russian.
	stellaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бэ-дабинь"}

	// stellaSimplifiedChinesePhrase represents Stella's phrase in Simplified Chinese.
	stellaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "羊毛"}

	// stellaSpanishPhrase represents Stella's phrase in Spanish.
	stellaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "no veeeas"}

	// stellaTraditionalChinesePhrase represents Stella's phrase in Traditional Chinese.
	stellaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羊毛"}
)

var (
	// stellaPhrase represents Stella's phrases in different languages.
	stellaPhrase = nook.Languages{
		language.AmericanEnglish:      stellaAmericanEnglishPhrase,
		language.CanadianFrench:       stellaCanadianFrenchPhrase,
		language.Dutch:                stellaDutchPhrase,
		language.French:               stellaFrenchPhrase,
		language.German:               stellaGermanPhrase,
		language.Italian:              stellaItalianPhrase,
		language.Japanese:             stellaJapanesePhrase,
		language.Korean:               stellaKoreanPhrase,
		language.LatinAmericanSpanish: stellaLatinAmericanSpanishPhrase,
		language.Russian:              stellaRussianPhrase,
		language.SimplifiedChinese:    stellaSimplifiedChinesePhrase,
		language.Spanish:              stellaSpanishPhrase,
		language.TraditionalChinese:   stellaTraditionalChinesePhrase}
)

var (
	// Stella represents the character Stella with her complete information.
	Stella = nook.Villager{
		Character:   stellaCharacter,
		Personality: personality.Normal,
		Phrase:      stellaPhrase}
)
