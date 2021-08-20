package koala

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
		Value:    "Koko"}

	ozzieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ozzie"}

	ozzieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Koko"}

	ozzieGermanName = nook.Name{
		Language: language.German,
		Value:    "Oskar"}

	ozzieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ozzy"}

	ozzieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドングリ"}

	ozzieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "동동이"}

	ozzieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Koloa"}

	ozzieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Оззи"}

	ozzieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿栗"}

	ozzieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Koloa"}

	ozzieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿栗"}
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
		Animal:   animal.Koala,
		Birthday: ozzieBirthday,
		Code:     ozzieCode,
		Key:      character.Ozzie,
		Gender:   gender.Male,
		Name:     ozzieName,
		Special:  false}
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
		language.AmericanEnglish:      ozzieAmericanEnglishPhrase,
		language.CanadianFrench:       ozzieCanadianFrenchPhrase,
		language.Dutch:                ozzieDutchPhrase,
		language.French:               ozzieFrenchPhrase,
		language.German:               ozzieGermanPhrase,
		language.Italian:              ozzieItalianPhrase,
		language.Japanese:             ozzieJapanesePhrase,
		language.Korean:               ozzieKoreanPhrase,
		language.LatinAmericanSpanish: ozzieLatinAmericanSpanishPhrase,
		language.Russian:              ozzieRussianPhrase,
		language.SimplifiedChinese:    ozzieSimplifiedChinesePhrase,
		language.Spanish:              ozzieSpanishPhrase,
		language.TraditionalChinese:   ozzieTraditionalChinesePhrase}
)

var (
	Ozzie = nook.Villager{
		Character:   ozzieCharacter,
		Personality: personality.Lazy,
		Phrase:      ozziePhrase}
)
