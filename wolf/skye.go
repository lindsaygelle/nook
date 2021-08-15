package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	skyeBirthday = nook.Birthday{
		Day:   24,
		Month: time.March}
)

var (
	skyeCode = nook.Code{
		Value: "wol09"}
)

var (
	skyeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Skye"}

	skyeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marilou"}

	skyeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Skye"}

	skyeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marilou"}

	skyeGermanName = nook.Name{
		Language: language.German,
		Value:    "Sabine"}

	skyeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lupilla"}

	skyeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リリィ"}

	skyeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "릴리"}

	skyeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Alderia"}

	skyeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Скай"}

	skyeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "百合"}

	skyeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Alderia"}

	skyeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "百合"}
)

var (
	skyeName = nook.Languages{
		language.AmericanEnglish:      skyeAmericanEnglishName,
		language.CanadianFrench:       skyeCanadianFrenchName,
		language.Dutch:                skyeDutchName,
		language.French:               skyeFrenchName,
		language.German:               skyeGermanName,
		language.Italian:              skyeItalianName,
		language.Japanese:             skyeJapaneseName,
		language.Korean:               skyeKoreanName,
		language.LatinAmericanSpanish: skyeLatinAmericanSpanishName,
		language.Russian:              skyeRussianName,
		language.SimplifiedChinese:    skyeSimplifiedChineseName,
		language.Spanish:              skyeSpanishName,
		language.TraditionalChinese:   skyeTraditionalChineseName}
)

var (
	skyeCharacter = nook.Character{
		Animal:   Wolf,
		Birthday: skyeBirthday,
		Code:     skyeCode,
		Gender:   nook.Female,
		Name:     skyeName}
)

var (
	skyeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "airmail"}

	skyeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mumuseau"}

	skyeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "luchtpost"}

	skyeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mumuseau"}

	skyeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ahiii"}

	skyeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uuuhlalà"}

	skyeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "フワワ"}

	skyeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "후와와"}

	skyeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "auuu"}

	skyeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пилот"}

	skyeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "花花"}

	skyeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "auuu"}

	skyeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "花花"}
)

var (
	skyePhrase = nook.Languages{
		language.AmericanEnglish:      skyeAmericanEnglishName,
		language.CanadianFrench:       skyeCanadianFrenchName,
		language.Dutch:                skyeDutchName,
		language.French:               skyeFrenchName,
		language.German:               skyeGermanName,
		language.Italian:              skyeItalianName,
		language.Japanese:             skyeJapaneseName,
		language.Korean:               skyeKoreanName,
		language.LatinAmericanSpanish: skyeLatinAmericanSpanishName,
		language.Russian:              skyeRussianName,
		language.SimplifiedChinese:    skyeSimplifiedChineseName,
		language.Spanish:              skyeSpanishName,
		language.TraditionalChinese:   skyeTraditionalChineseName}
)

var (
	Skye = nook.Villager{
		Character:   skyeCharacter,
		Personality: nook.Normal,
		Phrase:      skyePhrase}
)
