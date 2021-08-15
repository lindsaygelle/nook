package cow

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pattyBirthday = nook.Birthday{
		Day:   10,
		Month: time.May}
)

var (
	pattyCode = nook.Code{
		Value: "cow00"}
)

var (
	pattyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Patty"}

	pattyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Margauxoui oui"}

	pattyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pattykoebeest"}

	pattyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Margauxallez"}

	pattyGermanName = nook.Name{
		Language: language.German,
		Value:    "Patriciamuuhna"}

	pattyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pattytruuulà"}

	pattyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カルピだモー"}

	pattyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "밀크음메"}

	pattyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vacarenamuuuu"}

	pattyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пэттиму-му"}

	pattySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿排牟"}

	pattySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vacarenamuuuu"}

	pattyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿排牟"}
)

var (
	pattyName = nook.Languages{
		language.AmericanEnglish:      pattyAmericanEnglishName,
		language.CanadianFrench:       pattyCanadianFrenchName,
		language.Dutch:                pattyDutchName,
		language.French:               pattyFrenchName,
		language.German:               pattyGermanName,
		language.Italian:              pattyItalianName,
		language.Japanese:             pattyJapaneseName,
		language.Korean:               pattyKoreanName,
		language.LatinAmericanSpanish: pattyLatinAmericanSpanishName,
		language.Russian:              pattyRussianName,
		language.SimplifiedChinese:    pattySimplifiedChineseName,
		language.Spanish:              pattySpanishName,
		language.TraditionalChinese:   pattyTraditionalChineseName}
)

var (
	pattyCharacter = nook.Character{
		Animal:   Cow,
		Birthday: pattyBirthday,
		Code:     pattyCode,
		Gender:   nook.Female,
		Name:     pattyName}
)

var (
	pattyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "how-now"}

	pattyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "oui oui"}

	pattyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "koebeest"}

	pattyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "allez"}

	pattyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "muuhna"}

	pattyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "truuulà"}

	pattyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だモー"}

	pattyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "음메"}

	pattyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "muuuu"}

	pattyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "му-му"}

	pattySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "牟"}

	pattySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muuuu"}

	pattyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "牟"}
)

var (
	pattyPhrase = nook.Languages{
		language.AmericanEnglish:      pattyAmericanEnglishName,
		language.CanadianFrench:       pattyCanadianFrenchName,
		language.Dutch:                pattyDutchName,
		language.French:               pattyFrenchName,
		language.German:               pattyGermanName,
		language.Italian:              pattyItalianName,
		language.Japanese:             pattyJapaneseName,
		language.Korean:               pattyKoreanName,
		language.LatinAmericanSpanish: pattyLatinAmericanSpanishName,
		language.Russian:              pattyRussianName,
		language.SimplifiedChinese:    pattySimplifiedChineseName,
		language.Spanish:              pattySpanishName,
		language.TraditionalChinese:   pattyTraditionalChineseName}
)

var (
	Patty = nook.Villager{
		Character:   pattyCharacter,
		Personality: nook.Peppy,
		Phrase:      pattyPhrase}
)
