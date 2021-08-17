package bearcub

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
	baroldBirthday = nook.Birthday{
		Day:   2,
		Month: time.March}
)

var (
	baroldCode = nook.Code{
		Value: "cbr16"}
)

var (
	baroldAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Barold"}

	baroldCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Manu"}

	baroldDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Barold"}

	baroldFrenchName = nook.Name{
		Language: language.French,
		Value:    "Manu"}

	baroldGermanName = nook.Name{
		Language: language.German,
		Value:    "Hubert"}

	baroldItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Placido"}

	baroldJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ニッシー"}

	baroldKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곰시"}

	baroldLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Plácido"}

	baroldRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Барольд"}

	baroldSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿西"}

	baroldSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Plácido"}

	baroldTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿西"}
)

var (
	baroldName = nook.Languages{
		language.AmericanEnglish:      baroldAmericanEnglishName,
		language.CanadianFrench:       baroldCanadianFrenchName,
		language.Dutch:                baroldDutchName,
		language.French:               baroldFrenchName,
		language.German:               baroldGermanName,
		language.Italian:              baroldItalianName,
		language.Japanese:             baroldJapaneseName,
		language.Korean:               baroldKoreanName,
		language.LatinAmericanSpanish: baroldLatinAmericanSpanishName,
		language.Russian:              baroldRussianName,
		language.SimplifiedChinese:    baroldSimplifiedChineseName,
		language.Spanish:              baroldSpanishName,
		language.TraditionalChinese:   baroldTraditionalChineseName}
)

var (
	baroldCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: baroldBirthday,
		Code:     baroldCode,
		Key:      character.Barold,
		Gender:   gender.Male,
		Name:     baroldName}
)

var (
	baroldAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cubby"}

	baroldCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "dodododu"}

	baroldDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "welpje"}

	baroldFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "dodododu"}

	baroldGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mwuuu"}

	baroldItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "babalù"}

	baroldJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "いっそ"}

	baroldKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아니그냥"}

	baroldLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "apurri"}

	baroldRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "берлога"}

	baroldSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "算了"}

	baroldSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "apurri"}

	baroldTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "算了"}
)

var (
	baroldPhrase = nook.Languages{
		language.AmericanEnglish:      baroldAmericanEnglishPhrase,
		language.CanadianFrench:       baroldCanadianFrenchPhrase,
		language.Dutch:                baroldDutchPhrase,
		language.French:               baroldFrenchPhrase,
		language.German:               baroldGermanPhrase,
		language.Italian:              baroldItalianPhrase,
		language.Japanese:             baroldJapanesePhrase,
		language.Korean:               baroldKoreanPhrase,
		language.LatinAmericanSpanish: baroldLatinAmericanSpanishPhrase,
		language.Russian:              baroldRussianPhrase,
		language.SimplifiedChinese:    baroldSimplifiedChinesePhrase,
		language.Spanish:              baroldSpanishPhrase,
		language.TraditionalChinese:   baroldTraditionalChinesePhrase}
)

var (
	Barold = nook.Villager{
		Character:   baroldCharacter,
		Personality: personality.Lazy,
		Phrase:      baroldPhrase}
)
