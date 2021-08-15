package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	juneBirthday = nook.Birthday{
		Day:   21,
		Month: time.May}
)

var (
	juneCode = nook.Code{
		Value: "cbr13"}
)

var (
	juneAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "June"}

	juneCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Agnèscacou"}

	juneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Junewolkje"}

	juneFrenchName = nook.Name{
		Language: language.French,
		Value:    "Agnèscacou"}

	juneGermanName = nook.Name{
		Language: language.German,
		Value:    "Junaschönheit"}

	juneItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Junegrrande"}

	juneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メイアルネ"}

	juneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "메이있죠"}

	juneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hanalulúlurilá"}

	juneRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джунрадуга-дуга"}

	juneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿妹有耶"}

	juneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hanalulúcielo"}

	juneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿妹有耶"}
)

var (
	juneName = nook.Languages{
		language.AmericanEnglish:      juneAmericanEnglishName,
		language.CanadianFrench:       juneCanadianFrenchName,
		language.Dutch:                juneDutchName,
		language.French:               juneFrenchName,
		language.German:               juneGermanName,
		language.Italian:              juneItalianName,
		language.Japanese:             juneJapaneseName,
		language.Korean:               juneKoreanName,
		language.LatinAmericanSpanish: juneLatinAmericanSpanishName,
		language.Russian:              juneRussianName,
		language.SimplifiedChinese:    juneSimplifiedChineseName,
		language.Spanish:              juneSpanishName,
		language.TraditionalChinese:   juneTraditionalChineseName}
)

var (
	juneCharacter = nook.Character{
		Animal:   Bearcub,
		Birthday: juneBirthday,
		Code:     juneCode,
		Gender:   nook.Female,
		Name:     juneName}
)

var (
	juneAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "rainbow"}

	juneCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cacou"}

	juneDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wolkje"}

	juneFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cacou"}

	juneGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schönheit"}

	juneItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrande"}

	juneJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アルネ"}

	juneKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "있죠"}

	juneLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "lurilá"}

	juneRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "радуга-дуга"}

	juneSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "有耶"}

	juneSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cielo"}

	juneTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "有耶"}
)

var (
	junePhrase = nook.Languages{
		language.AmericanEnglish:      juneAmericanEnglishName,
		language.CanadianFrench:       juneCanadianFrenchName,
		language.Dutch:                juneDutchName,
		language.French:               juneFrenchName,
		language.German:               juneGermanName,
		language.Italian:              juneItalianName,
		language.Japanese:             juneJapaneseName,
		language.Korean:               juneKoreanName,
		language.LatinAmericanSpanish: juneLatinAmericanSpanishName,
		language.Russian:              juneRussianName,
		language.SimplifiedChinese:    juneSimplifiedChineseName,
		language.Spanish:              juneSpanishName,
		language.TraditionalChinese:   juneTraditionalChineseName}
)

var (
	June = nook.Villager{
		Character:   juneCharacter,
		Personality: nook.Normal,
		Phrase:      junePhrase}
)
