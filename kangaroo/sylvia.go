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
		Value:    "Madsid'la balle"}

	sylviaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sylviaboing"}

	sylviaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Madsid'la balle"}

	sylviaGermanName = nook.Name{
		Language: language.German,
		Value:    "Sylviaflottflott"}

	sylviaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Silviasaltolà"}

	sylviaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シルビアよねー"}

	sylviaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "실비아이자슥"}

	sylviaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cándidado-doing"}

	sylviaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сильвияпрыг"}

	sylviaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "希薇亚对耶"}

	sylviaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cándidapúgil"}

	sylviaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "希薇亞對耶"}
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
		Value:    "joeyboing"}

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
