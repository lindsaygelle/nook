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
	// auroraBirthday represents aurora birthday.
	auroraBirthday = nook.Birthday{
		Day:   27,
		Month: time.January}
)

var (
	// auroraCode represents aurora code.
	auroraCode = nook.Code{
		Value: "pgn00"}
)

var (
	// auroraAmericanEnglishName represents aurora american english name.
	auroraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Aurora"}

	// auroraCanadianFrenchName represents aurora canadian french name.
	auroraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Aurore"}

	// auroraDutchName represents aurora dutch name.
	auroraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Aurora"}

	// auroraFrenchName represents aurora french name.
	auroraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Aurore"}

	// auroraGermanName represents aurora german name.
	auroraGermanName = nook.Name{
		Language: language.German,
		Value:    "Sonja"}

	// auroraItalianName represents aurora italian name.
	auroraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Aurora"}

	// auroraJapaneseName represents aurora japanese name.
	auroraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オーロラ"}

	// auroraKoreanName represents aurora korean name.
	auroraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "오로라"}

	// auroraLatinAmericanSpanishName represents aurora latin american spanish name.
	auroraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aurora"}

	// auroraRussianName represents aurora russian name.
	auroraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аврора"}

	// auroraSimplifiedChineseName represents aurora simplified chinese name.
	auroraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "欧若拉"}

	// auroraSpanishName represents aurora spanish name.
	auroraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aurora"}

	// auroraTraditionalChineseName represents aurora traditional chinese name.
	auroraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "歐若拉"}
)

var (
	// auroraName represents aurora name.
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
	// auroraCharacter represents aurora character.
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
	// auroraAmericanEnglishPhrase represents aurora american english phrase.
	auroraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "b-b-baby"}

	// auroraCanadianFrenchPhrase represents aurora canadian french phrase.
	auroraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "b-b-bébé"}

	// auroraDutchPhrase represents aurora dutch phrase.
	auroraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "b-b-baby"}

	// auroraFrenchPhrase represents aurora french phrase.
	auroraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "b-b-bébé"}

	// auroraGermanPhrase represents aurora german phrase.
	auroraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "b-bestens"}

	// auroraItalianPhrase represents aurora italian phrase.
	auroraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "b-b-baby"}

	// auroraJapanesePhrase represents aurora japanese phrase.
	auroraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だジョー"}

	// auroraKoreanPhrase represents aurora korean phrase.
	auroraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "랄라"}

	// auroraLatinAmericanSpanishPhrase represents aurora latin american spanish phrase.
	auroraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pescadí"}

	// auroraRussianPhrase represents aurora russian phrase.
	auroraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "к-к-крошка"}

	// auroraSimplifiedChinesePhrase represents aurora simplified chinese phrase.
	auroraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "若"}

	// auroraSpanishPhrase represents aurora spanish phrase.
	auroraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pescadito"}

	// auroraTraditionalChinesePhrase represents aurora traditional chinese phrase.
	auroraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "若"}
)

var (
	// auroraPhrase represents aurora phrase.
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
	// Aurora represents aurora.
	Aurora = nook.Villager{
		Character:   auroraCharacter,
		Personality: personality.Normal,
		Phrase:      auroraPhrase}
)
