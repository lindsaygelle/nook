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
	janeBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	janeCode = nook.Code{
		Value: ""}
)

var (
	janeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jane"}

	janeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	janeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	janeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jane"}

	janeGermanName = nook.Name{
		Language: language.German,
		Value:    "Jane"}

	janeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Jane"}

	janeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フィーバー"}

	janeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	janeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	janeRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	janeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "白毛"}

	janeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jane"}

	janeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	janeCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: janeBirthday,
		Code:     janeCode,
		Key:      character.Jane,
		Gender:   gender.Female,
		Name:     janeName}
)

var (
	janeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chimp"}

	janeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	janeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	janeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ouistiti"}

	janeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "freund"}

	janeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "scimmietta"}

	janeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だゴリ"}

	janeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	janeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	janeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	janeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘣嘣"}

	janeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chimpa"}

	janeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	Jane = nook.Villager{
		Character:   janeCharacter,
		Personality: personality.Snooty,
		Phrase:      janePhrase}
)
