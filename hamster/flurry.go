package hamster

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Emmaabajoujou"}

	flurryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Flurrydonsbal"}

	flurryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Emmaabajoujou"}

	flurryGermanName = nook.Name{
		Language: language.German,
		Value:    "Emiliemillimilli"}

	flurryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Meringafondue"}

	flurryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゆきみなのです"}

	flurryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "뽀야미뽀드득"}

	flurryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lluvianomnom"}

	flurryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фларрипушистик"}

	flurrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪美对啊"}

	flurrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lluvianomnom"}

	flurryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪美對啊"}
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
		Animal:   Hamster,
		Birthday: flurryBirthday,
		Code:     flurryCode,
		Gender:   nook.Female,
		Name:     flurryName}
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
	Flurry = nook.Villager{
		Character:   flurryCharacter,
		Personality: nook.Normal,
		Phrase:      flurryPhrase}
)
