package duck

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
	pompomBirthday = nook.Birthday{
		Day:   11,
		Month: time.February}
)

var (
	pompomCode = nook.Code{
		Value: "duk05"}
)

var (
	pompomAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pompom"}

	pompomCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pompon"}

	pompomDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pompom"}

	pompomFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pompon"}

	pompomGermanName = nook.Name{
		Language: language.German,
		Value:    "Erika"}

	pompomItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Erica"}

	pompomJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のりっぺ"}

	pompomKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "주디"}

	pompomLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Flopi"}

	pompomRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Помпом"}

	pompomSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "海苔裴"}

	pompomSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Flopi"}

	pompomTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "海苔裴"}
)

var (
	pompomName = nook.Languages{
		language.AmericanEnglish:      pompomAmericanEnglishName,
		language.CanadianFrench:       pompomCanadianFrenchName,
		language.Dutch:                pompomDutchName,
		language.French:               pompomFrenchName,
		language.German:               pompomGermanName,
		language.Italian:              pompomItalianName,
		language.Japanese:             pompomJapaneseName,
		language.Korean:               pompomKoreanName,
		language.LatinAmericanSpanish: pompomLatinAmericanSpanishName,
		language.Russian:              pompomRussianName,
		language.SimplifiedChinese:    pompomSimplifiedChineseName,
		language.Spanish:              pompomSpanishName,
		language.TraditionalChinese:   pompomTraditionalChineseName}
)

var (
	pompomCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: pompomBirthday,
		Code:     pompomCode,
		Key:      character.Pompom,
		Gender:   gender.Female,
		Name:     pompomName}
)

var (
	pompomAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "rah rah"}

	pompomCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "rah rah"}

	pompomDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hup hup"}

	pompomFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "rah rah"}

	pompomGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krah"}

	pompomItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quoquiquà"}

	pompomJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だっピ"}

	pompomKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "다삐"}

	pompomLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ra-rá"}

	pompomRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "крякушки"}

	pompomSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "裴裴"}

	pompomSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sinplumas"}

	pompomTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "裴裴"}
)

var (
	pompomPhrase = nook.Languages{
		language.AmericanEnglish:      pompomAmericanEnglishPhrase,
		language.CanadianFrench:       pompomCanadianFrenchPhrase,
		language.Dutch:                pompomDutchPhrase,
		language.French:               pompomFrenchPhrase,
		language.German:               pompomGermanPhrase,
		language.Italian:              pompomItalianPhrase,
		language.Japanese:             pompomJapanesePhrase,
		language.Korean:               pompomKoreanPhrase,
		language.LatinAmericanSpanish: pompomLatinAmericanSpanishPhrase,
		language.Russian:              pompomRussianPhrase,
		language.SimplifiedChinese:    pompomSimplifiedChinesePhrase,
		language.Spanish:              pompomSpanishPhrase,
		language.TraditionalChinese:   pompomTraditionalChinesePhrase}
)

var (
	Pompom = nook.Villager{
		Character:   pompomCharacter,
		Personality: personality.Peppy,
		Phrase:      pompomPhrase}
)
