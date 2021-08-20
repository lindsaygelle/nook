package eagle

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
		Value:    "Faust"}

	averyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Avery"}

	averyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Faust"}

	averyGermanName = nook.Name{
		Language: language.German,
		Value:    "Quetzal"}

	averyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Falco"}

	averyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クスケチャ"}

	averyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쿠스케처"}

	averyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cuzco"}

	averyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Авери"}

	averySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安谷斯"}

	averySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cuzco"}

	averyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安谷斯"}
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
		Animal:   animal.Eagle,
		Birthday: averyBirthday,
		Code:     averyCode,
		Key:      character.Avery,
		Gender:   gender.Male,
		Name:     averyName,
		Special:  false}
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
		language.AmericanEnglish:      averyAmericanEnglishPhrase,
		language.CanadianFrench:       averyCanadianFrenchPhrase,
		language.Dutch:                averyDutchPhrase,
		language.French:               averyFrenchPhrase,
		language.German:               averyGermanPhrase,
		language.Italian:              averyItalianPhrase,
		language.Japanese:             averyJapanesePhrase,
		language.Korean:               averyKoreanPhrase,
		language.LatinAmericanSpanish: averyLatinAmericanSpanishPhrase,
		language.Russian:              averyRussianPhrase,
		language.SimplifiedChinese:    averySimplifiedChinesePhrase,
		language.Spanish:              averySpanishPhrase,
		language.TraditionalChinese:   averyTraditionalChinesePhrase}
)

var (
	Avery = nook.Villager{
		Character:   averyCharacter,
		Personality: personality.Cranky,
		Phrase:      averyPhrase}
)
