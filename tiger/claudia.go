package tiger

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	claudiaBirthday = nook.Birthday{
		Day:   22,
		Month: time.November}
)

var (
	claudiaCode = nook.Code{
		Value: "tig05"}
)

var (
	claudiaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Claudia"}

	claudiaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Vaninagrou grou"}

	claudiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Claudiaolala"}

	claudiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Vaninagrou grou"}

	claudiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Lillykchhh"}

	claudiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lillasmack"}

	claudiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マリリンアハ～ン"}

	claudiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "신디잇힝"}

	claudiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lilumindundi"}

	claudiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клаудияо-ля-ля"}

	claudiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "马丽莲啊哈"}

	claudiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lilumindundi"}

	claudiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "馬麗蓮啊哈"}
)

var (
	claudiaName = nook.Languages{
		language.AmericanEnglish:      claudiaAmericanEnglishName,
		language.CanadianFrench:       claudiaCanadianFrenchName,
		language.Dutch:                claudiaDutchName,
		language.French:               claudiaFrenchName,
		language.German:               claudiaGermanName,
		language.Italian:              claudiaItalianName,
		language.Japanese:             claudiaJapaneseName,
		language.Korean:               claudiaKoreanName,
		language.LatinAmericanSpanish: claudiaLatinAmericanSpanishName,
		language.Russian:              claudiaRussianName,
		language.SimplifiedChinese:    claudiaSimplifiedChineseName,
		language.Spanish:              claudiaSpanishName,
		language.TraditionalChinese:   claudiaTraditionalChineseName}
)

var (
	claudiaCharacter = nook.Character{
		Animal:   Tiger,
		Birthday: claudiaBirthday,
		Code:     claudiaCode,
		Gender:   nook.Female,
		Name:     claudiaName}
)

var (
	claudiaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ooh la la"}

	claudiaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grou grou"}

	claudiaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "olala"}

	claudiaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grou grou"}

	claudiaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kchhh"}

	claudiaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "smack"}

	claudiaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アハ～ン"}

	claudiaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "잇힝"}

	claudiaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mindundi"}

	claudiaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "о-ля-ля"}

	claudiaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啊哈"}

	claudiaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mindundi"}

	claudiaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啊哈"}
)

var (
	claudiaPhrase = nook.Languages{
		language.AmericanEnglish:      claudiaAmericanEnglishName,
		language.CanadianFrench:       claudiaCanadianFrenchName,
		language.Dutch:                claudiaDutchName,
		language.French:               claudiaFrenchName,
		language.German:               claudiaGermanName,
		language.Italian:              claudiaItalianName,
		language.Japanese:             claudiaJapaneseName,
		language.Korean:               claudiaKoreanName,
		language.LatinAmericanSpanish: claudiaLatinAmericanSpanishName,
		language.Russian:              claudiaRussianName,
		language.SimplifiedChinese:    claudiaSimplifiedChineseName,
		language.Spanish:              claudiaSpanishName,
		language.TraditionalChinese:   claudiaTraditionalChineseName}
)

var (
	Claudia = nook.Villager{
		Character:   claudiaCharacter,
		Personality: nook.Snooty,
		Phrase:      claudiaPhrase}
)
