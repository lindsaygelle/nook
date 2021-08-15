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
		Value:    "Mariloumumuseau"}

	skyeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Skyeluchtpost"}

	skyeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mariloumumuseau"}

	skyeGermanName = nook.Name{
		Language: language.German,
		Value:    "Sabineahiii"}

	skyeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lupillauuuhlalà"}

	skyeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リリィフワワ"}

	skyeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "릴리후와와"}

	skyeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Alderiaauuu"}

	skyeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Скайпилот"}

	skyeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "百合花花"}

	skyeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Alderiaauuu"}

	skyeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "百合花花"}
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
