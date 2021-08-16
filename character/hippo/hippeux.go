package hippo

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
		Value:    "Paulito"}

	hippeuxDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hippeux"}

	hippeuxFrenchName = nook.Name{
		Language: language.French,
		Value:    "Paulito"}

	hippeuxGermanName = nook.Name{
		Language: language.German,
		Value:    "Herbert"}

	hippeuxItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Poppy"}

	hippeuxJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ディビッド"}

	hippeuxKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "데이빗"}

	hippeuxLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hipóleo"}

	hippeuxRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гиппо"}

	hippeuxSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "戴维"}

	hippeuxSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hipóleo"}

	hippeuxTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "戴維"}
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
		Animal:   animal.Hippo,
		Birthday: hippeuxBirthday,
		Code:     hippeuxCode,
		Key:      character.Hippeux,
		Gender:   gender.Male,
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
		Personality: personality.Smug,
		Phrase:      hippeuxPhrase}
)
