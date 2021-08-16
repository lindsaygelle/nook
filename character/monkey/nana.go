package monkey

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	nanaBirthday = nook.Birthday{
		Day:   23,
		Month: time.August}
)

var (
	nanaCode = nook.Code{
		Value: "mnk01"}
)

var (
	nanaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nana"}

	nanaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mireille"}

	nanaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nana"}

	nanaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mireille"}

	nanaGermanName = nook.Name{
		Language: language.German,
		Value:    "Dorothea"}

	nanaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nanà"}

	nanaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チッチ"}

	nanaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "키키"}

	nanaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nana"}

	nanaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нана"}

	nanaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "七七"}

	nanaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nana"}

	nanaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "七七"}
)

var (
	nanaName = nook.Languages{
		language.AmericanEnglish:      nanaAmericanEnglishName,
		language.CanadianFrench:       nanaCanadianFrenchName,
		language.Dutch:                nanaDutchName,
		language.French:               nanaFrenchName,
		language.German:               nanaGermanName,
		language.Italian:              nanaItalianName,
		language.Japanese:             nanaJapaneseName,
		language.Korean:               nanaKoreanName,
		language.LatinAmericanSpanish: nanaLatinAmericanSpanishName,
		language.Russian:              nanaRussianName,
		language.SimplifiedChinese:    nanaSimplifiedChineseName,
		language.Spanish:              nanaSpanishName,
		language.TraditionalChinese:   nanaTraditionalChineseName}
)

var (
	nanaCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: nanaBirthday,
		Code:     nanaCode,
		Gender:   gender.Female,
		Name:     nanaName}
)

var (
	nanaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "po po"}

	nanaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "po po"}

	nanaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "banaan"}

	nanaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "po po"}

	nanaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "spitzi"}

	nanaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "po po"}

	nanaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウキャ"}

	nanaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "끼끼"}

	nanaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ukiú"}

	nanaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пу-пу"}

	nanaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唔唔"}

	nanaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "monada"}

	nanaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唔唔"}
)

var (
	nanaPhrase = nook.Languages{
		language.AmericanEnglish:      nanaAmericanEnglishName,
		language.CanadianFrench:       nanaCanadianFrenchName,
		language.Dutch:                nanaDutchName,
		language.French:               nanaFrenchName,
		language.German:               nanaGermanName,
		language.Italian:              nanaItalianName,
		language.Japanese:             nanaJapaneseName,
		language.Korean:               nanaKoreanName,
		language.LatinAmericanSpanish: nanaLatinAmericanSpanishName,
		language.Russian:              nanaRussianName,
		language.SimplifiedChinese:    nanaSimplifiedChineseName,
		language.Spanish:              nanaSpanishName,
		language.TraditionalChinese:   nanaTraditionalChineseName}
)

var (
	Nana = nook.Villager{
		Character:   nanaCharacter,
		Personality: personality.Normal,
		Phrase:      nanaPhrase}
)
