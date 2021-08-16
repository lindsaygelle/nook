package bearcub

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
		Value:    "Marty"}

	martyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marty"}

	martyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marty"}

	martyGermanName = nook.Name{
		Language: language.German,
		Value:    "Marty"}

	martyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marty"}

	martyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マーティー"}

	martyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마티"}

	martyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marty"}

	martyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Марти"}

	martySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玛丁"}

	martySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marty"}

	martyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑪丁"}
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
		Animal:   animal.Bearcub,
		Birthday: martyBirthday,
		Code:     martyCode,
		Key:      character.Marty,
		Gender:   gender.Male,
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
		Personality: personality.Lazy,
		Phrase:      martyPhrase}
)
