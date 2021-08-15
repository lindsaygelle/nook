package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Nadinecarotte"}

	francineDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Francinewortels"}

	francineFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nadinecarotte"}

	francineGermanName = nook.Name{
		Language: language.German,
		Value:    "Manuhey man"}

	francineItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Francalulalà"}

	francineJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フランソワルララ"}

	francineKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "프랑소와쇼봉"}

	francineLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Natachakaninchen"}

	francineRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Франсинморковка"}

	francineSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "法蓝琪噜啦啦"}

	francineSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Natachakaninchen"}

	francineTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "法藍琪嚕啦啦"}
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
		Animal:   Rabbit,
		Birthday: francineBirthday,
		Code:     francineCode,
		Gender:   nook.Female,
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
		Personality: nook.Snooty,
		Phrase:      francinePhrase}
)
