package bird

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
	twiggyBirthday = nook.Birthday{
		Day:   13,
		Month: time.July}
)

var (
	twiggyCode = nook.Code{
		Value: "brd03"}
)

var (
	twiggyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Twiggy"}

	twiggyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Titi"}

	twiggyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Twiggy"}

	twiggyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Titi"}

	twiggyGermanName = nook.Name{
		Language: language.German,
		Value:    "Twiggy"}

	twiggyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Titti"}

	twiggyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピーチク"}

	twiggyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핀틱"}

	twiggyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tití"}

	twiggyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Твигги"}

	twiggySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "叽叽"}

	twiggySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tití"}

	twiggyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘰嘰"}
)

var (
	twiggyName = nook.Languages{
		language.AmericanEnglish:      twiggyAmericanEnglishName,
		language.CanadianFrench:       twiggyCanadianFrenchName,
		language.Dutch:                twiggyDutchName,
		language.French:               twiggyFrenchName,
		language.German:               twiggyGermanName,
		language.Italian:              twiggyItalianName,
		language.Japanese:             twiggyJapaneseName,
		language.Korean:               twiggyKoreanName,
		language.LatinAmericanSpanish: twiggyLatinAmericanSpanishName,
		language.Russian:              twiggyRussianName,
		language.SimplifiedChinese:    twiggySimplifiedChineseName,
		language.Spanish:              twiggySpanishName,
		language.TraditionalChinese:   twiggyTraditionalChineseName}
)

var (
	twiggyCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: twiggyBirthday,
		Code:     twiggyCode,
		Key:      character.Twiggy,
		Gender:   gender.Female,
		Name:     twiggyName,
		Special:  false}
)

var (
	twiggyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cheepers"}

	twiggyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coolos"}

	twiggyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "fwieeet"}

	twiggyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "coolos"}

	twiggyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zwirtschi"}

	twiggyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ciiip"}

	twiggyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ッピ"}

	twiggyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "크"}

	twiggyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tirití"}

	twiggyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тю-вить"}

	twiggySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "叽"}

	twiggySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tirití"}

	twiggyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘰"}
)

var (
	twiggyPhrase = nook.Languages{
		language.AmericanEnglish:      twiggyAmericanEnglishPhrase,
		language.CanadianFrench:       twiggyCanadianFrenchPhrase,
		language.Dutch:                twiggyDutchPhrase,
		language.French:               twiggyFrenchPhrase,
		language.German:               twiggyGermanPhrase,
		language.Italian:              twiggyItalianPhrase,
		language.Japanese:             twiggyJapanesePhrase,
		language.Korean:               twiggyKoreanPhrase,
		language.LatinAmericanSpanish: twiggyLatinAmericanSpanishPhrase,
		language.Russian:              twiggyRussianPhrase,
		language.SimplifiedChinese:    twiggySimplifiedChinesePhrase,
		language.Spanish:              twiggySpanishPhrase,
		language.TraditionalChinese:   twiggyTraditionalChinesePhrase}
)

var (
	Twiggy = nook.Villager{
		Character:   twiggyCharacter,
		Personality: personality.Peppy,
		Phrase:      twiggyPhrase}
)
