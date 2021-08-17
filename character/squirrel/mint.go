package squirrel

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
	mintBirthday = nook.Birthday{
		Day:   2,
		Month: time.May}
)

var (
	mintCode = nook.Code{
		Value: "squ09"}
)

var (
	mintAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mint"}

	mintCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Amande"}

	mintDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mint"}

	mintFrenchName = nook.Name{
		Language: language.French,
		Value:    "Amande"}

	mintGermanName = nook.Name{
		Language: language.German,
		Value:    "Marika"}

	mintItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mentulla"}

	mintJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミント"}

	mintKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "민트"}

	mintLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Menta"}

	mintRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Минт"}

	mintSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小敏"}

	mintSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Menta"}

	mintTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小敏"}
)

var (
	mintName = nook.Languages{
		language.AmericanEnglish:      mintAmericanEnglishName,
		language.CanadianFrench:       mintCanadianFrenchName,
		language.Dutch:                mintDutchName,
		language.French:               mintFrenchName,
		language.German:               mintGermanName,
		language.Italian:              mintItalianName,
		language.Japanese:             mintJapaneseName,
		language.Korean:               mintKoreanName,
		language.LatinAmericanSpanish: mintLatinAmericanSpanishName,
		language.Russian:              mintRussianName,
		language.SimplifiedChinese:    mintSimplifiedChineseName,
		language.Spanish:              mintSpanishName,
		language.TraditionalChinese:   mintTraditionalChineseName}
)

var (
	mintCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: mintBirthday,
		Code:     mintCode,
		Key:      character.Mint,
		Gender:   gender.Female,
		Name:     mintName}
)

var (
	mintAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ahhhhhh"}

	mintCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ahhhhhh"}

	mintDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ahhhhhh"}

	mintFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ahhhhhh"}

	mintGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knabber"}

	mintItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squercia"}

	mintJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "うっふん"}

	mintKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우훗"}

	mintLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bellotuá"}

	mintRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ах-ах"}

	mintSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喔哼"}

	mintSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bichito"}

	mintTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喔哼"}
)

var (
	mintPhrase = nook.Languages{
		language.AmericanEnglish:      mintAmericanEnglishPhrase,
		language.CanadianFrench:       mintCanadianFrenchPhrase,
		language.Dutch:                mintDutchPhrase,
		language.French:               mintFrenchPhrase,
		language.German:               mintGermanPhrase,
		language.Italian:              mintItalianPhrase,
		language.Japanese:             mintJapanesePhrase,
		language.Korean:               mintKoreanPhrase,
		language.LatinAmericanSpanish: mintLatinAmericanSpanishPhrase,
		language.Russian:              mintRussianPhrase,
		language.SimplifiedChinese:    mintSimplifiedChinesePhrase,
		language.Spanish:              mintSpanishPhrase,
		language.TraditionalChinese:   mintTraditionalChinesePhrase}
)

var (
	Mint = nook.Villager{
		Character:   mintCharacter,
		Personality: personality.Snooty,
		Phrase:      mintPhrase}
)
