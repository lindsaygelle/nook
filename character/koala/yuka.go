package koala

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
	yukaBirthday = nook.Birthday{
		Day:   20,
		Month: time.July}
)

var (
	yukaCode = nook.Code{
		Value: "kal00"}
)

var (
	yukaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Yuka"}

	yukaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Calypso"}

	yukaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Yuka"}

	yukaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Calypso"}

	yukaGermanName = nook.Name{
		Language: language.German,
		Value:    "Ute"}

	yukaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Yuka"}

	yukaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ユーカリ"}

	yukaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "유카리"}

	yukaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Yuka"}

	yukaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Юка"}

	yukaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "尤嘉莉"}

	yukaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Yuka"}

	yukaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "尤嘉莉"}
)

var (
	yukaName = nook.Languages{
		language.AmericanEnglish:      yukaAmericanEnglishName,
		language.CanadianFrench:       yukaCanadianFrenchName,
		language.Dutch:                yukaDutchName,
		language.French:               yukaFrenchName,
		language.German:               yukaGermanName,
		language.Italian:              yukaItalianName,
		language.Japanese:             yukaJapaneseName,
		language.Korean:               yukaKoreanName,
		language.LatinAmericanSpanish: yukaLatinAmericanSpanishName,
		language.Russian:              yukaRussianName,
		language.SimplifiedChinese:    yukaSimplifiedChineseName,
		language.Spanish:              yukaSpanishName,
		language.TraditionalChinese:   yukaTraditionalChineseName}
)

var (
	yukaCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: yukaBirthday,
		Code:     yukaCode,
		Key:      character.Yuka,
		Gender:   gender.Female,
		Name:     yukaName,
		Special:  false}
)

var (
	yukaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tsk tsk"}

	yukaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tss tss"}

	yukaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tss tss"}

	yukaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tss tss"}

	yukaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grins"}

	yukaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tsé tsé"}

	yukaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アラ"}

	yukaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어맛"}

	yukaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tsé tsé"}

	yukaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ц-ц"}

	yukaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啊啦"}

	yukaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pues"}

	yukaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啊啦"}
)

var (
	yukaPhrase = nook.Languages{
		language.AmericanEnglish:      yukaAmericanEnglishPhrase,
		language.CanadianFrench:       yukaCanadianFrenchPhrase,
		language.Dutch:                yukaDutchPhrase,
		language.French:               yukaFrenchPhrase,
		language.German:               yukaGermanPhrase,
		language.Italian:              yukaItalianPhrase,
		language.Japanese:             yukaJapanesePhrase,
		language.Korean:               yukaKoreanPhrase,
		language.LatinAmericanSpanish: yukaLatinAmericanSpanishPhrase,
		language.Russian:              yukaRussianPhrase,
		language.SimplifiedChinese:    yukaSimplifiedChinesePhrase,
		language.Spanish:              yukaSpanishPhrase,
		language.TraditionalChinese:   yukaTraditionalChinesePhrase}
)

var (
	Yuka = nook.Villager{
		Character:   yukaCharacter,
		Personality: personality.Snooty,
		Phrase:      yukaPhrase}
)
