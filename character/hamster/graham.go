package hamster

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
	// grahamBirthday represents graham birthday.
	grahamBirthday = nook.Birthday{
		Day:   20,
		Month: time.June}
)

var (
	// grahamCode represents graham code.
	grahamCode = nook.Code{
		Value: "ham02"}
)

var (
	// grahamAmericanEnglishName represents graham american english name.
	grahamAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Graham"}

	// grahamCanadianFrenchName represents graham canadian french name.
	grahamCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Graham"}

	// grahamDutchName represents graham dutch name.
	grahamDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Graham"}

	// grahamFrenchName represents graham french name.
	grahamFrenchName = nook.Name{
		Language: language.French,
		Value:    "Graham"}

	// grahamGermanName represents graham german name.
	grahamGermanName = nook.Name{
		Language: language.German,
		Value:    "Günther"}

	// grahamItalianName represents graham italian name.
	grahamItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Evaristo"}

	// grahamJapaneseName represents graham japanese name.
	grahamJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グラハム"}

	// grahamKoreanName represents graham korean name.
	grahamKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "글라햄"}

	// grahamLatinAmericanSpanishName represents graham latin american spanish name.
	grahamLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Roelio"}

	// grahamRussianName represents graham russian name.
	grahamRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Грэм"}

	// grahamSimplifiedChineseName represents graham simplified chinese name.
	grahamSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "麦麦"}

	// grahamSpanishName represents graham spanish name.
	grahamSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Roelio"}

	// grahamTraditionalChineseName represents graham traditional chinese name.
	grahamTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "麥麥"}
)

var (
	// grahamName represents graham name.
	grahamName = nook.Languages{
		language.AmericanEnglish:      grahamAmericanEnglishName,
		language.CanadianFrench:       grahamCanadianFrenchName,
		language.Dutch:                grahamDutchName,
		language.French:               grahamFrenchName,
		language.German:               grahamGermanName,
		language.Italian:              grahamItalianName,
		language.Japanese:             grahamJapaneseName,
		language.Korean:               grahamKoreanName,
		language.LatinAmericanSpanish: grahamLatinAmericanSpanishName,
		language.Russian:              grahamRussianName,
		language.SimplifiedChinese:    grahamSimplifiedChineseName,
		language.Spanish:              grahamSpanishName,
		language.TraditionalChinese:   grahamTraditionalChineseName}
)

var (
	// grahamCharacter represents graham character.
	grahamCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: grahamBirthday,
		Code:     grahamCode,
		Key:      character.Graham,
		Gender:   gender.Male,
		Name:     grahamName,
		Special:  false}
)

var (
	// grahamAmericanEnglishPhrase represents graham american english phrase.
	grahamAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "indeed"}

	// grahamCanadianFrenchPhrase represents graham canadian french phrase.
	grahamCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "au pif"}

	// grahamDutchPhrase represents graham dutch phrase.
	grahamDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "inderdaad"}

	// grahamFrenchPhrase represents graham french phrase.
	grahamFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "au pif"}

	// grahamGermanPhrase represents graham german phrase.
	grahamGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hähäm"}

	// grahamItalianPhrase represents graham italian phrase.
	grahamItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "paella"}

	// grahamJapanesePhrase represents graham japanese phrase.
	grahamJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですぞ"}

	// grahamKoreanPhrase represents graham korean phrase.
	grahamKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇다고"}

	// grahamLatinAmericanSpanishPhrase represents graham latin american spanish phrase.
	grahamLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñisqui"}

	// grahamRussianPhrase represents graham russian phrase.
	grahamRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "резонно"}

	// grahamSimplifiedChinesePhrase represents graham simplified chinese phrase.
	grahamSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就是这样"}

	// grahamSpanishPhrase represents graham spanish phrase.
	grahamSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñisqui"}

	// grahamTraditionalChinesePhrase represents graham traditional chinese phrase.
	grahamTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就是這樣"}
)

var (
	// grahamPhrase represents graham phrase.
	grahamPhrase = nook.Languages{
		language.AmericanEnglish:      grahamAmericanEnglishPhrase,
		language.CanadianFrench:       grahamCanadianFrenchPhrase,
		language.Dutch:                grahamDutchPhrase,
		language.French:               grahamFrenchPhrase,
		language.German:               grahamGermanPhrase,
		language.Italian:              grahamItalianPhrase,
		language.Japanese:             grahamJapanesePhrase,
		language.Korean:               grahamKoreanPhrase,
		language.LatinAmericanSpanish: grahamLatinAmericanSpanishPhrase,
		language.Russian:              grahamRussianPhrase,
		language.SimplifiedChinese:    grahamSimplifiedChinesePhrase,
		language.Spanish:              grahamSpanishPhrase,
		language.TraditionalChinese:   grahamTraditionalChinesePhrase}
)

var (
	// Graham represents graham.
	Graham = nook.Villager{
		Character:   grahamCharacter,
		Personality: personality.Smug,
		Phrase:      grahamPhrase}
)
