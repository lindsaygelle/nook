package cat

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
		Value:    "ピエール"}

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
		Animal:   animal.Cat,
		Birthday: pierreBirthday,
		Code:     pierreCode,
		Key:      character.Pierre,
		Gender:   gender.Male,
		Name:     pierreName}
)

var (
	pierreAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ってね"}

	pierreCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	pierreDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	pierreFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	pierreGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	pierreItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	pierreJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ってね"}

	pierreKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	pierreLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	pierreRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	pierreSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	pierreSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	pierreTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	pierrePhrase = nook.Languages{
		language.AmericanEnglish:      pierreAmericanEnglishPhrase,
		language.CanadianFrench:       pierreCanadianFrenchPhrase,
		language.Dutch:                pierreDutchPhrase,
		language.French:               pierreFrenchPhrase,
		language.German:               pierreGermanPhrase,
		language.Italian:              pierreItalianPhrase,
		language.Japanese:             pierreJapanesePhrase,
		language.Korean:               pierreKoreanPhrase,
		language.LatinAmericanSpanish: pierreLatinAmericanSpanishPhrase,
		language.Russian:              pierreRussianPhrase,
		language.SimplifiedChinese:    pierreSimplifiedChinesePhrase,
		language.Spanish:              pierreSpanishPhrase,
		language.TraditionalChinese:   pierreTraditionalChinesePhrase}
)

var (
	Pierre = nook.Villager{
		Character:   pierreCharacter,
		Personality: personality.Jock,
		Phrase:      pierrePhrase}
)
