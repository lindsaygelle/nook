package rhinoceros

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
		Value:    "Rhino"}

	spikeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Spike"}

	spikeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rhino"}

	spikeGermanName = nook.Name{
		Language: language.German,
		Value:    "Atze"}

	spikeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bruto"}

	spikeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スクワット"}

	spikeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스쿼트"}

	spikeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cornio"}

	spikeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Спайк"}

	spikeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿蹲"}

	spikeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cornio"}

	spikeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿蹲"}
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
		Animal:   animal.Rhinoceros,
		Birthday: spikeBirthday,
		Code:     spikeCode,
		Key:      character.Spike,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      spikeAmericanEnglishPhrase,
		language.CanadianFrench:       spikeCanadianFrenchPhrase,
		language.Dutch:                spikeDutchPhrase,
		language.French:               spikeFrenchPhrase,
		language.German:               spikeGermanPhrase,
		language.Italian:              spikeItalianPhrase,
		language.Japanese:             spikeJapanesePhrase,
		language.Korean:               spikeKoreanPhrase,
		language.LatinAmericanSpanish: spikeLatinAmericanSpanishPhrase,
		language.Russian:              spikeRussianPhrase,
		language.SimplifiedChinese:    spikeSimplifiedChinesePhrase,
		language.Spanish:              spikeSpanishPhrase,
		language.TraditionalChinese:   spikeTraditionalChinesePhrase}
)

var (
	Spike = nook.Villager{
		Character:   spikeCharacter,
		Personality: personality.Cranky,
		Phrase:      spikePhrase}
)
