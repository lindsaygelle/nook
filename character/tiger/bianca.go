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
	// biancaBirthday represents Bianca's birthday.
	biancaBirthday = nook.Birthday{
		Day:   13,
		Month: time.December}
)

var (
	// biancaCode represents Bianca's unique code.
	biancaCode = nook.Code{
		Value: "tig06"}
)

var (
	// biancaAmericanEnglishName represents Bianca's name in American English.
	biancaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bianca"}

	// biancaCanadianFrenchName represents Bianca's name in Canadian French.
	biancaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Noémie"}

	// biancaDutchName represents Bianca's name in Dutch.
	biancaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bianca"}

	// biancaFrenchName represents Bianca's name in French.
	biancaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Noémie"}

	// biancaGermanName represents Bianca's name in German.
	biancaGermanName = nook.Name{
		Language: language.German,
		Value:    "Bettina"}

	// biancaItalianName represents Bianca's name in Italian.
	biancaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grattina"}

	// biancaJapaneseName represents Bianca's name in Japanese.
	biancaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "コユキ"}

	// biancaKoreanName represents Bianca's name in Korean.
	biancaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "백희"}

	// biancaLatinAmericanSpanishName represents Bianca's name in Latin American Spanish.
	biancaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bianca"}

	// biancaRussianName represents Bianca's name in Russian.
	biancaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бьянка"}

	// biancaSimplifiedChineseName represents Bianca's name in Simplified Chinese.
	biancaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小雪"}

	// biancaSpanishName represents Bianca's name in Spanish.
	biancaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bianca"}

	// biancaTraditionalChineseName represents Bianca's name in Traditional Chinese.
	biancaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小雪"}
)

var (
	// biancaName represents Bianca's name in different languages.
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
	// biancaCharacter represents Bianca's character information.
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
	// biancaAmericanEnglishPhrase represents Bianca's phrase in American English.
	biancaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "glimmer"}

	// biancaCanadianFrenchPhrase represents Bianca's phrase in Canadian French.
	biancaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "la vie"}

	// biancaDutchPhrase represents Bianca's phrase in Dutch.
	biancaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "witje"}

	// biancaFrenchPhrase represents Bianca's phrase in French.
	biancaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croky"}

	// biancaGermanPhrase represents Bianca's phrase in German.
	biancaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "prrr"}

	// biancaItalianPhrase represents Bianca's phrase in Italian.
	biancaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yeppa"}

	// biancaJapanesePhrase represents Bianca's phrase in Japanese.
	biancaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でヒョウ"}

	// biancaKoreanPhrase represents Bianca's phrase in Korean.
	biancaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그래호"}

	// biancaLatinAmericanSpanishPhrase represents Bianca's phrase in Latin American Spanish.
	biancaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ruarri"}

	// biancaRussianPhrase represents Bianca's phrase in Russian.
	biancaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хлоп-хлоп"}

	// biancaSimplifiedChinesePhrase represents Bianca's phrase in Simplified Chinese.
	biancaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪豹"}

	// biancaSpanishPhrase represents Bianca's phrase in Spanish.
	biancaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guiñito"}

	// biancaTraditionalChinesePhrase represents Bianca's phrase in Traditional Chinese.
	biancaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪豹"}
)

var (
	// biancaPhrase represents Bianca's phrase in different languages.
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
	// Bianca represents the character Bianca with her complete information.
	Bianca = nook.Villager{
		Character:   biancaCharacter,
		Personality: personality.Peppy,
		Phrase:      biancaPhrase}
)
