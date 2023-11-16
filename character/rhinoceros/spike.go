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
	// spikeBirthday represents Spike's birthday (June 17th).
	spikeBirthday = nook.Birthday{
		Day:   17,
		Month: time.June}
)

var (
	// spikeCode represents Spike's unique code ("rhn02").
	spikeCode = nook.Code{
		Value: "rhn02"}
)

var (
	// spikeAmericanEnglishName represents Spike's name in American English.
	spikeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Spike"}

	// spikeCanadianFrenchName represents Spike's name in Canadian French.
	spikeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rhino"}

	// spikeDutchName represents Spike's name in Dutch.
	spikeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Spike"}

	// spikeFrenchName represents Spike's name in French.
	spikeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rhino"}

	// spikeGermanName represents Spike's name in German.
	spikeGermanName = nook.Name{
		Language: language.German,
		Value:    "Atze"}

	// spikeItalianName represents Spike's name in Italian.
	spikeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bruto"}

	// spikeJapaneseName represents Spike's name in Japanese.
	spikeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スクワット"}

	// spikeKoreanName represents Spike's name in Korean.
	spikeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스쿼트"}

	// spikeLatinAmericanSpanishName represents Spike's name in Latin American Spanish.
	spikeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cornio"}

	// spikeRussianName represents Spike's name in Russian.
	spikeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Спайк"}

	// spikeSimplifiedChineseName represents Spike's name in Simplified Chinese.
	spikeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿蹲"}

	// spikeSpanishName represents Spike's name in Spanish.
	spikeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cornio"}

	// spikeTraditionalChineseName represents Spike's name in Traditional Chinese.
	spikeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿蹲"}
)

var (
	// spikeName represents Spike's name in different languages.
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
	// spikeCharacter represents Spike's character information.
	spikeCharacter = nook.Character{
		Animal:   animal.Rhinoceros,
		Birthday: spikeBirthday,
		Code:     spikeCode,
		Key:      character.Spike,
		Gender:   gender.Male,
		Name:     spikeName,
		Special:  false}
)

var (
	// spikeAmericanEnglishPhrase represents Spike's phrase in American English.
	spikeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "punk"}

	// spikeCanadianFrenchPhrase represents Spike's phrase in Canadian French.
	spikeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "groumph"}

	// spikeDutchPhrase represents Spike's phrase in Dutch.
	spikeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schelm"}

	// spikeFrenchPhrase represents Spike's phrase in French.
	spikeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "saperlotte"}

	// spikeGermanPhrase represents Spike's phrase in German.
	spikeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "punk"}

	// spikeItalianPhrase represents Spike's phrase in Italian.
	spikeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pachi"}

	// spikeJapanesePhrase represents Spike's phrase in Japanese.
	spikeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でヤンキ"}

	// spikeKoreanPhrase represents Spike's phrase in Korean.
	spikeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "빠직"}

	// spikeLatinAmericanSpanishPhrase represents Spike's phrase in Latin American Spanish.
	spikeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "papúm"}

	// spikeRussianPhrase represents Spike's phrase in Russian.
	spikeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бродяга"}

	// spikeSimplifiedChinesePhrase represents Spike's phrase in Simplified Chinese.
	spikeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "混混"}

	// spikeSpanishPhrase represents Spike's phrase in Spanish.
	spikeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "granuja"}

	// spikeTraditionalChinesePhrase represents Spike's phrase in Traditional Chinese.
	spikeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "混混"}
)

var (
	// spikePhrase represents Spike's phrase in different languages.
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
	// Spike represents the character Spike with his complete information.
	Spike = nook.Villager{
		Character:   spikeCharacter,
		Personality: personality.Cranky,
		Phrase:      spikePhrase}
)
