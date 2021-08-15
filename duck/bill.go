package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Chocochoupichou"}

	billDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Billkwaker"}

	billFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chocochoupichou"}

	billGermanName = nook.Name{
		Language: language.German,
		Value:    "Billquako"}

	billItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ginoquaquo"}

	billJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピータンだね"}

	billKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "코코아탕"}

	billLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paquitocuaco"}

	billRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Биллкряко"}

	billSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "皮蛋蛋是啊"}

	billSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paquitocuaco"}

	billTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "皮蛋蛋是啊"}
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
		Animal:   Duck,
		Birthday: billBirthday,
		Code:     billCode,
		Gender:   nook.Male,
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
		Personality: nook.Jock,
		Phrase:      billPhrase}
)
