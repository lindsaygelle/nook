package gorilla

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
		Value:    "Primo"}

	boydDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Boyd"}

	boydFrenchName = nook.Name{
		Language: language.French,
		Value:    "Primo"}

	boydGermanName = nook.Name{
		Language: language.German,
		Value:    "Boyd"}

	boydItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Brando"}

	boydJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボイド"}

	boydKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "보이드"}

	boydLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bunga"}

	boydRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бойд"}

	boydSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "空空"}

	boydSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bunga"}

	boydTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "空空"}
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
		Animal:   animal.Gorilla,
		Birthday: boydBirthday,
		Code:     boydCode,
		Key:      character.Boyd,
		Gender:   gender.Male,
		Name:     boydName,
		Special:  false}
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
		language.AmericanEnglish:      boydAmericanEnglishPhrase,
		language.CanadianFrench:       boydCanadianFrenchPhrase,
		language.Dutch:                boydDutchPhrase,
		language.French:               boydFrenchPhrase,
		language.German:               boydGermanPhrase,
		language.Italian:              boydItalianPhrase,
		language.Japanese:             boydJapanesePhrase,
		language.Korean:               boydKoreanPhrase,
		language.LatinAmericanSpanish: boydLatinAmericanSpanishPhrase,
		language.Russian:              boydRussianPhrase,
		language.SimplifiedChinese:    boydSimplifiedChinesePhrase,
		language.Spanish:              boydSpanishPhrase,
		language.TraditionalChinese:   boydTraditionalChinesePhrase}
)

var (
	Boyd = nook.Villager{
		Character:   boydCharacter,
		Personality: personality.Cranky,
		Phrase:      boydPhrase}
)
