package penguin

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
		Value:    "Aurore"}

	auroraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Aurora"}

	auroraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Aurore"}

	auroraGermanName = nook.Name{
		Language: language.German,
		Value:    "Sonja"}

	auroraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Aurora"}

	auroraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オーロラ"}

	auroraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "오로라"}

	auroraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aurora"}

	auroraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аврора"}

	auroraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "欧若拉"}

	auroraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aurora"}

	auroraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "歐若拉"}
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
		Animal:   animal.Penguin,
		Birthday: auroraBirthday,
		Code:     auroraCode,
		Key:      character.Aurora,
		Gender:   gender.Female,
		Name:     auroraName,
		Special:  false}
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
		language.AmericanEnglish:      auroraAmericanEnglishPhrase,
		language.CanadianFrench:       auroraCanadianFrenchPhrase,
		language.Dutch:                auroraDutchPhrase,
		language.French:               auroraFrenchPhrase,
		language.German:               auroraGermanPhrase,
		language.Italian:              auroraItalianPhrase,
		language.Japanese:             auroraJapanesePhrase,
		language.Korean:               auroraKoreanPhrase,
		language.LatinAmericanSpanish: auroraLatinAmericanSpanishPhrase,
		language.Russian:              auroraRussianPhrase,
		language.SimplifiedChinese:    auroraSimplifiedChinesePhrase,
		language.Spanish:              auroraSpanishPhrase,
		language.TraditionalChinese:   auroraTraditionalChinesePhrase}
)

var (
	Aurora = nook.Villager{
		Character:   auroraCharacter,
		Personality: personality.Normal,
		Phrase:      auroraPhrase}
)
