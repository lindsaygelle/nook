package tiger

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
	biancaBirthday = nook.Birthday{
		Day:   13,
		Month: time.December}
)

var (
	biancaCode = nook.Code{
		Value: "tig06"}
)

var (
	biancaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bianca"}

	biancaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Noémie"}

	biancaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bianca"}

	biancaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Noémie"}

	biancaGermanName = nook.Name{
		Language: language.German,
		Value:    "Bettina"}

	biancaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grattina"}

	biancaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "コユキ"}

	biancaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "백희"}

	biancaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bianca"}

	biancaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бьянка"}

	biancaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小雪"}

	biancaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bianca"}

	biancaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小雪"}
)

var (
	biancaName = nook.Languages{
		language.AmericanEnglish:      biancaAmericanEnglishName,
		language.CanadianFrench:       biancaCanadianFrenchName,
		language.Dutch:                biancaDutchName,
		language.French:               biancaFrenchName,
		language.German:               biancaGermanName,
		language.Italian:              biancaItalianName,
		language.Japanese:             biancaJapaneseName,
		language.Korean:               biancaKoreanName,
		language.LatinAmericanSpanish: biancaLatinAmericanSpanishName,
		language.Russian:              biancaRussianName,
		language.SimplifiedChinese:    biancaSimplifiedChineseName,
		language.Spanish:              biancaSpanishName,
		language.TraditionalChinese:   biancaTraditionalChineseName}
)

var (
	biancaCharacter = nook.Character{
		Animal:   animal.Tiger,
		Birthday: biancaBirthday,
		Code:     biancaCode,
		Key:      character.Bianca,
		Gender:   gender.Female,
		Name:     biancaName,
		Special:  false}
)

var (
	biancaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "glimmer"}

	biancaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "la vie"}

	biancaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "witje"}

	biancaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croky"}

	biancaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "prrr"}

	biancaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yeppa"}

	biancaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でヒョウ"}

	biancaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그래호"}

	biancaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ruarri"}

	biancaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хлоп-хлоп"}

	biancaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪豹"}

	biancaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guiñito"}

	biancaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪豹"}
)

var (
	biancaPhrase = nook.Languages{
		language.AmericanEnglish:      biancaAmericanEnglishPhrase,
		language.CanadianFrench:       biancaCanadianFrenchPhrase,
		language.Dutch:                biancaDutchPhrase,
		language.French:               biancaFrenchPhrase,
		language.German:               biancaGermanPhrase,
		language.Italian:              biancaItalianPhrase,
		language.Japanese:             biancaJapanesePhrase,
		language.Korean:               biancaKoreanPhrase,
		language.LatinAmericanSpanish: biancaLatinAmericanSpanishPhrase,
		language.Russian:              biancaRussianPhrase,
		language.SimplifiedChinese:    biancaSimplifiedChinesePhrase,
		language.Spanish:              biancaSpanishPhrase,
		language.TraditionalChinese:   biancaTraditionalChinesePhrase}
)

var (
	Bianca = nook.Villager{
		Character:   biancaCharacter,
		Personality: personality.Peppy,
		Phrase:      biancaPhrase}
)
