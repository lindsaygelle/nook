package bearcub

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
	// martyBirthday represents marty birthday.
	martyBirthday = nook.Birthday{
		Day:   16,
		Month: time.April}
)

var (
	// martyCode represents marty code.
	martyCode = nook.Code{
		Value: "cbr18"}
)

var (
	// martyAmericanEnglishName represents marty american english name.
	martyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Marty"}

	// martyCanadianFrenchName represents marty canadian french name.
	martyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marty"}

	// martyDutchName represents marty dutch name.
	martyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marty"}

	// martyFrenchName represents marty french name.
	martyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marty"}

	// martyGermanName represents marty german name.
	martyGermanName = nook.Name{
		Language: language.German,
		Value:    "Marty"}

	// martyItalianName represents marty italian name.
	martyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marty"}

	// martyJapaneseName represents marty japanese name.
	martyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マーティー"}

	// martyKoreanName represents marty korean name.
	martyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마티"}

	// martyLatinAmericanSpanishName represents marty latin american spanish name.
	martyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marty"}

	// martyRussianName represents marty russian name.
	martyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Марти"}

	// martySimplifiedChineseName represents marty simplified chinese name.
	martySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玛丁"}

	// martySpanishName represents marty spanish name.
	martySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marty"}

	// martyTraditionalChineseName represents marty traditional chinese name.
	martyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑪丁"}
)

var (
	// martyName represents marty name.
	martyName = nook.Languages{
		language.AmericanEnglish:      martyAmericanEnglishName,
		language.CanadianFrench:       martyCanadianFrenchName,
		language.Dutch:                martyDutchName,
		language.French:               martyFrenchName,
		language.German:               martyGermanName,
		language.Italian:              martyItalianName,
		language.Japanese:             martyJapaneseName,
		language.Korean:               martyKoreanName,
		language.LatinAmericanSpanish: martyLatinAmericanSpanishName,
		language.Russian:              martyRussianName,
		language.SimplifiedChinese:    martySimplifiedChineseName,
		language.Spanish:              martySpanishName,
		language.TraditionalChinese:   martyTraditionalChineseName}
)

var (
	// martyCharacter represents marty character.
	martyCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: martyBirthday,
		Code:     martyCode,
		Key:      character.Marty,
		Gender:   gender.Male,
		Name:     martyName,
		Special:  false}
)

var (
	// martyAmericanEnglishPhrase represents marty american english phrase.
	martyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pompom"}

	// martyCanadianFrenchPhrase represents marty canadian french phrase.
	martyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pom pom"}

	// martyDutchPhrase represents marty dutch phrase.
	martyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pommetje"}

	// martyFrenchPhrase represents marty french phrase.
	martyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pom pom"}

	// martyGermanPhrase represents marty german phrase.
	martyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pompudding"}

	// martyItalianPhrase represents marty italian phrase.
	martyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pom pom"}

	// martyJapanesePhrase represents marty japanese phrase.
	martyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ポムッ"}

	// martyKoreanPhrase represents marty korean phrase.
	martyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "폼폼"}

	// martyLatinAmericanSpanishPhrase represents marty latin american spanish phrase.
	martyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pudin"}

	// martyRussianPhrase represents marty russian phrase.
	martyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "помпончик"}

	// martySimplifiedChinesePhrase represents marty simplified chinese phrase.
	martySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "布丁"}

	// martySpanishPhrase represents marty spanish phrase.
	martySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pudin"}

	// martyTraditionalChinesePhrase represents marty traditional chinese phrase.
	martyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "布丁"}
)

var (
	// martyPhrase represents marty phrase.
	martyPhrase = nook.Languages{
		language.AmericanEnglish:      martyAmericanEnglishPhrase,
		language.CanadianFrench:       martyCanadianFrenchPhrase,
		language.Dutch:                martyDutchPhrase,
		language.French:               martyFrenchPhrase,
		language.German:               martyGermanPhrase,
		language.Italian:              martyItalianPhrase,
		language.Japanese:             martyJapanesePhrase,
		language.Korean:               martyKoreanPhrase,
		language.LatinAmericanSpanish: martyLatinAmericanSpanishPhrase,
		language.Russian:              martyRussianPhrase,
		language.SimplifiedChinese:    martySimplifiedChinesePhrase,
		language.Spanish:              martySpanishPhrase,
		language.TraditionalChinese:   martyTraditionalChinesePhrase}
)

var (
	// Marty represents marty.
	Marty = nook.Villager{
		Character:   martyCharacter,
		Personality: personality.Lazy,
		Phrase:      martyPhrase}
)
