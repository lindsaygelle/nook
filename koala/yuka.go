package koala

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Calypsotss tss"}

	yukaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Yukatss tss"}

	yukaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Calypsotss tss"}

	yukaGermanName = nook.Name{
		Language: language.German,
		Value:    "Utegrins"}

	yukaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Yukatsé tsé"}

	yukaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ユーカリアラ"}

	yukaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "유카리어맛"}

	yukaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Yukatsé tsé"}

	yukaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Юкац-ц"}

	yukaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "尤嘉莉啊啦"}

	yukaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Yukapues"}

	yukaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "尤嘉莉啊啦"}
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
		Animal:   Koala,
		Birthday: yukaBirthday,
		Code:     yukaCode,
		Gender:   nook.Female,
		Name:     yukaName}
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
	Yuka = nook.Villager{
		Character:   yukaCharacter,
		Personality: nook.Snooty,
		Phrase:      yukaPhrase}
)
