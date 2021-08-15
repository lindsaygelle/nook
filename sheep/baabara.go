package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Bêêttychêêêri"}

	baabaraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Baabaraschaaaapje"}

	baabaraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bêêttychêêêri"}

	baabaraGermanName = nook.Name{
		Language: language.German,
		Value:    "Babsidäääää"}

	baabaraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Beabreeezza"}

	baabaraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トロワアンドゥ"}

	baabaraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "트로와앤쥬"}

	baabaraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Beelénpeluuusas"}

	baabaraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бабарамила-ашка"}

	baabaraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "华尔滋一二三"}

	baabaraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Beelénpeeelusa"}

	baabaraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "華爾滋一二三"}
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
		Animal:   Sheep,
		Birthday: baabaraBirthday,
		Code:     baabaraCode,
		Gender:   nook.Female,
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
		Personality: nook.Snooty,
		Phrase:      baabaraPhrase}
)
