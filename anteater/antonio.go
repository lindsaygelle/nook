package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	antonioBirthday = nook.Birthday{
		Day:   20,
		Month: time.October}
)

var (
	antonioCode = nook.Code{
		Value: "ant01"}
)

var (
	antonioAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Antonio"}

	antonioCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Antoniopouet"}

	antonioDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Antoniosnuit"}

	antonioFrenchName = nook.Name{
		Language: language.French,
		Value:    "Antoniopouet"}

	antonioGermanName = nook.Name{
		Language: language.German,
		Value:    "Siggischlürrrf"}

	antonioItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Antoniohonk"}

	antonioJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マコトホントに"}

	antonioKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "퍼머거진짜로"}

	antonioLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Antoniofufuf"}

	antonioRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Антониого-го-го"}

	antonioSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿诚真诚"}

	antonioSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Antoniofufuf"}

	antonioTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿誠真誠"}
)

var (
	antonioName = nook.Languages{
		language.AmericanEnglish:      antonioAmericanEnglishName,
		language.CanadianFrench:       antonioCanadianFrenchName,
		language.Dutch:                antonioDutchName,
		language.French:               antonioFrenchName,
		language.German:               antonioGermanName,
		language.Italian:              antonioItalianName,
		language.Japanese:             antonioJapaneseName,
		language.Korean:               antonioKoreanName,
		language.LatinAmericanSpanish: antonioLatinAmericanSpanishName,
		language.Russian:              antonioRussianName,
		language.SimplifiedChinese:    antonioSimplifiedChineseName,
		language.Spanish:              antonioSpanishName,
		language.TraditionalChinese:   antonioTraditionalChineseName}
)

var (
	antonioCharacter = nook.Character{
		Animal:   Anteater,
		Birthday: antonioBirthday,
		Code:     antonioCode,
		Gender:   nook.Male,
		Name:     antonioName}
)

var (
	antonioAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "honk"}

	antonioCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pouet"}

	antonioDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuit"}

	antonioFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pouet"}

	antonioGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schlürrrf"}

	antonioItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "honk"}

	antonioJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ホントに"}

	antonioKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "진짜로"}

	antonioLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fufuf"}

	antonioRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "го-го-го"}

	antonioSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "真诚"}

	antonioSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fufuf"}

	antonioTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "真誠"}
)

var (
	antonioPhrase = nook.Languages{
		language.AmericanEnglish:      antonioAmericanEnglishName,
		language.CanadianFrench:       antonioCanadianFrenchName,
		language.Dutch:                antonioDutchName,
		language.French:               antonioFrenchName,
		language.German:               antonioGermanName,
		language.Italian:              antonioItalianName,
		language.Japanese:             antonioJapaneseName,
		language.Korean:               antonioKoreanName,
		language.LatinAmericanSpanish: antonioLatinAmericanSpanishName,
		language.Russian:              antonioRussianName,
		language.SimplifiedChinese:    antonioSimplifiedChineseName,
		language.Spanish:              antonioSpanishName,
		language.TraditionalChinese:   antonioTraditionalChineseName}
)

var (
	Antonio = nook.Villager{
		Character:   antonioCharacter,
		Personality: nook.Jock,
		Phrase:      antonioPhrase}
)
