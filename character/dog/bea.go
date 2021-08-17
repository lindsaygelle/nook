package dog

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
		Value:    "Béa"}

	beaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bea"}

	beaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Béa"}

	beaGermanName = nook.Name{
		Language: language.German,
		Value:    "Bea"}

	beaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cucciola"}

	beaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ベーグル"}

	beaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "베이글"}

	beaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bea"}

	beaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Беа"}

	beaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贝果"}

	beaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bea"}

	beaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "貝果"}
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
		Animal:   animal.Dog,
		Birthday: beaBirthday,
		Code:     beaCode,
		Key:      character.Bea,
		Gender:   gender.Female,
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
		language.AmericanEnglish:      beaAmericanEnglishPhrase,
		language.CanadianFrench:       beaCanadianFrenchPhrase,
		language.Dutch:                beaDutchPhrase,
		language.French:               beaFrenchPhrase,
		language.German:               beaGermanPhrase,
		language.Italian:              beaItalianPhrase,
		language.Japanese:             beaJapanesePhrase,
		language.Korean:               beaKoreanPhrase,
		language.LatinAmericanSpanish: beaLatinAmericanSpanishPhrase,
		language.Russian:              beaRussianPhrase,
		language.SimplifiedChinese:    beaSimplifiedChinesePhrase,
		language.Spanish:              beaSpanishPhrase,
		language.TraditionalChinese:   beaTraditionalChinesePhrase}
)

var (
	Bea = nook.Villager{
		Character:   beaCharacter,
		Personality: personality.Normal,
		Phrase:      beaPhrase}
)
