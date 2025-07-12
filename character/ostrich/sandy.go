package ostrich

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
	// sandyBirthday represents sandy birthday.
	sandyBirthday = nook.Birthday{
		Day:   21,
		Month: time.October}
)

var (
	// sandyCode represents sandy code.
	sandyCode = nook.Code{
		Value: "ost02"}
)

var (
	// sandyAmericanEnglishName represents sandy american english name.
	sandyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sandy"}

	// sandyCanadianFrenchName represents sandy canadian french name.
	sandyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ottie"}

	// sandyDutchName represents sandy dutch name.
	sandyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sandy"}

	// sandyFrenchName represents sandy french name.
	sandyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ottie"}

	// sandyGermanName represents sandy german name.
	sandyGermanName = nook.Name{
		Language: language.German,
		Value:    "Senta"}

	// sandyItalianName represents sandy italian name.
	sandyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Alessia"}

	// sandyJapaneseName represents sandy japanese name.
	sandyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラン"}

	// sandyKoreanName represents sandy korean name.
	sandyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "샌디"}

	// sandyLatinAmericanSpanishName represents sandy latin american spanish name.
	sandyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Priscila"}

	// sandyRussianName represents sandy russian name.
	sandyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сэнди"}

	// sandySimplifiedChineseName represents sandy simplified chinese name.
	sandySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小岚"}

	// sandySpanishName represents sandy spanish name.
	sandySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Priscila"}

	// sandyTraditionalChineseName represents sandy traditional chinese name.
	sandyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小嵐"}
)

var (
	// sandyName represents sandy name.
	sandyName = nook.Languages{
		language.AmericanEnglish:      sandyAmericanEnglishName,
		language.CanadianFrench:       sandyCanadianFrenchName,
		language.Dutch:                sandyDutchName,
		language.French:               sandyFrenchName,
		language.German:               sandyGermanName,
		language.Italian:              sandyItalianName,
		language.Japanese:             sandyJapaneseName,
		language.Korean:               sandyKoreanName,
		language.LatinAmericanSpanish: sandyLatinAmericanSpanishName,
		language.Russian:              sandyRussianName,
		language.SimplifiedChinese:    sandySimplifiedChineseName,
		language.Spanish:              sandySpanishName,
		language.TraditionalChinese:   sandyTraditionalChineseName}
)

var (
	// sandyCharacter represents sandy character.
	sandyCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: sandyBirthday,
		Code:     sandyCode,
		Key:      character.Sandy,
		Gender:   gender.Female,
		Name:     sandyName,
		Special:  false}
)

var (
	// sandyAmericanEnglishPhrase represents sandy american english phrase.
	sandyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "speedy"}

	// sandyCanadianFrenchPhrase represents sandy canadian french phrase.
	sandyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "euh quoi"}

	// sandyDutchPhrase represents sandy dutch phrase.
	sandyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "raprap"}

	// sandyFrenchPhrase represents sandy french phrase.
	sandyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "euh quoi"}

	// sandyGermanPhrase represents sandy german phrase.
	sandyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flottilda"}

	// sandyItalianPhrase represents sandy italian phrase.
	sandyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "vruum"}

	// sandyJapanesePhrase represents sandy japanese phrase.
	sandyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だチョー"}

	// sandyKoreanPhrase represents sandy korean phrase.
	sandyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그렇다조"}

	// sandyLatinAmericanSpanishPhrase represents sandy latin american spanish phrase.
	sandyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fiuuun"}

	// sandyRussianPhrase represents sandy russian phrase.
	sandyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "быстро"}

	// sandySimplifiedChinesePhrase represents sandy simplified chinese phrase.
	sandySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鸵鸟"}

	// sandySpanishPhrase represents sandy spanish phrase.
	sandySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fiuuun"}

	// sandyTraditionalChinesePhrase represents sandy traditional chinese phrase.
	sandyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鴕鳥"}
)

var (
	// sandyPhrase represents sandy phrase.
	sandyPhrase = nook.Languages{
		language.AmericanEnglish:      sandyAmericanEnglishPhrase,
		language.CanadianFrench:       sandyCanadianFrenchPhrase,
		language.Dutch:                sandyDutchPhrase,
		language.French:               sandyFrenchPhrase,
		language.German:               sandyGermanPhrase,
		language.Italian:              sandyItalianPhrase,
		language.Japanese:             sandyJapanesePhrase,
		language.Korean:               sandyKoreanPhrase,
		language.LatinAmericanSpanish: sandyLatinAmericanSpanishPhrase,
		language.Russian:              sandyRussianPhrase,
		language.SimplifiedChinese:    sandySimplifiedChinesePhrase,
		language.Spanish:              sandySpanishPhrase,
		language.TraditionalChinese:   sandyTraditionalChinesePhrase}
)

var (
	// Sandy represents sandy.
	Sandy = nook.Villager{
		Character:   sandyCharacter,
		Personality: personality.Normal,
		Phrase:      sandyPhrase}
)
