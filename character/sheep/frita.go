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
	// fritaBirthday represents Frita's birthday (July 16th).
	fritaBirthday = nook.Birthday{
		Day:   16,
		Month: time.July}
)

var (
	// fritaCode represents Frita's unique code ("shp11").
	fritaCode = nook.Code{
		Value: "shp11"}
)

var (
	// fritaAmericanEnglishName represents Frita's name in American English.
	fritaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Frita"}

	// fritaCanadianFrenchName represents Frita's name in Canadian French.
	fritaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Clarabêl"}

	// fritaDutchName represents Frita's name in Dutch.
	fritaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Frita"}

	// fritaFrenchName represents Frita's name in French.
	fritaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Clarabêl"}

	// fritaGermanName represents Frita's name in German.
	fritaGermanName = nook.Name{
		Language: language.German,
		Value:    "Martina"}

	// fritaItalianName represents Frita's name in Italian.
	fritaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Landa"}

	// fritaJapaneseName represents Frita's name in Japanese.
	fritaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ウェンディ"}

	// fritaKoreanName represents Frita's name in Korean.
	fritaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "웬디"}

	// fritaLatinAmericanSpanishName represents Frita's name in Latin American Spanish.
	fritaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ovinia"}

	// fritaRussianName represents Frita's name in Russian.
	fritaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фрита"}

	// fritaSimplifiedChineseName represents Frita's name in Simplified Chinese.
	fritaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "温蒂"}

	// fritaSpanishName represents Frita's name in Spanish.
	fritaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ovinia"}

	// fritaTraditionalChineseName represents Frita's name in Traditional Chinese.
	fritaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "溫蒂"}
)

var (
	// fritaName represents Frita's name in different languages.
	fritaName = nook.Languages{
		language.AmericanEnglish:      fritaAmericanEnglishName,
		language.CanadianFrench:       fritaCanadianFrenchName,
		language.Dutch:                fritaDutchName,
		language.French:               fritaFrenchName,
		language.German:               fritaGermanName,
		language.Italian:              fritaItalianName,
		language.Japanese:             fritaJapaneseName,
		language.Korean:               fritaKoreanName,
		language.LatinAmericanSpanish: fritaLatinAmericanSpanishName,
		language.Russian:              fritaRussianName,
		language.SimplifiedChinese:    fritaSimplifiedChineseName,
		language.Spanish:              fritaSpanishName,
		language.TraditionalChinese:   fritaTraditionalChineseName}
)

var (
	// fritaCharacter represents Frita's character information.
	fritaCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: fritaBirthday,
		Code:     fritaCode,
		Key:      character.Frita,
		Gender:   gender.Female,
		Name:     fritaName,
		Special:  false}
)

var (
	// fritaAmericanEnglishPhrase represents Frita's phrase in American English.
	fritaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "oh ewe"}

	// fritaCanadianFrenchPhrase represents Frita's phrase in Canadian French.
	fritaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mouêêê"}

	// fritaDutchPhrase represents Frita's phrase in Dutch.
	fritaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ooi zeg"}

	// fritaFrenchPhrase represents Frita's phrase in French.
	fritaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mouêêê"}

	// fritaGermanPhrase represents Frita's phrase in German.
	fritaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knuddel"}

	// fritaItalianPhrase represents Frita's phrase in Italian.
	fritaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beee la la"}

	// fritaJapanesePhrase represents Frita's phrase in Japanese.
	fritaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だポテ"}

	// fritaKoreanPhrase represents Frita's phrase in Korean.
	fritaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "포테토"}

	// fritaLatinAmericanSpanishPhrase represents Frita's phrase in Latin American Spanish.
	fritaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "berebé"}

	// fritaRussianPhrase represents Frita's phrase in Russian.
	fritaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мутончик"}

	// fritaSimplifiedChinesePhrase represents Frita's phrase in Simplified Chinese.
	fritaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "薯薯"}

	// fritaSpanishPhrase represents Frita's phrase in Spanish.
	fritaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "berebé"}

	// fritaTraditionalChinesePhrase represents Frita's phrase in Traditional Chinese.
	fritaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "薯薯"}
)

var (
	// fritaPhrase represents Frita's phrases in different languages.
	fritaPhrase = nook.Languages{
		language.AmericanEnglish:      fritaAmericanEnglishPhrase,
		language.CanadianFrench:       fritaCanadianFrenchPhrase,
		language.Dutch:                fritaDutchPhrase,
		language.French:               fritaFrenchPhrase,
		language.German:               fritaGermanPhrase,
		language.Italian:              fritaItalianPhrase,
		language.Japanese:             fritaJapanesePhrase,
		language.Korean:               fritaKoreanPhrase,
		language.LatinAmericanSpanish: fritaLatinAmericanSpanishPhrase,
		language.Russian:              fritaRussianPhrase,
		language.SimplifiedChinese:    fritaSimplifiedChinesePhrase,
		language.Spanish:              fritaSpanishPhrase,
		language.TraditionalChinese:   fritaTraditionalChinesePhrase}
)

var (
	// Frita represents the character Frita with her complete information.
	Frita = nook.Villager{
		Character:   fritaCharacter,
		Personality: personality.BigSister,
		Phrase:      fritaPhrase}
)
