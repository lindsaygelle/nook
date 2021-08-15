package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	croqueBirthday = nook.Birthday{
		Day:   18,
		Month: time.July}
)

var (
	croqueCode = nook.Code{
		Value: "flg17"}
)

var (
	croqueAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Croque"}

	croqueCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Carlosouaoua"}

	croqueDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Croqueja dááág"}

	croqueFrenchName = nook.Name{
		Language: language.French,
		Value:    "Carloscrapoupou"}

	croqueGermanName = nook.Name{
		Language: language.German,
		Value:    "Carloschleck"}

	croqueItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gracidoumpf"}

	croqueJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タイシしからば"}

	croqueKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "투투툴툴"}

	croqueLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ranolfoyastá"}

	croqueRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кроквроде того"}

	croqueSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "太子子曰"}

	croqueSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ranolfoyastá"}

	croqueTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "太子子曰"}
)

var (
	croqueName = nook.Languages{
		language.AmericanEnglish:      croqueAmericanEnglishName,
		language.CanadianFrench:       croqueCanadianFrenchName,
		language.Dutch:                croqueDutchName,
		language.French:               croqueFrenchName,
		language.German:               croqueGermanName,
		language.Italian:              croqueItalianName,
		language.Japanese:             croqueJapaneseName,
		language.Korean:               croqueKoreanName,
		language.LatinAmericanSpanish: croqueLatinAmericanSpanishName,
		language.Russian:              croqueRussianName,
		language.SimplifiedChinese:    croqueSimplifiedChineseName,
		language.Spanish:              croqueSpanishName,
		language.TraditionalChinese:   croqueTraditionalChineseName}
)

var (
	croqueCharacter = nook.Character{
		Animal:   Frog,
		Birthday: croqueBirthday,
		Code:     croqueCode,
		Gender:   nook.Male,
		Name:     croqueName}
)

var (
	croqueAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "as if"}

	croqueCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ouaoua"}

	croqueDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ja dááág"}

	croqueFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "crapoupou"}

	croqueGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schleck"}

	croqueItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "umpf"}

	croqueJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "しからば"}

	croqueKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "툴툴"}

	croqueLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "yastá"}

	croqueRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вроде того"}

	croqueSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "子曰"}

	croqueSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "yastá"}

	croqueTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "子曰"}
)

var (
	croquePhrase = nook.Languages{
		language.AmericanEnglish:      croqueAmericanEnglishName,
		language.CanadianFrench:       croqueCanadianFrenchName,
		language.Dutch:                croqueDutchName,
		language.French:               croqueFrenchName,
		language.German:               croqueGermanName,
		language.Italian:              croqueItalianName,
		language.Japanese:             croqueJapaneseName,
		language.Korean:               croqueKoreanName,
		language.LatinAmericanSpanish: croqueLatinAmericanSpanishName,
		language.Russian:              croqueRussianName,
		language.SimplifiedChinese:    croqueSimplifiedChineseName,
		language.Spanish:              croqueSpanishName,
		language.TraditionalChinese:   croqueTraditionalChineseName}
)

var (
	Croque = nook.Villager{
		Character:   croqueCharacter,
		Personality: nook.Cranky,
		Phrase:      croquePhrase}
)
