package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	// wlinkBirthday represents W. Link's birthday (December 2nd).
	wlinkBirthday = nook.Birthday{
		Day:   2,
		Month: time.December}
)

var (
	// wlinkCode represents W. Link's unique code ("wol11").
	wlinkCode = nook.Code{
		Value: "wol11"}
)

var (
	// wlinkAmericanEnglishName represents W. Link's name in American English.
	wlinkAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "W. Link"}

	// wlinkCanadianFrenchName represents W. Link's name in Canadian French.
	wlinkCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Link l."}

	// wlinkDutchName represents W. Link's name in Dutch.
	wlinkDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// wlinkFrenchName represents W. Link's name in French.
	wlinkFrenchName = nook.Name{
		Language: language.French,
		Value:    "Link l."}

	// wlinkGermanName represents W. Link's name in German.
	wlinkGermanName = nook.Name{
		Language: language.German,
		Value:    "W. Link"}

	// wlinkItalianName represents W. Link's name in Italian.
	wlinkItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Link l."}

	// wlinkJapaneseName represents W. Link's name in Japanese.
	wlinkJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ウルフリンク"}

	// wlinkKoreanName represents W. Link's name in Korean.
	wlinkKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "울프 링크"}

	// wlinkLatinAmericanSpanishName represents W. Link's name in Latin American Spanish.
	wlinkLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Link L."}

	// wlinkRussianName represents W. Link's name in Russian.
	wlinkRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// wlinkSimplifiedChineseName represents W. Link's name in Simplified Chinese.
	wlinkSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// wlinkSpanishName represents W. Link's name in Spanish.
	wlinkSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Link L."}

	// wlinkTraditionalChineseName represents W. Link's name in Traditional Chinese.
	wlinkTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// wlinkName represents W. Link's name in different languages.
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
	// wlinkCharacter represents W. Link's character information.
	wlinkCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: wlinkBirthday,
		Code:     wlinkCode,
		Key:      character.WLink,
		Gender:   gender.Male,
		Name:     wlinkName,
		Special:  false}
)

var (
	// wlinkAmericanEnglishPhrase represents W. Link's phrase in American English.
	wlinkAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ruff ruff"}

	// wlinkCanadianFrenchPhrase represents W. Link's phrase in Canadian French.
	wlinkCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "crépuscule"}

	// wlinkDutchPhrase represents W. Link's phrase in Dutch.
	wlinkDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// wlinkFrenchPhrase represents W. Link's phrase in French.
	wlinkFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "crépuscule"}

	// wlinkGermanPhrase represents W. Link's phrase in German.
	wlinkGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnupper"}

	// wlinkItalianPhrase represents W. Link's phrase in Italian.
	wlinkItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ruff ruff"}

	// wlinkJapanesePhrase represents W. Link's phrase in Japanese.
	wlinkJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "くんくん"}

	// wlinkKoreanPhrase represents W. Link's phrase in Korean.
	wlinkKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아우우"}

	// wlinkLatinAmericanSpanishPhrase represents W. Link's phrase in Latin American Spanish.
	wlinkLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cachorro"}

	// wlinkRussianPhrase represents W. Link's phrase in Russian.
	wlinkRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// wlinkSimplifiedChinesePhrase represents W. Link's phrase in Simplified Chinese.
	wlinkSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// wlinkSpanishPhrase represents W. Link's phrase in Spanish.
	wlinkSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cachorro"}

	// wlinkTraditionalChinesePhrase represents W. Link's phrase in Traditional Chinese.
	wlinkTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// wlinkPhrase represents W. Link's phrases in different languages.
	wlinkPhrase = nook.Languages{
		language.AmericanEnglish:      wlinkAmericanEnglishPhrase,
		language.CanadianFrench:       wlinkCanadianFrenchPhrase,
		language.Dutch:                wlinkDutchPhrase,
		language.French:               wlinkFrenchPhrase,
		language.German:               wlinkGermanPhrase,
		language.Italian:              wlinkItalianPhrase,
		language.Japanese:             wlinkJapanesePhrase,
		language.Korean:               wlinkKoreanPhrase,
		language.LatinAmericanSpanish: wlinkLatinAmericanSpanishPhrase,
		language.Russian:              wlinkRussianPhrase,
		language.SimplifiedChinese:    wlinkSimplifiedChinesePhrase,
		language.Spanish:              wlinkSpanishPhrase,
		language.TraditionalChinese:   wlinkTraditionalChinesePhrase}
)

var (
	// WLink represents the character W. Link with his complete information.
	WLink = nook.Villager{
		Character:   wlinkCharacter,
		Personality: personality.Smug,
		Phrase:      wlinkPhrase}
)
