package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	camofrogBirthday = nook.Birthday{
		Day:   5,
		Month: time.June}
)

var (
	camofrogCode = nook.Code{
		Value: "flg03"}
)

var (
	camofrogAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Camofrog"}

	camofrogCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Miloskaki"}

	camofrogDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Camofroggeef acht"}

	camofrogFrenchName = nook.Name{
		Language: language.French,
		Value:    "Milostaïaut"}

	camofrogGermanName = nook.Name{
		Language: language.German,
		Value:    "Tarnoquak-quak"}

	camofrogItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cramboattacroak"}

	camofrogJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フルメタルわい"}

	camofrogKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "충성와이"}

	camofrogLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Comandocroooa"}

	camofrogRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Камофрогшагом-квак"}

	camofrogSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "迷仔喂"}

	camofrogSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Comandorecluta"}

	camofrogTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "迷仔喂"}
)

var (
	camofrogName = nook.Languages{
		language.AmericanEnglish:      camofrogAmericanEnglishName,
		language.CanadianFrench:       camofrogCanadianFrenchName,
		language.Dutch:                camofrogDutchName,
		language.French:               camofrogFrenchName,
		language.German:               camofrogGermanName,
		language.Italian:              camofrogItalianName,
		language.Japanese:             camofrogJapaneseName,
		language.Korean:               camofrogKoreanName,
		language.LatinAmericanSpanish: camofrogLatinAmericanSpanishName,
		language.Russian:              camofrogRussianName,
		language.SimplifiedChinese:    camofrogSimplifiedChineseName,
		language.Spanish:              camofrogSpanishName,
		language.TraditionalChinese:   camofrogTraditionalChineseName}
)

var (
	camofrogCharacter = nook.Character{
		Animal:   Frog,
		Birthday: camofrogBirthday,
		Code:     camofrogCode,
		Gender:   nook.Male,
		Name:     camofrogName}
)

var (
	camofrogAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ten-hut"}

	camofrogCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kaki"}

	camofrogDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "geef acht"}

	camofrogFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "taïaut"}

	camofrogGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quak-quak"}

	camofrogItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "attacroak"}

	camofrogJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "わい"}

	camofrogKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "와이"}

	camofrogLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "croooa"}

	camofrogRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шагом-квак"}

	camofrogSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喂"}

	camofrogSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "recluta"}

	camofrogTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喂"}
)

var (
	camofrogPhrase = nook.Languages{
		language.AmericanEnglish:      camofrogAmericanEnglishName,
		language.CanadianFrench:       camofrogCanadianFrenchName,
		language.Dutch:                camofrogDutchName,
		language.French:               camofrogFrenchName,
		language.German:               camofrogGermanName,
		language.Italian:              camofrogItalianName,
		language.Japanese:             camofrogJapaneseName,
		language.Korean:               camofrogKoreanName,
		language.LatinAmericanSpanish: camofrogLatinAmericanSpanishName,
		language.Russian:              camofrogRussianName,
		language.SimplifiedChinese:    camofrogSimplifiedChineseName,
		language.Spanish:              camofrogSpanishName,
		language.TraditionalChinese:   camofrogTraditionalChineseName}
)

var (
	Camofrog = nook.Villager{
		Character:   camofrogCharacter,
		Personality: nook.Cranky,
		Phrase:      camofrogPhrase}
)
