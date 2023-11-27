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
	// freyaBirthday represents Freya's birthday.
	freyaBirthday = nook.Birthday{
		Day:   14,
		Month: time.December}
)

var (
	// freyaCode represents Freya's unique code.
	freyaCode = nook.Code{
		Value: "wol05"}
)

var (
	// freyaAmericanEnglishName represents Freya's name in American English.
	freyaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Freya"}

	// freyaCanadianFrenchName represents Freya's name in Canadian French.
	freyaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Luppa"}

	// freyaDutchName represents Freya's name in Dutch.
	freyaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Freya"}

	// freyaFrenchName represents Freya's name in French.
	freyaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Luppa"}

	// freyaGermanName represents Freya's name in German.
	freyaGermanName = nook.Name{
		Language: language.German,
		Value:    "Freya"}

	// freyaItalianName represents Freya's name in Italian.
	freyaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Freya"}

	// freyaJapaneseName represents Freya's name in Japanese.
	freyaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ツンドラ"}

	// freyaKoreanName represents Freya's name in Korean.
	freyaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "산드라"}

	// freyaLatinAmericanSpanishName represents Freya's name in Latin American Spanish.
	freyaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lupita"}

	// freyaRussianName represents Freya's name in Russian.
	freyaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фрея"}

	// freyaSimplifiedChineseName represents Freya's name in Simplified Chinese.
	freyaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "冰冰"}

	// freyaSpanishName represents Freya's name in Spanish.
	freyaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lupita"}

	// freyaTraditionalChineseName represents Freya's name in Traditional Chinese.
	freyaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "冰冰"}
)

var (
	// freyaName represents Freya's name in different languages.
	freyaName = nook.Languages{
		language.AmericanEnglish:      freyaAmericanEnglishName,
		language.CanadianFrench:       freyaCanadianFrenchName,
		language.Dutch:                freyaDutchName,
		language.French:               freyaFrenchName,
		language.German:               freyaGermanName,
		language.Italian:              freyaItalianName,
		language.Japanese:             freyaJapaneseName,
		language.Korean:               freyaKoreanName,
		language.LatinAmericanSpanish: freyaLatinAmericanSpanishName,
		language.Russian:              freyaRussianName,
		language.SimplifiedChinese:    freyaSimplifiedChineseName,
		language.Spanish:              freyaSpanishName,
		language.TraditionalChinese:   freyaTraditionalChineseName}
)

var (
	// freyaCharacter represents Freya's character information.
	freyaCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: freyaBirthday,
		Code:     freyaCode,
		Key:      character.Freya,
		Gender:   gender.Female,
		Name:     freyaName,
		Special:  false}
)

var (
	// freyaAmericanEnglishPhrase represents Freya's phrase in American English.
	freyaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "uff da"}

	// freyaCanadianFrenchPhrase represents Freya's phrase in Canadian French.
	freyaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon loulou"}

	// freyaDutchPhrase represents Freya's phrase in Dutch.
	freyaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "fjord"}

	// freyaFrenchPhrase represents Freya's phrase in French.
	freyaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon loulou"}

	// freyaGermanPhrase represents Freya's phrase in German.
	freyaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ähmmm"}

	// freyaItalianPhrase represents Freya's phrase in Italian.
	freyaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uuuff"}

	// freyaJapanesePhrase represents Freya's phrase in Japanese.
	freyaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのだわ"}

	// freyaKoreanPhrase represents Freya's phrase in Korean.
	freyaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "유행이야"}

	// freyaLatinAmericanSpanishPhrase represents Freya's phrase in Latin American Spanish.
	freyaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "auuuff"}

	// freyaRussianPhrase represents Freya's phrase in Russian.
	freyaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрусть"}

	// freyaSimplifiedChinesePhrase represents Freya's phrase in Simplified Chinese.
	freyaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就是说哇"}

	// freyaSpanishPhrase represents Freya's phrase in Spanish.
	freyaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "auuuff"}

	// freyaTraditionalChinesePhrase represents Freya's phrase in Traditional Chinese.
	freyaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就是說哇"}
)

var (
	// freyaPhrase represents Freya's phrase in different languages.
	freyaPhrase = nook.Languages{
		language.AmericanEnglish:      freyaAmericanEnglishPhrase,
		language.CanadianFrench:       freyaCanadianFrenchPhrase,
		language.Dutch:                freyaDutchPhrase,
		language.French:               freyaFrenchPhrase,
		language.German:               freyaGermanPhrase,
		language.Italian:              freyaItalianPhrase,
		language.Japanese:             freyaJapanesePhrase,
		language.Korean:               freyaKoreanPhrase,
		language.LatinAmericanSpanish: freyaLatinAmericanSpanishPhrase,
		language.Russian:              freyaRussianPhrase,
		language.SimplifiedChinese:    freyaSimplifiedChinesePhrase,
		language.Spanish:              freyaSpanishPhrase,
		language.TraditionalChinese:   freyaTraditionalChinesePhrase}
)

var (
	// Freya represents the character Freya with her complete information.
	Freya = nook.Villager{
		Character:   freyaCharacter,
		Personality: personality.Snooty,
		Phrase:      freyaPhrase}
)
