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
	// mintBirthday represents Mint's birthday (May 2nd).
	mintBirthday = nook.Birthday{
		Day:   2,
		Month: time.May}
)

var (
	// mintCode represents Mint's unique code ("squ09").
	mintCode = nook.Code{
		Value: "squ09"}
)

var (
	// mintAmericanEnglishName represents Mint's name in American English.
	mintAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mint"}

	// mintCanadianFrenchName represents Mint's name in Canadian French.
	mintCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Amande"}

	// mintDutchName represents Mint's name in Dutch.
	mintDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mint"}

	// mintFrenchName represents Mint's name in French.
	mintFrenchName = nook.Name{
		Language: language.French,
		Value:    "Amande"}

	// mintGermanName represents Mint's name in German.
	mintGermanName = nook.Name{
		Language: language.German,
		Value:    "Marika"}

	// mintItalianName represents Mint's name in Italian.
	mintItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mentulla"}

	// mintJapaneseName represents Mint's name in Japanese.
	mintJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミント"}

	// mintKoreanName represents Mint's name in Korean.
	mintKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "민트"}

	// mintLatinAmericanSpanishName represents Mint's name in Latin American Spanish.
	mintLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Menta"}

	// mintRussianName represents Mint's name in Russian.
	mintRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Минт"}

	// mintSimplifiedChineseName represents Mint's name in Simplified Chinese.
	mintSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小敏"}

	// mintSpanishName represents Mint's name in Spanish.
	mintSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Menta"}

	// mintTraditionalChineseName represents Mint's name in Traditional Chinese.
	mintTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小敏"}
)

var (
	// mintName represents Mint's name in different languages.
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
	// mintCharacter represents Mint's character information.
	mintCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: mintBirthday,
		Code:     mintCode,
		Key:      character.Mint,
		Gender:   gender.Female,
		Name:     mintName,
		Special:  false}
)

var (
	// mintAmericanEnglishPhrase represents Mint's phrase in American English.
	mintAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ahhhhhh"}

	// mintCanadianFrenchPhrase represents Mint's phrase in Canadian French.
	mintCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ahhhhhh"}

	// mintDutchPhrase represents Mint's phrase in Dutch.
	mintDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ahhhhhh"}

	// mintFrenchPhrase represents Mint's phrase in French.
	mintFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ahhhhhh"}

	// mintGermanPhrase represents Mint's phrase in German.
	mintGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knabber"}

	// mintItalianPhrase represents Mint's phrase in Italian.
	mintItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squercia"}

	// mintJapanesePhrase represents Mint's phrase in Japanese.
	mintJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "うっふん"}

	// mintKoreanPhrase represents Mint's phrase in Korean.
	mintKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우훗"}

	// mintLatinAmericanSpanishPhrase represents Mint's phrase in Latin American Spanish.
	mintLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bellotuá"}

	// mintRussianPhrase represents Mint's phrase in Russian.
	mintRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ах-ах"}

	// mintSimplifiedChinesePhrase represents Mint's phrase in Simplified Chinese.
	mintSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喔哼"}

	// mintSpanishPhrase represents Mint's phrase in Spanish.
	mintSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bichito"}

	// mintTraditionalChinesePhrase represents Mint's phrase in Traditional Chinese.
	mintTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喔哼"}
)

var (
	// mintPhrase represents Mint's phrase in different languages.
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
	// Mint represents the character Mint with her complete information.
	Mint = nook.Villager{
		Character:   mintCharacter,
		Personality: personality.Snooty,
		Phrase:      mintPhrase}
)
