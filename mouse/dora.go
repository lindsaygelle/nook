package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	doraBirthday = nook.Birthday{
		Day:   18,
		Month: time.February}
)

var (
	doraCode = nook.Code{
		Value: "mus00"}
)

var (
	doraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dora"}

	doraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Doracoui-coui"}

	doraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dorapieperig"}

	doraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Doratatouille"}

	doraGermanName = nook.Name{
		Language: language.German,
		Value:    "Doraquieek"}

	doraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Serenasquiksquik"}

	doraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "とめだべ"}

	doraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티미찍찍"}

	doraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Doricuiqui"}

	doraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дорапип-пип"}

	doraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "杜美美吧"}

	doraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Doricuiqui"}

	doraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "杜美美吧"}
)

var (
	doraName = nook.Languages{
		language.AmericanEnglish:      doraAmericanEnglishName,
		language.CanadianFrench:       doraCanadianFrenchName,
		language.Dutch:                doraDutchName,
		language.French:               doraFrenchName,
		language.German:               doraGermanName,
		language.Italian:              doraItalianName,
		language.Japanese:             doraJapaneseName,
		language.Korean:               doraKoreanName,
		language.LatinAmericanSpanish: doraLatinAmericanSpanishName,
		language.Russian:              doraRussianName,
		language.SimplifiedChinese:    doraSimplifiedChineseName,
		language.Spanish:              doraSpanishName,
		language.TraditionalChinese:   doraTraditionalChineseName}
)

var (
	doraCharacter = nook.Character{
		Animal:   Mouse,
		Birthday: doraBirthday,
		Code:     doraCode,
		Gender:   nook.Female,
		Name:     doraName}
)

var (
	doraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "squeaky"}

	doraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coui-coui"}

	doraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pieperig"}

	doraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tatouille"}

	doraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quieek"}

	doraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squiksquik"}

	doraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だべ"}

	doraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "찍찍"}

	doraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuiqui"}

	doraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пип-пип"}

	doraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "美吧"}

	doraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuiqui"}

	doraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "美吧"}
)

var (
	doraPhrase = nook.Languages{
		language.AmericanEnglish:      doraAmericanEnglishName,
		language.CanadianFrench:       doraCanadianFrenchName,
		language.Dutch:                doraDutchName,
		language.French:               doraFrenchName,
		language.German:               doraGermanName,
		language.Italian:              doraItalianName,
		language.Japanese:             doraJapaneseName,
		language.Korean:               doraKoreanName,
		language.LatinAmericanSpanish: doraLatinAmericanSpanishName,
		language.Russian:              doraRussianName,
		language.SimplifiedChinese:    doraSimplifiedChineseName,
		language.Spanish:              doraSpanishName,
		language.TraditionalChinese:   doraTraditionalChineseName}
)

var (
	Dora = nook.Villager{
		Character:   doraCharacter,
		Personality: nook.Normal,
		Phrase:      doraPhrase}
)
