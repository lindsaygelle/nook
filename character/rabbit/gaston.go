package rabbit

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
	gastonBirthday = nook.Birthday{
		Day:   28,
		Month: time.October}
)

var (
	gastonCode = nook.Code{
		Value: "rbt04"}
)

var (
	gastonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gaston"}

	gastonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gaston"}

	gastonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gaston"}

	gastonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gaston"}

	gastonGermanName = nook.Name{
		Language: language.German,
		Value:    "Gaston"}

	gastonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gaston"}

	gastonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モサキチ"}

	gastonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "대길"}

	gastonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gastón"}

	gastonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гастон"}

	gastonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "猛兔"}

	gastonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gastón"}

	gastonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "猛兔"}
)

var (
	gastonName = nook.Languages{
		language.AmericanEnglish:      gastonAmericanEnglishName,
		language.CanadianFrench:       gastonCanadianFrenchName,
		language.Dutch:                gastonDutchName,
		language.French:               gastonFrenchName,
		language.German:               gastonGermanName,
		language.Italian:              gastonItalianName,
		language.Japanese:             gastonJapaneseName,
		language.Korean:               gastonKoreanName,
		language.LatinAmericanSpanish: gastonLatinAmericanSpanishName,
		language.Russian:              gastonRussianName,
		language.SimplifiedChinese:    gastonSimplifiedChineseName,
		language.Spanish:              gastonSpanishName,
		language.TraditionalChinese:   gastonTraditionalChineseName}
)

var (
	gastonCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: gastonBirthday,
		Code:     gastonCode,
		Key:      character.Gaston,
		Gender:   gender.Male,
		Name:     gastonName,
		Special:  false}
)

var (
	gastonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mon chou"}

	gastonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cahuète"}

	gastonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mon amour"}

	gastonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cahuète"}

	gastonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "herzblatt"}

	gastonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cherie"}

	gastonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ムホッ"}

	gastonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쿨럭"}

	gastonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mon chou"}

	gastonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мон зайшу"}

	gastonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唔贺"}

	gastonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mon chou"}

	gastonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唔賀"}
)

var (
	gastonPhrase = nook.Languages{
		language.AmericanEnglish:      gastonAmericanEnglishPhrase,
		language.CanadianFrench:       gastonCanadianFrenchPhrase,
		language.Dutch:                gastonDutchPhrase,
		language.French:               gastonFrenchPhrase,
		language.German:               gastonGermanPhrase,
		language.Italian:              gastonItalianPhrase,
		language.Japanese:             gastonJapanesePhrase,
		language.Korean:               gastonKoreanPhrase,
		language.LatinAmericanSpanish: gastonLatinAmericanSpanishPhrase,
		language.Russian:              gastonRussianPhrase,
		language.SimplifiedChinese:    gastonSimplifiedChinesePhrase,
		language.Spanish:              gastonSpanishPhrase,
		language.TraditionalChinese:   gastonTraditionalChinesePhrase}
)

var (
	Gaston = nook.Villager{
		Character:   gastonCharacter,
		Personality: personality.Cranky,
		Phrase:      gastonPhrase}
)
