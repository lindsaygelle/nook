package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Wolfgang"}

	wolfgangDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wolfgang"}

	wolfgangFrenchName = nook.Name{
		Language: language.French,
		Value:    "Wolfgang"}

	wolfgangGermanName = nook.Name{
		Language: language.German,
		Value:    "Weber"}

	wolfgangItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Wolfgang"}

	wolfgangJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロボ"}

	wolfgangKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "로보"}

	wolfgangLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Wolfi"}

	wolfgangRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Вольфи"}

	wolfgangSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗博"}

	wolfgangSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Wolfi"}

	wolfgangTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅博"}
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
		Animal:   animal.Wolf,
		Birthday: wolfgangBirthday,
		Code:     wolfgangCode,
		Gender:   gender.Male,
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
		Personality: personality.Cranky,
		Phrase:      wolfgangPhrase}
)
