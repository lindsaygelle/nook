package rhinoceros

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	tankBirthday = nook.Birthday{
		Day:   6,
		Month: time.May}
)

var (
	tankCode = nook.Code{
		Value: "rhn00"}
)

var (
	tankAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tank"}

	tankCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Benkaboum"}

	tankDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "TankklaBAM"}

	tankFrenchName = nook.Name{
		Language: language.French,
		Value:    "Benkaboum"}

	tankGermanName = nook.Name{
		Language: language.German,
		Value:    "Frankrumpel"}

	tankItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rinaldozamperù"}

	tankJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "くるぶしですサイ"}

	tankKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "탱크아뿔소"}

	tankLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aníbaltapumba"}

	tankRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Танкбабах"}

	tankSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿足犀犀"}

	tankSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aníbaltapumba"}

	tankTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿足犀犀"}
)

var (
	tankName = nook.Languages{
		language.AmericanEnglish:      tankAmericanEnglishName,
		language.CanadianFrench:       tankCanadianFrenchName,
		language.Dutch:                tankDutchName,
		language.French:               tankFrenchName,
		language.German:               tankGermanName,
		language.Italian:              tankItalianName,
		language.Japanese:             tankJapaneseName,
		language.Korean:               tankKoreanName,
		language.LatinAmericanSpanish: tankLatinAmericanSpanishName,
		language.Russian:              tankRussianName,
		language.SimplifiedChinese:    tankSimplifiedChineseName,
		language.Spanish:              tankSpanishName,
		language.TraditionalChinese:   tankTraditionalChineseName}
)

var (
	tankCharacter = nook.Character{
		Animal:   Rhinoceros,
		Birthday: tankBirthday,
		Code:     tankCode,
		Gender:   nook.Male,
		Name:     tankName}
)

var (
	tankAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "kerPOW"}

	tankCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kaboum"}

	tankDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "klaBAM"}

	tankFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "kaboum"}

	tankGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "rumpel"}

	tankItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zamperù"}

	tankJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですサイ"}

	tankKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아뿔소"}

	tankLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tapumba"}

	tankRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бабах"}

	tankSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "犀犀"}

	tankSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tapumba"}

	tankTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "犀犀"}
)

var (
	tankPhrase = nook.Languages{
		language.AmericanEnglish:      tankAmericanEnglishName,
		language.CanadianFrench:       tankCanadianFrenchName,
		language.Dutch:                tankDutchName,
		language.French:               tankFrenchName,
		language.German:               tankGermanName,
		language.Italian:              tankItalianName,
		language.Japanese:             tankJapaneseName,
		language.Korean:               tankKoreanName,
		language.LatinAmericanSpanish: tankLatinAmericanSpanishName,
		language.Russian:              tankRussianName,
		language.SimplifiedChinese:    tankSimplifiedChineseName,
		language.Spanish:              tankSpanishName,
		language.TraditionalChinese:   tankTraditionalChineseName}
)

var (
	Tank = nook.Villager{
		Character:   tankCharacter,
		Personality: nook.Jock,
		Phrase:      tankPhrase}
)
