package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	tiffanyBirthday = nook.Birthday{
		Day:   9,
		Month: time.January}
)

var (
	tiffanyCode = nook.Code{
		Value: "rbt07"}
)

var (
	tiffanyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tiffany"}

	tiffanyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tiphaine"}

	tiffanyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tiffany"}

	tiffanyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tiphaine"}

	tiffanyGermanName = nook.Name{
		Language: language.German,
		Value:    "Michelle"}

	tiffanyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Stefania"}

	tiffanyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バズレー"}

	tiffanyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바슬레"}

	tiffanyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tiffany"}

	tiffanyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тиффани"}

	tiffanySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大姐头"}

	tiffanySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tiffany"}

	tiffanyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大姐頭"}
)

var (
	tiffanyName = nook.Languages{
		language.AmericanEnglish:      tiffanyAmericanEnglishName,
		language.CanadianFrench:       tiffanyCanadianFrenchName,
		language.Dutch:                tiffanyDutchName,
		language.French:               tiffanyFrenchName,
		language.German:               tiffanyGermanName,
		language.Italian:              tiffanyItalianName,
		language.Japanese:             tiffanyJapaneseName,
		language.Korean:               tiffanyKoreanName,
		language.LatinAmericanSpanish: tiffanyLatinAmericanSpanishName,
		language.Russian:              tiffanyRussianName,
		language.SimplifiedChinese:    tiffanySimplifiedChineseName,
		language.Spanish:              tiffanySpanishName,
		language.TraditionalChinese:   tiffanyTraditionalChineseName}
)

var (
	tiffanyCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: tiffanyBirthday,
		Code:     tiffanyCode,
		Gender:   gender.Female,
		Name:     tiffanyName}
)

var (
	tiffanyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bunbun"}

	tiffanyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pinpin"}

	tiffanyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knijnelijn"}

	tiffanyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pinpin"}

	tiffanyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "öhrchen"}

	tiffanyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bunbun"}

	tiffanyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ってさ"}

	tiffanyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라던데"}

	tiffanyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "toin-lalá"}

	tiffanyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайчонок"}

	tiffanySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "我说呀"}

	tiffanySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pompón"}

	tiffanyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "我說呀"}
)

var (
	tiffanyPhrase = nook.Languages{
		language.AmericanEnglish:      tiffanyAmericanEnglishName,
		language.CanadianFrench:       tiffanyCanadianFrenchName,
		language.Dutch:                tiffanyDutchName,
		language.French:               tiffanyFrenchName,
		language.German:               tiffanyGermanName,
		language.Italian:              tiffanyItalianName,
		language.Japanese:             tiffanyJapaneseName,
		language.Korean:               tiffanyKoreanName,
		language.LatinAmericanSpanish: tiffanyLatinAmericanSpanishName,
		language.Russian:              tiffanyRussianName,
		language.SimplifiedChinese:    tiffanySimplifiedChineseName,
		language.Spanish:              tiffanySpanishName,
		language.TraditionalChinese:   tiffanyTraditionalChineseName}
)

var (
	Tiffany = nook.Villager{
		Character:   tiffanyCharacter,
		Personality: personality.Snooty,
		Phrase:      tiffanyPhrase}
)
