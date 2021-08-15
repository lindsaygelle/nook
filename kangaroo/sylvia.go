package kangaroo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sylviaBirthday = nook.Birthday{
		Day:   3,
		Month: time.May}
)

var (
	sylviaCode = nook.Code{
		Value: "kgr06"}
)

var (
	sylviaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sylvia"}

	sylviaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Madsi"}

	sylviaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sylvia"}

	sylviaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Madsi"}

	sylviaGermanName = nook.Name{
		Language: language.German,
		Value:    "Sylvia"}

	sylviaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Silvia"}

	sylviaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シルビア"}

	sylviaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "실비아"}

	sylviaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cándida"}

	sylviaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сильвия"}

	sylviaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "希薇亚"}

	sylviaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cándida"}

	sylviaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "希薇亞"}
)

var (
	sylviaName = nook.Languages{
		language.AmericanEnglish:      sylviaAmericanEnglishName,
		language.CanadianFrench:       sylviaCanadianFrenchName,
		language.Dutch:                sylviaDutchName,
		language.French:               sylviaFrenchName,
		language.German:               sylviaGermanName,
		language.Italian:              sylviaItalianName,
		language.Japanese:             sylviaJapaneseName,
		language.Korean:               sylviaKoreanName,
		language.LatinAmericanSpanish: sylviaLatinAmericanSpanishName,
		language.Russian:              sylviaRussianName,
		language.SimplifiedChinese:    sylviaSimplifiedChineseName,
		language.Spanish:              sylviaSpanishName,
		language.TraditionalChinese:   sylviaTraditionalChineseName}
)

var (
	sylviaCharacter = nook.Character{
		Animal:   Kangaroo,
		Birthday: sylviaBirthday,
		Code:     sylviaCode,
		Gender:   nook.Female,
		Name:     sylviaName}
)

var (
	sylviaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "joey"}

	sylviaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "d'la balle"}

	sylviaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boing"}

	sylviaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "d'la balle"}

	sylviaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flottflott"}

	sylviaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "saltolà"}

	sylviaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よねー"}

	sylviaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "이자슥"}

	sylviaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "do-doing"}

	sylviaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "прыг"}

	sylviaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "对耶"}

	sylviaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "púgil"}

	sylviaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "對耶"}
)

var (
	sylviaPhrase = nook.Languages{
		language.AmericanEnglish:      sylviaAmericanEnglishName,
		language.CanadianFrench:       sylviaCanadianFrenchName,
		language.Dutch:                sylviaDutchName,
		language.French:               sylviaFrenchName,
		language.German:               sylviaGermanName,
		language.Italian:              sylviaItalianName,
		language.Japanese:             sylviaJapaneseName,
		language.Korean:               sylviaKoreanName,
		language.LatinAmericanSpanish: sylviaLatinAmericanSpanishName,
		language.Russian:              sylviaRussianName,
		language.SimplifiedChinese:    sylviaSimplifiedChineseName,
		language.Spanish:              sylviaSpanishName,
		language.TraditionalChinese:   sylviaTraditionalChineseName}
)

var (
	Sylvia = nook.Villager{
		Character:   sylviaCharacter,
		Personality: nook.BigSister,
		Phrase:      sylviaPhrase}
)
