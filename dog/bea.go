package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	beaBirthday = nook.Birthday{
		Day:   15,
		Month: time.October}
)

var (
	beaCode = nook.Code{
		Value: "dog10"}
)

var (
	beaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bea"}

	beaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Béamon chaton"}

	beaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Beabingo"}

	beaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Béamon chaton"}

	beaGermanName = nook.Name{
		Language: language.German,
		Value:    "Beabingo"}

	beaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cucciolabingo"}

	beaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ベーグルグー"}

	beaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "베이글쫀득"}

	beaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Beagufi-guf"}

	beaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Беадинго"}

	beaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贝果宾果"}

	beaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Beahuesín"}

	beaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "貝果賓果"}
)

var (
	beaName = nook.Languages{
		language.AmericanEnglish:      beaAmericanEnglishName,
		language.CanadianFrench:       beaCanadianFrenchName,
		language.Dutch:                beaDutchName,
		language.French:               beaFrenchName,
		language.German:               beaGermanName,
		language.Italian:              beaItalianName,
		language.Japanese:             beaJapaneseName,
		language.Korean:               beaKoreanName,
		language.LatinAmericanSpanish: beaLatinAmericanSpanishName,
		language.Russian:              beaRussianName,
		language.SimplifiedChinese:    beaSimplifiedChineseName,
		language.Spanish:              beaSpanishName,
		language.TraditionalChinese:   beaTraditionalChineseName}
)

var (
	beaCharacter = nook.Character{
		Animal:   Dog,
		Birthday: beaBirthday,
		Code:     beaCode,
		Gender:   nook.Female,
		Name:     beaName}
)

var (
	beaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bingo"}

	beaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon chaton"}

	beaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bingo"}

	beaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon chaton"}

	beaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bingo"}

	beaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bingo"}

	beaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "グー"}

	beaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쫀득"}

	beaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gufi-guf"}

	beaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "динго"}

	beaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宾果"}

	beaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "huesín"}

	beaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "賓果"}
)

var (
	beaPhrase = nook.Languages{
		language.AmericanEnglish:      beaAmericanEnglishName,
		language.CanadianFrench:       beaCanadianFrenchName,
		language.Dutch:                beaDutchName,
		language.French:               beaFrenchName,
		language.German:               beaGermanName,
		language.Italian:              beaItalianName,
		language.Japanese:             beaJapaneseName,
		language.Korean:               beaKoreanName,
		language.LatinAmericanSpanish: beaLatinAmericanSpanishName,
		language.Russian:              beaRussianName,
		language.SimplifiedChinese:    beaSimplifiedChineseName,
		language.Spanish:              beaSpanishName,
		language.TraditionalChinese:   beaTraditionalChineseName}
)

var (
	Bea = nook.Villager{
		Character:   beaCharacter,
		Personality: nook.Normal,
		Phrase:      beaPhrase}
)
