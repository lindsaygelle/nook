package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Berthagloups"}

	berthaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Berthablubber"}

	berthaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bertharololol"}

	berthaGermanName = nook.Name{
		Language: language.German,
		Value:    "Bertablubber"}

	berthaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Umbertablup"}

	berthaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "あんこそうです"}

	berthaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "베티맞습니다"}

	berthaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bertaglub-glub"}

	berthaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бертабултых"}

	berthaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "豆沙没错"}

	berthaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bertaglub-glub"}

	berthaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "豆沙沒錯"}
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
		Animal:   Hippo,
		Birthday: berthaBirthday,
		Code:     berthaCode,
		Gender:   nook.Female,
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
		Personality: nook.Normal,
		Phrase:      berthaPhrase}
)
