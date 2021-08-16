package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	joeBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	joeCode = nook.Code{
		Value: ""}
)

var (
	joeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Joe"}

	joeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	joeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	joeFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	joeGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	joeItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	joeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジョー"}

	joeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	joeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	joeRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	joeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	joeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	joeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	joeName = nook.Languages{
		language.AmericanEnglish:      joeAmericanEnglishName,
		language.CanadianFrench:       joeCanadianFrenchName,
		language.Dutch:                joeDutchName,
		language.French:               joeFrenchName,
		language.German:               joeGermanName,
		language.Italian:              joeItalianName,
		language.Japanese:             joeJapaneseName,
		language.Korean:               joeKoreanName,
		language.LatinAmericanSpanish: joeLatinAmericanSpanishName,
		language.Russian:              joeRussianName,
		language.SimplifiedChinese:    joeSimplifiedChineseName,
		language.Spanish:              joeSpanishName,
		language.TraditionalChinese:   joeTraditionalChineseName}
)

var (
	joeCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: joeBirthday,
		Code:     joeCode,
		Gender:   gender.Male,
		Name:     joeName}
)

var (
	joeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "やれやれ"}

	joeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	joeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	joeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	joeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	joeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	joeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やれやれ"}

	joeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	joeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	joeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	joeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	joeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	joeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	joePhrase = nook.Languages{
		language.AmericanEnglish:      joeAmericanEnglishName,
		language.CanadianFrench:       joeCanadianFrenchName,
		language.Dutch:                joeDutchName,
		language.French:               joeFrenchName,
		language.German:               joeGermanName,
		language.Italian:              joeItalianName,
		language.Japanese:             joeJapaneseName,
		language.Korean:               joeKoreanName,
		language.LatinAmericanSpanish: joeLatinAmericanSpanishName,
		language.Russian:              joeRussianName,
		language.SimplifiedChinese:    joeSimplifiedChineseName,
		language.Spanish:              joeSpanishName,
		language.TraditionalChinese:   joeTraditionalChineseName}
)

var (
	Joe = nook.Villager{
		Character:   joeCharacter,
		Personality: personality.Cranky,
		Phrase:      joePhrase}
)
