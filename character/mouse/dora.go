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
		Value:    "Dora"}

	doraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dora"}

	doraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dora"}

	doraGermanName = nook.Name{
		Language: language.German,
		Value:    "Dora"}

	doraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Serena"}

	doraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "とめ"}

	doraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티미"}

	doraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dori"}

	doraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дора"}

	doraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "杜美"}

	doraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Dori"}

	doraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "杜美"}
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
		Animal:   animal.Mouse,
		Birthday: doraBirthday,
		Code:     doraCode,
		Key:      character.Dora,
		Gender:   gender.Female,
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
		language.AmericanEnglish:      doraAmericanEnglishPhrase,
		language.CanadianFrench:       doraCanadianFrenchPhrase,
		language.Dutch:                doraDutchPhrase,
		language.French:               doraFrenchPhrase,
		language.German:               doraGermanPhrase,
		language.Italian:              doraItalianPhrase,
		language.Japanese:             doraJapanesePhrase,
		language.Korean:               doraKoreanPhrase,
		language.LatinAmericanSpanish: doraLatinAmericanSpanishPhrase,
		language.Russian:              doraRussianPhrase,
		language.SimplifiedChinese:    doraSimplifiedChinesePhrase,
		language.Spanish:              doraSpanishPhrase,
		language.TraditionalChinese:   doraTraditionalChinesePhrase}
)

var (
	Dora = nook.Villager{
		Character:   doraCharacter,
		Personality: personality.Normal,
		Phrase:      doraPhrase}
)
