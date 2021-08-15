package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	zoeBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	zoeCode = nook.Code{
		Value: ""}
)

var (
	zoeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Zoe"}

	zoeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	zoeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	zoeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Zoé"}

	zoeGermanName = nook.Name{
		Language: language.German,
		Value:    "Zoe"}

	zoeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zoe"}

	zoeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビーフン"}

	zoeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	zoeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	zoeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	zoeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蔓蔓"}

	zoeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Naricia"}

	zoeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	zoeName = nook.Languages{
		language.AmericanEnglish:      zoeAmericanEnglishName,
		language.CanadianFrench:       zoeCanadianFrenchName,
		language.Dutch:                zoeDutchName,
		language.French:               zoeFrenchName,
		language.German:               zoeGermanName,
		language.Italian:              zoeItalianName,
		language.Japanese:             zoeJapaneseName,
		language.Korean:               zoeKoreanName,
		language.LatinAmericanSpanish: zoeLatinAmericanSpanishName,
		language.Russian:              zoeRussianName,
		language.SimplifiedChinese:    zoeSimplifiedChineseName,
		language.Spanish:              zoeSpanishName,
		language.TraditionalChinese:   zoeTraditionalChineseName}
)

var (
	zoeCharacter = nook.Character{
		Animal:   Anteater,
		Birthday: zoeBirthday,
		Code:     zoeCode,
		Gender:   nook.Female,
		Name:     zoeName}
)

var (
	zoeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	zoeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	zoeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	zoeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "fouf fouf"}

	zoeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnieef"}

	zoeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oiiinfff"}

	zoeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だモン"}

	zoeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	zoeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	zoeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	zoeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "没错"}

	zoeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fufuf"}

	zoeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	zoePhrase = nook.Languages{
		language.AmericanEnglish:      zoeAmericanEnglishName,
		language.CanadianFrench:       zoeCanadianFrenchName,
		language.Dutch:                zoeDutchName,
		language.French:               zoeFrenchName,
		language.German:               zoeGermanName,
		language.Italian:              zoeItalianName,
		language.Japanese:             zoeJapaneseName,
		language.Korean:               zoeKoreanName,
		language.LatinAmericanSpanish: zoeLatinAmericanSpanishName,
		language.Russian:              zoeRussianName,
		language.SimplifiedChinese:    zoeSimplifiedChineseName,
		language.Spanish:              zoeSpanishName,
		language.TraditionalChinese:   zoeTraditionalChineseName}
)

var (
	Zoe = nook.Villager{
		Character:   zoeCharacter,
		Personality: nook.Normal,
		Phrase:      zoePhrase}
)
