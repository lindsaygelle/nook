package hamster

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
	rodneyBirthday = nook.Birthday{
		Day:   10,
		Month: time.November}
)

var (
	rodneyCode = nook.Code{
		Value: "ham03"}
)

var (
	rodneyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rodney"}

	rodneyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chico"}

	rodneyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rodney"}

	rodneyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chico"}

	rodneyGermanName = nook.Name{
		Language: language.German,
		Value:    "Reinhold"}

	rodneyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosikkio"}

	rodneyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジミー"}

	rodneyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "지미"}

	rodneyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chusquis"}

	rodneyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Родни"}

	rodneySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "米米"}

	rodneySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chusquis"}

	rodneyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "米米"}
)

var (
	rodneyName = nook.Languages{
		language.AmericanEnglish:      rodneyAmericanEnglishName,
		language.CanadianFrench:       rodneyCanadianFrenchName,
		language.Dutch:                rodneyDutchName,
		language.French:               rodneyFrenchName,
		language.German:               rodneyGermanName,
		language.Italian:              rodneyItalianName,
		language.Japanese:             rodneyJapaneseName,
		language.Korean:               rodneyKoreanName,
		language.LatinAmericanSpanish: rodneyLatinAmericanSpanishName,
		language.Russian:              rodneyRussianName,
		language.SimplifiedChinese:    rodneySimplifiedChineseName,
		language.Spanish:              rodneySpanishName,
		language.TraditionalChinese:   rodneyTraditionalChineseName}
)

var (
	rodneyCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: rodneyBirthday,
		Code:     rodneyCode,
		Key:      character.Rodney,
		Gender:   gender.Male,
		Name:     rodneyName}
)

var (
	rodneyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "le ham"}

	rodneyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "yes sir"}

	rodneyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "le ham"}

	rodneyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "caries"}

	rodneyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mjamjam"}

	rodneyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "flop"}

	rodneyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヒュー"}

	rodneyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "휴우"}

	rodneyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "adoya"}

	rodneyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ле хом"}

	rodneySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咻咻"}

	rodneySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "adoya"}

	rodneyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咻咻"}
)

var (
	rodneyPhrase = nook.Languages{
		language.AmericanEnglish:      rodneyAmericanEnglishPhrase,
		language.CanadianFrench:       rodneyCanadianFrenchPhrase,
		language.Dutch:                rodneyDutchPhrase,
		language.French:               rodneyFrenchPhrase,
		language.German:               rodneyGermanPhrase,
		language.Italian:              rodneyItalianPhrase,
		language.Japanese:             rodneyJapanesePhrase,
		language.Korean:               rodneyKoreanPhrase,
		language.LatinAmericanSpanish: rodneyLatinAmericanSpanishPhrase,
		language.Russian:              rodneyRussianPhrase,
		language.SimplifiedChinese:    rodneySimplifiedChinesePhrase,
		language.Spanish:              rodneySpanishPhrase,
		language.TraditionalChinese:   rodneyTraditionalChinesePhrase}
)

var (
	Rodney = nook.Villager{
		Character:   rodneyCharacter,
		Personality: personality.Smug,
		Phrase:      rodneyPhrase}
)
