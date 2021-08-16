package tiger

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Vanina"}

	claudiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Claudia"}

	claudiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Vanina"}

	claudiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Lilly"}

	claudiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lilla"}

	claudiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マリリン"}

	claudiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "신디"}

	claudiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lilu"}

	claudiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клаудия"}

	claudiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "马丽莲"}

	claudiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lilu"}

	claudiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "馬麗蓮"}
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
		Animal:   animal.Tiger,
		Birthday: claudiaBirthday,
		Code:     claudiaCode,
		Gender:   gender.Female,
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
		Personality: personality.Snooty,
		Phrase:      claudiaPhrase}
)
