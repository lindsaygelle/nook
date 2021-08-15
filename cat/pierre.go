package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pierreBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	pierreCode = nook.Code{
		Value: ""}
)

var (
	pierreAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pierre"}

	pierreCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	pierreDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	pierreFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	pierreGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	pierreItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	pierreJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピエールってね"}

	pierreKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	pierreLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	pierreRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	pierreSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	pierreSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	pierreTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	pierreName = nook.Languages{
		language.AmericanEnglish:      pierreAmericanEnglishName,
		language.CanadianFrench:       pierreCanadianFrenchName,
		language.Dutch:                pierreDutchName,
		language.French:               pierreFrenchName,
		language.German:               pierreGermanName,
		language.Italian:              pierreItalianName,
		language.Japanese:             pierreJapaneseName,
		language.Korean:               pierreKoreanName,
		language.LatinAmericanSpanish: pierreLatinAmericanSpanishName,
		language.Russian:              pierreRussianName,
		language.SimplifiedChinese:    pierreSimplifiedChineseName,
		language.Spanish:              pierreSpanishName,
		language.TraditionalChinese:   pierreTraditionalChineseName}
)

var (
	pierreCharacter = nook.Character{
		Animal:   Cat,
		Birthday: pierreBirthday,
		Code:     pierreCode,
		Gender:   nook.Male,
		Name:     pierreName}
)

var (
	pierreAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	pierreCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	pierreDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	pierreFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	pierreGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	pierreItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	pierreJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ってね"}

	pierreKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	pierreLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	pierreRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	pierreSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	pierreSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	pierreTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	pierrePhrase = nook.Languages{
		language.AmericanEnglish:      pierreAmericanEnglishName,
		language.CanadianFrench:       pierreCanadianFrenchName,
		language.Dutch:                pierreDutchName,
		language.French:               pierreFrenchName,
		language.German:               pierreGermanName,
		language.Italian:              pierreItalianName,
		language.Japanese:             pierreJapaneseName,
		language.Korean:               pierreKoreanName,
		language.LatinAmericanSpanish: pierreLatinAmericanSpanishName,
		language.Russian:              pierreRussianName,
		language.SimplifiedChinese:    pierreSimplifiedChineseName,
		language.Spanish:              pierreSpanishName,
		language.TraditionalChinese:   pierreTraditionalChineseName}
)

var (
	Pierre = nook.Villager{
		Character:   pierreCharacter,
		Personality: nook.Jock,
		Phrase:      pierrePhrase}
)
