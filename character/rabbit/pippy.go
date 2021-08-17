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
	pippyBirthday = nook.Birthday{
		Day:   14,
		Month: time.June}
)

var (
	pippyCode = nook.Code{
		Value: "rbt06"}
)

var (
	pippyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pippy"}

	pippyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nadia"}

	pippyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pippy"}

	pippyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nadia"}

	pippyGermanName = nook.Name{
		Language: language.German,
		Value:    "Lotta"}

	pippyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pippi"}

	pippyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロッタ"}

	pippyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "로타"}

	pippyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Méloni"}

	pippyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пиппи"}

	pippySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗塔"}

	pippySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Méloni"}

	pippyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅塔"}
)

var (
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
	pippyCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: pippyBirthday,
		Code:     pippyCode,
		Key:      character.Pippy,
		Gender:   gender.Female,
		Name:     pippyName}
)

var (
	pippyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "li'l hare"}

	pippyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ti' lièvre"}

	pippyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "haasje"}

	pippyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ti' lièvre"}

	pippyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hihi"}

	pippyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gnaf"}

	pippyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのサ"}

	pippyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "울랄라"}

	pippyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "requeboing"}

	pippyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайчишка"}

	pippySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就这样啦"}

	pippySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pipita"}

	pippyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就這樣啦"}
)

var (
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
	Pippy = nook.Villager{
		Character:   pippyCharacter,
		Personality: personality.Peppy,
		Phrase:      pippyPhrase}
)
