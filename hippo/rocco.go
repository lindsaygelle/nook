package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	roccoBirthday = nook.Birthday{
		Day:   18,
		Month: time.August}
)

var (
	roccoCode = nook.Code{
		Value: "hip00"}
)

var (
	roccoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rocco"}

	roccoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bebelbabi"}

	roccoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Roccohippo"}

	roccoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bebelbabi"}

	roccoGermanName = nook.Name{
		Language: language.German,
		Value:    "Richipfffft"}

	roccoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Roccoipper"}

	roccoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゴンザレスだぎゃー"}

	roccoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곤잘레스그러마"}

	roccoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rocohiponó"}

	roccoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Роккогиппи"}

	roccoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "孔在来觉得是啦"}

	roccoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rocobatracio"}

	roccoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "孔在來覺得是啦"}
)

var (
	roccoName = nook.Languages{
		language.AmericanEnglish:      roccoAmericanEnglishName,
		language.CanadianFrench:       roccoCanadianFrenchName,
		language.Dutch:                roccoDutchName,
		language.French:               roccoFrenchName,
		language.German:               roccoGermanName,
		language.Italian:              roccoItalianName,
		language.Japanese:             roccoJapaneseName,
		language.Korean:               roccoKoreanName,
		language.LatinAmericanSpanish: roccoLatinAmericanSpanishName,
		language.Russian:              roccoRussianName,
		language.SimplifiedChinese:    roccoSimplifiedChineseName,
		language.Spanish:              roccoSpanishName,
		language.TraditionalChinese:   roccoTraditionalChineseName}
)

var (
	roccoCharacter = nook.Character{
		Animal:   Hippo,
		Birthday: roccoBirthday,
		Code:     roccoCode,
		Gender:   nook.Male,
		Name:     roccoName}
)

var (
	roccoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hippie"}

	roccoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "babi"}

	roccoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hippo"}

	roccoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "babi"}

	roccoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pfffft"}

	roccoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ipper"}

	roccoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だぎゃー"}

	roccoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그러마"}

	roccoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hiponó"}

	roccoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гиппи"}

	roccoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "觉得是啦"}

	roccoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "batracio"}

	roccoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "覺得是啦"}
)

var (
	roccoPhrase = nook.Languages{
		language.AmericanEnglish:      roccoAmericanEnglishName,
		language.CanadianFrench:       roccoCanadianFrenchName,
		language.Dutch:                roccoDutchName,
		language.French:               roccoFrenchName,
		language.German:               roccoGermanName,
		language.Italian:              roccoItalianName,
		language.Japanese:             roccoJapaneseName,
		language.Korean:               roccoKoreanName,
		language.LatinAmericanSpanish: roccoLatinAmericanSpanishName,
		language.Russian:              roccoRussianName,
		language.SimplifiedChinese:    roccoSimplifiedChineseName,
		language.Spanish:              roccoSpanishName,
		language.TraditionalChinese:   roccoTraditionalChineseName}
)

var (
	Rocco = nook.Villager{
		Character:   roccoCharacter,
		Personality: nook.Cranky,
		Phrase:      roccoPhrase}
)
