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
	// dottyBirthday represents Dotty's birthday.
	dottyBirthday = nook.Birthday{
		Day:   14,
		Month: time.March}
)

var (
	// dottyCode represents Dotty's unique code ("rbt01").
	dottyCode = nook.Code{
		Value: "rbt01"}
)

var (
	// dottyAmericanEnglishName represents Dotty's name in American English.
	dottyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dotty"}

	// dottyCanadianFrenchName represents Dotty's name in Canadian French.
	dottyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dorothée"}

	// dottyDutchName represents Dotty's name in Dutch.
	dottyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dotty"}

	// dottyFrenchName represents Dotty's name in French.
	dottyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dorothée"}

	// dottyGermanName represents Dotty's name in German.
	dottyGermanName = nook.Name{
		Language: language.German,
		Value:    "Doro"}

	// dottyItalianName represents Dotty's name in Italian.
	dottyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dotty"}

	// dottyJapaneseName represents Dotty's name in Japanese.
	dottyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マーサ"}

	// dottyKoreanName represents Dotty's name in Korean.
	dottyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "서머"}

	// dottyLatinAmericanSpanishName represents Dotty's name in Latin American Spanish.
	dottyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Katia"}

	// dottyRussianName represents Dotty's name in Russian.
	dottyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дотти"}

	// dottySimplifiedChineseName represents Dotty's name in Simplified Chinese.
	dottySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玛莎"}

	// dottySpanishName represents Dotty's name in Spanish.
	dottySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Katia"}

	// dottyTraditionalChineseName represents Dotty's name in Traditional Chinese.
	dottyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑪莎"}
)

var (
	// dottyName represents Dotty's name in different languages.
	dottyName = nook.Languages{
		language.AmericanEnglish:      dottyAmericanEnglishName,
		language.CanadianFrench:       dottyCanadianFrenchName,
		language.Dutch:                dottyDutchName,
		language.French:               dottyFrenchName,
		language.German:               dottyGermanName,
		language.Italian:              dottyItalianName,
		language.Japanese:             dottyJapaneseName,
		language.Korean:               dottyKoreanName,
		language.LatinAmericanSpanish: dottyLatinAmericanSpanishName,
		language.Russian:              dottyRussianName,
		language.SimplifiedChinese:    dottySimplifiedChineseName,
		language.Spanish:              dottySpanishName,
		language.TraditionalChinese:   dottyTraditionalChineseName}
)

var (
	// dottyCharacter represents Dotty's character information.
	dottyCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: dottyBirthday,
		Code:     dottyCode,
		Key:      character.Dotty,
		Gender:   gender.Female,
		Name:     dottyName,
		Special:  false}
)

var (
	// dottyAmericanEnglishPhrase represents Dotty's phrase in American English.
	dottyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wee one"}

	// dottyCanadianFrenchPhrase represents Dotty's phrase in Canadian French.
	dottyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "lapinou"}

	// dottyDutchPhrase represents Dotty's phrase in Dutch.
	dottyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "lamprei"}

	// dottyFrenchPhrase represents Dotty's phrase in French.
	dottyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "lapinou"}

	// dottyGermanPhrase represents Dotty's phrase in German.
	dottyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "uiui"}

	// dottyItalianPhrase represents Dotty's phrase in Italian.
	dottyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "caroté"}

	// dottyJapanesePhrase represents Dotty's phrase in Japanese.
	dottyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ラン"}

	// dottyKoreanPhrase represents Dotty's phrase in Korean.
	dottyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "룰루랄라"}

	// dottyLatinAmericanSpanishPhrase represents Dotty's phrase in Latin American Spanish.
	dottyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "toini"}

	// dottyRussianPhrase represents Dotty's phrase in Russian.
	dottyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "малыш"}

	// dottySimplifiedChinesePhrase represents Dotty's phrase in Simplified Chinese.
	dottySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啷"}

	// dottySpanishPhrase represents Dotty's phrase in Spanish.
	dottySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trompis"}

	// dottyTraditionalChinesePhrase represents Dotty's phrase in Traditional Chinese.
	dottyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啷"}
)

var (
	// dottyPhrase represents Dotty's phrase in different languages.
	dottyPhrase = nook.Languages{
		language.AmericanEnglish:      dottyAmericanEnglishPhrase,
		language.CanadianFrench:       dottyCanadianFrenchPhrase,
		language.Dutch:                dottyDutchPhrase,
		language.French:               dottyFrenchPhrase,
		language.German:               dottyGermanPhrase,
		language.Italian:              dottyItalianPhrase,
		language.Japanese:             dottyJapanesePhrase,
		language.Korean:               dottyKoreanPhrase,
		language.LatinAmericanSpanish: dottyLatinAmericanSpanishPhrase,
		language.Russian:              dottyRussianPhrase,
		language.SimplifiedChinese:    dottySimplifiedChinesePhrase,
		language.Spanish:              dottySpanishPhrase,
		language.TraditionalChinese:   dottyTraditionalChinesePhrase}
)

var (
	// Dotty represents the character Dotty with her complete information.
	Dotty = nook.Villager{
		Character:   dottyCharacter,
		Personality: personality.Peppy,
		Phrase:      dottyPhrase}
)
