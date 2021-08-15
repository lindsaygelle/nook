package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	callyBirthday = nook.Birthday{
		Day:   4,
		Month: time.September}
)

var (
	callyCode = nook.Code{
		Value: "squ11"}
)

var (
	callyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cally"}

	callyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Céliacacahuète"}

	callyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Callycashew"}

	callyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Céliayeeeees"}

	callyGermanName = nook.Name{
		Language: language.German,
		Value:    "Hörnchenhuiii"}

	callyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "RosaSCRUNCH"}

	callyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パセリララー"}

	callyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파슬리랄랄라"}

	callyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Almendrañip-ñip"}

	callyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэллиух ты"}

	callySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小芹啦啦"}

	callySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Almendramacadamia"}

	callyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小芹啦啦"}
)

var (
	callyName = nook.Languages{
		language.AmericanEnglish:      callyAmericanEnglishName,
		language.CanadianFrench:       callyCanadianFrenchName,
		language.Dutch:                callyDutchName,
		language.French:               callyFrenchName,
		language.German:               callyGermanName,
		language.Italian:              callyItalianName,
		language.Japanese:             callyJapaneseName,
		language.Korean:               callyKoreanName,
		language.LatinAmericanSpanish: callyLatinAmericanSpanishName,
		language.Russian:              callyRussianName,
		language.SimplifiedChinese:    callySimplifiedChineseName,
		language.Spanish:              callySpanishName,
		language.TraditionalChinese:   callyTraditionalChineseName}
)

var (
	callyCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: callyBirthday,
		Code:     callyCode,
		Gender:   nook.Female,
		Name:     callyName}
)

var (
	callyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "WHEE"}

	callyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cacahuète"}

	callyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "cashew"}

	callyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "yeeeees"}

	callyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "huiii"}

	callyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "SCRUNCH"}

	callyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ララー"}

	callyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "랄랄라"}

	callyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñip-ñip"}

	callyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ух ты"}

	callySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啦啦"}

	callySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "macadamia"}

	callyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啦啦"}
)

var (
	callyPhrase = nook.Languages{
		language.AmericanEnglish:      callyAmericanEnglishName,
		language.CanadianFrench:       callyCanadianFrenchName,
		language.Dutch:                callyDutchName,
		language.French:               callyFrenchName,
		language.German:               callyGermanName,
		language.Italian:              callyItalianName,
		language.Japanese:             callyJapaneseName,
		language.Korean:               callyKoreanName,
		language.LatinAmericanSpanish: callyLatinAmericanSpanishName,
		language.Russian:              callyRussianName,
		language.SimplifiedChinese:    callySimplifiedChineseName,
		language.Spanish:              callySpanishName,
		language.TraditionalChinese:   callyTraditionalChineseName}
)

var (
	Cally = nook.Villager{
		Character:   callyCharacter,
		Personality: nook.Normal,
		Phrase:      callyPhrase}
)
