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
	// montyBirthday represents monty birthday.
	montyBirthday = nook.Birthday{
		Day:   7,
		Month: time.December}
)

var (
	// montyCode represents monty code.
	montyCode = nook.Code{
		Value: "mnk04"}
)

var (
	// montyAmericanEnglishName represents monty american english name.
	montyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Monty"}

	// montyCanadianFrenchName represents monty canadian french name.
	montyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lourant"}

	// montyDutchName represents monty dutch name.
	montyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Monty"}

	// montyFrenchName represents monty french name.
	montyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lourant"}

	// montyGermanName represents monty german name.
	montyGermanName = nook.Name{
		Language: language.German,
		Value:    "Daniel"}

	// montyItalianName represents monty italian name.
	montyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mimmo"}

	// montyJapaneseName represents monty japanese name.
	montyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サルモンティ"}

	// montyKoreanName represents monty korean name.
	montyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "몽티"}

	// montyLatinAmericanSpanishName represents monty latin american spanish name.
	montyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Burton"}

	// montyRussianName represents monty russian name.
	montyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Монти"}

	// montySimplifiedChineseName represents monty simplified chinese name.
	montySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "皮猴"}

	// montySpanishName represents monty spanish name.
	montySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Burton"}

	// montyTraditionalChineseName represents monty traditional chinese name.
	montyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "皮猴"}
)

var (
	// montyName represents monty name.
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
	// montyCharacter represents monty character.
	montyCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: montyBirthday,
		Code:     montyCode,
		Key:      character.Monty,
		Gender:   gender.Male,
		Name:     montyName,
		Special:  false}
)

var (
	// montyAmericanEnglishPhrase represents monty american english phrase.
	montyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "g'tang"}

	// montyCanadianFrenchPhrase represents monty canadian french phrase.
	montyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "l'outang"}

	// montyDutchPhrase represents monty dutch phrase.
	montyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "orang"}

	// montyFrenchPhrase represents monty french phrase.
	montyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "l'outang"}

	// montyGermanPhrase represents monty german phrase.
	montyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "abbo"}

	// montyItalianPhrase represents monty italian phrase.
	montyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "g'tang"}

	// montyJapanesePhrase represents monty japanese phrase.
	montyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "バナーナ"}

	// montyKoreanPhrase represents monty korean phrase.
	montyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "바나나"}

	// montyLatinAmericanSpanishPhrase represents monty latin american spanish phrase.
	montyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uuuuh-ah"}

	// montyRussianPhrase represents monty russian phrase.
	montyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ух-эх"}

	// montySimplifiedChinesePhrase represents monty simplified chinese phrase.
	montySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "香蕉"}

	// montySpanishPhrase represents monty spanish phrase.
	montySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uuuuh-ah"}

	// montyTraditionalChinesePhrase represents monty traditional chinese phrase.
	montyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "香蕉"}
)

var (
	// montyPhrase represents monty phrase.
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
	// Monty represents monty.
	Monty = nook.Villager{
		Character:   montyCharacter,
		Personality: personality.Cranky,
		Phrase:      montyPhrase}
)
