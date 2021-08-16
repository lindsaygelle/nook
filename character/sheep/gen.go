package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	genBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	genCode = nook.Code{
		Value: ""}
)

var (
	genAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gen"}

	genCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	genDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	genFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	genGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	genItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	genJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゲン"}

	genKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	genLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	genRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	genSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	genSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	genTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	genName = nook.Languages{
		language.AmericanEnglish:      genAmericanEnglishName,
		language.CanadianFrench:       genCanadianFrenchName,
		language.Dutch:                genDutchName,
		language.French:               genFrenchName,
		language.German:               genGermanName,
		language.Italian:              genItalianName,
		language.Japanese:             genJapaneseName,
		language.Korean:               genKoreanName,
		language.LatinAmericanSpanish: genLatinAmericanSpanishName,
		language.Russian:              genRussianName,
		language.SimplifiedChinese:    genSimplifiedChineseName,
		language.Spanish:              genSpanishName,
		language.TraditionalChinese:   genTraditionalChineseName}
)

var (
	genCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: genBirthday,
		Code:     genCode,
		Gender:   gender.Male,
		Name:     genName}
)

var (
	genAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "っしゃあ"}

	genCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	genDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	genFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	genGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	genItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	genJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "っしゃあ"}

	genKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	genLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	genRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	genSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	genSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	genTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	genPhrase = nook.Languages{
		language.AmericanEnglish:      genAmericanEnglishName,
		language.CanadianFrench:       genCanadianFrenchName,
		language.Dutch:                genDutchName,
		language.French:               genFrenchName,
		language.German:               genGermanName,
		language.Italian:              genItalianName,
		language.Japanese:             genJapaneseName,
		language.Korean:               genKoreanName,
		language.LatinAmericanSpanish: genLatinAmericanSpanishName,
		language.Russian:              genRussianName,
		language.SimplifiedChinese:    genSimplifiedChineseName,
		language.Spanish:              genSpanishName,
		language.TraditionalChinese:   genTraditionalChineseName}
)

var (
	Gen = nook.Villager{
		Character:   genCharacter,
		Personality: personality.Jock,
		Phrase:      genPhrase}
)
