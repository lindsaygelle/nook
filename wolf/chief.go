package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Chefmmmph"}

	chiefDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chiefhuil"}

	chiefFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chefmmmph"}

	chiefGermanName = nook.Name{
		Language: language.German,
		Value:    "Saschaharrumff"}

	chiefItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Artiglioarrrgh"}

	chiefJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チーフやんか"}

	chiefKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "대장그렇잖여"}

	chiefLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Zoiloejem"}

	chiefRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чиффыр-р-рф"}

	chiefSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "施傅不是嘛"}

	chiefSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Zoiloejem"}

	chiefTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "施傅不是嘛"}
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
		Animal:   Wolf,
		Birthday: chiefBirthday,
		Code:     chiefCode,
		Gender:   nook.Male,
		Name:     chiefName}
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
	Chief = nook.Villager{
		Character:   chiefCharacter,
		Personality: nook.Cranky,
		Phrase:      chiefPhrase}
)
