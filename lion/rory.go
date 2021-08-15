package lion

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Herculeexcellent"}

	roryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rorykoninklijk"}

	roryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Herculecriniou"}

	roryGermanName = nook.Name{
		Language: language.German,
		Value:    "Leongraaah"}

	roryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ruggeroroar"}

	roryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アーサーナハッ"}

	roryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아더크릉"}

	roryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rogelioticotaco"}

	roryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рорикапитально"}

	rorySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亚瑟哪哈"}

	rorySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rogelioticotaco"}

	roryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亞瑟哪哈"}
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
		Animal:   Lion,
		Birthday: roryBirthday,
		Code:     roryCode,
		Gender:   nook.Male,
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
		Personality: nook.Jock,
		Phrase:      roryPhrase}
)
