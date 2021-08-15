package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	hippeuxBirthday = nook.Birthday{
		Day:   15,
		Month: time.October}
)

var (
	hippeuxCode = nook.Code{
		Value: "hip09"}
)

var (
	hippeuxAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hippeux"}

	hippeuxCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Paulitotam-tam"}

	hippeuxDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hippeuxtuurlijk"}

	hippeuxFrenchName = nook.Name{
		Language: language.French,
		Value:    "Paulitotam-tam"}

	hippeuxGermanName = nook.Name{
		Language: language.German,
		Value:    "Herberthipphipp"}

	hippeuxItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Poppyvale"}

	hippeuxJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ディビッドイェス"}

	hippeuxKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "데이빗예스"}

	hippeuxLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hipóleochofchof"}

	hippeuxRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гиппоестессно"}

	hippeuxSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "戴维Yes"}

	hippeuxSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hipóleochofchof"}

	hippeuxTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "戴維Yes"}
)

var (
	hippeuxName = nook.Languages{
		language.AmericanEnglish:      hippeuxAmericanEnglishName,
		language.CanadianFrench:       hippeuxCanadianFrenchName,
		language.Dutch:                hippeuxDutchName,
		language.French:               hippeuxFrenchName,
		language.German:               hippeuxGermanName,
		language.Italian:              hippeuxItalianName,
		language.Japanese:             hippeuxJapaneseName,
		language.Korean:               hippeuxKoreanName,
		language.LatinAmericanSpanish: hippeuxLatinAmericanSpanishName,
		language.Russian:              hippeuxRussianName,
		language.SimplifiedChinese:    hippeuxSimplifiedChineseName,
		language.Spanish:              hippeuxSpanishName,
		language.TraditionalChinese:   hippeuxTraditionalChineseName}
)

var (
	hippeuxCharacter = nook.Character{
		Animal:   Hippo,
		Birthday: hippeuxBirthday,
		Code:     hippeuxCode,
		Gender:   nook.Male,
		Name:     hippeuxName}
)

var (
	hippeuxAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "natch"}

	hippeuxCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tam-tam"}

	hippeuxDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "tuurlijk"}

	hippeuxFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tam-tam"}

	hippeuxGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hipphipp"}

	hippeuxItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "vale"}

	hippeuxJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "イェス"}

	hippeuxKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "예스"}

	hippeuxLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chofchof"}

	hippeuxRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "естессно"}

	hippeuxSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Yes"}

	hippeuxSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chofchof"}

	hippeuxTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Yes"}
)

var (
	hippeuxPhrase = nook.Languages{
		language.AmericanEnglish:      hippeuxAmericanEnglishName,
		language.CanadianFrench:       hippeuxCanadianFrenchName,
		language.Dutch:                hippeuxDutchName,
		language.French:               hippeuxFrenchName,
		language.German:               hippeuxGermanName,
		language.Italian:              hippeuxItalianName,
		language.Japanese:             hippeuxJapaneseName,
		language.Korean:               hippeuxKoreanName,
		language.LatinAmericanSpanish: hippeuxLatinAmericanSpanishName,
		language.Russian:              hippeuxRussianName,
		language.SimplifiedChinese:    hippeuxSimplifiedChineseName,
		language.Spanish:              hippeuxSpanishName,
		language.TraditionalChinese:   hippeuxTraditionalChineseName}
)

var (
	Hippeux = nook.Villager{
		Character:   hippeuxCharacter,
		Personality: nook.Smug,
		Phrase:      hippeuxPhrase}
)
