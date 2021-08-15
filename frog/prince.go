package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	princeBirthday = nook.Birthday{
		Day:   21,
		Month: time.July}
)

var (
	princeCode = nook.Code{
		Value: "flg12"}
)

var (
	princeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Prince"}

	princeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Princeboiiiiing"}

	princeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Princeblurp"}

	princeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Princeboiiiiing"}

	princeGermanName = nook.Name{
		Language: language.German,
		Value:    "Prinzgurp"}

	princeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Principeberp"}

	princeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カールですだ"}

	princeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "카일이리오슈"}

	princeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Felipebuuurp"}

	princeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Принсквак-квак"}

	princeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "青呱是的"}

	princeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Felipebuuurp"}

	princeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "青呱是的"}
)

var (
	princeName = nook.Languages{
		language.AmericanEnglish:      princeAmericanEnglishName,
		language.CanadianFrench:       princeCanadianFrenchName,
		language.Dutch:                princeDutchName,
		language.French:               princeFrenchName,
		language.German:               princeGermanName,
		language.Italian:              princeItalianName,
		language.Japanese:             princeJapaneseName,
		language.Korean:               princeKoreanName,
		language.LatinAmericanSpanish: princeLatinAmericanSpanishName,
		language.Russian:              princeRussianName,
		language.SimplifiedChinese:    princeSimplifiedChineseName,
		language.Spanish:              princeSpanishName,
		language.TraditionalChinese:   princeTraditionalChineseName}
)

var (
	princeCharacter = nook.Character{
		Animal:   Frog,
		Birthday: princeBirthday,
		Code:     princeCode,
		Gender:   nook.Male,
		Name:     princeName}
)

var (
	princeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "burrup"}

	princeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "boiiiiing"}

	princeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "blurp"}

	princeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "boiiiiing"}

	princeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gurp"}

	princeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "berp"}

	princeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですだ"}

	princeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "이리오슈"}

	princeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "buuurp"}

	princeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "квак-квак"}

	princeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是的"}

	princeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "buuurp"}

	princeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是的"}
)

var (
	princePhrase = nook.Languages{
		language.AmericanEnglish:      princeAmericanEnglishName,
		language.CanadianFrench:       princeCanadianFrenchName,
		language.Dutch:                princeDutchName,
		language.French:               princeFrenchName,
		language.German:               princeGermanName,
		language.Italian:              princeItalianName,
		language.Japanese:             princeJapaneseName,
		language.Korean:               princeKoreanName,
		language.LatinAmericanSpanish: princeLatinAmericanSpanishName,
		language.Russian:              princeRussianName,
		language.SimplifiedChinese:    princeSimplifiedChineseName,
		language.Spanish:              princeSpanishName,
		language.TraditionalChinese:   princeTraditionalChineseName}
)

var (
	Prince = nook.Villager{
		Character:   princeCharacter,
		Personality: nook.Lazy,
		Phrase:      princePhrase}
)
