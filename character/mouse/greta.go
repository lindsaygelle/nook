package mouse

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
		Value:    "Greta"}

	gretaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Greta"}

	gretaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Greta"}

	gretaGermanName = nook.Name{
		Language: language.German,
		Value:    "Gretel"}

	gretaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Greta"}

	gretaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ふくこ"}

	gretaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "복자"}

	gretaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jimena"}

	gretaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Грета"}

	gretaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "如意"}

	gretaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jimena"}

	gretaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "如意"}
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
		Animal:   animal.Mouse,
		Birthday: gretaBirthday,
		Code:     gretaCode,
		Key:      character.Greta,
		Gender:   gender.Female,
		Name:     gretaName,
		Special:  false}
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
		language.AmericanEnglish:      gretaAmericanEnglishPhrase,
		language.CanadianFrench:       gretaCanadianFrenchPhrase,
		language.Dutch:                gretaDutchPhrase,
		language.French:               gretaFrenchPhrase,
		language.German:               gretaGermanPhrase,
		language.Italian:              gretaItalianPhrase,
		language.Japanese:             gretaJapanesePhrase,
		language.Korean:               gretaKoreanPhrase,
		language.LatinAmericanSpanish: gretaLatinAmericanSpanishPhrase,
		language.Russian:              gretaRussianPhrase,
		language.SimplifiedChinese:    gretaSimplifiedChinesePhrase,
		language.Spanish:              gretaSpanishPhrase,
		language.TraditionalChinese:   gretaTraditionalChinesePhrase}
)

var (
	Greta = nook.Villager{
		Character:   gretaCharacter,
		Personality: personality.Snooty,
		Phrase:      gretaPhrase}
)
