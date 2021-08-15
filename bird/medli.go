package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	medliBirthday = nook.Birthday{
		Day:   13,
		Month: time.December}
)

var (
	medliCode = nook.Code{
		Value: "brd19"}
)

var (
	medliAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Medli"}

	medliCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Médolielalalyre"}

	medliDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	medliFrenchName = nook.Name{
		Language: language.French,
		Value:    "Médolielalalyre"}

	medliGermanName = nook.Name{
		Language: language.German,
		Value:    "Medoliehilfhilf"}

	medliItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Famirèplinplon"}

	medliJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メドリくらさい"}

	medliKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "메들리핑그르르"}

	medliLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Medliplimplín"}

	medliRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	medliSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	medliSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Medliplimplín"}

	medliTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	medliName = nook.Languages{
		language.AmericanEnglish:      medliAmericanEnglishName,
		language.CanadianFrench:       medliCanadianFrenchName,
		language.Dutch:                medliDutchName,
		language.French:               medliFrenchName,
		language.German:               medliGermanName,
		language.Italian:              medliItalianName,
		language.Japanese:             medliJapaneseName,
		language.Korean:               medliKoreanName,
		language.LatinAmericanSpanish: medliLatinAmericanSpanishName,
		language.Russian:              medliRussianName,
		language.SimplifiedChinese:    medliSimplifiedChineseName,
		language.Spanish:              medliSpanishName,
		language.TraditionalChinese:   medliTraditionalChineseName}
)

var (
	medliCharacter = nook.Character{
		Animal:   Bird,
		Birthday: medliBirthday,
		Code:     medliCode,
		Gender:   nook.Female,
		Name:     medliName}
)

var (
	medliAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "gimme"}

	medliCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "lalalyre"}

	medliDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	medliFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "lalalyre"}

	medliGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hilfhilf"}

	medliItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "plinplon"}

	medliJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "くらさい"}

	medliKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "핑그르르"}

	medliLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "plimplín"}

	medliRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	medliSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	medliSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "plimplín"}

	medliTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	medliPhrase = nook.Languages{
		language.AmericanEnglish:      medliAmericanEnglishName,
		language.CanadianFrench:       medliCanadianFrenchName,
		language.Dutch:                medliDutchName,
		language.French:               medliFrenchName,
		language.German:               medliGermanName,
		language.Italian:              medliItalianName,
		language.Japanese:             medliJapaneseName,
		language.Korean:               medliKoreanName,
		language.LatinAmericanSpanish: medliLatinAmericanSpanishName,
		language.Russian:              medliRussianName,
		language.SimplifiedChinese:    medliSimplifiedChineseName,
		language.Spanish:              medliSpanishName,
		language.TraditionalChinese:   medliTraditionalChineseName}
)

var (
	Medli = nook.Villager{
		Character:   medliCharacter,
		Personality: nook.Normal,
		Phrase:      medliPhrase}
)
