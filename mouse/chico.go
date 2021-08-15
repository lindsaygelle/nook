package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	chicoBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	chicoCode = nook.Code{
		Value: ""}
)

var (
	chicoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chico"}

	chicoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	chicoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	chicoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rapido"}

	chicoGermanName = nook.Name{
		Language: language.German,
		Value:    "Pablo"}

	chicoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ciro"}

	chicoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チューボー"}

	chicoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	chicoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	chicoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	chicoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吱吱"}

	chicoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chico"}

	chicoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	chicoName = nook.Languages{
		language.AmericanEnglish:      chicoAmericanEnglishName,
		language.CanadianFrench:       chicoCanadianFrenchName,
		language.Dutch:                chicoDutchName,
		language.French:               chicoFrenchName,
		language.German:               chicoGermanName,
		language.Italian:              chicoItalianName,
		language.Japanese:             chicoJapaneseName,
		language.Korean:               chicoKoreanName,
		language.LatinAmericanSpanish: chicoLatinAmericanSpanishName,
		language.Russian:              chicoRussianName,
		language.SimplifiedChinese:    chicoSimplifiedChineseName,
		language.Spanish:              chicoSpanishName,
		language.TraditionalChinese:   chicoTraditionalChineseName}
)

var (
	chicoCharacter = nook.Character{
		Animal:   Mouse,
		Birthday: chicoBirthday,
		Code:     chicoCode,
		Gender:   nook.Male,
		Name:     chicoName}
)

var (
	chicoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	chicoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	chicoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	chicoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ma croûte"}

	chicoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "griiins"}

	chicoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "caciotta"}

	chicoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ナリ"}

	chicoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	chicoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	chicoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	chicoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哗哗"}

	chicoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "quesito"}

	chicoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	chicoPhrase = nook.Languages{
		language.AmericanEnglish:      chicoAmericanEnglishName,
		language.CanadianFrench:       chicoCanadianFrenchName,
		language.Dutch:                chicoDutchName,
		language.French:               chicoFrenchName,
		language.German:               chicoGermanName,
		language.Italian:              chicoItalianName,
		language.Japanese:             chicoJapaneseName,
		language.Korean:               chicoKoreanName,
		language.LatinAmericanSpanish: chicoLatinAmericanSpanishName,
		language.Russian:              chicoRussianName,
		language.SimplifiedChinese:    chicoSimplifiedChineseName,
		language.Spanish:              chicoSpanishName,
		language.TraditionalChinese:   chicoTraditionalChineseName}
)

var (
	Chico = nook.Villager{
		Character:   chicoCharacter,
		Personality: nook.Lazy,
		Phrase:      chicoPhrase}
)
