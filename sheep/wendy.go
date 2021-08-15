package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	wendyBirthday = nook.Birthday{
		Day:   15,
		Month: time.August}
)

var (
	wendyCode = nook.Code{
		Value: "shp09"}
)

var (
	wendyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wendy"}

	wendyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Karenbê-bê-bê"}

	wendyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wendylammetje"}

	wendyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Karenbê-bê-bê"}

	wendyGermanName = nook.Name{
		Language: language.German,
		Value:    "Wolliwuschel"}

	wendyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Agnolabeee-ehi"}

	wendyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みぞれイシシ"}

	wendyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "눈송이꺄르르"}

	wendyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Franelabee-biis"}

	wendyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Вендиовечка"}

	wendySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪花一丝丝"}

	wendySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Franelajoyita"}

	wendyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪花一絲絲"}
)

var (
	wendyName = nook.Languages{
		language.AmericanEnglish:      wendyAmericanEnglishName,
		language.CanadianFrench:       wendyCanadianFrenchName,
		language.Dutch:                wendyDutchName,
		language.French:               wendyFrenchName,
		language.German:               wendyGermanName,
		language.Italian:              wendyItalianName,
		language.Japanese:             wendyJapaneseName,
		language.Korean:               wendyKoreanName,
		language.LatinAmericanSpanish: wendyLatinAmericanSpanishName,
		language.Russian:              wendyRussianName,
		language.SimplifiedChinese:    wendySimplifiedChineseName,
		language.Spanish:              wendySpanishName,
		language.TraditionalChinese:   wendyTraditionalChineseName}
)

var (
	wendyCharacter = nook.Character{
		Animal:   Sheep,
		Birthday: wendyBirthday,
		Code:     wendyCode,
		Gender:   nook.Female,
		Name:     wendyName}
)

var (
	wendyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "lambkins"}

	wendyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bê-bê-bê"}

	wendyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "lammetje"}

	wendyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bê-bê-bê"}

	wendyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wuschel"}

	wendyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beee-ehi"}

	wendyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "イシシ"}

	wendyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꺄르르"}

	wendyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bee-biis"}

	wendyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "овечка"}

	wendySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "一丝丝"}

	wendySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "joyita"}

	wendyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "一絲絲"}
)

var (
	wendyPhrase = nook.Languages{
		language.AmericanEnglish:      wendyAmericanEnglishName,
		language.CanadianFrench:       wendyCanadianFrenchName,
		language.Dutch:                wendyDutchName,
		language.French:               wendyFrenchName,
		language.German:               wendyGermanName,
		language.Italian:              wendyItalianName,
		language.Japanese:             wendyJapaneseName,
		language.Korean:               wendyKoreanName,
		language.LatinAmericanSpanish: wendyLatinAmericanSpanishName,
		language.Russian:              wendyRussianName,
		language.SimplifiedChinese:    wendySimplifiedChineseName,
		language.Spanish:              wendySpanishName,
		language.TraditionalChinese:   wendyTraditionalChineseName}
)

var (
	Wendy = nook.Villager{
		Character:   wendyCharacter,
		Personality: nook.Peppy,
		Phrase:      wendyPhrase}
)
