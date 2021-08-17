package duck

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
	mirandaBirthday = nook.Birthday{
		Day:   23,
		Month: time.April}
)

var (
	mirandaCode = nook.Code{
		Value: "duk12"}
)

var (
	mirandaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Miranda"}

	mirandaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maëllis"}

	mirandaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Miranda"}

	mirandaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maëllis"}

	mirandaGermanName = nook.Name{
		Language: language.German,
		Value:    "Tanya"}

	mirandaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Miranda"}

	mirandaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミランダ"}

	mirandaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미란다"}

	mirandaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Miranda"}

	mirandaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Миранда"}

	mirandaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "米兰达"}

	mirandaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Miranda"}

	mirandaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "米蘭達"}
)

var (
	mirandaName = nook.Languages{
		language.AmericanEnglish:      mirandaAmericanEnglishName,
		language.CanadianFrench:       mirandaCanadianFrenchName,
		language.Dutch:                mirandaDutchName,
		language.French:               mirandaFrenchName,
		language.German:               mirandaGermanName,
		language.Italian:              mirandaItalianName,
		language.Japanese:             mirandaJapaneseName,
		language.Korean:               mirandaKoreanName,
		language.LatinAmericanSpanish: mirandaLatinAmericanSpanishName,
		language.Russian:              mirandaRussianName,
		language.SimplifiedChinese:    mirandaSimplifiedChineseName,
		language.Spanish:              mirandaSpanishName,
		language.TraditionalChinese:   mirandaTraditionalChineseName}
)

var (
	mirandaCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: mirandaBirthday,
		Code:     mirandaCode,
		Key:      character.Miranda,
		Gender:   gender.Female,
		Name:     mirandaName}
)

var (
	mirandaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quackulous"}

	mirandaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "couac"}

	mirandaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwektakel"}

	mirandaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "couac"}

	mirandaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brotbrot"}

	mirandaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quaquao"}

	mirandaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なにさ"}

	mirandaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "괜찮아"}

	mirandaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cua-cuá"}

	mirandaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "покрясающе"}

	mirandaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "怎么回事"}

	mirandaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cua-cuá"}

	mirandaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "怎麼回事"}
)

var (
	mirandaPhrase = nook.Languages{
		language.AmericanEnglish:      mirandaAmericanEnglishPhrase,
		language.CanadianFrench:       mirandaCanadianFrenchPhrase,
		language.Dutch:                mirandaDutchPhrase,
		language.French:               mirandaFrenchPhrase,
		language.German:               mirandaGermanPhrase,
		language.Italian:              mirandaItalianPhrase,
		language.Japanese:             mirandaJapanesePhrase,
		language.Korean:               mirandaKoreanPhrase,
		language.LatinAmericanSpanish: mirandaLatinAmericanSpanishPhrase,
		language.Russian:              mirandaRussianPhrase,
		language.SimplifiedChinese:    mirandaSimplifiedChinesePhrase,
		language.Spanish:              mirandaSpanishPhrase,
		language.TraditionalChinese:   mirandaTraditionalChinesePhrase}
)

var (
	Miranda = nook.Villager{
		Character:   mirandaCharacter,
		Personality: personality.Snooty,
		Phrase:      mirandaPhrase}
)
