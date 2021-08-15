package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Moumoutebêêêê quoi"}

	woolioGermanName = nook.Name{
		Language: language.German,
		Value:    "Wullebizääää"}

	woolioItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Moerosbaallo"}

	woolioJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホシオヨロシク"}

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
		Value:    "星星帮忙"}

	woolioSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Esquilopaaasssa"}

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
		Animal:   Sheep,
		Birthday: woolioBirthday,
		Code:     woolioCode,
		Gender:   nook.Male,
		Name:     woolioName}
)

var (
	woolioAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	woolioCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	woolioDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	woolioLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	woolioRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	woolioSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "帮忙"}

	woolioSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "paaasssa"}

	woolioTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	woolioPhrase = nook.Languages{
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
	Woolio = nook.Villager{
		Character:   woolioCharacter,
		Personality: nook.Jock,
		Phrase:      woolioPhrase}
)
