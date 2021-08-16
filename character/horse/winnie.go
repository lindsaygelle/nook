package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	winnieBirthday = nook.Birthday{
		Day:   31,
		Month: time.January}
)

var (
	winnieCode = nook.Code{
		Value: "hrs05"}
)

var (
	winnieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Winnie"}

	winnieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Anne"}

	winnieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Winnie"}

	winnieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Anne"}

	winnieGermanName = nook.Name{
		Language: language.German,
		Value:    "Walli"}

	winnieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Clara"}

	winnieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マキバスター"}

	winnieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "카로틴"}

	winnieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Soonia"}

	winnieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Винни"}

	winnieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "马星星"}

	winnieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Soonia"}

	winnieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "馬星星"}
)

var (
	winnieName = nook.Languages{
		language.AmericanEnglish:      winnieAmericanEnglishName,
		language.CanadianFrench:       winnieCanadianFrenchName,
		language.Dutch:                winnieDutchName,
		language.French:               winnieFrenchName,
		language.German:               winnieGermanName,
		language.Italian:              winnieItalianName,
		language.Japanese:             winnieJapaneseName,
		language.Korean:               winnieKoreanName,
		language.LatinAmericanSpanish: winnieLatinAmericanSpanishName,
		language.Russian:              winnieRussianName,
		language.SimplifiedChinese:    winnieSimplifiedChineseName,
		language.Spanish:              winnieSpanishName,
		language.TraditionalChinese:   winnieTraditionalChineseName}
)

var (
	winnieCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: winnieBirthday,
		Code:     winnieCode,
		Gender:   gender.Female,
		Name:     winnieName}
)

var (
	winnieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hay-OK"}

	winnieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "youp là"}

	winnieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knol"}

	winnieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "youp là"}

	winnieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "heu-ja"}

	winnieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "biadina"}

	winnieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "みゃあ"}

	winnieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "당근당근"}

	winnieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "okeimakei"}

	winnieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тык-дык"}

	winnieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘶嘶"}

	winnieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "okeimakei"}

	winnieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘶嘶"}
)

var (
	winniePhrase = nook.Languages{
		language.AmericanEnglish:      winnieAmericanEnglishName,
		language.CanadianFrench:       winnieCanadianFrenchName,
		language.Dutch:                winnieDutchName,
		language.French:               winnieFrenchName,
		language.German:               winnieGermanName,
		language.Italian:              winnieItalianName,
		language.Japanese:             winnieJapaneseName,
		language.Korean:               winnieKoreanName,
		language.LatinAmericanSpanish: winnieLatinAmericanSpanishName,
		language.Russian:              winnieRussianName,
		language.SimplifiedChinese:    winnieSimplifiedChineseName,
		language.Spanish:              winnieSpanishName,
		language.TraditionalChinese:   winnieTraditionalChineseName}
)

var (
	Winnie = nook.Villager{
		Character:   winnieCharacter,
		Personality: personality.Peppy,
		Phrase:      winniePhrase}
)
