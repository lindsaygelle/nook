package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	willowBirthday = nook.Birthday{
		Day:   26,
		Month: time.November}
)

var (
	willowCode = nook.Code{
		Value: "shp07"}
)

var (
	willowAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Willow"}

	willowCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maï"}

	willowDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Willow"}

	willowFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maï"}

	willowGermanName = nook.Name{
		Language: language.German,
		Value:    "Natascha"}

	willowItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Belinda"}

	willowJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マリー"}

	willowKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마리"}

	willowLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cuqui"}

	willowRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уиллоу"}

	willowSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "梅丽诺"}

	willowSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cuqui"}

	willowTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "梅麗諾"}
)

var (
	willowName = nook.Languages{
		language.AmericanEnglish:      willowAmericanEnglishName,
		language.CanadianFrench:       willowCanadianFrenchName,
		language.Dutch:                willowDutchName,
		language.French:               willowFrenchName,
		language.German:               willowGermanName,
		language.Italian:              willowItalianName,
		language.Japanese:             willowJapaneseName,
		language.Korean:               willowKoreanName,
		language.LatinAmericanSpanish: willowLatinAmericanSpanishName,
		language.Russian:              willowRussianName,
		language.SimplifiedChinese:    willowSimplifiedChineseName,
		language.Spanish:              willowSpanishName,
		language.TraditionalChinese:   willowTraditionalChineseName}
)

var (
	willowCharacter = nook.Character{
		Animal:   Sheep,
		Birthday: willowBirthday,
		Code:     willowCode,
		Gender:   nook.Female,
		Name:     willowName}
)

var (
	willowAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bo peep"}

	willowCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tourni"}

	willowDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wollebol"}

	willowFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tourni"}

	willowGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "blö-ö-ök"}

	willowItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "babeeebí"}

	willowJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですのよ"}

	willowKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "됐거든"}

	willowLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "babeeebí"}

	willowRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пасемся"}

	willowSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是的唷"}

	willowSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "babeeebí"}

	willowTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是的唷"}
)

var (
	willowPhrase = nook.Languages{
		language.AmericanEnglish:      willowAmericanEnglishName,
		language.CanadianFrench:       willowCanadianFrenchName,
		language.Dutch:                willowDutchName,
		language.French:               willowFrenchName,
		language.German:               willowGermanName,
		language.Italian:              willowItalianName,
		language.Japanese:             willowJapaneseName,
		language.Korean:               willowKoreanName,
		language.LatinAmericanSpanish: willowLatinAmericanSpanishName,
		language.Russian:              willowRussianName,
		language.SimplifiedChinese:    willowSimplifiedChineseName,
		language.Spanish:              willowSpanishName,
		language.TraditionalChinese:   willowTraditionalChineseName}
)

var (
	Willow = nook.Villager{
		Character:   willowCharacter,
		Personality: nook.Snooty,
		Phrase:      willowPhrase}
)
