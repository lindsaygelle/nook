package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	clydeBirthday = nook.Birthday{
		Day:   1,
		Month: time.May}
)

var (
	clydeCode = nook.Code{
		Value: "hrs10"}
)

var (
	clydeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Clyde"}

	clydeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dorianmouh-ouh"}

	clydeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Clydeklipklop"}

	clydeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dorianmouh-ouh"}

	clydeGermanName = nook.Name{
		Language: language.German,
		Value:    "Tommischauder"}

	clydeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lalloclop clop"}

	clydeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "デースケぴろぴろ"}

	clydeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마철이다그닥"}

	clydeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Moeclip-clop"}

	clydeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клайдцок-цок"}

	clydeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "笛笛吹卷卷"}

	clydeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Moetroteras"}

	clydeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "笛笛吹捲捲"}
)

var (
	clydeName = nook.Languages{
		language.AmericanEnglish:      clydeAmericanEnglishName,
		language.CanadianFrench:       clydeCanadianFrenchName,
		language.Dutch:                clydeDutchName,
		language.French:               clydeFrenchName,
		language.German:               clydeGermanName,
		language.Italian:              clydeItalianName,
		language.Japanese:             clydeJapaneseName,
		language.Korean:               clydeKoreanName,
		language.LatinAmericanSpanish: clydeLatinAmericanSpanishName,
		language.Russian:              clydeRussianName,
		language.SimplifiedChinese:    clydeSimplifiedChineseName,
		language.Spanish:              clydeSpanishName,
		language.TraditionalChinese:   clydeTraditionalChineseName}
)

var (
	clydeCharacter = nook.Character{
		Animal:   Horse,
		Birthday: clydeBirthday,
		Code:     clydeCode,
		Gender:   nook.Male,
		Name:     clydeName}
)

var (
	clydeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "clip clawp"}

	clydeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mouh-ouh"}

	clydeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "klipklop"}

	clydeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mouh-ouh"}

	clydeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schauder"}

	clydeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "clop clop"}

	clydeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぴろぴろ"}

	clydeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "다그닥"}

	clydeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "clip-clop"}

	clydeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "цок-цок"}

	clydeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吹卷卷"}

	clydeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "troteras"}

	clydeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吹捲捲"}
)

var (
	clydePhrase = nook.Languages{
		language.AmericanEnglish:      clydeAmericanEnglishName,
		language.CanadianFrench:       clydeCanadianFrenchName,
		language.Dutch:                clydeDutchName,
		language.French:               clydeFrenchName,
		language.German:               clydeGermanName,
		language.Italian:              clydeItalianName,
		language.Japanese:             clydeJapaneseName,
		language.Korean:               clydeKoreanName,
		language.LatinAmericanSpanish: clydeLatinAmericanSpanishName,
		language.Russian:              clydeRussianName,
		language.SimplifiedChinese:    clydeSimplifiedChineseName,
		language.Spanish:              clydeSpanishName,
		language.TraditionalChinese:   clydeTraditionalChineseName}
)

var (
	Clyde = nook.Villager{
		Character:   clydeCharacter,
		Personality: nook.Lazy,
		Phrase:      clydePhrase}
)
