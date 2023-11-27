package squirrel

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
	// sylvanaBirthday represents Sylvana's birthday.
	sylvanaBirthday = nook.Birthday{
		Day:   22,
		Month: time.October}
)

var (
	// sylvanaCode represents Sylvana's unique code.
	sylvanaCode = nook.Code{
		Value: "squ14"}
)

var (
	// sylvanaAmericanEnglishName represents Sylvana's name in American English.
	sylvanaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sylvana"}

	// sylvanaCanadianFrenchName represents Sylvana's name in Canadian French.
	sylvanaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mounia"}

	// sylvanaDutchName represents Sylvana's name in Dutch.
	sylvanaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sylvana"}

	// sylvanaFrenchName represents Sylvana's name in French.
	sylvanaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mounia"}

	// sylvanaGermanName represents Sylvana's name in German.
	sylvanaGermanName = nook.Name{
		Language: language.German,
		Value:    "Maren"}

	// sylvanaItalianName represents Sylvana's name in Italian.
	sylvanaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Silvana"}

	// sylvanaJapaneseName represents Sylvana's name in Japanese.
	sylvanaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "もんぺ"}

	// sylvanaKoreanName represents Sylvana's name in Korean.
	sylvanaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "실바나"}

	// sylvanaLatinAmericanSpanishName represents Sylvana's name in Latin American Spanish.
	sylvanaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Silvana"}

	// sylvanaRussianName represents Sylvana's name in Russian.
	sylvanaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сильвана"}

	// sylvanaSimplifiedChineseName represents Sylvana's name in Simplified Chinese.
	sylvanaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "孟珮"}

	// sylvanaSpanishName represents Sylvana's name in Spanish.
	sylvanaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Silvana"}

	// sylvanaTraditionalChineseName represents Sylvana's name in Traditional Chinese.
	sylvanaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "孟珮"}
)

var (
	// sylvanaName represents Sylvana's name in different languages.
	sylvanaName = nook.Languages{
		language.AmericanEnglish:      sylvanaAmericanEnglishName,
		language.CanadianFrench:       sylvanaCanadianFrenchName,
		language.Dutch:                sylvanaDutchName,
		language.French:               sylvanaFrenchName,
		language.German:               sylvanaGermanName,
		language.Italian:              sylvanaItalianName,
		language.Japanese:             sylvanaJapaneseName,
		language.Korean:               sylvanaKoreanName,
		language.LatinAmericanSpanish: sylvanaLatinAmericanSpanishName,
		language.Russian:              sylvanaRussianName,
		language.SimplifiedChinese:    sylvanaSimplifiedChineseName,
		language.Spanish:              sylvanaSpanishName,
		language.TraditionalChinese:   sylvanaTraditionalChineseName}
)

var (
	// sylvanaCharacter represents Sylvana's character information.
	sylvanaCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: sylvanaBirthday,
		Code:     sylvanaCode,
		Key:      character.Sylvana,
		Gender:   gender.Female,
		Name:     sylvanaName,
		Special:  false}
)

var (
	// sylvanaAmericanEnglishPhrase represents Sylvana's phrase in American English.
	sylvanaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hubbub"}

	// sylvanaCanadianFrenchPhrase represents Sylvana's phrase in Canadian French.
	sylvanaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grignote"}

	// sylvanaDutchPhrase represents Sylvana's phrase in Dutch.
	sylvanaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jammie"}

	// sylvanaFrenchPhrase represents Sylvana's phrase in French.
	sylvanaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grignote"}

	// sylvanaGermanPhrase represents Sylvana's phrase in German.
	sylvanaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "jammi"}

	// sylvanaItalianPhrase represents Sylvana's phrase in Italian.
	sylvanaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sgranocc"}

	// sylvanaJapanesePhrase represents Sylvana's phrase in Japanese.
	sylvanaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ひゅん"}

	// sylvanaKoreanPhrase represents Sylvana's phrase in Korean.
	sylvanaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "퓨우"}

	// sylvanaLatinAmericanSpanishPhrase represents Sylvana's phrase in Latin American Spanish.
	sylvanaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "yupip"}

	// sylvanaRussianPhrase represents Sylvana's phrase in Russian.
	sylvanaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрум"}

	// sylvanaSimplifiedChinesePhrase represents Sylvana's phrase in Simplified Chinese.
	sylvanaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咻"}

	// sylvanaSpanishPhrase represents Sylvana's phrase in Spanish.
	sylvanaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "caldito"}

	// sylvanaTraditionalChinesePhrase represents Sylvana's phrase in Traditional Chinese.
	sylvanaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咻"}
)

var (
	// sylvanaPhrase represents Sylvana's phrases in different languages.
	sylvanaPhrase = nook.Languages{
		language.AmericanEnglish:      sylvanaAmericanEnglishPhrase,
		language.CanadianFrench:       sylvanaCanadianFrenchPhrase,
		language.Dutch:                sylvanaDutchPhrase,
		language.French:               sylvanaFrenchPhrase,
		language.German:               sylvanaGermanPhrase,
		language.Italian:              sylvanaItalianPhrase,
		language.Japanese:             sylvanaJapanesePhrase,
		language.Korean:               sylvanaKoreanPhrase,
		language.LatinAmericanSpanish: sylvanaLatinAmericanSpanishPhrase,
		language.Russian:              sylvanaRussianPhrase,
		language.SimplifiedChinese:    sylvanaSimplifiedChinesePhrase,
		language.Spanish:              sylvanaSpanishPhrase,
		language.TraditionalChinese:   sylvanaTraditionalChinesePhrase}
)

var (
	// Sylvana represents the character Sylvana with her complete information.
	Sylvana = nook.Villager{
		Character:   sylvanaCharacter,
		Personality: personality.Normal,
		Phrase:      sylvanaPhrase}
)
