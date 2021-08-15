package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	genDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	genFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	genGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	genItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	genJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゲン"}

	genKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	genLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	genRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	genSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	genSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	genTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Sheep,
		Birthday: genBirthday,
		Code:     genCode,
		Gender:   nook.Male,
		Name:     genName}
)

var (
	genAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	genCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	genDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	genFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	genGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	genItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	genJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "っしゃあ"}

	genKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	genLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	genRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	genSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	genSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	genTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Jock,
		Phrase:      genPhrase}
)
