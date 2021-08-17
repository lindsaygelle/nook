package wolf

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
	tarouBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	tarouCode = nook.Code{
		Value: ""}
)

var (
	tarouAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tarou"}

	tarouCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	tarouDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	tarouFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	tarouGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	tarouItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	tarouJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タロウ"}

	tarouKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	tarouLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	tarouRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	tarouSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	tarouSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	tarouTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	tarouName = nook.Languages{
		language.AmericanEnglish:      tarouAmericanEnglishName,
		language.CanadianFrench:       tarouCanadianFrenchName,
		language.Dutch:                tarouDutchName,
		language.French:               tarouFrenchName,
		language.German:               tarouGermanName,
		language.Italian:              tarouItalianName,
		language.Japanese:             tarouJapaneseName,
		language.Korean:               tarouKoreanName,
		language.LatinAmericanSpanish: tarouLatinAmericanSpanishName,
		language.Russian:              tarouRussianName,
		language.SimplifiedChinese:    tarouSimplifiedChineseName,
		language.Spanish:              tarouSpanishName,
		language.TraditionalChinese:   tarouTraditionalChineseName}
)

var (
	tarouCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: tarouBirthday,
		Code:     tarouCode,
		Key:      character.Tarou,
		Gender:   gender.Male,
		Name:     tarouName}
)

var (
	tarouAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ワオーン"}

	tarouCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	tarouDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	tarouFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	tarouGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	tarouItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	tarouJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワオーン"}

	tarouKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	tarouLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	tarouRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	tarouSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	tarouSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	tarouTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	tarouPhrase = nook.Languages{
		language.AmericanEnglish:      tarouAmericanEnglishPhrase,
		language.CanadianFrench:       tarouCanadianFrenchPhrase,
		language.Dutch:                tarouDutchPhrase,
		language.French:               tarouFrenchPhrase,
		language.German:               tarouGermanPhrase,
		language.Italian:              tarouItalianPhrase,
		language.Japanese:             tarouJapanesePhrase,
		language.Korean:               tarouKoreanPhrase,
		language.LatinAmericanSpanish: tarouLatinAmericanSpanishPhrase,
		language.Russian:              tarouRussianPhrase,
		language.SimplifiedChinese:    tarouSimplifiedChinesePhrase,
		language.Spanish:              tarouSpanishPhrase,
		language.TraditionalChinese:   tarouTraditionalChinesePhrase}
)

var (
	Tarou = nook.Villager{
		Character:   tarouCharacter,
		Personality: personality.Jock,
		Phrase:      tarouPhrase}
)
