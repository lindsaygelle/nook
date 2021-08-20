package hamster

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
	flurryBirthday = nook.Birthday{
		Day:   30,
		Month: time.January}
)

var (
	flurryCode = nook.Code{
		Value: "ham06"}
)

var (
	flurryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flurry"}

	flurryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Emma"}

	flurryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Flurry"}

	flurryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Emma"}

	flurryGermanName = nook.Name{
		Language: language.German,
		Value:    "Emilie"}

	flurryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Meringa"}

	flurryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゆきみ"}

	flurryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뽀야미"}

	flurryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lluvia"}

	flurryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фларри"}

	flurrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪美"}

	flurrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lluvia"}

	flurryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪美"}
)

var (
	flurryName = nook.Languages{
		language.AmericanEnglish:      flurryAmericanEnglishName,
		language.CanadianFrench:       flurryCanadianFrenchName,
		language.Dutch:                flurryDutchName,
		language.French:               flurryFrenchName,
		language.German:               flurryGermanName,
		language.Italian:              flurryItalianName,
		language.Japanese:             flurryJapaneseName,
		language.Korean:               flurryKoreanName,
		language.LatinAmericanSpanish: flurryLatinAmericanSpanishName,
		language.Russian:              flurryRussianName,
		language.SimplifiedChinese:    flurrySimplifiedChineseName,
		language.Spanish:              flurrySpanishName,
		language.TraditionalChinese:   flurryTraditionalChineseName}
)

var (
	flurryCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: flurryBirthday,
		Code:     flurryCode,
		Key:      character.Flurry,
		Gender:   gender.Female,
		Name:     flurryName,
		Special:  false}
)

var (
	flurryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "powderpuff"}

	flurryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "abajoujou"}

	flurryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "donsbal"}

	flurryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "abajoujou"}

	flurryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "millimilli"}

	flurryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fondue"}

	flurryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのです"}

	flurryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뽀드득"}

	flurryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nomnom"}

	flurryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пушистик"}

	flurrySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "对啊"}

	flurrySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "nomnom"}

	flurryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "對啊"}
)

var (
	flurryPhrase = nook.Languages{
		language.AmericanEnglish:      flurryAmericanEnglishPhrase,
		language.CanadianFrench:       flurryCanadianFrenchPhrase,
		language.Dutch:                flurryDutchPhrase,
		language.French:               flurryFrenchPhrase,
		language.German:               flurryGermanPhrase,
		language.Italian:              flurryItalianPhrase,
		language.Japanese:             flurryJapanesePhrase,
		language.Korean:               flurryKoreanPhrase,
		language.LatinAmericanSpanish: flurryLatinAmericanSpanishPhrase,
		language.Russian:              flurryRussianPhrase,
		language.SimplifiedChinese:    flurrySimplifiedChinesePhrase,
		language.Spanish:              flurrySpanishPhrase,
		language.TraditionalChinese:   flurryTraditionalChinesePhrase}
)

var (
	Flurry = nook.Villager{
		Character:   flurryCharacter,
		Personality: personality.Normal,
		Phrase:      flurryPhrase}
)
