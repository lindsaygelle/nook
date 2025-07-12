package cat

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
	// punchyBirthday represents punchy birthday.
	punchyBirthday = nook.Birthday{
		Day:   11,
		Month: time.April}
)

var (
	// punchyCode represents punchy code.
	punchyCode = nook.Code{
		Value: "cat06"}
)

var (
	// punchyAmericanEnglishName represents punchy american english name.
	punchyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Punchy"}

	// punchyCanadianFrenchName represents punchy canadian french name.
	punchyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cédric"}

	// punchyDutchName represents punchy dutch name.
	punchyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Punchy"}

	// punchyFrenchName represents punchy french name.
	punchyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cédric"}

	// punchyGermanName represents punchy german name.
	punchyGermanName = nook.Name{
		Language: language.German,
		Value:    "Julian"}

	// punchyItalianName represents punchy italian name.
	punchyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Felix"}

	// punchyJapaneseName represents punchy japanese name.
	punchyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビンタ"}

	// punchyKoreanName represents punchy korean name.
	punchyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "빙티"}

	// punchyLatinAmericanSpanishName represents punchy latin american spanish name.
	punchyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Félix"}

	// punchyRussianName represents punchy russian name.
	punchyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Панчи"}

	// punchySimplifiedChineseName represents punchy simplified chinese name.
	punchySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "尔光"}

	// punchySpanishName represents punchy spanish name.
	punchySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Félix"}

	// punchyTraditionalChineseName represents punchy traditional chinese name.
	punchyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "爾光"}
)

var (
	// punchyName represents punchy name.
	punchyName = nook.Languages{
		language.AmericanEnglish:      punchyAmericanEnglishName,
		language.CanadianFrench:       punchyCanadianFrenchName,
		language.Dutch:                punchyDutchName,
		language.French:               punchyFrenchName,
		language.German:               punchyGermanName,
		language.Italian:              punchyItalianName,
		language.Japanese:             punchyJapaneseName,
		language.Korean:               punchyKoreanName,
		language.LatinAmericanSpanish: punchyLatinAmericanSpanishName,
		language.Russian:              punchyRussianName,
		language.SimplifiedChinese:    punchySimplifiedChineseName,
		language.Spanish:              punchySpanishName,
		language.TraditionalChinese:   punchyTraditionalChineseName}
)

var (
	// punchyCharacter represents punchy character.
	punchyCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: punchyBirthday,
		Code:     punchyCode,
		Key:      character.Punchy,
		Gender:   gender.Male,
		Name:     punchyName,
		Special:  false}
)

var (
	// punchyAmericanEnglishPhrase represents punchy american english phrase.
	punchyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mrmpht"}

	// punchyCanadianFrenchPhrase represents punchy canadian french phrase.
	punchyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mimine"}

	// punchyDutchPhrase represents punchy dutch phrase.
	punchyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "plof"}

	// punchyFrenchPhrase represents punchy french phrase.
	punchyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "blébléblé"}

	// punchyGermanPhrase represents punchy german phrase.
	punchyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mrmpft"}

	// punchyItalianPhrase represents punchy italian phrase.
	punchyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "maramao"}

	// punchyJapanesePhrase represents punchy japanese phrase.
	punchyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だのら"}

	// punchyKoreanPhrase represents punchy korean phrase.
	punchyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "노라줘"}

	// punchyLatinAmericanSpanishPhrase represents punchy latin american spanish phrase.
	punchyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gruuuah"}

	// punchyRussianPhrase represents punchy russian phrase.
	punchyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "отдыхаем"}

	// punchySimplifiedChinesePhrase represents punchy simplified chinese phrase.
	punchySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "晃晃"}

	// punchySpanishPhrase represents punchy spanish phrase.
	punchySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "arenque"}

	// punchyTraditionalChinesePhrase represents punchy traditional chinese phrase.
	punchyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "晃晃"}
)

var (
	// punchyPhrase represents punchy phrase.
	punchyPhrase = nook.Languages{
		language.AmericanEnglish:      punchyAmericanEnglishPhrase,
		language.CanadianFrench:       punchyCanadianFrenchPhrase,
		language.Dutch:                punchyDutchPhrase,
		language.French:               punchyFrenchPhrase,
		language.German:               punchyGermanPhrase,
		language.Italian:              punchyItalianPhrase,
		language.Japanese:             punchyJapanesePhrase,
		language.Korean:               punchyKoreanPhrase,
		language.LatinAmericanSpanish: punchyLatinAmericanSpanishPhrase,
		language.Russian:              punchyRussianPhrase,
		language.SimplifiedChinese:    punchySimplifiedChinesePhrase,
		language.Spanish:              punchySpanishPhrase,
		language.TraditionalChinese:   punchyTraditionalChinesePhrase}
)

var (
	// Punchy represents punchy.
	Punchy = nook.Villager{
		Character:   punchyCharacter,
		Personality: personality.Lazy,
		Phrase:      punchyPhrase}
)
