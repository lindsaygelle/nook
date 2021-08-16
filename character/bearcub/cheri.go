package bearcub

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
	cheriBirthday = nook.Birthday{
		Day:   17,
		Month: time.March}
)

var (
	cheriCode = nook.Code{
		Value: "cbr10"}
)

var (
	cheriAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cheri"}

	cheriCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rosalie"}

	cheriDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cheri"}

	cheriFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosalie"}

	cheriGermanName = nook.Name{
		Language: language.German,
		Value:    "Claudia"}

	cheriItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bonbon"}

	cheriJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アセロラ"}

	cheriKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아세로라"}

	cheriLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cerecita"}

	cheriRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шери"}

	cheriSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "樱桃"}

	cheriSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cerecita"}

	cheriTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "櫻桃"}
)

var (
	cheriName = nook.Languages{
		language.AmericanEnglish:      cheriAmericanEnglishName,
		language.CanadianFrench:       cheriCanadianFrenchName,
		language.Dutch:                cheriDutchName,
		language.French:               cheriFrenchName,
		language.German:               cheriGermanName,
		language.Italian:              cheriItalianName,
		language.Japanese:             cheriJapaneseName,
		language.Korean:               cheriKoreanName,
		language.LatinAmericanSpanish: cheriLatinAmericanSpanishName,
		language.Russian:              cheriRussianName,
		language.SimplifiedChinese:    cheriSimplifiedChineseName,
		language.Spanish:              cheriSpanishName,
		language.TraditionalChinese:   cheriTraditionalChineseName}
)

var (
	cheriCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: cheriBirthday,
		Code:     cheriCode,
		Key:      character.Cheri,
		Gender:   gender.Female,
		Name:     cheriName}
)

var (
	cheriAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tralala"}

	cheriCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tralala"}

	cheriDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tralala"}

	cheriFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tralala"}

	cheriGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tralala"}

	cheriItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "trallalà"}

	cheriJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なんてね"}

	cheriKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어쩜이래"}

	cheriLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tralará"}

	cheriRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тра-ля-ля"}

	cheriSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "开玩笑的"}

	cheriSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tralará"}

	cheriTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "開玩笑的"}
)

var (
	cheriPhrase = nook.Languages{
		language.AmericanEnglish:      cheriAmericanEnglishName,
		language.CanadianFrench:       cheriCanadianFrenchName,
		language.Dutch:                cheriDutchName,
		language.French:               cheriFrenchName,
		language.German:               cheriGermanName,
		language.Italian:              cheriItalianName,
		language.Japanese:             cheriJapaneseName,
		language.Korean:               cheriKoreanName,
		language.LatinAmericanSpanish: cheriLatinAmericanSpanishName,
		language.Russian:              cheriRussianName,
		language.SimplifiedChinese:    cheriSimplifiedChineseName,
		language.Spanish:              cheriSpanishName,
		language.TraditionalChinese:   cheriTraditionalChineseName}
)

var (
	Cheri = nook.Villager{
		Character:   cheriCharacter,
		Personality: personality.Peppy,
		Phrase:      cheriPhrase}
)
