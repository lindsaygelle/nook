package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	chowBirthday = nook.Birthday{
		Day:   22,
		Month: time.July}
)

var (
	chowCode = nook.Code{
		Value: "bea03"}
)

var (
	chowAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chow"}

	chowCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chulinlala"}

	chowDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chowkiai"}

	chowFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chulinlala"}

	chowGermanName = nook.Name{
		Language: language.German,
		Value:    "Changhiijaa"}

	chowItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chowchowehilà"}

	chowJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャウヤンアルヨ"}

	chowKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "츄양그럼"}

	chowLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pandogrinchu"}

	chowRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чауай-яй"}

	chowSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朝阳有喔"}

	chowSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pandoándale-oso"}

	chowTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朝陽有喔"}
)

var (
	chowName = nook.Languages{
		language.AmericanEnglish:      chowAmericanEnglishName,
		language.CanadianFrench:       chowCanadianFrenchName,
		language.Dutch:                chowDutchName,
		language.French:               chowFrenchName,
		language.German:               chowGermanName,
		language.Italian:              chowItalianName,
		language.Japanese:             chowJapaneseName,
		language.Korean:               chowKoreanName,
		language.LatinAmericanSpanish: chowLatinAmericanSpanishName,
		language.Russian:              chowRussianName,
		language.SimplifiedChinese:    chowSimplifiedChineseName,
		language.Spanish:              chowSpanishName,
		language.TraditionalChinese:   chowTraditionalChineseName}
)

var (
	chowCharacter = nook.Character{
		Animal:   Bear,
		Birthday: chowBirthday,
		Code:     chowCode,
		Gender:   nook.Male,
		Name:     chowName}
)

var (
	chowAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "aiya"}

	chowCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "lala"}

	chowDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kiai"}

	chowFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "lala"}

	chowGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hiijaa"}

	chowItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ehilà"}

	chowJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アルヨ"}

	chowKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그럼"}

	chowLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grinchu"}

	chowRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ай-яй"}

	chowSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "有喔"}

	chowSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ándale-oso"}

	chowTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "有喔"}
)

var (
	chowPhrase = nook.Languages{
		language.AmericanEnglish:      chowAmericanEnglishName,
		language.CanadianFrench:       chowCanadianFrenchName,
		language.Dutch:                chowDutchName,
		language.French:               chowFrenchName,
		language.German:               chowGermanName,
		language.Italian:              chowItalianName,
		language.Japanese:             chowJapaneseName,
		language.Korean:               chowKoreanName,
		language.LatinAmericanSpanish: chowLatinAmericanSpanishName,
		language.Russian:              chowRussianName,
		language.SimplifiedChinese:    chowSimplifiedChineseName,
		language.Spanish:              chowSpanishName,
		language.TraditionalChinese:   chowTraditionalChineseName}
)

var (
	Chow = nook.Villager{
		Character:   chowCharacter,
		Personality: nook.Cranky,
		Phrase:      chowPhrase}
)
