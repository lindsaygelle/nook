package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	wartjrBirthday = nook.Birthday{
		Day:   21,
		Month: time.August}
)

var (
	wartjrCode = nook.Code{
		Value: "flg05"}
)

var (
	wartjrAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wart Jr."}

	wartjrCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Crakos"}

	wartjrDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wart Jr."}

	wartjrFrenchName = nook.Name{
		Language: language.French,
		Value:    "Crakos"}

	wartjrGermanName = nook.Name{
		Language: language.German,
		Value:    "Warzi"}

	wartjrItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Porro"}

	wartjrJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サム"}

	wartjrKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "샘"}

	wartjrLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Saponcio"}

	wartjrRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Варт-мл."}

	wartjrSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "山姆"}

	wartjrSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Saponcio"}

	wartjrTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "山姆"}
)

var (
	wartjrName = nook.Languages{
		language.AmericanEnglish:      wartjrAmericanEnglishName,
		language.CanadianFrench:       wartjrCanadianFrenchName,
		language.Dutch:                wartjrDutchName,
		language.French:               wartjrFrenchName,
		language.German:               wartjrGermanName,
		language.Italian:              wartjrItalianName,
		language.Japanese:             wartjrJapaneseName,
		language.Korean:               wartjrKoreanName,
		language.LatinAmericanSpanish: wartjrLatinAmericanSpanishName,
		language.Russian:              wartjrRussianName,
		language.SimplifiedChinese:    wartjrSimplifiedChineseName,
		language.Spanish:              wartjrSpanishName,
		language.TraditionalChinese:   wartjrTraditionalChineseName}
)

var (
	wartjrCharacter = nook.Character{
		Animal:   Frog,
		Birthday: wartjrBirthday,
		Code:     wartjrCode,
		Gender:   nook.Male,
		Name:     wartjrName}
)

var (
	wartjrAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grr-ribbit"}

	wartjrCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "croa-croa"}

	wartjrDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brulkikker"}

	wartjrFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croa-croa"}

	wartjrGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grupap"}

	wartjrItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cracragnam"}

	wartjrJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だぎゃ"}

	wartjrKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "므흣"}

	wartjrLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grroac"}

	wartjrRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ква-ква"}

	wartjrSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "山山"}

	wartjrSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grroac"}

	wartjrTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "山山"}
)

var (
	wartjrPhrase = nook.Languages{
		language.AmericanEnglish:      wartjrAmericanEnglishName,
		language.CanadianFrench:       wartjrCanadianFrenchName,
		language.Dutch:                wartjrDutchName,
		language.French:               wartjrFrenchName,
		language.German:               wartjrGermanName,
		language.Italian:              wartjrItalianName,
		language.Japanese:             wartjrJapaneseName,
		language.Korean:               wartjrKoreanName,
		language.LatinAmericanSpanish: wartjrLatinAmericanSpanishName,
		language.Russian:              wartjrRussianName,
		language.SimplifiedChinese:    wartjrSimplifiedChineseName,
		language.Spanish:              wartjrSpanishName,
		language.TraditionalChinese:   wartjrTraditionalChineseName}
)

var (
	WartJr = nook.Villager{
		Character:   wartjrCharacter,
		Personality: nook.Cranky,
		Phrase:      wartjrPhrase}
)
