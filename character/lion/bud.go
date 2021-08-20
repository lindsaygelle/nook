package lion

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
		Value:    "Léonard"}

	budDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bud"}

	budFrenchName = nook.Name{
		Language: language.French,
		Value:    "Léonard"}

	budGermanName = nook.Name{
		Language: language.German,
		Value:    "Dieter"}

	budItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Leo"}

	budJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グラさん"}

	budKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "선글"}

	budLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Surfleo"}

	budRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бад"}

	budSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莫敬"}

	budSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Surfleo"}

	budTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莫敬"}
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
		Animal:   animal.Lion,
		Birthday: budBirthday,
		Code:     budCode,
		Key:      character.Bud,
		Gender:   gender.Male,
		Name:     budName,
		Special:  false}
)

var (
	budAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "dood"}

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
		language.AmericanEnglish:      budAmericanEnglishPhrase,
		language.CanadianFrench:       budCanadianFrenchPhrase,
		language.Dutch:                budDutchPhrase,
		language.French:               budFrenchPhrase,
		language.German:               budGermanPhrase,
		language.Italian:              budItalianPhrase,
		language.Japanese:             budJapanesePhrase,
		language.Korean:               budKoreanPhrase,
		language.LatinAmericanSpanish: budLatinAmericanSpanishPhrase,
		language.Russian:              budRussianPhrase,
		language.SimplifiedChinese:    budSimplifiedChinesePhrase,
		language.Spanish:              budSpanishPhrase,
		language.TraditionalChinese:   budTraditionalChinesePhrase}
)

var (
	Bud = nook.Villager{
		Character:   budCharacter,
		Personality: personality.Jock,
		Phrase:      budPhrase}
)
