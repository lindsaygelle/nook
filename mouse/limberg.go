package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	limbergBirthday = nook.Birthday{
		Day:   17,
		Month: time.October}
)

var (
	limbergCode = nook.Code{
		Value: "mus01"}
)

var (
	limbergAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Limberg"}

	limbergCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gruyèrescriiich"}

	limbergDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Limbergpiepsels"}

	limbergFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gruyèrescriiich"}

	limbergGermanName = nook.Name{
		Language: language.German,
		Value:    "Rafaelpfiffikus"}

	limbergItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosicosquick"}

	limbergJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "らっきょてやんで"}

	limbergKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "단무지긍가벼"}

	limbergLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Camemberscroch"}

	limbergRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лимбергфью-фью"}

	limbergSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "韭菜非也"}

	limbergSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Camembercaramba"}

	limbergTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "韭菜非也"}
)

var (
	limbergName = nook.Languages{
		language.AmericanEnglish:      limbergAmericanEnglishName,
		language.CanadianFrench:       limbergCanadianFrenchName,
		language.Dutch:                limbergDutchName,
		language.French:               limbergFrenchName,
		language.German:               limbergGermanName,
		language.Italian:              limbergItalianName,
		language.Japanese:             limbergJapaneseName,
		language.Korean:               limbergKoreanName,
		language.LatinAmericanSpanish: limbergLatinAmericanSpanishName,
		language.Russian:              limbergRussianName,
		language.SimplifiedChinese:    limbergSimplifiedChineseName,
		language.Spanish:              limbergSpanishName,
		language.TraditionalChinese:   limbergTraditionalChineseName}
)

var (
	limbergCharacter = nook.Character{
		Animal:   Mouse,
		Birthday: limbergBirthday,
		Code:     limbergCode,
		Gender:   nook.Male,
		Name:     limbergName}
)

var (
	limbergAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "squinky"}

	limbergCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "scriiich"}

	limbergDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "piepsels"}

	limbergFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "scriiich"}

	limbergGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pfiffikus"}

	limbergItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squick"}

	limbergJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "てやんで"}

	limbergKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "긍가벼"}

	limbergLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "scroch"}

	limbergRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фью-фью"}

	limbergSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "非也"}

	limbergSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "caramba"}

	limbergTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "非也"}
)

var (
	limbergPhrase = nook.Languages{
		language.AmericanEnglish:      limbergAmericanEnglishName,
		language.CanadianFrench:       limbergCanadianFrenchName,
		language.Dutch:                limbergDutchName,
		language.French:               limbergFrenchName,
		language.German:               limbergGermanName,
		language.Italian:              limbergItalianName,
		language.Japanese:             limbergJapaneseName,
		language.Korean:               limbergKoreanName,
		language.LatinAmericanSpanish: limbergLatinAmericanSpanishName,
		language.Russian:              limbergRussianName,
		language.SimplifiedChinese:    limbergSimplifiedChineseName,
		language.Spanish:              limbergSpanishName,
		language.TraditionalChinese:   limbergTraditionalChineseName}
)

var (
	Limberg = nook.Villager{
		Character:   limbergCharacter,
		Personality: nook.Cranky,
		Phrase:      limbergPhrase}
)
