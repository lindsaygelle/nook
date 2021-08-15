package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	leighBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	leighCode = nook.Code{
		Value: ""}
)

var (
	leighAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leigh"}

	leighCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	leighDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	leighFrenchName = nook.Name{
		Language: language.French,
		Value:    "Huguette"}

	leighGermanName = nook.Name{
		Language: language.German,
		Value:    "Loretta"}

	leighItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Piumetta"}

	leighJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ディアナ"}

	leighKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	leighLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	leighRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	leighSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "画眉"}

	leighSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tabita"}

	leighTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	leighName = nook.Languages{
		language.AmericanEnglish:      leighAmericanEnglishName,
		language.CanadianFrench:       leighCanadianFrenchName,
		language.Dutch:                leighDutchName,
		language.French:               leighFrenchName,
		language.German:               leighGermanName,
		language.Italian:              leighItalianName,
		language.Japanese:             leighJapaneseName,
		language.Korean:               leighKoreanName,
		language.LatinAmericanSpanish: leighLatinAmericanSpanishName,
		language.Russian:              leighRussianName,
		language.SimplifiedChinese:    leighSimplifiedChineseName,
		language.Spanish:              leighSpanishName,
		language.TraditionalChinese:   leighTraditionalChineseName}
)

var (
	leighCharacter = nook.Character{
		Animal:   Chicken,
		Birthday: leighBirthday,
		Code:     leighCode,
		Gender:   nook.Female,
		Name:     leighName}
)

var (
	leighAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	leighCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	leighDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	leighFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poussinet"}

	leighGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "goldi"}

	leighItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cocò"}

	leighJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だもん"}

	leighKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	leighLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	leighRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	leighSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啄啄"}

	leighSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chuli"}

	leighTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	leighPhrase = nook.Languages{
		language.AmericanEnglish:      leighAmericanEnglishName,
		language.CanadianFrench:       leighCanadianFrenchName,
		language.Dutch:                leighDutchName,
		language.French:               leighFrenchName,
		language.German:               leighGermanName,
		language.Italian:              leighItalianName,
		language.Japanese:             leighJapaneseName,
		language.Korean:               leighKoreanName,
		language.LatinAmericanSpanish: leighLatinAmericanSpanishName,
		language.Russian:              leighRussianName,
		language.SimplifiedChinese:    leighSimplifiedChineseName,
		language.Spanish:              leighSpanishName,
		language.TraditionalChinese:   leighTraditionalChineseName}
)

var (
	Leigh = nook.Villager{
		Character:   leighCharacter,
		Personality: nook.Peppy,
		Phrase:      leighPhrase}
)
