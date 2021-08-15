package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	driftBirthday = nook.Birthday{
		Day:   9,
		Month: time.October}
)

var (
	driftCode = nook.Code{
		Value: "flg04"}
)

var (
	driftAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Drift"}

	driftCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gordonbraaaa"}

	driftDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Driftgozer"}

	driftFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gordonbraaaa"}

	driftGermanName = nook.Name{
		Language: language.German,
		Value:    "Casparquak"}

	driftItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Prosperogirin"}

	driftJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドクゲロゲロ"}

	driftKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "덕개굴개굴"}

	driftLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Surfíndribit"}

	driftRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дрифтбро"}

	driftSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毒仔聒聒"}

	driftSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Surfíncolega"}

	driftTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毒仔嘓嘓"}
)

var (
	driftName = nook.Languages{
		language.AmericanEnglish:      driftAmericanEnglishName,
		language.CanadianFrench:       driftCanadianFrenchName,
		language.Dutch:                driftDutchName,
		language.French:               driftFrenchName,
		language.German:               driftGermanName,
		language.Italian:              driftItalianName,
		language.Japanese:             driftJapaneseName,
		language.Korean:               driftKoreanName,
		language.LatinAmericanSpanish: driftLatinAmericanSpanishName,
		language.Russian:              driftRussianName,
		language.SimplifiedChinese:    driftSimplifiedChineseName,
		language.Spanish:              driftSpanishName,
		language.TraditionalChinese:   driftTraditionalChineseName}
)

var (
	driftCharacter = nook.Character{
		Animal:   Frog,
		Birthday: driftBirthday,
		Code:     driftCode,
		Gender:   nook.Male,
		Name:     driftName}
)

var (
	driftAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "brahdribbit"}

	driftCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "braaaa"}

	driftDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gozer"}

	driftFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "braaaa"}

	driftGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quak"}

	driftItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "girin"}

	driftJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ゲロゲロ"}

	driftKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "개굴개굴"}

	driftLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "dribit"}

	driftRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бро"}

	driftSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "聒聒"}

	driftSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "colega"}

	driftTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘓嘓"}
)

var (
	driftPhrase = nook.Languages{
		language.AmericanEnglish:      driftAmericanEnglishName,
		language.CanadianFrench:       driftCanadianFrenchName,
		language.Dutch:                driftDutchName,
		language.French:               driftFrenchName,
		language.German:               driftGermanName,
		language.Italian:              driftItalianName,
		language.Japanese:             driftJapaneseName,
		language.Korean:               driftKoreanName,
		language.LatinAmericanSpanish: driftLatinAmericanSpanishName,
		language.Russian:              driftRussianName,
		language.SimplifiedChinese:    driftSimplifiedChineseName,
		language.Spanish:              driftSpanishName,
		language.TraditionalChinese:   driftTraditionalChineseName}
)

var (
	Drift = nook.Villager{
		Character:   driftCharacter,
		Personality: nook.Jock,
		Phrase:      driftPhrase}
)
