package hippo

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
	// bittyBirthday represents bitty birthday.
	bittyBirthday = nook.Birthday{
		Day:   6,
		Month: time.October}
)

var (
	// bittyCode represents bitty code.
	bittyCode = nook.Code{
		Value: "hip05"}
)

var (
	// bittyAmericanEnglishName represents bitty american english name.
	bittyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bitty"}

	// bittyCanadianFrenchName represents bitty canadian french name.
	bittyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Potama"}

	// bittyDutchName represents bitty dutch name.
	bittyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bitty"}

	// bittyFrenchName represents bitty french name.
	bittyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Potama"}

	// bittyGermanName represents bitty german name.
	bittyGermanName = nook.Name{
		Language: language.German,
		Value:    "Biggi"}

	// bittyItalianName represents bitty italian name.
	bittyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosetta"}

	// bittyJapaneseName represents bitty japanese name.
	bittyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エーミー"}

	// bittyKoreanName represents bitty korean name.
	bittyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "비티"}

	// bittyLatinAmericanSpanishName represents bitty latin american spanish name.
	bittyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ofelia"}

	// bittyRussianName represents bitty russian name.
	bittyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Битти"}

	// bittySimplifiedChineseName represents bitty simplified chinese name.
	bittySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艾咪"}

	// bittySpanishName represents bitty spanish name.
	bittySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ofelia"}

	// bittyTraditionalChineseName represents bitty traditional chinese name.
	bittyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艾咪"}
)

var (
	// bittyName represents bitty name.
	bittyName = nook.Languages{
		language.AmericanEnglish:      bittyAmericanEnglishName,
		language.CanadianFrench:       bittyCanadianFrenchName,
		language.Dutch:                bittyDutchName,
		language.French:               bittyFrenchName,
		language.German:               bittyGermanName,
		language.Italian:              bittyItalianName,
		language.Japanese:             bittyJapaneseName,
		language.Korean:               bittyKoreanName,
		language.LatinAmericanSpanish: bittyLatinAmericanSpanishName,
		language.Russian:              bittyRussianName,
		language.SimplifiedChinese:    bittySimplifiedChineseName,
		language.Spanish:              bittySpanishName,
		language.TraditionalChinese:   bittyTraditionalChineseName}
)

var (
	// bittyCharacter represents bitty character.
	bittyCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: bittyBirthday,
		Code:     bittyCode,
		Key:      character.Bitty,
		Gender:   gender.Female,
		Name:     bittyName,
		Special:  false}
)

var (
	// bittyAmericanEnglishPhrase represents bitty american english phrase.
	bittyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "my dear"}

	// bittyCanadianFrenchPhrase represents bitty canadian french phrase.
	bittyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "joli cœur"}

	// bittyDutchPhrase represents bitty dutch phrase.
	bittyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "liefje"}

	// bittyFrenchPhrase represents bitty french phrase.
	bittyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "joli cœur"}

	// bittyGermanPhrase represents bitty german phrase.
	bittyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mei o mei"}

	// bittyItalianPhrase represents bitty italian phrase.
	bittyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bibù"}

	// bittyJapanesePhrase represents bitty japanese phrase.
	bittyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だピー"}

	// bittyKoreanPhrase represents bitty korean phrase.
	bittyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쩝쩝"}

	// bittyLatinAmericanSpanishPhrase represents bitty latin american spanish phrase.
	bittyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hipopop"}

	// bittyRussianPhrase represents bitty russian phrase.
	bittyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дорогуша"}

	// bittySimplifiedChinesePhrase represents bitty simplified chinese phrase.
	bittySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哔哔"}

	// bittySpanishPhrase represents bitty spanish phrase.
	bittySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "corassón"}

	// bittyTraditionalChinesePhrase represents bitty traditional chinese phrase.
	bittyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗶嗶"}
)

var (
	// bittyPhrase represents bitty phrase.
	bittyPhrase = nook.Languages{
		language.AmericanEnglish:      bittyAmericanEnglishPhrase,
		language.CanadianFrench:       bittyCanadianFrenchPhrase,
		language.Dutch:                bittyDutchPhrase,
		language.French:               bittyFrenchPhrase,
		language.German:               bittyGermanPhrase,
		language.Italian:              bittyItalianPhrase,
		language.Japanese:             bittyJapanesePhrase,
		language.Korean:               bittyKoreanPhrase,
		language.LatinAmericanSpanish: bittyLatinAmericanSpanishPhrase,
		language.Russian:              bittyRussianPhrase,
		language.SimplifiedChinese:    bittySimplifiedChinesePhrase,
		language.Spanish:              bittySpanishPhrase,
		language.TraditionalChinese:   bittyTraditionalChinesePhrase}
)

var (
	// Bitty represents bitty.
	Bitty = nook.Villager{
		Character:   bittyCharacter,
		Personality: personality.Snooty,
		Phrase:      bittyPhrase}
)
