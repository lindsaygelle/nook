package rhinoceros

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
		Value:    "Ben"}

	tankDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tank"}

	tankFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ben"}

	tankGermanName = nook.Name{
		Language: language.German,
		Value:    "Frank"}

	tankItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rinaldo"}

	tankJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "くるぶし"}

	tankKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "탱크"}

	tankLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aníbal"}

	tankRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Танк"}

	tankSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿足"}

	tankSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aníbal"}

	tankTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿足"}
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
		Animal:   animal.Rhinoceros,
		Birthday: tankBirthday,
		Code:     tankCode,
		Key:      character.Tank,
		Gender:   gender.Male,
		Name:     tankName,
		Special:  false}
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
		language.AmericanEnglish:      tankAmericanEnglishPhrase,
		language.CanadianFrench:       tankCanadianFrenchPhrase,
		language.Dutch:                tankDutchPhrase,
		language.French:               tankFrenchPhrase,
		language.German:               tankGermanPhrase,
		language.Italian:              tankItalianPhrase,
		language.Japanese:             tankJapanesePhrase,
		language.Korean:               tankKoreanPhrase,
		language.LatinAmericanSpanish: tankLatinAmericanSpanishPhrase,
		language.Russian:              tankRussianPhrase,
		language.SimplifiedChinese:    tankSimplifiedChinesePhrase,
		language.Spanish:              tankSpanishPhrase,
		language.TraditionalChinese:   tankTraditionalChinesePhrase}
)

var (
	Tank = nook.Villager{
		Character:   tankCharacter,
		Personality: personality.Jock,
		Phrase:      tankPhrase}
)
