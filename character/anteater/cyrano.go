package anteater

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
		Value:    "Cyrano"}

	cyranoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cyrano"}

	cyranoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cyrano"}

	cyranoGermanName = nook.Name{
		Language: language.German,
		Value:    "Theo"}

	cyranoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cirano"}

	cyranoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "さくらじま"}

	cyranoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사지마"}

	cyranoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cirano"}

	cyranoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сирано"}

	cyranoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阳明"}

	cyranoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cirano"}

	cyranoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "陽明"}
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
		Animal:   animal.Anteater,
		Birthday: cyranoBirthday,
		Code:     cyranoCode,
		Key:      character.Cyrano,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      cyranoAmericanEnglishPhrase,
		language.CanadianFrench:       cyranoCanadianFrenchPhrase,
		language.Dutch:                cyranoDutchPhrase,
		language.French:               cyranoFrenchPhrase,
		language.German:               cyranoGermanPhrase,
		language.Italian:              cyranoItalianPhrase,
		language.Japanese:             cyranoJapanesePhrase,
		language.Korean:               cyranoKoreanPhrase,
		language.LatinAmericanSpanish: cyranoLatinAmericanSpanishPhrase,
		language.Russian:              cyranoRussianPhrase,
		language.SimplifiedChinese:    cyranoSimplifiedChinesePhrase,
		language.Spanish:              cyranoSpanishPhrase,
		language.TraditionalChinese:   cyranoTraditionalChinesePhrase}
)

var (
	Cyrano = nook.Villager{
		Character:   cyranoCharacter,
		Personality: personality.Cranky,
		Phrase:      cyranoPhrase}
)
