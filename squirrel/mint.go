package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Amandeahhhhhh"}

	mintDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Mintahhhhhh"}

	mintFrenchName = nook.Name{
		Language: language.French,
		Value:    "Amandeahhhhhh"}

	mintGermanName = nook.Name{
		Language: language.German,
		Value:    "Marikaknabber"}

	mintItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mentullasquercia"}

	mintJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミントうっふん"}

	mintKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "민트우훗"}

	mintLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mentabellotuá"}

	mintRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Минтах-ах"}

	mintSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小敏喔哼"}

	mintSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mentabichito"}

	mintTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小敏喔哼"}
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
		Animal:   Squirrel,
		Birthday: mintBirthday,
		Code:     mintCode,
		Gender:   nook.Female,
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
	Mint = nook.Villager{
		Character:   mintCharacter,
		Personality: nook.Snooty,
		Phrase:      mintPhrase}
)
