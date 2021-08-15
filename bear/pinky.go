package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Rosineouah"}

	pinkyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pinkywauw"}

	pinkyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosineouah"}

	pinkyGermanName = nook.Name{
		Language: language.German,
		Value:    "Piaschätzchen"}

	pinkyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pinkycubetto"}

	pinkyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タンタンわぉ"}

	pinkyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "링링와"}

	pinkyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Violetaguah"}

	pinkyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пинкину и ну"}

	pinkySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丹丹哇喔"}

	pinkySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Violetaterronillo"}

	pinkyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "丹丹哇喔"}
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
		Animal:   Bear,
		Birthday: pinkyBirthday,
		Code:     pinkyCode,
		Gender:   nook.Female,
		Name:     pinkyName}
)

var (
	pinkyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cubbiewah"}

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
	Pinky = nook.Villager{
		Character:   pinkyCharacter,
		Personality: nook.Peppy,
		Phrase:      pinkyPhrase}
)
