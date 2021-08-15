package eagle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	averyBirthday = nook.Birthday{
		Day:   22,
		Month: time.February}
)

var (
	averyCode = nook.Code{
		Value: "pbr05"}
)

var (
	averyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Avery"}

	averyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Faustgrrraak"}

	averyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Averykra-hoe"}

	averyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Faustphélès"}

	averyGermanName = nook.Name{
		Language: language.German,
		Value:    "Quetzalflapp"}

	averyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Falcogawk"}

	averyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クスケチャアリョイ"}

	averyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쿠스케처내놔"}

	averyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cuzcogrrraak"}

	averyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аверикурлы"}

	averySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安谷斯溜溜"}

	averySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cuzcogrrraak"}

	averyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安谷斯溜溜"}
)

var (
	averyName = nook.Languages{
		language.AmericanEnglish:      averyAmericanEnglishName,
		language.CanadianFrench:       averyCanadianFrenchName,
		language.Dutch:                averyDutchName,
		language.French:               averyFrenchName,
		language.German:               averyGermanName,
		language.Italian:              averyItalianName,
		language.Japanese:             averyJapaneseName,
		language.Korean:               averyKoreanName,
		language.LatinAmericanSpanish: averyLatinAmericanSpanishName,
		language.Russian:              averyRussianName,
		language.SimplifiedChinese:    averySimplifiedChineseName,
		language.Spanish:              averySpanishName,
		language.TraditionalChinese:   averyTraditionalChineseName}
)

var (
	averyCharacter = nook.Character{
		Animal:   Eagle,
		Birthday: averyBirthday,
		Code:     averyCode,
		Gender:   nook.Male,
		Name:     averyName}
)

var (
	averyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "skree-haw"}

	averyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grrraak"}

	averyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kra-hoe"}

	averyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "phélès"}

	averyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flapp"}

	averyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gawk"}

	averyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アリョイ"}

	averyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "내놔"}

	averyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrraak"}

	averyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "курлы"}

	averySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "溜溜"}

	averySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "grrraak"}

	averyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "溜溜"}
)

var (
	averyPhrase = nook.Languages{
		language.AmericanEnglish:      averyAmericanEnglishName,
		language.CanadianFrench:       averyCanadianFrenchName,
		language.Dutch:                averyDutchName,
		language.French:               averyFrenchName,
		language.German:               averyGermanName,
		language.Italian:              averyItalianName,
		language.Japanese:             averyJapaneseName,
		language.Korean:               averyKoreanName,
		language.LatinAmericanSpanish: averyLatinAmericanSpanishName,
		language.Russian:              averyRussianName,
		language.SimplifiedChinese:    averySimplifiedChineseName,
		language.Spanish:              averySpanishName,
		language.TraditionalChinese:   averyTraditionalChineseName}
)

var (
	Avery = nook.Villager{
		Character:   averyCharacter,
		Personality: nook.Cranky,
		Phrase:      averyPhrase}
)
