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
	// felyneBirthday represents felyne birthday.
	felyneBirthday = nook.Birthday{
		Day:   6,
		Month: time.January}
)

var (
	// felyneCode represents felyne code.
	felyneCode = nook.Code{
		Value: "cat22"}
)

var (
	// felyneAmericanEnglishName represents felyne american english name.
	felyneAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Felyne"}

	// felyneCanadianFrenchName represents felyne canadian french name.
	felyneCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Felyne"}

	// felyneDutchName represents felyne dutch name.
	felyneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// felyneFrenchName represents felyne french name.
	felyneFrenchName = nook.Name{
		Language: language.French,
		Value:    "Felyne"}

	// felyneGermanName represents felyne german name.
	felyneGermanName = nook.Name{
		Language: language.German,
		Value:    "Felyne"}

	// felyneItalianName represents felyne italian name.
	felyneItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Felyne"}

	// felyneJapaneseName represents felyne japanese name.
	felyneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイルー"}

	// felyneKoreanName represents felyne korean name.
	felyneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이루"}

	// felyneLatinAmericanSpanishName represents felyne latin american spanish name.
	felyneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Felyne"}

	// felyneRussianName represents felyne russian name.
	felyneRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// felyneSimplifiedChineseName represents felyne simplified chinese name.
	felyneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// felyneSpanishName represents felyne spanish name.
	felyneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Felyne"}

	// felyneTraditionalChineseName represents felyne traditional chinese name.
	felyneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// felyneName represents felyne name.
	felyneName = nook.Languages{
		language.AmericanEnglish:      felyneAmericanEnglishName,
		language.CanadianFrench:       felyneCanadianFrenchName,
		language.Dutch:                felyneDutchName,
		language.French:               felyneFrenchName,
		language.German:               felyneGermanName,
		language.Italian:              felyneItalianName,
		language.Japanese:             felyneJapaneseName,
		language.Korean:               felyneKoreanName,
		language.LatinAmericanSpanish: felyneLatinAmericanSpanishName,
		language.Russian:              felyneRussianName,
		language.SimplifiedChinese:    felyneSimplifiedChineseName,
		language.Spanish:              felyneSpanishName,
		language.TraditionalChinese:   felyneTraditionalChineseName}
)

var (
	// felyneCharacter represents felyne character.
	felyneCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: felyneBirthday,
		Code:     felyneCode,
		Key:      character.Felyne,
		Gender:   gender.Male,
		Name:     felyneName,
		Special:  false}
)

var (
	// felyneAmericanEnglishPhrase represents felyne american english phrase.
	felyneAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nya"}

	// felyneCanadianFrenchPhrase represents felyne canadian french phrase.
	felyneCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miaou"}

	// felyneDutchPhrase represents felyne dutch phrase.
	felyneDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// felyneFrenchPhrase represents felyne french phrase.
	felyneFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miaou"}

	// felyneGermanPhrase represents felyne german phrase.
	felyneGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "miau"}

	// felyneItalianPhrase represents felyne italian phrase.
	felyneItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mià"}

	// felyneJapanesePhrase represents felyne japanese phrase.
	felyneJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ニャ"}

	// felyneKoreanPhrase represents felyne korean phrase.
	felyneKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냥"}

	// felyneLatinAmericanSpanishPhrase represents felyne latin american spanish phrase.
	felyneLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "miaumo"}

	// felyneRussianPhrase represents felyne russian phrase.
	felyneRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// felyneSimplifiedChinesePhrase represents felyne simplified chinese phrase.
	felyneSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// felyneSpanishPhrase represents felyne spanish phrase.
	felyneSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "miaumo"}

	// felyneTraditionalChinesePhrase represents felyne traditional chinese phrase.
	felyneTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// felynePhrase represents felyne phrase.
	felynePhrase = nook.Languages{
		language.AmericanEnglish:      felyneAmericanEnglishPhrase,
		language.CanadianFrench:       felyneCanadianFrenchPhrase,
		language.Dutch:                felyneDutchPhrase,
		language.French:               felyneFrenchPhrase,
		language.German:               felyneGermanPhrase,
		language.Italian:              felyneItalianPhrase,
		language.Japanese:             felyneJapanesePhrase,
		language.Korean:               felyneKoreanPhrase,
		language.LatinAmericanSpanish: felyneLatinAmericanSpanishPhrase,
		language.Russian:              felyneRussianPhrase,
		language.SimplifiedChinese:    felyneSimplifiedChinesePhrase,
		language.Spanish:              felyneSpanishPhrase,
		language.TraditionalChinese:   felyneTraditionalChinesePhrase}
)

var (
	// Felyne represents felyne.
	Felyne = nook.Villager{
		Character:   felyneCharacter,
		Personality: personality.Lazy,
		Phrase:      felynePhrase}
)
