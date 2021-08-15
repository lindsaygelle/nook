package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	ursalaBirthday = nook.Birthday{
		Day:   16,
		Month: time.January}
)

var (
	ursalaCode = nook.Code{
		Value: "bea08"}
)

var (
	ursalaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ursala"}

	ursalaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Oursulagroumpf"}

	ursalaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ursalaoemf"}

	ursalaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Oursulagroumpf"}

	ursalaGermanName = nook.Name{
		Language: language.German,
		Value:    "Ursulabrumbrum"}

	ursalaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ursulagruuunf"}

	ursalaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ネーヤやーねぇ"}

	ursalaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "네이아그라지"}

	ursalaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Úrsulagrumf"}

	ursalaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Урсалагррум"}

	ursalaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "妮雅呀呐"}

	ursalaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Úrsulagrumf"}

	ursalaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "妮雅呀吶"}
)

var (
	ursalaName = nook.Languages{
		language.AmericanEnglish:      ursalaAmericanEnglishName,
		language.CanadianFrench:       ursalaCanadianFrenchName,
		language.Dutch:                ursalaDutchName,
		language.French:               ursalaFrenchName,
		language.German:               ursalaGermanName,
		language.Italian:              ursalaItalianName,
		language.Japanese:             ursalaJapaneseName,
		language.Korean:               ursalaKoreanName,
		language.LatinAmericanSpanish: ursalaLatinAmericanSpanishName,
		language.Russian:              ursalaRussianName,
		language.SimplifiedChinese:    ursalaSimplifiedChineseName,
		language.Spanish:              ursalaSpanishName,
		language.TraditionalChinese:   ursalaTraditionalChineseName}
)

var (
	ursalaCharacter = nook.Character{
		Animal:   Bear,
		Birthday: ursalaBirthday,
		Code:     ursalaCode,
		Gender:   nook.Female,
		Name:     ursalaName}
)

var (
	ursalaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grooomph"}

	ursalaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "groumpf"}

	ursalaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oemf"}

	ursalaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "groumpf"}

	ursalaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brumbrum"}

	ursalaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gruuunf"}

	ursalaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やーねぇ"}

	ursalaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그라지"}

	ursalaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grumf"}

	ursalaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гррум"}

	ursalaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呀呐"}

	ursalaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grumf"}

	ursalaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呀吶"}
)

var (
	ursalaPhrase = nook.Languages{
		language.AmericanEnglish:      ursalaAmericanEnglishName,
		language.CanadianFrench:       ursalaCanadianFrenchName,
		language.Dutch:                ursalaDutchName,
		language.French:               ursalaFrenchName,
		language.German:               ursalaGermanName,
		language.Italian:              ursalaItalianName,
		language.Japanese:             ursalaJapaneseName,
		language.Korean:               ursalaKoreanName,
		language.LatinAmericanSpanish: ursalaLatinAmericanSpanishName,
		language.Russian:              ursalaRussianName,
		language.SimplifiedChinese:    ursalaSimplifiedChineseName,
		language.Spanish:              ursalaSpanishName,
		language.TraditionalChinese:   ursalaTraditionalChineseName}
)

var (
	Ursala = nook.Villager{
		Character:   ursalaCharacter,
		Personality: nook.BigSister,
		Phrase:      ursalaPhrase}
)
