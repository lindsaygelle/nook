package kangaroo

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
	// sylviaBirthday represents sylvia birthday.
	sylviaBirthday = nook.Birthday{
		Day:   3,
		Month: time.May}
)

var (
	// sylviaCode represents sylvia code.
	sylviaCode = nook.Code{
		Value: "kgr06"}
)

var (
	// sylviaAmericanEnglishName represents sylvia american english name.
	sylviaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sylvia"}

	// sylviaCanadianFrenchName represents sylvia canadian french name.
	sylviaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Madsi"}

	// sylviaDutchName represents sylvia dutch name.
	sylviaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sylvia"}

	// sylviaFrenchName represents sylvia french name.
	sylviaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Madsi"}

	// sylviaGermanName represents sylvia german name.
	sylviaGermanName = nook.Name{
		Language: language.German,
		Value:    "Sylvia"}

	// sylviaItalianName represents sylvia italian name.
	sylviaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Silvia"}

	// sylviaJapaneseName represents sylvia japanese name.
	sylviaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シルビア"}

	// sylviaKoreanName represents sylvia korean name.
	sylviaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "실비아"}

	// sylviaLatinAmericanSpanishName represents sylvia latin american spanish name.
	sylviaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cándida"}

	// sylviaRussianName represents sylvia russian name.
	sylviaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сильвия"}

	// sylviaSimplifiedChineseName represents sylvia simplified chinese name.
	sylviaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "希薇亚"}

	// sylviaSpanishName represents sylvia spanish name.
	sylviaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cándida"}

	// sylviaTraditionalChineseName represents sylvia traditional chinese name.
	sylviaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "希薇亞"}
)

var (
	// sylviaName represents sylvia name.
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
	// sylviaCharacter represents sylvia character.
	sylviaCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: sylviaBirthday,
		Code:     sylviaCode,
		Key:      character.Sylvia,
		Gender:   gender.Female,
		Name:     sylviaName,
		Special:  false}
)

var (
	// sylviaAmericanEnglishPhrase represents sylvia american english phrase.
	sylviaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "joey"}

	// sylviaCanadianFrenchPhrase represents sylvia canadian french phrase.
	sylviaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "d'la balle"}

	// sylviaDutchPhrase represents sylvia dutch phrase.
	sylviaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boing"}

	// sylviaFrenchPhrase represents sylvia french phrase.
	sylviaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "d'la balle"}

	// sylviaGermanPhrase represents sylvia german phrase.
	sylviaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flottflott"}

	// sylviaItalianPhrase represents sylvia italian phrase.
	sylviaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "saltolà"}

	// sylviaJapanesePhrase represents sylvia japanese phrase.
	sylviaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よねー"}

	// sylviaKoreanPhrase represents sylvia korean phrase.
	sylviaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "이자슥"}

	// sylviaLatinAmericanSpanishPhrase represents sylvia latin american spanish phrase.
	sylviaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "do-doing"}

	// sylviaRussianPhrase represents sylvia russian phrase.
	sylviaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "прыг"}

	// sylviaSimplifiedChinesePhrase represents sylvia simplified chinese phrase.
	sylviaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "对耶"}

	// sylviaSpanishPhrase represents sylvia spanish phrase.
	sylviaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "púgil"}

	// sylviaTraditionalChinesePhrase represents sylvia traditional chinese phrase.
	sylviaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "對耶"}
)

var (
	// sylviaPhrase represents sylvia phrase.
	sylviaPhrase = nook.Languages{
		language.AmericanEnglish:      sylviaAmericanEnglishPhrase,
		language.CanadianFrench:       sylviaCanadianFrenchPhrase,
		language.Dutch:                sylviaDutchPhrase,
		language.French:               sylviaFrenchPhrase,
		language.German:               sylviaGermanPhrase,
		language.Italian:              sylviaItalianPhrase,
		language.Japanese:             sylviaJapanesePhrase,
		language.Korean:               sylviaKoreanPhrase,
		language.LatinAmericanSpanish: sylviaLatinAmericanSpanishPhrase,
		language.Russian:              sylviaRussianPhrase,
		language.SimplifiedChinese:    sylviaSimplifiedChinesePhrase,
		language.Spanish:              sylviaSpanishPhrase,
		language.TraditionalChinese:   sylviaTraditionalChinesePhrase}
)

var (
	// Sylvia represents sylvia.
	Sylvia = nook.Villager{
		Character:   sylviaCharacter,
		Personality: personality.BigSister,
		Phrase:      sylviaPhrase}
)
