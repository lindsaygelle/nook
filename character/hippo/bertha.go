package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	berthaBirthday = nook.Birthday{
		Day:   25,
		Month: time.April}
)

var (
	berthaCode = nook.Code{
		Value: "hip03"}
)

var (
	berthaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bertha"}

	berthaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bertha"}

	berthaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bertha"}

	berthaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bertha"}

	berthaGermanName = nook.Name{
		Language: language.German,
		Value:    "Berta"}

	berthaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Umberta"}

	berthaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "あんこ"}

	berthaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "베티"}

	berthaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Berta"}

	berthaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Берта"}

	berthaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "豆沙"}

	berthaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Berta"}

	berthaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "豆沙"}
)

var (
	berthaName = nook.Languages{
		language.AmericanEnglish:      berthaAmericanEnglishName,
		language.CanadianFrench:       berthaCanadianFrenchName,
		language.Dutch:                berthaDutchName,
		language.French:               berthaFrenchName,
		language.German:               berthaGermanName,
		language.Italian:              berthaItalianName,
		language.Japanese:             berthaJapaneseName,
		language.Korean:               berthaKoreanName,
		language.LatinAmericanSpanish: berthaLatinAmericanSpanishName,
		language.Russian:              berthaRussianName,
		language.SimplifiedChinese:    berthaSimplifiedChineseName,
		language.Spanish:              berthaSpanishName,
		language.TraditionalChinese:   berthaTraditionalChineseName}
)

var (
	berthaCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: berthaBirthday,
		Code:     berthaCode,
		Gender:   gender.Female,
		Name:     berthaName}
)

var (
	berthaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bloop"}

	berthaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gloups"}

	berthaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "blubber"}

	berthaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "rololol"}

	berthaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "blubber"}

	berthaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blup"}

	berthaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "そうです"}

	berthaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "맞습니다"}

	berthaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "glub-glub"}

	berthaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бултых"}

	berthaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "没错"}

	berthaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "glub-glub"}

	berthaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "沒錯"}
)

var (
	berthaPhrase = nook.Languages{
		language.AmericanEnglish:      berthaAmericanEnglishName,
		language.CanadianFrench:       berthaCanadianFrenchName,
		language.Dutch:                berthaDutchName,
		language.French:               berthaFrenchName,
		language.German:               berthaGermanName,
		language.Italian:              berthaItalianName,
		language.Japanese:             berthaJapaneseName,
		language.Korean:               berthaKoreanName,
		language.LatinAmericanSpanish: berthaLatinAmericanSpanishName,
		language.Russian:              berthaRussianName,
		language.SimplifiedChinese:    berthaSimplifiedChineseName,
		language.Spanish:              berthaSpanishName,
		language.TraditionalChinese:   berthaTraditionalChineseName}
)

var (
	Bertha = nook.Villager{
		Character:   berthaCharacter,
		Personality: personality.Normal,
		Phrase:      berthaPhrase}
)
