package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	baabaraBirthday = nook.Birthday{
		Day:   28,
		Month: time.March}
)

var (
	baabaraCode = nook.Code{
		Value: "shp01"}
)

var (
	baabaraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Baabara"}

	baabaraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bêêtty"}

	baabaraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Baabara"}

	baabaraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bêêtty"}

	baabaraGermanName = nook.Name{
		Language: language.German,
		Value:    "Babsi"}

	baabaraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bea"}

	baabaraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トロワ"}

	baabaraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "트로와"}

	baabaraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Beelén"}

	baabaraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бабара"}

	baabaraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "华尔滋"}

	baabaraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Beelén"}

	baabaraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "華爾滋"}
)

var (
	baabaraName = nook.Languages{
		language.AmericanEnglish:      baabaraAmericanEnglishName,
		language.CanadianFrench:       baabaraCanadianFrenchName,
		language.Dutch:                baabaraDutchName,
		language.French:               baabaraFrenchName,
		language.German:               baabaraGermanName,
		language.Italian:              baabaraItalianName,
		language.Japanese:             baabaraJapaneseName,
		language.Korean:               baabaraKoreanName,
		language.LatinAmericanSpanish: baabaraLatinAmericanSpanishName,
		language.Russian:              baabaraRussianName,
		language.SimplifiedChinese:    baabaraSimplifiedChineseName,
		language.Spanish:              baabaraSpanishName,
		language.TraditionalChinese:   baabaraTraditionalChineseName}
)

var (
	baabaraCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: baabaraBirthday,
		Code:     baabaraCode,
		Gender:   gender.Female,
		Name:     baabaraName}
)

var (
	baabaraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "daahling"}

	baabaraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chêêêri"}

	baabaraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schaaaapje"}

	baabaraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chêêêri"}

	baabaraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "däääää"}

	baabaraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "breeezza"}

	baabaraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アンドゥ"}

	baabaraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "앤쥬"}

	baabaraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "peluuusas"}

	baabaraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мила-ашка"}

	baabaraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "一二三"}

	baabaraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "peeelusa"}

	baabaraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "一二三"}
)

var (
	baabaraPhrase = nook.Languages{
		language.AmericanEnglish:      baabaraAmericanEnglishName,
		language.CanadianFrench:       baabaraCanadianFrenchName,
		language.Dutch:                baabaraDutchName,
		language.French:               baabaraFrenchName,
		language.German:               baabaraGermanName,
		language.Italian:              baabaraItalianName,
		language.Japanese:             baabaraJapaneseName,
		language.Korean:               baabaraKoreanName,
		language.LatinAmericanSpanish: baabaraLatinAmericanSpanishName,
		language.Russian:              baabaraRussianName,
		language.SimplifiedChinese:    baabaraSimplifiedChineseName,
		language.Spanish:              baabaraSpanishName,
		language.TraditionalChinese:   baabaraTraditionalChineseName}
)

var (
	Baabara = nook.Villager{
		Character:   baabaraCharacter,
		Personality: personality.Snooty,
		Phrase:      baabaraPhrase}
)
