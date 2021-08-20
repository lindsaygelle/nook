package ostrich

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
		Value:    "Justine"}

	floraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Flora"}

	floraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Justine"}

	floraGermanName = nook.Name{
		Language: language.German,
		Value:    "Flora"}

	floraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosalba"}

	floraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フララ"}

	floraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "플라라"}

	floraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Azucena"}

	floraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Флора"}

	floraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "胡拉拉"}

	floraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Azucena"}

	floraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "胡拉拉"}
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
		Animal:   animal.Ostrich,
		Birthday: floraBirthday,
		Code:     floraCode,
		Key:      character.Flora,
		Gender:   gender.Female,
		Name:     floraName,
		Special:  false}
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
		language.AmericanEnglish:      floraAmericanEnglishPhrase,
		language.CanadianFrench:       floraCanadianFrenchPhrase,
		language.Dutch:                floraDutchPhrase,
		language.French:               floraFrenchPhrase,
		language.German:               floraGermanPhrase,
		language.Italian:              floraItalianPhrase,
		language.Japanese:             floraJapanesePhrase,
		language.Korean:               floraKoreanPhrase,
		language.LatinAmericanSpanish: floraLatinAmericanSpanishPhrase,
		language.Russian:              floraRussianPhrase,
		language.SimplifiedChinese:    floraSimplifiedChinesePhrase,
		language.Spanish:              floraSpanishPhrase,
		language.TraditionalChinese:   floraTraditionalChinesePhrase}
)

var (
	Flora = nook.Villager{
		Character:   floraCharacter,
		Personality: personality.Peppy,
		Phrase:      floraPhrase}
)
