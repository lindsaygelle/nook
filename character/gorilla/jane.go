package gorilla

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
	// janeBirthday represents jane birthday.
	janeBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// janeCode represents jane code.
	janeCode = nook.Code{
		Value: ""}
)

var (
	// janeAmericanEnglishName represents jane american english name.
	janeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jane"}

	// janeCanadianFrenchName represents jane canadian french name.
	janeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// janeDutchName represents jane dutch name.
	janeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// janeFrenchName represents jane french name.
	janeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jane"}

	// janeGermanName represents jane german name.
	janeGermanName = nook.Name{
		Language: language.German,
		Value:    "Jane"}

	// janeItalianName represents jane italian name.
	janeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jane"}

	// janeJapaneseName represents jane japanese name.
	janeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フィーバー"}

	// janeKoreanName represents jane korean name.
	janeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// janeLatinAmericanSpanishName represents jane latin american spanish name.
	janeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// janeRussianName represents jane russian name.
	janeRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// janeSimplifiedChineseName represents jane simplified chinese name.
	janeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "白毛"}

	// janeSpanishName represents jane spanish name.
	janeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jane"}

	// janeTraditionalChineseName represents jane traditional chinese name.
	janeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// janeName represents jane name.
	janeName = nook.Languages{
		language.AmericanEnglish:      janeAmericanEnglishName,
		language.CanadianFrench:       janeCanadianFrenchName,
		language.Dutch:                janeDutchName,
		language.French:               janeFrenchName,
		language.German:               janeGermanName,
		language.Italian:              janeItalianName,
		language.Japanese:             janeJapaneseName,
		language.Korean:               janeKoreanName,
		language.LatinAmericanSpanish: janeLatinAmericanSpanishName,
		language.Russian:              janeRussianName,
		language.SimplifiedChinese:    janeSimplifiedChineseName,
		language.Spanish:              janeSpanishName,
		language.TraditionalChinese:   janeTraditionalChineseName}
)

var (
	// janeCharacter represents jane character.
	janeCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: janeBirthday,
		Code:     janeCode,
		Key:      character.Jane,
		Gender:   gender.Female,
		Name:     janeName,
		Special:  false}
)

var (
	// janeAmericanEnglishPhrase represents jane american english phrase.
	janeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chimp"}

	// janeCanadianFrenchPhrase represents jane canadian french phrase.
	janeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// janeDutchPhrase represents jane dutch phrase.
	janeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// janeFrenchPhrase represents jane french phrase.
	janeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ouistiti"}

	// janeGermanPhrase represents jane german phrase.
	janeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "freund"}

	// janeItalianPhrase represents jane italian phrase.
	janeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "scimmietta"}

	// janeJapanesePhrase represents jane japanese phrase.
	janeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だゴリ"}

	// janeKoreanPhrase represents jane korean phrase.
	janeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// janeLatinAmericanSpanishPhrase represents jane latin american spanish phrase.
	janeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// janeRussianPhrase represents jane russian phrase.
	janeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// janeSimplifiedChinesePhrase represents jane simplified chinese phrase.
	janeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘣嘣"}

	// janeSpanishPhrase represents jane spanish phrase.
	janeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chimpa"}

	// janeTraditionalChinesePhrase represents jane traditional chinese phrase.
	janeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// janePhrase represents jane phrase.
	janePhrase = nook.Languages{
		language.AmericanEnglish:      janeAmericanEnglishPhrase,
		language.CanadianFrench:       janeCanadianFrenchPhrase,
		language.Dutch:                janeDutchPhrase,
		language.French:               janeFrenchPhrase,
		language.German:               janeGermanPhrase,
		language.Italian:              janeItalianPhrase,
		language.Japanese:             janeJapanesePhrase,
		language.Korean:               janeKoreanPhrase,
		language.LatinAmericanSpanish: janeLatinAmericanSpanishPhrase,
		language.Russian:              janeRussianPhrase,
		language.SimplifiedChinese:    janeSimplifiedChinesePhrase,
		language.Spanish:              janeSpanishPhrase,
		language.TraditionalChinese:   janeTraditionalChinesePhrase}
)

var (
	// Jane represents jane.
	Jane = nook.Villager{
		Character:   janeCharacter,
		Personality: personality.Snooty,
		Phrase:      janePhrase}
)
