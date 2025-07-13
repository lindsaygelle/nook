package octopus

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
	// zuckerBirthday represents zucker birthday.
	zuckerBirthday = nook.Birthday{
		Day:   8,
		Month: time.March}
)

var (
	// zuckerCode represents zucker code.
	zuckerCode = nook.Code{
		Value: "ocp02"}
)

var (
	// zuckerAmericanEnglishName represents zucker american english name.
	zuckerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Zucker"}

	// zuckerCanadianFrenchName represents zucker canadian french name.
	zuckerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marvin"}

	// zuckerDutchName represents zucker dutch name.
	zuckerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Zucker"}

	// zuckerFrenchName represents zucker french name.
	zuckerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marvin"}

	// zuckerGermanName represents zucker german name.
	zuckerGermanName = nook.Name{
		Language: language.German,
		Value:    "Ottokar"}

	// zuckerItalianName represents zucker italian name.
	zuckerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Poldo"}

	// zuckerJapaneseName represents zucker japanese name.
	zuckerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タコヤ"}

	// zuckerKoreanName represents zucker korean name.
	zuckerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "탁호"}

	// zuckerLatinAmericanSpanishName represents zucker latin american spanish name.
	zuckerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paulino"}

	// zuckerRussianName represents zucker russian name.
	zuckerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Цукер"}

	// zuckerSimplifiedChineseName represents zucker simplified chinese name.
	zuckerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "章丸丸"}

	// zuckerSpanishName represents zucker spanish name.
	zuckerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paulino"}

	// zuckerTraditionalChineseName represents zucker traditional chinese name.
	zuckerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "章丸丸"}
)

var (
	// zuckerName represents zucker name.
	zuckerName = nook.Languages{
		language.AmericanEnglish:      zuckerAmericanEnglishName,
		language.CanadianFrench:       zuckerCanadianFrenchName,
		language.Dutch:                zuckerDutchName,
		language.French:               zuckerFrenchName,
		language.German:               zuckerGermanName,
		language.Italian:              zuckerItalianName,
		language.Japanese:             zuckerJapaneseName,
		language.Korean:               zuckerKoreanName,
		language.LatinAmericanSpanish: zuckerLatinAmericanSpanishName,
		language.Russian:              zuckerRussianName,
		language.SimplifiedChinese:    zuckerSimplifiedChineseName,
		language.Spanish:              zuckerSpanishName,
		language.TraditionalChinese:   zuckerTraditionalChineseName}
)

var (
	// zuckerCharacter represents zucker character.
	zuckerCharacter = nook.Character{
		Animal:   animal.Octopus,
		Birthday: zuckerBirthday,
		Code:     zuckerCode,
		Key:      character.Zucker,
		Gender:   gender.Male,
		Name:     zuckerName,
		Special:  false}
)

var (
	// zuckerAmericanEnglishPhrase represents zucker american english phrase.
	zuckerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bloop"}

	// zuckerCanadianFrenchPhrase represents zucker canadian french phrase.
	zuckerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sploush"}

	// zuckerDutchPhrase represents zucker dutch phrase.
	zuckerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bloep"}

	// zuckerFrenchPhrase represents zucker french phrase.
	zuckerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sploush"}

	// zuckerGermanPhrase represents zucker german phrase.
	zuckerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "saug"}

	// zuckerItalianPhrase represents zucker italian phrase.
	zuckerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blub"}

	// zuckerJapanesePhrase represents zucker japanese phrase.
	zuckerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "せやねん"}

	// zuckerKoreanPhrase represents zucker korean phrase.
	zuckerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "약히"}

	// zuckerLatinAmericanSpanishPhrase represents zucker latin american spanish phrase.
	zuckerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "churubú"}

	// zuckerRussianPhrase represents zucker russian phrase.
	zuckerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "буль"}

	// zuckerSimplifiedChinesePhrase represents zucker simplified chinese phrase.
	zuckerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "认同"}

	// zuckerSpanishPhrase represents zucker spanish phrase.
	zuckerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "churubú"}

	// zuckerTraditionalChinesePhrase represents zucker traditional chinese phrase.
	zuckerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "認同"}
)

var (
	// zuckerPhrase represents zucker phrase.
	zuckerPhrase = nook.Languages{
		language.AmericanEnglish:      zuckerAmericanEnglishPhrase,
		language.CanadianFrench:       zuckerCanadianFrenchPhrase,
		language.Dutch:                zuckerDutchPhrase,
		language.French:               zuckerFrenchPhrase,
		language.German:               zuckerGermanPhrase,
		language.Italian:              zuckerItalianPhrase,
		language.Japanese:             zuckerJapanesePhrase,
		language.Korean:               zuckerKoreanPhrase,
		language.LatinAmericanSpanish: zuckerLatinAmericanSpanishPhrase,
		language.Russian:              zuckerRussianPhrase,
		language.SimplifiedChinese:    zuckerSimplifiedChinesePhrase,
		language.Spanish:              zuckerSpanishPhrase,
		language.TraditionalChinese:   zuckerTraditionalChinesePhrase}
)

var (
	// Zucker represents zucker.
	Zucker = nook.Villager{
		Character:   zuckerCharacter,
		Personality: personality.Lazy,
		Phrase:      zuckerPhrase}
)
