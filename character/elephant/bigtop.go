package elephant

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
	// bigtopBirthday represents bigtop birthday.
	bigtopBirthday = nook.Birthday{
		Day:   3,
		Month: time.October}
)

var (
	// bigtopCode represents bigtop code.
	bigtopCode = nook.Code{
		Value: "elp02"}
)

var (
	// bigtopAmericanEnglishName represents bigtop american english name.
	bigtopAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Big Top"}

	// bigtopCanadianFrenchName represents bigtop canadian french name.
	bigtopCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Miles"}

	// bigtopDutchName represents bigtop dutch name.
	bigtopDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Big Top"}

	// bigtopFrenchName represents bigtop french name.
	bigtopFrenchName = nook.Name{
		Language: language.French,
		Value:    "Miles"}

	// bigtopGermanName represents bigtop german name.
	bigtopGermanName = nook.Name{
		Language: language.German,
		Value:    "Benni"}

	// bigtopItalianName represents bigtop italian name.
	bigtopItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grande C"}

	// bigtopJapaneseName represents bigtop japanese name.
	bigtopJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "３ごう"}

	// bigtopKoreanName represents bigtop korean name.
	bigtopKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "3호"}

	// bigtopLatinAmericanSpanishName represents bigtop latin american spanish name.
	bigtopLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Déivid"}

	// bigtopRussianName represents bigtop russian name.
	bigtopRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Биг-Топ"}

	// bigtopSimplifiedChineseName represents bigtop simplified chinese name.
	bigtopSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿三"}

	// bigtopSpanishName represents bigtop spanish name.
	bigtopSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Déivid"}

	// bigtopTraditionalChineseName represents bigtop traditional chinese name.
	bigtopTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿三"}
)

var (
	// bigtopName represents bigtop name.
	bigtopName = nook.Languages{
		language.AmericanEnglish:      bigtopAmericanEnglishName,
		language.CanadianFrench:       bigtopCanadianFrenchName,
		language.Dutch:                bigtopDutchName,
		language.French:               bigtopFrenchName,
		language.German:               bigtopGermanName,
		language.Italian:              bigtopItalianName,
		language.Japanese:             bigtopJapaneseName,
		language.Korean:               bigtopKoreanName,
		language.LatinAmericanSpanish: bigtopLatinAmericanSpanishName,
		language.Russian:              bigtopRussianName,
		language.SimplifiedChinese:    bigtopSimplifiedChineseName,
		language.Spanish:              bigtopSpanishName,
		language.TraditionalChinese:   bigtopTraditionalChineseName}
)

var (
	// bigtopCharacter represents bigtop character.
	bigtopCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: bigtopBirthday,
		Code:     bigtopCode,
		Key:      character.BigTop,
		Gender:   gender.Male,
		Name:     bigtopName,
		Special:  false}
)

var (
	// bigtopAmericanEnglishPhrase represents bigtop american english phrase.
	bigtopAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "villain"}

	// bigtopCanadianFrenchPhrase represents bigtop canadian french phrase.
	bigtopCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "vipère"}

	// bigtopDutchPhrase represents bigtop dutch phrase.
	bigtopDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boef"}

	// bigtopFrenchPhrase represents bigtop french phrase.
	bigtopFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "vipère"}

	// bigtopGermanPhrase represents bigtop german phrase.
	bigtopGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "truuu"}

	// bigtopItalianPhrase represents bigtop italian phrase.
	bigtopItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "furfante"}

	// bigtopJapanesePhrase represents bigtop japanese phrase.
	bigtopJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "うりゃー"}

	// bigtopKoreanPhrase represents bigtop korean phrase.
	bigtopKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "으랏차"}

	// bigtopLatinAmericanSpanishPhrase represents bigtop latin american spanish phrase.
	bigtopLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "duuud"}

	// bigtopRussianPhrase represents bigtop russian phrase.
	bigtopRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хитрюга"}

	// bigtopSimplifiedChinesePhrase represents bigtop simplified chinese phrase.
	bigtopSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈啊"}

	// bigtopSpanishPhrase represents bigtop spanish phrase.
	bigtopSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mutante"}

	// bigtopTraditionalChinesePhrase represents bigtop traditional chinese phrase.
	bigtopTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈啊"}
)

var (
	// bigtopPhrase represents bigtop phrase.
	bigtopPhrase = nook.Languages{
		language.AmericanEnglish:      bigtopAmericanEnglishPhrase,
		language.CanadianFrench:       bigtopCanadianFrenchPhrase,
		language.Dutch:                bigtopDutchPhrase,
		language.French:               bigtopFrenchPhrase,
		language.German:               bigtopGermanPhrase,
		language.Italian:              bigtopItalianPhrase,
		language.Japanese:             bigtopJapanesePhrase,
		language.Korean:               bigtopKoreanPhrase,
		language.LatinAmericanSpanish: bigtopLatinAmericanSpanishPhrase,
		language.Russian:              bigtopRussianPhrase,
		language.SimplifiedChinese:    bigtopSimplifiedChinesePhrase,
		language.Spanish:              bigtopSpanishPhrase,
		language.TraditionalChinese:   bigtopTraditionalChinesePhrase}
)

var (
	// BigTop represents big top.
	BigTop = nook.Villager{
		Character:   bigtopCharacter,
		Personality: personality.Lazy,
		Phrase:      bigtopPhrase}
)
