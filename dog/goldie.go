package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	goldieBirthday = nook.Birthday{
		Day:   27,
		Month: time.December}
)

var (
	goldieCode = nook.Code{
		Value: "dog00"}
)

var (
	goldieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Goldie"}

	goldieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mirza"}

	goldieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Goldie"}

	goldieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mirza"}

	goldieGermanName = nook.Name{
		Language: language.German,
		Value:    "Bienchen"}

	goldieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dora"}

	goldieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャラメル"}

	goldieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "카라멜"}

	goldieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tere"}

	goldieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Голди"}

	goldieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "牛奶糖"}

	goldieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tere"}

	goldieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "牛奶糖"}
)

var (
	goldieName = nook.Languages{
		language.AmericanEnglish:      goldieAmericanEnglishName,
		language.CanadianFrench:       goldieCanadianFrenchName,
		language.Dutch:                goldieDutchName,
		language.French:               goldieFrenchName,
		language.German:               goldieGermanName,
		language.Italian:              goldieItalianName,
		language.Japanese:             goldieJapaneseName,
		language.Korean:               goldieKoreanName,
		language.LatinAmericanSpanish: goldieLatinAmericanSpanishName,
		language.Russian:              goldieRussianName,
		language.SimplifiedChinese:    goldieSimplifiedChineseName,
		language.Spanish:              goldieSpanishName,
		language.TraditionalChinese:   goldieTraditionalChineseName}
)

var (
	goldieCharacter = nook.Character{
		Animal:   Dog,
		Birthday: goldieBirthday,
		Code:     goldieCode,
		Gender:   nook.Female,
		Name:     goldieName}
)

var (
	goldieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "woof"}

	goldieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ouaf"}

	goldieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "woef"}

	goldieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ouaf"}

	goldieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wuff"}

	goldieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "wuf wuf"}

	goldieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワン"}

	goldieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "왈왈"}

	goldieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guau-guau"}

	goldieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тяв"}

	goldieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "汪"}

	goldieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guau-guau"}

	goldieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "汪"}
)

var (
	goldiePhrase = nook.Languages{
		language.AmericanEnglish:      goldieAmericanEnglishName,
		language.CanadianFrench:       goldieCanadianFrenchName,
		language.Dutch:                goldieDutchName,
		language.French:               goldieFrenchName,
		language.German:               goldieGermanName,
		language.Italian:              goldieItalianName,
		language.Japanese:             goldieJapaneseName,
		language.Korean:               goldieKoreanName,
		language.LatinAmericanSpanish: goldieLatinAmericanSpanishName,
		language.Russian:              goldieRussianName,
		language.SimplifiedChinese:    goldieSimplifiedChineseName,
		language.Spanish:              goldieSpanishName,
		language.TraditionalChinese:   goldieTraditionalChineseName}
)

var (
	Goldie = nook.Villager{
		Character:   goldieCharacter,
		Personality: nook.Normal,
		Phrase:      goldiePhrase}
)
