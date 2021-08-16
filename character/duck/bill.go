package duck

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
	billBirthday = nook.Birthday{
		Day:   1,
		Month: time.February}
)

var (
	billCode = nook.Code{
		Value: "duk00"}
)

var (
	billAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bill"}

	billCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Choco"}

	billDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bill"}

	billFrenchName = nook.Name{
		Language: language.French,
		Value:    "Choco"}

	billGermanName = nook.Name{
		Language: language.German,
		Value:    "Bill"}

	billItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gino"}

	billJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピータン"}

	billKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "코코아"}

	billLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paquito"}

	billRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Билл"}

	billSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "皮蛋"}

	billSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paquito"}

	billTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "皮蛋"}
)

var (
	billName = nook.Languages{
		language.AmericanEnglish:      billAmericanEnglishName,
		language.CanadianFrench:       billCanadianFrenchName,
		language.Dutch:                billDutchName,
		language.French:               billFrenchName,
		language.German:               billGermanName,
		language.Italian:              billItalianName,
		language.Japanese:             billJapaneseName,
		language.Korean:               billKoreanName,
		language.LatinAmericanSpanish: billLatinAmericanSpanishName,
		language.Russian:              billRussianName,
		language.SimplifiedChinese:    billSimplifiedChineseName,
		language.Spanish:              billSpanishName,
		language.TraditionalChinese:   billTraditionalChineseName}
)

var (
	billCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: billBirthday,
		Code:     billCode,
		Key:      character.Bill,
		Gender:   gender.Male,
		Name:     billName}
)

var (
	billAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quacko"}

	billCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "choupichou"}

	billDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwaker"}

	billFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "choupichou"}

	billGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quako"}

	billItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quaquo"}

	billJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だね"}

	billKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "탕"}

	billLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuaco"}

	billRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кряко"}

	billSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蛋是啊"}

	billSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuaco"}

	billTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蛋是啊"}
)

var (
	billPhrase = nook.Languages{
		language.AmericanEnglish:      billAmericanEnglishName,
		language.CanadianFrench:       billCanadianFrenchName,
		language.Dutch:                billDutchName,
		language.French:               billFrenchName,
		language.German:               billGermanName,
		language.Italian:              billItalianName,
		language.Japanese:             billJapaneseName,
		language.Korean:               billKoreanName,
		language.LatinAmericanSpanish: billLatinAmericanSpanishName,
		language.Russian:              billRussianName,
		language.SimplifiedChinese:    billSimplifiedChineseName,
		language.Spanish:              billSpanishName,
		language.TraditionalChinese:   billTraditionalChineseName}
)

var (
	Bill = nook.Villager{
		Character:   billCharacter,
		Personality: personality.Jock,
		Phrase:      billPhrase}
)
