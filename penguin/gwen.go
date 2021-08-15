package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	gwenBirthday = nook.Birthday{
		Day:   23,
		Month: time.January}
)

var (
	gwenCode = nook.Code{
		Value: "pgn05"}
)

var (
	gwenAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gwen"}

	gwenCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gwen"}

	gwenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gwen"}

	gwenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gwen"}

	gwenGermanName = nook.Name{
		Language: language.German,
		Value:    "Judith"}

	gwenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gelinda"}

	gwenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ポーラ"}

	gwenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "폴라"}

	gwenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fabiola"}

	gwenRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гвен"}

	gwenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宝拉"}

	gwenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fabiola"}

	gwenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "寶拉"}
)

var (
	gwenName = nook.Languages{
		language.AmericanEnglish:      gwenAmericanEnglishName,
		language.CanadianFrench:       gwenCanadianFrenchName,
		language.Dutch:                gwenDutchName,
		language.French:               gwenFrenchName,
		language.German:               gwenGermanName,
		language.Italian:              gwenItalianName,
		language.Japanese:             gwenJapaneseName,
		language.Korean:               gwenKoreanName,
		language.LatinAmericanSpanish: gwenLatinAmericanSpanishName,
		language.Russian:              gwenRussianName,
		language.SimplifiedChinese:    gwenSimplifiedChineseName,
		language.Spanish:              gwenSpanishName,
		language.TraditionalChinese:   gwenTraditionalChineseName}
)

var (
	gwenCharacter = nook.Character{
		Animal:   Penguin,
		Birthday: gwenBirthday,
		Code:     gwenCode,
		Gender:   nook.Female,
		Name:     gwenName}
)

var (
	gwenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "h-h-h-hon"}

	gwenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "h-h-h-hon"}

	gwenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "d-d-duifje"}

	gwenFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "h-h-h-hon"}

	gwenGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "sssweetie"}

	gwenItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "a-a-amor"}

	gwenJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウフフ"}

	gwenKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우훗훗"}

	gwenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ki-ki-kiú"}

	gwenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "м-милашка"}

	gwenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唔呼呼"}

	gwenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "crustáceo"}

	gwenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唔呼呼"}
)

var (
	gwenPhrase = nook.Languages{
		language.AmericanEnglish:      gwenAmericanEnglishName,
		language.CanadianFrench:       gwenCanadianFrenchName,
		language.Dutch:                gwenDutchName,
		language.French:               gwenFrenchName,
		language.German:               gwenGermanName,
		language.Italian:              gwenItalianName,
		language.Japanese:             gwenJapaneseName,
		language.Korean:               gwenKoreanName,
		language.LatinAmericanSpanish: gwenLatinAmericanSpanishName,
		language.Russian:              gwenRussianName,
		language.SimplifiedChinese:    gwenSimplifiedChineseName,
		language.Spanish:              gwenSpanishName,
		language.TraditionalChinese:   gwenTraditionalChineseName}
)

var (
	Gwen = nook.Villager{
		Character:   gwenCharacter,
		Personality: nook.Snooty,
		Phrase:      gwenPhrase}
)
