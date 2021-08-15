package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	daisyBirthday = nook.Birthday{
		Day:   16,
		Month: time.November}
)

var (
	daisyCode = nook.Code{
		Value: "dog07"}
)

var (
	daisyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Daisy"}

	daisyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Naomie"}

	daisyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Daisy"}

	daisyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Naomie"}

	daisyGermanName = nook.Name{
		Language: language.German,
		Value:    "Doris"}

	daisyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fiorella"}

	daisyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バニラ"}

	daisyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바닐라"}

	daisyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Luisa"}

	daisyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дейзи"}

	daisySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "香草"}

	daisySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Luisa"}

	daisyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "香草"}
)

var (
	daisyName = nook.Languages{
		language.AmericanEnglish:      daisyAmericanEnglishName,
		language.CanadianFrench:       daisyCanadianFrenchName,
		language.Dutch:                daisyDutchName,
		language.French:               daisyFrenchName,
		language.German:               daisyGermanName,
		language.Italian:              daisyItalianName,
		language.Japanese:             daisyJapaneseName,
		language.Korean:               daisyKoreanName,
		language.LatinAmericanSpanish: daisyLatinAmericanSpanishName,
		language.Russian:              daisyRussianName,
		language.SimplifiedChinese:    daisySimplifiedChineseName,
		language.Spanish:              daisySpanishName,
		language.TraditionalChinese:   daisyTraditionalChineseName}
)

var (
	daisyCharacter = nook.Character{
		Animal:   Dog,
		Birthday: daisyBirthday,
		Code:     daisyCode,
		Gender:   nook.Female,
		Name:     daisyName}
)

var (
	daisyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bow-WOW"}

	daisyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon chou"}

	daisyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hazewind"}

	daisyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon chou"}

	daisyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wauwau"}

	daisyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bau WOW"}

	daisyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だよね"}

	daisyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇죠"}

	daisyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guaucito"}

	daisyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тяв-ого"}

	daisySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "对不对"}

	daisySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guaay"}

	daisyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "對不對"}
)

var (
	daisyPhrase = nook.Languages{
		language.AmericanEnglish:      daisyAmericanEnglishName,
		language.CanadianFrench:       daisyCanadianFrenchName,
		language.Dutch:                daisyDutchName,
		language.French:               daisyFrenchName,
		language.German:               daisyGermanName,
		language.Italian:              daisyItalianName,
		language.Japanese:             daisyJapaneseName,
		language.Korean:               daisyKoreanName,
		language.LatinAmericanSpanish: daisyLatinAmericanSpanishName,
		language.Russian:              daisyRussianName,
		language.SimplifiedChinese:    daisySimplifiedChineseName,
		language.Spanish:              daisySpanishName,
		language.TraditionalChinese:   daisyTraditionalChineseName}
)

var (
	Daisy = nook.Villager{
		Character:   daisyCharacter,
		Personality: nook.Normal,
		Phrase:      daisyPhrase}
)
