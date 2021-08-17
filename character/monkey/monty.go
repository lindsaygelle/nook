package monkey

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
		Value:    "Lourant"}

	montyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Monty"}

	montyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lourant"}

	montyGermanName = nook.Name{
		Language: language.German,
		Value:    "Daniel"}

	montyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mimmo"}

	montyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サルモンティ"}

	montyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "몽티"}

	montyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Burton"}

	montyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Монти"}

	montySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "皮猴"}

	montySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Burton"}

	montyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "皮猴"}
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
		Animal:   animal.Monkey,
		Birthday: montyBirthday,
		Code:     montyCode,
		Key:      character.Monty,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      montyAmericanEnglishPhrase,
		language.CanadianFrench:       montyCanadianFrenchPhrase,
		language.Dutch:                montyDutchPhrase,
		language.French:               montyFrenchPhrase,
		language.German:               montyGermanPhrase,
		language.Italian:              montyItalianPhrase,
		language.Japanese:             montyJapanesePhrase,
		language.Korean:               montyKoreanPhrase,
		language.LatinAmericanSpanish: montyLatinAmericanSpanishPhrase,
		language.Russian:              montyRussianPhrase,
		language.SimplifiedChinese:    montySimplifiedChinesePhrase,
		language.Spanish:              montySpanishPhrase,
		language.TraditionalChinese:   montyTraditionalChinesePhrase}
)

var (
	Monty = nook.Villager{
		Character:   montyCharacter,
		Personality: personality.Cranky,
		Phrase:      montyPhrase}
)
