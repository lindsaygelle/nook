package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	felyneBirthday = nook.Birthday{
		Day:   6,
		Month: time.January}
)

var (
	felyneCode = nook.Code{
		Value: "cat22"}
)

var (
	felyneAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Felyne"}

	felyneCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Felyne"}

	felyneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	felyneFrenchName = nook.Name{
		Language: language.French,
		Value:    "Felyne"}

	felyneGermanName = nook.Name{
		Language: language.German,
		Value:    "Felyne"}

	felyneItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Felyne"}

	felyneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイルー"}

	felyneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이루"}

	felyneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Felyne"}

	felyneRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	felyneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	felyneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Felyne"}

	felyneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
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
	felyneCharacter = nook.Character{
		Animal:   Cat,
		Birthday: felyneBirthday,
		Code:     felyneCode,
		Gender:   nook.Male,
		Name:     felyneName}
)

var (
	felyneAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nya"}

	felyneCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miaou"}

	felyneDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	felyneFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miaou"}

	felyneGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "miau"}

	felyneItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mià"}

	felyneJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ニャ"}

	felyneKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냥"}

	felyneLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "miaumo"}

	felyneRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	felyneSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	felyneSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "miaumo"}

	felyneTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	felynePhrase = nook.Languages{
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
	Felyne = nook.Villager{
		Character:   felyneCharacter,
		Personality: nook.Lazy,
		Phrase:      felynePhrase}
)
