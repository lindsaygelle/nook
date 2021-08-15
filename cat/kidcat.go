package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	kidcatBirthday = nook.Birthday{
		Day:   1,
		Month: time.August}
)

var (
	kidcatCode = nook.Code{
		Value: "cat10"}
)

var (
	kidcatAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kid Cat"}

	kidcatCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Câlinpsst"}

	kidcatDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kid Catpsst"}

	kidcatFrenchName = nook.Name{
		Language: language.French,
		Value:    "Câlinpsst"}

	kidcatGermanName = nook.Name{
		Language: language.German,
		Value:    "Pitmaunz"}

	kidcatItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zampapsst"}

	kidcatJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "１ごうとぉっ"}

	kidcatKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "1호얍"}

	kidcatLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gatománpsst"}

	kidcatRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кид Кэтшш"}

	kidcatSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿一喝"}

	kidcatSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gatománsoldado"}

	kidcatTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿一喝"}
)

var (
	kidcatName = nook.Languages{
		language.AmericanEnglish:      kidcatAmericanEnglishName,
		language.CanadianFrench:       kidcatCanadianFrenchName,
		language.Dutch:                kidcatDutchName,
		language.French:               kidcatFrenchName,
		language.German:               kidcatGermanName,
		language.Italian:              kidcatItalianName,
		language.Japanese:             kidcatJapaneseName,
		language.Korean:               kidcatKoreanName,
		language.LatinAmericanSpanish: kidcatLatinAmericanSpanishName,
		language.Russian:              kidcatRussianName,
		language.SimplifiedChinese:    kidcatSimplifiedChineseName,
		language.Spanish:              kidcatSpanishName,
		language.TraditionalChinese:   kidcatTraditionalChineseName}
)

var (
	kidcatCharacter = nook.Character{
		Animal:   Cat,
		Birthday: kidcatBirthday,
		Code:     kidcatCode,
		Gender:   nook.Male,
		Name:     kidcatName}
)

var (
	kidcatAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "psst"}

	kidcatCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "psst"}

	kidcatDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "psst"}

	kidcatFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "psst"}

	kidcatGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "maunz"}

	kidcatItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "psst"}

	kidcatJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とぉっ"}

	kidcatKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "얍"}

	kidcatLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "psst"}

	kidcatRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шш"}

	kidcatSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喝"}

	kidcatSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "soldado"}

	kidcatTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喝"}
)

var (
	kidcatPhrase = nook.Languages{
		language.AmericanEnglish:      kidcatAmericanEnglishName,
		language.CanadianFrench:       kidcatCanadianFrenchName,
		language.Dutch:                kidcatDutchName,
		language.French:               kidcatFrenchName,
		language.German:               kidcatGermanName,
		language.Italian:              kidcatItalianName,
		language.Japanese:             kidcatJapaneseName,
		language.Korean:               kidcatKoreanName,
		language.LatinAmericanSpanish: kidcatLatinAmericanSpanishName,
		language.Russian:              kidcatRussianName,
		language.SimplifiedChinese:    kidcatSimplifiedChineseName,
		language.Spanish:              kidcatSpanishName,
		language.TraditionalChinese:   kidcatTraditionalChineseName}
)

var (
	KidCat = nook.Villager{
		Character:   kidcatCharacter,
		Personality: nook.Jock,
		Phrase:      kidcatPhrase}
)
