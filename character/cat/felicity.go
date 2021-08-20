package cat

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
	felicityBirthday = nook.Birthday{
		Day:   30,
		Month: time.March}
)

var (
	felicityCode = nook.Code{
		Value: "cat17"}
)

var (
	felicityAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Felicity"}

	felicityCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maud"}

	felicityDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Felicity"}

	felicityFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maud"}

	felicityGermanName = nook.Name{
		Language: language.German,
		Value:    "Minka"}

	felicityItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Milly"}

	felicityJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みかっち"}

	felicityKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "예링"}

	felicityLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Micha"}

	felicityRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фелисити"}

	felicitySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "美佳"}

	felicitySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Micha"}

	felicityTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "美佳"}
)

var (
	felicityName = nook.Languages{
		language.AmericanEnglish:      felicityAmericanEnglishName,
		language.CanadianFrench:       felicityCanadianFrenchName,
		language.Dutch:                felicityDutchName,
		language.French:               felicityFrenchName,
		language.German:               felicityGermanName,
		language.Italian:              felicityItalianName,
		language.Japanese:             felicityJapaneseName,
		language.Korean:               felicityKoreanName,
		language.LatinAmericanSpanish: felicityLatinAmericanSpanishName,
		language.Russian:              felicityRussianName,
		language.SimplifiedChinese:    felicitySimplifiedChineseName,
		language.Spanish:              felicitySpanishName,
		language.TraditionalChinese:   felicityTraditionalChineseName}
)

var (
	felicityCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: felicityBirthday,
		Code:     felicityCode,
		Key:      character.Felicity,
		Gender:   gender.Female,
		Name:     felicityName,
		Special:  false}
)

var (
	felicityAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mimimi"}

	felicityCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chatounet"}

	felicityDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mimimi"}

	felicityFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chatounet"}

	felicityGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mauzi"}

	felicityItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "meeeow"}

	felicityJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ね"}

	felicityKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우앙"}

	felicityLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "miauito"}

	felicityRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ми-ми-ми"}

	felicitySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哪"}

	felicitySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "miauito"}

	felicityTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哪"}
)

var (
	felicityPhrase = nook.Languages{
		language.AmericanEnglish:      felicityAmericanEnglishPhrase,
		language.CanadianFrench:       felicityCanadianFrenchPhrase,
		language.Dutch:                felicityDutchPhrase,
		language.French:               felicityFrenchPhrase,
		language.German:               felicityGermanPhrase,
		language.Italian:              felicityItalianPhrase,
		language.Japanese:             felicityJapanesePhrase,
		language.Korean:               felicityKoreanPhrase,
		language.LatinAmericanSpanish: felicityLatinAmericanSpanishPhrase,
		language.Russian:              felicityRussianPhrase,
		language.SimplifiedChinese:    felicitySimplifiedChinesePhrase,
		language.Spanish:              felicitySpanishPhrase,
		language.TraditionalChinese:   felicityTraditionalChinesePhrase}
)

var (
	Felicity = nook.Villager{
		Character:   felicityCharacter,
		Personality: personality.Peppy,
		Phrase:      felicityPhrase}
)
