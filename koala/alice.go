package koala

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	aliceBirthday = nook.Birthday{
		Day:   19,
		Month: time.August}
)

var (
	aliceCode = nook.Code{
		Value: "kal01"}
)

var (
	aliceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Alice"}

	aliceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Alice"}

	aliceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Alice"}

	aliceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Alice"}

	aliceGermanName = nook.Name{
		Language: language.German,
		Value:    "Konny"}

	aliceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Alice"}

	aliceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メルボルン"}

	aliceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "멜버른"}

	aliceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Zelanda"}

	aliceRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Алиса"}

	aliceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莫儿"}

	aliceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Zelanda"}

	aliceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莫兒"}
)

var (
	aliceName = nook.Languages{
		language.AmericanEnglish:      aliceAmericanEnglishName,
		language.CanadianFrench:       aliceCanadianFrenchName,
		language.Dutch:                aliceDutchName,
		language.French:               aliceFrenchName,
		language.German:               aliceGermanName,
		language.Italian:              aliceItalianName,
		language.Japanese:             aliceJapaneseName,
		language.Korean:               aliceKoreanName,
		language.LatinAmericanSpanish: aliceLatinAmericanSpanishName,
		language.Russian:              aliceRussianName,
		language.SimplifiedChinese:    aliceSimplifiedChineseName,
		language.Spanish:              aliceSpanishName,
		language.TraditionalChinese:   aliceTraditionalChineseName}
)

var (
	aliceCharacter = nook.Character{
		Animal:   Koala,
		Birthday: aliceBirthday,
		Code:     aliceCode,
		Gender:   nook.Female,
		Name:     aliceName}
)

var (
	aliceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "guvnor"}

	aliceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chef"}

	aliceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "chef"}

	aliceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chef"}

	aliceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnarchhh"}

	aliceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tesorino"}

	aliceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キラリ"}

	aliceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "반짝"}

	aliceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "eucaliup"}

	aliceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "лидер"}

	aliceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "晶亮"}

	aliceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tú"}

	aliceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "晶亮"}
)

var (
	alicePhrase = nook.Languages{
		language.AmericanEnglish:      aliceAmericanEnglishName,
		language.CanadianFrench:       aliceCanadianFrenchName,
		language.Dutch:                aliceDutchName,
		language.French:               aliceFrenchName,
		language.German:               aliceGermanName,
		language.Italian:              aliceItalianName,
		language.Japanese:             aliceJapaneseName,
		language.Korean:               aliceKoreanName,
		language.LatinAmericanSpanish: aliceLatinAmericanSpanishName,
		language.Russian:              aliceRussianName,
		language.SimplifiedChinese:    aliceSimplifiedChineseName,
		language.Spanish:              aliceSpanishName,
		language.TraditionalChinese:   aliceTraditionalChineseName}
)

var (
	Alice = nook.Villager{
		Character:   aliceCharacter,
		Personality: nook.Normal,
		Phrase:      alicePhrase}
)
