package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	boydBirthday = nook.Birthday{
		Day:   1,
		Month: time.October}
)

var (
	boydCode = nook.Code{
		Value: "gor05"}
)

var (
	boydAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Boyd"}

	boydCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Primopelage"}

	boydDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Boydoh-oh"}

	boydFrenchName = nook.Name{
		Language: language.French,
		Value:    "Primopelage"}

	boydGermanName = nook.Name{
		Language: language.German,
		Value:    "Boydäffäff"}

	boydItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Brandobam bam"}

	boydJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボイドおうおう"}

	boydKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "보이드오우오우"}

	boydLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bungaunga"}

	boydRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бойдух-ох"}

	boydSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "空空喔喔"}

	boydSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bungaunga"}

	boydTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "空空喔喔"}
)

var (
	boydName = nook.Languages{
		language.AmericanEnglish:      boydAmericanEnglishName,
		language.CanadianFrench:       boydCanadianFrenchName,
		language.Dutch:                boydDutchName,
		language.French:               boydFrenchName,
		language.German:               boydGermanName,
		language.Italian:              boydItalianName,
		language.Japanese:             boydJapaneseName,
		language.Korean:               boydKoreanName,
		language.LatinAmericanSpanish: boydLatinAmericanSpanishName,
		language.Russian:              boydRussianName,
		language.SimplifiedChinese:    boydSimplifiedChineseName,
		language.Spanish:              boydSpanishName,
		language.TraditionalChinese:   boydTraditionalChineseName}
)

var (
	boydCharacter = nook.Character{
		Animal:   Gorilla,
		Birthday: boydBirthday,
		Code:     boydCode,
		Gender:   nook.Male,
		Name:     boydName}
)

var (
	boydAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "uh-oh"}

	boydCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pelage"}

	boydDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oh-oh"}

	boydFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pelage"}

	boydGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "äffäff"}

	boydItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bam bam"}

	boydJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "おうおう"}

	boydKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오우오우"}

	boydLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "unga"}

	boydRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ух-ох"}

	boydSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喔喔"}

	boydSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "unga"}

	boydTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喔喔"}
)

var (
	boydPhrase = nook.Languages{
		language.AmericanEnglish:      boydAmericanEnglishName,
		language.CanadianFrench:       boydCanadianFrenchName,
		language.Dutch:                boydDutchName,
		language.French:               boydFrenchName,
		language.German:               boydGermanName,
		language.Italian:              boydItalianName,
		language.Japanese:             boydJapaneseName,
		language.Korean:               boydKoreanName,
		language.LatinAmericanSpanish: boydLatinAmericanSpanishName,
		language.Russian:              boydRussianName,
		language.SimplifiedChinese:    boydSimplifiedChineseName,
		language.Spanish:              boydSpanishName,
		language.TraditionalChinese:   boydTraditionalChineseName}
)

var (
	Boyd = nook.Villager{
		Character:   boydCharacter,
		Personality: nook.Cranky,
		Phrase:      boydPhrase}
)
