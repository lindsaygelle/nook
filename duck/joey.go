package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	joeyBirthday = nook.Birthday{
		Day:   3,
		Month: time.January}
)

var (
	joeyCode = nook.Code{
		Value: "duk01"}
)

var (
	joeyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Joey"}

	joeyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Josephbeeeuuuurk"}

	joeyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Joeykèèèk"}

	joeyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Josephbeeeuuuurk"}

	joeyGermanName = nook.Name{
		Language: language.German,
		Value:    "Kalletschiiirp"}

	joeyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pinosquack"}

	joeyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リチャードでヤンス"}

	joeyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리처드그래유"}

	joeyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pascualuac-uac"}

	joeyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джоуикряк"}

	joeySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "李察鸭鸭"}

	joeySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pascualuac-uac"}

	joeyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "李察鴨鴨"}
)

var (
	joeyName = nook.Languages{
		language.AmericanEnglish:      joeyAmericanEnglishName,
		language.CanadianFrench:       joeyCanadianFrenchName,
		language.Dutch:                joeyDutchName,
		language.French:               joeyFrenchName,
		language.German:               joeyGermanName,
		language.Italian:              joeyItalianName,
		language.Japanese:             joeyJapaneseName,
		language.Korean:               joeyKoreanName,
		language.LatinAmericanSpanish: joeyLatinAmericanSpanishName,
		language.Russian:              joeyRussianName,
		language.SimplifiedChinese:    joeySimplifiedChineseName,
		language.Spanish:              joeySpanishName,
		language.TraditionalChinese:   joeyTraditionalChineseName}
)

var (
	joeyCharacter = nook.Character{
		Animal:   Duck,
		Birthday: joeyBirthday,
		Code:     joeyCode,
		Gender:   nook.Male,
		Name:     joeyName}
)

var (
	joeyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bleeeeeck"}

	joeyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "beeeuuuurk"}

	joeyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kèèèk"}

	joeyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "beeeuuuurk"}

	joeyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tschiiirp"}

	joeyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squack"}

	joeyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でヤンス"}

	joeyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그래유"}

	joeyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uac-uac"}

	joeyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кряк"}

	joeySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鸭鸭"}

	joeySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uac-uac"}

	joeyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鴨鴨"}
)

var (
	joeyPhrase = nook.Languages{
		language.AmericanEnglish:      joeyAmericanEnglishName,
		language.CanadianFrench:       joeyCanadianFrenchName,
		language.Dutch:                joeyDutchName,
		language.French:               joeyFrenchName,
		language.German:               joeyGermanName,
		language.Italian:              joeyItalianName,
		language.Japanese:             joeyJapaneseName,
		language.Korean:               joeyKoreanName,
		language.LatinAmericanSpanish: joeyLatinAmericanSpanishName,
		language.Russian:              joeyRussianName,
		language.SimplifiedChinese:    joeySimplifiedChineseName,
		language.Spanish:              joeySpanishName,
		language.TraditionalChinese:   joeyTraditionalChineseName}
)

var (
	Joey = nook.Villager{
		Character:   joeyCharacter,
		Personality: nook.Lazy,
		Phrase:      joeyPhrase}
)
