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
	fritaBirthday = nook.Birthday{
		Day:   16,
		Month: time.July}
)

var (
	fritaCode = nook.Code{
		Value: "shp11"}
)

var (
	fritaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Frita"}

	fritaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Clarabêl"}

	fritaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Frita"}

	fritaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Clarabêl"}

	fritaGermanName = nook.Name{
		Language: language.German,
		Value:    "Martina"}

	fritaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Landa"}

	fritaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ウェンディ"}

	fritaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "웬디"}

	fritaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ovinia"}

	fritaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фрита"}

	fritaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "温蒂"}

	fritaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ovinia"}

	fritaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "溫蒂"}
)

var (
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
	fritaCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: fritaBirthday,
		Code:     fritaCode,
		Key:      character.Frita,
		Gender:   gender.Female,
		Name:     fritaName}
)

var (
	fritaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "oh ewe"}

	fritaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mouêêê"}

	fritaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ooi zeg"}

	fritaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mouêêê"}

	fritaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knuddel"}

	fritaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beee la la"}

	fritaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だポテ"}

	fritaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "포테토"}

	fritaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "berebé"}

	fritaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мутончик"}

	fritaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "薯薯"}

	fritaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "berebé"}

	fritaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "薯薯"}
)

var (
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
	Frita = nook.Villager{
		Character:   fritaCharacter,
		Personality: personality.BigSister,
		Phrase:      fritaPhrase}
)
