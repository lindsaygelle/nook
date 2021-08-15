package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	wolfgangBirthday = nook.Birthday{
		Day:   25,
		Month: time.November}
)

var (
	wolfgangCode = nook.Code{
		Value: "wol02"}
)

var (
	wolfgangAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wolfgang"}

	wolfgangCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Wolfgangsnurffff"}

	wolfgangDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wolfgangsnauw"}

	wolfgangFrenchName = nook.Name{
		Language: language.French,
		Value:    "Wolfgangsnurffff"}

	wolfgangGermanName = nook.Name{
		Language: language.German,
		Value:    "Weberknurrr"}

	wolfgangItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Wolfgangzanna"}

	wolfgangJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロボのな"}

	wolfgangKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "로보거봐"}

	wolfgangLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Wolfigrauuuh"}

	wolfgangRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Вольфиау-рр"}

	wolfgangSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗博罗罗"}

	wolfgangSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Wolfigrauuuh"}

	wolfgangTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅博羅羅"}
)

var (
	wolfgangName = nook.Languages{
		language.AmericanEnglish:      wolfgangAmericanEnglishName,
		language.CanadianFrench:       wolfgangCanadianFrenchName,
		language.Dutch:                wolfgangDutchName,
		language.French:               wolfgangFrenchName,
		language.German:               wolfgangGermanName,
		language.Italian:              wolfgangItalianName,
		language.Japanese:             wolfgangJapaneseName,
		language.Korean:               wolfgangKoreanName,
		language.LatinAmericanSpanish: wolfgangLatinAmericanSpanishName,
		language.Russian:              wolfgangRussianName,
		language.SimplifiedChinese:    wolfgangSimplifiedChineseName,
		language.Spanish:              wolfgangSpanishName,
		language.TraditionalChinese:   wolfgangTraditionalChineseName}
)

var (
	wolfgangCharacter = nook.Character{
		Animal:   Wolf,
		Birthday: wolfgangBirthday,
		Code:     wolfgangCode,
		Gender:   nook.Male,
		Name:     wolfgangName}
)

var (
	wolfgangAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snarrrl"}

	wolfgangCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "snurffff"}

	wolfgangDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snauw"}

	wolfgangFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "snurffff"}

	wolfgangGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knurrr"}

	wolfgangItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zanna"}

	wolfgangJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "のな"}

	wolfgangKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "거봐"}

	wolfgangLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grauuuh"}

	wolfgangRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ау-рр"}

	wolfgangSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗罗"}

	wolfgangSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grauuuh"}

	wolfgangTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅羅"}
)

var (
	wolfgangPhrase = nook.Languages{
		language.AmericanEnglish:      wolfgangAmericanEnglishName,
		language.CanadianFrench:       wolfgangCanadianFrenchName,
		language.Dutch:                wolfgangDutchName,
		language.French:               wolfgangFrenchName,
		language.German:               wolfgangGermanName,
		language.Italian:              wolfgangItalianName,
		language.Japanese:             wolfgangJapaneseName,
		language.Korean:               wolfgangKoreanName,
		language.LatinAmericanSpanish: wolfgangLatinAmericanSpanishName,
		language.Russian:              wolfgangRussianName,
		language.SimplifiedChinese:    wolfgangSimplifiedChineseName,
		language.Spanish:              wolfgangSpanishName,
		language.TraditionalChinese:   wolfgangTraditionalChineseName}
)

var (
	Wolfgang = nook.Villager{
		Character:   wolfgangCharacter,
		Personality: nook.Cranky,
		Phrase:      wolfgangPhrase}
)
