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
	// hopperBirthday represents hopper birthday.
	hopperBirthday = nook.Birthday{
		Day:   6,
		Month: time.April}
)

var (
	// hopperCode represents hopper code.
	hopperCode = nook.Code{
		Value: "pgn03"}
)

var (
	// hopperAmericanEnglishName represents hopper american english name.
	hopperAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hopper"}

	// hopperCanadianFrenchName represents hopper canadian french name.
	hopperCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Victor"}

	// hopperDutchName represents hopper dutch name.
	hopperDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hopper"}

	// hopperFrenchName represents hopper french name.
	hopperFrenchName = nook.Name{
		Language: language.French,
		Value:    "Victor"}

	// hopperGermanName represents hopper german name.
	hopperGermanName = nook.Name{
		Language: language.German,
		Value:    "Hauke"}

	// hopperItalianName represents hopper italian name.
	hopperItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pinguaio"}

	// hopperJapaneseName represents hopper japanese name.
	hopperJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダルマン"}

	// hopperKoreanName represents hopper korean name.
	hopperKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "달만이"}

	// hopperLatinAmericanSpanishName represents hopper latin american spanish name.
	hopperLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Güiñón"}

	// hopperRussianName represents hopper russian name.
	hopperRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хоппер"}

	// hopperSimplifiedChineseName represents hopper simplified chinese name.
	hopperSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "达满"}

	// hopperSpanishName represents hopper spanish name.
	hopperSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Güiñón"}

	// hopperTraditionalChineseName represents hopper traditional chinese name.
	hopperTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "達滿"}
)

var (
	// hopperName represents hopper name.
	hopperName = nook.Languages{
		language.AmericanEnglish:      hopperAmericanEnglishName,
		language.CanadianFrench:       hopperCanadianFrenchName,
		language.Dutch:                hopperDutchName,
		language.French:               hopperFrenchName,
		language.German:               hopperGermanName,
		language.Italian:              hopperItalianName,
		language.Japanese:             hopperJapaneseName,
		language.Korean:               hopperKoreanName,
		language.LatinAmericanSpanish: hopperLatinAmericanSpanishName,
		language.Russian:              hopperRussianName,
		language.SimplifiedChinese:    hopperSimplifiedChineseName,
		language.Spanish:              hopperSpanishName,
		language.TraditionalChinese:   hopperTraditionalChineseName}
)

var (
	// hopperCharacter represents hopper character.
	hopperCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: hopperBirthday,
		Code:     hopperCode,
		Key:      character.Hopper,
		Gender:   gender.Male,
		Name:     hopperName,
		Special:  false}
)

var (
	// hopperAmericanEnglishPhrase represents hopper american english phrase.
	hopperAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "slushie"}

	// hopperCanadianFrenchPhrase represents hopper canadian french phrase.
	hopperCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "manette"}

	// hopperDutchPhrase represents hopper dutch phrase.
	hopperDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ĳsje"}

	// hopperFrenchPhrase represents hopper french phrase.
	hopperFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "manette"}

	// hopperGermanPhrase represents hopper german phrase.
	hopperGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "slaschi"}

	// hopperItalianPhrase represents hopper italian phrase.
	hopperItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sguish"}

	// hopperJapanesePhrase represents hopper japanese phrase.
	hopperJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ちぅねん"}

	// hopperKoreanPhrase represents hopper korean phrase.
	hopperKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아따"}

	// hopperLatinAmericanSpanishPhrase represents hopper latin american spanish phrase.
	hopperLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "esploch"}

	// hopperRussianPhrase represents hopper russian phrase.
	hopperRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ледок"}

	// hopperSimplifiedChinesePhrase represents hopper simplified chinese phrase.
	hopperSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "禅"}

	// hopperSpanishPhrase represents hopper spanish phrase.
	hopperSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "esploch"}

	// hopperTraditionalChinesePhrase represents hopper traditional chinese phrase.
	hopperTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "禪"}
)

var (
	// hopperPhrase represents hopper phrase.
	hopperPhrase = nook.Languages{
		language.AmericanEnglish:      hopperAmericanEnglishPhrase,
		language.CanadianFrench:       hopperCanadianFrenchPhrase,
		language.Dutch:                hopperDutchPhrase,
		language.French:               hopperFrenchPhrase,
		language.German:               hopperGermanPhrase,
		language.Italian:              hopperItalianPhrase,
		language.Japanese:             hopperJapanesePhrase,
		language.Korean:               hopperKoreanPhrase,
		language.LatinAmericanSpanish: hopperLatinAmericanSpanishPhrase,
		language.Russian:              hopperRussianPhrase,
		language.SimplifiedChinese:    hopperSimplifiedChinesePhrase,
		language.Spanish:              hopperSpanishPhrase,
		language.TraditionalChinese:   hopperTraditionalChinesePhrase}
)

var (
	// Hopper represents hopper.
	Hopper = nook.Villager{
		Character:   hopperCharacter,
		Personality: personality.Cranky,
		Phrase:      hopperPhrase}
)
