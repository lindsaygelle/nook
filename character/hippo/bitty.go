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
	bittyBirthday = nook.Birthday{
		Day:   6,
		Month: time.October}
)

var (
	bittyCode = nook.Code{
		Value: "hip05"}
)

var (
	bittyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bitty"}

	bittyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Potama"}

	bittyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bitty"}

	bittyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Potama"}

	bittyGermanName = nook.Name{
		Language: language.German,
		Value:    "Biggi"}

	bittyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosetta"}

	bittyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エーミー"}

	bittyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "비티"}

	bittyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ofelia"}

	bittyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Битти"}

	bittySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艾咪"}

	bittySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ofelia"}

	bittyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艾咪"}
)

var (
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
	bittyCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: bittyBirthday,
		Code:     bittyCode,
		Key:      character.Bitty,
		Gender:   gender.Female,
		Name:     bittyName}
)

var (
	bittyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "my dear"}

	bittyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "joli cœur"}

	bittyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "liefje"}

	bittyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "joli cœur"}

	bittyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mei o mei"}

	bittyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bibù"}

	bittyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だピー"}

	bittyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쩝쩝"}

	bittyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hipopop"}

	bittyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дорогуша"}

	bittySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哔哔"}

	bittySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "corassón"}

	bittyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗶嗶"}
)

var (
	bittyPhrase = nook.Languages{
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
	Bitty = nook.Villager{
		Character:   bittyCharacter,
		Personality: personality.Snooty,
		Phrase:      bittyPhrase}
)
