package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	francineBirthday = nook.Birthday{
		Day:   22,
		Month: time.January}
)

var (
	francineCode = nook.Code{
		Value: "rbt12"}
)

var (
	francineAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Francine"}

	francineCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nadine"}

	francineDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Francine"}

	francineFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nadine"}

	francineGermanName = nook.Name{
		Language: language.German,
		Value:    "Manu"}

	francineItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Franca"}

	francineJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フランソワ"}

	francineKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "프랑소와"}

	francineLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Natacha"}

	francineRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Франсин"}

	francineSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "法蓝琪"}

	francineSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Natacha"}

	francineTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "法藍琪"}
)

var (
	francineName = nook.Languages{
		language.AmericanEnglish:      francineAmericanEnglishName,
		language.CanadianFrench:       francineCanadianFrenchName,
		language.Dutch:                francineDutchName,
		language.French:               francineFrenchName,
		language.German:               francineGermanName,
		language.Italian:              francineItalianName,
		language.Japanese:             francineJapaneseName,
		language.Korean:               francineKoreanName,
		language.LatinAmericanSpanish: francineLatinAmericanSpanishName,
		language.Russian:              francineRussianName,
		language.SimplifiedChinese:    francineSimplifiedChineseName,
		language.Spanish:              francineSpanishName,
		language.TraditionalChinese:   francineTraditionalChineseName}
)

var (
	francineCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: francineBirthday,
		Code:     francineCode,
		Gender:   gender.Female,
		Name:     francineName}
)

var (
	francineAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "karat"}

	francineCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "carotte"}

	francineDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wortels"}

	francineFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "carotte"}

	francineGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hey man"}

	francineItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "lulalà"}

	francineJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ルララ"}

	francineKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쇼봉"}

	francineLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kaninchen"}

	francineRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "морковка"}

	francineSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噜啦啦"}

	francineSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "kaninchen"}

	francineTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嚕啦啦"}
)

var (
	francinePhrase = nook.Languages{
		language.AmericanEnglish:      francineAmericanEnglishName,
		language.CanadianFrench:       francineCanadianFrenchName,
		language.Dutch:                francineDutchName,
		language.French:               francineFrenchName,
		language.German:               francineGermanName,
		language.Italian:              francineItalianName,
		language.Japanese:             francineJapaneseName,
		language.Korean:               francineKoreanName,
		language.LatinAmericanSpanish: francineLatinAmericanSpanishName,
		language.Russian:              francineRussianName,
		language.SimplifiedChinese:    francineSimplifiedChineseName,
		language.Spanish:              francineSpanishName,
		language.TraditionalChinese:   francineTraditionalChineseName}
)

var (
	Francine = nook.Villager{
		Character:   francineCharacter,
		Personality: personality.Snooty,
		Phrase:      francinePhrase}
)
