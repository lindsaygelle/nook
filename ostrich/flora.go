package ostrich

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	floraBirthday = nook.Birthday{
		Day:   9,
		Month: time.February}
)

var (
	floraCode = nook.Code{
		Value: "ost09"}
)

var (
	floraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flora"}

	floraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Justinerosi rosa"}

	floraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Floraroze"}

	floraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Justineguiboles"}

	floraGermanName = nook.Name{
		Language: language.German,
		Value:    "Floraflappflapp"}

	floraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosalbasivuplé"}

	floraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フララだミン"}

	floraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "플라라밍고밍고"}

	floraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Azucenacroqui"}

	floraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Флорарозочка"}

	floraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "胡拉拉红鹤"}

	floraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Azucenacroqui"}

	floraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "胡拉拉紅鶴"}
)

var (
	floraName = nook.Languages{
		language.AmericanEnglish:      floraAmericanEnglishName,
		language.CanadianFrench:       floraCanadianFrenchName,
		language.Dutch:                floraDutchName,
		language.French:               floraFrenchName,
		language.German:               floraGermanName,
		language.Italian:              floraItalianName,
		language.Japanese:             floraJapaneseName,
		language.Korean:               floraKoreanName,
		language.LatinAmericanSpanish: floraLatinAmericanSpanishName,
		language.Russian:              floraRussianName,
		language.SimplifiedChinese:    floraSimplifiedChineseName,
		language.Spanish:              floraSpanishName,
		language.TraditionalChinese:   floraTraditionalChineseName}
)

var (
	floraCharacter = nook.Character{
		Animal:   Ostrich,
		Birthday: floraBirthday,
		Code:     floraCode,
		Gender:   nook.Female,
		Name:     floraName}
)

var (
	floraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pinky"}

	floraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "rosi rosa"}

	floraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "roze"}

	floraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "guiboles"}

	floraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flappflapp"}

	floraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sivuplé"}

	floraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だミン"}

	floraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "밍고밍고"}

	floraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "croqui"}

	floraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "розочка"}

	floraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "红鹤"}

	floraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "croqui"}

	floraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "紅鶴"}
)

var (
	floraPhrase = nook.Languages{
		language.AmericanEnglish:      floraAmericanEnglishName,
		language.CanadianFrench:       floraCanadianFrenchName,
		language.Dutch:                floraDutchName,
		language.French:               floraFrenchName,
		language.German:               floraGermanName,
		language.Italian:              floraItalianName,
		language.Japanese:             floraJapaneseName,
		language.Korean:               floraKoreanName,
		language.LatinAmericanSpanish: floraLatinAmericanSpanishName,
		language.Russian:              floraRussianName,
		language.SimplifiedChinese:    floraSimplifiedChineseName,
		language.Spanish:              floraSpanishName,
		language.TraditionalChinese:   floraTraditionalChineseName}
)

var (
	Flora = nook.Villager{
		Character:   floraCharacter,
		Personality: nook.Peppy,
		Phrase:      floraPhrase}
)
