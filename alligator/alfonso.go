package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Alphonsemoi moi"}

	alfonsoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Alfonsoit⁠'⁠sa me"}

	alfonsoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Alphonsemoi moi"}

	alfonsoGermanName = nook.Name{
		Language: language.German,
		Value:    "Markusahhhrg"}

	alfonsoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Alfonsoproprio"}

	alfonsoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アルベルトだワニ"}

	alfonsoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "알베르트나아거"}

	alfonsoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kaimánñam-ñam"}

	alfonsoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Альфонсоэто я"}

	alfonsoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿泥鳄泥"}

	alfonsoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kaimánñam-ñam"}

	alfonsoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿泥鱷泥"}
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
		Animal:   Alligator,
		Birthday: alfonsoBirthday,
		Code:     alfonsoCode,
		Gender:   nook.Male,
		Name:     alfonsoName}
)

var (
	alfonsoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "it's a meit'sa me"}

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
	Alfonso = nook.Villager{
		Character:   alfonsoCharacter,
		Personality: nook.Lazy,
		Phrase:      alfonsoPhrase}
)
