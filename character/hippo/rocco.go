package hippo

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
		Value:    "Bebel"}

	roccoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rocco"}

	roccoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bebel"}

	roccoGermanName = nook.Name{
		Language: language.German,
		Value:    "Richi"}

	roccoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rocco"}

	roccoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゴンザレス"}

	roccoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곤잘레스"}

	roccoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Roco"}

	roccoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рокко"}

	roccoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "孔在来"}

	roccoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Roco"}

	roccoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "孔在來"}
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
		Animal:   animal.Hippo,
		Birthday: roccoBirthday,
		Code:     roccoCode,
		Key:      character.Rocco,
		Gender:   gender.Male,
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
		Personality: personality.Cranky,
		Phrase:      roccoPhrase}
)
