package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	cyranoBirthday = nook.Birthday{
		Day:   9,
		Month: time.March}
)

var (
	cyranoCode = nook.Code{
		Value: "ant00"}
)

var (
	cyranoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cyrano"}

	cyranoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "CyranoATCHOUM"}

	cyranoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cyranoha-TSJOE"}

	cyranoFrenchName = nook.Name{
		Language: language.French,
		Value:    "CyranoATCHOUM"}

	cyranoGermanName = nook.Name{
		Language: language.German,
		Value:    "Theoschneuf"}

	cyranoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ciranoett-CCIÙ"}

	cyranoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "さくらじまでごわす"}

	cyranoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사지마임돠"}

	cyranoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ciranoachús"}

	cyranoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сираноапчхи"}

	cyranoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阳明有的"}

	cyranoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ciranoachús"}

	cyranoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "陽明有的"}
)

var (
	cyranoName = nook.Languages{
		language.AmericanEnglish:      cyranoAmericanEnglishName,
		language.CanadianFrench:       cyranoCanadianFrenchName,
		language.Dutch:                cyranoDutchName,
		language.French:               cyranoFrenchName,
		language.German:               cyranoGermanName,
		language.Italian:              cyranoItalianName,
		language.Japanese:             cyranoJapaneseName,
		language.Korean:               cyranoKoreanName,
		language.LatinAmericanSpanish: cyranoLatinAmericanSpanishName,
		language.Russian:              cyranoRussianName,
		language.SimplifiedChinese:    cyranoSimplifiedChineseName,
		language.Spanish:              cyranoSpanishName,
		language.TraditionalChinese:   cyranoTraditionalChineseName}
)

var (
	cyranoCharacter = nook.Character{
		Animal:   Anteater,
		Birthday: cyranoBirthday,
		Code:     cyranoCode,
		Gender:   nook.Male,
		Name:     cyranoName}
)

var (
	cyranoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ah-CHOO"}

	cyranoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ATCHOUM"}

	cyranoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ha-TSJOE"}

	cyranoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ATCHOUM"}

	cyranoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schneuf"}

	cyranoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ett-CCIÙ"}

	cyranoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でごわす"}

	cyranoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "임돠"}

	cyranoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "achús"}

	cyranoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "апчхи"}

	cyranoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "有的"}

	cyranoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "achús"}

	cyranoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "有的"}
)

var (
	cyranoPhrase = nook.Languages{
		language.AmericanEnglish:      cyranoAmericanEnglishName,
		language.CanadianFrench:       cyranoCanadianFrenchName,
		language.Dutch:                cyranoDutchName,
		language.French:               cyranoFrenchName,
		language.German:               cyranoGermanName,
		language.Italian:              cyranoItalianName,
		language.Japanese:             cyranoJapaneseName,
		language.Korean:               cyranoKoreanName,
		language.LatinAmericanSpanish: cyranoLatinAmericanSpanishName,
		language.Russian:              cyranoRussianName,
		language.SimplifiedChinese:    cyranoSimplifiedChineseName,
		language.Spanish:              cyranoSpanishName,
		language.TraditionalChinese:   cyranoTraditionalChineseName}
)

var (
	Cyrano = nook.Villager{
		Character:   cyranoCharacter,
		Personality: nook.Cranky,
		Phrase:      cyranoPhrase}
)
