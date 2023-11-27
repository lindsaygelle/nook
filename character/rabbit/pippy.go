package rabbit

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
	// pippyBirthday represents Pippy's birthday.
	pippyBirthday = nook.Birthday{
		Day:   14,
		Month: time.June}
)

var (
	// pippyCode represents Pippy's unique code.
	pippyCode = nook.Code{
		Value: "rbt06"}
)

var (
	// pippyAmericanEnglishName represents Pippy's name in American English.
	pippyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pippy"}

	// pippyCanadianFrenchName represents Pippy's name in Canadian French.
	pippyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nadia"}

	// pippyDutchName represents Pippy's name in Dutch.
	pippyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pippy"}

	// pippyFrenchName represents Pippy's name in French.
	pippyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nadia"}

	// pippyGermanName represents Pippy's name in German.
	pippyGermanName = nook.Name{
		Language: language.German,
		Value:    "Lotta"}

	// pippyItalianName represents Pippy's name in Italian.
	pippyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pippi"}

	// pippyJapaneseName represents Pippy's name in Japanese.
	pippyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロッタ"}

	// pippyKoreanName represents Pippy's name in Korean.
	pippyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "로타"}

	// pippyLatinAmericanSpanishName represents Pippy's name in Latin American Spanish.
	pippyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Méloni"}

	// pippyRussianName represents Pippy's name in Russian.
	pippyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пиппи"}

	// pippySimplifiedChineseName represents Pippy's name in Simplified Chinese.
	pippySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗塔"}

	// pippySpanishName represents Pippy's name in Spanish.
	pippySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Méloni"}

	// pippyTraditionalChineseName represents Pippy's name in Traditional Chinese.
	pippyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅塔"}
)

var (
	// pippyName represents Pippy's name in different languages.
	pippyName = nook.Languages{
		language.AmericanEnglish:      pippyAmericanEnglishName,
		language.CanadianFrench:       pippyCanadianFrenchName,
		language.Dutch:                pippyDutchName,
		language.French:               pippyFrenchName,
		language.German:               pippyGermanName,
		language.Italian:              pippyItalianName,
		language.Japanese:             pippyJapaneseName,
		language.Korean:               pippyKoreanName,
		language.LatinAmericanSpanish: pippyLatinAmericanSpanishName,
		language.Russian:              pippyRussianName,
		language.SimplifiedChinese:    pippySimplifiedChineseName,
		language.Spanish:              pippySpanishName,
		language.TraditionalChinese:   pippyTraditionalChineseName}
)

var (
	// pippyCharacter represents Pippy's character information.
	pippyCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: pippyBirthday,
		Code:     pippyCode,
		Key:      character.Pippy,
		Gender:   gender.Female,
		Name:     pippyName,
		Special:  false}
)

var (
	// pippyAmericanEnglishPhrase represents Pippy's phrase in American English.
	pippyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "li'l hare"}

	// pippyCanadianFrenchPhrase represents Pippy's phrase in Canadian French.
	pippyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ti' lièvre"}

	// pippyDutchPhrase represents Pippy's phrase in Dutch.
	pippyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "haasje"}

	// pippyFrenchPhrase represents Pippy's phrase in French.
	pippyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ti' lièvre"}

	// pippyGermanPhrase represents Pippy's phrase in German.
	pippyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hihi"}

	// pippyItalianPhrase represents Pippy's phrase in Italian.
	pippyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnaf"}

	// pippyJapanesePhrase represents Pippy's phrase in Japanese.
	pippyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのサ"}

	// pippyKoreanPhrase represents Pippy's phrase in Korean.
	pippyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "울랄라"}

	// pippyLatinAmericanSpanishPhrase represents Pippy's phrase in Latin American Spanish.
	pippyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "requeboing"}

	// pippyRussianPhrase represents Pippy's phrase in Russian.
	pippyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайчишка"}

	// pippySimplifiedChinesePhrase represents Pippy's phrase in Simplified Chinese.
	pippySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就这样啦"}

	// pippySpanishPhrase represents Pippy's phrase in Spanish.
	pippySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pipita"}

	// pippyTraditionalChinesePhrase represents Pippy's phrase in Traditional Chinese.
	pippyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就這樣啦"}
)

var (
	// pippyPhrase represents Pippy's phrases in different languages.
	pippyPhrase = nook.Languages{
		language.AmericanEnglish:      pippyAmericanEnglishPhrase,
		language.CanadianFrench:       pippyCanadianFrenchPhrase,
		language.Dutch:                pippyDutchPhrase,
		language.French:               pippyFrenchPhrase,
		language.German:               pippyGermanPhrase,
		language.Italian:              pippyItalianPhrase,
		language.Japanese:             pippyJapanesePhrase,
		language.Korean:               pippyKoreanPhrase,
		language.LatinAmericanSpanish: pippyLatinAmericanSpanishPhrase,
		language.Russian:              pippyRussianPhrase,
		language.SimplifiedChinese:    pippySimplifiedChinesePhrase,
		language.Spanish:              pippySpanishPhrase,
		language.TraditionalChinese:   pippyTraditionalChinesePhrase}
)

var (
	// Pippy represents the character Pippy with her complete information.
	Pippy = nook.Villager{
		Character:   pippyCharacter,
		Personality: personality.Peppy,
		Phrase:      pippyPhrase}
)
