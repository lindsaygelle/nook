package alligator

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
	alfonsoBirthday = nook.Birthday{
		Day:   9,
		Month: time.June}
)

var (
	alfonsoCode = nook.Code{
		Value: "crd00"}
)

var (
	alfonsoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Alfonso"}

	alfonsoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Alphonse"}

	alfonsoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Alfonso"}

	alfonsoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Alphonse"}

	alfonsoGermanName = nook.Name{
		Language: language.German,
		Value:    "Markus"}

	alfonsoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Alfonso"}

	alfonsoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アルベルト"}

	alfonsoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "알베르트"}

	alfonsoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kaimán"}

	alfonsoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Альфонсо"}

	alfonsoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿泥"}

	alfonsoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kaimán"}

	alfonsoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿泥"}
)

var (
	alfonsoName = nook.Languages{
		language.AmericanEnglish:      alfonsoAmericanEnglishName,
		language.CanadianFrench:       alfonsoCanadianFrenchName,
		language.Dutch:                alfonsoDutchName,
		language.French:               alfonsoFrenchName,
		language.German:               alfonsoGermanName,
		language.Italian:              alfonsoItalianName,
		language.Japanese:             alfonsoJapaneseName,
		language.Korean:               alfonsoKoreanName,
		language.LatinAmericanSpanish: alfonsoLatinAmericanSpanishName,
		language.Russian:              alfonsoRussianName,
		language.SimplifiedChinese:    alfonsoSimplifiedChineseName,
		language.Spanish:              alfonsoSpanishName,
		language.TraditionalChinese:   alfonsoTraditionalChineseName}
)

var (
	alfonsoCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: alfonsoBirthday,
		Code:     alfonsoCode,
		Key:      character.Alfonso,
		Gender:   gender.Male,
		Name:     alfonsoName}
)

var (
	alfonsoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "it's a me"}

	alfonsoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "moi moi"}

	alfonsoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "it⁠'⁠sa me"}

	alfonsoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "moi moi"}

	alfonsoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ahhhrg"}

	alfonsoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "proprio"}

	alfonsoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だワニ"}

	alfonsoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "나아거"}

	alfonsoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñam-ñam"}

	alfonsoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "это я"}

	alfonsoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鳄泥"}

	alfonsoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñam-ñam"}

	alfonsoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鱷泥"}
)

var (
	alfonsoPhrase = nook.Languages{
		language.AmericanEnglish:      alfonsoAmericanEnglishPhrase,
		language.CanadianFrench:       alfonsoCanadianFrenchPhrase,
		language.Dutch:                alfonsoDutchPhrase,
		language.French:               alfonsoFrenchPhrase,
		language.German:               alfonsoGermanPhrase,
		language.Italian:              alfonsoItalianPhrase,
		language.Japanese:             alfonsoJapanesePhrase,
		language.Korean:               alfonsoKoreanPhrase,
		language.LatinAmericanSpanish: alfonsoLatinAmericanSpanishPhrase,
		language.Russian:              alfonsoRussianPhrase,
		language.SimplifiedChinese:    alfonsoSimplifiedChinesePhrase,
		language.Spanish:              alfonsoSpanishPhrase,
		language.TraditionalChinese:   alfonsoTraditionalChinesePhrase}
)

var (
	Alfonso = nook.Villager{
		Character:   alfonsoCharacter,
		Personality: personality.Lazy,
		Phrase:      alfonsoPhrase}
)
