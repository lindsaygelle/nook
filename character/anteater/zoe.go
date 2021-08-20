package anteater

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
		Value:    ""}

	zoeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

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
		Value:    ""}

	zoeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	zoeRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	zoeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蔓蔓"}

	zoeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Naricia"}

	zoeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
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
		Animal:   animal.Anteater,
		Birthday: zoeBirthday,
		Code:     zoeCode,
		Key:      character.Zoe,
		Gender:   gender.Female,
		Name:     zoeName,
		Special:  false}
)

var (
	zoeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "whiiifff"}

	zoeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	zoeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

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
		Value:    ""}

	zoeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	zoeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	zoeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "没错"}

	zoeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fufuf"}

	zoeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	zoePhrase = nook.Languages{
		language.AmericanEnglish:      zoeAmericanEnglishPhrase,
		language.CanadianFrench:       zoeCanadianFrenchPhrase,
		language.Dutch:                zoeDutchPhrase,
		language.French:               zoeFrenchPhrase,
		language.German:               zoeGermanPhrase,
		language.Italian:              zoeItalianPhrase,
		language.Japanese:             zoeJapanesePhrase,
		language.Korean:               zoeKoreanPhrase,
		language.LatinAmericanSpanish: zoeLatinAmericanSpanishPhrase,
		language.Russian:              zoeRussianPhrase,
		language.SimplifiedChinese:    zoeSimplifiedChinesePhrase,
		language.Spanish:              zoeSpanishPhrase,
		language.TraditionalChinese:   zoeTraditionalChinesePhrase}
)

var (
	Zoe = nook.Villager{
		Character:   zoeCharacter,
		Personality: personality.Normal,
		Phrase:      zoePhrase}
)
