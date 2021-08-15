package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Crokomiam miam"}

	bootsDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bootssmikkel"}

	bootsFrenchName = nook.Name{
		Language: language.French,
		Value:    "Crokomariolle"}

	bootsGermanName = nook.Name{
		Language: language.German,
		Value:    "Tilmannschmatzi"}

	bootsItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Croccognammi"}

	bootsJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホウサクだぴょん"}

	bootsKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "풍작만세"}

	bootsLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brauliocrococroco"}

	bootsRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бутскусь"}

	bootsSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丰年跳"}

	bootsSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brauliocrococroco"}

	bootsTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "豐年跳"}
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
		Animal:   Alligator,
		Birthday: bootsBirthday,
		Code:     bootsCode,
		Gender:   nook.Male,
		Name:     bootsName}
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
	Boots = nook.Villager{
		Character:   bootsCharacter,
		Personality: nook.Jock,
		Phrase:      bootsPhrase}
)
