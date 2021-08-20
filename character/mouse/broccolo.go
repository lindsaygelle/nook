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
	broccoloBirthday = nook.Birthday{
		Day:   30,
		Month: time.June}
)

var (
	broccoloCode = nook.Code{
		Value: "mus12"}
)

var (
	broccoloAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Broccolo"}

	broccoloCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Steven"}

	broccoloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Broccolo"}

	broccoloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Steven"}

	broccoloGermanName = nook.Name{
		Language: language.German,
		Value:    "Fritzi"}

	broccoloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Franco"}

	broccoloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブロッコリー"}

	broccoloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "브로콜리"}

	broccoloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brócoli"}

	broccoloRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Брокколо"}

	broccoloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茜蓝"}

	broccoloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brócoli"}

	broccoloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "茜藍"}
)

var (
	broccoloName = nook.Languages{
		language.AmericanEnglish:      broccoloAmericanEnglishName,
		language.CanadianFrench:       broccoloCanadianFrenchName,
		language.Dutch:                broccoloDutchName,
		language.French:               broccoloFrenchName,
		language.German:               broccoloGermanName,
		language.Italian:              broccoloItalianName,
		language.Japanese:             broccoloJapaneseName,
		language.Korean:               broccoloKoreanName,
		language.LatinAmericanSpanish: broccoloLatinAmericanSpanishName,
		language.Russian:              broccoloRussianName,
		language.SimplifiedChinese:    broccoloSimplifiedChineseName,
		language.Spanish:              broccoloSpanishName,
		language.TraditionalChinese:   broccoloTraditionalChineseName}
)

var (
	broccoloCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: broccoloBirthday,
		Code:     broccoloCode,
		Key:      character.Broccolo,
		Gender:   gender.Male,
		Name:     broccoloName,
		Special:  false}
)

var (
	broccoloAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "eat it"}

	broccoloCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "flim-flam"}

	broccoloDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "smakelijk"}

	broccoloFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "flim-flam"}

	broccoloGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knabba"}

	broccoloItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squikko"}

	broccoloJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ピコ"}

	broccoloKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "앗차"}

	broccoloLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "flinflán"}

	broccoloRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ням-ням"}

	broccoloSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "微微"}

	broccoloSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "dientes"}

	broccoloTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "微微"}
)

var (
	broccoloPhrase = nook.Languages{
		language.AmericanEnglish:      broccoloAmericanEnglishPhrase,
		language.CanadianFrench:       broccoloCanadianFrenchPhrase,
		language.Dutch:                broccoloDutchPhrase,
		language.French:               broccoloFrenchPhrase,
		language.German:               broccoloGermanPhrase,
		language.Italian:              broccoloItalianPhrase,
		language.Japanese:             broccoloJapanesePhrase,
		language.Korean:               broccoloKoreanPhrase,
		language.LatinAmericanSpanish: broccoloLatinAmericanSpanishPhrase,
		language.Russian:              broccoloRussianPhrase,
		language.SimplifiedChinese:    broccoloSimplifiedChinesePhrase,
		language.Spanish:              broccoloSpanishPhrase,
		language.TraditionalChinese:   broccoloTraditionalChinesePhrase}
)

var (
	Broccolo = nook.Villager{
		Character:   broccoloCharacter,
		Personality: personality.Lazy,
		Phrase:      broccoloPhrase}
)
