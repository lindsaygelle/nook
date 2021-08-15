package hamster

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Chicoyes sir"}

	rodneyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rodneyle ham"}

	rodneyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chicocaries"}

	rodneyGermanName = nook.Name{
		Language: language.German,
		Value:    "Reinholdmjamjam"}

	rodneyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosikkioflop"}

	rodneyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジミーヒュー"}

	rodneyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "지미휴우"}

	rodneyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chusquisadoya"}

	rodneyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Родниле хом"}

	rodneySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "米米咻咻"}

	rodneySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chusquisadoya"}

	rodneyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "米米咻咻"}
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
		Animal:   Hamster,
		Birthday: rodneyBirthday,
		Code:     rodneyCode,
		Gender:   nook.Male,
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
	Rodney = nook.Villager{
		Character:   rodneyCharacter,
		Personality: nook.Smug,
		Phrase:      rodneyPhrase}
)
