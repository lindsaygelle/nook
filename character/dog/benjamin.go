package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	benjaminBirthday = nook.Birthday{
		Day:   3,
		Month: time.August}
)

var (
	benjaminCode = nook.Code{
		Value: "dog16"}
)

var (
	benjaminAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Benjamin"}

	benjaminCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bernardo"}

	benjaminDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Benjamin"}

	benjaminFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bernardo"}

	benjaminGermanName = nook.Name{
		Language: language.German,
		Value:    "Wastl"}

	benjaminItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bernardo"}

	benjaminJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハチ"}

	benjaminKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "땡칠이"}

	benjaminLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bernardo"}

	benjaminRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бернардо"}

	benjaminSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小八"}

	benjaminSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bernardo"}

	benjaminTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小八"}
)

var (
	benjaminName = nook.Languages{
		language.AmericanEnglish:      benjaminAmericanEnglishName,
		language.CanadianFrench:       benjaminCanadianFrenchName,
		language.Dutch:                benjaminDutchName,
		language.French:               benjaminFrenchName,
		language.German:               benjaminGermanName,
		language.Italian:              benjaminItalianName,
		language.Japanese:             benjaminJapaneseName,
		language.Korean:               benjaminKoreanName,
		language.LatinAmericanSpanish: benjaminLatinAmericanSpanishName,
		language.Russian:              benjaminRussianName,
		language.SimplifiedChinese:    benjaminSimplifiedChineseName,
		language.Spanish:              benjaminSpanishName,
		language.TraditionalChinese:   benjaminTraditionalChineseName}
)

var (
	benjaminCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: benjaminBirthday,
		Code:     benjaminCode,
		Gender:   gender.Male,
		Name:     benjaminName}
)

var (
	benjaminAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "alrighty"}

	benjaminCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "croquettes"}

	benjaminDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oké dan"}

	benjaminFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croquettes"}

	benjaminGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "winsel"}

	benjaminItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bauuau"}

	benjaminJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ではでは"}

	benjaminKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그럼그럼"}

	benjaminLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tirirí"}

	benjaminRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пор-рядок"}

	benjaminSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "那么那么"}

	benjaminSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tirirí"}

	benjaminTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "那麼那麼"}
)

var (
	benjaminPhrase = nook.Languages{
		language.AmericanEnglish:      benjaminAmericanEnglishName,
		language.CanadianFrench:       benjaminCanadianFrenchName,
		language.Dutch:                benjaminDutchName,
		language.French:               benjaminFrenchName,
		language.German:               benjaminGermanName,
		language.Italian:              benjaminItalianName,
		language.Japanese:             benjaminJapaneseName,
		language.Korean:               benjaminKoreanName,
		language.LatinAmericanSpanish: benjaminLatinAmericanSpanishName,
		language.Russian:              benjaminRussianName,
		language.SimplifiedChinese:    benjaminSimplifiedChineseName,
		language.Spanish:              benjaminSpanishName,
		language.TraditionalChinese:   benjaminTraditionalChineseName}
)

var (
	Benjamin = nook.Villager{
		Character:   benjaminCharacter,
		Personality: personality.Lazy,
		Phrase:      benjaminPhrase}
)
