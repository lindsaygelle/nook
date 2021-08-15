package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	jacquesBirthday = nook.Birthday{
		Day:   22,
		Month: time.June}
)

var (
	jacquesCode = nook.Code{
		Value: "brd16"}
)

var (
	jacquesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jacques"}

	jacquesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jackywhy not"}

	jacquesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jacqueszut alors"}

	jacquesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jackyploc ploc"}

	jacquesGermanName = nook.Name{
		Language: language.German,
		Value:    "Pierrepiepiep"}

	jacquesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zampieropio pio"}

	jacquesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジョッキーチェケラ"}

	jacquesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쪼끼체키라웃"}

	jacquesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gorrelmochapó"}

	jacquesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Жакфу ты"}

	jacquesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "赵奇唷唷唷"}

	jacquesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gorrelmochapó"}

	jacquesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "趙奇唷唷唷"}
)

var (
	jacquesName = nook.Languages{
		language.AmericanEnglish:      jacquesAmericanEnglishName,
		language.CanadianFrench:       jacquesCanadianFrenchName,
		language.Dutch:                jacquesDutchName,
		language.French:               jacquesFrenchName,
		language.German:               jacquesGermanName,
		language.Italian:              jacquesItalianName,
		language.Japanese:             jacquesJapaneseName,
		language.Korean:               jacquesKoreanName,
		language.LatinAmericanSpanish: jacquesLatinAmericanSpanishName,
		language.Russian:              jacquesRussianName,
		language.SimplifiedChinese:    jacquesSimplifiedChineseName,
		language.Spanish:              jacquesSpanishName,
		language.TraditionalChinese:   jacquesTraditionalChineseName}
)

var (
	jacquesCharacter = nook.Character{
		Animal:   Bird,
		Birthday: jacquesBirthday,
		Code:     jacquesCode,
		Gender:   nook.Male,
		Name:     jacquesName}
)

var (
	jacquesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zut alors"}

	jacquesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "why not"}

	jacquesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zut alors"}

	jacquesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ploc ploc"}

	jacquesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "piepiep"}

	jacquesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pio pio"}

	jacquesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "チェケラ"}

	jacquesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "체키라웃"}

	jacquesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chapó"}

	jacquesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фу ты"}

	jacquesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唷唷唷"}

	jacquesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chapó"}

	jacquesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唷唷唷"}
)

var (
	jacquesPhrase = nook.Languages{
		language.AmericanEnglish:      jacquesAmericanEnglishName,
		language.CanadianFrench:       jacquesCanadianFrenchName,
		language.Dutch:                jacquesDutchName,
		language.French:               jacquesFrenchName,
		language.German:               jacquesGermanName,
		language.Italian:              jacquesItalianName,
		language.Japanese:             jacquesJapaneseName,
		language.Korean:               jacquesKoreanName,
		language.LatinAmericanSpanish: jacquesLatinAmericanSpanishName,
		language.Russian:              jacquesRussianName,
		language.SimplifiedChinese:    jacquesSimplifiedChineseName,
		language.Spanish:              jacquesSpanishName,
		language.TraditionalChinese:   jacquesTraditionalChineseName}
)

var (
	Jacques = nook.Villager{
		Character:   jacquesCharacter,
		Personality: nook.Smug,
		Phrase:      jacquesPhrase}
)
