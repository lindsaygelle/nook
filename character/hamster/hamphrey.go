package hamster

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
	// hamphreyBirthday represents hamphrey birthday.
	hamphreyBirthday = nook.Birthday{
		Day:   25,
		Month: time.February}
)

var (
	// hamphreyCode represents hamphrey code.
	hamphreyCode = nook.Code{
		Value: "ham07"}
)

var (
	// hamphreyAmericanEnglishName represents hamphrey american english name.
	hamphreyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hamphrey"}

	// hamphreyCanadianFrenchName represents hamphrey canadian french name.
	hamphreyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Charles"}

	// hamphreyDutchName represents hamphrey dutch name.
	hamphreyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hamphrey"}

	// hamphreyFrenchName represents hamphrey french name.
	hamphreyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Charles"}

	// hamphreyGermanName represents hamphrey german name.
	hamphreyGermanName = nook.Name{
		Language: language.German,
		Value:    "Heinrich"}

	// hamphreyItalianName represents hamphrey italian name.
	hamphreyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nerone"}

	// hamphreyJapaneseName represents hamphrey japanese name.
	hamphreyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハムジ"}

	// hamphreyKoreanName represents hamphrey korean name.
	hamphreyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "햄쥐"}

	// hamphreyLatinAmericanSpanishName represents hamphrey latin american spanish name.
	hamphreyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Arsenio"}

	// hamphreyRussianName represents hamphrey russian name.
	hamphreyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хэмфри"}

	// hamphreySimplifiedChineseName represents hamphrey simplified chinese name.
	hamphreySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小仓"}

	// hamphreySpanishName represents hamphrey spanish name.
	hamphreySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Arsenio"}

	// hamphreyTraditionalChineseName represents hamphrey traditional chinese name.
	hamphreyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小倉"}
)

var (
	// hamphreyName represents hamphrey name.
	hamphreyName = nook.Languages{
		language.AmericanEnglish:      hamphreyAmericanEnglishName,
		language.CanadianFrench:       hamphreyCanadianFrenchName,
		language.Dutch:                hamphreyDutchName,
		language.French:               hamphreyFrenchName,
		language.German:               hamphreyGermanName,
		language.Italian:              hamphreyItalianName,
		language.Japanese:             hamphreyJapaneseName,
		language.Korean:               hamphreyKoreanName,
		language.LatinAmericanSpanish: hamphreyLatinAmericanSpanishName,
		language.Russian:              hamphreyRussianName,
		language.SimplifiedChinese:    hamphreySimplifiedChineseName,
		language.Spanish:              hamphreySpanishName,
		language.TraditionalChinese:   hamphreyTraditionalChineseName}
)

var (
	// hamphreyCharacter represents hamphrey character.
	hamphreyCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: hamphreyBirthday,
		Code:     hamphreyCode,
		Key:      character.Hamphrey,
		Gender:   gender.Male,
		Name:     hamphreyName,
		Special:  false}
)

var (
	// hamphreyAmericanEnglishPhrase represents hamphrey american english phrase.
	hamphreyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snort"}

	// hamphreyCanadianFrenchPhrase represents hamphrey canadian french phrase.
	hamphreyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gna gna"}

	// hamphreyDutchPhrase represents hamphrey dutch phrase.
	hamphreyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuifsnuif"}

	// hamphreyFrenchPhrase represents hamphrey french phrase.
	hamphreyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "balivernes"}

	// hamphreyGermanPhrase represents hamphrey german phrase.
	hamphreyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "squiep"}

	// hamphreyItalianPhrase represents hamphrey italian phrase.
	hamphreyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "spatapumf"}

	// hamphreyJapanesePhrase represents hamphrey japanese phrase.
	hamphreyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "カーッ"}

	// hamphreyKoreanPhrase represents hamphrey korean phrase.
	hamphreyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "카악"}

	// hamphreyLatinAmericanSpanishPhrase represents hamphrey latin american spanish phrase.
	hamphreyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "noquenó"}

	// hamphreyRussianPhrase represents hamphrey russian phrase.
	hamphreyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фырк"}

	// hamphreySimplifiedChinesePhrase represents hamphrey simplified chinese phrase.
	hamphreySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "怒"}

	// hamphreySpanishPhrase represents hamphrey spanish phrase.
	hamphreySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "andaquenó"}

	// hamphreyTraditionalChinesePhrase represents hamphrey traditional chinese phrase.
	hamphreyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "怒"}
)

var (
	// hamphreyPhrase represents hamphrey phrase.
	hamphreyPhrase = nook.Languages{
		language.AmericanEnglish:      hamphreyAmericanEnglishPhrase,
		language.CanadianFrench:       hamphreyCanadianFrenchPhrase,
		language.Dutch:                hamphreyDutchPhrase,
		language.French:               hamphreyFrenchPhrase,
		language.German:               hamphreyGermanPhrase,
		language.Italian:              hamphreyItalianPhrase,
		language.Japanese:             hamphreyJapanesePhrase,
		language.Korean:               hamphreyKoreanPhrase,
		language.LatinAmericanSpanish: hamphreyLatinAmericanSpanishPhrase,
		language.Russian:              hamphreyRussianPhrase,
		language.SimplifiedChinese:    hamphreySimplifiedChinesePhrase,
		language.Spanish:              hamphreySpanishPhrase,
		language.TraditionalChinese:   hamphreyTraditionalChinesePhrase}
)

var (
	// Hamphrey represents hamphrey.
	Hamphrey = nook.Villager{
		Character:   hamphreyCharacter,
		Personality: personality.Cranky,
		Phrase:      hamphreyPhrase}
)
