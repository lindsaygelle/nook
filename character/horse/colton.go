package horse

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
	coltonBirthday = nook.Birthday{
		Day:   22,
		Month: time.May}
)

var (
	coltonCode = nook.Code{
		Value: "hrs11"}
)

var (
	coltonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Colton"}

	coltonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tony"}

	coltonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Colton"}

	coltonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tony"}

	coltonGermanName = nook.Name{
		Language: language.German,
		Value:    "Carsten"}

	coltonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Furio"}

	coltonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アンソニー"}

	coltonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안소니"}

	coltonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Furio"}

	coltonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Колтон"}

	coltonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安索尼"}

	coltonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Furio"}

	coltonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安索尼"}
)

var (
	coltonName = nook.Languages{
		language.AmericanEnglish:      coltonAmericanEnglishName,
		language.CanadianFrench:       coltonCanadianFrenchName,
		language.Dutch:                coltonDutchName,
		language.French:               coltonFrenchName,
		language.German:               coltonGermanName,
		language.Italian:              coltonItalianName,
		language.Japanese:             coltonJapaneseName,
		language.Korean:               coltonKoreanName,
		language.LatinAmericanSpanish: coltonLatinAmericanSpanishName,
		language.Russian:              coltonRussianName,
		language.SimplifiedChinese:    coltonSimplifiedChineseName,
		language.Spanish:              coltonSpanishName,
		language.TraditionalChinese:   coltonTraditionalChineseName}
)

var (
	coltonCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: coltonBirthday,
		Code:     coltonCode,
		Key:      character.Colton,
		Gender:   gender.Male,
		Name:     coltonName}
)

var (
	coltonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "check it"}

	coltonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hue hue"}

	coltonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hortsik"}

	coltonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hue hue"}

	coltonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hamham"}

	coltonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hiii"}

	coltonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ごらんよ"}

	coltonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "이거봐"}

	coltonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jia"}

	coltonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "глянь"}

	coltonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "请看"}

	coltonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jia"}

	coltonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "請看"}
)

var (
	coltonPhrase = nook.Languages{
		language.AmericanEnglish:      coltonAmericanEnglishPhrase,
		language.CanadianFrench:       coltonCanadianFrenchPhrase,
		language.Dutch:                coltonDutchPhrase,
		language.French:               coltonFrenchPhrase,
		language.German:               coltonGermanPhrase,
		language.Italian:              coltonItalianPhrase,
		language.Japanese:             coltonJapanesePhrase,
		language.Korean:               coltonKoreanPhrase,
		language.LatinAmericanSpanish: coltonLatinAmericanSpanishPhrase,
		language.Russian:              coltonRussianPhrase,
		language.SimplifiedChinese:    coltonSimplifiedChinesePhrase,
		language.Spanish:              coltonSpanishPhrase,
		language.TraditionalChinese:   coltonTraditionalChinesePhrase}
)

var (
	Colton = nook.Villager{
		Character:   coltonCharacter,
		Personality: personality.Smug,
		Phrase:      coltonPhrase}
)
