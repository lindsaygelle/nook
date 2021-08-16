package koala

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	melbaBirthday = nook.Birthday{
		Day:   12,
		Month: time.April}
)

var (
	melbaCode = nook.Code{
		Value: "kal02"}
)

var (
	melbaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Melba"}

	melbaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Melba"}

	melbaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Melba"}

	melbaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Melba"}

	melbaGermanName = nook.Name{
		Language: language.German,
		Value:    "Kornelia"}

	melbaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Melba"}

	melbaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アデレード"}

	melbaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아델레이드"}

	melbaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Melba"}

	melbaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мельба"}

	melbaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿得"}

	melbaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Melba"}

	melbaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿得"}
)

var (
	melbaName = nook.Languages{
		language.AmericanEnglish:      melbaAmericanEnglishName,
		language.CanadianFrench:       melbaCanadianFrenchName,
		language.Dutch:                melbaDutchName,
		language.French:               melbaFrenchName,
		language.German:               melbaGermanName,
		language.Italian:              melbaItalianName,
		language.Japanese:             melbaJapaneseName,
		language.Korean:               melbaKoreanName,
		language.LatinAmericanSpanish: melbaLatinAmericanSpanishName,
		language.Russian:              melbaRussianName,
		language.SimplifiedChinese:    melbaSimplifiedChineseName,
		language.Spanish:              melbaSpanishName,
		language.TraditionalChinese:   melbaTraditionalChineseName}
)

var (
	melbaCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: melbaBirthday,
		Code:     melbaCode,
		Gender:   gender.Female,
		Name:     melbaName}
)

var (
	melbaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "toasty"}

	melbaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pêchou"}

	melbaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "toastje"}

	melbaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pêchou"}

	melbaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zubba"}

	melbaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "toasty"}

	melbaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とっても"}

	melbaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "정말로요"}

	melbaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pichú"}

	melbaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрустяшка"}

	melbaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不得了"}

	melbaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tesoro"}

	melbaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不得了"}
)

var (
	melbaPhrase = nook.Languages{
		language.AmericanEnglish:      melbaAmericanEnglishName,
		language.CanadianFrench:       melbaCanadianFrenchName,
		language.Dutch:                melbaDutchName,
		language.French:               melbaFrenchName,
		language.German:               melbaGermanName,
		language.Italian:              melbaItalianName,
		language.Japanese:             melbaJapaneseName,
		language.Korean:               melbaKoreanName,
		language.LatinAmericanSpanish: melbaLatinAmericanSpanishName,
		language.Russian:              melbaRussianName,
		language.SimplifiedChinese:    melbaSimplifiedChineseName,
		language.Spanish:              melbaSpanishName,
		language.TraditionalChinese:   melbaTraditionalChineseName}
)

var (
	Melba = nook.Villager{
		Character:   melbaCharacter,
		Personality: personality.Normal,
		Phrase:      melbaPhrase}
)
