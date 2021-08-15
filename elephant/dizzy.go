package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	dizzyBirthday = nook.Birthday{
		Day:   14,
		Month: time.July}
)

var (
	dizzyCode = nook.Code{
		Value: "elp01"}
)

var (
	dizzyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dizzy"}

	dizzyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pachy"}

	dizzyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dizzy"}

	dizzyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pachy"}

	dizzyGermanName = nook.Name{
		Language: language.German,
		Value:    "Bastian"}

	dizzyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Annibale"}

	dizzyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヒュージ"}

	dizzyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "휴지"}

	dizzyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Quique"}

	dizzyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Диззи"}

	dizzySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巨巨"}

	dizzySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Quique"}

	dizzyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巨巨"}
)

var (
	dizzyName = nook.Languages{
		language.AmericanEnglish:      dizzyAmericanEnglishName,
		language.CanadianFrench:       dizzyCanadianFrenchName,
		language.Dutch:                dizzyDutchName,
		language.French:               dizzyFrenchName,
		language.German:               dizzyGermanName,
		language.Italian:              dizzyItalianName,
		language.Japanese:             dizzyJapaneseName,
		language.Korean:               dizzyKoreanName,
		language.LatinAmericanSpanish: dizzyLatinAmericanSpanishName,
		language.Russian:              dizzyRussianName,
		language.SimplifiedChinese:    dizzySimplifiedChineseName,
		language.Spanish:              dizzySpanishName,
		language.TraditionalChinese:   dizzyTraditionalChineseName}
)

var (
	dizzyCharacter = nook.Character{
		Animal:   Elephant,
		Birthday: dizzyBirthday,
		Code:     dizzyCode,
		Gender:   nook.Male,
		Name:     dizzyName}
)

var (
	dizzyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "woo-oo"}

	dizzyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "woutchou"}

	dizzyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "trompet"}

	dizzyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "woutchou"}

	dizzyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "taraaaa"}

	dizzyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ullà"}

	dizzyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だゾウ"}

	dizzyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "휴휴"}

	dizzyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "wuuulá"}

	dizzyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тру-ру"}

	dizzySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "象象"}

	dizzySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "naricita"}

	dizzyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "象象"}
)

var (
	dizzyPhrase = nook.Languages{
		language.AmericanEnglish:      dizzyAmericanEnglishName,
		language.CanadianFrench:       dizzyCanadianFrenchName,
		language.Dutch:                dizzyDutchName,
		language.French:               dizzyFrenchName,
		language.German:               dizzyGermanName,
		language.Italian:              dizzyItalianName,
		language.Japanese:             dizzyJapaneseName,
		language.Korean:               dizzyKoreanName,
		language.LatinAmericanSpanish: dizzyLatinAmericanSpanishName,
		language.Russian:              dizzyRussianName,
		language.SimplifiedChinese:    dizzySimplifiedChineseName,
		language.Spanish:              dizzySpanishName,
		language.TraditionalChinese:   dizzyTraditionalChineseName}
)

var (
	Dizzy = nook.Villager{
		Character:   dizzyCharacter,
		Personality: nook.Lazy,
		Phrase:      dizzyPhrase}
)
