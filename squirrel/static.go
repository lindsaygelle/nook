package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	staticBirthday = nook.Birthday{
		Day:   9,
		Month: time.July}
)

var (
	staticCode = nook.Code{
		Value: "squ08"}
)

var (
	staticAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Static"}

	staticCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Électrokrzzt"}

	staticDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Staticbliksems"}

	staticFrenchName = nook.Name{
		Language: language.French,
		Value:    "Électrokrzzt"}

	staticGermanName = nook.Name{
		Language: language.German,
		Value:    "Rudolfkrazz"}

	staticItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Protonesguscio"}

	staticJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スパークピカッ"}

	staticKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스파크콰광"}

	staticLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Protónkreee"}

	staticRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Статиквж-ж-ж"}

	staticSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "闪电闪闪"}

	staticSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Protónavería"}

	staticTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "閃電閃閃"}
)

var (
	staticName = nook.Languages{
		language.AmericanEnglish:      staticAmericanEnglishName,
		language.CanadianFrench:       staticCanadianFrenchName,
		language.Dutch:                staticDutchName,
		language.French:               staticFrenchName,
		language.German:               staticGermanName,
		language.Italian:              staticItalianName,
		language.Japanese:             staticJapaneseName,
		language.Korean:               staticKoreanName,
		language.LatinAmericanSpanish: staticLatinAmericanSpanishName,
		language.Russian:              staticRussianName,
		language.SimplifiedChinese:    staticSimplifiedChineseName,
		language.Spanish:              staticSpanishName,
		language.TraditionalChinese:   staticTraditionalChineseName}
)

var (
	staticCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: staticBirthday,
		Code:     staticCode,
		Gender:   nook.Male,
		Name:     staticName}
)

var (
	staticAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "krzzt"}

	staticCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "krzzt"}

	staticDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bliksems"}

	staticFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "krzzt"}

	staticGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krazz"}

	staticItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sguscio"}

	staticJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ピカッ"}

	staticKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "콰광"}

	staticLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kreee"}

	staticRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вж-ж-ж"}

	staticSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "闪闪"}

	staticSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "avería"}

	staticTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "閃閃"}
)

var (
	staticPhrase = nook.Languages{
		language.AmericanEnglish:      staticAmericanEnglishName,
		language.CanadianFrench:       staticCanadianFrenchName,
		language.Dutch:                staticDutchName,
		language.French:               staticFrenchName,
		language.German:               staticGermanName,
		language.Italian:              staticItalianName,
		language.Japanese:             staticJapaneseName,
		language.Korean:               staticKoreanName,
		language.LatinAmericanSpanish: staticLatinAmericanSpanishName,
		language.Russian:              staticRussianName,
		language.SimplifiedChinese:    staticSimplifiedChineseName,
		language.Spanish:              staticSpanishName,
		language.TraditionalChinese:   staticTraditionalChineseName}
)

var (
	Static = nook.Villager{
		Character:   staticCharacter,
		Personality: nook.Cranky,
		Phrase:      staticPhrase}
)
