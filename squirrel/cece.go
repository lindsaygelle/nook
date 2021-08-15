package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	ceceBirthday = nook.Birthday{
		Day:   28,
		Month: time.May}
)

var (
	ceceCode = nook.Code{
		Value: "squ19"}
)

var (
	ceceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cece"}

	ceceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kalamaréôte"}

	ceceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	ceceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kalamaréôte"}

	ceceGermanName = nook.Name{
		Language: language.German,
		Value:    "A-Chanaioli"}

	ceceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ceciliacalacala"}

	ceceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "なぎさよろしく"}

	ceceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나기사잘지내자"}

	ceceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ardunacosmar"}

	ceceRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	ceceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	ceceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ardunacosmar"}

	ceceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	ceceName = nook.Languages{
		language.AmericanEnglish:      ceceAmericanEnglishName,
		language.CanadianFrench:       ceceCanadianFrenchName,
		language.Dutch:                ceceDutchName,
		language.French:               ceceFrenchName,
		language.German:               ceceGermanName,
		language.Italian:              ceceItalianName,
		language.Japanese:             ceceJapaneseName,
		language.Korean:               ceceKoreanName,
		language.LatinAmericanSpanish: ceceLatinAmericanSpanishName,
		language.Russian:              ceceRussianName,
		language.SimplifiedChinese:    ceceSimplifiedChineseName,
		language.Spanish:              ceceSpanishName,
		language.TraditionalChinese:   ceceTraditionalChineseName}
)

var (
	ceceCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: ceceBirthday,
		Code:     ceceCode,
		Gender:   nook.Female,
		Name:     ceceName}
)

var (
	ceceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "stay fresh"}

	ceceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "maréôte"}

	ceceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	ceceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "maréôte"}

	ceceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "aioli"}

	ceceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "calacala"}

	ceceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よろしく"}

	ceceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "잘지내자"}

	ceceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cosmar"}

	ceceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	ceceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	ceceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cosmar"}

	ceceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	cecePhrase = nook.Languages{
		language.AmericanEnglish:      ceceAmericanEnglishName,
		language.CanadianFrench:       ceceCanadianFrenchName,
		language.Dutch:                ceceDutchName,
		language.French:               ceceFrenchName,
		language.German:               ceceGermanName,
		language.Italian:              ceceItalianName,
		language.Japanese:             ceceJapaneseName,
		language.Korean:               ceceKoreanName,
		language.LatinAmericanSpanish: ceceLatinAmericanSpanishName,
		language.Russian:              ceceRussianName,
		language.SimplifiedChinese:    ceceSimplifiedChineseName,
		language.Spanish:              ceceSpanishName,
		language.TraditionalChinese:   ceceTraditionalChineseName}
)

var (
	Cece = nook.Villager{
		Character:   ceceCharacter,
		Personality: nook.Peppy,
		Phrase:      cecePhrase}
)
