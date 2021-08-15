package monkey

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Mireillepo po"}

	nanaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nanabanaan"}

	nanaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mireillepo po"}

	nanaGermanName = nook.Name{
		Language: language.German,
		Value:    "Dorotheaspitzi"}

	nanaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nanàpo po"}

	nanaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チッチウキャ"}

	nanaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "키키끼끼"}

	nanaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nanaukiú"}

	nanaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нанапу-пу"}

	nanaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "七七唔唔"}

	nanaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nanamonada"}

	nanaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "七七唔唔"}
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
		Animal:   Monkey,
		Birthday: nanaBirthday,
		Code:     nanaCode,
		Gender:   nook.Female,
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
		Personality: nook.Normal,
		Phrase:      nanaPhrase}
)
