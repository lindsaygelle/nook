package bear

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
	pinkyBirthday = nook.Birthday{
		Day:   9,
		Month: time.September}
)

var (
	pinkyCode = nook.Code{
		Value: "bea01"}
)

var (
	pinkyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pinky"}

	pinkyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rosine"}

	pinkyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pinky"}

	pinkyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosine"}

	pinkyGermanName = nook.Name{
		Language: language.German,
		Value:    "Pia"}

	pinkyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pinky"}

	pinkyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タンタン"}

	pinkyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "링링"}

	pinkyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Violeta"}

	pinkyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пинки"}

	pinkySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丹丹"}

	pinkySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Violeta"}

	pinkyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "丹丹"}
)

var (
	pinkyName = nook.Languages{
		language.AmericanEnglish:      pinkyAmericanEnglishName,
		language.CanadianFrench:       pinkyCanadianFrenchName,
		language.Dutch:                pinkyDutchName,
		language.French:               pinkyFrenchName,
		language.German:               pinkyGermanName,
		language.Italian:              pinkyItalianName,
		language.Japanese:             pinkyJapaneseName,
		language.Korean:               pinkyKoreanName,
		language.LatinAmericanSpanish: pinkyLatinAmericanSpanishName,
		language.Russian:              pinkyRussianName,
		language.SimplifiedChinese:    pinkySimplifiedChineseName,
		language.Spanish:              pinkySpanishName,
		language.TraditionalChinese:   pinkyTraditionalChineseName}
)

var (
	pinkyCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: pinkyBirthday,
		Code:     pinkyCode,
		Key:      character.Pinky,
		Gender:   gender.Female,
		Name:     pinkyName}
)

var (
	pinkyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cubbie"}

	pinkyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ouah"}

	pinkyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wauw"}

	pinkyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ouah"}

	pinkyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schätzchen"}

	pinkyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cubetto"}

	pinkyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "わぉ"}

	pinkyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "와"}

	pinkyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guah"}

	pinkyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ну и ну"}

	pinkySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哇喔"}

	pinkySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "terronillo"}

	pinkyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哇喔"}
)

var (
	pinkyPhrase = nook.Languages{
		language.AmericanEnglish:      pinkyAmericanEnglishPhrase,
		language.CanadianFrench:       pinkyCanadianFrenchPhrase,
		language.Dutch:                pinkyDutchPhrase,
		language.French:               pinkyFrenchPhrase,
		language.German:               pinkyGermanPhrase,
		language.Italian:              pinkyItalianPhrase,
		language.Japanese:             pinkyJapanesePhrase,
		language.Korean:               pinkyKoreanPhrase,
		language.LatinAmericanSpanish: pinkyLatinAmericanSpanishPhrase,
		language.Russian:              pinkyRussianPhrase,
		language.SimplifiedChinese:    pinkySimplifiedChinesePhrase,
		language.Spanish:              pinkySpanishPhrase,
		language.TraditionalChinese:   pinkyTraditionalChinesePhrase}
)

var (
	Pinky = nook.Villager{
		Character:   pinkyCharacter,
		Personality: personality.Peppy,
		Phrase:      pinkyPhrase}
)
