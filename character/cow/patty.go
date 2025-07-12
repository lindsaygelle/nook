package cow

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
	// pattyBirthday represents patty birthday.
	pattyBirthday = nook.Birthday{
		Day:   10,
		Month: time.May}
)

var (
	// pattyCode represents patty code.
	pattyCode = nook.Code{
		Value: "cow00"}
)

var (
	// pattyAmericanEnglishName represents patty american english name.
	pattyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Patty"}

	// pattyCanadianFrenchName represents patty canadian french name.
	pattyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Margaux"}

	// pattyDutchName represents patty dutch name.
	pattyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Patty"}

	// pattyFrenchName represents patty french name.
	pattyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Margaux"}

	// pattyGermanName represents patty german name.
	pattyGermanName = nook.Name{
		Language: language.German,
		Value:    "Patricia"}

	// pattyItalianName represents patty italian name.
	pattyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Patty"}

	// pattyJapaneseName represents patty japanese name.
	pattyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カルピ"}

	// pattyKoreanName represents patty korean name.
	pattyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "밀크"}

	// pattyLatinAmericanSpanishName represents patty latin american spanish name.
	pattyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vacarena"}

	// pattyRussianName represents patty russian name.
	pattyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пэтти"}

	// pattySimplifiedChineseName represents patty simplified chinese name.
	pattySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿排"}

	// pattySpanishName represents patty spanish name.
	pattySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vacarena"}

	// pattyTraditionalChineseName represents patty traditional chinese name.
	pattyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿排"}
)

var (
	// pattyName represents patty name.
	pattyName = nook.Languages{
		language.AmericanEnglish:      pattyAmericanEnglishName,
		language.CanadianFrench:       pattyCanadianFrenchName,
		language.Dutch:                pattyDutchName,
		language.French:               pattyFrenchName,
		language.German:               pattyGermanName,
		language.Italian:              pattyItalianName,
		language.Japanese:             pattyJapaneseName,
		language.Korean:               pattyKoreanName,
		language.LatinAmericanSpanish: pattyLatinAmericanSpanishName,
		language.Russian:              pattyRussianName,
		language.SimplifiedChinese:    pattySimplifiedChineseName,
		language.Spanish:              pattySpanishName,
		language.TraditionalChinese:   pattyTraditionalChineseName}
)

var (
	// pattyCharacter represents patty character.
	pattyCharacter = nook.Character{
		Animal:   animal.Cow,
		Birthday: pattyBirthday,
		Code:     pattyCode,
		Key:      character.Patty,
		Gender:   gender.Female,
		Name:     pattyName,
		Special:  false}
)

var (
	// pattyAmericanEnglishPhrase represents patty american english phrase.
	pattyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "how-now"}

	// pattyCanadianFrenchPhrase represents patty canadian french phrase.
	pattyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "oui oui"}

	// pattyDutchPhrase represents patty dutch phrase.
	pattyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "koebeest"}

	// pattyFrenchPhrase represents patty french phrase.
	pattyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "allez"}

	// pattyGermanPhrase represents patty german phrase.
	pattyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "muuhna"}

	// pattyItalianPhrase represents patty italian phrase.
	pattyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "truuulà"}

	// pattyJapanesePhrase represents patty japanese phrase.
	pattyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だモー"}

	// pattyKoreanPhrase represents patty korean phrase.
	pattyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "음메"}

	// pattyLatinAmericanSpanishPhrase represents patty latin american spanish phrase.
	pattyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "muuuu"}

	// pattyRussianPhrase represents patty russian phrase.
	pattyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "му-му"}

	// pattySimplifiedChinesePhrase represents patty simplified chinese phrase.
	pattySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "牟"}

	// pattySpanishPhrase represents patty spanish phrase.
	pattySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muuuu"}

	// pattyTraditionalChinesePhrase represents patty traditional chinese phrase.
	pattyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "牟"}
)

var (
	// pattyPhrase represents patty phrase.
	pattyPhrase = nook.Languages{
		language.AmericanEnglish:      pattyAmericanEnglishPhrase,
		language.CanadianFrench:       pattyCanadianFrenchPhrase,
		language.Dutch:                pattyDutchPhrase,
		language.French:               pattyFrenchPhrase,
		language.German:               pattyGermanPhrase,
		language.Italian:              pattyItalianPhrase,
		language.Japanese:             pattyJapanesePhrase,
		language.Korean:               pattyKoreanPhrase,
		language.LatinAmericanSpanish: pattyLatinAmericanSpanishPhrase,
		language.Russian:              pattyRussianPhrase,
		language.SimplifiedChinese:    pattySimplifiedChinesePhrase,
		language.Spanish:              pattySpanishPhrase,
		language.TraditionalChinese:   pattyTraditionalChinesePhrase}
)

var (
	// Patty represents patty.
	Patty = nook.Villager{
		Character:   pattyCharacter,
		Personality: personality.Peppy,
		Phrase:      pattyPhrase}
)
