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

var (
	bootsBirthday = nook.Birthday{
		Day:   7,
		Month: time.August}
)

var (
	bootsCode = nook.Code{
		Value: "crd02"}
)

var (
	bootsAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Boots"}

	bootsCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Croko"}

	bootsDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Boots"}

	bootsFrenchName = nook.Name{
		Language: language.French,
		Value:    "Croko"}

	bootsGermanName = nook.Name{
		Language: language.German,
		Value:    "Tilmann"}

	bootsItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Crocco"}

	bootsJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホウサク"}

	bootsKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "풍작"}

	bootsLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Braulio"}

	bootsRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бутс"}

	bootsSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丰年"}

	bootsSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Braulio"}

	bootsTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "豐年"}
)

var (
	bootsName = nook.Languages{
		language.AmericanEnglish:      bootsAmericanEnglishName,
		language.CanadianFrench:       bootsCanadianFrenchName,
		language.Dutch:                bootsDutchName,
		language.French:               bootsFrenchName,
		language.German:               bootsGermanName,
		language.Italian:              bootsItalianName,
		language.Japanese:             bootsJapaneseName,
		language.Korean:               bootsKoreanName,
		language.LatinAmericanSpanish: bootsLatinAmericanSpanishName,
		language.Russian:              bootsRussianName,
		language.SimplifiedChinese:    bootsSimplifiedChineseName,
		language.Spanish:              bootsSpanishName,
		language.TraditionalChinese:   bootsTraditionalChineseName}
)

var (
	bootsCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: bootsBirthday,
		Code:     bootsCode,
		Key:      character.Boots,
		Gender:   gender.Male,
		Name:     bootsName,
		Special:  false}
)

var (
	bootsAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "munchie"}

	bootsCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miam miam"}

	bootsDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "smikkel"}

	bootsFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mariolle"}

	bootsGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schmatzi"}

	bootsItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnammi"}

	bootsJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だぴょん"}

	bootsKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "만세"}

	bootsLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "crococroco"}

	bootsRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кусь"}

	bootsSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "跳"}

	bootsSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "crococroco"}

	bootsTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "跳"}
)

var (
	bootsPhrase = nook.Languages{
		language.AmericanEnglish:      bootsAmericanEnglishPhrase,
		language.CanadianFrench:       bootsCanadianFrenchPhrase,
		language.Dutch:                bootsDutchPhrase,
		language.French:               bootsFrenchPhrase,
		language.German:               bootsGermanPhrase,
		language.Italian:              bootsItalianPhrase,
		language.Japanese:             bootsJapanesePhrase,
		language.Korean:               bootsKoreanPhrase,
		language.LatinAmericanSpanish: bootsLatinAmericanSpanishPhrase,
		language.Russian:              bootsRussianPhrase,
		language.SimplifiedChinese:    bootsSimplifiedChinesePhrase,
		language.Spanish:              bootsSpanishPhrase,
		language.TraditionalChinese:   bootsTraditionalChinesePhrase}
)

var (
	Boots = nook.Villager{
		Character:   bootsCharacter,
		Personality: personality.Jock,
		Phrase:      bootsPhrase}
)
