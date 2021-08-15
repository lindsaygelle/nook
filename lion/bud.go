package lion

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	budBirthday = nook.Birthday{
		Day:   8,
		Month: time.August}
)

var (
	budCode = nook.Code{
		Value: "lon00"}
)

var (
	budAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bud"}

	budCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Léonardmon man"}

	budDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Budmaaaan"}

	budFrenchName = nook.Name{
		Language: language.French,
		Value:    "Léonardm'pote"}

	budGermanName = nook.Name{
		Language: language.German,
		Value:    "Dieterkumpel"}

	budItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Leoruggi"}

	budJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グラさんメ～ン"}

	budKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "선글요～맨"}

	budLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Surfleogrrrau"}

	budRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бадче-е-ел"}

	budSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莫敬男人"}

	budSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Surfleosabes"}

	budTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莫敬男人"}
)

var (
	budName = nook.Languages{
		language.AmericanEnglish:      budAmericanEnglishName,
		language.CanadianFrench:       budCanadianFrenchName,
		language.Dutch:                budDutchName,
		language.French:               budFrenchName,
		language.German:               budGermanName,
		language.Italian:              budItalianName,
		language.Japanese:             budJapaneseName,
		language.Korean:               budKoreanName,
		language.LatinAmericanSpanish: budLatinAmericanSpanishName,
		language.Russian:              budRussianName,
		language.SimplifiedChinese:    budSimplifiedChineseName,
		language.Spanish:              budSpanishName,
		language.TraditionalChinese:   budTraditionalChineseName}
)

var (
	budCharacter = nook.Character{
		Animal:   Lion,
		Birthday: budBirthday,
		Code:     budCode,
		Gender:   nook.Male,
		Name:     budName}
)

var (
	budAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "doodmaaanshredded"}

	budCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon man"}

	budDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "maaaan"}

	budFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "m'pote"}

	budGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kumpel"}

	budItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ruggi"}

	budJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "メ～ン"}

	budKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "요～맨"}

	budLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrrau"}

	budRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "че-е-ел"}

	budSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "男人"}

	budSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sabes"}

	budTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "男人"}
)

var (
	budPhrase = nook.Languages{
		language.AmericanEnglish:      budAmericanEnglishName,
		language.CanadianFrench:       budCanadianFrenchName,
		language.Dutch:                budDutchName,
		language.French:               budFrenchName,
		language.German:               budGermanName,
		language.Italian:              budItalianName,
		language.Japanese:             budJapaneseName,
		language.Korean:               budKoreanName,
		language.LatinAmericanSpanish: budLatinAmericanSpanishName,
		language.Russian:              budRussianName,
		language.SimplifiedChinese:    budSimplifiedChineseName,
		language.Spanish:              budSpanishName,
		language.TraditionalChinese:   budTraditionalChineseName}
)

var (
	Bud = nook.Villager{
		Character:   budCharacter,
		Personality: nook.Jock,
		Phrase:      budPhrase}
)
