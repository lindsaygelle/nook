package sheep

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
	woolioBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	woolioCode = nook.Code{
		Value: ""}
)

var (
	woolioAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Woolio"}

	woolioCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	woolioDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	woolioFrenchName = nook.Name{
		Language: language.French,
		Value:    "Moumoute"}

	woolioGermanName = nook.Name{
		Language: language.German,
		Value:    "Wulle"}

	woolioItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Moero"}

	woolioJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホシオ"}

	woolioKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	woolioLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	woolioRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	woolioSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "星星"}

	woolioSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Esquilo"}

	woolioTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	woolioName = nook.Languages{
		language.AmericanEnglish:      woolioAmericanEnglishName,
		language.CanadianFrench:       woolioCanadianFrenchName,
		language.Dutch:                woolioDutchName,
		language.French:               woolioFrenchName,
		language.German:               woolioGermanName,
		language.Italian:              woolioItalianName,
		language.Japanese:             woolioJapaneseName,
		language.Korean:               woolioKoreanName,
		language.LatinAmericanSpanish: woolioLatinAmericanSpanishName,
		language.Russian:              woolioRussianName,
		language.SimplifiedChinese:    woolioSimplifiedChineseName,
		language.Spanish:              woolioSpanishName,
		language.TraditionalChinese:   woolioTraditionalChineseName}
)

var (
	woolioCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: woolioBirthday,
		Code:     woolioCode,
		Key:      character.Woolio,
		Gender:   gender.Male,
		Name:     woolioName,
		Special:  false}
)

var (
	woolioAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "biz-aaa"}

	woolioCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	woolioDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	woolioFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bêêêê quoi"}

	woolioGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bizääää"}

	woolioItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sbaallo"}

	woolioJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヨロシク"}

	woolioKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	woolioLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	woolioRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	woolioSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "帮忙"}

	woolioSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "paaasssa"}

	woolioTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	woolioPhrase = nook.Languages{
		language.AmericanEnglish:      woolioAmericanEnglishPhrase,
		language.CanadianFrench:       woolioCanadianFrenchPhrase,
		language.Dutch:                woolioDutchPhrase,
		language.French:               woolioFrenchPhrase,
		language.German:               woolioGermanPhrase,
		language.Italian:              woolioItalianPhrase,
		language.Japanese:             woolioJapanesePhrase,
		language.Korean:               woolioKoreanPhrase,
		language.LatinAmericanSpanish: woolioLatinAmericanSpanishPhrase,
		language.Russian:              woolioRussianPhrase,
		language.SimplifiedChinese:    woolioSimplifiedChinesePhrase,
		language.Spanish:              woolioSpanishPhrase,
		language.TraditionalChinese:   woolioTraditionalChinesePhrase}
)

var (
	Woolio = nook.Villager{
		Character:   woolioCharacter,
		Personality: personality.Jock,
		Phrase:      woolioPhrase}
)
