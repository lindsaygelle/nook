package bull

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Léondugong"}

	chuckGermanName = nook.Name{
		Language: language.German,
		Value:    "Pacodaddel"}

	chuckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bullomummamia"}

	chuckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビフテキってんだ"}

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
		Value:    "拉面伙计"}

	chuckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tauriciocobarde"}

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
		Animal:   Bull,
		Birthday: chuckBirthday,
		Code:     chuckCode,
		Gender:   nook.Male,
		Name:     chuckName}
)

var (
	chuckAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	chuckCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	chuckDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	chuckLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	chuckRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	chuckSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "伙计"}

	chuckSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cobarde"}

	chuckTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Cranky,
		Phrase:      chuckPhrase}
)
