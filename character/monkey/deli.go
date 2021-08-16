package monkey

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	deliBirthday = nook.Birthday{
		Day:   24,
		Month: time.May}
)

var (
	deliCode = nook.Code{
		Value: "mnk08"}
)

var (
	deliAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Deli"}

	deliCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Magogo"}

	deliDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Deli"}

	deliFrenchName = nook.Name{
		Language: language.French,
		Value:    "Magogo"}

	deliGermanName = nook.Name{
		Language: language.German,
		Value:    "Anton"}

	deliItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bobo"}

	deliJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "デリー"}

	deliKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "델리"}

	deliLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cana"}

	deliRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дели"}

	deliSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "德里"}

	deliSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cana"}

	deliTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "德里"}
)

var (
	deliName = nook.Languages{
		language.AmericanEnglish:      deliAmericanEnglishName,
		language.CanadianFrench:       deliCanadianFrenchName,
		language.Dutch:                deliDutchName,
		language.French:               deliFrenchName,
		language.German:               deliGermanName,
		language.Italian:              deliItalianName,
		language.Japanese:             deliJapaneseName,
		language.Korean:               deliKoreanName,
		language.LatinAmericanSpanish: deliLatinAmericanSpanishName,
		language.Russian:              deliRussianName,
		language.SimplifiedChinese:    deliSimplifiedChineseName,
		language.Spanish:              deliSpanishName,
		language.TraditionalChinese:   deliTraditionalChineseName}
)

var (
	deliCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: deliBirthday,
		Code:     deliCode,
		Gender:   gender.Male,
		Name:     deliName}
)

var (
	deliAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "monch"}

	deliCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "liaaaane"}

	deliDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "smak"}

	deliFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "liaaaane"}

	deliGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "miekmiek"}

	deliItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "panzé"}

	deliJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だぜ"}

	deliKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "흠냐리"}

	deliLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñim-ñim"}

	deliRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чавк-чавк"}

	deliSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "去吧"}

	deliSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "millón"}

	deliTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "去吧"}
)

var (
	deliPhrase = nook.Languages{
		language.AmericanEnglish:      deliAmericanEnglishName,
		language.CanadianFrench:       deliCanadianFrenchName,
		language.Dutch:                deliDutchName,
		language.French:               deliFrenchName,
		language.German:               deliGermanName,
		language.Italian:              deliItalianName,
		language.Japanese:             deliJapaneseName,
		language.Korean:               deliKoreanName,
		language.LatinAmericanSpanish: deliLatinAmericanSpanishName,
		language.Russian:              deliRussianName,
		language.SimplifiedChinese:    deliSimplifiedChineseName,
		language.Spanish:              deliSpanishName,
		language.TraditionalChinese:   deliTraditionalChineseName}
)

var (
	Deli = nook.Villager{
		Character:   deliCharacter,
		Personality: personality.Lazy,
		Phrase:      deliPhrase}
)
