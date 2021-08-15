package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	hazelBirthday = nook.Birthday{
		Day:   30,
		Month: time.August}
)

var (
	hazelCode = nook.Code{
		Value: "squ18"}
)

var (
	hazelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hazel"}

	hazelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pamelatac tac"}

	hazelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hazelbrauw"}

	hazelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pamelatac tac"}

	hazelGermanName = nook.Name{
		Language: language.German,
		Value:    "Nadinekekeke"}

	hazelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cigliolablink"}

	hazelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイリスだもんね"}

	hazelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이리스알아알아"}

	hazelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Anacardatac-tac"}

	hazelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хейзелфундук"}

	hazelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艾栗丝当然哦"}

	hazelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Anacardacastaña"}

	hazelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艾栗絲當然囉"}
)

var (
	hazelName = nook.Languages{
		language.AmericanEnglish:      hazelAmericanEnglishName,
		language.CanadianFrench:       hazelCanadianFrenchName,
		language.Dutch:                hazelDutchName,
		language.French:               hazelFrenchName,
		language.German:               hazelGermanName,
		language.Italian:              hazelItalianName,
		language.Japanese:             hazelJapaneseName,
		language.Korean:               hazelKoreanName,
		language.LatinAmericanSpanish: hazelLatinAmericanSpanishName,
		language.Russian:              hazelRussianName,
		language.SimplifiedChinese:    hazelSimplifiedChineseName,
		language.Spanish:              hazelSpanishName,
		language.TraditionalChinese:   hazelTraditionalChineseName}
)

var (
	hazelCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: hazelBirthday,
		Code:     hazelCode,
		Gender:   nook.Female,
		Name:     hazelName}
)

var (
	hazelAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "uni-wow"}

	hazelCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tac tac"}

	hazelDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brauw"}

	hazelFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tac tac"}

	hazelGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kekeke"}

	hazelItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blink"}

	hazelJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だもんね"}

	hazelKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "알아알아"}

	hazelLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tac-tac"}

	hazelRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фундук"}

	hazelSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "当然哦"}

	hazelSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "castaña"}

	hazelTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "當然囉"}
)

var (
	hazelPhrase = nook.Languages{
		language.AmericanEnglish:      hazelAmericanEnglishName,
		language.CanadianFrench:       hazelCanadianFrenchName,
		language.Dutch:                hazelDutchName,
		language.French:               hazelFrenchName,
		language.German:               hazelGermanName,
		language.Italian:              hazelItalianName,
		language.Japanese:             hazelJapaneseName,
		language.Korean:               hazelKoreanName,
		language.LatinAmericanSpanish: hazelLatinAmericanSpanishName,
		language.Russian:              hazelRussianName,
		language.SimplifiedChinese:    hazelSimplifiedChineseName,
		language.Spanish:              hazelSpanishName,
		language.TraditionalChinese:   hazelTraditionalChineseName}
)

var (
	Hazel = nook.Villager{
		Character:   hazelCharacter,
		Personality: nook.BigSister,
		Phrase:      hazelPhrase}
)
