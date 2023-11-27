package rhinoceros

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
	// merengueBirthday represents Merengue's birthday.
	merengueBirthday = nook.Birthday{
		Day:   19,
		Month: time.March}
)

var (
	// merengueCode represents Merengue's unique code.
	merengueCode = nook.Code{
		Value: "rhn07"}
)

var (
	// merengueAmericanEnglishName represents Merengue's name in American English.
	merengueAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Merengue"}

	// merengueCanadianFrenchName represents Merengue's name in Canadian French.
	merengueCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Patty"}

	// merengueDutchName represents Merengue's name in Dutch.
	merengueDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Merengue"}

	// merengueFrenchName represents Merengue's name in French.
	merengueFrenchName = nook.Name{
		Language: language.French,
		Value:    "Patty"}

	// merengueGermanName represents Merengue's name in German.
	merengueGermanName = nook.Name{
		Language: language.German,
		Value:    "Maria"}

	// merengueItalianName represents Merengue's name in Italian.
	merengueItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Irina"}

	// merengueJapaneseName represents Merengue's name in Japanese.
	merengueJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パティ"}

	// merengueKoreanName represents Merengue's name in Korean.
	merengueKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스트로베리"}

	// merengueLatinAmericanSpanishName represents Merengue's name in Latin American Spanish.
	merengueLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Natura"}

	// merengueRussianName represents Merengue's name in Russian.
	merengueRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Меренг"}

	// merengueSimplifiedChineseName represents Merengue's name in Simplified Chinese.
	merengueSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "草莓"}

	// merengueSpanishName represents Merengue's name in Spanish.
	merengueSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Natura"}

	// merengueTraditionalChineseName represents Merengue's name in Traditional Chinese.
	merengueTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "草莓"}
)

var (
	// merengueName represents Merengue's name in different languages.
	merengueName = nook.Languages{
		language.AmericanEnglish:      merengueAmericanEnglishName,
		language.CanadianFrench:       merengueCanadianFrenchName,
		language.Dutch:                merengueDutchName,
		language.French:               merengueFrenchName,
		language.German:               merengueGermanName,
		language.Italian:              merengueItalianName,
		language.Japanese:             merengueJapaneseName,
		language.Korean:               merengueKoreanName,
		language.LatinAmericanSpanish: merengueLatinAmericanSpanishName,
		language.Russian:              merengueRussianName,
		language.SimplifiedChinese:    merengueSimplifiedChineseName,
		language.Spanish:              merengueSpanishName,
		language.TraditionalChinese:   merengueTraditionalChineseName}
)

var (
	// merengueCharacter represents Merengue's character information.
	merengueCharacter = nook.Character{
		Animal:   animal.Rhinoceros,
		Birthday: merengueBirthday,
		Code:     merengueCode,
		Key:      character.Merengue,
		Gender:   gender.Female,
		Name:     merengueName,
		Special:  false}
)

var (
	// merengueAmericanEnglishPhrase represents Merengue's phrase in American English.
	merengueAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "shortcake"}

	// merengueCanadianFrenchPhrase represents Merengue's phrase in Canadian French.
	merengueCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "confiture"}

	// merengueDutchPhrase represents Merengue's phrase in Dutch.
	merengueDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schuimpje"}

	// merengueFrenchPhrase represents Merengue's phrase in French.
	merengueFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ma fraise"}

	// merengueGermanPhrase represents Merengue's phrase in German.
	merengueGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "trarapp"}

	// merengueItalianPhrase represents Merengue's phrase in Italian.
	merengueItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "rien"}

	// merengueJapanesePhrase represents Merengue's phrase in Japanese.
	merengueJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でシュガ"}

	// merengueKoreanPhrase represents Merengue's phrase in Korean.
	merengueKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "베리베리"}

	// merengueLatinAmericanSpanishPhrase represents Merengue's phrase in Latin American Spanish.
	merengueLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fruti"}

	// merengueRussianPhrase represents Merengue's phrase in Russian.
	merengueRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кексик"}

	// merengueSimplifiedChinesePhrase represents Merengue's phrase in Simplified Chinese.
	merengueSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "糖糖"}

	// merengueSpanishPhrase represents Merengue's phrase in Spanish.
	merengueSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fresitas"}

	// merengueTraditionalChinesePhrase represents Merengue's phrase in Traditional Chinese.
	merengueTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "糖糖"}
)

var (
	// merenguePhrase represents Merengue's phrases in different languages.
	merenguePhrase = nook.Languages{
		language.AmericanEnglish:      merengueAmericanEnglishPhrase,
		language.CanadianFrench:       merengueCanadianFrenchPhrase,
		language.Dutch:                merengueDutchPhrase,
		language.French:               merengueFrenchPhrase,
		language.German:               merengueGermanPhrase,
		language.Italian:              merengueItalianPhrase,
		language.Japanese:             merengueJapanesePhrase,
		language.Korean:               merengueKoreanPhrase,
		language.LatinAmericanSpanish: merengueLatinAmericanSpanishPhrase,
		language.Russian:              merengueRussianPhrase,
		language.SimplifiedChinese:    merengueSimplifiedChinesePhrase,
		language.Spanish:              merengueSpanishPhrase,
		language.TraditionalChinese:   merengueTraditionalChinesePhrase}
)

var (
	// Merengue represents the character Merengue with her complete information.
	Merengue = nook.Villager{
		Character:   merengueCharacter,
		Personality: personality.Normal,
		Phrase:      merenguePhrase}
)
