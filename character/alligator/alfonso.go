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
	// alfonsoBirthday represents Alfonso's birthday.
	alfonsoBirthday = nook.Birthday{
		Day:   9,
		Month: time.June}
)

var (
	// alfonsoCode represents Alfonso's unique code.
	alfonsoCode = nook.Code{
		Value: "crd00"}
)

var (
	// alfonsoAmericanEnglishName represents Alfonso's name in American English.
	alfonsoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Alfonso"}

	// alfonsoCanadianFrenchName represents Alfonso's name in Canadian French.
	alfonsoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Alphonse"}

	// alfonsoDutchName represents Alfonso's name in Dutch.
	alfonsoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Alfonso"}

	// alfonsoFrenchName represents Alfonso's name in French.
	alfonsoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Alphonse"}

	// alfonsoGermanName represents Alfonso's name in German.
	alfonsoGermanName = nook.Name{
		Language: language.German,
		Value:    "Markus"}

	// alfonsoItalianName represents Alfonso's name in Italian.
	alfonsoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Alfonso"}

	// alfonsoJapaneseName represents Alfonso's name in Japanese.
	alfonsoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アルベルト"}

	// alfonsoKoreanName represents Alfonso's name in Korean.
	alfonsoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "알베르트"}

	// alfonsoLatinAmericanSpanishName represents Alfonso's name in Latin American Spanish.
	alfonsoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kaimán"}

	// alfonsoRussianName represents Alfonso's name in Russian.
	alfonsoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Альфонсо"}

	// alfonsoSimplifiedChineseName represents Alfonso's name in Simplified Chinese.
	alfonsoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿泥"}

	// alfonsoSpanishName represents Alfonso's name in Spanish.
	alfonsoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kaimán"}

	// alfonsoTraditionalChineseName represents Alfonso's name in Traditional Chinese.
	alfonsoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿泥"}
)

var (
	// alfonsoName represents Alfonso's name in different languages.
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
	// alfonsoCharacter represents Alfonso's character information.
	alfonsoCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: alfonsoBirthday,
		Code:     alfonsoCode,
		Key:      character.Alfonso,
		Gender:   gender.Male,
		Name:     alfonsoName,
		Special:  false}
)

var (
	// alfonsoAmericanEnglishPhrase represents Alfonso's phrase in American English.
	alfonsoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "it's a me"}

	// alfonsoCanadianFrenchPhrase represents Alfonso's phrase in Canadian French.
	alfonsoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "moi moi"}

	// alfonsoDutchPhrase represents Alfonso's phrase in Dutch.
	alfonsoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "it'sa me"}

	// alfonsoFrenchPhrase represents Alfonso's phrase in French.
	alfonsoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "moi moi"}

	// alfonsoGermanPhrase represents Alfonso's phrase in German.
	alfonsoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ahhhrg"}

	// alfonsoItalianPhrase represents Alfonso's phrase in Italian.
	alfonsoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "proprio"}

	// alfonsoJapanesePhrase represents Alfonso's phrase in Japanese.
	alfonsoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だワニ"}

	// alfonsoKoreanPhrase represents Alfonso's phrase in Korean.
	alfonsoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "나아거"}

	// alfonsoLatinAmericanSpanishPhrase represents Alfonso's phrase in Latin American Spanish.
	alfonsoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñam-ñam"}

	// alfonsoRussianPhrase represents Alfonso's phrase in Russian.
	alfonsoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "это я"}

	// alfonsoSimplifiedChinesePhrase represents Alfonso's phrase in Simplified Chinese.
	alfonsoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鳄泥"}

	// alfonsoSpanishPhrase represents Alfonso's phrase in Spanish.
	alfonsoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñam-ñam"}

	// alfonsoTraditionalChinesePhrase represents Alfonso's phrase in Traditional Chinese.
	alfonsoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鱷泥"}
)

var (
	// alfonsoPhrase represents Alfonso's phrases in different languages.
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
	// Alfonso represents the character Alfonso with his complete information.
	Alfonso = nook.Villager{
		Character:   alfonsoCharacter,
		Personality: personality.Lazy,
		Phrase:      alfonsoPhrase}
)
