package mouse

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
		Value:    "Gruyère"}

	limbergDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Limberg"}

	limbergFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gruyère"}

	limbergGermanName = nook.Name{
		Language: language.German,
		Value:    "Rafael"}

	limbergItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosico"}

	limbergJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "らっきょ"}

	limbergKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "단무지"}

	limbergLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Camember"}

	limbergRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лимберг"}

	limbergSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "韭菜"}

	limbergSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Camember"}

	limbergTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "韭菜"}
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
		Animal:   animal.Mouse,
		Birthday: limbergBirthday,
		Code:     limbergCode,
		Key:      character.Limberg,
		Gender:   gender.Male,
		Name:     limbergName,
		Special:  false}
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
		language.AmericanEnglish:      limbergAmericanEnglishPhrase,
		language.CanadianFrench:       limbergCanadianFrenchPhrase,
		language.Dutch:                limbergDutchPhrase,
		language.French:               limbergFrenchPhrase,
		language.German:               limbergGermanPhrase,
		language.Italian:              limbergItalianPhrase,
		language.Japanese:             limbergJapanesePhrase,
		language.Korean:               limbergKoreanPhrase,
		language.LatinAmericanSpanish: limbergLatinAmericanSpanishPhrase,
		language.Russian:              limbergRussianPhrase,
		language.SimplifiedChinese:    limbergSimplifiedChinesePhrase,
		language.Spanish:              limbergSpanishPhrase,
		language.TraditionalChinese:   limbergTraditionalChinesePhrase}
)

var (
	Limberg = nook.Villager{
		Character:   limbergCharacter,
		Personality: personality.Cranky,
		Phrase:      limbergPhrase}
)
