package monkey

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	montyBirthday = nook.Birthday{
		Day:   7,
		Month: time.December}
)

var (
	montyCode = nook.Code{
		Value: "mnk04"}
)

var (
	montyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Monty"}

	montyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lourantl'outang"}

	montyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Montyorang"}

	montyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lourantl'outang"}

	montyGermanName = nook.Name{
		Language: language.German,
		Value:    "Danielabbo"}

	montyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mimmog'tang"}

	montyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サルモンティバナーナ"}

	montyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "몽티바나나"}

	montyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Burtonuuuuh-ah"}

	montyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Монтиух-эх"}

	montySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "皮猴香蕉"}

	montySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Burtonuuuuh-ah"}

	montyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "皮猴香蕉"}
)

var (
	montyName = nook.Languages{
		language.AmericanEnglish:      montyAmericanEnglishName,
		language.CanadianFrench:       montyCanadianFrenchName,
		language.Dutch:                montyDutchName,
		language.French:               montyFrenchName,
		language.German:               montyGermanName,
		language.Italian:              montyItalianName,
		language.Japanese:             montyJapaneseName,
		language.Korean:               montyKoreanName,
		language.LatinAmericanSpanish: montyLatinAmericanSpanishName,
		language.Russian:              montyRussianName,
		language.SimplifiedChinese:    montySimplifiedChineseName,
		language.Spanish:              montySpanishName,
		language.TraditionalChinese:   montyTraditionalChineseName}
)

var (
	montyCharacter = nook.Character{
		Animal:   Monkey,
		Birthday: montyBirthday,
		Code:     montyCode,
		Gender:   nook.Male,
		Name:     montyName}
)

var (
	montyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "g'tang"}

	montyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "l'outang"}

	montyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "orang"}

	montyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "l'outang"}

	montyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "abbo"}

	montyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "g'tang"}

	montyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "バナーナ"}

	montyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "바나나"}

	montyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uuuuh-ah"}

	montyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ух-эх"}

	montySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "香蕉"}

	montySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uuuuh-ah"}

	montyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "香蕉"}
)

var (
	montyPhrase = nook.Languages{
		language.AmericanEnglish:      montyAmericanEnglishName,
		language.CanadianFrench:       montyCanadianFrenchName,
		language.Dutch:                montyDutchName,
		language.French:               montyFrenchName,
		language.German:               montyGermanName,
		language.Italian:              montyItalianName,
		language.Japanese:             montyJapaneseName,
		language.Korean:               montyKoreanName,
		language.LatinAmericanSpanish: montyLatinAmericanSpanishName,
		language.Russian:              montyRussianName,
		language.SimplifiedChinese:    montySimplifiedChineseName,
		language.Spanish:              montySpanishName,
		language.TraditionalChinese:   montyTraditionalChineseName}
)

var (
	Monty = nook.Villager{
		Character:   montyCharacter,
		Personality: nook.Cranky,
		Phrase:      montyPhrase}
)
