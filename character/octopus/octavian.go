package octopus

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
	octavianBirthday = nook.Birthday{
		Day:   20,
		Month: time.September}
)

var (
	octavianCode = nook.Code{
		Value: "ocp00"}
)

var (
	octavianAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Octavian"}

	octavianCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Octave"}

	octavianDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Octavian"}

	octavianFrenchName = nook.Name{
		Language: language.French,
		Value:    "Octave"}

	octavianGermanName = nook.Name{
		Language: language.German,
		Value:    "Ottfried"}

	octavianItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ottavio"}

	octavianJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "おくたろう"}

	octavianKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "문복"}

	octavianLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Octavio"}

	octavianRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Октавиан"}

	octavianSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "张瑜"}

	octavianSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Octavio"}

	octavianTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "張瑜"}
)

var (
	octavianName = nook.Languages{
		language.AmericanEnglish:      octavianAmericanEnglishName,
		language.CanadianFrench:       octavianCanadianFrenchName,
		language.Dutch:                octavianDutchName,
		language.French:               octavianFrenchName,
		language.German:               octavianGermanName,
		language.Italian:              octavianItalianName,
		language.Japanese:             octavianJapaneseName,
		language.Korean:               octavianKoreanName,
		language.LatinAmericanSpanish: octavianLatinAmericanSpanishName,
		language.Russian:              octavianRussianName,
		language.SimplifiedChinese:    octavianSimplifiedChineseName,
		language.Spanish:              octavianSpanishName,
		language.TraditionalChinese:   octavianTraditionalChineseName}
)

var (
	octavianCharacter = nook.Character{
		Animal:   animal.Octopus,
		Birthday: octavianBirthday,
		Code:     octavianCode,
		Key:      character.Octavian,
		Gender:   gender.Male,
		Name:     octavianName}
)

var (
	octavianAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sucker"}

	octavianCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chenapan"}

	octavianDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zuignap"}

	octavianFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chenapan"}

	octavianGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schling"}

	octavianItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "perdindi"}

	octavianJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "タコ"}

	octavianKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쭉쭉"}

	octavianLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "glop-glop"}

	octavianRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "все пропало"}

	octavianSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "章鱼"}

	octavianSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chupatinta"}

	octavianTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "章魚"}
)

var (
	octavianPhrase = nook.Languages{
		language.AmericanEnglish:      octavianAmericanEnglishPhrase,
		language.CanadianFrench:       octavianCanadianFrenchPhrase,
		language.Dutch:                octavianDutchPhrase,
		language.French:               octavianFrenchPhrase,
		language.German:               octavianGermanPhrase,
		language.Italian:              octavianItalianPhrase,
		language.Japanese:             octavianJapanesePhrase,
		language.Korean:               octavianKoreanPhrase,
		language.LatinAmericanSpanish: octavianLatinAmericanSpanishPhrase,
		language.Russian:              octavianRussianPhrase,
		language.SimplifiedChinese:    octavianSimplifiedChinesePhrase,
		language.Spanish:              octavianSpanishPhrase,
		language.TraditionalChinese:   octavianTraditionalChinesePhrase}
)

var (
	Octavian = nook.Villager{
		Character:   octavianCharacter,
		Personality: personality.Cranky,
		Phrase:      octavianPhrase}
)
