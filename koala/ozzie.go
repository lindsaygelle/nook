package koala

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	ozzieBirthday = nook.Birthday{
		Day:   7,
		Month: time.May}
)

var (
	ozzieCode = nook.Code{
		Value: "kal05"}
)

var (
	ozzieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ozzie"}

	ozzieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kokokoa là là"}

	ozzieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ozzieouwe beer"}

	ozzieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kokokoa là là"}

	ozzieGermanName = nook.Name{
		Language: language.German,
		Value:    "Oskaraltes Haus"}

	ozzieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ozzyulabadula"}

	ozzieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドングリククッ"}

	ozzieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "동동이큐큐"}

	ozzieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Koloakoahhh"}

	ozzieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Оззимедведище"}

	ozzieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿栗颗颗"}

	ozzieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Koloaosete"}

	ozzieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿栗顆顆"}
)

var (
	ozzieName = nook.Languages{
		language.AmericanEnglish:      ozzieAmericanEnglishName,
		language.CanadianFrench:       ozzieCanadianFrenchName,
		language.Dutch:                ozzieDutchName,
		language.French:               ozzieFrenchName,
		language.German:               ozzieGermanName,
		language.Italian:              ozzieItalianName,
		language.Japanese:             ozzieJapaneseName,
		language.Korean:               ozzieKoreanName,
		language.LatinAmericanSpanish: ozzieLatinAmericanSpanishName,
		language.Russian:              ozzieRussianName,
		language.SimplifiedChinese:    ozzieSimplifiedChineseName,
		language.Spanish:              ozzieSpanishName,
		language.TraditionalChinese:   ozzieTraditionalChineseName}
)

var (
	ozzieCharacter = nook.Character{
		Animal:   Koala,
		Birthday: ozzieBirthday,
		Code:     ozzieCode,
		Gender:   nook.Male,
		Name:     ozzieName}
)

var (
	ozzieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ol' bear"}

	ozzieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "koa là là"}

	ozzieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ouwe beer"}

	ozzieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "koa là là"}

	ozzieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "altes Haus"}

	ozzieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ulabadula"}

	ozzieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ククッ"}

	ozzieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "큐큐"}

	ozzieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "koahhh"}

	ozzieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "медведище"}

	ozzieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "颗颗"}

	ozzieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "osete"}

	ozzieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "顆顆"}
)

var (
	ozziePhrase = nook.Languages{
		language.AmericanEnglish:      ozzieAmericanEnglishName,
		language.CanadianFrench:       ozzieCanadianFrenchName,
		language.Dutch:                ozzieDutchName,
		language.French:               ozzieFrenchName,
		language.German:               ozzieGermanName,
		language.Italian:              ozzieItalianName,
		language.Japanese:             ozzieJapaneseName,
		language.Korean:               ozzieKoreanName,
		language.LatinAmericanSpanish: ozzieLatinAmericanSpanishName,
		language.Russian:              ozzieRussianName,
		language.SimplifiedChinese:    ozzieSimplifiedChineseName,
		language.Spanish:              ozzieSpanishName,
		language.TraditionalChinese:   ozzieTraditionalChineseName}
)

var (
	Ozzie = nook.Villager{
		Character:   ozzieCharacter,
		Personality: nook.Lazy,
		Phrase:      ozziePhrase}
)
