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
		Value:    "Naomiemon chou"}

	daisyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Daisyhazewind"}

	daisyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Naomiemon chou"}

	daisyGermanName = nook.Name{
		Language: language.German,
		Value:    "Doriswauwau"}

	daisyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fiorellabau WOW"}

	daisyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バニラだよね"}

	daisyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바닐라그렇죠"}

	daisyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Luisaguaucito"}

	daisyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дейзитяв-ого"}

	daisySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "香草对不对"}

	daisySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Luisaguaay"}

	daisyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "香草對不對"}
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
