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
	dottyBirthday = nook.Birthday{
		Day:   14,
		Month: time.March}
)

var (
	dottyCode = nook.Code{
		Value: "rbt01"}
)

var (
	dottyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dotty"}

	dottyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dorothée"}

	dottyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dotty"}

	dottyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dorothée"}

	dottyGermanName = nook.Name{
		Language: language.German,
		Value:    "Doro"}

	dottyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dotty"}

	dottyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マーサ"}

	dottyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "서머"}

	dottyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Katia"}

	dottyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дотти"}

	dottySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玛莎"}

	dottySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Katia"}

	dottyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑪莎"}
)

var (
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
	dottyCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: dottyBirthday,
		Code:     dottyCode,
		Key:      character.Dotty,
		Gender:   gender.Female,
		Name:     dottyName}
)

var (
	dottyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wee one"}

	dottyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "lapinou"}

	dottyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "lamprei"}

	dottyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "lapinou"}

	dottyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "uiui"}

	dottyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "caroté"}

	dottyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ラン"}

	dottyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "룰루랄라"}

	dottyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "toini"}

	dottyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "малыш"}

	dottySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啷"}

	dottySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trompis"}

	dottyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啷"}
)

var (
	dottyPhrase = nook.Languages{
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
	Dotty = nook.Villager{
		Character:   dottyCharacter,
		Personality: personality.Peppy,
		Phrase:      dottyPhrase}
)
