package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	tarouDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	tarouFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	tarouGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	tarouItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	tarouJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タロウ"}

	tarouKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	tarouLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	tarouRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	tarouSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	tarouSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	tarouTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Wolf,
		Birthday: tarouBirthday,
		Code:     tarouCode,
		Gender:   nook.Male,
		Name:     tarouName}
)

var (
	tarouAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	tarouCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	tarouDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	tarouFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	tarouGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	tarouItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	tarouJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワオーン"}

	tarouKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	tarouLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	tarouRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	tarouSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	tarouSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	tarouTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	tarouPhrase = nook.Languages{
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
	Tarou = nook.Villager{
		Character:   tarouCharacter,
		Personality: nook.Jock,
		Phrase:      tarouPhrase}
)
