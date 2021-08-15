package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	butchBirthday = nook.Birthday{
		Day:   1,
		Month: time.November}
)

var (
	butchCode = nook.Code{
		Value: "dog01"}
)

var (
	butchAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Butch"}

	butchCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "AvrilWROOOOUF"}

	butchDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "ButchBROEEEF"}

	butchFrenchName = nook.Name{
		Language: language.French,
		Value:    "AvrilWROOOOUF"}

	butchGermanName = nook.Name{
		Language: language.German,
		Value:    "HassoGRRUFF"}

	butchItalianName = nook.Name{
		Language: language.Italian,
		Value:    "AttilaROOOOOF"}

	butchJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジョンノン"}

	butchKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "존아니"}

	butchLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brunogruf-gruf"}

	butchRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бутчррр-гав"}

	butchSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "约翰侬"}

	butchSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brunogruf-gruf"}

	butchTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "約翰儂"}
)

var (
	butchName = nook.Languages{
		language.AmericanEnglish:      butchAmericanEnglishName,
		language.CanadianFrench:       butchCanadianFrenchName,
		language.Dutch:                butchDutchName,
		language.French:               butchFrenchName,
		language.German:               butchGermanName,
		language.Italian:              butchItalianName,
		language.Japanese:             butchJapaneseName,
		language.Korean:               butchKoreanName,
		language.LatinAmericanSpanish: butchLatinAmericanSpanishName,
		language.Russian:              butchRussianName,
		language.SimplifiedChinese:    butchSimplifiedChineseName,
		language.Spanish:              butchSpanishName,
		language.TraditionalChinese:   butchTraditionalChineseName}
)

var (
	butchCharacter = nook.Character{
		Animal:   Dog,
		Birthday: butchBirthday,
		Code:     butchCode,
		Gender:   nook.Male,
		Name:     butchName}
)

var (
	butchAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ROOOOOWF"}

	butchCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "WROOOOUF"}

	butchDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "BROEEEF"}

	butchFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "WROOOOUF"}

	butchGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "GRRUFF"}

	butchItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ROOOOOF"}

	butchJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ノン"}

	butchKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아니"}

	butchLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gruf-gruf"}

	butchRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ррр-гав"}

	butchSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "侬"}

	butchSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gruf-gruf"}

	butchTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "儂"}
)

var (
	butchPhrase = nook.Languages{
		language.AmericanEnglish:      butchAmericanEnglishName,
		language.CanadianFrench:       butchCanadianFrenchName,
		language.Dutch:                butchDutchName,
		language.French:               butchFrenchName,
		language.German:               butchGermanName,
		language.Italian:              butchItalianName,
		language.Japanese:             butchJapaneseName,
		language.Korean:               butchKoreanName,
		language.LatinAmericanSpanish: butchLatinAmericanSpanishName,
		language.Russian:              butchRussianName,
		language.SimplifiedChinese:    butchSimplifiedChineseName,
		language.Spanish:              butchSpanishName,
		language.TraditionalChinese:   butchTraditionalChineseName}
)

var (
	Butch = nook.Villager{
		Character:   butchCharacter,
		Personality: nook.Cranky,
		Phrase:      butchPhrase}
)
