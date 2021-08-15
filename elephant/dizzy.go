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
		Value:    "Pachywoutchou"}

	dizzyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dizzytrompet"}

	dizzyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pachywoutchou"}

	dizzyGermanName = nook.Name{
		Language: language.German,
		Value:    "Bastiantaraaaa"}

	dizzyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Annibaleullà"}

	dizzyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヒュージだゾウ"}

	dizzyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "휴지휴휴"}

	dizzyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Quiquewuuulá"}

	dizzyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Диззитру-ру"}

	dizzySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巨巨象象"}

	dizzySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Quiquenaricita"}

	dizzyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巨巨象象"}
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
