package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	auroraBirthday = nook.Birthday{
		Day:   27,
		Month: time.January}
)

var (
	auroraCode = nook.Code{
		Value: "pgn00"}
)

var (
	auroraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Aurora"}

	auroraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Auroreb-b-bébé"}

	auroraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Aurorab-b-baby"}

	auroraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Auroreb-b-bébé"}

	auroraGermanName = nook.Name{
		Language: language.German,
		Value:    "Sonjab-bestens"}

	auroraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Aurorab-b-baby"}

	auroraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オーロラだジョー"}

	auroraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "오로라랄라"}

	auroraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aurorapescadí"}

	auroraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аврорак-к-крошка"}

	auroraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "欧若拉若"}

	auroraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aurorapescadito"}

	auroraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "歐若拉若"}
)

var (
	auroraName = nook.Languages{
		language.AmericanEnglish:      auroraAmericanEnglishName,
		language.CanadianFrench:       auroraCanadianFrenchName,
		language.Dutch:                auroraDutchName,
		language.French:               auroraFrenchName,
		language.German:               auroraGermanName,
		language.Italian:              auroraItalianName,
		language.Japanese:             auroraJapaneseName,
		language.Korean:               auroraKoreanName,
		language.LatinAmericanSpanish: auroraLatinAmericanSpanishName,
		language.Russian:              auroraRussianName,
		language.SimplifiedChinese:    auroraSimplifiedChineseName,
		language.Spanish:              auroraSpanishName,
		language.TraditionalChinese:   auroraTraditionalChineseName}
)

var (
	auroraCharacter = nook.Character{
		Animal:   Penguin,
		Birthday: auroraBirthday,
		Code:     auroraCode,
		Gender:   nook.Female,
		Name:     auroraName}
)

var (
	auroraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "b-b-baby"}

	auroraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "b-b-bébé"}

	auroraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "b-b-baby"}

	auroraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "b-b-bébé"}

	auroraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "b-bestens"}

	auroraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "b-b-baby"}

	auroraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だジョー"}

	auroraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "랄라"}

	auroraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pescadí"}

	auroraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "к-к-крошка"}

	auroraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "若"}

	auroraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pescadito"}

	auroraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "若"}
)

var (
	auroraPhrase = nook.Languages{
		language.AmericanEnglish:      auroraAmericanEnglishName,
		language.CanadianFrench:       auroraCanadianFrenchName,
		language.Dutch:                auroraDutchName,
		language.French:               auroraFrenchName,
		language.German:               auroraGermanName,
		language.Italian:              auroraItalianName,
		language.Japanese:             auroraJapaneseName,
		language.Korean:               auroraKoreanName,
		language.LatinAmericanSpanish: auroraLatinAmericanSpanishName,
		language.Russian:              auroraRussianName,
		language.SimplifiedChinese:    auroraSimplifiedChineseName,
		language.Spanish:              auroraSpanishName,
		language.TraditionalChinese:   auroraTraditionalChineseName}
)

var (
	Aurora = nook.Villager{
		Character:   auroraCharacter,
		Personality: nook.Normal,
		Phrase:      auroraPhrase}
)
