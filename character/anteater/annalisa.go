package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// annalisaBirthday represents Annalisa's birthday (February 6th).
var (
	annalisaBirthday = nook.Birthday{
		Day:   6,
		Month: time.February}
)

// annalisaCode represents Annalisa's unique code ("ant08").
var (
	annalisaCode = nook.Code{
		Value: "ant08"}
)

// Different names for Annalisa in various languages.
var (
	annalisaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Annalisa"}

	annalisaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Roberta"}

	annalisaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Annalisa"}

	annalisaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Roberta"}

	annalisaGermanName = nook.Name{
		Language: language.German,
		Value:    "Annalena"}

	annalisaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Anita"}

	annalisaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みやび"}

	annalisaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "설백"}

	annalisaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Alba"}

	annalisaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аннализа"}

	annalisaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小雅"}

	annalisaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Alba"}

	annalisaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小雅"}
)

// annalisaName represents Annalisa's name in different languages.
var (
	annalisaName = nook.Languages{
		language.AmericanEnglish:      annalisaAmericanEnglishName,
		language.CanadianFrench:       annalisaCanadianFrenchName,
		language.Dutch:                annalisaDutchName,
		language.French:               annalisaFrenchName,
		language.German:               annalisaGermanName,
		language.Italian:              annalisaItalianName,
		language.Japanese:             annalisaJapaneseName,
		language.Korean:               annalisaKoreanName,
		language.LatinAmericanSpanish: annalisaLatinAmericanSpanishName,
		language.Russian:              annalisaRussianName,
		language.SimplifiedChinese:    annalisaSimplifiedChineseName,
		language.Spanish:              annalisaSpanishName,
		language.TraditionalChinese:   annalisaTraditionalChineseName}
)

// annalisaCharacter represents Annalisa's character information.
var (
	annalisaCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: annalisaBirthday,
		Code:     annalisaCode,
		Key:      character.Annalisa,
		Gender:   gender.Female,
		Name:     annalisaName,
		Special:  false}
)

// Different phrases spoken by Annalisa in various languages.
var (
	annalisaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "gumdrop"}

	annalisaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tap tap"}

	annalisaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "dropje"}

	annalisaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tap tap"}

	annalisaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnüffel"}

	annalisaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "chic"}

	annalisaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "をかし"}

	annalisaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하양하양"}

	annalisaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "snifffi"}

	annalisaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "жуй-жуй"}

	annalisaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "优雅"}

	annalisaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "nubecillas"}

	annalisaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "優雅"}
)

// annalisaPhrase represents Annalisa's phrases in different languages.
var (
	annalisaPhrase = nook.Languages{
		language.AmericanEnglish:      annalisaAmericanEnglishPhrase,
		language.CanadianFrench:       annalisaCanadianFrenchPhrase,
		language.Dutch:                annalisaDutchPhrase,
		language.French:               annalisaFrenchPhrase,
		language.German:               annalisaGermanPhrase,
		language.Italian:              annalisaItalianPhrase,
		language.Japanese:             annalisaJapanesePhrase,
		language.Korean:               annalisaKoreanPhrase,
		language.LatinAmericanSpanish: annalisaLatinAmericanSpanishPhrase,
		language.Russian:              annalisaRussianPhrase,
		language.SimplifiedChinese:    annalisaSimplifiedChinesePhrase,
		language.Spanish:              annalisaSpanishPhrase,
		language.TraditionalChinese:   annalisaTraditionalChinesePhrase}
)

// Annalisa represents the character Annalisa with her complete information.
var (
	Annalisa = nook.Villager{
		Character:   annalisaCharacter,
		Personality: personality.Normal,
		Phrase:      annalisaPhrase}
)
