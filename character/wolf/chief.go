package wolf

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
	chiefBirthday = nook.Birthday{
		Day:   19,
		Month: time.December}
)

var (
	chiefCode = nook.Code{
		Value: "wol00"}
)

var (
	chiefAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chief"}

	chiefCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chef"}

	chiefDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chief"}

	chiefFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chef"}

	chiefGermanName = nook.Name{
		Language: language.German,
		Value:    "Sascha"}

	chiefItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Artiglio"}

	chiefJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チーフ"}

	chiefKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "대장"}

	chiefLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Zoilo"}

	chiefRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чиф"}

	chiefSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "施傅"}

	chiefSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Zoilo"}

	chiefTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "施傅"}
)

var (
	chiefName = nook.Languages{
		language.AmericanEnglish:      chiefAmericanEnglishName,
		language.CanadianFrench:       chiefCanadianFrenchName,
		language.Dutch:                chiefDutchName,
		language.French:               chiefFrenchName,
		language.German:               chiefGermanName,
		language.Italian:              chiefItalianName,
		language.Japanese:             chiefJapaneseName,
		language.Korean:               chiefKoreanName,
		language.LatinAmericanSpanish: chiefLatinAmericanSpanishName,
		language.Russian:              chiefRussianName,
		language.SimplifiedChinese:    chiefSimplifiedChineseName,
		language.Spanish:              chiefSpanishName,
		language.TraditionalChinese:   chiefTraditionalChineseName}
)

var (
	chiefCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: chiefBirthday,
		Code:     chiefCode,
		Key:      character.Chief,
		Gender:   gender.Male,
		Name:     chiefName,
		Special:  false}
)

var (
	chiefAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "harrumph"}

	chiefCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mmmph"}

	chiefDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "huil"}

	chiefFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mmmph"}

	chiefGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "harrumff"}

	chiefItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "arrrgh"}

	chiefJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やんか"}

	chiefKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇잖여"}

	chiefLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ejem"}

	chiefRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фыр-р-рф"}

	chiefSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不是嘛"}

	chiefSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ejem"}

	chiefTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不是嘛"}
)

var (
	chiefPhrase = nook.Languages{
		language.AmericanEnglish:      chiefAmericanEnglishPhrase,
		language.CanadianFrench:       chiefCanadianFrenchPhrase,
		language.Dutch:                chiefDutchPhrase,
		language.French:               chiefFrenchPhrase,
		language.German:               chiefGermanPhrase,
		language.Italian:              chiefItalianPhrase,
		language.Japanese:             chiefJapanesePhrase,
		language.Korean:               chiefKoreanPhrase,
		language.LatinAmericanSpanish: chiefLatinAmericanSpanishPhrase,
		language.Russian:              chiefRussianPhrase,
		language.SimplifiedChinese:    chiefSimplifiedChinesePhrase,
		language.Spanish:              chiefSpanishPhrase,
		language.TraditionalChinese:   chiefTraditionalChinesePhrase}
)

var (
	Chief = nook.Villager{
		Character:   chiefCharacter,
		Personality: personality.Cranky,
		Phrase:      chiefPhrase}
)
