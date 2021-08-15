package rhinoceros

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	spikeBirthday = nook.Birthday{
		Day:   17,
		Month: time.June}
)

var (
	spikeCode = nook.Code{
		Value: "rhn02"}
)

var (
	spikeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Spike"}

	spikeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rhinogroumph"}

	spikeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Spikeschelm"}

	spikeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rhinosaperlotte"}

	spikeGermanName = nook.Name{
		Language: language.German,
		Value:    "Atzepunk"}

	spikeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Brutopachi"}

	spikeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スクワットでヤンキ"}

	spikeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스쿼트빠직"}

	spikeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Corniopapúm"}

	spikeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Спайкбродяга"}

	spikeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿蹲混混"}

	spikeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Corniogranuja"}

	spikeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿蹲混混"}
)

var (
	spikeName = nook.Languages{
		language.AmericanEnglish:      spikeAmericanEnglishName,
		language.CanadianFrench:       spikeCanadianFrenchName,
		language.Dutch:                spikeDutchName,
		language.French:               spikeFrenchName,
		language.German:               spikeGermanName,
		language.Italian:              spikeItalianName,
		language.Japanese:             spikeJapaneseName,
		language.Korean:               spikeKoreanName,
		language.LatinAmericanSpanish: spikeLatinAmericanSpanishName,
		language.Russian:              spikeRussianName,
		language.SimplifiedChinese:    spikeSimplifiedChineseName,
		language.Spanish:              spikeSpanishName,
		language.TraditionalChinese:   spikeTraditionalChineseName}
)

var (
	spikeCharacter = nook.Character{
		Animal:   Rhinoceros,
		Birthday: spikeBirthday,
		Code:     spikeCode,
		Gender:   nook.Male,
		Name:     spikeName}
)

var (
	spikeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "punk"}

	spikeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "groumph"}

	spikeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schelm"}

	spikeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "saperlotte"}

	spikeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "punk"}

	spikeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pachi"}

	spikeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でヤンキ"}

	spikeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "빠직"}

	spikeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "papúm"}

	spikeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бродяга"}

	spikeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "混混"}

	spikeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "granuja"}

	spikeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "混混"}
)

var (
	spikePhrase = nook.Languages{
		language.AmericanEnglish:      spikeAmericanEnglishName,
		language.CanadianFrench:       spikeCanadianFrenchName,
		language.Dutch:                spikeDutchName,
		language.French:               spikeFrenchName,
		language.German:               spikeGermanName,
		language.Italian:              spikeItalianName,
		language.Japanese:             spikeJapaneseName,
		language.Korean:               spikeKoreanName,
		language.LatinAmericanSpanish: spikeLatinAmericanSpanishName,
		language.Russian:              spikeRussianName,
		language.SimplifiedChinese:    spikeSimplifiedChineseName,
		language.Spanish:              spikeSpanishName,
		language.TraditionalChinese:   spikeTraditionalChineseName}
)

var (
	Spike = nook.Villager{
		Character:   spikeCharacter,
		Personality: nook.Cranky,
		Phrase:      spikePhrase}
)
