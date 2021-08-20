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
	merengueBirthday = nook.Birthday{
		Day:   19,
		Month: time.March}
)

var (
	merengueCode = nook.Code{
		Value: "rhn07"}
)

var (
	merengueAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Merengue"}

	merengueCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Patty"}

	merengueDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Merengue"}

	merengueFrenchName = nook.Name{
		Language: language.French,
		Value:    "Patty"}

	merengueGermanName = nook.Name{
		Language: language.German,
		Value:    "Maria"}

	merengueItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Irina"}

	merengueJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パティ"}

	merengueKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스트로베리"}

	merengueLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Natura"}

	merengueRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Меренг"}

	merengueSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "草莓"}

	merengueSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Natura"}

	merengueTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "草莓"}
)

var (
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
	merengueAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "shortcake"}

	merengueCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "confiture"}

	merengueDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schuimpje"}

	merengueFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ma fraise"}

	merengueGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "trarapp"}

	merengueItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "rien"}

	merengueJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でシュガ"}

	merengueKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "베리베리"}

	merengueLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fruti"}

	merengueRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кексик"}

	merengueSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "糖糖"}

	merengueSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fresitas"}

	merengueTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "糖糖"}
)

var (
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
	Merengue = nook.Villager{
		Character:   merengueCharacter,
		Personality: personality.Normal,
		Phrase:      merenguePhrase}
)
