package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Maudchatounet"}

	felicityDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Felicitymimimi"}

	felicityFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maudchatounet"}

	felicityGermanName = nook.Name{
		Language: language.German,
		Value:    "Minkamauzi"}

	felicityItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Millymeeeow"}

	felicityJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みかっちね"}

	felicityKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "예링우앙"}

	felicityLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Michamiauito"}

	felicityRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фелиситими-ми-ми"}

	felicitySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "美佳哪"}

	felicitySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Michamiauito"}

	felicityTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "美佳哪"}
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
		Animal:   Cat,
		Birthday: felicityBirthday,
		Code:     felicityCode,
		Gender:   nook.Female,
		Name:     felicityName}
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
	Felicity = nook.Villager{
		Character:   felicityCharacter,
		Personality: nook.Peppy,
		Phrase:      felicityPhrase}
)
