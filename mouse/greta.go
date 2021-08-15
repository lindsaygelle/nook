package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	gretaBirthday = nook.Birthday{
		Day:   5,
		Month: time.September}
)

var (
	gretaCode = nook.Code{
		Value: "mus16"}
)

var (
	gretaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Greta"}

	gretaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gretaokka"}

	gretaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gretapiepzak"}

	gretaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gretatéma"}

	gretaGermanName = nook.Name{
		Language: language.German,
		Value:    "Gretelnagnag"}

	gretaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gretarattaplan"}

	gretaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ふくこおほほ"}

	gretaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "복자오호호"}

	gretaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jimenañiñi-ñiá"}

	gretaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гретаписк"}

	gretaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "如意呵呵呵"}

	gretaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jimenabocadito"}

	gretaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "如意呵呵呵"}
)

var (
	gretaName = nook.Languages{
		language.AmericanEnglish:      gretaAmericanEnglishName,
		language.CanadianFrench:       gretaCanadianFrenchName,
		language.Dutch:                gretaDutchName,
		language.French:               gretaFrenchName,
		language.German:               gretaGermanName,
		language.Italian:              gretaItalianName,
		language.Japanese:             gretaJapaneseName,
		language.Korean:               gretaKoreanName,
		language.LatinAmericanSpanish: gretaLatinAmericanSpanishName,
		language.Russian:              gretaRussianName,
		language.SimplifiedChinese:    gretaSimplifiedChineseName,
		language.Spanish:              gretaSpanishName,
		language.TraditionalChinese:   gretaTraditionalChineseName}
)

var (
	gretaCharacter = nook.Character{
		Animal:   Mouse,
		Birthday: gretaBirthday,
		Code:     gretaCode,
		Gender:   nook.Female,
		Name:     gretaName}
)

var (
	gretaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yelp"}

	gretaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "okka"}

	gretaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "piepzak"}

	gretaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "téma"}

	gretaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nagnag"}

	gretaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "rattaplan"}

	gretaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "おほほ"}

	gretaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오호호"}

	gretaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñiñi-ñiá"}

	gretaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "писк"}

	gretaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呵呵呵"}

	gretaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bocadito"}

	gretaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呵呵呵"}
)

var (
	gretaPhrase = nook.Languages{
		language.AmericanEnglish:      gretaAmericanEnglishName,
		language.CanadianFrench:       gretaCanadianFrenchName,
		language.Dutch:                gretaDutchName,
		language.French:               gretaFrenchName,
		language.German:               gretaGermanName,
		language.Italian:              gretaItalianName,
		language.Japanese:             gretaJapaneseName,
		language.Korean:               gretaKoreanName,
		language.LatinAmericanSpanish: gretaLatinAmericanSpanishName,
		language.Russian:              gretaRussianName,
		language.SimplifiedChinese:    gretaSimplifiedChineseName,
		language.Spanish:              gretaSpanishName,
		language.TraditionalChinese:   gretaTraditionalChineseName}
)

var (
	Greta = nook.Villager{
		Character:   gretaCharacter,
		Personality: nook.Snooty,
		Phrase:      gretaPhrase}
)
