package rhinoceros

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Pattyconfiture"}

	merengueDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Merengueschuimpje"}

	merengueFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pattyma fraise"}

	merengueGermanName = nook.Name{
		Language: language.German,
		Value:    "Mariatrarapp"}

	merengueItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Irinarien"}

	merengueJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パティでシュガ"}

	merengueKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스트로베리베리베리"}

	merengueLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Naturafruti"}

	merengueRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Меренгкексик"}

	merengueSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "草莓糖糖"}

	merengueSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Naturafresitas"}

	merengueTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "草莓糖糖"}
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
		Animal:   Rhinoceros,
		Birthday: merengueBirthday,
		Code:     merengueCode,
		Gender:   nook.Female,
		Name:     merengueName}
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
	Merengue = nook.Villager{
		Character:   merengueCharacter,
		Personality: nook.Normal,
		Phrase:      merenguePhrase}
)
