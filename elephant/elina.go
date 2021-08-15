package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	elinaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	elinaCode = nook.Code{
		Value: ""}
)

var (
	elinaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Elina"}

	elinaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	elinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	elinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Fanfangarri"}

	elinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Elenazwerg"}

	elinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Elliepesciolino"}

	elinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビンディナマステ"}

	elinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	elinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	elinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	elinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	elinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Elinacamarón"}

	elinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	elinaName = nook.Languages{
		language.AmericanEnglish:      elinaAmericanEnglishName,
		language.CanadianFrench:       elinaCanadianFrenchName,
		language.Dutch:                elinaDutchName,
		language.French:               elinaFrenchName,
		language.German:               elinaGermanName,
		language.Italian:              elinaItalianName,
		language.Japanese:             elinaJapaneseName,
		language.Korean:               elinaKoreanName,
		language.LatinAmericanSpanish: elinaLatinAmericanSpanishName,
		language.Russian:              elinaRussianName,
		language.SimplifiedChinese:    elinaSimplifiedChineseName,
		language.Spanish:              elinaSpanishName,
		language.TraditionalChinese:   elinaTraditionalChineseName}
)

var (
	elinaCharacter = nook.Character{
		Animal:   Elephant,
		Birthday: elinaBirthday,
		Code:     elinaCode,
		Gender:   nook.Female,
		Name:     elinaName}
)

var (
	elinaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	elinaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	elinaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	elinaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "garri"}

	elinaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zwerg"}

	elinaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pesciolino"}

	elinaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ナマステ"}

	elinaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	elinaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	elinaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	elinaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	elinaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "camarón"}

	elinaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	elinaPhrase = nook.Languages{
		language.AmericanEnglish:      elinaAmericanEnglishName,
		language.CanadianFrench:       elinaCanadianFrenchName,
		language.Dutch:                elinaDutchName,
		language.French:               elinaFrenchName,
		language.German:               elinaGermanName,
		language.Italian:              elinaItalianName,
		language.Japanese:             elinaJapaneseName,
		language.Korean:               elinaKoreanName,
		language.LatinAmericanSpanish: elinaLatinAmericanSpanishName,
		language.Russian:              elinaRussianName,
		language.SimplifiedChinese:    elinaSimplifiedChineseName,
		language.Spanish:              elinaSpanishName,
		language.TraditionalChinese:   elinaTraditionalChineseName}
)

var (
	Elina = nook.Villager{
		Character:   elinaCharacter,
		Personality: nook.Peppy,
		Phrase:      elinaPhrase}
)
