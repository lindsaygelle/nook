package penguin

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
	// gwenBirthday represents gwen birthday.
	gwenBirthday = nook.Birthday{
		Day:   23,
		Month: time.January}
)

var (
	// gwenCode represents gwen code.
	gwenCode = nook.Code{
		Value: "pgn05"}
)

var (
	// gwenAmericanEnglishName represents gwen american english name.
	gwenAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gwen"}

	// gwenCanadianFrenchName represents gwen canadian french name.
	gwenCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gwen"}

	// gwenDutchName represents gwen dutch name.
	gwenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gwen"}

	// gwenFrenchName represents gwen french name.
	gwenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gwen"}

	// gwenGermanName represents gwen german name.
	gwenGermanName = nook.Name{
		Language: language.German,
		Value:    "Judith"}

	// gwenItalianName represents gwen italian name.
	gwenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gelinda"}

	// gwenJapaneseName represents gwen japanese name.
	gwenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ポーラ"}

	// gwenKoreanName represents gwen korean name.
	gwenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "폴라"}

	// gwenLatinAmericanSpanishName represents gwen latin american spanish name.
	gwenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fabiola"}

	// gwenRussianName represents gwen russian name.
	gwenRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гвен"}

	// gwenSimplifiedChineseName represents gwen simplified chinese name.
	gwenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宝拉"}

	// gwenSpanishName represents gwen spanish name.
	gwenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fabiola"}

	// gwenTraditionalChineseName represents gwen traditional chinese name.
	gwenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "寶拉"}
)

var (
	// gwenName represents gwen name.
	gwenName = nook.Languages{
		language.AmericanEnglish:      gwenAmericanEnglishName,
		language.CanadianFrench:       gwenCanadianFrenchName,
		language.Dutch:                gwenDutchName,
		language.French:               gwenFrenchName,
		language.German:               gwenGermanName,
		language.Italian:              gwenItalianName,
		language.Japanese:             gwenJapaneseName,
		language.Korean:               gwenKoreanName,
		language.LatinAmericanSpanish: gwenLatinAmericanSpanishName,
		language.Russian:              gwenRussianName,
		language.SimplifiedChinese:    gwenSimplifiedChineseName,
		language.Spanish:              gwenSpanishName,
		language.TraditionalChinese:   gwenTraditionalChineseName}
)

var (
	// gwenCharacter represents gwen character.
	gwenCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: gwenBirthday,
		Code:     gwenCode,
		Key:      character.Gwen,
		Gender:   gender.Female,
		Name:     gwenName,
		Special:  false}
)

var (
	// gwenAmericanEnglishPhrase represents gwen american english phrase.
	gwenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "h-h-h-hon"}

	// gwenCanadianFrenchPhrase represents gwen canadian french phrase.
	gwenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "h-h-h-hon"}

	// gwenDutchPhrase represents gwen dutch phrase.
	gwenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "d-d-duifje"}

	// gwenFrenchPhrase represents gwen french phrase.
	gwenFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "h-h-h-hon"}

	// gwenGermanPhrase represents gwen german phrase.
	gwenGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "sssweetie"}

	// gwenItalianPhrase represents gwen italian phrase.
	gwenItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "a-a-amor"}

	// gwenJapanesePhrase represents gwen japanese phrase.
	gwenJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウフフ"}

	// gwenKoreanPhrase represents gwen korean phrase.
	gwenKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우훗훗"}

	// gwenLatinAmericanSpanishPhrase represents gwen latin american spanish phrase.
	gwenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ki-ki-kiú"}

	// gwenRussianPhrase represents gwen russian phrase.
	gwenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "м-милашка"}

	// gwenSimplifiedChinesePhrase represents gwen simplified chinese phrase.
	gwenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唔呼呼"}

	// gwenSpanishPhrase represents gwen spanish phrase.
	gwenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "crustáceo"}

	// gwenTraditionalChinesePhrase represents gwen traditional chinese phrase.
	gwenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唔呼呼"}
)

var (
	// gwenPhrase represents gwen phrase.
	gwenPhrase = nook.Languages{
		language.AmericanEnglish:      gwenAmericanEnglishPhrase,
		language.CanadianFrench:       gwenCanadianFrenchPhrase,
		language.Dutch:                gwenDutchPhrase,
		language.French:               gwenFrenchPhrase,
		language.German:               gwenGermanPhrase,
		language.Italian:              gwenItalianPhrase,
		language.Japanese:             gwenJapanesePhrase,
		language.Korean:               gwenKoreanPhrase,
		language.LatinAmericanSpanish: gwenLatinAmericanSpanishPhrase,
		language.Russian:              gwenRussianPhrase,
		language.SimplifiedChinese:    gwenSimplifiedChinesePhrase,
		language.Spanish:              gwenSpanishPhrase,
		language.TraditionalChinese:   gwenTraditionalChinesePhrase}
)

var (
	// Gwen represents gwen.
	Gwen = nook.Villager{
		Character:   gwenCharacter,
		Personality: personality.Snooty,
		Phrase:      gwenPhrase}
)
