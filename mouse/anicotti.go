package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	anicottiBirthday = nook.Birthday{
		Day:   24,
		Month: time.February}
)

var (
	anicottiCode = nook.Code{
		Value: "mus10"}
)

var (
	anicottiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Anicotti"}

	anicottiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Anniedis donc"}

	anicottiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Anicotticannoli"}

	anicottiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Anniedis donc"}

	anicottiGermanName = nook.Name{
		Language: language.German,
		Value:    "Evapiepser"}

	anicottiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Squittacippoli"}

	anicottiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラザニアルンルン"}

	anicottiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "라자냐룰룰"}

	anicottiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Clorindacloricló"}

	anicottiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аникоттиканноли"}

	anicottiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗萱儿开心"}

	anicottiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Clorindacloricló"}

	anicottiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅萱兒開心"}
)

var (
	anicottiName = nook.Languages{
		language.AmericanEnglish:      anicottiAmericanEnglishName,
		language.CanadianFrench:       anicottiCanadianFrenchName,
		language.Dutch:                anicottiDutchName,
		language.French:               anicottiFrenchName,
		language.German:               anicottiGermanName,
		language.Italian:              anicottiItalianName,
		language.Japanese:             anicottiJapaneseName,
		language.Korean:               anicottiKoreanName,
		language.LatinAmericanSpanish: anicottiLatinAmericanSpanishName,
		language.Russian:              anicottiRussianName,
		language.SimplifiedChinese:    anicottiSimplifiedChineseName,
		language.Spanish:              anicottiSpanishName,
		language.TraditionalChinese:   anicottiTraditionalChineseName}
)

var (
	anicottiCharacter = nook.Character{
		Animal:   Mouse,
		Birthday: anicottiBirthday,
		Code:     anicottiCode,
		Gender:   nook.Female,
		Name:     anicottiName}
)

var (
	anicottiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cannoli"}

	anicottiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "dis donc"}

	anicottiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "cannoli"}

	anicottiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "dis donc"}

	anicottiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "piepser"}

	anicottiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cippoli"}

	anicottiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ルンルン"}

	anicottiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "룰룰"}

	anicottiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cloricló"}

	anicottiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "канноли"}

	anicottiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "开心"}

	anicottiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cloricló"}

	anicottiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "開心"}
)

var (
	anicottiPhrase = nook.Languages{
		language.AmericanEnglish:      anicottiAmericanEnglishName,
		language.CanadianFrench:       anicottiCanadianFrenchName,
		language.Dutch:                anicottiDutchName,
		language.French:               anicottiFrenchName,
		language.German:               anicottiGermanName,
		language.Italian:              anicottiItalianName,
		language.Japanese:             anicottiJapaneseName,
		language.Korean:               anicottiKoreanName,
		language.LatinAmericanSpanish: anicottiLatinAmericanSpanishName,
		language.Russian:              anicottiRussianName,
		language.SimplifiedChinese:    anicottiSimplifiedChineseName,
		language.Spanish:              anicottiSpanishName,
		language.TraditionalChinese:   anicottiTraditionalChineseName}
)

var (
	Anicotti = nook.Villager{
		Character:   anicottiCharacter,
		Personality: nook.Peppy,
		Phrase:      anicottiPhrase}
)
