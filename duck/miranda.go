package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Maëlliscouac"}

	mirandaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mirandakwektakel"}

	mirandaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maëlliscouac"}

	mirandaGermanName = nook.Name{
		Language: language.German,
		Value:    "Tanyabrotbrot"}

	mirandaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mirandaquaquao"}

	mirandaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミランダなにさ"}

	mirandaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미란다괜찮아"}

	mirandaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mirandacua-cuá"}

	mirandaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мирандапокрясающе"}

	mirandaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "米兰达怎么回事"}

	mirandaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mirandacua-cuá"}

	mirandaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "米蘭達怎麼回事"}
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
		Animal:   Duck,
		Birthday: mirandaBirthday,
		Code:     mirandaCode,
		Gender:   nook.Female,
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
	Miranda = nook.Villager{
		Character:   mirandaCharacter,
		Personality: nook.Snooty,
		Phrase:      mirandaPhrase}
)
