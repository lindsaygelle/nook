package kangaroo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rooneyBirthday = nook.Birthday{
		Day:   1,
		Month: time.December}
)

var (
	rooneyCode = nook.Code{
		Value: "kgr09"}
)

var (
	rooneyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rooney"}

	rooneyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mikesavate"}

	rooneyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rooneyboks"}

	rooneyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mikesavate"}

	rooneyGermanName = nook.Name{
		Language: language.German,
		Value:    "Robertbonk"}

	rooneyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Balzakbalz balz"}

	rooneyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マイクやるぜ"}

	rooneyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마이크아자"}

	rooneyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cerillochimpón"}

	rooneyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рунихук"}

	rooneySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麦克做就对了"}

	rooneySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cerillochimpón"}

	rooneyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麥克做就對了"}
)

var (
	rooneyName = nook.Languages{
		language.AmericanEnglish:      rooneyAmericanEnglishName,
		language.CanadianFrench:       rooneyCanadianFrenchName,
		language.Dutch:                rooneyDutchName,
		language.French:               rooneyFrenchName,
		language.German:               rooneyGermanName,
		language.Italian:              rooneyItalianName,
		language.Japanese:             rooneyJapaneseName,
		language.Korean:               rooneyKoreanName,
		language.LatinAmericanSpanish: rooneyLatinAmericanSpanishName,
		language.Russian:              rooneyRussianName,
		language.SimplifiedChinese:    rooneySimplifiedChineseName,
		language.Spanish:              rooneySpanishName,
		language.TraditionalChinese:   rooneyTraditionalChineseName}
)

var (
	rooneyCharacter = nook.Character{
		Animal:   Kangaroo,
		Birthday: rooneyBirthday,
		Code:     rooneyCode,
		Gender:   nook.Male,
		Name:     rooneyName}
)

var (
	rooneyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "punches"}

	rooneyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "savate"}

	rooneyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boks"}

	rooneyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "savate"}

	rooneyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bonk"}

	rooneyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "balz balz"}

	rooneyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やるぜ"}

	rooneyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아자"}

	rooneyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chimpón"}

	rooneyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хук"}

	rooneySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "做就对了"}

	rooneySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chimpón"}

	rooneyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "做就對了"}
)

var (
	rooneyPhrase = nook.Languages{
		language.AmericanEnglish:      rooneyAmericanEnglishName,
		language.CanadianFrench:       rooneyCanadianFrenchName,
		language.Dutch:                rooneyDutchName,
		language.French:               rooneyFrenchName,
		language.German:               rooneyGermanName,
		language.Italian:              rooneyItalianName,
		language.Japanese:             rooneyJapaneseName,
		language.Korean:               rooneyKoreanName,
		language.LatinAmericanSpanish: rooneyLatinAmericanSpanishName,
		language.Russian:              rooneyRussianName,
		language.SimplifiedChinese:    rooneySimplifiedChineseName,
		language.Spanish:              rooneySpanishName,
		language.TraditionalChinese:   rooneyTraditionalChineseName}
)

var (
	Rooney = nook.Villager{
		Character:   rooneyCharacter,
		Personality: nook.Cranky,
		Phrase:      rooneyPhrase}
)
