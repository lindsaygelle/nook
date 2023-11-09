package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// delBirthday represents Del's birthday (May 27th).
var (
	delBirthday = nook.Birthday{
		Day:   27,
		Month: time.May}
)

// delCode represents Del's unique code ("crd04").
var (
	delCode = nook.Code{
		Value: "crd04"}
)

// Different names for Del in various languages.
var (
	delAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Del"}

	delCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hector"}

	delDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Del"}

	delFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hector"}

	delGermanName = nook.Name{
		Language: language.German,
		Value:    "Krokki"}

	delItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Krokki"}

	delJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヤマト"}

	delKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파도맨"}

	delLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Croco"}

	delRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дел"}

	delSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大和"}

	delSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Croco"}

	delTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大和"}
)

// delName represents Del's name in different languages.
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

// delCharacter represents Del's character information.
var (
	delCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: delBirthday,
		Code:     delCode,
		Key:      character.Del,
		Gender:   gender.Male,
		Name:     delName,
		Special:  false}
)

// Different phrases spoken by Del in various languages.
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

// delPhrase represents Del's phrases in different languages.
var (
	delPhrase = nook.Languages{
		language.AmericanEnglish:      delAmericanEnglishPhrase,
		language.CanadianFrench:       delCanadianFrenchPhrase,
		language.Dutch:                delDutchPhrase,
		language.French:               delFrenchPhrase,
		language.German:               delGermanPhrase,
		language.Italian:              delItalianPhrase,
		language.Japanese:             delJapanesePhrase,
		language.Korean:               delKoreanPhrase,
		language.LatinAmericanSpanish: delLatinAmericanSpanishPhrase,
		language.Russian:              delRussianPhrase,
		language.SimplifiedChinese:    delSimplifiedChinesePhrase,
		language.Spanish:              delSpanishPhrase,
		language.TraditionalChinese:   delTraditionalChinesePhrase}
)

// Del represents the character Del with his complete information.
var (
	Del = nook.Villager{
		Character:   delCharacter,
		Personality: personality.Cranky,
		Phrase:      delPhrase}
)
