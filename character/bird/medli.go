package bird

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
		Value:    "Médolie"}

	medliDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	medliFrenchName = nook.Name{
		Language: language.French,
		Value:    "Médolie"}

	medliGermanName = nook.Name{
		Language: language.German,
		Value:    "Medolie"}

	medliItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Famirè"}

	medliJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メドリ"}

	medliKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "메들리"}

	medliLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Medli"}

	medliRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	medliSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	medliSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Medli"}

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
		Animal:   animal.Bird,
		Birthday: medliBirthday,
		Code:     medliCode,
		Key:      character.Medli,
		Gender:   gender.Female,
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
		Value:    ""}

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
		Value:    ""}

	medliSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	medliSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "plimplín"}

	medliTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	medliPhrase = nook.Languages{
		language.AmericanEnglish:      medliAmericanEnglishPhrase,
		language.CanadianFrench:       medliCanadianFrenchPhrase,
		language.Dutch:                medliDutchPhrase,
		language.French:               medliFrenchPhrase,
		language.German:               medliGermanPhrase,
		language.Italian:              medliItalianPhrase,
		language.Japanese:             medliJapanesePhrase,
		language.Korean:               medliKoreanPhrase,
		language.LatinAmericanSpanish: medliLatinAmericanSpanishPhrase,
		language.Russian:              medliRussianPhrase,
		language.SimplifiedChinese:    medliSimplifiedChinesePhrase,
		language.Spanish:              medliSpanishPhrase,
		language.TraditionalChinese:   medliTraditionalChinesePhrase}
)

var (
	Medli = nook.Villager{
		Character:   medliCharacter,
		Personality: personality.Normal,
		Phrase:      medliPhrase}
)
