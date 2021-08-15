package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	martyBirthday = nook.Birthday{
		Day:   16,
		Month: time.April}
)

var (
	martyCode = nook.Code{
		Value: "cbr18"}
)

var (
	martyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Marty"}

	martyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Martypom pom"}

	martyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Martypommetje"}

	martyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Martypom pom"}

	martyGermanName = nook.Name{
		Language: language.German,
		Value:    "Martypompudding"}

	martyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Martypom pom"}

	martyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マーティーポムッ"}

	martyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마티폼폼"}

	martyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Martypudin"}

	martyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мартипомпончик"}

	martySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玛丁布丁"}

	martySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Martypudin"}

	martyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑪丁布丁"}
)

var (
	martyName = nook.Languages{
		language.AmericanEnglish:      martyAmericanEnglishName,
		language.CanadianFrench:       martyCanadianFrenchName,
		language.Dutch:                martyDutchName,
		language.French:               martyFrenchName,
		language.German:               martyGermanName,
		language.Italian:              martyItalianName,
		language.Japanese:             martyJapaneseName,
		language.Korean:               martyKoreanName,
		language.LatinAmericanSpanish: martyLatinAmericanSpanishName,
		language.Russian:              martyRussianName,
		language.SimplifiedChinese:    martySimplifiedChineseName,
		language.Spanish:              martySpanishName,
		language.TraditionalChinese:   martyTraditionalChineseName}
)

var (
	martyCharacter = nook.Character{
		Animal:   Bearcub,
		Birthday: martyBirthday,
		Code:     martyCode,
		Gender:   nook.Male,
		Name:     martyName}
)

var (
	martyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pompom"}

	martyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pom pom"}

	martyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pommetje"}

	martyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pom pom"}

	martyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pompudding"}

	martyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pom pom"}

	martyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ポムッ"}

	martyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "폼폼"}

	martyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pudin"}

	martyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "помпончик"}

	martySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "布丁"}

	martySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pudin"}

	martyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "布丁"}
)

var (
	martyPhrase = nook.Languages{
		language.AmericanEnglish:      martyAmericanEnglishName,
		language.CanadianFrench:       martyCanadianFrenchName,
		language.Dutch:                martyDutchName,
		language.French:               martyFrenchName,
		language.German:               martyGermanName,
		language.Italian:              martyItalianName,
		language.Japanese:             martyJapaneseName,
		language.Korean:               martyKoreanName,
		language.LatinAmericanSpanish: martyLatinAmericanSpanishName,
		language.Russian:              martyRussianName,
		language.SimplifiedChinese:    martySimplifiedChineseName,
		language.Spanish:              martySpanishName,
		language.TraditionalChinese:   martyTraditionalChineseName}
)

var (
	Marty = nook.Villager{
		Character:   martyCharacter,
		Personality: nook.Lazy,
		Phrase:      martyPhrase}
)
