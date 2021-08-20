package alligator

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
	dragoBirthday = nook.Birthday{
		Day:   12,
		Month: time.February}
)

var (
	dragoCode = nook.Code{
		Value: "crd08"}
)

var (
	dragoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Drago"}

	dragoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Drago"}

	dragoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Drago"}

	dragoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Drago"}

	dragoGermanName = nook.Name{
		Language: language.German,
		Value:    "Frederik"}

	dragoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dragonio"}

	dragoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タツオ"}

	dragoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "용남이"}

	dragoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dragonio"}

	dragoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Драго"}

	dragoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿龙"}

	dragoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Dragonio"}

	dragoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿龍"}
)

var (
	dragoName = nook.Languages{
		language.AmericanEnglish:      dragoAmericanEnglishName,
		language.CanadianFrench:       dragoCanadianFrenchName,
		language.Dutch:                dragoDutchName,
		language.French:               dragoFrenchName,
		language.German:               dragoGermanName,
		language.Italian:              dragoItalianName,
		language.Japanese:             dragoJapaneseName,
		language.Korean:               dragoKoreanName,
		language.LatinAmericanSpanish: dragoLatinAmericanSpanishName,
		language.Russian:              dragoRussianName,
		language.SimplifiedChinese:    dragoSimplifiedChineseName,
		language.Spanish:              dragoSpanishName,
		language.TraditionalChinese:   dragoTraditionalChineseName}
)

var (
	dragoCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: dragoBirthday,
		Code:     dragoCode,
		Key:      character.Drago,
		Gender:   gender.Male,
		Name:     dragoName,
		Special:  false}
)

var (
	dragoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "burrrn"}

	dragoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ducroc"}

	dragoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "puf"}

	dragoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ducroc"}

	dragoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hömpf"}

	dragoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tic tac"}

	dragoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "えーと"}

	dragoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "띠용띠용"}

	dragoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñami"}

	dragoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "огонь"}

	dragoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "然后"}

	dragoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "quemarrr"}

	dragoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "然後"}
)

var (
	dragoPhrase = nook.Languages{
		language.AmericanEnglish:      dragoAmericanEnglishPhrase,
		language.CanadianFrench:       dragoCanadianFrenchPhrase,
		language.Dutch:                dragoDutchPhrase,
		language.French:               dragoFrenchPhrase,
		language.German:               dragoGermanPhrase,
		language.Italian:              dragoItalianPhrase,
		language.Japanese:             dragoJapanesePhrase,
		language.Korean:               dragoKoreanPhrase,
		language.LatinAmericanSpanish: dragoLatinAmericanSpanishPhrase,
		language.Russian:              dragoRussianPhrase,
		language.SimplifiedChinese:    dragoSimplifiedChinesePhrase,
		language.Spanish:              dragoSpanishPhrase,
		language.TraditionalChinese:   dragoTraditionalChinesePhrase}
)

var (
	Drago = nook.Villager{
		Character:   dragoCharacter,
		Personality: personality.Lazy,
		Phrase:      dragoPhrase}
)
