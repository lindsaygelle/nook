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
	// staticBirthday represents Static's birthday (July 9th).
	staticBirthday = nook.Birthday{
		Day:   9,
		Month: time.July}
)

var (
	// staticCode represents Static's unique code ("squ08").
	staticCode = nook.Code{
		Value: "squ08"}
)

var (
	// staticAmericanEnglishName represents Static's name in American English.
	staticAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Static"}

	// staticCanadianFrenchName represents Static's name in Canadian French.
	staticCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Électro"}

	// staticDutchName represents Static's name in Dutch.
	staticDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Static"}

	// staticFrenchName represents Static's name in French.
	staticFrenchName = nook.Name{
		Language: language.French,
		Value:    "Électro"}

	// staticGermanName represents Static's name in German.
	staticGermanName = nook.Name{
		Language: language.German,
		Value:    "Rudolf"}

	// staticItalianName represents Static's name in Italian.
	staticItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Protone"}

	// staticJapaneseName represents Static's name in Japanese.
	staticJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スパーク"}

	// staticKoreanName represents Static's name in Korean.
	staticKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스파크"}

	// staticLatinAmericanSpanishName represents Static's name in Latin American Spanish.
	staticLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Protón"}

	// staticRussianName represents Static's name in Russian.
	staticRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Статик"}

	// staticSimplifiedChineseName represents Static's name in Simplified Chinese.
	staticSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "闪电"}

	// staticSpanishName represents Static's name in Spanish.
	staticSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Protón"}

	// staticTraditionalChineseName represents Static's name in Traditional Chinese.
	staticTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "閃電"}
)

var (
	// staticName represents Static's name in different languages.
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
	// staticCharacter represents Static's character information.
	staticCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: staticBirthday,
		Code:     staticCode,
		Key:      character.Static,
		Gender:   gender.Male,
		Name:     staticName,
		Special:  false}
)

var (
	// staticAmericanEnglishPhrase represents Static's phrase in American English.
	staticAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "krzzt"}

	// staticCanadianFrenchPhrase represents Static's phrase in Canadian French.
	staticCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "krzzt"}

	// staticDutchPhrase represents Static's phrase in Dutch.
	staticDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bliksems"}

	// staticFrenchPhrase represents Static's phrase in French.
	staticFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "krzzt"}

	// staticGermanPhrase represents Static's phrase in German.
	staticGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krazz"}

	// staticItalianPhrase represents Static's phrase in Italian.
	staticItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sguscio"}

	// staticJapanesePhrase represents Static's phrase in Japanese.
	staticJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ピカッ"}

	// staticKoreanPhrase represents Static's phrase in Korean.
	staticKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "콰광"}

	// staticLatinAmericanSpanishPhrase represents Static's phrase in Latin American Spanish.
	staticLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kreee"}

	// staticRussianPhrase represents Static's phrase in Russian.
	staticRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вж-ж-ж"}

	// staticSimplifiedChinesePhrase represents Static's phrase in Simplified Chinese.
	staticSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "闪闪"}

	// staticSpanishPhrase represents Static's phrase in Spanish.
	staticSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "avería"}

	// staticTraditionalChinesePhrase represents Static's phrase in Traditional Chinese.
	staticTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "閃閃"}
)

var (
	// staticPhrase represents Static's phrases in different languages.
	staticPhrase = nook.Languages{
		language.AmericanEnglish:      staticAmericanEnglishPhrase,
		language.CanadianFrench:       staticCanadianFrenchPhrase,
		language.Dutch:                staticDutchPhrase,
		language.French:               staticFrenchPhrase,
		language.German:               staticGermanPhrase,
		language.Italian:              staticItalianPhrase,
		language.Japanese:             staticJapanesePhrase,
		language.Korean:               staticKoreanPhrase,
		language.LatinAmericanSpanish: staticLatinAmericanSpanishPhrase,
		language.Russian:              staticRussianPhrase,
		language.SimplifiedChinese:    staticSimplifiedChinesePhrase,
		language.Spanish:              staticSpanishPhrase,
		language.TraditionalChinese:   staticTraditionalChinesePhrase}
)

var (
	// Static represents the character Static with his complete information.
	Static = nook.Villager{
		Character:   staticCharacter,
		Personality: personality.Cranky,
		Phrase:      staticPhrase}
)
