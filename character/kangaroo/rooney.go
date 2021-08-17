package kangaroo

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
		Value:    "Mike"}

	rooneyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rooney"}

	rooneyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mike"}

	rooneyGermanName = nook.Name{
		Language: language.German,
		Value:    "Robert"}

	rooneyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Balzak"}

	rooneyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マイク"}

	rooneyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마이크"}

	rooneyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cerillo"}

	rooneyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Руни"}

	rooneySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麦克"}

	rooneySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cerillo"}

	rooneyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麥克"}
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
		Animal:   animal.Kangaroo,
		Birthday: rooneyBirthday,
		Code:     rooneyCode,
		Key:      character.Rooney,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      rooneyAmericanEnglishPhrase,
		language.CanadianFrench:       rooneyCanadianFrenchPhrase,
		language.Dutch:                rooneyDutchPhrase,
		language.French:               rooneyFrenchPhrase,
		language.German:               rooneyGermanPhrase,
		language.Italian:              rooneyItalianPhrase,
		language.Japanese:             rooneyJapanesePhrase,
		language.Korean:               rooneyKoreanPhrase,
		language.LatinAmericanSpanish: rooneyLatinAmericanSpanishPhrase,
		language.Russian:              rooneyRussianPhrase,
		language.SimplifiedChinese:    rooneySimplifiedChinesePhrase,
		language.Spanish:              rooneySpanishPhrase,
		language.TraditionalChinese:   rooneyTraditionalChinesePhrase}
)

var (
	Rooney = nook.Villager{
		Character:   rooneyCharacter,
		Personality: personality.Cranky,
		Phrase:      rooneyPhrase}
)
