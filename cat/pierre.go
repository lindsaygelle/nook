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
		Value:    "N/A"}

	pierreDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	pierreFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	pierreGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	pierreItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	pierreJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピエール"}

	pierreKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	pierreLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	pierreRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	pierreSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	pierreSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	pierreTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
