package dog

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
	biskitBirthday = nook.Birthday{
		Day:   13,
		Month: time.May}
)

var (
	biskitCode = nook.Code{
		Value: "dog03"}
)

var (
	biskitAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Biskit"}

	biskitCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Crocket"}

	biskitDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Biskit"}

	biskitFrenchName = nook.Name{
		Language: language.French,
		Value:    "Crocket"}

	biskitGermanName = nook.Name{
		Language: language.German,
		Value:    "Keks"}

	biskitItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Buffetto"}

	biskitJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロビン"}

	biskitKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "로빈"}

	biskitLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Amnesio"}

	biskitRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бискит"}

	biskitSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗宾"}

	biskitSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Amnesio"}

	biskitTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅賓"}
)

var (
	biskitName = nook.Languages{
		language.AmericanEnglish:      biskitAmericanEnglishName,
		language.CanadianFrench:       biskitCanadianFrenchName,
		language.Dutch:                biskitDutchName,
		language.French:               biskitFrenchName,
		language.German:               biskitGermanName,
		language.Italian:              biskitItalianName,
		language.Japanese:             biskitJapaneseName,
		language.Korean:               biskitKoreanName,
		language.LatinAmericanSpanish: biskitLatinAmericanSpanishName,
		language.Russian:              biskitRussianName,
		language.SimplifiedChinese:    biskitSimplifiedChineseName,
		language.Spanish:              biskitSpanishName,
		language.TraditionalChinese:   biskitTraditionalChineseName}
)

var (
	biskitCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: biskitBirthday,
		Code:     biskitCode,
		Key:      character.Biskit,
		Gender:   gender.Male,
		Name:     biskitName}
)

var (
	biskitAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "dawg"}

	biskitCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "clair"}

	biskitDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "makker"}

	biskitFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "clair"}

	biskitGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kläff"}

	biskitItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "brrranco"}

	biskitJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だイヌ"}

	biskitKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "멍멍"}

	biskitLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guau"}

	biskitRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пр-риятяв"}

	biskitSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "狗狗"}

	biskitSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guau"}

	biskitTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "狗狗"}
)

var (
	biskitPhrase = nook.Languages{
		language.AmericanEnglish:      biskitAmericanEnglishName,
		language.CanadianFrench:       biskitCanadianFrenchName,
		language.Dutch:                biskitDutchName,
		language.French:               biskitFrenchName,
		language.German:               biskitGermanName,
		language.Italian:              biskitItalianName,
		language.Japanese:             biskitJapaneseName,
		language.Korean:               biskitKoreanName,
		language.LatinAmericanSpanish: biskitLatinAmericanSpanishName,
		language.Russian:              biskitRussianName,
		language.SimplifiedChinese:    biskitSimplifiedChineseName,
		language.Spanish:              biskitSpanishName,
		language.TraditionalChinese:   biskitTraditionalChineseName}
)

var (
	Biskit = nook.Villager{
		Character:   biskitCharacter,
		Personality: personality.Lazy,
		Phrase:      biskitPhrase}
)
