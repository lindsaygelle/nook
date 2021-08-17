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
	gladysBirthday = nook.Birthday{
		Day:   15,
		Month: time.January}
)

var (
	gladysCode = nook.Code{
		Value: "ost01"}
)

var (
	gladysAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gladys"}

	gladysCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gladys"}

	gladysDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gladys"}

	gladysFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gladys"}

	gladysGermanName = nook.Name{
		Language: language.German,
		Value:    "Sandra"}

	gladysItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marta"}

	gladysJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちとせ"}

	gladysKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "빅토리아"}

	gladysLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gladis"}

	gladysRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Глэдис"}

	gladysSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "仙仙"}

	gladysSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gladis"}

	gladysTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "仙仙"}
)

var (
	gladysName = nook.Languages{
		language.AmericanEnglish:      gladysAmericanEnglishName,
		language.CanadianFrench:       gladysCanadianFrenchName,
		language.Dutch:                gladysDutchName,
		language.French:               gladysFrenchName,
		language.German:               gladysGermanName,
		language.Italian:              gladysItalianName,
		language.Japanese:             gladysJapaneseName,
		language.Korean:               gladysKoreanName,
		language.LatinAmericanSpanish: gladysLatinAmericanSpanishName,
		language.Russian:              gladysRussianName,
		language.SimplifiedChinese:    gladysSimplifiedChineseName,
		language.Spanish:              gladysSpanishName,
		language.TraditionalChinese:   gladysTraditionalChineseName}
)

var (
	gladysCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: gladysBirthday,
		Code:     gladysCode,
		Key:      character.Gladys,
		Gender:   gender.Female,
		Name:     gladysName}
)

var (
	gladysAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "stretch"}

	gladysCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "codêt"}

	gladysDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "en strek"}

	gladysFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "codêt"}

	gladysGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ajajajjj"}

	gladysItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "stretch"}

	gladysJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですの"}

	gladysKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇지라"}

	gladysLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "plumiplú"}

	gladysRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "растяжка"}

	gladysSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "千岁"}

	gladysSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "plumillas"}

	gladysTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "千歲"}
)

var (
	gladysPhrase = nook.Languages{
		language.AmericanEnglish:      gladysAmericanEnglishPhrase,
		language.CanadianFrench:       gladysCanadianFrenchPhrase,
		language.Dutch:                gladysDutchPhrase,
		language.French:               gladysFrenchPhrase,
		language.German:               gladysGermanPhrase,
		language.Italian:              gladysItalianPhrase,
		language.Japanese:             gladysJapanesePhrase,
		language.Korean:               gladysKoreanPhrase,
		language.LatinAmericanSpanish: gladysLatinAmericanSpanishPhrase,
		language.Russian:              gladysRussianPhrase,
		language.SimplifiedChinese:    gladysSimplifiedChinesePhrase,
		language.Spanish:              gladysSpanishPhrase,
		language.TraditionalChinese:   gladysTraditionalChinesePhrase}
)

var (
	Gladys = nook.Villager{
		Character:   gladysCharacter,
		Personality: personality.Normal,
		Phrase:      gladysPhrase}
)
