package eagle

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
	apolloBirthday = nook.Birthday{
		Day:   4,
		Month: time.July}
)

var (
	apolloCode = nook.Code{
		Value: "pbr00"}
)

var (
	apolloAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Apollo"}

	apolloCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Apollon"}

	apolloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Apollo"}

	apolloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Apollon"}

	apolloGermanName = nook.Name{
		Language: language.German,
		Value:    "Apollo"}

	apolloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Apollo"}

	apolloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アポロ"}

	apolloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아폴로"}

	apolloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Apolo"}

	apolloRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аполло"}

	apolloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿波罗"}

	apolloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Apolo"}

	apolloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿波羅"}
)

var (
	apolloName = nook.Languages{
		language.AmericanEnglish:      apolloAmericanEnglishName,
		language.CanadianFrench:       apolloCanadianFrenchName,
		language.Dutch:                apolloDutchName,
		language.French:               apolloFrenchName,
		language.German:               apolloGermanName,
		language.Italian:              apolloItalianName,
		language.Japanese:             apolloJapaneseName,
		language.Korean:               apolloKoreanName,
		language.LatinAmericanSpanish: apolloLatinAmericanSpanishName,
		language.Russian:              apolloRussianName,
		language.SimplifiedChinese:    apolloSimplifiedChineseName,
		language.Spanish:              apolloSpanishName,
		language.TraditionalChinese:   apolloTraditionalChineseName}
)

var (
	apolloCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: apolloBirthday,
		Code:     apolloCode,
		Key:      character.Apollo,
		Gender:   gender.Male,
		Name:     apolloName,
		Special:  false}
)

var (
	apolloAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pah"}

	apolloCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "trraaaa"}

	apolloDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "poeh"}

	apolloFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "trraaaa"}

	apolloGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ahhh"}

	apolloItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pah"}

	apolloJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だワイ"}

	apolloKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "다꿇어"}

	apolloLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "rapahhh"}

	apolloRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пф-ф"}

	apolloSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唔咿"}

	apolloSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "rapaz"}

	apolloTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唔咿"}
)

var (
	apolloPhrase = nook.Languages{
		language.AmericanEnglish:      apolloAmericanEnglishPhrase,
		language.CanadianFrench:       apolloCanadianFrenchPhrase,
		language.Dutch:                apolloDutchPhrase,
		language.French:               apolloFrenchPhrase,
		language.German:               apolloGermanPhrase,
		language.Italian:              apolloItalianPhrase,
		language.Japanese:             apolloJapanesePhrase,
		language.Korean:               apolloKoreanPhrase,
		language.LatinAmericanSpanish: apolloLatinAmericanSpanishPhrase,
		language.Russian:              apolloRussianPhrase,
		language.SimplifiedChinese:    apolloSimplifiedChinesePhrase,
		language.Spanish:              apolloSpanishPhrase,
		language.TraditionalChinese:   apolloTraditionalChinesePhrase}
)

var (
	Apollo = nook.Villager{
		Character:   apolloCharacter,
		Personality: personality.Cranky,
		Phrase:      apolloPhrase}
)
