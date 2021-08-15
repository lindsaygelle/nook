package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	antonioBirthday = nook.Birthday{
		Day:   20,
		Month: time.October}
)

var (
	antonioCode = nook.Code{
		Value: "ant01"}
)

var (
	antonioAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Antonio"}

	antonioCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Antonio"}

	antonioDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Antonio"}

	antonioFrenchName = nook.Name{
		Language: language.French,
		Value:    "Antonio"}

	antonioGermanName = nook.Name{
		Language: language.German,
		Value:    "Siggi"}

	antonioItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Antonio"}

	antonioJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マコト"}

	antonioKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "퍼머거"}

	antonioLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Antonio"}

	antonioRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Антонио"}

	antonioSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿诚"}

	antonioSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Antonio"}

	antonioTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿誠"}
)

var (
	antonioName = nook.Languages{
		language.AmericanEnglish:      antonioAmericanEnglishName,
		language.CanadianFrench:       antonioCanadianFrenchName,
		language.Dutch:                antonioDutchName,
		language.French:               antonioFrenchName,
		language.German:               antonioGermanName,
		language.Italian:              antonioItalianName,
		language.Japanese:             antonioJapaneseName,
		language.Korean:               antonioKoreanName,
		language.LatinAmericanSpanish: antonioLatinAmericanSpanishName,
		language.Russian:              antonioRussianName,
		language.SimplifiedChinese:    antonioSimplifiedChineseName,
		language.Spanish:              antonioSpanishName,
		language.TraditionalChinese:   antonioTraditionalChineseName}
)

var (
	antonioCharacter = nook.Character{
		Animal:   Anteater,
		Birthday: antonioBirthday,
		Code:     antonioCode,
		Gender:   nook.Male,
		Name:     antonioName}
)

var (
	antonioAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "honk"}

	antonioCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pouet"}

	antonioDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuit"}

	antonioFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pouet"}

	antonioGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schlürrrf"}

	antonioItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "honk"}

	antonioJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ホントに"}

	antonioKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "진짜로"}

	antonioLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fufuf"}

	antonioRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "го-го-го"}

	antonioSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "真诚"}

	antonioSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fufuf"}

	antonioTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "真誠"}
)

var (
	antonioPhrase = nook.Languages{
		language.AmericanEnglish:      antonioAmericanEnglishName,
		language.CanadianFrench:       antonioCanadianFrenchName,
		language.Dutch:                antonioDutchName,
		language.French:               antonioFrenchName,
		language.German:               antonioGermanName,
		language.Italian:              antonioItalianName,
		language.Japanese:             antonioJapaneseName,
		language.Korean:               antonioKoreanName,
		language.LatinAmericanSpanish: antonioLatinAmericanSpanishName,
		language.Russian:              antonioRussianName,
		language.SimplifiedChinese:    antonioSimplifiedChineseName,
		language.Spanish:              antonioSpanishName,
		language.TraditionalChinese:   antonioTraditionalChineseName}
)

var (
	Antonio = nook.Villager{
		Character:   antonioCharacter,
		Personality: nook.Jock,
		Phrase:      antonioPhrase}
)
