package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	janeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	janeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	janeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	janeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "白毛"}

	janeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jane"}

	janeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Gorilla,
		Birthday: janeBirthday,
		Code:     janeCode,
		Gender:   nook.Female,
		Name:     janeName}
)

var (
	janeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	janeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	janeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	janeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	janeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	janeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘣嘣"}

	janeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chimpa"}

	janeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	janePhrase = nook.Languages{
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
	Jane = nook.Villager{
		Character:   janeCharacter,
		Personality: nook.Snooty,
		Phrase:      janePhrase}
)
