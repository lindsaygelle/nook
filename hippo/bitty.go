package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Potamajoli cœur"}

	bittyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bittyliefje"}

	bittyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Potamajoli cœur"}

	bittyGermanName = nook.Name{
		Language: language.German,
		Value:    "Biggimei o mei"}

	bittyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosettabibù"}

	bittyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エーミーだピー"}

	bittyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "비티쩝쩝"}

	bittyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ofeliahipopop"}

	bittyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Биттидорогуша"}

	bittySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艾咪哔哔"}

	bittySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ofeliacorassón"}

	bittyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艾咪嗶嗶"}
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
		Animal:   Hippo,
		Birthday: bittyBirthday,
		Code:     bittyCode,
		Gender:   nook.Female,
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
		Personality: nook.Snooty,
		Phrase:      bittyPhrase}
)
