package eagle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	ameliaBirthday = nook.Birthday{
		Day:   19,
		Month: time.November}
)

var (
	ameliaCode = nook.Code{
		Value: "pbr01"}
)

var (
	ameliaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Amelia"}

	ameliaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Aurélie"}

	ameliaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Amelia"}

	ameliaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Aurélie"}

	ameliaGermanName = nook.Name{
		Language: language.German,
		Value:    "Adelheid"}

	ameliaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Amelia"}

	ameliaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アンデス"}

	ameliaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안데스"}

	ameliaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Amelia"}

	ameliaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Амелия"}

	ameliaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安地斯"}

	ameliaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Amelia"}

	ameliaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安地斯"}
)

var (
	ameliaName = nook.Languages{
		language.AmericanEnglish:      ameliaAmericanEnglishName,
		language.CanadianFrench:       ameliaCanadianFrenchName,
		language.Dutch:                ameliaDutchName,
		language.French:               ameliaFrenchName,
		language.German:               ameliaGermanName,
		language.Italian:              ameliaItalianName,
		language.Japanese:             ameliaJapaneseName,
		language.Korean:               ameliaKoreanName,
		language.LatinAmericanSpanish: ameliaLatinAmericanSpanishName,
		language.Russian:              ameliaRussianName,
		language.SimplifiedChinese:    ameliaSimplifiedChineseName,
		language.Spanish:              ameliaSpanishName,
		language.TraditionalChinese:   ameliaTraditionalChineseName}
)

var (
	ameliaCharacter = nook.Character{
		Animal:   Eagle,
		Birthday: ameliaBirthday,
		Code:     ameliaCode,
		Gender:   nook.Female,
		Name:     ameliaName}
)

var (
	ameliaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "eaglet"}

	ameliaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "croooa"}

	ameliaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "broeder"}

	ameliaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croooa"}

	ameliaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "adlerauge"}

	ameliaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mini mini"}

	ameliaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "カラカラ"}

	ameliaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "엄멈머"}

	ameliaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cro-ah"}

	ameliaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вот"}

	ameliaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卡拉卡拉"}

	ameliaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chinchilla"}

	ameliaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "卡拉卡拉"}
)

var (
	ameliaPhrase = nook.Languages{
		language.AmericanEnglish:      ameliaAmericanEnglishName,
		language.CanadianFrench:       ameliaCanadianFrenchName,
		language.Dutch:                ameliaDutchName,
		language.French:               ameliaFrenchName,
		language.German:               ameliaGermanName,
		language.Italian:              ameliaItalianName,
		language.Japanese:             ameliaJapaneseName,
		language.Korean:               ameliaKoreanName,
		language.LatinAmericanSpanish: ameliaLatinAmericanSpanishName,
		language.Russian:              ameliaRussianName,
		language.SimplifiedChinese:    ameliaSimplifiedChineseName,
		language.Spanish:              ameliaSpanishName,
		language.TraditionalChinese:   ameliaTraditionalChineseName}
)

var (
	Amelia = nook.Villager{
		Character:   ameliaCharacter,
		Personality: nook.Snooty,
		Phrase:      ameliaPhrase}
)
