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
		Value:    "Célia"}

	callyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cally"}

	callyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Célia"}

	callyGermanName = nook.Name{
		Language: language.German,
		Value:    "Hörnchen"}

	callyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosa"}

	callyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パセリ"}

	callyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파슬리"}

	callyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Almendra"}

	callyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэлли"}

	callySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小芹"}

	callySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Almendra"}

	callyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小芹"}
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
		Animal:   animal.Squirrel,
		Birthday: callyBirthday,
		Code:     callyCode,
		Key:      character.Cally,
		Gender:   gender.Female,
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
		Personality: personality.Normal,
		Phrase:      callyPhrase}
)
