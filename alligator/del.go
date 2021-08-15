package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	delBirthday = nook.Birthday{
		Day:   27,
		Month: time.May}
)

var (
	delCode = nook.Code{
		Value: "crd04"}
)

var (
	delAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Del"}

	delCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hectorclap"}

	delDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Delgronk"}

	delFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hectorclap"}

	delGermanName = nook.Name{
		Language: language.German,
		Value:    "Krokkihapp"}

	delItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Krokkichomp"}

	delJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヤマトプシュー"}

	delKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파도맨처얼썩"}

	delLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Crocochomp"}

	delRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Делщелк"}

	delSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大和噗咻"}

	delSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Crocochomp"}

	delTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大和噗咻"}
)

var (
	delName = nook.Languages{
		language.AmericanEnglish:      delAmericanEnglishName,
		language.CanadianFrench:       delCanadianFrenchName,
		language.Dutch:                delDutchName,
		language.French:               delFrenchName,
		language.German:               delGermanName,
		language.Italian:              delItalianName,
		language.Japanese:             delJapaneseName,
		language.Korean:               delKoreanName,
		language.LatinAmericanSpanish: delLatinAmericanSpanishName,
		language.Russian:              delRussianName,
		language.SimplifiedChinese:    delSimplifiedChineseName,
		language.Spanish:              delSpanishName,
		language.TraditionalChinese:   delTraditionalChineseName}
)

var (
	delCharacter = nook.Character{
		Animal:   Alligator,
		Birthday: delBirthday,
		Code:     delCode,
		Gender:   nook.Male,
		Name:     delName}
)

var (
	delAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "gronk"}

	delCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "clap"}

	delDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gronk"}

	delFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "clap"}

	delGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "happ"}

	delItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "chomp"}

	delJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "プシュー"}

	delKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "처얼썩"}

	delLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chomp"}

	delRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "щелк"}

	delSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗咻"}

	delSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chomp"}

	delTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗咻"}
)

var (
	delPhrase = nook.Languages{
		language.AmericanEnglish:      delAmericanEnglishName,
		language.CanadianFrench:       delCanadianFrenchName,
		language.Dutch:                delDutchName,
		language.French:               delFrenchName,
		language.German:               delGermanName,
		language.Italian:              delItalianName,
		language.Japanese:             delJapaneseName,
		language.Korean:               delKoreanName,
		language.LatinAmericanSpanish: delLatinAmericanSpanishName,
		language.Russian:              delRussianName,
		language.SimplifiedChinese:    delSimplifiedChineseName,
		language.Spanish:              delSpanishName,
		language.TraditionalChinese:   delTraditionalChineseName}
)

var (
	Del = nook.Villager{
		Character:   delCharacter,
		Personality: nook.Cranky,
		Phrase:      delPhrase}
)
