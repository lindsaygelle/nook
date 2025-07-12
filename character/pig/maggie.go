package pig

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
	// maggieBirthday represents maggie birthday.
	maggieBirthday = nook.Birthday{
		Day:   3,
		Month: time.September}
)

var (
	// maggieCode represents maggie code.
	maggieCode = nook.Code{
		Value: "pig10"}
)

var (
	// maggieAmericanEnglishName represents maggie american english name.
	maggieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Maggie"}

	// maggieCanadianFrenchName represents maggie canadian french name.
	maggieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marjorie"}

	// maggieDutchName represents maggie dutch name.
	maggieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Maggie"}

	// maggieFrenchName represents maggie french name.
	maggieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marjorie"}

	// maggieGermanName represents maggie german name.
	maggieGermanName = nook.Name{
		Language: language.German,
		Value:    "Magda"}

	// maggieItalianName represents maggie italian name.
	maggieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marcella"}

	// maggieJapaneseName represents maggie japanese name.
	maggieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マーガレット"}

	// maggieKoreanName represents maggie korean name.
	maggieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마가렛"}

	// maggieLatinAmericanSpanishName represents maggie latin american spanish name.
	maggieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Margarí"}

	// maggieRussianName represents maggie russian name.
	maggieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мэгги"}

	// maggieSimplifiedChineseName represents maggie simplified chinese name.
	maggieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玛格"}

	// maggieSpanishName represents maggie spanish name.
	maggieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Margarí"}

	// maggieTraditionalChineseName represents maggie traditional chinese name.
	maggieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑪格"}
)

var (
	// maggieName represents maggie name.
	maggieName = nook.Languages{
		language.AmericanEnglish:      maggieAmericanEnglishName,
		language.CanadianFrench:       maggieCanadianFrenchName,
		language.Dutch:                maggieDutchName,
		language.French:               maggieFrenchName,
		language.German:               maggieGermanName,
		language.Italian:              maggieItalianName,
		language.Japanese:             maggieJapaneseName,
		language.Korean:               maggieKoreanName,
		language.LatinAmericanSpanish: maggieLatinAmericanSpanishName,
		language.Russian:              maggieRussianName,
		language.SimplifiedChinese:    maggieSimplifiedChineseName,
		language.Spanish:              maggieSpanishName,
		language.TraditionalChinese:   maggieTraditionalChineseName}
)

var (
	// maggieCharacter represents maggie character.
	maggieCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: maggieBirthday,
		Code:     maggieCode,
		Key:      character.Maggie,
		Gender:   gender.Female,
		Name:     maggieName,
		Special:  false}
)

var (
	// maggieAmericanEnglishPhrase represents maggie american english phrase.
	maggieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "schep"}

	// maggieCanadianFrenchPhrase represents maggie canadian french phrase.
	maggieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grouik"}

	// maggieDutchPhrase represents maggie dutch phrase.
	maggieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knor-knor"}

	// maggieFrenchPhrase represents maggie french phrase.
	maggieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grouik"}

	// maggieGermanPhrase represents maggie german phrase.
	maggieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kringel"}

	// maggieItalianPhrase represents maggie italian phrase.
	maggieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sissì"}

	// maggieJapanesePhrase represents maggie japanese phrase.
	maggieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "うん"}

	// maggieKoreanPhrase represents maggie korean phrase.
	maggieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "유후"}

	// maggieLatinAmericanSpanishPhrase represents maggie latin american spanish phrase.
	maggieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oinki"}

	// maggieRussianPhrase represents maggie russian phrase.
	maggieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюля"}

	// maggieSimplifiedChinesePhrase represents maggie simplified chinese phrase.
	maggieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯"}

	// maggieSpanishPhrase represents maggie spanish phrase.
	maggieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "girasol"}

	// maggieTraditionalChinesePhrase represents maggie traditional chinese phrase.
	maggieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯"}
)

var (
	// maggiePhrase represents maggie phrase.
	maggiePhrase = nook.Languages{
		language.AmericanEnglish:      maggieAmericanEnglishPhrase,
		language.CanadianFrench:       maggieCanadianFrenchPhrase,
		language.Dutch:                maggieDutchPhrase,
		language.French:               maggieFrenchPhrase,
		language.German:               maggieGermanPhrase,
		language.Italian:              maggieItalianPhrase,
		language.Japanese:             maggieJapanesePhrase,
		language.Korean:               maggieKoreanPhrase,
		language.LatinAmericanSpanish: maggieLatinAmericanSpanishPhrase,
		language.Russian:              maggieRussianPhrase,
		language.SimplifiedChinese:    maggieSimplifiedChinesePhrase,
		language.Spanish:              maggieSpanishPhrase,
		language.TraditionalChinese:   maggieTraditionalChinesePhrase}
)

var (
	// Maggie represents maggie.
	Maggie = nook.Villager{
		Character:   maggieCharacter,
		Personality: personality.Normal,
		Phrase:      maggiePhrase}
)
