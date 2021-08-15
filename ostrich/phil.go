package ostrich

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	philBirthday = nook.Birthday{
		Day:   27,
		Month: time.November}
)

var (
	philCode = nook.Code{
		Value: "ost07"}
)

var (
	philAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Phil"}

	philCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Phil"}

	philDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Phil"}

	philFrenchName = nook.Name{
		Language: language.French,
		Value:    "Phil"}

	philGermanName = nook.Name{
		Language: language.German,
		Value:    "Ingo"}

	philItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Enzo"}

	philJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ケイン"}

	philKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "케인"}

	philLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Amalio"}

	philRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фил"}

	philSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "凯恩"}

	philSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Amalio"}

	philTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "凱恩"}
)

var (
	philName = nook.Languages{
		language.AmericanEnglish:      philAmericanEnglishName,
		language.CanadianFrench:       philCanadianFrenchName,
		language.Dutch:                philDutchName,
		language.French:               philFrenchName,
		language.German:               philGermanName,
		language.Italian:              philItalianName,
		language.Japanese:             philJapaneseName,
		language.Korean:               philKoreanName,
		language.LatinAmericanSpanish: philLatinAmericanSpanishName,
		language.Russian:              philRussianName,
		language.SimplifiedChinese:    philSimplifiedChineseName,
		language.Spanish:              philSpanishName,
		language.TraditionalChinese:   philTraditionalChineseName}
)

var (
	philCharacter = nook.Character{
		Animal:   Ostrich,
		Birthday: philBirthday,
		Code:     philCode,
		Gender:   nook.Male,
		Name:     philName}
)

var (
	philAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hurk"}

	philCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bécot"}

	philDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "struis"}

	philFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bécot"}

	philGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zwoing"}

	philItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tambien"}

	philJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ホロロ"}

	philKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "호롤로"}

	philLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jurk"}

	philRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хех"}

	philSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "轰隆隆"}

	philSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jurk"}

	philTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "轟隆隆"}
)

var (
	philPhrase = nook.Languages{
		language.AmericanEnglish:      philAmericanEnglishName,
		language.CanadianFrench:       philCanadianFrenchName,
		language.Dutch:                philDutchName,
		language.French:               philFrenchName,
		language.German:               philGermanName,
		language.Italian:              philItalianName,
		language.Japanese:             philJapaneseName,
		language.Korean:               philKoreanName,
		language.LatinAmericanSpanish: philLatinAmericanSpanishName,
		language.Russian:              philRussianName,
		language.SimplifiedChinese:    philSimplifiedChineseName,
		language.Spanish:              philSpanishName,
		language.TraditionalChinese:   philTraditionalChineseName}
)

var (
	Phil = nook.Villager{
		Character:   philCharacter,
		Personality: nook.Smug,
		Phrase:      philPhrase}
)
