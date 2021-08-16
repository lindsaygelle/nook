package lion

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
	roryBirthday = nook.Birthday{
		Day:   7,
		Month: time.August}
)

var (
	roryCode = nook.Code{
		Value: "lon07"}
)

var (
	roryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rory"}

	roryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hercule"}

	roryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rory"}

	roryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hercule"}

	roryGermanName = nook.Name{
		Language: language.German,
		Value:    "Leon"}

	roryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ruggero"}

	roryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アーサー"}

	roryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아더"}

	roryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rogelio"}

	roryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рори"}

	rorySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亚瑟"}

	rorySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rogelio"}

	roryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亞瑟"}
)

var (
	roryName = nook.Languages{
		language.AmericanEnglish:      roryAmericanEnglishName,
		language.CanadianFrench:       roryCanadianFrenchName,
		language.Dutch:                roryDutchName,
		language.French:               roryFrenchName,
		language.German:               roryGermanName,
		language.Italian:              roryItalianName,
		language.Japanese:             roryJapaneseName,
		language.Korean:               roryKoreanName,
		language.LatinAmericanSpanish: roryLatinAmericanSpanishName,
		language.Russian:              roryRussianName,
		language.SimplifiedChinese:    rorySimplifiedChineseName,
		language.Spanish:              rorySpanishName,
		language.TraditionalChinese:   roryTraditionalChineseName}
)

var (
	roryCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: roryBirthday,
		Code:     roryCode,
		Key:      character.Rory,
		Gender:   gender.Male,
		Name:     roryName}
)

var (
	roryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "capital"}

	roryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "excellent"}

	roryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "koninklijk"}

	roryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "criniou"}

	roryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "graaah"}

	roryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "roar"}

	roryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ナハッ"}

	roryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "크릉"}

	roryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ticotaco"}

	roryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "капитально"}

	rorySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哪哈"}

	rorySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ticotaco"}

	roryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哪哈"}
)

var (
	roryPhrase = nook.Languages{
		language.AmericanEnglish:      roryAmericanEnglishName,
		language.CanadianFrench:       roryCanadianFrenchName,
		language.Dutch:                roryDutchName,
		language.French:               roryFrenchName,
		language.German:               roryGermanName,
		language.Italian:              roryItalianName,
		language.Japanese:             roryJapaneseName,
		language.Korean:               roryKoreanName,
		language.LatinAmericanSpanish: roryLatinAmericanSpanishName,
		language.Russian:              roryRussianName,
		language.SimplifiedChinese:    rorySimplifiedChineseName,
		language.Spanish:              rorySpanishName,
		language.TraditionalChinese:   roryTraditionalChineseName}
)

var (
	Rory = nook.Villager{
		Character:   roryCharacter,
		Personality: personality.Jock,
		Phrase:      roryPhrase}
)
