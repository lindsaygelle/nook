package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	boomerBirthday = nook.Birthday{
		Day:   7,
		Month: time.February}
)

var (
	boomerCode = nook.Code{
		Value: "pgn10"}
)

var (
	boomerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Boomer"}

	boomerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ethan"}

	boomerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Boomer"}

	boomerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ethan"}

	boomerGermanName = nook.Name{
		Language: language.German,
		Value:    "Max"}

	boomerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Icaro"}

	boomerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ショーイ"}

	boomerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "팽기"}

	boomerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Serafino"}

	boomerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бумер"}

	boomerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "秀益"}

	boomerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Serafino"}

	boomerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "秀益"}
)

var (
	boomerName = nook.Languages{
		language.AmericanEnglish:      boomerAmericanEnglishName,
		language.CanadianFrench:       boomerCanadianFrenchName,
		language.Dutch:                boomerDutchName,
		language.French:               boomerFrenchName,
		language.German:               boomerGermanName,
		language.Italian:              boomerItalianName,
		language.Japanese:             boomerJapaneseName,
		language.Korean:               boomerKoreanName,
		language.LatinAmericanSpanish: boomerLatinAmericanSpanishName,
		language.Russian:              boomerRussianName,
		language.SimplifiedChinese:    boomerSimplifiedChineseName,
		language.Spanish:              boomerSpanishName,
		language.TraditionalChinese:   boomerTraditionalChineseName}
)

var (
	boomerCharacter = nook.Character{
		Animal:   Penguin,
		Birthday: boomerBirthday,
		Code:     boomerCode,
		Gender:   nook.Male,
		Name:     boomerName}
)

var (
	boomerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "human"}

	boomerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "humanoïde"}

	boomerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mens"}

	boomerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "boudiou"}

	boomerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "menschlein"}

	boomerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "glacius"}

	boomerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ツーツツ"}

	boomerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "에헴헴"}

	boomerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "glaciux"}

	boomerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "человек"}

	boomerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "轧轧"}

	boomerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "humanoide"}

	boomerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "軋軋"}
)

var (
	boomerPhrase = nook.Languages{
		language.AmericanEnglish:      boomerAmericanEnglishName,
		language.CanadianFrench:       boomerCanadianFrenchName,
		language.Dutch:                boomerDutchName,
		language.French:               boomerFrenchName,
		language.German:               boomerGermanName,
		language.Italian:              boomerItalianName,
		language.Japanese:             boomerJapaneseName,
		language.Korean:               boomerKoreanName,
		language.LatinAmericanSpanish: boomerLatinAmericanSpanishName,
		language.Russian:              boomerRussianName,
		language.SimplifiedChinese:    boomerSimplifiedChineseName,
		language.Spanish:              boomerSpanishName,
		language.TraditionalChinese:   boomerTraditionalChineseName}
)

var (
	Boomer = nook.Villager{
		Character:   boomerCharacter,
		Personality: nook.Lazy,
		Phrase:      boomerPhrase}
)
