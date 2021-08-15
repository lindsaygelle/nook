package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	jeremiahBirthday = nook.Birthday{
		Day:   8,
		Month: time.July}
)

var (
	jeremiahCode = nook.Code{
		Value: "flg07"}
)

var (
	jeremiahAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jeremiah"}

	jeremiahCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jérémiebrooap"}

	jeremiahDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jeremiahrikkekik"}

	jeremiahFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jérémiebrooap"}

	jeremiahGermanName = nook.Name{
		Language: language.German,
		Value:    "Jörgnidip"}

	jeremiahItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Geremiafri fri"}

	jeremiahJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クワトロにゃむ"}

	jeremiahKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "드리미딩동댕"}

	jeremiahLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jeremíascroc-croc"}

	jeremiahRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джеремиябуль-буль"}

	jeremiahSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿刻嚼嚼"}

	jeremiahSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jeremíascroc-croc"}

	jeremiahTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿刻嚼嚼"}
)

var (
	jeremiahName = nook.Languages{
		language.AmericanEnglish:      jeremiahAmericanEnglishName,
		language.CanadianFrench:       jeremiahCanadianFrenchName,
		language.Dutch:                jeremiahDutchName,
		language.French:               jeremiahFrenchName,
		language.German:               jeremiahGermanName,
		language.Italian:              jeremiahItalianName,
		language.Japanese:             jeremiahJapaneseName,
		language.Korean:               jeremiahKoreanName,
		language.LatinAmericanSpanish: jeremiahLatinAmericanSpanishName,
		language.Russian:              jeremiahRussianName,
		language.SimplifiedChinese:    jeremiahSimplifiedChineseName,
		language.Spanish:              jeremiahSpanishName,
		language.TraditionalChinese:   jeremiahTraditionalChineseName}
)

var (
	jeremiahCharacter = nook.Character{
		Animal:   Frog,
		Birthday: jeremiahBirthday,
		Code:     jeremiahCode,
		Gender:   nook.Male,
		Name:     jeremiahName}
)

var (
	jeremiahAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nee-deep"}

	jeremiahCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "brooap"}

	jeremiahDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "rikkekik"}

	jeremiahFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "brooap"}

	jeremiahGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nidip"}

	jeremiahItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fri fri"}

	jeremiahJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "にゃむ"}

	jeremiahKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "딩동댕"}

	jeremiahLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "croc-croc"}

	jeremiahRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "буль-буль"}

	jeremiahSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嚼嚼"}

	jeremiahSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "croc-croc"}

	jeremiahTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嚼嚼"}
)

var (
	jeremiahPhrase = nook.Languages{
		language.AmericanEnglish:      jeremiahAmericanEnglishName,
		language.CanadianFrench:       jeremiahCanadianFrenchName,
		language.Dutch:                jeremiahDutchName,
		language.French:               jeremiahFrenchName,
		language.German:               jeremiahGermanName,
		language.Italian:              jeremiahItalianName,
		language.Japanese:             jeremiahJapaneseName,
		language.Korean:               jeremiahKoreanName,
		language.LatinAmericanSpanish: jeremiahLatinAmericanSpanishName,
		language.Russian:              jeremiahRussianName,
		language.SimplifiedChinese:    jeremiahSimplifiedChineseName,
		language.Spanish:              jeremiahSpanishName,
		language.TraditionalChinese:   jeremiahTraditionalChineseName}
)

var (
	Jeremiah = nook.Villager{
		Character:   jeremiahCharacter,
		Personality: nook.Lazy,
		Phrase:      jeremiahPhrase}
)
