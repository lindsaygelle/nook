package bull

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	chuckBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	chuckCode = nook.Code{
		Value: ""}
)

var (
	chuckAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chuck"}

	chuckCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	chuckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	chuckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Léon"}

	chuckGermanName = nook.Name{
		Language: language.German,
		Value:    "Paco"}

	chuckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bullo"}

	chuckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビフテキ"}

	chuckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	chuckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	chuckRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	chuckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "拉面"}

	chuckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tauricio"}

	chuckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	chuckName = nook.Languages{
		language.AmericanEnglish:      chuckAmericanEnglishName,
		language.CanadianFrench:       chuckCanadianFrenchName,
		language.Dutch:                chuckDutchName,
		language.French:               chuckFrenchName,
		language.German:               chuckGermanName,
		language.Italian:              chuckItalianName,
		language.Japanese:             chuckJapaneseName,
		language.Korean:               chuckKoreanName,
		language.LatinAmericanSpanish: chuckLatinAmericanSpanishName,
		language.Russian:              chuckRussianName,
		language.SimplifiedChinese:    chuckSimplifiedChineseName,
		language.Spanish:              chuckSpanishName,
		language.TraditionalChinese:   chuckTraditionalChineseName}
)

var (
	chuckCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: chuckBirthday,
		Code:     chuckCode,
		Gender:   gender.Male,
		Name:     chuckName}
)

var (
	chuckAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "jerky"}

	chuckCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	chuckDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	chuckFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "dugong"}

	chuckGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "daddel"}

	chuckItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mummamia"}

	chuckJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ってんだ"}

	chuckKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	chuckLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	chuckRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	chuckSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "伙计"}

	chuckSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cobarde"}

	chuckTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	chuckPhrase = nook.Languages{
		language.AmericanEnglish:      chuckAmericanEnglishName,
		language.CanadianFrench:       chuckCanadianFrenchName,
		language.Dutch:                chuckDutchName,
		language.French:               chuckFrenchName,
		language.German:               chuckGermanName,
		language.Italian:              chuckItalianName,
		language.Japanese:             chuckJapaneseName,
		language.Korean:               chuckKoreanName,
		language.LatinAmericanSpanish: chuckLatinAmericanSpanishName,
		language.Russian:              chuckRussianName,
		language.SimplifiedChinese:    chuckSimplifiedChineseName,
		language.Spanish:              chuckSpanishName,
		language.TraditionalChinese:   chuckTraditionalChineseName}
)

var (
	Chuck = nook.Villager{
		Character:   chuckCharacter,
		Personality: personality.Cranky,
		Phrase:      chuckPhrase}
)
