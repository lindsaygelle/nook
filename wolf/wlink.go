package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	wlinkBirthday = nook.Birthday{
		Day:   2,
		Month: time.December}
)

var (
	wlinkCode = nook.Code{
		Value: "wol11"}
)

var (
	wlinkAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "W. Link"}

	wlinkCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Link l."}

	wlinkDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	wlinkFrenchName = nook.Name{
		Language: language.French,
		Value:    "Link l."}

	wlinkGermanName = nook.Name{
		Language: language.German,
		Value:    "W. Link"}

	wlinkItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Link l."}

	wlinkJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ウルフリンク"}

	wlinkKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "울프 링크"}

	wlinkLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Link L."}

	wlinkRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	wlinkSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	wlinkSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Link L."}

	wlinkTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	wlinkName = nook.Languages{
		language.AmericanEnglish:      wlinkAmericanEnglishName,
		language.CanadianFrench:       wlinkCanadianFrenchName,
		language.Dutch:                wlinkDutchName,
		language.French:               wlinkFrenchName,
		language.German:               wlinkGermanName,
		language.Italian:              wlinkItalianName,
		language.Japanese:             wlinkJapaneseName,
		language.Korean:               wlinkKoreanName,
		language.LatinAmericanSpanish: wlinkLatinAmericanSpanishName,
		language.Russian:              wlinkRussianName,
		language.SimplifiedChinese:    wlinkSimplifiedChineseName,
		language.Spanish:              wlinkSpanishName,
		language.TraditionalChinese:   wlinkTraditionalChineseName}
)

var (
	wlinkCharacter = nook.Character{
		Animal:   Wolf,
		Birthday: wlinkBirthday,
		Code:     wlinkCode,
		Gender:   nook.Male,
		Name:     wlinkName}
)

var (
	wlinkAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ruff ruff"}

	wlinkCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "crépuscule"}

	wlinkDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	wlinkFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "crépuscule"}

	wlinkGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnupper"}

	wlinkItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ruff ruff"}

	wlinkJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "くんくん"}

	wlinkKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아우우"}

	wlinkLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cachorro"}

	wlinkRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	wlinkSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	wlinkSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cachorro"}

	wlinkTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	wlinkPhrase = nook.Languages{
		language.AmericanEnglish:      wlinkAmericanEnglishName,
		language.CanadianFrench:       wlinkCanadianFrenchName,
		language.Dutch:                wlinkDutchName,
		language.French:               wlinkFrenchName,
		language.German:               wlinkGermanName,
		language.Italian:              wlinkItalianName,
		language.Japanese:             wlinkJapaneseName,
		language.Korean:               wlinkKoreanName,
		language.LatinAmericanSpanish: wlinkLatinAmericanSpanishName,
		language.Russian:              wlinkRussianName,
		language.SimplifiedChinese:    wlinkSimplifiedChineseName,
		language.Spanish:              wlinkSpanishName,
		language.TraditionalChinese:   wlinkTraditionalChineseName}
)

var (
	WLink = nook.Villager{
		Character:   wlinkCharacter,
		Personality: nook.Smug,
		Phrase:      wlinkPhrase}
)
